<template>
  <el-form :model="formData" size="small">
    <el-form-item
      :label-width="formLabelWidth"
      label="公司名称"
    >
      <el-input
        v-model="formData.TenantName"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      :label-width="formLabelWidth"
      label="公司地址"
    >
      <el-input
        v-model="formData.TenantAddress"
        auto-complete="off"
      />
    </el-form-item>
    <el-row>
      <el-col :span="12"><div class="grid-content bg-purple"> <el-form-item
        :label-width="formLabelWidth"
        label="组织机构代码"
      >
        <el-input
          v-model="formData.OrganizationCode"
          auto-complete="off"
        />
      </el-form-item></div>
      </el-col>
      <el-col :span="12"><el-form-item
        :label-width="formLabelWidth"
        label="营业执照"
      >
        <el-input
          v-model="formData.BusinessLisenceUrl"
          auto-complete="off"
        />
      </el-form-item></el-col>
    </el-row>
    <el-row>
      <el-col :span="12"> <el-form-item
        :label-width="formLabelWidth"
        label="公司税号"
      >
        <el-input
          v-model="formData.TaxFileNumber"
          auto-complete="off"
        />
      </el-form-item>
      </el-col>
      <el-col :span="12"> <el-form-item
        :label-width="formLabelWidth"
        label="公司联系人"
      >
        <el-input
          v-model="formData.LinkMan"
          auto-complete="off"
        />
      </el-form-item></el-col>
    </el-row>
    <el-row>
      <el-col :span="12"> <el-form-item
        :label-width="formLabelWidth"
        label="联系电话"
      >
        <el-input
          v-model="formData.LinkPhone"
          auto-complete="off"
        />
      </el-form-item>
      </el-col>
      <el-col :span="12"><el-form-item
        :label-width="formLabelWidth"
        label="联系邮箱"
      >
        <el-input
          v-model="formData.Email"
          auto-complete="off"
        />
      </el-form-item></el-col>
    </el-row>
    <el-form-item
      :label-width="formLabelWidth"
      label="系统名称"
    >
      <template>
        <el-checkbox-group v-model="checkedApplications" @change="handlecheckedAppChange">
          <el-checkbox v-for="(sys, index) in SysOptions" :label="sys.SysCode" :key="index">{{ sys.SysName }}</el-checkbox>
        </el-checkbox-group>
      </template>
    </el-form-item>
    <template>
      <el-tabs type="card">
        <el-tab-pane v-for="(sys, index) in SelectData" :label="sys.SysName" :key="index"><PermissionPage/></el-tab-pane>
      </el-tabs>
    </template>
  </el-form>
</template>
<script>
import { sysDataSelect } from '@/api/sysconfig'
import PermissionPage from './permission'
export default {
  components: { PermissionPage },
  props: {
    data: {
      type: Object,
      default: null
    }
  },
  data() {
    return {
      formLabelWidth: '100px',
      search: {
        tenId: this.data.id,
        sysCode: this.data.sysCode
      },
      formData: {},
      authData: [],
      SysOptions: [], // 系统checkbox数据
      checkedApplications: [],
      tabModel: '',
      editableTabs: [],
      SelectData: [], // 记录选择的系统数据 tab迭代使用
      lastSelect: ''
    }
  },
  created() {
    this.getSysData()
  },
  methods: {
    // 获取企业checkbox和套餐的radio数据
    getSysData() {
      sysDataSelect().then(response => {
        this.SysOptions = response.Data
      })
    },
    onChangeTop(index, topId, e) { // 父级change事件
      this.authData[index].mychecked = e// 父级勾选后，子级全部勾选或者取消
      if (e === false) this.authData[index].indeterminate = false // 去掉不确定状态
      var childrenArray = this.authData[index].childrenList
      if (childrenArray) {
        for (var i = 0, len = childrenArray.length; i < len; i++) { childrenArray[i].mychecked = e }
      }
    },
    onChangeSon(topIndex, sonId, topId, e) { // 子级change事件
      var childrenArray = this.authData[topIndex].childrenList
      var tickCount = 0
      var unTickCount = 0
      var len = childrenArray.length
      for (var i = 0; i < len; i++) {
        if (sonId === childrenArray[i].permissionId) childrenArray[i].mychecked = e
        if (childrenArray[i].mychecked === true) tickCount++
        if (childrenArray[i].mychecked === false) unTickCount++
      }
      if (tickCount === len) { // 子级全勾选
        this.authData[topIndex].mychecked = true
        this.authData[topIndex].indeterminate = false
      } else if (unTickCount === len) { // 子级全不勾选
        this.authData[topIndex].mychecked = false
        this.authData[topIndex].indeterminate = false
      } else {
        this.authData[topIndex].mychecked = true
        this.authData[topIndex].indeterminate = true // 添加不确定状态
      }
    },
    // 系统信息选择事件
    handlecheckedAppChange(val) {
      this.SelectData = []
      for (const i of val) {
        for (const option of this.SysOptions) {
          if (option.SysCode === i) {
            var result = {}
            result.SysName = option.SysName + '权限配置'
            result.SysCode = option.SysCode
            this.SelectData.push(result)
          }
        }
      }
    }
  }
}
</script>
<style rel="stylesheet/scss" lang="scss">
.el-row {
  &:last-child {
    margin-bottom: 0;
  }
}
.el-checkbox{
  margin:2% 5%
}
</style>
