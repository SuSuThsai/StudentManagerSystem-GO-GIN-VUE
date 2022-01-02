<template>
  <div>
    <el-table
      :data="data"
      style="width: 100%">
      <el-table-column
        sortable
        prop="start_day"
        label="开始日期"
        width="100">
        <!--        <template slot-scope="scope">{{ scope.row.date }}</template>-->
      </el-table-column>
      <el-table-column
        sortable
        prop="end_day"
        label="结束日期"
        width="100">
        <!--        <template slot-scope="scope">{{ scope.row.date }}</template>-->
      </el-table-column>
      <el-table-column
        sortable
        prop="last_for"
        label="持续天数"
        width="100">
        <!--        <template slot-scope="scope">{{ scope.row.date }}</template>-->
      </el-table-column>
      <el-table-column
        sortable
        prop="userid"
        label="学号"
        width="100"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        sortable
        prop="username"
        label="姓名"
        width="100"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        prop="reason"
        label="请假原因"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="200">
        <template slot-scope="scope">
          <el-row type="flex">
            <el-col :offset="1">
              <el-button @click="look(scope.row)" size="mini">查看</el-button>
            </el-col>
            <el-col>
              <el-tag :type="pizhun_type(scope.row)">{{pizhun_show(scope.row)}}</el-tag>
            </el-col>
          </el-row>
        </template>
      </el-table-column>
    </el-table>
    <!--  <div style="margin-top: 20px">-->
    <!--&lt;!&ndash;    <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button>&ndash;&gt;-->
    <!--&lt;!&ndash;    <el-button @click="toggleSelection()">取消选择</el-button>&ndash;&gt;-->
    <!--  </div>-->
  </div>
</template>

<script>
export default {
  name: "qinjia",
  data() {
    return {
      data: [],
    }
  },
  methods: {
    look(t){
      this.$router.push({
        name: 'teachers_qinjia_look',
        params:{
          userid:t.userid,
          username:t.username,
          reason:t.reason,
          status:t.status,
          id:t.ID
        }
      })
    },
    pizhun_show(t){
      if (t.status==1){
        return "等待批准"
      }
      if(t.status==2){
        return "已批准"
      }
      if(t.status==3){
        return "未批准"
      }
      return "错误状态"
    },pizhun_type(t){
      if (t.status==1){
        return "primary"
      }
      if(t.status==2){
        return "success"
      }
      if(t.status==3){
        return "danger"
      }
      return "错误状态"
    }
  },
  mounted() {
    if (this.data.length != 0) {
      return
    }
    const api = 'http://localhost:3000/api/v1/admin/leavecheckclass'
    this.axios.get(api).then(res => {
      this.data = res.data.data
      console.log(this.data)
    }).catch(err => {
      this.$message({
        message: "失败:" + err.response.data.message,
        type: 'error'
      })
      console.log('err:', err.response.data)
    })

  }
}
</script>

<style scoped>

</style>
