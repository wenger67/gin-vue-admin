<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增DeviceConfig</el-button>
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
    <el-table-column label="ID" prop="ID" sortable min-width="60"></el-table-column>
    <el-table-column label="配置键值" prop="key" sortable min-width="60"></el-table-column>
    <el-table-column label="配置内容" prop="value" sortable min-width="60"></el-table-column>
    <el-table-column label="配置说明" prop="comment" sortable min-width="60"></el-table-column>
      <el-table-column label="日期" min-width="60">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" width="200">
        <template slot-scope="scope">
          <el-button @click="updateAdDeviceConfig(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteAdDeviceConfig(scope.row)">删除</el-button>
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
        <el-form-item label="配置项" prop="key">
          <el-input v-model="formData.key" placeholder="请输入配置项" clearable :style="{width: '100%'}"></el-input>
        </el-form-item>
        <el-form-item label="配置值" prop="value">
          <el-input v-model="formData.value" placeholder="请输入配置值" clearable :style="{width: '100%'}">
          </el-input>
        </el-form-item>
        <el-form-item label="配置说明" prop="comment">
          <el-input v-model="formData.comment" type="textarea" placeholder="请输入配置说明"
                    :autosize="{minRows: 4, maxRows: 4}" :style="{width: '100%'}"></el-input>
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
    createAdDeviceConfig,
    deleteAdDeviceConfig,
    deleteAdDeviceConfigByIds,
    updateAdDeviceConfig,
    findAdDeviceConfig,
    getAdDeviceConfigList
} from "@/api/adDeviceConfig";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "AdDeviceConfig",
  mixins: [infoList],
  data() {
    return {
      listApi: getAdDeviceConfigList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        key: undefined,
        value: undefined,
        comment: undefined,
      },
      rules: {
        key: [{
          required: true,
          message: '请输入配置项',
          trigger: 'blur'
        }],
        value: [{
          required: true,
          message: '请输入配置值',
          trigger: 'blur'
        }],
        comment: [{
          required: true,
          message: '请输入配置说明',
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
        const res = await deleteAdDeviceConfigByIds({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async updateAdDeviceConfig(row) {
      const res = await findAdDeviceConfig({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.readDeviceConfig;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          key:null,
          value:null,
          comment:null,
      };
    },
    async deleteAdDeviceConfig(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteAdDeviceConfig({ ID: row.ID });
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
          res = await createAdDeviceConfig(this.formData);
          break;
        case "update":
          res = await updateAdDeviceConfig(this.formData);
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
    }
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
</style>