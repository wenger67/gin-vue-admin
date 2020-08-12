<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增userAdmin表</el-button>
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
      <el-table-column label="uuid" prop="uuid" min-width="60"></el-table-column>
      <el-table-column label="登陆手机号" prop="phoneNumber" min-width="60"></el-table-column>
      <el-table-column label="用户真名" prop="realName" min-width="60"></el-table-column>
      <el-table-column label="用户昵称" prop="nickName" min-width="60"></el-table-column>
      <el-table-column label="用户头像" min-width="60">
        <template slot-scope="scope">
          <el-image :src="scope.row.avatar" lazy/>
        </template>
      </el-table-column>
      <el-table-column label="用户所属公司" prop="company.fullName" min-width="60"></el-table-column>
      <el-table-column label="用户住址" prop="address" min-width="60"></el-table-column>
      <el-table-column label="用户角色" prop="authority.authorityName" min-width="60"></el-table-column>
      <el-table-column label="日期" min-width="60">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" width="200">
        <template slot-scope="scope">
          <el-button @click="updateUserAdmin(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteUserAdmin(scope.row)">删除</el-button>
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
      <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px"
               label-position="left">
        <el-form-item label="手机号" prop="phoneNumber">
          <el-input v-model="formData.phoneNumber" placeholder="请输入手机号" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="formData.password" placeholder="请输入密码" clearable show-password
                    :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="真实姓名" prop="realName">
          <el-input v-model="formData.realName" placeholder="请输入真实姓名" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="昵称" prop="nickName">
          <el-input v-model="formData.nickName" placeholder="请输入昵称" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="所属公司" prop="companyId">
          <el-select v-model="formData.companyId" placeholder="请选择所属公司" filterable clearable
                     :style="{width: '100%'}">
            <el-option v-for="item in companyOptions"
                       :key="item.ID"
                       :label="item.fullName"
                       :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="角色" prop="authorityId">
          <el-select v-model="formData.authorityId" placeholder="请选择角色" filterable clearable
                     :style="{width: '100%'}">
            <el-option v-for="item in authorityOptions"
                       :key="item.authorityId"
                       :label="item.authorityName"
                       :value="item.authorityId">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="住址" prop="address">
          <el-input v-model="formData.address" placeholder="请输入住址" clearable :style="{width: '100%'}">
          </el-input>
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
    getUser,
    getUserList,
    deleteUserList,
    deleteUser,
    createUser
  } from "@/api/user";
  import { getCompanyList } from "@/api/company";
  import { getAuthorityList } from "@/api/authority";
  import { formatTimeToStr } from "@/utils/data";
  import infoList from "@/components/mixins/infoList";

  export default {
    name: "list",
    mixins: [infoList],
    data() {
      const checkUsername = (rule, value, callback) => {
        if (!(/^1[3456789]\d{9}$/.test(value))) {
          return callback(new Error("请输入正确的用户名"));
        } else {
          callback();
        }
      };
      const checkPassword = (rule, value, callback) => {
        if (value.length < 6 || value.length > 12) {
          return callback(new Error("请输入正确的密码"));
        } else {
          callback();
        }
      };
      return {
        listApi: getUserList,
        dialogFormVisible: false,
        type: "",
        deleteVisible: false,
        multipleSelection: [],
        companyOptions:[],
        authorityOptions:[],
        formData: {
          phoneNumber:null,password:null,realName:null,nickName:null,
          companyId:null, address:null,authorityId:null
        },
        rules: {
          phoneNumber: [{
            validator: checkUsername,
            trigger: 'blur'
          }],
          password: [{
            validator: checkPassword,
            trigger: 'blur'
          }],
          realName: [{
            required: true,
            message: '请输入真实姓名',
            trigger: 'blur'
          }],
          nickName: [{
            required: true,
            message: '请输入昵称',
            trigger: 'blur'
          }],
          companyId: [{
            required: true,
            message: '请选择所属公司',
            trigger: 'change'
          }],
          authorityId: [{
            required: true,
            message: '请选择角色',
            trigger: 'change'
          }],
          address: [{
            required: true,
            message: '请输入住址',
            trigger: 'blur'
          }],
        },
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
      async onDelete() {
        const ids = []
        this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
        const res = await deleteUserList({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
      async updateUserAdmin(row) {
        const res = await getUser({ ID: row.ID });
        this.type = "update";
        if (res.code === 0) {
          this.formData = res.data.reuser;
          this.dialogFormVisible = true;
        }
      },
      closeDialog() {
        this.dialogFormVisible = false;
        this.formData = {
          phoneNumber:null,password:null,realName:null,nickName:null,
          companyId:null, address:null,authorityId:null
        };
      },
      async deleteUserAdmin(row) {
        this.$confirm('this operation is dangerous, continue ?', 'Hint', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        .then(async () => {
          const res = await deleteUser({ ID: row.ID });
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
            res = await createUser(this.formData);
            break;
          case "update":
            res = await updateUserAdmin(this.formData);
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
      async getCompanyOptions() {
        this.companyOptions = []
        let res = await getCompanyList({
          // TODO get list by querystring
          page: 1,
          pageSize: 9999
        })
        if (res.code === 0) {
          this.companyOptions = res.data.list
        }
      },
      async getAuthorityOptions() {
        this.authorityOptions = []
        let res = await getAuthorityList({
          // TODO get list by querystring
          page: 1,
          pageSize: 9999
        })
        if (res.code === 0) {
          this.authorityOptions = res.data.list
        }
      }
    },
    async created() {
      await this.getTableData()
      await this.getAuthorityOptions()
      await this.getCompanyOptions()
    }
  };
</script>

<style>
</style>