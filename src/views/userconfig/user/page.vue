<template>
  <div>
    <!-- 查询 form start -->
    <el-form
      :inline="true"
      :model="search"
      class="demo-form-inline"
    ><el-form-item label="系统名称">
      <el-input
        v-model="search.sysName"
        placeholder="系统名称"
      />
    </el-form-item>
      <el-form-item label="角色名称">
        <el-input
          v-model="search.roleName"
          placeholder="套餐名称"
        />
      </el-form-item>
      <el-form-item label="登录名">
        <el-input
          v-model="search.userName"
          placeholder="用户名/邮箱/手机号"
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
    <el-row id="action_line" style="margin-bottom:10px">
      <el-button @click="dialogFormVisible = true">新增用户</el-button>
      <el-button type="primary" @click="handleDeleteUser">删除用户</el-button>
      <el-button type="success" @click="handleAddExistUser">新增</el-button>
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
        prop="Id"
        label="用户编号"
        align="center"
      />
      <el-table-column
        prop="UserName"
        label="登录名"
        align="center"
      />
      <el-table-column
        prop="PhoneNumber"
        label="手机号"
        align="center"
      />
      <el-table-column
        :show-overflow-tooltip="true"
        prop="EmailAddress"
        label="邮箱"
        align="center"
      />
      <el-table-column
        prop="SysName"
        label="系统名称"
        align="center"
      /><el-table-column
        prop="RoleName"
        label="角色名称"
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
      width="25%"
      @close="handleCloseDialog"
      @open="handleOpenDialog"
    > <h4 v-if="type==='detail'" slot="title">用户详情</h4>
      <h4 v-else-if="type==='update'" slot="title">修改用户</h4>
      <h4 v-else-if="type==='updateNew'" slot="title">更新用户</h4>
      <h4 v-else slot="title">新增用户</h4>
      <div v-if="type==='detail'"><DetailPage ref="userData" :data="form"/></div>
      <div v-else-if="type==='update'"><UpdatePage ref="userData" :data="form"/></div>
      <div v-else-if="type==='insert'"><SavePage ref="userData" :data="form"/></div>
      <div v-else><UpdateNewPage ref="userData" :data="form"/></div>
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
import { getUserList, updateUserInfo, addUserInfo, updateUserValidStatus, deleteUser } from '@/api/user'

import DetailPage from './dialogview/detail'
import SavePage from './dialogview/save'
import UpdatePage from './dialogview/update'
import UpdateNewPage from './dialogview/updateuser'
export default {
  components: { DetailPage, SavePage, UpdatePage, UpdateNewPage },
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
      options: [],
      authData: [],
      multipleSelection: [],
      formRules: {
        sysCode: [{ required: true, trigger: 'change', message: '系统为必填项' }],
        setMealName: [{ required: true, trigger: 'blur', message: '套餐名为必填项' }, { max: 20, message: '输入内容最大长度为20', trigger: 'blur' }]
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
      this.form.Id = row.Id
    },
    // 获取列表数据
    getList() {
      getUserList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    handleClose(done) {
      done()
    },

    // 重置按钮
    onReset() {
      this.search = { pageSize: 10, offset: 0 }
      this.getList()
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`)
    },
    handleCurrentChange(val) {
      var pageSize = this.search.pageSize
      this.search.offset = (val > 1 ? (val - 1) * pageSize : 0)
      getUserList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 保存系统信息
    saveData() {
      this.$refs.userData.validData()
      const valid = this.form.valid
      if (valid) {
        // 新增操作
        if (this.type === 'insert') {
          addUserInfo(this.form.formData).then(response => {
            if (response.Result === 0) {
              this.$message.error(response.Message)
            } else {
              this.dialogFormVisible = false
              this.getList()
            }
          })
        } else {
          updateUserInfo(this.form.formData).then(response => {
            if (response.Result === 0) {
              this.$message.error(response.Message)
            } else {
              this.dialogFormVisible = false
              this.getList()
            }
          })
        }
      }
    },
    // dialog 取消按钮
    handleCancle() {
      this.dialogFormVisible = false
      this.dialogInfoVisable = false
    },
    // 表格选择框的改变事件 监听
    selectChangeFun(selection) {
      var data = []
      for (var i = 0; i < selection.length; i++) {
        data.push(selection[i].Id)
      }
      this.multipleSelection = data
    },
    // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form = {}
      this.authData = []
      this.type = 'insert'
      this.$refs.userData.validData()
      this.$refs.userData.cleanData()
    },
    // 监听dialog的打开事件
    handleOpenDialog() {
      this.$refs.userData.getSysList()
      this.$refs.userData.cancleValid()
    },
    // 双击点击事件
    handleRowClick(row, event) {
      this.type = 'detail'
      this.dialogFormVisible = true
      this.form.Id = row.Id
    },
    // 表格按钮切换事件
    handleSwitchChange(row, index) {
      updateUserValidStatus(row).then(response => {
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
      deleteUser(id).then(response => {
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
    // 批量禁用套餐
    handleDeleteUser() {
      deleteUser(this.multipleSelection.toString()).then(response => {
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
    handleAddExistUser(data) {
      if (this.multipleSelection.length !== 1) {
        this.$message({
          showClose: true,
          message: '有且只有一个用户',
          type: 'warning'
        })
        return
      }
      this.form.Id = this.multipleSelection.toString()
      this.type = 'updateNew'
      this.dialogFormVisible = true
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
