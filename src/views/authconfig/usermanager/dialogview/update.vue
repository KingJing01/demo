<template>
  <el-form ref="userFormUpdate" :model="formData" :rules="formRules" size="small" >
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
    <div class="content">
      <div class="power">
        <h4>模块配置</h4>
        <el-row>
          <el-col :span="7">菜单功能</el-col>
          <el-col :span="17">权限名称</el-col>
        </el-row>
        <div v-for="(permissionTop, topIndex) in authData" :key="topIndex">
          <el-row>
            <el-col :span="6">
              <p class="checkGroup" style="width:99%;">
                <el-checkbox :indeterminate="permissionTop.indeterminate" :key="topIndex" v-model="permissionTop.mychecked" :label="permissionTop.permissionId" class="auth_check" @change="onChangeTop(topIndex, permissionTop.permissionId, $event)">{{ permissionTop.permissionName }}</el-checkbox>
            </p></el-col>
            <el-col :span="18">
              <el-checkbox v-for="permissionSon in permissionTop.childrenList" v-model="permissionSon.mychecked" :label="permissionSon.permissionId" :key="permissionSon.permissionId" @change="onChangeSon(topIndex, permissionSon.permissionId, permissionTop.permissionId, $event)">{{ permissionSon.permissionName }}</el-checkbox>
            </el-col>
        </el-row></div>
      </div>
    </div>
  </el-form>
</template>
<script>
import { getUserInfo, getUserPermission } from '@/api/usermanage'
export default {
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
      formRules: {
        TenantName: [{ required: true, trigger: 'blur', message: '公司名称为必填项' },
          { min: 3, max: 40, message: '输入内容最大长度为40', trigger: 'blur' }],
        TenantAddress: [{ max: 40, message: '输入内容最大长度为40', trigger: 'blur' }],
        Email: [{ required: true, trigger: 'blur', message: '邮件为必输项' },
          { pattern: /^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$/, trigger: 'blur', message: '请输入正确格式的邮箱' }],
        LinkPhone: [{ required: true, trigger: 'blur', message: '联系人电话为必输项' },
          { pattern: /^1[345678]\d{9}$/, trigger: 'blur', message: '请输入正确格式的手机号' }]
      }
    }
  },
  created() {
    this.getUserData()
  },
  methods: {
    getUserData() {
      getUserInfo(this.search.tenId).then(response => {
        this.formData = response.Data
        this.data.formData = response.Data
      })
      getUserPermission(this.search).then(response => {
        this.authData = response.Data
        this.data.authData = response.Data
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
    validData() {
      this.$refs.userFormUpdate.validate(valid => {
        this.data.valid = valid
      })
    },
    cancleValid() {
      this.$refs['userFormUpdate'].resetFields()
    },
    cleanData() {
      this.SelectData = [] // 记录选择的系统数据 tab迭代使用
      this.childPerSelect = []// 记录历史选择的套餐及权限信息
      this.checkedApplications = []
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
