<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                                      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增lift表</el-button>
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
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="电梯别名" prop="nickName" width="120"></el-table-column> 
    
    <el-table-column label="电梯编号" prop="code" width="120"></el-table-column> 
    
    <el-table-column label="电梯安装公司id" prop="installerId" width="120"></el-table-column> 
    
    <el-table-column label="电梯维保公司id" prop="maintainerId" width="120"></el-table-column> 
    
    <el-table-column label="电梯年检公司id" prop="checkerId" width="120"></el-table-column> 
    
    <el-table-column label="电梯使用公司id" prop="ownerId" width="120"></el-table-column> 
    
    <el-table-column label="电梯出厂时间" prop="factoryTime" width="120"></el-table-column> 
    
    <el-table-column label="电梯安装时间" prop="installTime" width="120"></el-table-column> 
    
    <el-table-column label="电梯年检时间" prop="checkTime" width="120"></el-table-column> 
    
    <el-table-column label="电梯型号" prop="liftModelId" width="120"></el-table-column> 
    
    <el-table-column label="电梯类别" prop="categoryId" width="120"></el-table-column> 
    
    <el-table-column label="总楼层" prop="floorCount" width="120"></el-table-column> 
    
    <el-table-column label="电梯位置地理经度" prop="latitude" width="120"></el-table-column> 
    
    <el-table-column label="电梯位置地理纬度" prop="longitude" width="120"></el-table-column> 
    
    <el-table-column label="地址id" prop="addressId" width="120"></el-table-column> 
    
    <el-table-column label="区域id" prop="regionId" width="120"></el-table-column> 
    
    <el-table-column label="楼栋" prop="building" width="120"></el-table-column> 
    
    <el-table-column label="单元" prop="cell" width="120"></el-table-column> 
    
    <el-table-column label="广告机设备id" prop="adDeviceId" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateLift(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteLift(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
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
    createLift,
    deleteLift,
    deleteLiftByIds,
    updateLift,
    findLift,
    getLiftList
} from "@/api/lift";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Lift",
  mixins: [infoList],
  data() {
    return {
      listApi: getLiftList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
        nickName:null,code:null,installerId:null,maintainerId:null,checkerId:null,ownerId:null,factoryTime:null,installTime:null,checkTime:null,liftModelId:null,categoryId:null,floorCount:null,latitude:null,longitude:null,addressId:null,regionId:null,building:null,cell:null,adDeviceId:null,
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
        const res = await deleteLiftByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateLift(row) {
      const res = await findLift({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.relift;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          nickName:null,
          code:null,
          installerId:null,
          maintainerId:null,
          checkerId:null,
          ownerId:null,
          factoryTime:null,
          installTime:null,
          checkTime:null,
          liftModelId:null,
          categoryId:null,
          floorCount:null,
          latitude:null,
          longitude:null,
          addressId:null,
          regionId:null,
          building:null,
          cell:null,
          adDeviceId:null,
      };
    },
    async deleteLift(row) {
      this.visible = false;
      const res = await deleteLift({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createLift(this.formData);
          break;
        case "update":
          res = await updateLift(this.formData);
          break;
        default:
          res = await createLift(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
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