<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
      <el-table-column type="selection" min-width="20"></el-table-column>
      <el-table-column label="序号" prop="ID" sortable min-width="30" align="center"></el-table-column> 
      <el-table-column label="电梯" prop="lift.nickName" sortable min-width="50" align="center"></el-table-column> 
      <el-table-column label="记录类别" prop="category.categoryName" sortable min-width="60" align="center"></el-table-column> 
      <el-table-column label="开始时间" sortable min-width="80" align="center">
        <template slot-scope="scope">{{scope.row.startTime|formatDate}}</template>
      </el-table-column> 

      <el-table-column label="图片记录" prop="medias" sortable min-width="120" align="center">
        <template slot-scope="scope">
          <viewer v-if="scope.row.medias" :images="previewImages" @inited="inited" ref="viewer" class="images clearfix">
            <el-carousel height="150px" indicator-position="none">
              <el-carousel-item v-for="item in scope.row.medias" :key="item.url">
                <el-image :src="item.url" fit="cover" @click="handlePreview(scope.row.medias)" style="width:200px;height:150px"/>
              </el-carousel-item>
            </el-carousel>            
          </viewer>          
          <span v-else>---</span>
        </template>     
      </el-table-column> 
      <el-table-column label="文字记录" prop="content" sortable min-width="80" align="center"></el-table-column> 
      <el-table-column label="记录人员" prop="recorder.realName" sortable min-width="60" align="center"></el-table-column> 
      <el-table-column label="结束时间" prop="endTime" sortable min-width="80" align="center">
        <template slot-scope="scope">{{scope.row.endTime|formatDate}}</template>        
      </el-table-column> 
      <el-table-column label="按钮组" fixed="right" min-width="100" >
        <template slot-scope="scope">
          <el-button @click="process(scope.row)" size="small" :type="getButtonType(scope.row)">{{ scope.row|formatTitle }}</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteLiftRecord(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-steps :active="stepActive" align-center>
        <el-step v-for="item in stepItems"
          :key="item.key"
          :title="item.title"
          :description="item.description"
          :icon="item.icon">
        </el-step>
      </el-steps>
      <el-form class="form-data" :model="formData" ref="formData" :label-position="labelPosition" label-width="100px" size="medium" label-postion="left">
        <!-- created -->
        <el-form-item v-if="stepActive == 0" label="电梯:" prop="liftId">
          <el-select v-model="formData.liftId" filter clearable placeholder="请选择电梯" >
            <el-option v-for="item in liftOptions"
              :key="item.ID"
              :label="item.nickName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive == 0" label="记录类型:" prop="categoryId">
          <el-select v-model="formData.categoryId" filter clearable placeholder="请选择记录类型">
            <el-option v-for="item in categoryOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="电梯:" prop="liftId">
          <span>{{formData.lift.nickName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="记录类型:" prop="categoryId">
          <span>{{formData.category.categoryName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="创建时间:" prop="CreatedAt">
          <span >{{formData.CreatedAt|formatDate}}</span>
        </el-form-item>
        <!-- started -->
        <el-form-item v-if="stepActive == 1" label="人  员:" prop="workerId">
          <el-select v-model="formData.workerId" filter clearable placeholder="请选择人员">
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 1" label="人  员:" prop="workerId">
          <span>{{formData.worker.realName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 1" label="开始时间:" prop="startTime">
          <span >{{formData.startTime|formatDate}}</span>
        </el-form-item>
        <!-- ended -->
        <el-form-item v-if="stepActive == 2" label="内  容:" prop="content">
          <el-input v-model="formData.content" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder=""></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive > 2" label="内  容:" prop="content">
          <el-input v-model="formData.content" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder="please input record content"></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive == 2" label="媒体文件" prop="medias">
          <el-upload
            :action="`${path}/fileUploadAndDownload/upload?storage=local`"
            :on-remove="handleRemove"
            :on-success="handleUploadSuccess"
            multiple
            :headers="{ 'x-token': token }"
            :file-list="uploadFileList"
            list-type="picture-card">
            <i class="el-icon-plus"></i>
          </el-upload>
        </el-form-item>
        <el-form-item v-if="stepActive > 2" label="媒体文件">
          <el-carousel height="150px">
            <el-carousel-item v-for="item in uploadFileList" :key="item.uid">
              <img :src="item.url" />
            </el-carousel-item>
          </el-carousel>   
        </el-form-item>
        <!-- reviewed -->
        <el-form-item v-if="stepActive == 3" label="审核人员" prop="recorderId">
          <el-select v-model="formData.recorderId" filter clearable placeholder="请选择审核人员">
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="完结时间" prop="endTime">
          <span>{{formData.endTime|formatDate}}</span>
        </el-form-item> 
        <el-form-item v-if="stepActive > 3" label="审核人员" prop="recorderId">
          <span>{{formData.recorder.realName}}</span>
        </el-form-item> 
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
const path = process.env.VUE_APP_BASE_API;
import {
    createLiftRecord,
    deleteLiftRecord,
    deleteLiftRecordByIds,
    updateLiftRecord,
    findLiftRecord,
    getLiftRecordList
} from "@/api/liftRecord";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { getUserList } from '@/api/user';
import { getCategoriesList } from '@/api/categories';
import { getLiftList } from "@/api/lift";
import { mapGetters } from "vuex";
import Subject from "@/utils/subject"

export default {
  name: "LiftRecord",
  mixins: [infoList],
  data() {
    return {
      listApi: getLiftRecordList,
      dialogFormVisible: false,
      previewDialogVisible: false,
      previewImages:[],
      deleteVisible: false,
      labelPosition:"right",
      path: path,
      uploadFileList: [],
      stepItems: [
        {key: 1, title: "1. 创建", description: "创建记录", icon:"el-icon-edit"},
        {key: 2, title: "2. 开始", description: "开始记录", icon:"el-icon-upload"},
        {key: 3, title: "3. 完结", description: "提交记录", icon:"el-icon-s-promotion"},
        {key: 4, title: "4. 审核", description: "审核记录", icon:"el-icon-check"}
      ],
      categoryOptions:[],
      userOptions:[],
      liftOptions:[],
      multipleSelection: [],
      formData: {
        liftId:null,lift:{}, categoryId:null, category:{},
        medias:"",content:"",startTime:"0001-01-01T00:00:00Z",endTime:"0001-01-01T00:00:00Z",
        workerId:null,worker:{}, recorderId:0, recorder:{}, progress:0
      }
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"]),
    stepActive: function() {
      return this.formData.progress
    }
  },
  filters: {
    formatDate: function(time) {
      if (time == "0001-01-01T00:00:00Z") {
        return "---"
      }
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    },
    formatTitle: function(row) {
      switch(row.progress) {
        case 1:
          return "待开始"
        case 2:
          return "待完成"
        case 3:
          return "待审核"
        case 4:
          return "已完成"
      }
    }
  },
  methods: {
      onSubmit() {
        this.page = 1
        this.pageSize = 10            
        this.getTableData()
      },
      inited(viewer) {
        this.$viewer = viewer
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      getButtonType(row) {
        if (row.recorderId == 0) {
          return "warning"
        } else return "success"
      },
      handleRemove(file, fileList){
        console.log(fileList)        
        this._.remove(this.uploadFileList, function(item) {
          return item.uid == file.uid  
        })
      },
      handlePreview(medias) {
        this.previewImages = []
        this.previewImages = this._.map(this._.filter(medias, function(o){return o.tag == "jpg"}), "url")
        this.$viewer.show()
      },
      handleUploadSuccess(response, file, fileList) {
        this.uploadFileList.push({name: file.name, url: response.data.file.url})
        console.log(fileList)
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteLiftRecordByIds({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async process(row) {
      const res = await findLiftRecord({ ID: row.ID });
      if (res.code === 0) {
        this.formData = res.data.liftRecord;
        this.uploadFileList = this._.filter(this.formData.medias, ['tag', 'jpg'])
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        liftId:null,lift:{}, categoryId:null, category:{},
        medias:"",content:"",startTime:"0001-01-01T00:00:00Z",endTime:"0001-01-01T00:00:00Z",
        workerId:null,worker:{}, recorderId:0, recorder:{}
      };
    },
    async deleteLiftRecord(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteLiftRecord({ ID: row.ID });
        if (res.code === 0) {
          this.$message({
            type: "success",
            message: "删除成功"
          });
          await this.getTableData();
        }
      })
      .catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    async enterDialog() {
      let res;
      let param = {}
      switch (this.stepActive) {
        case 0:
          param.liftId = this.formData.liftId
          param.categoryId = this.formData.categoryId
          res = await createLiftRecord(param);
          break;
        case 1:
          param.recordId = this.formData.ID
          param.workerId = this.formData.workerId
          res = await updateLiftRecord(param);
          break;
        case 2:
          param.recordId = this.formData.ID
          param.content = this.formData.content
          param.medias = JSON.stringify(this.uploadFileList)
          res = await updateLiftRecord(param);
          break;
        case 3:
          param.recordId = this.formData.ID
          param.recorderId = this.formData.recorderId
          res = await updateLiftRecord(param)
          break;
      }
      if (res.code === 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        await this.getTableData();
      }
    },
    openDialog() {
      this.dialogFormVisible = true;
    },
    async getUserOptions() {
      let res = await getUserList({
        page: 1,
        pageSize: 9999
      })
      if (res.code == 0) {
        this.userOptions = res.data.list
      }
    },
    async getLiftOptions() {
      let res = await getLiftList({
        page: 1,
        pageSize: 9999
      })
      if (res.code == 0) {
        this.liftOptions = res.data.list
      }
    },
    async getCategoryOptions() {
      let res = await getCategoriesList({
        ID: Subject.SubjectLiftRecordType,
        page: 1,
        pageSize: 9999
      })
      if (res.code == 0) {
        this.categoryOptions = res.data.list
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getUserOptions()  
    await this.getCategoryOptions()
    await this.getLiftOptions()
  }
};
</script>

<style lang="less">
.previewImage {
    cursor: pointer;
    margin: 5px;
    display: inline-block;
}

.form-data {
  margin-top: 36px;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  .el-form-item {
    width: 60%;
    margin-bottom: 8px;
    .el-form-item__label {
      display: flex;
      align-content: center;
      width: 20%;
      padding: 0;
    }
    .el-form-item__content {
      width: 60%;
      .el-select {
        width: 100%;
      }
    }
  }
}

</style>