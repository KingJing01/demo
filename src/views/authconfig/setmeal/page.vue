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
      @selection-change="selectChangeFun"
      @row-dblclick="handleRowClick"
    >
      <el-table-column
        type="selection"
        width="auto"
        align="center"
      />
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
      /><el-table-column
        v-if="hide"
        prop="SysCode"
        label="系统编码"
        align="center"
      />
      <el-table-column
        prop="PermissionText"
        label="操作名称"
        align="center"
        show-overflow-tooltip="true"
      />
      <el-table-column
        :formatter="formatText"
        prop="IsDeleted"
        label="是否有效"
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
        :page-size="search.pageSize"
        :total="search.pageTotal"
        layout="total, prev, pager, next"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"/>
    </div>
    <el-row id="action_line">
      <el-button @click="dialogFormVisible = true">新增套餐</el-button>
      <el-button type="primary" @click="handleDeleteSetMeal">禁用套餐</el-button>
    </el-row>

    <!-- 分页控件  end -->
    <!-- 弹出层 信息录入和修改  start -->
    <el-dialog
      :visible.sync="dialogFormVisible"
      width="40%"
      @close="handleCloseDialog"
    > <h4 v-if="type==='detail'" slot="title">套餐详情</h4>
      <h4 v-else-if="type==='update'" slot="title">修改套餐</h4>
      <h4 v-else slot="title">新增套餐</h4>
      <el-form :model="form">
        <div v-if="type!='detail'">
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
            label="系统名称"
          >
            <el-select v-model="form.sysCode" placeholder="请选择" @change="changeSysSelect">
              <el-option
                v-for="item in options"
                :key="item.SysCode"
                :label="item.SysName"
                :value="item.SysCode"/>
            </el-select>
          </el-form-item>
        </div>
        <div v-else>
          <span> {{ form.sysName }}  {{ form.setMealName }} </span>
        </div>
      </el-form>
      <template>
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
      </div></template>
      <div
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="handleCancle">取 消</el-button>
        <el-button
          v-show="type!='detail'"
          type="primary"
          @click="saveData"
        >确 定</el-button>
      </div>
    </el-dialog>
    <!-- 弹出层 信息录入和修改  end -->
  </div>
</template>
<script>
import { getSetMealList, addSetMealInfo, deleteSetMeal, updateSetMealInfo } from '@/api/setmeal'
import { sysDataSelect } from '@/api/sysconfig'
import { getPerInfoBySysCode, getPerInfoBySysCodeUpdate } from '@/api/permission'
import { transPermisionCheckedData } from '@/api/utils'
export default {
  data() {
    return {
      type: 'insert',
      tableData: [],
      search: {
        setMealName: '',
        sysName: '',
        pageSize: 5,
        offset: 0
      },
      form: {
        perName: '',
        setMealName: '',
        setMealCode: '',
        perId: '',
        sysCode: '',
        id: ''
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '120px',
      dialogInfoVisable: false,
      options: [],
      authData: [],
      multipleSelection: []
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
      this.type = 'update'
      this.dialogFormVisible = true
      this.form.setMealName = row.SetMealName
      this.form.setMealCode = row.SetMealCode
      this.form.sysCode = row.SysCode
      this.form.id = row.Id
      getPerInfoBySysCodeUpdate(row.SysCode, row.SetMealCode).then(response => {
        this.authData = response.Data
      })
    },
    // 获取列表数据
    getList() {
      getSetMealList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    handleClose(done) {
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
      var pageSize = this.search.pageSize
      this.search.offset = (val > 1 ? (val - 1) * pageSize : 0)
      getSetMealList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 保存系统信息
    saveData() {
      var transData = transPermisionCheckedData(this.authData)
      if (transData.perName === '') {
        this.$message({
          message: '请选择操作权限',
          type: 'warning'
        })
        return false
      }
      this.form.perId = transData.perId
      this.form.perName = transData.perName
      if (this.dialogInfoVisable === false) {
        // 新增操作
        if (this.type === 'insert') {
          addSetMealInfo(this.form).then(response => {
            this.dialogFormVisible = false
            this.getList()
          })
        } else {
          updateSetMealInfo(this.form).then(response => {
            this.dialogFormVisible = false
            this.getList()
          })
        }
      } else {
        console.log('修改信息')
      }
    },
    // 系统信息验重
    checkRepeat() {
      console.log('修改信息')
    },
    // dialog 取消按钮
    handleCancle() {
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
    },
    // 表格选择框的改变事件 监听
    selectChangeFun(selection) {
      var data = []
      for (var i = 0; i < selection.length; i++) {
        data.push(selection[i].Id)
      }
      this.multipleSelection = data
    },
    // 批量禁用套餐
    handleDeleteSetMeal() {
      deleteSetMeal(this.multipleSelection.toString()).then(response => {
        this.getList()
      })
    },
    // 文本格式转换
    formatText(row, column) {
      const data = row[column.property]
      return data === 0 ? '是' : '否'
    },
    // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form.perName = ''
      this.form.perId = ''
      this.form.sysCode = ''
      this.form.setMealName = ''
      this.authData = []
      this.type = 'insert'
    },
    // 双击点击事件
    handleRowClick(row, event) {
      this.type = 'detail'
      this.dialogFormVisible = true
      this.form.setMealName = row.SetMealName
      this.form.sysName = row.SysName
      getPerInfoBySysCodeUpdate(row.SysCode, row.SetMealCode).then(response => {
        this.authData = response.Data
      })
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
