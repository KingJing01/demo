<template>
  <div>
    <!-- 查询 form start -->
    <el-form
      :inline="true"
      :model="search"
      class="demo-form-inline"
    >
      <el-form-item label="套餐名称">
        <el-input
          v-model="search.setMealName"
          placeholder="套餐名称"
        />
      </el-form-item>
      <el-form-item label="系统名称">
        <el-input
          v-model="search.sysName"
          placeholder="系统名称"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          @click="onSubmit"
        >查询</el-button>
        <el-button @click="onReset">重置</el-button>
      </el-form-item>
    </el-form>
    <!-- 查询 form end -->
    <!-- 基础权限列表  start -->
    <el-table
      :data="tableData"
      style="width: 90%"
      border
    >
      <el-table-column
        type="selection"
        width="auto"
        align="center"/>
      <el-table-column
        prop="SetMealCode"
        label="套餐编码"
        align="center"
      />
      <el-table-column
        prop="SetMealName"
        label="套餐名称"
        align="center"
      />
      <el-table-column
        prop="SysName"
        label="系统名称"
        align="center"
      />
      <el-table-column
        prop="PermissionText"
        label="操作名称"
        align="center"

      />
      <el-table-column
        prop="Id"
        label="操作"
        align="center"
      >
        <template slot-scope="scope">
          <el-button
            type="text"
            size="small"
            @click="handleClick(scope.row)"
          >编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
    <!-- 基础权限信息列表  end -->
    <!-- 分页控件  start -->
    <div class="block">
      <el-pagination
        :page-size="20"
        :pager-count="11"
        :total="1000"
        layout="prev, pager, next"/>
    </div>
    <el-row id="action_line">
      <el-button @click="dialogFormVisible = true">新增套餐</el-button>
      <el-button type="primary">删除套餐</el-button>
    </el-row>

    <!-- 分页控件  end -->
    <!-- 弹出层 信息录入和修改  start -->
    <el-dialog
      :visible.sync="dialogFormVisible"
      title="套餐配置"
      width="40%"
    >
      <el-form :model="form">
        <el-form-item
          :label-width="formLabelWidth"
          label="套餐名"
        >
          <el-input
            v-model="form.setMealName"
            autocomplete="off"
          />
        </el-form-item>

        <el-form-item
          :label-width="formLabelWidth"
          label="菜单编码"
        >
          <el-select v-model="sysInfoSelect" placeholder="请选择" @change="changeSysSelect">
            <el-option
              v-for="item in options"
              :key="item.SysCode"
              :label="item.SysName"
              :value="item.SysCode"/>
          </el-select>
        </el-form-item>
      </el-form>
      <template>
        <div class="content">
          <div class="power">
            <h4>权限设置</h4>
            <el-row>
              <el-col :span="7">菜单功能</el-col>
              <el-col :span="17">权限名称</el-col>
            </el-row>
            <div v-for="(permissionTop, topIndex) in authData" :key="topIndex" >
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
      </div></template>
      <div
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="handleCancle">取 消</el-button>
        <el-button
          type="primary"
          @click="saveData"
        >确 定</el-button>
      </div>
    </el-dialog>
    <!-- 弹出层 信息录入和修改  end -->
  </div>
</template>
<script>
import { getSetMealList } from '@/api/setmeal'
import { sysDataSelect } from '@/api/sysconfig'
import { getPerInfoBySysCode } from '@/api/permission'
export default {
  data() {
    return {
      tableData: [],
      search: {
        setMealName: '',
        sysName: '',
        pageSize: 10,
        offset: 0
      },
      form: {
        id: '',
        setMealName: '',
        menuCode: '',
        sysUrl: '',
        IsValid: true
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '120px',
      dialogInfoVisable: false,
      insertAct: true,
      radio1: '',
      options: [],
      sysInfoSelect: '',
      authData: [],
      activeName: 'accountManage',
      people: '',
      phoneNum: '',
      isIndeterminate: true
    }
  },
  created() {
    this.getList()
    this.getSysData()
  },
  methods: {
    // 列表查询
    onSubmit() {
      this.getList()
    },
    // 编辑事件
    handleClick(row) {
      this.dialogFormVisible = true
      this.form.sysCode = row.SysCode
      this.form.sysName = row.SysName
      this.form.id = row.Id
      this.form.sysUrl = row.SysUrl
      this.insertAct = false
      if (row.IsValid === 0) {
        this.form.IsValid = true
      } else {
        this.form.IsValid = false
      }
    },
    // 获取列表数据
    getList() {
      getSetMealList(this.search).then(response => {
        this.tableData = response.Data.list
      })
    },
    handleClose(done) {
      this.form.sysName = ''
      this.form.sysCode = ''
      this.form.IsValid = true
      this.dialogInfoVisable = false
      done()
    },

    // 重置按钮
    onReset() {
      this.search.setMealName = ''
      this.search.sysName = ''
      this.search.pageSize = 10
      this.search.offset = 0
      this.dialogInfoVisable = false
      this.getList()
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`)
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`)
    },
    // 保存系统信息
    saveData() {
      if (this.dialogInfoVisable === false) {
        if (this.insertAct === true) {
          console.log('修改信息')
        } else {
          console.log('修改信息')
        }
      }
    },
    // 系统信息验重
    checkRepeat() {
      console.log('修改信息')
    },
    // dialog 取消按钮
    handleCancle() {
      this.form.sysName = ''
      this.form.sysCode = ''
      this.form.IsValid = true
      this.dialogFormVisible = false
      this.dialogInfoVisable = false
    },
    // 系统下拉初始化方法
    getSysData() {
      sysDataSelect().then(response => {
        this.options = response.Data
      })
    },
    // 绑定 系统下拉的值修改事件
    changeSysSelect(val) {
      getPerInfoBySysCode(val).then(response => {
        this.authData = response.Data
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
    }

  }
}
</script>
<style rel="stylesheet/scss" lang="scss" scoped>

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}

#action_line{
  margin-top:1%
}
#dialogInfo {
  color:  red
}

.el-checkbox{
  margin:2% 5%
}
</style>
