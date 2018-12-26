<template>
  <el-tabs v-model="activeName" class="auth_tab" @tab-click="handleClick" >
    <el-tab-pane label="系统配置" name="first"><SysConfig ref="mychildFirst"/></el-tab-pane>
    <el-tab-pane label="基础权限" name="second"><BasePermission ref="mychildSecond"/></el-tab-pane>
    <el-tab-pane label="用户套餐" name="third"><UserPackage ref="mychildThird"/></el-tab-pane>
    <el-tab-pane label="平台用户管理" name="fourth"><UserManager ref="mychildFourth"/></el-tab-pane>
  </el-tabs>
</template>
<script>
import BasePermission from './basepermissions/page'
import SysConfig from './sysconfig/page'
import UserManager from './usermanager/page'
import UserPackage from './setmeal/page'
export default {
  components: { BasePermission, SysConfig, UserManager, UserPackage },
  data() {
    return {
      activeName: 'first'
    }
  },
  methods: {
    handleClick(tab, event) {
      if (tab.name === 'first') {
        this.$refs.mychildFirst.getList()
      } else if (tab.name === 'second') {
        this.$refs.mychildSecond.getList()
      } else if (tab.name === 'third') {
        this.$refs.mychildThird.getList()
      } else if (tab.name === 'fourth') {
        this.$refs.mychildFourth.getList()
      }
    }
  }
}

</script>
<style rel="stylesheet/scss" lang="scss" scoped>
.auth_tab {
  margin-top: 20px;
  margin-left: 20px
}

</style>
