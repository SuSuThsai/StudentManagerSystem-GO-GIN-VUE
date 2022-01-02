<template>
  <div>
    <el-table
      :data="data"
      style="width: 100%">
      <el-table-column
        sortable
        prop="creat_day"
        label="日期"
        width="150">
        <!--        <template slot-scope="scope">{{ scope.row.date }}</template>-->
      </el-table-column>
      <el-table-column
        sortable
        prop="update_day"
        label="最近更新"
        width="200">
        <!--        <template slot-scope="scope">{{ scope.row.date }}</template>-->
      </el-table-column>
      <el-table-column
        prop="title"
        label="实习通知标题"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        prop="text"
        label="通知内容"
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
  name: "shixitonzhi",
  data(){
    return{
      data:[]
    }
  },
  methods:{
    look(t){
      this.$router.push({
        name: 'students_shixitonzhi_look',
        params:{
          title:t.title,
          text:t.text,
          creat_day:t.creat_day,
          update_day:t.update_day,
          userid:t.userid,
          username:t.username
        }
      })
    },
  },
  mounted() {
    if (this.data.length != 0) {
      return
    }
    const api = 'http://localhost:3000/api/v1/student/noticecheckclass'
    this.axios.get(api).then(res => {
      this.data = res.data.data
    }).catch(err => {
      this.$message({
        message: "获取个人实习报告失败:" + err.response.data.message,
        type: 'error'
      })
      console.log('err:', err.response.data)
    })
  }
}
</script>

<style scoped>

</style>
