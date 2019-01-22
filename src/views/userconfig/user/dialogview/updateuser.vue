<template>
  <el-form ref="userForm" :model="form" size="small" label-width="100px">
    <el-form-item
      label="登录名"
      prop="UserName"
    >
      <el-input
        v-model="form.UserName"
        :disabled="true"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      label="手机号"
      prop="PhoneNumber"
    >
      <el-input
        v-model="form.PhoneNumber"
        :disabled="true"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      label="邮箱"
      prop="EmailAddress"
    >
      <el-input
        v-model="form.EmailAddress"
        :disabled="true"
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
          <el-radio-group v-model="form.RoleCode">
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
      options: {},
      form: {},
      radioData: {},
      checkedApplications: []
    }
  },
  created() {
    this.getUserInfoById()
  },
  methods: {
    // 获取系统下拉数据
    getSysList(sysCode) {
      sysDataSelect(1).then(response => {
        var data = response.Data
        data.splice(data.findIndex(item => item.SysCode === sysCode), 1)
        this.options = data
      })
    },
    // 获取用户的详情信息
    getUserInfoById() {
      getUserInfoById(this.userId).then(response => {
        this.form = response.Data
        this.getSysList(response.Data.SysCode)
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
    }
  }
}
</script>

