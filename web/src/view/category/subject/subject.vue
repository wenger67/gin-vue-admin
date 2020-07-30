<template>
  <div>
    <div class="search-term">
        <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
          <el-form-item label="Subject Name">
            <el-input placeholder="Subject Name" v-model="searchInfo.subjectName"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">Search</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="openDialog('addSubject')">Create Subject</el-button>
          </el-form-item>
        </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="ID" min-width="60" prop="ID" sortable="custom"></el-table-column>
      <el-table-column label="主体" min-width="60" prop="subjectName" sortable="custom"></el-table-column>
      <el-table-column label="日期" width="180">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column fixed="right" label="按钮组" width="200">
        <template slot-scope="scope">
          <el-button type="primary" @click="editSubject(scope.row)" size="small" icon="el-icon-edit">Edit</el-button>
          <el-button type="danger" @click="deleteSubject(scope.row)" size="small" icon="el-icon-delete">Delete</el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="page"
      :page-sizes="[10, 30, 50, 100]"
      :page-size="pageSize"
      layout="total, sizes, prev, pager, next, jumper"
      :style="{float:'right', padding:'20px'}"
      :total="total">
    </el-pagination>
    
    <el-dialog
      :before-close="closeDialog"
      :title="dialogTitle"
      :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="subjectForm">
        <el-form-item label="name" prop="subjectName">
          <el-input autocomplete="off" v-model="form.subjectName" ref="nameInput"></el-input>
        </el-form-item>
      </el-form>
      
      <span class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </span>
    </el-dialog>
    
  </div>
</template>

<script>

import {
  getSubjectById,
  getSubjectList,
  createSubject,
  updateSubject,
  deleteSubject
} from '@/api/subject'
import infoList from '@/components/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
import { formatTimeToStr } from "@/utils/data";

export default {
    name:"Subject",
    mixins: [infoList],
    data() {
      return {
        listApi: getSubjectList,
        dialogFormVisible: false,
        dialogTitle: 'Create Subject',
        form: {
          subjectName: ''
        },
        type: '',
        rules: {
          subjectName: [{
            required: true, message: 'please input category subject name', trigger: 'blur'
          }],
        }
      }
    },
    filters: {
      formatDate: function(time) {
        if (time != null && time != "") {
          var date = new Date(time);
          return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
        } else {
          return "";
        }
      },
    },
    methods: {
      sortChange({prop, order}) {
        if (prop) {
          this.searchInfo.orderKey = toSQLLine(prop)
          this.searchInfo.desc = order == 'descending'
        }
        this.getTableData()
      },
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.getTableData()
      },
      initForm() {
        this.$refs.subjectForm.resetFields()
        this.form = {
          subjectName: ''
        }
      },
      closeDialog() {
        this.initForm()
        this.dialogFormVisible = false
      },
      openDialog(type) {
        switch (type) {
          case 'addSubject':
            this.dialogTitle = 'Create Subject'
            break
          case 'edit':
            this.dialogTitle = 'Edit Subject'
            break
          default:
            break
        }
        this.type = type
        this.dialogFormVisible = true
      },
      async editSubject(row) {
        const res = await getSubjectById({ id: row.ID })
        this.form = res.data.subject
        this.openDialog('edit')
      },
      async deleteSubject(row) {
        this.$confirm('this operation is dangerous, continue ? ', 'hint', {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning'
        })
        .then(async () => {
          const res = await deleteSubject(row)
          if (res.code == 0) {
            this.$message({
              type: 'success',
              message: 'delete success!'
            })
            this.getTableData()
          }
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: 'delete canceled!'
          })
        })
      },
      async enterDialog() {
        this.$refs.subjectForm.validate(async valid => {
          if (valid) {
            switch (this.type) {
              case 'addSubject':
                {
                  const res = await createSubject(this.form)
                  if (res.code == 0) {
                    this.$message({
                      type: 'success',
                      message: 'create success',
                      showClose: true
                    })
                  }
                  this.getTableData()
                  this.closeDialog()
                }
                break
              case 'edit':
                {
                  const res = await updateSubject(this.form)
                  if (res.code == 0) {
                    this.$message({
                      type: 'success',
                      message: 'edit success',
                      showClose: true
                    })
                  }
                  this.getTableData()
                  this.closeDialog()
                }
                break
              default:
                {
                  this.$message({
                    type: 'error',
                    message: 'unknow operation',
                    showClose: true
                  })
                }
                break
            }
          }
        })
      }     
    },
    async created() {
      await this.getTableData()
    }
}
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