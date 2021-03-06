package utils

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"mime/multipart"
	"os"
	"panta/global"
	"strings"
	"time"
)

// 接收两个参数 一个文件流 一个 bucket 你的七牛云标准空间的名字
func Upload(file *multipart.FileHeader) (err error, path string, key string) {
	putPolicy := storage.PutPolicy{
		Scope: global.PantaConfig.Qiniu.Bucket,
	}
	mac := qbox.NewMac(global.PantaConfig.Qiniu.AccessKey, global.PantaConfig.Qiniu.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	f, e := file.Open()
	if e != nil {
		fmt.Println(e)
		return e, "", ""
	}
	dataLen := file.Size
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	err = formUploader.Put(context.Background(), &ret, upToken, fileKey, f, dataLen, &putExtra)
	if err != nil {
		global.PantaLog.Error("upload file fail:", err)
		return err, "", ""
	}
	return err, global.PantaConfig.Qiniu.ImgPath + "/" + ret.Key, ret.Key
}

func DeleteFileInQiniu(key string) error {
	mac := qbox.NewMac(global.PantaConfig.Qiniu.AccessKey, global.PantaConfig.Qiniu.SecretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	// cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(global.PantaConfig.Qiniu.Bucket, key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DeleteFileInDisk(url string) (err error) {
	root, _ := os.Getwd()
	prefix := root + "/resource"
	splits := strings.Split(url, "/upload")
	if len(splits) < 2 {
		return fmt.Errorf("url: %s  illegal", url)
	} else {
		path := prefix + "/upload" + splits[1]
		if exist, _ := PathExists(path); exist {
			return os.Remove(path)
		} else {
			return fmt.Errorf("path: %s  not exist", path)
		}
	}
}