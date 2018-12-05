<template>
  <div>
    <el-form
      :inline="true"
      :model="form"
      class="demo-form-inline"
    >
      <el-form-item label="系统代码">
        <el-input
          v-model="form.sysCode"
          placeholder="系统代码"
        />
      </el-form-item>
      <el-form-item label="系统名称">
        <el-input
          v-model="form.sysName"
          placeholder="系统名称"
        />
      </el-form-item>
      <el-form-item>
        <el-button
          type="primary"
          @click="onSubmit"
        >查询</el-button>
        <el-button
          type="primary"
          @click="dialogFormVisible  = true"
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
        prop="IsValid"
        label="是否有效"
        align="center"
      :formatter="formatText"

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
        layout="prev, pager, next"
        aligen="center"
      />
    </div>

    <el-dialog
      title="新增系统"
      width="40%"
      :visible.sync="dialogFormVisible"
    >
      <el-form :model="form">
        <el-form-item
          label="系统代码"
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="form.name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item
          label="系统名称"
          :label-width="formLabelWidth"
        >
          <el-input
            v-model="form.name"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item
          label=""
          :label-width="formLabelWidth"
        >
          <el-checkbox v-model="checked">是否有效</el-checkbox>
        </el-form-item>

      </el-form>
      <div
        slot="footer"
        class="dialog-footer"
      >
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button
          type="primary"
          @click="dialogFormVisible = false"
        >确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import { getListData } from "@/api/sysconfig";
export default {
  data() {
    return {
      tableData: [],
      form: {
        sysCode: "",
        sysName: ""
      },
      dialogTableVisible: false,
      dialogFormVisible: false,
      formLabelWidth: "120px",
      checked: true
    };
  },
  created() {
    this.getList();
  },
  methods: {
    onSubmit() {
      console.log("submit!");
    },
    handleClick(row) {
      console.log(row);
    },
    getList() {
      getListData().then(response => {
        this.tableData = response.Data.list;
      });
    },
    handleClose(done) {
      done();
    },
    formatText(row, column){
      const data = row[column.property]
      return data==0?'是':'否'
    }
  }
};
</script>
<style>
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
</style>
