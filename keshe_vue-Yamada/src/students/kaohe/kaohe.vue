<template>
  <div>
    <el-table
      :data="data"
      style="width: 100%">
      <el-table-column
        sortable
        prop="userid"
        label="学号"
        width="200">
      </el-table-column>
      <el-table-column
        prop="username"
        label="姓名"
        width="150">
      </el-table-column>
      <el-table-column
        sortable
        prop="working_ability"
        label="工作能力分值"
        width="150"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        sortable
        prop="working_attitude"
        label="工作态度分值"
        width="150"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        sortable
        prop="communication"
        label="沟通能力分值"
        width="150"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        sortable
        prop="total_score"
        label="总成绩(工作能力*0.5+工作态度*0.25+沟通能力*0.25)"
        show-overflow-tooltip>
      </el-table-column>
      <!--      <el-table-column-->
      <!--        fixed="right"-->
      <!--        label="操作"-->
      <!--        width="200">-->
      <!--        <template slot-scope="scope">-->
      <!--          <el-col :offset="1">-->
      <!--            <el-button @click="look(scope.row)" size="mini">查看</el-button>-->
      <!--          </el-col>-->
      <!--        </template>-->
      <!--      </el-table-column>-->
    </el-table>
    <!--  <div style="margin-top: 20px">-->
    <!--&lt;!&ndash;    <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button>&ndash;&gt;-->
    <!--&lt;!&ndash;    <el-button @click="toggleSelection()">取消选择</el-button>&ndash;&gt;-->
    <!--  </div>-->
  </div>
</template>

<script>
export default {
  name: "students_kaohe",
  data() {
    return {
      data: [],
    }
  },
  methods: {
  },
  mounted() {
    if (this.data.length != 0) {
      return
    }
    const api = 'http://localhost:3000/api/v1/student/assessmentcheckitself'
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
