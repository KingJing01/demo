<template>
  <el-tabs v-model="activeName" class="auth_tab" @tab-click="handleClick" >
    <el-tab-pane label="角色配置" name="first"><RoleManage ref="mychildFirst"/></el-tab-pane>
    <el-tab-pane label="用户配置" name="second"><UserManage ref="mychildSecond"/></el-tab-pane>
</el-tabs></template>
<script>
import RoleManage from './role/page'
import UserManage from './user/page'
export default {
  components: { RoleManage, UserManage },
  data() {
    return {
      activeName: 'first'
    }
  },
  methods: {
    handleClick(tab, event) {
    /*   if (tab.name === 'first') {
        this.$refs.mychildFirst.getList()
      } else {
        this.$refs.mychildSecond.getList()
      } */
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
