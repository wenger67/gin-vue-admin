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
           <template slot-scope="scope">{{scope.row.CreatedAt|formatDateTime}}</template>
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
      <el-row :gutter="15">
        <el-form ref="elForm" :model="formData" :rules="rules" size="medium" label-width="100px"
                 label-position="left">
          <el-col :span="12">
            <el-form-item label="别名" prop="nickName">
              <el-input v-model="formData.nickName" placeholder="请输入别名" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="编号" prop="code">
              <el-input v-model="formData.code" placeholder="请输入编号" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="安装公司" prop="installerId">
              <el-select v-model="formData.installerId" placeholder="请选择安装公司" filterable clearable
                         :style="{width: '100%'}">
                <el-option
                        v-for="item in companyOptions"
                        :key="item.ID"
                        :label="item.fullName"
                        :value="item.ID">
                  <span style="float: left">{{ item.fullName }}</span>
                  <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.category.categoryName }}</span>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="维保公司" prop="maintainerId">
              <el-select v-model="formData.maintainerId" placeholder="请选择维保公司" filterable clearable
                         :style="{width: '100%'}">
                <el-option
                        v-for="item in companyOptions"
                        :key="item.ID"
                        :label="item.fullName"
                        :value="item.ID">
                  <span style="float: left">{{ item.fullName }}</span>
                  <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.category.categoryName }}</span>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="年检单位" prop="checkerId">
              <el-select v-model="formData.checkerId" placeholder="请选择年检单位" filterable clearable
                         :style="{width: '100%'}">
                <el-option
                        v-for="item in companyOptions"
                        :key="item.ID"
                        :label="item.fullName"
                        :value="item.ID">
                  <span style="float: left">{{ item.fullName }}</span>
                  <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.category.categoryName }}</span>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="使用单位" prop="ownerId">
              <el-select v-model="formData.ownerId" placeholder="请选择使用单位" filterable clearable
                         :style="{width: '100%'}">
                <el-option
                        v-for="item in companyOptions"
                        :key="item.ID"
                        :label="item.fullName"
                        :value="item.ID">
                  <span style="float: left">{{ item.fullName }}</span>
                  <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.category.categoryName }}</span>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="出厂时间" prop="factoryTime">
              <el-date-picker v-model="formData.factoryTime"
                              :style="{width: '100%'}" placeholder="请选择出厂时间" clearable></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="安装时间" prop="installTime">
              <el-date-picker v-model="formData.installTime"
                              :style="{width: '100%'}" placeholder="请选择安装时间" clearable></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="年检时间" prop="checkTime">
              <el-date-picker v-model="formData.checkTime"
                              :style="{width: '100%'}" placeholder="请选择年检时间" clearable></el-date-picker>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="模型" prop="liftModelId">
              <el-select v-model="formData.liftModelId" placeholder="请选择模型" filterable clearable
                         :style="{width: '100%'}">
                <el-option
                        v-for="item in modalOptions"
                        :key="item.ID"
                        :label="item.modal"
                        :value="item.ID">
                  <span style="float: left">{{ item.brand }}</span>
                  <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.modal }}</span>
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="用途类别" prop="categoryId">
              <el-select v-model="formData.categoryId" placeholder="请选择用途类别" filterable clearable
                         :style="{width: '100%'}">
                <el-option
                        v-for="item in usageOptions"
                        :key="item.ID"
                        :label="item.categoryName"
                        :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="地址划分" prop="addressId">
              <el-select v-model="formData.addressId" placeholder="请选择地址划分" clearable
                         :style="{width: '100%'}">
                <el-option v-for="(item, index) in addressOptions" :key="index" :label="item.addressName"
                           :value="item.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="9">
            <el-form-item label="智能设备" prop="adDeviceId">
              <el-select v-model="formData.adDeviceId" placeholder="请选择智能设备" filterable clearable
                         :style="{width: '100%'}">
                <el-option
                        v-for="item in deviceOptions"
                        :key="item.ID"
                        :label="item.ID"
                        :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label-width="64px" label="总楼层" prop="floorCount">
              <el-input v-model="formData.floorCount" placeholder="请输入总楼层" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label-width="64px" label="楼栋号" prop="building">
              <el-input v-model="formData.building" placeholder="请输入楼栋号" clearable :style="{width: '100%'}">
              </el-input>
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label-width="64px" label="单元号" prop="cell">
              <el-input v-model="formData.cell" placeholder="请输入单元号" clearable :style="{width: '100%'}">
              </el-input>
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
    createLift,
    deleteLift,
    deleteLiftByIds,
    updateLift,
    findLift,
    getLiftList
} from "@/api/lift";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import {getCompanyList} from "../../../api/company";
import {getLiftModelList} from "../../../api/liftModel";
import {getCategoriesList} from "../../../api/categories";
import {getAddressList} from "../../../api/address";
import {getAdDeviceList} from "../../../api/adDevice";

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
      multipleSelection: [],
      formData: {
        nickName: undefined,
        code: undefined,
        installerId: undefined,
        maintainerId: undefined,
        checkerId: undefined,
        ownerId: undefined,
        factoryTime: null,
        installTime: null,
        checkTime: null,
        liftModelId: undefined,
        categoryId: undefined,
        addressId: undefined,
        adDeviceId: undefined,
        floorCount: undefined,
        building: undefined,
        cell: undefined,
      },
      rules: {
        nickName: [{
          required: true,
          message: '请输入别名',
          trigger: 'blur'
        }],
        code: [{
          required: true,
          message: '请输入编号',
          trigger: 'blur'
        }],
        installerId: [{
          required: true,
          message: '请选择安装公司',
          trigger: 'change'
        }],
        maintainerId: [{
          required: true,
          message: '请选择维保公司',
          trigger: 'change'
        }],
        checkerId: [{
          required: true,
          message: '请选择年检单位',
          trigger: 'change'
        }],
        ownerId: [{
          required: true,
          message: '请选择使用单位',
          trigger: 'change'
        }],
        factoryTime: [{
          required: true,
          message: '请选择出厂时间',
          trigger: 'change'
        }],
        installTime: [{
          required: true,
          message: '请选择安装时间',
          trigger: 'change'
        }],
        checkTime: [{
          required: true,
          message: '请选择年检时间',
          trigger: 'change'
        }],
        liftModelId: [{
          required: true,
          message: '请选择模型',
          trigger: 'change'
        }],
        categoryId: [{
          required: true,
          message: '请选择用途类别',
          trigger: 'change'
        }],
        addressId: [{
          required: true,
          message: '请选择地址划分',
          trigger: 'change'
        }],
        adDeviceId: [{
          required: true,
          message: '请选择智能设备',
          trigger: 'change'
        }],
        floorCount: [{
          required: true,
          message: '请输入总楼层',
          trigger: 'blur'
        }],
        building: [{
          required: true,
          message: '请输入楼栋号',
          trigger: 'blur'
        }],
        cell: [{
          required: true,
          message: '请输入单元号',
          trigger: 'blur'
        }],
      },
      optionParams: {
        page: 1,
        pageSize: 9999
      },
      companyOptions: [],
      modalOptions: [],
      usageOptions: [],
      addressOptions: [],
      deviceOptions: []
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time !== "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd");
      } else {
        return "";
      }
    },
    formatDateTime: function(time) {
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
        const res = await deleteLiftByIds({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async updateLift(row) {
      const res = await findLift({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
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
      if (res.code === 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        await this.getTableData();
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
      let res = await getCompanyList(this.optionParams)
      if (res.code === 0)
        this.companyOptions = res.data.list
    },
    async getModalOptions() {
      let res = await getLiftModelList(this.optionParams)
      if (res.code === 0)
        this.modalOptions = res.data.list
    },
    async getUsageOptions() {
      let params = this.optionParams
      params.ID = 100
      let res = await getCategoriesList(params)
      if (res.code === 0)
        this.usageOptions = res.data.list
    },
    async getAddressOptions() {
      let res = await getAddressList(this.optionParams)
      if (res.code === 0)
        this.addressOptions = res.data.list
    },
    async getDeviceOptions() {
      let res = await getAdDeviceList(this.optionParams)
      if (res.code === 0)
        this.deviceOptions = res.data.list
    },
  },
  async created() {
    await this.getTableData()
    await this.getModalOptions()
    await this.getCompanyOptions()
    await this.getDeviceOptions()
    await this.getUsageOptions()
    await this.getAddressOptions()
  }
};
</script>

<style>
</style>