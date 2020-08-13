<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">   
        <el-form-item label="Subject">
          <el-select v-model="searchInfo.ID" filterable placeholder="Select a subject">
            <el-option v-for="item in subjectOptions"
              :key="item.ID"
              :label="item.subjectName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增类别</el-button>
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
      <el-table-column label="ID" prop="ID" width="120"></el-table-column> 
      <el-table-column label="主体" prop="categorySubject.subjectName" min-width="60"></el-table-column> 
      <el-table-column label="类别" prop="categoryName" min-width="60"></el-table-column> 
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" width="200" fixed="right">
        <template slot-scope="scope">
          <el-button @click="updateCategories(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" size="mini" @click="deleteCategories(scope.row)">删除</el-button>
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

    <el-dialog 
      :before-close="closeDialog" 
      :visible.sync="dialogFormVisible" 
      :title="dialogTitle">
      <el-form :inline=true :model="formData" :rules="rules" ref="categoryForm" label-width="80px">
        <el-form-item label="Subject">
          <el-select v-model="formData.categorySubjectId" filterable placeholder="Select a subject">
            <el-option v-for="item in subjectOptions"
              :key="item.ID"
              :label="item.subjectName"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>        
        <el-form-item label="name" prop="categoryName">
          <el-input autocomplete="off" v-model="formData.categoryName"></el-input>
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
    createCategories,
    deleteCategories,
    deleteCategoriesByIds,
    updateCategories,
    findCategories,
    getCategoriesList
} from "@/api/categories";
import { getAllSubjects } from "@/api/subject";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Categories",
  mixins: [infoList],
  data() {
    return {
      listApi: getCategoriesList,
      dialogFormVisible: false,
      visible: false,
      dialogTitle: '',
      type: "",
      rules: {
        categorySubjectId: [{
          required: true, message: 'please input category subject name', trigger: 'blur'
        }],
        categoryName: [{
          required: true, message: 'please input category subject name', trigger: 'blur'
        }]
      },      
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        categorySubjectId:null,
        categoryName:null,
      },
      subjectOptions: []
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
      const res = await deleteCategoriesByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        await this.getTableData()
      }
    },
    async updateCategories(row) {
      const res = await findCategories({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.recategories;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          categorySubjectId:null,
          categoryName:null,
      };
    },
    async deleteCategories(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteCategories({ ID: row.ID });
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
          res = await createCategories(this.formData);
          break;
        case "update":
          res = await updateCategories(this.formData);
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
      this.dialogTitle = 'Create Category'
      this.dialogFormVisible = true;
    },
    async getSubjects() {
      const res = await getAllSubjects();
      if (res.code === 0) {
        this.subjectOptions = res.data.subjectList;
      }
    }
  },
  created() {
    this.getSubjects();
    this.getTableData();}
};
</script>

<style scoped lang="scss">
    .button-box {
      padding: 10px, 20px;
      .el-button {
        float: right;
      }
    }
    .el-tag--mini {
      margin-left: 5px;
    }
    .warning {
      color: #dc143c;
    }
</style>