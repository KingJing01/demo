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
        <el-tab-pane v-for="(sys, index) in SelectData" :label="sys.SysName" :key="index"><PermissionPage :sys-code="sys.SysCode" @listen="getChildEvent"/></el-tab-pane>
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
      formData: {},
      authData: [],
      SysOptions: [], // 系统checkbox数据
      checkedApplications: [],
      tabModel: '',
      editableTabs: [],
      SelectData: [], // 记录选择的系统数据 tab迭代使用
      childPerSelect: []// 记录历史选择的套餐及权限信息
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
    },
    getChildEvent(val) {
      if (this.childPerSelect.length === 0) {
        this.childPerSelect.push(val)
      } else {
        var data = this.childPerSelect.filter(function(item) {
          return item.sysCode === val.sysCode
        })
        if (data.length === 0) {
          this.childPerSelect.push(val)
        } else {
          for (const index in this.childPerSelect) {
            if (this.childPerSelect[index].sysCode === val.sysCode) {
              this.childPerSelect[index].data = val.data
              break
            }
          }
        }
      }
      this.data.authData = this.childPerSelect
      this.data.formData = this.formData
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
