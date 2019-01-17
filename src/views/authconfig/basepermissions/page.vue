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
    <el-row id="action_line">
      <el-button @click="handleInsert"> 新增菜单</el-button>
      <el-button type="primary" @click="handleSetMeal">生成套餐</el-button>
    </el-row>
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
        :show-overflow-tooltip="true"
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

    <!-- 分页控件  end -->
    <!-- 弹出层 信息录入和修改  start -->
    <el-dialog
      :close-on-click-modal="false"
      :visible.sync="dialogFormVisible"
      width="30%"
      @open="handleOpenDialog"
      @close="handleCloseDialog"
    ><h4 v-if="type==='detail'" slot="title">菜单详情</h4>
      <h4 v-else-if="type==='update'" slot="title">修改菜单</h4>
      <h4 v-else slot="title">新增菜单</h4>
      <el-form ref="perForm" :model="form" :rules="formRules" :disabled="type=='detail'?true:false" size="small">
        <el-form-item
          :label-width="formLabelWidth"
          label="系统名称"
          prop="SysCode"
        >
          <el-select v-model="form.SysCode" placeholder="请选择" @change="changeSysSelect">
            <el-option
              v-for="item in options"
              :key="item.SysCode"
              :label="item.SysName"
              :value="item.SysCode"/>
          </el-select>
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="菜单编码"
          prop="Name"
        >
          <el-input
            v-model="form.Name"
            auto-complete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="菜单名称"
          prop="DisplayName"
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
          <template slot-scope="scope">
            <el-input v-model="scope.row.DisplayName" :disabled="type=='detail'?true:false" size="mini" placeholder="请输入内容" @change="handleEdit(scope.$index, scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column label="权限缩写">
          <template slot-scope="scope">
            <el-input v-model="scope.row.Name" :disabled="type=='detail'?true:false" size="mini" placeholder="请输入内容" @change="handleEdit(scope.$index, scope.row)"/>
          </template>
        </el-table-column>
        <el-table-column v-if="type!='detail'" label="操作">
          <template slot-scope="scope">
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
import { getMenuList, addPerInfo, getPerInfoByMenuId, updatePerInfo } from '@/api/permission'
import { sysDataSelect } from '@/api/sysconfig'
export default {
  data() {
    return {
      type: 'insert',
      tableData: [],
      search: {
        pageSize: 10,
        offset: 0
      },
      form: {
        PerData: []
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '100px',
      dialogInfoVisable: false,
      selection: [], // 列表选择框的信息
      selectionSysCodes: [], // 勾选的系统信息
      options: [], //  系统下拉数据
      formRules: {
        SysCode: [{ required: true, trigger: 'change', message: '系统为必填项' }],
        Name: [{ required: true, trigger: 'blur', message: '菜单编码为必填项' }, { min: 3, max: 20, message: '输入内容最大长度为20', trigger: 'blur' }],
        DisplayName: [{ required: true, trigger: 'blur', message: '菜单名称为必填项' }, { min: 3, max: 20, message: '输入内容最大长度为20', trigger: 'blur' }]
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
    // 新增事件
    handleInsert() {
      this.dialogFormVisible = true
    },
    // 编辑事件
    handleClick(row) {
      this.type = 'update'
      this.dialogFormVisible = true
      getPerInfoByMenuId(row.Id).then(response => {
        this.form = response.Data
        if (!response.Data.PerData) this.form.PerData = []
      })
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
      getMenuList(this.search).then(response => {
        this.tableData = response.Data.list
        this.search.pageTotal = response.Data.total
      })
    },
    // 保存系统信息
    saveData() {
      this.$refs.perForm.validate(valid => {
        if (valid) {
          if (this.type === 'insert') {
            addPerInfo(this.form).then(response => {
              if (response.Result === 0) {
                this.$message.error(response.Message)
              } else {
                this.dialogFormVisible = false
                this.getList()
              }
            })
          } else {
            updatePerInfo(this.form).then(response => {
              if (response.Result === 0) {
                this.$message.error(response.Message)
              } else {
                this.dialogFormVisible = false
                this.getList()
              }
            })
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
    // 监听dialog的关闭事件
    handleCloseDialog() {
      this.form = {}
      this.form.PerData = []
      this.type = 'insert'
      this.$refs['perForm'].resetFields()
    },
    // 监听dialog的打开事件
    handleOpenDialog() {
      sysDataSelect(0).then(response => {
        this.options = response.Data
      })
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
    handleSelectionChange(select) {
      var data = []
      var sysData = []
      for (var i = 0; i < select.length; i++) {
        data.push(select[i].Id)
        sysData.push(select[i].SysCode)
      }
      this.selection = data
      this.selectionSysCodes = sysData
    },
    // 双击点击事件
    handleRowClick(row, event) {
      this.type = 'detail'
      this.dialogFormVisible = true
      getPerInfoByMenuId(row.Id).then(response => {
        this.form = response.Data
      })
    },
    // 下拉框的数据
    changeSysSelect(val) {
      let obj = {}
      obj = this.options.find((item) => {
        return item.SysCode === val
      })
      this.form.SysName = obj.SysName
    },
    // 生成套餐
    handleSetMeal() {
      var arr = Array.from(new Set(this.selectionSysCodes))
      if (arr.length > 1) {
        this.$message({ message: '请选取同一系统下的权限', type: 'warning' })
        return
      }
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
