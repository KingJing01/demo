<template>
  <div>
    <!-- 查询 form start -->
    <el-form
      :inline="true"
      :model="search"
      class="demo-form-inline"
    >
      <el-form-item label="菜单名称">
        <el-input
          v-model="search.menuName"
          placeholder="菜单名称"
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
      @selection-change	="handleSelectionChange"
      @row-dblclick="handleRowClick"
    >
      <el-table-column
        type="selection"
        width="auto"
        align="center"/>

      <el-table-column
        prop="SysName"
        label="系统名称"
        align="center"
      />
      <el-table-column
        prop="MenuName"
        label="菜单名称"
        align="center"
      />
      <el-table-column
        prop="MenuText"
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
        :page-size="search.pageSize"
        :total="search.pageTotal"
        layout="total, prev, pager, next"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"/>
    </div>
    <el-row id="action_line">
      <el-button @click="handleInsert"> 新增菜单</el-button>
      <el-button type="primary">生成套餐</el-button>
    </el-row>

    <!-- 分页控件  end -->
    <!-- 弹出层 信息录入和修改  start -->
    <el-dialog
      :visible.sync="dialogFormVisible"
      width="30%"
      @close="handleCloseDialog"
    ><h4 v-if="type==='detail'" slot="title">菜单详情</h4>
      <h4 v-else-if="type==='update'" slot="title">修改菜单</h4>
      <h4 v-else slot="title">新增菜单</h4>
      <el-form :model="form" :disabled="type=='detail'?true:false" size="small">
        <el-form-item
          :label-width="formLabelWidth"
          label="系统名称"
        >
          <el-input
            :disabled="true"
            v-model="form.SysName"
            auto-complete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="菜单编码"
        >
          <el-input
            v-model="form.Name"
            auto-complete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="菜单名称"
        >
          <el-input
            v-model="form.DisplayName"
            auto-complete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="操作名称"
        >
          <el-input
            :disabled="true"
            v-model="form.MenuText"
            auto-complete="off"
          />
        </el-form-item>
      </el-form>
      <el-button v-show="type=='detail'?false:true" type="primary" size="mini" icon="el-icon-circle-plus" @click="handlePerInsert"/>
      <el-table :data="form.PerData" class="tb-edit" style="width: 100%" highlight-current-row max-height="310" @row-click="handleTableCurrentChange">
        <el-table-column label="权限名称">
          <template scope="scope">
            <el-input v-model="scope.row.DisplayName" :disabled="type=='detail'?true:false" size="mini" placeholder="请输入内容" @change="handleEdit(scope.$index, scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column label="权限缩写">
          <template scope="scope">
            <el-input v-model="scope.row.Name" :disabled="type=='detail'?true:false" size="mini" placeholder="请输入内容" @change="handleEdit(scope.$index, scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column v-if="type!='detail'" label="操作">
          <template scope="scope">
            <el-button size="mini" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="handleCancle">取 消</el-button>
        <el-button
          v-show="type=='detail'?false:true"
          type="primary"
          @click="saveData"
        >确 定</el-button>
      </div>
    </el-dialog>
    <!-- 弹出层 信息录入和修改  end -->
  </div>
</template>
<script>
import { getMenuList, addPerInfo, getPerInfoByMenuId } from '@/api/permission'
export default {
  data() {
    return {
      type: 'insert',
      tableData: [],
      search: {
        pageSize: 5,
        offset: 0
      },
      form: {
        PerData: []
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '100px',
      dialogInfoVisable: false,
      selection: [] // 列表选择框的信息
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
    // 新增事件
    handleInsert() {
      var select = this.selection
      if (select.length === 1) {
        this.form.SysCode = select[0].SysCode
        this.form.SysName = select[0].SysName
        this.dialogFormVisible = true
      } else if (select.length === 0) {
        this.$message({
          message: '请选择一条记录',
          type: 'warning'
        })
        return false
      } else {
        this.$message({
          message: '新增操作只能选择一条记录',
          type: 'warning'
        })
        return false
      }
    },
    // 编辑事件
    handleClick(row) {
      this.type = 'update'
      this.dialogFormVisible = true
      this.form.SysName = row.SysName
      this.form.id = row.Id
    },
    // 获取列表数据
    getList() {
      getMenuList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 重置按钮
    onReset() {
      this.search = { pageSize: 5, offset: 0 }
      this.dialogInfoVisable = false
      this.getList()
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`)
    },
    handleCurrentChange(val) {
      var pageSize = this.search.pageSize
      this.search.offset = (val > 1 ? (val - 1) * pageSize : 0)
      getMenuList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 保存系统信息
    saveData() {
      if (this.type === 'insert') {
        addPerInfo(this.form).then(response => {
          this.getList()
        })
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
    // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form = {}
      this.form.PerData = []
      this.type = 'insert'
    },
    handleTableCurrentChange(row, event, column) {
      console.log(row, event, column, event.currentTarget)
    },
    handleEdit(index, row) {
      console.log(index, row)
    },
    handleDelete(index, row) {
      this.form.PerData.splice(index, 1)
    },
    // 权限表上的新增点击事件
    handlePerInsert(column, event) {
      var data = { displayName: '', name: '' }
      this.form.PerData.push(data)
    },
    handleSelectionChange(selection) {
      this.selection = selection
    },
    // 双击点击事件
    handleRowClick(row, event) {
      this.type = 'detail'
      this.dialogFormVisible = true
      getPerInfoByMenuId(row.Id).then(response => {
        this.form = response.Data
      })
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

#action_line{
  margin-top:1%
}
#dialogInfo {
  color:  red
}
</style>
