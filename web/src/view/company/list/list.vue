<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                      
        <el-form-item label="Subject" prop="subjectId">
          <el-select v-model="searchInfo.subjectId" @change="getCategories()" placeholder="请选择类别主体">
            <el-option v-for="item in subjectOptions"
              :key="item.ID"
              :label="item.subjectName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Category" prop="categoryName">
          <el-select v-model="searchInfo.categoryId" filterable placeholder="请选择类别名称">
            <el-option v-for="item in categoryOptions"
              :key="item.ID"
              :label="item.categoryName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增company</el-button>
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
    <el-table-column type="selection" min-width="30" align="center"></el-table-column>
    <el-table-column sortable label="序号" prop="ID" min-width="50" align="center"></el-table-column>
    <el-table-column sortable label="简称" min-width="100" align="center">
      <template slot-scope="scope">
        <el-popover
          placement="top"
          trigger="click">
          <p>全称: {{ scope.row.fullName }}</p>
          <div slot="reference" class="name-wrapper">
            <el-tag size="medium">{{ scope.row.alias}}</el-tag>
          </div>
        </el-popover>
      </template>      
    </el-table-column>

    <el-table-column sortable label="法人" prop="legalPerson" min-width="60" align="center"></el-table-column>
    <el-table-column label="联系方式" prop="phone" min-width="80" align="center"></el-table-column> 
    <el-table-column sortable label="公司状态" prop="status" min-width="70" align="center"></el-table-column>
    <el-table-column sortable label="统一信用代码" prop="creditCode" min-width="100" align="center">
      <template slot-scope="scope">
        <el-popover
          placement="top"
          trigger="click">
          <p v-if="scope.row.regCode">工商注册号 : {{ scope.row.regCode }}</p>
          <p v-if="scope.row.orgCode">组织机构号 : {{ scope.row.orgCode }}</p>
          <p v-if="scope.row.taxCode">纳税人号: {{ scope.row.taxCode }}</p>
          <div slot="reference" class="name-wrapper">
            <el-tag size="medium">{{ scope.row.creditCode}}</el-tag>
          </div>
        </el-popover>
      </template>  
    </el-table-column>
    <el-table-column sortable label="注册地址" prop="address" min-width="200" align="center"></el-table-column>
    <el-table-column sortable label="类别" prop="category.categoryName" min-width="60" align="center"></el-table-column>
    <el-table-column sortable label="日期" :inline="true" min-width="100" align="center">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
      <el-table-column fixed="right" min-width="100" label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateCompany(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" size="mini" @click="deleteCompany(scope.row)">删除</el-button>
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
          <el-col :span="12">
            <el-form-item label="公司全称" prop="fullName">
              <el-input v-model="formData.fullName" placeholder="请输入公司全称" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="公司简称" prop="alias">
              <el-input v-model="formData.alias" placeholder="请输入公司简称" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="法人信息" prop="legalPerson">
              <el-input v-model="formData.legalPerson" placeholder="请输入法人信息" clearable
                :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="联系方式" prop="phone">
              <el-input v-model="formData.phone" placeholder="请输入联系方式" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="统一代码" prop="creditCode">
              <el-input v-model="formData.creditCode" placeholder="请输入统一社会信用代码" clearable
                :style="{width: '100%'}"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="工商注册号" prop="regCode">
              <el-input v-model="formData.regCode" placeholder="请输入工商注册号" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="组织机构号" prop="orgCode">
              <el-input v-model="formData.orgCode" placeholder="请输入组织机构号" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="纳税人识别" prop="taxCode">
              <el-input v-model="formData.taxCode" placeholder="请输入纳税人识别号" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="注册地址" prop="address">
              <el-input v-model="formData.address" placeholder="请输入注册地址" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="类别主体" prop="subject">
              <el-select v-model="formData.subject" @change="getCategories()" placeholder="请选择类别主体" clearable :style="{width: '100%'}">
                <el-option v-for="item in subjectOptions"
                  :key="item.ID"
                  :label="item.subjectName"
                  :value="item.ID">
                </el-option>
              </el-select>              
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="类别名称" prop="categoryId">
              <el-select v-model="formData.categoryId" filterable placeholder="请选择类别名称" clearable :style="{width: '100%'}">
                <el-option v-for="item in categoryOptions"
                  :key="item.ID"
                  :label="item.categoryName"
                  :value="item.ID">
                </el-option>
              </el-select>
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
    createCompany,
    deleteCompany,
    deleteCompanyByIds,
    updateCompany,
    findCompany,
    getCompanyList
} from "@/api/company";  //  此处请自行替换地址
import { getCategoriesList } from "@/api/categories";
import { getAllSubjects } from "@/api/subject";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Company",
  mixins: [infoList],
  data() {
    return {
      listApi: getCompanyList,
      dialogFormVisible: false,
      visible: false,
      type: "search",
      subjectOptions: [],
      categoryOptions: [],
      rules: {
        fullName: [{
          required: true,
          message: '请输入公司全称',
          trigger: 'blur'
        }],
        alias: [{
          required: true,
          message: '请输入公司简称',
          trigger: 'blur'
        }],
        legalPerson: [{
          required: true,
          message: '请输入法人信息',
          trigger: 'blur'
        }],
        phone: [{
          required: true,
          message: '请输入联系方式',
          trigger: 'blur'
        }],
        creditCode: [{
          required: true,
          message: '请输入统一社会信用代码',
          trigger: 'blur'
        }],
        regCode: [],
        orgCode: [],
        taxCode: [],
        address: [{
          required: true,
          message: '请输入注册地址',
          trigger: 'blur'
        }],
        subject: [{
          required: true,
          message: '请选择类别主体',
          trigger: 'change'
        }],
        categoryId: [{
          required: true,
          message: '请选择类别名称',
          trigger: 'change'
        }],
      },      
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        fullName:null,alias:null,legalPerson:null,phone:null,status:null,regCode:null,
        orgCode:null,creditCode:null,taxCode:null,address:null,categoryId:null,
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
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.searchInfo = {
          ID: this.searchInfo.categoryId
        }          
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteCompanyByIds({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async updateCompany(row) {
      const res = await findCompany({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.recompany;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          fullName:null,
          alias:null,
          legalPerson:null,
          phone:null,
          status:null,
          regCode:null,
          orgCode:null,
          creditCode:null,
          taxCode:null,
          address:null,
          categoryId:null,
      };
      this.type = "search";
    },
    async deleteCompany(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteCompany({ ID: row.ID });
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
      this.$refs.elForm.validate(async valid => {
        if (valid) {
          let res
          if (this.type === "create") {
             res = await createCompany(this.formData);
          } else if (this.type === "update") {
            res = await updateCompany(this.formData)
          }
          if (res.code === 0) {
            this.$message({
              type:"success",
              message:"创建/更改成功"
            })
            this.closeDialog();
            await this.getTableData();
          }          
        }
      })
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    },
    async getSubjects() {
      const res = await getAllSubjects();
      if (res.code === 0) {
        this.subjectOptions = res.data.subjectList;
      }
    },
    async getCategories() {
      const params = {
        page: 1,
        pageSize: 100,
        ID: this.type === "create" ? this.formData.subject : this.searchInfo.subjectId
      };
      const res = await getCategoriesList(params);
      if (res.code === 0) {
        this.categoryOptions = res.data.list;
      }

    }
  },
  async created() {
    await this.getSubjects();
    await this.getTableData();
  }
};
</script>

<style>
</style>