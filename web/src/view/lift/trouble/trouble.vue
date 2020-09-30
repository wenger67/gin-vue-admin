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
        <el-form-item class="export">
          <el-button type="primary" @click="exportXlsx">Export</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      tooltip-effect="dark"
    >
      <el-table-column type="selection" min-width="30"></el-table-column>
      <el-table-column label="序号" prop="ID" sortable min-width="80" align="center"></el-table-column> 
      <el-table-column label="创建" sortable min-width="120" align="center">
        <template slot-scope="scope">
          <el-popover
            placement="top"
            trigger="click">
            <p>来源类别: {{ scope.row.fromCategory.categoryName }}</p>
            <p>开始时间: {{ scope.row.startTime | formatDate }}</p>
            <p>发起人员: {{ scope.row.startUser.realName }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.lift.nickName}}</el-tag>
            </div>
          </el-popover>
        </template>
      </el-table-column> 
      <el-table-column label="响应" sortable min-width="100" align="center">
        <template slot-scope="scope">
          <el-popover v-if="scope.row.responseUserId"
            placement="top"
            trigger="click">
            <p>响应时间: {{ scope.row.responseTime | formatDate }}</p>
            <p>响应人员: {{ scope.row.responseUser.realName }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.responseUser.realName }}</el-tag>
            </div>
          </el-popover>
          <span class="no-data" v-else >无记录</span>
        </template>
      </el-table-column>   
          
      <el-table-column label="现场" sortable min-width="100" align="center">
        <template slot-scope="scope">
          <el-popover v-if="scope.row.sceneUserId"
            placement="top"
            trigger="click">
            <p>现场时间: {{ scope.row.sceneTime | formatDate }}</p>
            <p>现场人员: {{ scope.row.sceneUser.realName }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.sceneUser.realName }}</el-tag>
            </div>
          </el-popover>
          <span class="no-data" v-else >无记录</span>
        </template>
      </el-table-column> 
     
      <el-table-column label="解除" sortable min-width="100" align="center">
        <template slot-scope="scope">
          <el-popover v-if="scope.row.fixUserId"
            placement="top"
            trigger="click">
            <p>解除时间: {{ scope.row.fixTime | formatDate }}</p>
            <p>解除故障人员: {{ scope.row.fixUser.realName }}</p>
            <p>解除故障方式: {{ scope.row.fixCategory.categoryName }}</p>
            <p>故障原因: {{ scope.row.reasonCategory.categoryName }}</p>
            <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.fixCategory.categoryName }}</el-tag>
            </div>
          </el-popover>
          <span class="no-data" v-else >无记录</span>
        </template>
      </el-table-column> 

      <el-table-column label="故障详情" sortable min-width="200" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.content && scope.row.content != `null`">{{ scope.row.content }}</span>
          <span class="no-data" v-else >无记录</span>
        </template>
      </el-table-column>  -->
      <el-table-column label="图片记录" prop="medias" sortable min-width="200" align="center">
        <template slot-scope="scope">
          <viewer v-if="scope.row.medias.length" :images="previewImages" @inited="inited" ref="viewer" class="images clearfix">
            <el-carousel height="150px" indicator-position="none">
              <el-carousel-item v-for="item in scope.row.medias" :key="item.url">
                <el-image :src="item.url" fit="cover" @click="handlePreview(scope.row.medias)" style="width:200px;height:150px"/>
              </el-carousel-item>
            </el-carousel>            
          </viewer>          
          <span class="no-data" v-else >无记录</span>
        </template>     
      </el-table-column>     
      <el-table-column label="记录人员" prop="recorder.realName" sortable min-width="120" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.recorderId">{{ scope.row.recorder.realName }}</span>
          <span class="no-data" v-else >无记录</span>
        </template>
      </el-table-column> 
      <el-table-column label="反馈内容" sortable min-width="150" prop="feedbackContent" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.feedbackContent">{{ scope.row.feedbackContent }}</span>
          <span class="no-data" v-else >无记录</span>
        </template>
      </el-table-column> 
      <el-table-column label="评分" prop="feedbackRate" sortable min-width="180" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.feedbackRate">
            <el-rate v-model="scope.row.feedbackRate" disabled show-score :allow-half="allowHalf"
           :max="max" :colors="['#99A9BF', '#F7BA2A', '#FF9900']"></el-rate>
          </span>
          <span class="no-data" v-else >无记录</span>
        </template>
      </el-table-column> 
      <el-table-column label="日期" min-width="160" align="center">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" min-width="200" align="center">
        <template slot-scope="scope">
          <el-button @click="process(scope.row)" size="small" :type="getButtonType(scope.row)">{{ scope.row.progress|formatProgress }}</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteLiftTrouble(scope.row)">删除</el-button>
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
      <el-form :model="formData" ref="formData" class="form-data" label-width="100px" size="medium" label-postion="left" :style="{marginTop:'30px'}">
        <!-- created -->
        <el-form-item v-if="stepActive == 0" label="故障类型" prop="fromCategoryId">
          <el-select v-model="formData.fromCategoryId" filter clearable placeholder="请选择故障类型" >
            <el-option v-for="item in fromOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive == 0" label="电梯:" prop="liftId">
          <el-select v-model="formData.liftId" filter clearable placeholder="请选择电梯" >
            <el-option v-for="item in liftOptions"
              :key="item.ID"
              :label="item.nickName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive == 0" label="创建者" prop="startUserId">
          <el-select v-model="formData.startUserId" filter clearable placeholder="请选择创建者" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="故障类型" prop="fromCategoryId">
          <span>{{formData.fromCategory.categoryName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="电梯:" prop="liftId">
          <span>{{formData.lift.nickName}}</span>
        </el-form-item>        
        <el-form-item v-if="stepActive > 0" label="创建者" prop="startUserId">
          <span>{{formData.startUser.realName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="创建时间" prop="startTime">
          <span>{{formData.startTime|formatDate}}</span>
        </el-form-item>
        <!-- responsed -->
        <el-form-item v-if="stepActive == 1" label="响应人员" prop="responseUserId">
          <el-select v-model="formData.responseUserId" filter clearable placeholder="请选择响应人员" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 1" label="响应人员" prop="responseUserId">
          <span>{{formData.responseUser.realName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 1" label="响应时间" prop="responseTime">
          <span>{{formData.responseTime|formatDate}}</span>
        </el-form-item>
        <!-- scene -->
        <el-form-item v-if="stepActive == 2" label="到场人员" prop="sceneUserId">
          <el-select v-model="formData.sceneUserId" filter clearable placeholder="请选择到场人员" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 2" label="到场人员" prop="sceneUserId">
          <span>{{formData.sceneUser.realName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 2" label="到场时间" prop="sceneTime">
          <span>{{formData.sceneTime|formatDate}}</span>
        </el-form-item>
        <!-- fixed -->
        <el-form-item v-if="stepActive == 3" label="修复人员" prop="fixUserId">
          <el-select v-model="formData.fixUserId" filter clearable placeholder="请选择修复人员" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive == 3" label="修复方式" prop="fixCategoryId">
          <el-select v-model="formData.fixCategoryId" filter clearable placeholder="请选择修复方式" >
            <el-option v-for="item in fixOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive == 3" label="故障原因" prop="reasonCategoryId">
          <el-select v-model="formData.reasonCategoryId" filter clearable placeholder="请选择故障原因" >
            <el-option v-for="item in reasonOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="修复人员" prop="fixUserId">
          <span>{{formData.fixUser.realName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="修复方式" prop="fixCategoryId">
          <span>{{formData.fixCategory.categoryName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="故障原因" prop="reasonCategoryId">
          <span>{{formData.reasonCategory.categoryName}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="修复时间" prop="fixTime">
          <span>{{formData.fixTime|formatDate}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive == 3" label="内容:" prop="content">
          <el-input v-model="formData.content" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder="请输入内容"></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="内容:" prop="content">
          <span>{{formData.content}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive == 3" label="媒体文件" prop="medias">
          <el-upload
            :action="`${path}/fileUploadAndDownload/uploadList`"
            :on-remove="handleRemove"
            :on-success="handleUploadSuccess"
            multiple
            :headers="{ 'x-token': token }"
            :file-list="uploadFileList"
            list-type="picture-card"
            :data="uploadData">
            <i class="el-icon-plus"></i>
          </el-upload>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="媒体文件">
          <el-carousel height="150px">
            <el-carousel-item v-for="item in uploadFileList" :key="item.uid">
              <img :src="item.url" />
            </el-carousel-item>
          </el-carousel>   
        </el-form-item>
        <!-- feedback -->
        <el-form-item v-if="stepActive == 4" label="反馈内容" prop="feedbackContent">
          <el-input v-model="formData.feedbackContent" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder="请输入反馈内容"></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive > 4" label="反馈内容" prop="feedbackContent">
          <span>{{formData.feedbackContent}}</span>
        </el-form-item>
        <el-form-item v-if="stepActive == 4" label="反馈评分" prop="feedbackRate">
          <el-rate v-model="formData.feedbackRate" :allow-half="allowHalf" :max="max"
          :colors="['#99A9BF', '#F7BA2A', '#FF9900']"></el-rate>
        </el-form-item>
        <el-form-item v-if="stepActive > 4" label="反馈评分" prop="feedbackRate">
          <el-rate v-model="formData.feedbackRate" disabled show-score
           :max="max" :colors="['#99A9BF', '#F7BA2A', '#FF9900']"></el-rate>
        </el-form-item>
        <!-- reviewed -->
        <el-form-item v-if="stepActive == 5" label="审核人员" prop="recorderId">
          <el-select v-model="formData.recorderId" filter clearable placeholder="请选择审核人员" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 5" label="审核人员" prop="recorderId">
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
    createLiftTrouble,
    deleteLiftTrouble,
    deleteLiftTroubleByIds,
    updateLiftTrouble,
    findLiftTrouble,
    getLiftTroubleList
} from "@/api/liftTrouble";  //  此处请自行替换地址
import { deleteFile } from "@/api/fileUploadAndDownload";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { getUserList } from '@/api/user';
import { getCategoriesList } from '@/api/categories';
import { getLiftList } from "@/api/lift";
import { mapGetters } from "vuex";
import Subject from "@/utils/subject";
import XLSX from "xlsx";

export default {
  name: "LiftTrouble",
  mixins: [infoList],
  data() {
    return {
      listApi: getLiftTroubleList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      previewImages:[],
      path: path,
      uploadFileList: [],
      uploadData:{
        storage: "local",
        type: "trouble"
      },
      stepItems: [
        {key: 1, title: "1. 创建", description: "创建故障记录", icon:"el-icon-edit"},
        {key: 2, title: "2. 响应", description: "响应故障", icon:"el-icon-upload"},
        {key: 3, title: "3. 到场", description: "到场时间和人员", icon:"el-icon-s-promotion"},
        {key: 4, title: "4. 修复", description: "修复记录", icon:"el-icon-magic-stick"},
        {key: 5, title: "5. 反馈", description: "反馈记录", icon:"el-icon-s-comment"},
        {key: 6, title: "6. 审核", description: "审核记录", icon:"el-icon-check"}
      ],
      liftOptions:[],
      fromOptions:[],
      userOptions:[],
      fixOptions:[],
      reasonOptions:[],
      multipleSelection: [],formData: {
        liftId:null,lift:{}, fromCategoryId:null,fromCategory:{}, startTime:"0001-01-01T00:00:00Z",
        startUserId:null,startUser:{},responseTime:"0001-01-01T00:00:00Z",responseUserId:null,
        responseUser:{}, sceneTime:"0001-01-01T00:00:00Z",sceneUserId:null,sceneUser:{},
        fixTime:"0001-01-01T00:00:00Z",fixUserId:null,fixUser:{}, fixCategoryId:null,
        fixCategory:{},reasonCategoryId:null,reasonCategory:{}, content:null,progress:0,
        recorderId:null,recorder:{},feedbackContent:null,feedbackRate:null
      },
      allowHalf:true,
      max: 5
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
      if (time != null && time != "" && time != "0001-01-01T00:00:00Z") {
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
    formatProgress: function(progress) {
      switch(progress) {
        case 1:
          return "待响应"
        case 2:
          return "待到达"
        case 3:
          return "待修复"
        case 4:
          return "待反馈"
        case 5:
          return "待审核"
        case 6:
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
    async handleRemove(file){
      const item = this._.find(this.uploadFileList, ["uid", file.uid])  
      const res = await deleteFile({
        ID: item.ID,
        storage: "local"
      })
      if (res.code == 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })       
        this._.remove(this.uploadFileList, function(n){
          return n.uid == file.uid
        })   
      } else {
        this.$message({
          type: 'error',
          message: '删除失败:' + res.message
        })     
      }
    },
    handlePreview(medias) {
      this.previewImages = []
      this.previewImages = this._.map(this._.filter(medias, function(o){return o.tag == "jpg"}), "url")
      this.$viewer.show()
    },
    handleUploadSuccess(response) {
      this.uploadFileList.push(response.data.files[0])
    },
    exportXlsx() {
      var table = [["ID", "Lift", "ResponseUser"]] // xlsx header
      // xlsx data
      this.tableData.map((item, index) => {
        table[index + 1] = [item.ID, item.lift.nickName, item.responseUser.realName]
      })

      // format convert
      const ws = XLSX.utils.aoa_to_sheet(table)
      const wb = XLSX.utils.book_new()
      XLSX.utils.book_append_sheet(wb, ws, "sheet1");
      // save files
      XLSX.writeFile(wb, "lift_trouble_sheet.xlsx")
    },
    async onDelete() {
      const ids = []
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteLiftTroubleByIds({ ids })
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
      const res = await findLiftTrouble({ ID: row.ID });
      if (res.code === 0) {
        this.formData = res.data.reliftTrouble;
        
        switch (this.formData.progress) {
          case 1:
            this.formData.responseUserId = this.userOptions[0].ID
            break;
          case 2:
            this.formData.sceneUserId = this.userOptions[0].ID
            break;
          case 3:
            this.formData.fixUserId = this.userOptions[0].ID
            this.formData.fixCategoryId = this.fixOptions[0].ID
            this.formData.reasonCategoryId = this.reasonOptions[0].ID
            break;
          case 5:
            this.formData.recorderId = this.userOptions[0].ID
            break;
        }
        console.log(this.formData.medias)
        this.uploadFileList = this.formData.medias
        this.dialogFormVisible = true;
        // assign id for upload file field
        this.uploadData.id = row.ID
        this.uploadData.liftId = res.data.reliftTrouble.liftId
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        liftId:null,lift:{}, fromCategoryId:null,fromCategory:{}, startTime:"0001-01-01T00:00:00Z",
        startUserId:null,startUser:{},responseTime:"0001-01-01T00:00:00Z",responseUserId:null,
        responseUser:{}, sceneTime:"0001-01-01T00:00:00Z",sceneUserId:null,sceneUser:{},
        fixTime:"0001-01-01T00:00:00Z",fixUserId:null,fixUser:{}, fixCategoryId:null,
        fixCategory:{},reasonCategoryId:null,reasonCategory:{}, content:null,progress:0,
        recorderId:null,recorder:{},feedbackContent:null,feedbackRate:null
      };
    },
    async deleteLiftTrouble(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteLiftTrouble({ ID: row.ID });
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
      this.type = this.stepActive == 0 ? "create":"update"

      let res;
      switch (this.type) {
        case "create":
          res = await createLiftTrouble(this.formData);
          break;
        case "update":
          res = await updateLiftTrouble(this.formData);
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
      this.type = "create";
      this.formData.startUserId = this.userOptions[0].ID
      this.formData.liftId = this.liftOptions[0].ID
      this.formData.fromCategoryId = this.fromOptions[0].ID

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
    async getFromOptions(){
      let res = await getCategoriesList({
        ID: Subject.SubjectLiftTroubleSourceType,
        page: 1,
        pageSize: 9999
      })
      if (res.code == 0) {
        this.fromOptions = res.data.list
      }
    },
    async getFixOptions(){
      let res = await getCategoriesList({
        ID: Subject.SubjectLiftTroubleSolvedType,
        page: 1,
        pageSize: 9999
      })
      if (res.code == 0) {
        this.fixOptions = res.data.list
      }
    },
    async getReasonOptions(){
      let res = await getCategoriesList({
        ID: Subject.SubjectLiftTroubleReasonType,
        page: 1,
        pageSize: 9999
      })
      if (res.code == 0) {
        this.reasonOptions = res.data.list
      }
    },
  },
  async created() {
    await this.getTableData()
    await this.getLiftOptions()
    await this.getUserOptions()
    await this.getFromOptions()
    await this.getFixOptions()
    await this.getReasonOptions()
  }
};
</script>

<style lang="less">

.export {
  float: right;
}

.el-rate {
  display: inline;
}

.no-data {
  color: lightgray;
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