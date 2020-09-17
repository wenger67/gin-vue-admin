package utils

import (
	"gin-vue-admin/global"
	"os"
	"os/exec"
	"strings"
)

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @title    createDir
// @description   批量创建文件夹
// @auth                     （2020/04/05  20:22）
// @param     dirs            string
// @return    err             error

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.PantaLog.Debug("create directory ", v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				global.PantaLog.Error("create directory", v, " error:", err)
			}
		}
	}
	return err
}


 func RootPath() string {
     s, err := exec.LookPath(os.Args[0])
     if err != nil {
         global.PantaLog.Error("发生错误",err.Error())
 	 }
     i := strings.LastIndex(s, "\\")
     path := s[0 : i+1]
     return path
 }