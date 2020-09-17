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
        <el-form-item>
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
      <el-table-column label="ID" prop="ID" sortable min-width="60"></el-table-column> 
      <el-table-column label="故障进度" sortable min-width="120">
        <template slot-scope="scope">{{ scope.row.progress|formatProgress }}</template>
      </el-table-column>
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
      <el-table-column label="响应" sortable min-width="120" align="center">
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
          <span v-else>---</span>
        </template>
      </el-table-column>   
          
      <el-table-column label="现场" sortable min-width="120" align="center">
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
          <span v-else>---</span>
        </template>
      </el-table-column> 
     
      <el-table-column label="解除" sortable min-width="120" align="center">
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
          <span v-else>---</span>
        </template>
      </el-table-column> 

      <el-table-column label="故障详情" sortable min-width="200" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.content && scope.row.content != `null`">{{ scope.row.content }}</span>
          <span v-else>---</span>
        </template>
      </el-table-column>  -->
      <el-table-column label="图片记录" prop="medias" sortable min-width="200" align="center">
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
      <el-table-column label="记录人员" prop="recorder.realName" sortable min-width="120" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.recorderId">{{ scope.row.recorder.realName }}</span>
          <span v-else>---</span>
        </template>
      </el-table-column> 
      <el-table-column label="反馈内容" sortable min-width="150" prop="feedbackContent" align="center">
        <template slot-scope="scope">
          <span v-if="scope.row.feedbackContent">{{ scope.row.feedbackContent }}</span>
          <span v-else>---</span>
        </template>
      </el-table-column> 
      <el-table-column label="评分" prop="feedbackRate" sortable min-width="80">
        <template slot-scope="scope">
          <span v-if="scope.row.feedbackRate">{{ scope.row.feedbackRate }}</span>
          <span v-else>---</span>
        </template>
      </el-table-column> 
      <el-table-column label="日期" min-width="160">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" min-width="150">
        <template slot-scope="scope">
          <el-button @click="updateLiftTrouble(scope.row)" size="small" :type="getButtonType(scope.row)">{{ scope.row|formatTitle }}</el-button>
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
       <el-steps :active="stepActive" :style="{marginLeft:'30px', marginRight:'30px'}">
        <el-step v-for="item in stepItems"
          :key="item.key"
          :title="item.title"
          :description="item.description"
          :icon="item.icon">
        </el-step>
      </el-steps>
      <el-form :model="formData" ref="formData" label-width="100px" size="medium" label-postion="left" :style="{marginTop:'30px'}">
        <el-form-item v-if="stepActive == 0" label="lift id" prop="liftId">
          <el-select v-model="formData.liftId" filter clearable placeholder="please select lift" >
            <el-option v-for="item in liftOptions"
              :key="item.ID"
              :label="item.nickName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="lift id" prop="liftId">
          <el-input v-model="formData.lift.nickName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 0" label="from category" prop="fromCategoryId">
          <el-select v-model="formData.fromCategoryId" filter clearable placeholder="please select fromCategoryId" >
            <el-option v-for="item in fromOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="from category" prop="fromCategoryId">
          <el-input v-model="formData.fromCategory.categoryName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 0" label="creator" prop="startUserId">
          <el-select v-model="formData.startUserId" filter clearable placeholder="please select startUserId" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 0" label="creator" prop="startUserId">
          <el-input v-model="formData.startUser.realName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 1" label="response" prop="responseUserId">
          <el-select v-model="formData.responseUserId" filter clearable placeholder="please select responseUserId" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 1" label="response" prop="responseUserId">
          <el-input v-model="formData.responseUser.realName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 2" label="scene" prop="sceneUserId">
          <el-select v-model="formData.sceneUserId" filter clearable placeholder="please select sceneUserId" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 2" label="scene" prop="sceneUserId">
          <el-input v-model="formData.sceneUser.realName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 3" label="fixUser" prop="fixUserId">
          <el-select v-model="formData.fixUserId" filter clearable placeholder="please select fixUserId" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="fixUser" prop="fixUserId">
          <el-input v-model="formData.fixUser.realName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 3" label="fixCategory" prop="fixCategoryId">
          <el-select v-model="formData.fixCategoryId" filter clearable placeholder="please select fixCategoryId" >
            <el-option v-for="item in fixOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="fixCategory" prop="fixCategoryId">
          <el-input v-model="formData.fixCategory.categoryName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 3" label="reasonCategory" prop="reasonCategoryId">
          <el-select v-model="formData.reasonCategoryId" filter clearable placeholder="please select reasonCategoryId" >
            <el-option v-for="item in reasonOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="reasonCategory" prop="reasonCategoryId">
          <el-input v-model="formData.reasonCategory.categoryName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 3" label="content" prop="content">
          <el-input v-model="formData.content" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder="please input trouble content"></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive > 3" label="content" prop="content">
          <el-input v-model="formData.content" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 4" label="feedbackContent" prop="feedbackContent">
          <el-input v-model="formData.feedbackContent" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder="please input feedbackContent"></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive > 4" label="feedbackContent" prop="feedbackContent">
          <el-input v-model="formData.feedbackContent" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 4" label="feedbackRate" prop="feedbackRate">
           <el-input-number v-model="formData.feedbackRate" :min="1" :max="100" placeholder="please input feedbackRate"></el-input-number>
        </el-form-item>
        <el-form-item v-if="stepActive > 4" label="feedbackRate" prop="feedbackRate">
           <el-input-number v-model="formData.feedbackRate" :min="1" :max="100" placeholder="please input feedbackRate"></el-input-number>
        </el-form-item>

        <el-form-item v-if="stepActive == 5" label="recorder" prop="recorderId">
          <el-select v-model="formData.recorderId" filter clearable placeholder="please select recorderId" >
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive > 5" label="recorder" prop="recorderId">
          <el-input v-model="formData.recorder.realName" disabled placeholder=""></el-input>
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
import {
    createLiftTrouble,
    deleteLiftTrouble,
    deleteLiftTroubleByIds,
    updateLiftTrouble,
    findLiftTrouble,
    getLiftTroubleList
} from "@/api/liftTrouble";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { getUserList } from '@/api/user';
import { getCategoriesList } from '@/api/categories';
import { getLiftList } from "@/api/lift";
import Subject from "@/utils/subject";
import XLSX from "xlsx";
import FileSaver from "file-saver"

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
      stepItems: [
        {key: 1, title: "步骤 1", description: "create trouble", icon:"el-icon-edit"},
        {key: 2, title: "步骤 2", description: "response", icon:"el-icon-alarm-clock"},
        {key: 3, title: "步骤 3", description: "scene", icon:"el-icon-location-information"},
        {key: 4, title: "步骤 4", description: "fix", icon:"el-icon-s-tools"},
        {key: 5, title: "步骤 5", description: "feedback", icon:"el-icon-s-comment"},
        {key: 6, title: "步骤 6", description: "review", icon:"el-icon-s-promotion"}
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
      }
    };
  },
  computed: {
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
    formatTitle: function(row) {
      if (row.recorderId == 0) {
        return "InProcess"
      } else {
        return "View"
      }
    },
    formatProgress: function(progress) {
      switch(progress) {
        case 1:
          return "created, need response"
        case 2:
          return "responsed, need scene"
        case 3:
          return "scened, need fix"
        case 4:
          return "fixed, neesd feedback"
        case 5:
          return "feedbacked, need review"
        case 6:
          return "reviewed"
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
    handlePreview(medias) {
      this.previewImages = []
      this.previewImages = _.map(_.filter(medias, function(o){return o.tag == "jpg"}), "url")
      this.$viewer.show()
    },    
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    getButtonType(row) {
      if (row.recorderId == 0) {
        return "warning"
      } else return "success"
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
    async updateLiftTrouble(row) {
      const res = await findLiftTrouble({ ID: row.ID });
      if (res.code === 0) {
        this.formData = res.data.reliftTrouble;
        this.dialogFormVisible = true;
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
        ID: 105,
        page: 1,
        pageSize: 9999
      })
      if (res.code == 0) {
        this.fixOptions = res.data.list
      }
    },
    async getReasonOptions(){
      let res = await getCategoriesList({
        ID: 106,
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

<style>
</style>