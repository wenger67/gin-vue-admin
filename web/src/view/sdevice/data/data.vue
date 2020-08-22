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
    <el-table-column type="selection" width="55"></el-table-column>

    
    <el-table-column label="设备id" prop="deviceId" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="x轴加速度" prop="accx" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="y轴加速度" prop="accy" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="z轴加速度" prop="accz" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="x轴倾斜角" prop="degx" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="y轴倾斜角" prop="degy" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="z轴倾斜角" prop="degz" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="电梯速度" prop="speedz" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="电梯当前楼层" prop="floor" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="电梯门状态类型" prop="doorStateId" sortable min-width="60"></el-table-column> 
    
    <el-table-column label="电梯内是否有人" prop="peopleInside" min-width="60">
         <template slot-scope="scope">{{scope.row.peopleInside|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="电梯状态/故障类型" prop="troubleId" sortable min-width="60"></el-table-column> 
    
      <el-table-column label="日期" min-width="60">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" width="200">
        <template slot-scope="scope">
          <el-button @click="updateAdDeviceData(scope.row)" size="small" type="primary">变更</el-button>
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
  </div>
</template>

<script>
import {
    deleteAdDeviceData,
    deleteAdDeviceDataByIds,
    getAdDeviceDataList
} from "@/api/adDeviceData";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "AdDeviceData",
  mixins: [infoList],
  data() {
    return {
      listApi: getAdDeviceDataList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: []
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
      if (this.searchInfo.peopleInside===""){
        this.searchInfo.peopleInside=null
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
      const res = await deleteAdDeviceDataByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        await this.getTableData()
      }
    },
    async deleteAdDeviceData(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteAdDeviceData({ ID: row.ID });
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
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
</style>