<template>
  <div>
    <div> <el-radio v-model="radio" label="1">备选项</el-radio>
    <el-radio v-model="radio" label="2">备选项</el-radio></div>
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
                <el-checkbox :indeterminate="permissionTop.indeterminate" :key="topIndex" v-model="permissionTop.mychecked" :label="permissionTop.permissionId" :disabled="type=='detail'?true:false" class="auth_check" @change="onChangeTop(topIndex, permissionTop.permissionId, $event)">{{ permissionTop.permissionName }}</el-checkbox>
            </p></el-col>
            <el-col :span="18">
              <el-checkbox v-for="permissionSon in permissionTop.childrenList" v-model="permissionSon.mychecked" :label="permissionSon.permissionId" :key="permissionSon.permissionId" :disabled="type=='detail'?true:false" @change="onChangeSon(topIndex, permissionSon.permissionId, permissionTop.permissionId, $event)">{{ permissionSon.permissionName }}</el-checkbox>
            </el-col>
        </el-row></div>
      </div>
    </div>
</div></template>
<script>
export default {

}
</script>

