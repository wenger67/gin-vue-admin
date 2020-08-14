<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增adDevice表</el-button>
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
    <el-table-column label="在线状态" min-width="50" sortable>
      <template slot-scope="scope">
        <el-tag :type="scope.row.online === '1' ? 'success':'warning'">
          {{scope.row.online|formatBoolean}}
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column label="型号" prop="type.categoryName" sortable min-width="60"></el-table-column>
    <el-table-column label="生产商" prop="factory.fullName" sortable min-width="60"></el-table-column>
    <el-table-column label="出厂时间" sortable min-width="50">
      <template slot-scope="scope">{{scope.row.factoryTime|formatDate}}</template>
    </el-table-column>
    <el-table-column label="安装时间" sortable min-width="50">
      <template slot-scope="scope">{{scope.row.installTime|formatDate}}</template>
    </el-table-column>
    <el-table-column label="状态类别" prop="status.categoryName" sortable min-width="50"></el-table-column>
    <el-table-column label="上次离线时间" sortable min-width="70">
      <template slot-scope="scope">{{scope.row.lastOfflineTime|formatDateTime}}</template>
    </el-table-column>
    <el-table-column label="上次上线时间" sortable min-width="70">
      <template slot-scope="scope">{{scope.row.lastOnlineTime|formatDateTime}}</template>
    </el-table-column>
      <el-table-column label="日期" min-width="70">
        <template slot-scope="scope">{{scope.row.CreatedAt|formatDateTime}}</template>
      </el-table-column>
      <el-table-column label="按钮组" fixed="right" width="200">
        <template slot-scope="scope">
          <el-button @click="updateAdDevice(scope.row)" size="small" type="primary">变更</el-button>
          <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteAdDevice(scope.row)">删除</el-button>
          <!--  TODO owner && config infos         -->
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
        <el-form-item label="型号" prop="typeId">
          <el-select v-model="formData.typeId" placeholder="请选择型号" filterable clearable
                     :style="{width: '100%'}">
            <el-option
                    v-for="item in modalOptions"
                    :key="item.categoryName"
                    :label="item.categoryName"
                    :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="生产商" prop="factoryId">
          <el-select v-model="formData.factoryId" placeholder="请选择生产商" filterable clearable
                     :style="{width: '100%'}">
            <el-option
                    v-for="item in factoryOptions"
                    :key="item.ID"
                    :label="item.fullName"
                    :value="item.ID">
              <span style="float: left">{{ item.fullName }}</span>
              <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.category.categoryName }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="出厂时间" prop="factoryTime">
          <el-date-picker v-model="formData.factoryTime" type="datetime"
                          :style="{width: '100%'}" placeholder="请选择出厂时间" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="安装时间" prop="installTime">
          <el-date-picker v-model="formData.installTime" type="datetime"
                          :style="{width: '100%'}" placeholder="请选择安装时间" clearable></el-date-picker>
        </el-form-item>
        <el-form-item label="设备状态" prop="statusId">
          <el-select v-model="formData.statusId" placeholder="请输入设备状态" filterable clearable
                     :style="{width: '100%'}">
            <el-option
                    v-for="item in statusOptions"
                    :key="item.ID"
                    :label="item.categoryName"
                    :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="负责人" prop="owners">
          <el-select v-model="formData.owners" value-key="ID" placeholder="请选择负责人" multiple filterable clearable
                     :style="{width: '100%'}">
            <el-option
                    v-for="item in ownerOptions"
                    :key="item.ID"
                    :label="item.realName"
                    :value="item">
              <span style="float: left">{{ item.realName }}</span>
              <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.company.fullName }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="配置" prop="configs">
          <el-select v-model="formData.configs" value-key="ID" placeholder="请选择配置" multiple filterable clearable
                     :style="{width: '100%'}">
            <el-option
                    v-for="item in configOptions"
                    :key="item.ID"
                    :label="item.value"
                    :value="item">
              <span style="float: left">{{ item.key }}</span>
              <span style="float: right;color: #2d79a0;font-size: 13px">{{ item.value }}</span>
            </el-option>
          </el-select>
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
    createAdDevice,
    deleteAdDevice,
    deleteAdDeviceByIds,
    updateAdDevice,
    findAdDevice,
    getAdDeviceList
} from "@/api/adDevice";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import {getCategoriesList} from "../../../api/categories";
import {getCompanyList} from "../../../api/company";
import {getUserList} from "../../../api/user";
import {getAdDeviceConfigList} from "../../../api/adDeviceConfig";
import {mapGetters} from "vuex";

export default {
  name: "AdDevice",
  mixins: [infoList],
  data() {
    return {
      listApi: getAdDeviceList,
      dialogFormVisible: false,
      type: "",
      deleteVisible: false,
      modalOptions: [],
      statusOptions: [],
      factoryOptions: [],
      ownerOptions: [],
      configOptions: [],
      formData: {
        typeId: undefined,
        factoryId: undefined,
        factoryTime: null,
        installTime: null,
        statusId: undefined,
        owners: [],
        configs: undefined,
      },
      rules: {
        typeId: [{
          required: true,
          message: '请选择型号',
          trigger: 'change'
        }],
        factoryId: [{
          required: true,
          message: '请选择生产商',
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
        statusId: [{
          required: true,
          message: '请输入设备状态',
          trigger: 'change'
        }],
        owners: [{
          required: true,
          type: 'array',
          message: '请至少选择一个负责人',
          trigger: 'change'
        }],
        configs: [{
          required: true,
          message: '请选择配置',
          trigger: 'change'
        }],
      },
    };
  },
  computed: {
    ...mapGetters("user", ["userInfo"])
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
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10           
        if (this.searchInfo.online===""){
          this.searchInfo.online=null
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
        const res = await deleteAdDeviceByIds({ ids })
        if (res.code === 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          await this.getTableData()
        }
      },
    async updateAdDevice(row) {
      const res = await findAdDevice({ ID: row.ID });
      this.type = "update";
      if (res.code === 0) {
        this.formData = res.data.readDevice;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        typeId: undefined,
        factoryId: undefined,
        factoryTime: null,
        installTime: null,
        statusId: undefined,
        owners: [],
        configs: undefined,
      };
    },
    async deleteAdDevice(row) {
      this.$confirm('this operation is dangerous, continue ?', 'Hint', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      .then(async () => {
        const res = await deleteAdDevice({ ID: row.ID });
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
          res = await createAdDevice(this.formData);
          break;
        case "update":
          res = await updateAdDevice(this.formData);
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
    async getModalOptions() {
      let res = await getCategoriesList({
        page: 1,
        pageSize: 9999,
        ID: 109
      })
      if (res.code === 0) {
        this.modalOptions = res.data.list
      }
    },
    async getStatuOptions() {
      let res = await getCategoriesList({
        page: 1,
        pageSize: 9999,
        ID: 110
      })
      if (res.code === 0) {
        this.statusOptions = res.data.list
      }
    },
    async getFactoryOptions() {
      let res = await getCompanyList({
        page: 1,
        pageSize: 9999,
      })
      if (res.code === 0) {
        this.factoryOptions = res.data.list
      }
    },
    async getOwnerOptions() {
      let res = await getUserList({
        page: 1,
        pageSize: 9999,
      })
      if (res.code === 0) {
        this.ownerOptions = res.data.list
      }
    },
    async getConfigOptions() {
      let res = await getAdDeviceConfigList({
        page: 1,
        pageSize: 9999,
      })
      if (res.code === 0) {
        this.configOptions = res.data.list
      }
    },
  },
  async created() {
    await this.getTableData();
    await this.getConfigOptions()
    await this.getOwnerOptions()
    await this.getFactoryOptions()
    await this.getStatuOptions()
    await this.getModalOptions()
    // TODO websocket onopen event
    if (this.$socket !== undefined) {
      this.$socket.sendObj({
        target: ["server"],
        event: "report",
        data: this.userInfo.uuid
      })
    }
  }
};
</script>

<style>
</style>