<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">      
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增region</el-button>
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
    <el-table-column label="ID" prop="ID" sortable min_width="60"></el-table-column>
    <el-table-column label="省份" prop="province" sortable min_width="60"></el-table-column>
    <el-table-column label="城市" prop="city" sortable min_width="60"></el-table-column>
    <el-table-column label="行政区" prop="district" sortable min_width="60"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
      <el-table-column label="按钮组" fixed="right" witdh="200">
        <template slot-scope="scope">
          <el-button @click="updateRegion(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteRegion(scope.row)">删除</el-button>
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
    createRegion,
    deleteRegion,
    deleteRegionByIds,
    updateRegion,
    findRegion,
    getRegionList
} from "@/api/region";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Region",
  mixins: [infoList],
  data() {
    return {
      listApi: getRegionList,
      dialogFormVisible: false,
      visible: false,
      dialogTitle: "",
      type: "",
      deleteVisible: false,
      multipleSelection: [],
      formData: {
        province:null,city:null,district:null,
      },
      rules: {
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
      }
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
      const res = await deleteRegionByIds({ ids })
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        this.deleteVisible = false
        await this.getTableData()
      }
    },
    async updateRegion(row) {
      const res = await findRegion({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.reregion;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          province:null,
          city:null,
          district:null,
      };
    },
    async deleteRegion(row) {
      this.$confirm('this operation is dangerous, continue ? ', 'hint', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteRegion(row)
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: 'delete success!'
          })
          await this.getTableData()
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
      this.$refs.elForm.validate(async valid => {
        if (valid) {
          let res;
          switch (this.type) {
            case "create":
              res = await createRegion(this.formData);
              break;
            case "update":
              res = await updateRegion(this.formData);
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
        }
      })
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
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
      if (res.code === 0) {
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
      if (res.code === 0) {
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
      if (res.code === 0) {
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
</style>