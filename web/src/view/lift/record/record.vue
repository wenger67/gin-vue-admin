<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增liftRecord表</el-button>
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
    <el-table-column type="selection" width="55"></el-table-column>
    
      <el-table-column label="电梯" prop="lift.nickName" sortable min-width="60"></el-table-column> 
      <el-table-column label="记录类别" prop="category.categoryName" sortable min-width="60"></el-table-column> 
      <el-table-column label="开始时间" sortable min-width="80">
        <template slot-scope="scope">{{scope.row.startTime|formatDate}}</template>
      </el-table-column> 
      <el-table-column label="结束时间" prop="endTime" sortable min-width="60">
        <template slot-scope="scope">{{scope.row.endTime|formatDate}}</template>        
      </el-table-column> 
      <el-table-column label="图片记录" prop="images" sortable min-width="60"></el-table-column> 
      <el-table-column label="文字记录" prop="content" sortable min-width="60"></el-table-column> 
      <el-table-column label="操作人员" prop="worker.realName" sortable min-width="60"></el-table-column> 
      <el-table-column label="记录人员" prop="recorder.realName" sortable min-width="60"></el-table-column> 
      <el-table-column label="日期" min-width="80">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" min-width="200">
        <template slot-scope="scope">
          <el-button @click="updateLiftRecord(scope.row)" size="small" :type="getButtonType(scope.row)">{{ scope.row|formatTitle }}</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteItem(scope.row)">删除</el-button>
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
      <el-form :model="formData" ref="formData" label-width="100px" size="medium" label-postion="left">
        <el-form-item label="">
          <el-steps :active="stepActive" align-center>
            <el-step v-for="item in stepItems"
              :key="item.key"
              :title="item.title"
              :description="item.description"
              :icon="item.icon">
            </el-step>
          </el-steps>
        </el-form-item>
        <el-form-item v-if="stepActive == 0" label="lift id" prop="liftId">
          <el-select v-model="formData.liftId" filter clearable placeholder="please select lift" >
            <el-option v-for="item in liftOptions"
              :key="item.ID"
              :label="item.nickName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive != 0" label="lift id" prop="liftId">
          <el-input v-model="formData.lift.nickName" disabled placeholder=""></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive == 0" label="record type" prop="categoryId">
          <el-select v-model="formData.categoryId" filter clearable placeholder="please select record type">
            <el-option v-for="item in categoryOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive != 0" label="record type" prop="categoryId">
          <el-input v-model="formData.category.categoryName" disabled placeholder=""></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive == 0" label="worker" prop="workerId">
          <el-select v-model="formData.workerId" filter clearable placeholder="please select workerId">
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="stepActive != 0" label="worker" prop="workerId">
          <el-input v-model="formData.worker.realName" disabled placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 1" label="content" prop="content">
          <el-input v-model="formData.content" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder="please input record content"></el-input>
        </el-form-item>
        <el-form-item v-if="stepActive == 2" label="content" prop="content">
          <el-input v-model="formData.content" type="textarea" :autosize="{minRows: 4, maxRows: 4}" placeholder=""></el-input>
        </el-form-item>

        <el-form-item v-if="stepActive == 1" label="images" prop="images">
          <el-upload
            action="#"
            :on-remove="handleRemove"
            multiple
            list-type="picture-card">
            <i class="el-icon-plus"></i>
          </el-upload>
        </el-form-item>
        
        <el-form-item v-if="stepActive == 2" label="recorder" prop="recorderId">
          <el-select v-model="formData.recorderId" filter clearable placeholder="please select recorderId">
            <el-option v-for="item in userOptions"
              :key="item.ID"
              :label="item.realName"
              :value="item.ID">
            </el-option>
          </el-select>
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
    createLiftRecord,
    deleteLiftRecord,
    deleteLiftRecordByIds,
    fillLiftRecord,
    reviewLiftRecord,
    findLiftRecord,
    getLiftRecordList
} from "@/api/liftRecord";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { getUserList } from '@/api/user';
import { getCategoriesList } from '@/api/categories';
import { getLiftList } from "@/api/lift";

export default {
  name: "LiftRecord",
  mixins: [infoList],
  data() {
    return {
      listApi: getLiftRecordList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      stepItems: [
        {key: 1, title: "步骤 1", description: "create record", icon:"el-icon-edit"},
        {key: 2, title: "步骤 2", description: "fill record", icon:"el-icon-upload"},
        {key: 3, title: "步骤 3", description: "review record", icon:"el-icon-s-promotion"}
      ],
      categoryOptions:[],
      userOptions:[],
      liftOptions:[],
      multipleSelection: [],
      formData: {
        liftId:null,lift:{}, categoryId:null, category:{},
        images:"",content:"",startTime:"0001-01-01T00:00:00Z",endTime:"0001-01-01T00:00:00Z",
        workerId:null,worker:{}, recorderId:0, recorder:{}
      }
    };
  },
  computed: {
      stepActive: function() {
        if (this.formData.content == "" && this.formData.images == "" && this.formData.startTime == "0001-01-01T00:00:00Z") {
          return 0
        }
        if (this.formData.recorderId == 0 && this.formData.endTime == "0001-01-01T00:00:00Z") {
          return 1
        }

        return 2
      }
  },
  filters: {
    formatDate: function(time) {
      if (time == "0001-01-01T00:00:00Z") {
        return "NULL"
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
      if (row.recorderId == 0) {
        return "InProcess"
      } else {
        return "View"
      }
    }
  },
  methods: {
      onSubmit() {
        this.page = 1
        this.pageSize = 10            
        this.getTableData()
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
    async updateLiftRecord(row) {
      const res = await findLiftRecord({ ID: row.ID });
      if (res.code === 0) {
        this.formData = res.data.reliftRecord;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        liftId:null,lift:{}, categoryId:null, category:{},
        images:"",content:"",startTime:"0001-01-01T00:00:00Z",endTime:"0001-01-01T00:00:00Z",
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
      if (this.stepActive == 1) {
        this.type = "create"
      } else if (this.stepActive == 2) {
        this.type = "fill"
      } else this.type = "review"

      let res;
      let param = {}
      switch (this.type) {
        case "create":
          param.liftId = this.formData.liftId
          param.categoryId = this.formData.categoryId
          param.workerId = this.formData.workerId
          res = await createLiftRecord(param);
          break;
        case "fill":
          param.recordId = this.formData.ID
          param.content = this.formData.content
          param.images = this.formData.images
          res = await fillLiftRecord(param);
          break;
        case "review":
          param.recordId = this.formData.ID
          param.recorderId = this.formData.recorderId
          res = await reviewLiftRecord(param)
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
      console.log(this.stepActive)
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
        ID: 103, // TODO hard code
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

<style>

.el-form-item{
  text-align: center;
}
</style>