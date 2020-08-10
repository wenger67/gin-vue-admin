<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增liftModel表</el-button>
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
  
      <el-table-column label="电梯生产商" prop="company.fullName" min-width="60"></el-table-column>
      
      <el-table-column label="电梯品牌" prop="brand" min-width="60"></el-table-column> 
      
      <el-table-column label="电梯型号" prop="modal" min-width="60"></el-table-column> 
      
      <el-table-column label="电梯载重" prop="load" min-width="60"></el-table-column> 
      <el-table-column label="日期"  min-width="60">
          <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
   
      <el-table-column label="按钮组" fixed="right" width="200">
        <template slot-scope="scope">
          <el-button @click="updateLiftModel(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteLiftModel(scope.row)">删除</el-button>
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
      <el-row :gutter="15">
        <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px"
          label-position="left">
          <el-col :span="24">
            <el-form-item label="厂商" prop="factoryId">
              <el-select v-model="formData.factoryId" placeholder="请选择厂商" filterable clearable
                :style="{width: '100%'}">
                <el-option v-for="item in factoryOptions"
                           :key="item.value.fullName"
                           :label="item.value.fullName"
                           :value="item.value.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="品牌" prop="brand">
              <el-input v-model="formData.brand" placeholder="请输入品牌" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="型号" prop="modal">
              <el-input v-model="formData.modal" placeholder="请输入型号" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="载重" prop="load">
              <el-input-number v-model="formData.load" placeholder="请输入载重" clearable :style="{width: '100%'}">
              </el-input-number>
            </el-form-item>
          </el-col>
        </el-form>
      </el-row>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createLiftModel,
    deleteLiftModel,
    deleteLiftModelByIds,
    updateLiftModel,
    findLiftModel,
    getLiftModelList
} from "@/api/liftModel";
import {
  getCompanyList
} from "@/api/company";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "LiftModel",
  mixins: [infoList],
  data() {
    return {
      listApi: getLiftModelList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      factoryOptions: [],
      multipleSelection: [],
      formData: {
        company:{},brand:null,modal:null,load:null,
      },
      rules: {
        factoryId: [{
          required: true,
          message: '请选择厂商',
          trigger: 'change'
        }],
        brand: [{
          required: true,
          message: '请输入品牌',
          trigger: 'blur'
        }],
        modal: [{
          required: true,
          message: '请输入型号',
          trigger: 'blur'
        }],
        load: [{
          required: true,
          message: '请输入载重',
          trigger: 'blur'
        }]
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time !== "") {
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
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10        
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      handleClosePopover(index) {
        this.$refs[`popover-${index}`].doClose()
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteLiftModelByIds({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async updateLiftModel(row) {
      const res = await findLiftModel({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.reliftModel;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          company:{},
          brand:null,
          modal:null,
          load:null,
      };
    },
    async deleteLiftModel(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteLiftModel({ ID: row.ID });
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
      switch (this.type) {
        case "create":
          res = await createLiftModel(this.formData);
          break;
        case "update":
          res = await updateLiftModel(this.formData);
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
    async getFactoryOptions() {
        let params = {
          page: 1,
          pageSize: 9999
        }
        let res = await getCompanyList(params);
        if (res.code === 0) {
          res.data.list.forEach(element => {
            this.factoryOptions.push({value: element})
          })
        }
    }
  },
  async created() {
    await this.getTableData();
    await this.getFactoryOptions();
  }
};
</script>

<style>
</style>