<template>
  <div style="height:400px;overflow-y:scroll">
    <div> 套餐:
      <el-radio-group v-model="radio" @change="handleRadioChange">
        <el-radio v-for="data in radioData" :key="data.Name" :label="data.Name">{{ data.DisplayName }}</el-radio>
      </el-radio-group>
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
</div></div></template>
<script>
import { getSetMealRadio } from '@/api/setmeal'
import { getPerInfoBySysCodeUpdate } from '@/api/permission'
export default {
  props: {
    sysCode: {
      type: String,
      default: null
    }
  },
  data() {
    return {
      radio: '',
      authData: [],
      radioData: []
    }
  },
  created() {
    getSetMealRadio(this.sysCode).then(response => {
      this.radioData = response.Data
    })
  },
  methods: {
    // 单选按钮的触发事件
    handleRadioChange(val) {
      getPerInfoBySysCodeUpdate(this.sysCode, val).then(response => {
        this.authData = response.Data
        var resp = {}
        resp.sysCode = this.sysCode
        resp.data = this.authData
        this.$emit('listen', resp)
      })
    },
    onChangeTop(index, topId, e) { // 父级change事件
      this.authData[index].mychecked = e// 父级勾选后，子级全部勾选或者取消
      if (e === false) this.authData[index].indeterminate = false // 去掉不确定状态
      var childrenArray = this.authData[index].childrenList
      if (childrenArray) {
        for (var i = 0, len = childrenArray.length; i < len; i++) { childrenArray[i].mychecked = e }
      }
      var resp = {}
      resp.sysCode = this.sysCode
      resp.data = this.authData
      this.$emit('listen', resp)
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
      var resp = {}
      resp.sysCode = this.sysCode
      resp.data = this.authData
      this.$emit('listen', resp)
    }
  }
}
</script>

