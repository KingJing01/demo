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
      <el-button @click="dialogFormVisible = true"> 新增菜单</el-button>
      <el-button type="primary">生成套餐</el-button>
    </el-row>

    <!-- 分页控件  end -->
    <!-- 弹出层 信息录入和修改  start -->
    <el-dialog
      :visible.sync="dialogFormVisible"
      width="40%"
    ><h4 v-if="type==='detail'" slot="title">菜单详情</h4>
      <h4 v-else-if="type==='update'" slot="title">修改菜单</h4>
      <h4 v-else slot="title">新增菜单</h4>
      <el-form :model="form">
        <el-form-item
          :label-width="formLabelWidth"
          label="系统名称"
        >
          <el-input
            :disabled="true"
            v-model="form.sysName"
            auto-complete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="菜单编码"
        >
          <el-input
            v-model="form.menuCode"
            auto-complete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="菜单名称"
        >
          <el-input
            v-model="form.menuName"
            auto-complete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="操作名称"
        >
          <el-input
            :disabled="true"
            v-model="form.menuText"
            auto-complete="off"
          />
        </el-form-item>
      </el-form>
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
import { getMenuList } from '@/api/permission'
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
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '120px',
      dialogInfoVisable: false
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
      this.form.sysCode = row.SysCode
      this.form.sysName = row.SysName
      this.form.id = row.Id
      this.form.sysUrl = row.SysUrl
      if (row.IsValid === 0) {
        this.form.IsValid = true
      } else {
        this.form.IsValid = false
      }
    },
    // 获取列表数据
    getList() {
      getMenuList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
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
      this.search = { pageSize: 5, offset: 0 }
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
        if (this.type === 'insert') {
          console.log('新增信息')
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
      this.dialogFormVisible = false
      this.dialogInfoVisable = false
    },
    // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form = {}
      this.type = 'insert'
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
