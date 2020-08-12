<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增address表</el-button>
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
      <el-table-column label="ID" prop="ID" min-width="30"></el-table-column>
      <el-table-column label="Province" prop="region.province" min-width="60"></el-table-column>
      <el-table-column label="City" prop="region.city" min-width="60"></el-table-column>
      <el-table-column label="District" prop="region.district" min-width="60"></el-table-column> 
      <el-table-column label="地址 detail" prop="addressName" min-width="60"></el-table-column> 
      <el-table-column label="user amount" prop="userAmount" min-width="60"></el-table-column> 
      <el-table-column label="Tags" prop="tags" min-width="60">
        <template slot-scope="scope">
          <el-tag :key="tag.ID" v-for="tag in scope.row.tags">{{tag.categoryName}}</el-tag>        
        </template>
      </el-table-column> 
      <el-table-column label="日期" min-width="100">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" width="200">
        <template slot-scope="scope">
          <el-button @click="updateAddress(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteAddress(scope.row)">删除</el-button>
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
          <el-col :span="8">
            <el-form-item label="省份" prop="province">
              <el-select @change="getCityOptions" v-model="formData.region.province" filterable clearable placeholder="Please select province">
                <el-option v-for="item in provinceOptions"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="城市" prop="city">
              <el-select @change="getDistrictOptions" v-model="formData.region.city"  filterable clearable :disabled="cityOptionDisable" placeholder="Please select city">
                <el-option v-for="item in cityOptions"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="行政区" prop="district">
              <el-select @change="handleDistrictChange" v-model="formData.region.district" filterable clearable :disabled="districtOptionDisable" placeholder="Please select district">
                <el-option v-for="item in districtOptions"
                  :key="item.value"
                  :label="item.value"
                  :value="item.value">
                </el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="address" prop="address">
              <el-amap-search-box v-model="formData.addressName" class="search-box" :search-option="searchOption" :on-search-result="onSearchResult"></el-amap-search-box>
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
          <el-col :span="12">
            <el-form-item label="Amount" prop="userAmount">
              <el-input-number v-model="formData.userAmount" :min="1" label="Amount"></el-input-number>
            </el-form-item> 
          </el-col>
          <el-col :span="12">
            <el-form-item label="标签" prop="tags">
              <el-select v-model="formData.tags" value-key="ID" multiple placeholder="Please select tags">
                <el-option v-for="item in tagOptions"
                  :key="item.value"
                  :label="item.value"
                  :value="item.data">
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
    createAddress,
    deleteAddress,
    deleteAddressByIds,
    updateAddress,
    findAddress,
    getAddressList
} from "@/api/address";
import {
    getRegionList
} from "@/api/region";
import {
    getCategoriesList
} from "@/api/categories";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import VueAMap from "vue-amap";
let aMapManager = new VueAMap.AMapManager();

export default {
  name: "Address",
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
            click: () => {
              alert('click');
            }
          }
        }
      ],
      searchOption: {
        city: '上海',
        citylimit: true
      },
      mapCenter: [121.59996, 31.197646],
      listApi: getAddressList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      inputVisible: false,
      multipleSelection: [],
      provinceOptions: [],
      cityOptions: [],
      districtOptions: [],
      tagOptions: [],
      dynamicTags: [],
      cityOptionDisable: true,
      districtOptionDisable: true,
      tagSubjectId: 107,
      formData: {
        regionId: null, addressName:null, region: {province:null, city: null, district:null},
        location: null, userAmount: null, tags:[]
      },
      rules: {
        address: [{
          required: true,
          message: '请选择地址',
          trigger: 'change'
        }],
        tags: [{
          required: true,
          message: '请选择标签',
          trigger: 'change'
        }],
        userAmount: [{
          required: true,
          message: 'Please input amount',
          trigger: 'change'
        }],
        province: [{
          required: true,
          message: '请选择省份',
          trigger: 'blur'
        }],
        city: [{
          required: true,
          message: '请选择城市城市',
          trigger: 'blur'
        }],
        district: [{
          required: true,
          message: '请选择行政区',
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
    handleClosePopover(index) {
      this.$refs[`popover-${index}`].doClose()
    },
    async onDelete() {
      const ids = []
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteAddressByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        await this.getTableData()
      }
    },
    async updateAddress(row) {
      const res = await findAddress({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.readdress

        // TODO fill the origin data
        this.cityOptionDisable = false
        this.districtOptionDisable = false
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        regionId: null, addressName:null, region: {province:null, city: null, district:null},
        location: null, userAmount: null, tags:[]
      };
      this.cityOptions = []
      this.districtOptions = []
    },
    async deleteAddress(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteAddress({ ID: row.ID });
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
      // void create region one more time
      let tmp = this.formData
      delete tmp.region
      switch (this.type) {
        case "create":
          res = await createAddress(tmp);
          break;
        case "update":
          res = await updateAddress(tmp);
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
    async getProvinceOptions() {
      // clear right input component
      this.formData.region.city = ""
      this.formData.region.district = ""
      this.cityOptionDisable = true
      this.districtOptionDisable = true
      const params = {
        page: 1,
        pageSize: 100,
        province: "",
        groupKey: "province"
      }
      let res = await getRegionList(params)
      if (res.code === 0) {
        res.data.list.forEach(element => {
          // should have value field
          this.provinceOptions.push({value: element.province})
        })
      }
    },
    async getCityOptions() {
      // clear right input component
      this.formData.region.district = ""
      this.cityOptionDisable = false
      this.districtOptionDisable = true

      const params = {
        page: 1,
        pageSize: 100,
        province: this.formData.region.province,
        city: "",
        groupKey: "city"
      }
      let res = await getRegionList(params)
      if (res.code === 0) {
        res.data.list.forEach(element => {
          this.cityOptions.push({value: element.city})
        })
      }
    },
    async getDistrictOptions() {
      this.districtOptionDisable = false
      const params = {
        page: 1,
        pageSize: 100,
        city: this.formData.region.city,
        district: "",
        groupKey: "district"
      }
      // change map center
      this.searchOption.city = params.city
      this.aMapManager.getMap().setCity(params.city)
      let center = this.aMapManager.getMap().getCenter();
      this.mapCenter = [center.lat, center.lng]

      this.districtOptions = []
      let res = await getRegionList(params)
      if (res.code === 0) {
        res.data.list.forEach(element => {
          this.districtOptions.push({value: element.district, regionId: element.ID})
        })
      }
    },
    handleDistrictChange(val) {
      let tmp = this.districtOptions.filter(item => item.value === val);
      this.formData.regionId = tmp[0].regionId;
    },
    async getTags() {
      const params = {
        page: 1,
        pageSize: 100,
        ID: this.tagSubjectId
      }
      let res = await getCategoriesList(params);
      if (res.code === 0) {
        res.data.list.forEach(element => {
          this.tagOptions.push({value: element.categoryName, data:element})
        })
      }
    },
    onSearchResult(pois) {
      if (pois.length > 0) {
        /*pois.forEach(poi => {
          let {lng, lat} = poi;
          lngSum += lng;
          latSum += lat;
        });*/
        let center = {
          lng: pois[0].lng,
          lat: pois[0].lat
        };
        this.mapCenter = [center.lng, center.lat];
        this.circles[0].center = [center.lng, center.lat]
        this.aMapManager.getMap().setZoom(16)
        this.formData.location = center.lng + "," + center.lat
        this.formData.addressName = pois[0].address + pois[0].name
      }
    }
  },
  async created() {
    await this.getTableData()
    await this.getTags()
    await this.getProvinceOptions()    
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