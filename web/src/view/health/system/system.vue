<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">            
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
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
      <el-table-column type="selection" width="40" align="center"></el-table-column>
      <el-table-column sortable label="序号" prop="ID" min-width="50" align="center"/>
      <el-table-column label="电梯" prop="lift.nickName" sortable min-width="60" align="center"></el-table-column> 
      <el-table-column label="时间维度" prop="timeDimension" sortable min-width="80" align="center"></el-table-column> 
      <el-table-column label="维保维度" prop="maintainDimension" sortable min-width="80" align="center"></el-table-column> 
      <el-table-column label="人为维度" prop="humanDimension" sortable min-width="80" align="center"></el-table-column> 
      <el-table-column label="内在维度" prop="innerDimension" sortable min-width="80" align="center"></el-table-column> 
      <el-table-column label="数据维度" prop="sensorDimension" sortable min-width="80" align="center"></el-table-column> 
      <el-table-column label="日期" min-width="100" align="center">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" min-width="100" align="center">
        <template slot-scope="scope">
          <el-button @click="updateHealthSystem(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteHealthSystem(scope.row)">删除</el-button>
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
      此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createHealthSystem,
    deleteHealthSystem,
    deleteHealthSystemByIds,
    updateHealthSystem,
    findHealthSystem,
    getHealthSystemList
} from "@/api/healthSystem";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "HealthSystem",
  mixins: [infoList],
  data() {
    return {
      listApi: getHealthSystemList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
        liftId:null,timeDimension:null,maintainDimension:null,humanDimension:null,innerDimension:null,sensorDimension:null,
      }
    };
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
        const res = await deleteHealthSystemByIds({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async updateHealthSystem(row) {
      const res = await findHealthSystem({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.rehealthSystem;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          liftId:null,
          timeDimension:null,
          maintainDimension:null,
          humanDimension:null,
          innerDimension:null,
          sensorDimension:null,
      };
    },
    async deleteHealthSystem(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteHealthSystem({ ID: row.ID });
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
          res = await createHealthSystem(this.formData);
          break;
        case "update":
          res = await updateHealthSystem(this.formData);
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