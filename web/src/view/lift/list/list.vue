<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                                      
        <el-form-item>
        <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="refresh" type="primary">刷新</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增电梯</el-button>
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
      <el-table-column sortable label="序号" prop="ID" min-width="50" align="center"></el-table-column>
      <el-table-column label="别名" prop="nickName" min-width="70" align="center">
        <template slot-scope="scope">
          <el-popover
            placement="top"
            trigger="click">
            <p>安装公司: {{ scope.row.installer.fullName }}</p>
            <p>维保公司: {{ scope.row.maintainer.fullName }}</p>
            <p>年检公司: {{ scope.row.checker.fullName }}</p>
            <p>使用公司: {{ scope.row.owner.fullName }}</p>
            <p>出厂时间: {{ scope.row.factoryTime|formatDate }}</p>
           <div slot="reference" class="name-wrapper">
              <el-tag size="medium">{{ scope.row.nickName}}</el-tag>
            </div>
          </el-popover>
        </template>          
      </el-table-column>
      <el-table-column label="安装时间" prop="installTime" min-width="70" align="center">
        <template slot-scope="scope">{{scope.row.factoryTime|formatDate}}</template>
      </el-table-column>
      <el-table-column label="年检时间" prop="checkTime" min-width="70" align="center">
        <template slot-scope="scope">{{scope.row.factoryTime|formatDate}}</template>
      </el-table-column>
      <el-table-column label="型号" min-width="100" align="center">
        <template slot-scope="scope">
          {{ scope.row.liftModel.brand }}{{ scope.row.liftModel.modal }}
        </template>
      </el-table-column>
      <el-table-column label="类别" prop="category.categoryName" min-width="48" align="center"></el-table-column>
      <el-table-column label="总楼层" prop="floorCount" min-width="40" align="center"></el-table-column>
      <el-table-column label="地址" min-width="150" align="center">
        <template slot-scope="scope">
          {{ scope.row.address.addressName }}{{ scope.row.building }}{{ scope.row.cell }}
        </template>
      </el-table-column>
      <el-table-column label="设备" min-width="40">
        <template slot-scope="scope">
          <el-tag :type="scope.row.adDevice.online === true ? 'success':'warning'">
            {{scope.row.adDevice.online|formatBoolean}}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="日期" min-width="100" align="center">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDateTime}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" min-width="200" align="center">
        <template slot-scope="scope">
          <el-button type="primary" @click="assign(scope.row)" size="small" icon="el-icon-setting">Assign</el-button>
          <el-button @click="updateLift(scope.row)" icon="el-icon-edit" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="small" @click="deleteLift(scope.row)">删除</el-button>
          <el-button type="success" icon="el-icon-info" size="small" @click="toDetail(scope.row)">详情</el-button>
          <!--  TODO more operations  -->
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

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" :title="dialogTitle">
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
              <el-select v-model="formData.ownerId" placeholder="请选择使用单位" @change="change" filterable clearable
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
              <el-select v-model="formData.addressId" @change="handleSelectAddress" placeholder="请选择地址划分" clearable
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
              <el-input-number v-model="formData.floorCount" placeholder="请输入总楼层" clearable :style="{width: '100%'}">
              </el-input-number>
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
              <el-input-number v-model="formData.cell" placeholder="请输入单元号" clearable :style="{width: '100%'}">
              </el-input-number>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="address">
              <el-amap-search-box class="search-box" :search-option="searchOption" :on-search-result="onSearchResult"></el-amap-search-box>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-amap vid="amapDemo" :center="mapCenter" :amap-manager="aMapManager" :zoom="12" :events="events" class="amap-demo">
              <el-amap-circle
                      v-for="(circle, index) in circles"
                      :center="circle.center"
                      :radius="circle.radius"
                      :key="index"
                      :fill-opacity="circle.fillOpacity"
                      :events="circle.events">
              </el-amap-circle>
            </el-amap>
          </el-col>
        </el-form>
      </el-row>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog :before-close="closeAssignDialog" :visible.sync="assignDialogVisible" :title="assignDialogTitle">
      <el-row :gutter="15">
        <el-form ref="elAssignForm" :model="assignFormData" :rules="assignRules" size="medium" label-width="100px"
          label-position="left">
          <el-col :span="8">
            <el-form-item label="安装单位">
              {{assignFormData.installer != null ? assignFormData.installer.fullName:"null"}}              
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="安装人员" prop="installIds">
              <el-select v-model="assignFormData.installIds" @change="$forceUpdate()" placeholder="请选择下拉选择安装人员" multiple clearable
                :style="{width: '100%'}">
                <el-option v-for="item in installUserOptions"
                  :key="item.ID"
                  :label="item.realName"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-button type="primary" size="small" @click="assignBase(installCateogry)">Setting</el-button>
          </el-col>
                    
          <el-col :span="8">
            <el-form-item label="维保单位">
              {{assignFormData.maintainer != null ? assignFormData.maintainer.fullName:"null"}}              
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="维保人员" prop="maintainIds">
              <el-select v-model="assignFormData.maintainIds" @change="$forceUpdate()" placeholder="请选择下拉选择维保人员" multiple clearable
                :style="{width: '100%'}">
                <el-option v-for="item in maintainUserOptions"
                  :key="item.ID"
                  :label="item.realName"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-button type="primary" size="small" @click="assignBase(maintainCategory)">Setting</el-button>
          </el-col>

          <el-col :span="8">
            <el-form-item label="年检单位">
              {{assignFormData.checker != null ? assignFormData.checker.fullName:"null"}}              
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="年检人员" prop="checkIds">
              <el-select v-model="assignFormData.checkIds" @change="$forceUpdate()" placeholder="请选择下拉选择年检人员" multiple clearable
                :style="{width: '100%'}">
                <el-option v-for="item in checkUserOptions"
                  :key="item.ID"
                  :label="item.realName"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-button type="primary" size="small" @click="assignBase(checkCategory)">Setting</el-button>
          </el-col>

          <el-col :span="8">
            <el-form-item label="使用单位">
              {{assignFormData.owner != null ? assignFormData.owner.fullName:"null"}}              
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="管理人员" prop="ownIds">
              <el-select v-model="assignFormData.ownIds" @change="$forceUpdate()" placeholder="请选择下拉选择管理人员" multiple clearable
                :style="{width: '100%'}">
                <el-option v-for="item in ownUserOptions"
                  :key="item.ID"
                  :label="item.realName"
                  :value="item.ID">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="6">
            <el-button type="primary" size="small" @click="assignBase(ownerCategory)">Setting</el-button>
          </el-col>

        </el-form>
      </el-row>
      <div slot="footer">
        <el-button type="primary" @click="closeAssignDialog">确定</el-button>
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
} from "@/api/lift";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import {getCompanyList} from "@/api/company";
import {getLiftModelList} from "@/api/liftModel";
import {getCategoriesList} from "@/api/categories";
import {getAddressList} from "@/api/address";
import {getAdDeviceList} from "@/api/adDevice";
import VueAMap from "vue-amap";
import { getUserList } from '@/api/user';
import {getUserLiftList, deleteUserLiftByIds, createUserLift} from '@/api/userLift'
import _ from "lodash"

let aMapManager = new VueAMap.AMapManager();

export default {
  name: "Lift",
  mixins: [infoList],
  data() {
    return {
      aMapManager,
      events: {
        init(o) {
          // eslint-disable-next-line no-undef
          let marker = new AMap.Marker({
            position: [121.59996, 31.197646]
          });
          marker.setMap(o);
        }
      },
      circles: [
        {
          center: [121.5273285, 31.21515044],
          radius: 200,
          fillOpacity: 0.5,
          events: {
          }
        }
      ],
      searchOption: {
        city: '上海',
        citylimit: true
      },
      mapCenter: [121.59996, 31.197646],
      listApi: getLiftList,
      dialogFormVisible: false,
      dialogTitle: "",
      assignDialogVisible: false,
      assignDialogTitle: "",
      videoCallTitle: "",
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
      assignFormData: {
        installer: undefined,
        installUser: [],
        maintainer: undefined,
        maintainUser: [],
        checker: undefined,
        checkUser: [],
        user: undefined,
        ownerUser: [],
      },
      installUserOptions:[],
      maintainUserOptions:[],
      checkUserOptions:[],
      ownUserOptions:[],
      installCateogry:1104,
      maintainCategory: 1101,
      checkCategory: 1103,
      ownerCategory:1102,
      assignRules: {
        installIds: [{
          required: true,
          type: 'array',
          message: '请至少选择一个安装人员',
          trigger: 'blur'
        }],
        maintainIds: [{
          required: true,
          type: 'array',
          message: '请至少选择一个维保人员',
          trigger: 'blur'
        }],
        checkIds: [{
          required: true,
          type: 'array',
          message: '请至少选择一个年检人员',
          trigger: 'blur'
        }],
        ownIds: [{
          required: true,
          type: 'array',
          message: '请至少选择一个管理人员',
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
        return bool ? "在线" :"离线";
      } else {
        return "";
      }
    }
  },
  methods: {
    toDetail(row){
      this.$router.push({
        name:"liftDetail",
        params:{
          id:row.ID
        }
      })
    },
    onSubmit() {
      this.page = 1
      this.pageSize = 10
      this.getTableData()
    },
    change(){
      console.log("change")
      console.log(this.formData.ownerId)
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    handleSelectAddress() {
      let id = this.formData.addressId
      let address = this.addressOptions.filter(item => item.ID === id)
      // change map center
      this.searchOption.city = address[0].region.city
      this.aMapManager.getMap().setCity(address[0].region.city)
      let center = this.aMapManager.getMap().getCenter();
      this.mapCenter = [center.lat, center.lng]
    },
    onSearchResult(pois) {
      if (pois.length > 0) {
        let center = {
          lng: pois[0].lng,
          lat: pois[0].lat
        };
        this.mapCenter = [center.lng, center.lat];
        this.circles[0].center = [center.lng, center.lat]
        this.aMapManager.getMap().setZoom(16)
        this.formData.location = center.lng + "," + center.lat
      }
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
      this.dialogTitle = "Update Lift"
      if (res.code === 0) {
        this.formData = res.data.relift;
        this.dialogFormVisible = true;
      }
    },
    async assign(row) {
      // get options
      this.assignFormData = row
      let res = await getUserList({
        page: 1,
        pageSize: 9999,
        companyId: row.installerId
      })
      this.installUserOptions = res.data.list
      res = await getUserList({
        page: 1,
        pageSize: 9999,
        companyId: row.maintainerId
      })
      this.maintainUserOptions = res.data.list      
      res = await getUserList({
        page: 1,
        pageSize: 9999,
        companyId: row.checkerId
      })
      this.checkUserOptions = res.data.list
      res = await getUserList({
        page: 1,
        pageSize: 9999,
        companyId: row.ownerId
      })
      this.ownUserOptions = res.data.list
      // attach previous value
      res = await getUserLiftList({
        page: 1,
        pageSize: 9999,
        liftId: this.assignFormData.ID,
        categoryId: this.installCateogry
      })
      this.assignFormData.installIds = _.map(res.data.list, "userId")
      res = await getUserLiftList({
        page: 1,
        pageSize: 9999,
        liftId: this.assignFormData.ID,
        categoryId: this.maintainCategory
      })
      this.assignFormData.maintainIds = _.map(res.data.list, "userId")
      res = await getUserLiftList({
        page: 1,
        pageSize: 9999,
        liftId: this.assignFormData.ID,
        categoryId: this.checkCategory
      })
      this.assignFormData.checkIds = _.map(res.data.list, "userId")
      res = await getUserLiftList({
        page: 1,
        pageSize: 9999,
        liftId: this.assignFormData.ID,
        categoryId: this.ownerCategory
      })
      this.assignFormData.ownIds = _.map(res.data.list, "userId")
      this.assignDialogVisible = true;
      this.assignDialogTitle = "Assign User"
    },
    async assignBase(categoryId) {
      let ids
      if (categoryId == this.installCateogry) {
        ids = this.assignFormData.installIds
      } else if (categoryId == this.maintainCategory) {
        ids = this.assignFormData.maintainIds
      } else if (categoryId == this.checkCategory) {
        ids = this.assignFormData.checkIds
      } else if (categoryId == this.ownerCategory) {
        ids = this.assignFormData.ownIds
      }
      let res = await getUserLiftList({
        page: 1,
        pageSize: 9999,
        liftId: this.assignFormData.ID,
        categoryId: categoryId
      })
      let targetUserIds = ids
      let preUserIds = _.map(res.data.list, "userId")
      // nothing changed
      if (_.isEqual(targetUserIds, preUserIds)) {
        this.$message({
          type: "warning",
          message: "nothing changed"
        });
        return
      }
      
      // exec del and add 
      let delUserIds = _.difference(preUserIds, targetUserIds)
      let addUserIds = _.difference(targetUserIds, preUserIds)
      if (delUserIds.length > 0) {
        let delIds = _.map(_.filter(res.data.list, function(o) {return _.includes(delUserIds, o.userId)}), "ID")
        res = await deleteUserLiftByIds({
          ids: delIds
        })
      }
      for(const userId of addUserIds) {
        res = await createUserLift({
          userId: userId,
          liftId: this.assignFormData.ID,
          categoryId: categoryId
        })        
      }
      if (res.code === 0) {
          this.$message({
            type: "success",
            message: "setting成功"
          });
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
    closeAssignDialog() {
      this.assignDialogVisible = false;
    },
    async deleteLift(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteLift({ ID: row.ID });
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
      this.formData.owner = {}
      this.formData.maintainer = {}
      this.formData.checker = {}
      this.formData.installer = {}
      this.formData.liftModel = {}
      this.formData.category = {}
      this.formData.address = {}
      this.formData.adDevice = {}
      switch (this.type) {
        case "create":
          res = await createLift(this.formData);
          break;
        case "update":
          console.log(this.formData)
          res = await updateLift(this.formData);
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
      this.dialogTitle = "Create Lift"
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
  .amap-demo {
    height: 300px!important;
    padding: 10px 20px 20px 20px;
  }

  .search-box {
    position: relative;
  }
</style>