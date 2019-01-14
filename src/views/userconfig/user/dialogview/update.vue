<template>
  <el-form ref="userForm" :model="form" size="small" label-width="100px" >
    <el-form-item
      label="登录名"
      prop="UserName"
    >
      <el-input
        v-model="form.UserName"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      label="手机号"
      prop="PhoneNumber"
    >
      <el-input
        v-model="form.PhoneNumber"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      label="邮箱"
      prop="EmailAddress"
    >
      <el-input
        v-model="form.EmailAddress"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      label="系统名称"
    >
      <template>
        <el-checkbox-group v-model="checkedApplications" @change="handlecheckedAppChange">
          <el-checkbox v-for="(sys, index) in options" :label="sys.SysCode" :key="index">{{ sys.SysName }}</el-checkbox>
        </el-checkbox-group>
      </template>
    </el-form-item>

    <template>
      <div v-for="(radio, topIndex) in radioData" :key="topIndex">
        <el-form-item :label="radio.name" prop="resource">
          <el-radio-group v-model="radio.data">
            <el-radio v-for="child in radio.childrenList" :label="child.childCode" :key="child.childCode" @change="handleRadioChange(radio.key,child.childCode)">{{ child.childName }}</el-radio>
          </el-radio-group>
        </el-form-item>

    </div></template>
  </el-form>
</template>
<script>
import { sysDataSelect } from '@/api/sysconfig'
import { getRoleDataBySysCodes } from '@/api/role'
import { getUserInfoById } from '@/api/user'
export default {
  props: {
    data: {
      type: Object,
      default: null
    }
  },
  data: function() {
    return {
      userId: this.data.Id,
      form: {},
      radio: '',
      checkedApplications: [],
      radioData: [],
      options: {},
      selectData: new Map(),
      formRules: {
        EmailAddress: [{ required: true, trigger: 'change', message: '邮箱为必填项' }],
        PhoneNumber: [{ required: true, trigger: 'change', message: '手机号为必填项' }],
        UserName: [{ required: true, trigger: 'blur', message: '用户名为必填项' }, { max: 20, message: '输入内容最大长度为20', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getSysList()
    this.getUserInfoById()
  },
  methods: {
    // 获取系统下拉数据
    getSysList() {
      sysDataSelect().then(response => {
        this.options = response.Data
      })
    },
    // 获取用户的详情信息
    getUserInfoById() {
      getUserInfoById(this.userId).then(response => {
        this.form = response.Data
        getRoleDataBySysCodes(this.form.SysCode).then(response => {
          this.radioData = response.Data
        })
      })
    },
    cancleValid() {
      this.$refs['userForm'].resetFields()
    },
    validData() {
      this.$refs.userForm.validate(valid => {
        this.data.valid = valid
      })
    },
    // 系统修改刷新数据
    handlecheckedAppChange(val) {
      getRoleDataBySysCodes(val.join(',')).then(response => {
        this.radioData = response.Data
      })
    },
    // 单选按钮的修改事件
    handleRadioChange(name, code) {
      this.selectData.set(name, code)
      var sysCode = new Array(0)
      var roleId = new Array(0)
      this.selectData.forEach(function(item, key, mapObj) {
        sysCode.push(key)
        roleId.push(item)
      })
      this.form.roleIds = roleId
      this.form.sysCodes = sysCode
      this.data.formData = this.form
    }

  }
}
</script>

