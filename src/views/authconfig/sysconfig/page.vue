<template>
  <div>
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
    <el-table
      :data="tableData"
      style="width: 90%"
      border
    >
      <el-table-column
        type="index"
        label="序号"
        align="center"
        width="auto"
      />
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
    <div class="block">
      <el-pagination
        :page-size="20"
        :pager-count="11"
        :total="1000"
        layout="prev, pager, next"/>
    </div>

    <el-dialog
      :visible.sync="dialogFormVisible"
      title="新增系统"
      width="40%"
    >
      <el-form :model="form">
        <el-form-item
          :label-width="formLabelWidth"
          label="系统代码"
        >
          <el-input
            :disabled="true"
            v-model="form.name"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label="系统名称"
        >
          <el-input
            v-model="form.sysName"
            autocomplete="off"
            @change="checkRepeat"
          />
          <span v-if= "dialogInfo==true" id="dialogInfo">系统名称已经存在,请重新输入</span>
        </el-form-item>
        <el-form-item
          :label-width="formLabelWidth"
          label=""
        >
          <el-checkbox v-model="form.IsValid">是否有效</el-checkbox>
        </el-form-item>
      </el-form>
      <div
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button
          type="primary"
          @click="saveData"
        >确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { getListData, saveSysInfo, uniqueCheck } from '@/api/sysconfig'
export default {
  data() {
    return {
      tableData: [],
      search: {
        sysCode: '',
        sysName: '',
        pageSize: 10,
        offset: 0
      },
      form: {
        sysName: '',
        IsValid: true
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: '120px',
      dialogInfo: false
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
    handleClick(row) {
      console.log(row.Id)
    },
    // 获取列表数据
    getList() {
      getListData(this.search).then(response => {
        this.tableData = response.Data.list
      })
    },
    handleClose(done) {
      done()
    },
    // 文本格式转换
    formatText(row, column) {
      const data = row[column.property]
      return data === 0 ? '是' : '否'
    },
    // 重置按钮
    onReset() {
      this.search.sysCode = ''
      this.search.sysName = ''
      this.search.pageSize = 10
      this.search.offset = 0
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
      saveSysInfo(this.form).then(response => {
        this.dialogFormVisible = false
        this.getList()
      })
    },
    // 系统信息验重
    checkRepeat() {
      uniqueCheck(this.form.sysName).then(response => {
        if (response.Result > 0) {
          this.dialogInfo = true
        }
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

#action_item{
  float:right;
  margin-right:10%
}
#dialogInfo {
  color:  red
}
</style>
