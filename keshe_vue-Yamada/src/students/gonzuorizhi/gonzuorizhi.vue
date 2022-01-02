<template>
  <div>
    <div>
      <el-button @click="create" type="success">
        新建
      </el-button>
    </div>
    <el-table
      :data="data"
      style="width: 100%">
      <el-table-column
        sortable
        prop="creat_day"
        label="日期"
        width="100">
      </el-table-column>
      <el-table-column
        sortable
        prop="update_day"
        label="最近更改"
        width="100">
      </el-table-column>
      <el-table-column
        sortable
        prop="userid"
        label="学号"
        width="100"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        prop="username"
        label="姓名"
        width="150"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        prop="title"
        label="标题"
        width="200">
      </el-table-column>
      <el-table-column
        prop="manager"
        label="管理者">
      </el-table-column>
      <el-table-column
        prop="team_member"
        label="成员">
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="200">
        <template slot-scope="scope">
            <el-col :offset="1">
              <el-button @click="look(scope.row)" size="mini">编辑</el-button>
            </el-col>
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
  name: "students_gonzuorizhi",
  data() {
    return {
      data: [],
      userid:""
    }
  },
  methods: {
    create() {
      this.$router.push({
        name: 'students_gonzuorizhi_create',
      })
    },
    look(t){
      this.$router.push({
        name: 'students_gonzuorizhi_look',
        params:{
          id:t.ID,
          userid:t.userid,
          class:t.class,
          title:t.title,
          manager:t.manager,
          team_member:t.team_member,
          goal:t.goal,
          results:t.results,
          creat_day:t.creat_day,
          update_day:t.update_day,
        }
      })
    },
  },
  mounted() {
    if (this.data.length != 0) {
      return
    }
    const api = 'http://localhost:3000/api/v1/student/worklogcheckitself'
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
