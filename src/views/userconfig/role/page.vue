<template>
  <div>
    <!-- 查询 form start -->
    <el-form
      :inline="true"
      :model="search"
      class="demo-form-inline"
    >

      <el-form-item label="角色名称">
        <el-input
          v-model="search.roleName"
          placeholder="角色名称"
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
    <el-row id="action_line" style="margin-bottom:15px">
      <el-button type="primary" @click="dialogFormVisible=true">新增</el-button>
      <el-button type="default" @click="handleDeleteRole" >删除</el-button>
    </el-row>
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
        prop="RoleCode"
        label="角色编号"
        align="center"
      />
      <el-table-column
        prop="RoleName"
        label="角色名称"
        align="center"
      />
      <el-table-column
        prop="SysName"
        label="系统名称"
        align="center"
      />
      <el-table-column
        v-if="false"
        prop="SysCode"
        label="系统编码"
        align="center"
      />
      <el-table-column
        :show-overflow-tooltip="true"
        prop="AuthText"
        label="权限"
        align="center"
      />
      <el-table-column prop="IsValid" align="center" label="是否禁用">
        <template slot-scope="scope">
          <el-switch
            v-model="scope.row.IsValid"
            :active-value="0"
            :inactive-value="1"
            active-color="#13ce66"
            inactive-color="#ff4949"
            @change="handleSwitchChange(scope.row,scope.$index)"/>
        </template>
      </el-table-column>
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
          <el-button
            type="text"
            size="small"
            @click="handleDeleteClick(scope.row.Id)"
          >删除</el-button>
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

    <!-- 分页控件  end -->
    <!-- 弹出层 信息录入和修改  start -->
    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      width="40%"
      @close="handleCloseDialog"
      @open="handleOpenDialog"
    > <h4 v-if="type==='detail'" slot="title">角色详情</h4>
      <h4 v-else-if="type==='update'" slot="title">修改角色</h4>
      <h4 v-else slot="title">新增角色</h4>
      <el-form ref="roleForm" :model="form" :rules="formRules" :disabled="type=='detail'?true:false" size="small">
        <el-form-item
          :label-width="formLabelWidth"
          label="角色名称"
          prop="roleName"
        >
          <el-input
            v-model="form.roleName"
            auto-complete="off"
          />
        </el-form-item>

        <el-form-item
          :label-width="formLabelWidth"
          label="系统名称"
          prop="sysCode"
        >
          <el-select v-model="form.sysCode" placeholder="请选择" @change="changeSysSelect">
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
import { getRoleList, addRoleInfo, updateRoleInfo, updateValidStatus, deleteRole } from '@/api/role'
import { sysDataSelect } from '@/api/sysconfig'
import { getPerInfoBySysCode, getPerInfoByRoleId } from '@/api/permission'
import { transPermissionCheckedData } from '@/api/utils'
export default {
  data() {
    return {
      type: 'insert',
      tableData: [],
      search: {
        pageSize: 10,
        offset: 0
      },
      form: { },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '100px',
      dialogInfoVisable: false,
      options: [],
      authData: [],
      multipleSelection: [],
      formRules: {
        sysCode: [{ required: true, trigger: 'change', message: '系统为必填项' }],
        roleName: [{ required: true, trigger: 'blur', message: '角色名称为必填项' }, { max: 20, message: '输入内容最大长度为20', trigger: 'blur' }]
      }
    }
  },
  created() {
    this.getList()
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
      this.form.roleName = row.RoleName
      this.form.roleCode = row.RoleCode
      this.form.sysCode = row.SysCode
      this.form.id = row.Id
      getPerInfoByRoleId(row.Id, row.SysCode).then(response => {
        this.authData = response.Data
      })
    },
    // 获取列表数据
    getList() {
      getRoleList(this.search).then(response => {
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
      this.search = { pageSize: 10, offset: 0 }
      this.dialogInfoVisable = false
      this.getList()
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`)
    },
    handleCurrentChange(val) {
      var pageSize = this.search.pageSize
      this.search.offset = (val > 1 ? (val - 1) * pageSize : 0)
      getRoleList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 保存系统信息
    saveData() {
      this.$refs.roleForm.validate(valid => {
        if (valid) {
          var transData = transPermissionCheckedData(this.authData)
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
              addRoleInfo(this.form).then(response => {
                if (response.Result === 0) {
                  this.$message.error(response.Message)
                } else {
                  this.dialogFormVisible = false
                  this.getList()
                }
              })
            } else {
              updateRoleInfo(this.form).then(response => {
                if (response.Result === 0) {
                  this.$message.error(response.Message)
                } else {
                  this.dialogFormVisible = false
                  this.getList()
                }
              })
            }
          } else {
            console.log('修改信息')
          }
        }
      })
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
    handleDeleteRole() {
      deleteRole(this.multipleSelection.toString()).then(response => {
        if (response.Result === 1) {
          this.getList()
        } else {
          this.$message({
            showClose: true,
            message: '操作失败',
            type: 'warning'
          })
        }
      })
    },
    // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form = {}
      this.authData = []
      this.type = 'insert'
      this.$refs['roleForm'].resetFields()
    },
    // 监听dialog的打开事件
    handleOpenDialog() {
      sysDataSelect(1).then(response => {
        this.options = response.Data
      })
    },
    // 双击点击事件
    handleRowClick(row, event) {
      this.type = 'detail'
      this.dialogFormVisible = true
      this.form.roleName = row.RoleName
      this.form.sysCode = row.SysCode
      this.form.id = row.Id
      getPerInfoByRoleId(row.Id, row.SysCode).then(response => {
        this.authData = response.Data
      })
    },
    // 表格按钮切换事件
    handleSwitchChange(row, index) {
      updateValidStatus(row).then(response => {
        if (response.Result !== 1) {
          this.$message({
            showClose: true,
            message: '操作失败',
            type: 'warning'
          })
          this.tableData[index].IsValid = (row.IsValid === 1 ? 0 : 1)
        }
      })
    },
    // 删除操作
    handleDeleteClick(id) {
      deleteRole(id).then(response => {
        if (response.Result === 1) {
          this.getList()
        } else {
          this.$message({
            showClose: true,
            message: '操作失败',
            type: 'warning'
          })
        }
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
