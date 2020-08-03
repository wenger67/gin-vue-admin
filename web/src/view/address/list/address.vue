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
    <el-table-column label="region id" prop="regionId" min-width="60"></el-table-column> 
    <el-table-column label="地址 detail" prop="addressName" min-width="60"></el-table-column> 
    <el-table-column label="日期" min-width="60">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    <el-table-column label="user amount" prop="userAmount" min-width="60"></el-table-column> 
    
      <el-table-column label="按钮组" fixed="right" width="200">
        <template slot-scope="scope">
          <el-button @click="updateAddress(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteAddress(scope.row)">确定</el-button>
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
          <el-col :span="8">
            <el-form-item label="省份" prop="province">
              <el-autocomplete
                class="inline-input"
                v-model="formData.province"
                placeholder="Please input province"
                :fetch-suggestions="getProvinces"
                :trigger-on-focus="false"
                @select="handleSelect"
              ></el-autocomplete>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="城市" prop="city">
              <el-autocomplete
                class="inline-input"
                v-model="formData.city"
                placeholder="Please input city"
                :fetch-suggestions="getCities"
                :trigger-on-focus="false"
                @select="handleSelect"
              ></el-autocomplete>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="行政区" prop="district">
              <el-autocomplete
                class="inline-input"
                v-model="formData.district"
                placeholder="Please input district"
                :fetch-suggestions="getDistricts"
                :trigger-on-focus="false"
                @select="handleSelect"
              ></el-autocomplete>
            </el-form-item>
          </el-col>
          <el-col :span="24">
            <el-form-item label="地址" prop="address">
              <el-autocomplete
                class="inline-input"
                v-model="formData.address"
                placeholder="Please input address"
                :fetch-suggestions="getAddresses"
                :trigger-on-focus="false"
                @select="handleSelect"
              ></el-autocomplete>
            </el-form-item> 
          </el-col>
          <el-col :span="24">
            <el-form-item label="标签" prop="tags">
              <el-tag
                :key="tag"
                v-for="tag in dynamicTags"
                closable
                :disable-transitions="false"
                @close="handleClose(tag)">
                {{tag}}
              </el-tag>
              <el-input
                class="input-new-tag"
                v-if="inputVisible"
                v-model="inputValue"
                ref="saveTagInput"
                size="small"
                @keyup.enter.native="handleInputConfirm"
                @blur="handleInputConfirm"
              >
              </el-input>
              <el-button v-else class="button-new-tag" size="small" @click="showInput">+ New Tag</el-button>
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
    findRegion,
    getRegionList
} from "@/api/region";
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Address",
  mixins: [infoList],
  data() {
    return {
      listApi: getAddressList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      dynamicTags: [],
      inputValue: '',
      deleteVisible: false,
      inputVisible: false,
      multipleSelection: [],
      formData: {
        province:null, city:null, district:null,
        regionId:null,addressName:null,userAmount:null,
      },
      rules: {
        region: [{
          required: true,
          message: '请选择区域',
          trigger: 'blur'
        }],
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
        province: [{
          required: true,
          message: '请选择省份',
          trigger: 'change'
        }],
        city: [{
          required: true,
          message: '请选择城市城市',
          trigger: 'change'
        }],
        district: [{
          required: true,
          message: '请选择行政区',
          trigger: 'change'
        }],
      },      
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
        const res = await deleteAddressByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateAddress(row) {
      const res = await findAddress({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.readdress;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          regionId:null,
          addressName:null,
          userAmount:null,
      };
    },
    async deleteAddress(row) {
      this.visible = false;
      const res = await deleteAddress({ ID: row.ID });
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
          res = await createAddress(this.formData);
          break;
        case "update":
          res = await updateAddress(this.formData);
          break;
        default:
          res = await createAddress(this.formData);
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
    },
    handleClose(tag) {

    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick(_ => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {

    },
    async getRegions() {

    },
    async getAddresses() {

    },
    async getProvinces(queryString, cb) {
      // clear right input component
      this.formData.city = ""
      this.formData.district = ""

      let result = []
      const params = {
        page: 1,
        pageSize: 100,
        province: queryString,
        groupKey: "province"
      }
      let res = await getRegionList(params)
      if (res.code == 0) {
        res.data.list.forEach(element => {
          // should have value field
          result.push({value: element.province})
        })
      }
      cb(result)
    },
    async getCities(queryString, cb) {
      // clear right input component
      this.formData.district = ""
      let result = []
      const params = {
        page: 1,
        pageSize: 100,
        province: this.formData.province,
        city: queryString,
        groupKey: "city"
      }
      let res = await getRegionList(params)
      if (res.code == 0) {
        res.data.list.forEach(element => {
          result.push({value: element.city})
        })
      }
      cb(result);
    },
    async getDistricts(queryString, cb) {
      let result = []
      const params = {
        page: 1,
        pageSize: 100,
        city: this.formData.city,
        district: queryString,
        groupKey: "district"
      }
      let res = await getRegionList(params)
      if (res.code == 0) {
        res.data.list.forEach(element => {
          result.push({value: element.district})
        })
      }
      cb(result);
    },
    handleSelect(item) {
      this.province = item
    }    
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
  .el-tag + .el-tag {
    margin-left: 10px;
  }
  .button-new-tag {
    margin-left: 10px;
    height: 32px;
    line-height: 30px;
    padding-top: 0;
    padding-bottom: 0;
  }
  .input-new-tag {
    width: 90px;
    margin-left: 10px;
    vertical-align: bottom;
  }
</style>