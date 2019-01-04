<template>
  <el-form ref="mealForm" :model="form" :rules="formRules" size="small" label-width="100px">
    <el-form-item
      :label-width="formLabelWidth"
      label="登录名"
      prop="UserName"
    >
      <el-input
        v-model="form.UserName"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      :label-width="formLabelWidth"
      label="手机号"
      prop="PhoneNumber"
    >
      <el-input
        v-model="form.PhoneNumber"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      :label-width="formLabelWidth"
      label="邮箱"
      prop="EmailAddress"
    >
      <el-input
        v-model="form.EmailAddress"
        auto-complete="off"
      />
    </el-form-item>
    <el-form-item
      :label-width="formLabelWidth"
      label="系统名称"
      prop="SysCode"
    >
      <el-select v-model="form.SysCode" placeholder="请选择">
        <el-option
          v-for="item in options"
          :key="item.SysCode"
          :label="item.SysName"
          :value="item.SysCode"/>
      </el-select>
    </el-form-item>
  </el-form>
</template>
<script>
import { sysDataSelect } from '@/api/sysconfig'
export default {
  data: function() {
    return {
      options: {},
      form: {},
      formRules: {
        SysCode: [{ required: true, trigger: 'change', message: '系统为必填项' }],
        EmailAddress: [{ required: true, trigger: 'change', message: '邮箱为必填项' }],
        PhoneNumber: [{ required: true, trigger: 'change', message: '手机号为必填项' }],
        UserName: [{ required: true, trigger: 'blur', message: '用户名为必填项' }, { max: 20, message: '输入内容最大长度为20', trigger: 'blur' }]
      }
    }
  }, created() {
    this.getSysList()
  },
  methods: {
    // 获取系统下拉数据
    getSysList() {
      sysDataSelect().then(response => {
        this.options = response.Data
      })
    }
  }
}
</script>

