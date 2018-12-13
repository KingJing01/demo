<template>
  <div>
    <!-- 查询 form start -->
    <el-form
      :inline="true"
      :model="search"
      class="demo-form-inline"
    >
      <el-form-item label="系统代码">
        <el-input
          v-model="search.sysCode"
          placeholder="系统代码"
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
        prop="SysCode"
        label="系统代码"
        align="center"
      />
      <el-table-column
        prop="SysName"
        label="系统名称"
        align="center"
      />
      <el-table-column
        prop="SysUrl"
        label="系统地址"
        align="center"
      />
      <el-table-column
        :formatter="formatText"
        prop="IsValid"
        label="是否有效"
        align="center"
      />
      <el-table-column
        prop="address"
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
        :visible.sync="dialogFormVisible"
        width="40%"
        @close="handleCloseDialog"
      ><h4 v-if="type==='detail'" slot="title">系统详情</h4>
        <h4 v-else-if="type==='update'" slot="title">修改系统</h4>
        <h4 v-else slot="title">新增系统</h4>
        <el-form :model="form" size="small">
          <el-form-item
            :label-width="formLabelWidth"
            label="系统代码"
          >
            <el-input
              :disabled="true"
              v-model="form.sysCode"
              auto-complete="off"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="系统名称"
          >
            <el-input
              :disabled="type=='detail'?true:false"
              v-model="form.sysName"
              auto-complete="off"
              @change="checkRepeat"
            />
            <span v-if= "dialogInfoVisable==true" id="dialogInfo">系统名称已经存在,请重新输入</span>
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label="系统地址"
          >
            <el-input
              :disabled="type=='detail'?true:false"
              v-model="form.sysUrl"
              auto-complete="off"
            />
          </el-form-item>
          <el-form-item
            :label-width="formLabelWidth"
            label=""
          >
            <el-checkbox v-model="form.IsValid" :disabled="type=='detail'?true:false">是否有效</el-checkbox>
          </el-form-item>
        </el-form>
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
import { getListData, saveSysInfo, uniqueCheck, updateSysInfo } from '@/api/sysconfig'
export default {
  data() {
    return {
      tableData: [],
      search: {
        pageSize: 5,
        offset: 0
      },
      form: {
        IsValid: true
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '100px',
      dialogInfoVisable: false,
      type: 'insert'
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
      getListData(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 文本格式转换
    formatText(row, column) {
      const data = row[column.property]
      return data === 0 ? '是' : '否'
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
      getListData(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 保存系统信息
    saveData() {
      if (this.dialogInfoVisable === false) {
        if (this.type === 'insert') {
          saveSysInfo(this.form).then(response => {
            this.dialogFormVisible = false
            this.dialogInfoVisable = false
            this.getList()
          })
        } else {
          updateSysInfo(this.form).then(response => {
            this.dialogFormVisible = false
            this.dialogInfoVisable = false
            this.getList()
          })
        }
      }
    },
    // 系统信息验重
    checkRepeat() {
      uniqueCheck(this.form.sysName).then(response => {
        if (response.Result > 0) {
          this.dialogInfoVisable = true
        } else {
          this.dialogInfoVisable = false
        }
      })
    },
    // dialog 取消按钮
    handleCancle() {
      this.dialogFormVisible = false
      this.dialogInfoVisable = false
    },
    // 双击点击事件
    handleRowClick(row, event) {
      this.type = 'detail'
      this.dialogFormVisible = true
      this.form.sysCode = row.SysCode
      this.form.sysName = row.SysName
      this.form.sysUrl = row.SysUrl
      if (row.IsValid === 0) {
        this.form.IsValid = true
      } else {
        this.form.IsValid = false
      }
    }, // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form = { IsValid: true }
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

#action_item{
  float:right;
  margin-right:10%
}
#dialogInfo {
  color:  red
}
</style>
