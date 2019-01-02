<template>
  <el-form ref="userForm" :model="formData" :rules="formRules" size="small" >
    <el-form-item
      :label-width="formLabelWidth"
      label="公司名称"
      prop="TenantName"
    >
      <el-input
        v-model="formData.TenantName"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      :label-width="formLabelWidth"
      label="公司地址"
      prop="TenantAddress"
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
        prop="LinkPhone"
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
        prop="Email"
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
      <el-tabs type="card" @tab-remove="handleTabRemove" >
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
      childPerSelect: [], // 记录历史选择的套餐及权限信息
      formRules: {
        TenantName: [{ required: true, trigger: 'blur', message: '公司名称为必填项' },
          { min: 3, max: 40, message: '输入内容最大长度为40', trigger: 'blur' }],
        TenantAddress: [{ max: 40, message: '输入内容最大长度为40', trigger: 'blur' }],
        Email: [{ required: true, trigger: 'blur', message: '邮件为必输项' },
          { pattern: /^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$/, trigger: 'blur', message: '请输入正确格式的邮箱' }],
        LinkPhone: [{ required: true, trigger: 'blur', message: '联系人电话为必输项' },
          { pattern: /^1[345678]\d{9}$/, trigger: 'blur', message: '请输入正确格式的手机号' }],
        SysCode: [{ required: true, trigger: 'blur', message: '系统为必选项' }]
      }
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
      console.log('转换前' + this.childPerSelect)
      debugger
      var tempSelectData = []
      for (const i of val) {
        // 添加动态tab
        for (const option of this.SysOptions) {
          if (option.SysCode === i) {
            var result = {}
            result.SysName = option.SysName + '权限配置'
            result.SysCode = option.SysCode
            tempSelectData.push(result)
          }
        }
        if (val.length < this.childPerSelect.length) {
          var tempChildSelect = []
          for (const child of this.childPerSelect) {
            debugger
            if (child.sysCode === i) {
              var temp = {}
              temp.data = child.data
              temp.sysCode = child.sysCode
              tempChildSelect.push(temp)
            }
          }
          this.childPerSelect = tempChildSelect
        }
      }
      console.log('转换后' + this.childPerSelect)
      this.SelectData = tempSelectData
      this.data.authData = this.childPerSelect
      this.data.formData = this.formData
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
    },
    handleTabRemove(val) {
      console.log(val)
    },
    validData() {
      this.$refs.userForm.validate(valid => {
        this.data.valid = valid
      })
    },
    cancleValid() {
      this.$refs['userForm'].resetFields()
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
