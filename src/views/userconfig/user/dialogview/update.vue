<template>
  <el-form ref="userForm" :model="form" size="small" label-width="100px">
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
      prop="SysCode"
    >
      <el-select v-model="form.SysCode" placeholder="请选择" disabled>
        <el-option
          v-for="item in options"
          :key="item.SysCode"
          :label="item.SysName"
          :value="item.SysCode"/>
      </el-select>
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
      radioData: {}
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
    // 单选按钮的修改事件
    handleRadioChange(name, code) {
      this.form.RoleCode = code
      this.data.formData = this.form
    }
  }
}
</script>

