<template>
  <div>
    <!-- 查询 form start -->
    <el-form
      :inline="true"
      :model="search"
      class="demo-form-inline"
    >
      <el-form-item label="系统名称">
        <el-input
          v-model="search.sysName"
          placeholder="系统名称"
        />
      </el-form-item>
      <el-form-item label="用户名称">
        <el-input
          v-model="search.tenantName"
          placeholder="用户名称"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          @click="onSubmit"
        >查询</el-button>
        <el-button @click="onReset">重置</el-button>
      </el-form-item>
      <el-form-item id="action_item">
        <el-button
          type="success"
          @click="dialogFormVisible = true"
        >新增</el-button>
      </el-form-item>
    </el-form>
    <!-- 查询 form end -->
    <!-- 系统信息列表  start -->
    <el-table
      :data="tableData"
      style="width: 90%"
      border
      @row-dblclick="handleRowClick"
    >
      <el-table-column
        prop="TenantName"
        label="用户名称"
        align="center"
      />
      <el-table-column
        prop="SysName"
        label="系统名称"
        align="center"
      />
      <el-table-column
        :show-overflow-tooltip="true"
        prop="MenuText"
        label="菜单"
        align="center"
      />
      <el-table-column
        prop="Operator"
        label="录入人"
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
    <!-- 系统信息列表  end -->
    <!-- 分页控件  start -->
    <template>
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
      ><h4 v-if="type==='detail'" slot="title">用户详情</h4>
        <h4 v-else-if="type==='update'" slot="title">修改用户信息</h4>
        <h4 v-else slot="title">新增用户信息</h4>
        <div v-if="type==='detail'"><DetailPage :data="form"/></div>
        <div v-else-if="type==='update'"><UpdatePage ref="userData" :data="form"/></div>
        <div v-else><SavePage ref="userData" :data="form"/></div>
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
  </template></div>
</template>
<script>
import { getUserList, updateTenantInfo, saveTenantInfo } from '@/api/usermanage'
import { transPermissionCheckedData, transPermissionCheckedDataArr } from '@/api/utils'

import DetailPage from './dialogview/detail'
import SavePage from './dialogview/save'
import UpdatePage from './dialogview/update'

export default {
  components: { DetailPage, SavePage, UpdatePage },
  data() {
    return {
      tableData: [],
      search: {
        tenantName: '',
        sysName: '',
        pageSize: 10,
        offset: 0
      },
      form: {},
      dialogTableVisible: false,
      dialogFormVisible: false,
      dialogInfoVisable: false,
      type: 'insert',
      tenant: {}
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
      this.form.id = row.Id
      this.form.sysCode = row.SysCode
      this.dialogFormVisible = true
    },
    // 获取列表数据
    getList() {
      getUserList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
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
      getUserList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 保存/修改用户信息
    saveData() {
      this.$refs.userData.validData()
      const valid = this.form.valid
      if (valid) {
        if (this.type === 'insert') {
          const transData = transPermissionCheckedDataArr(this.form.authData)
          var perId = []
          var perName = []
          var perMenu = []
          var sysCode = []
          for (const index in transData) {
            var arrIndexData = transData[index]
            if (arrIndexData.perName === '') {
              this.$message({
                message: '系统中必须选择操作权限',
                type: 'warning'
              })
              return false
            } else {
              perId.push(arrIndexData.perId)
              perName.push(arrIndexData.perName)
              perMenu.push(arrIndexData.perMenu)
              sysCode.push(arrIndexData.sysCode)
            }
          }
          this.tenant = this.form.formData
          this.tenant.perMenu = perMenu
          this.tenant.perId = perId
          this.tenant.sysCode = sysCode
          saveTenantInfo(this.tenant).then(response => {
            if (response.Result === 0) {
              this.$message.error(response.Message)
            } else {
              this.dialogFormVisible = false
              this.getList()
              this.tenant = {}
            }
          })
        } else {
          const transData = transPermissionCheckedData(this.form.authData)
          if (transData.perName === '') {
            this.$message({
              message: '请选择操作权限',
              type: 'warning'
            })
            return false
          }
          this.tenant = this.form.formData
          this.tenant.perMenu = transData.perMenu
          this.tenant.perId = transData.perId
          this.tenant.sysCode = this.form.sysCode
          updateTenantInfo(this.tenant).then(response => {
            if (response.Result === 0) {
              this.$message.error(response.Message)
            } else {
              this.dialogFormVisible = false
              this.getList()
              this.tenant = {}
            }
          })
        }
      }
    },
    // dialog 取消按钮
    handleCancle() {
      this.dialogFormVisible = false
    },
    // 双击点击事件
    handleRowClick(row, event) {
      this.dialogFormVisible = true
      this.type = 'detail'
      this.form.sysCode = row.SysCode
      this.form.id = row.Id
    }, // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form = {}
      this.type = 'insert'
      this.$refs.userData.cancleValid()
      this.$refs.userData.cleanData()
    }
  }
}
</script>
<style rel="stylesheet/scss" lang="scss" scoped>
.el-row {
  margin-bottom: 20px;
  &:last-child {
    margin-bottom: 0;
  }
}

.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}

#action_item{
  float:right;
  margin-right:10%
}
#dialogInfo {
  color:  red
}
</style>
