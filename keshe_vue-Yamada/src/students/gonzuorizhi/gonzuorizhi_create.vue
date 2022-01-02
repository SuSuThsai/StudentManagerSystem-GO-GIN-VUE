<template>
  <div>
    <span>创建工作日志</span>
    <el-card>
      <el-main>
        <el-row type="flex">
          <span >标题：</span>
          <el-row>
            {{"本日"}}
          </el-row>
        </el-row>
        <div style="margin: 20px 0;"></div>
        <div style="margin: 20px 0;"></div>
        <el-row type="flex">
          <span >负责人：</span>
          <el-row>
            {{"本账户名"}}
          </el-row>
        </el-row>
        <div style="margin: 20px 0;"></div>
        <span>成员：</span>
        <el-input
          type="team_member"
          v-model="team_member">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>目的：</span>
        <el-input
          type="goal"
          autosize
          placeholder="请输入内容"
          v-model="goal">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>日志内容：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="results">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <el-button @click="create_">
          确认
        </el-button>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "shixibaogao_create",
  data(){
    return{
      username:"",
      title: "",
      manager: "",
      team_member: "",
      goal: "",
      results: "",
      creat_day: "",
      id: "",
      userid:""
    }
  },
  methods:{
    create_(){
      const api="http://localhost:3000/api/v1/student/worklogcheckpost"

      let e=new URLSearchParams()
      e.append("team_member",this.team_member)
      e.append("goal",this.goal)
      e.append("results",this.results)
      e.append("title",this.title)
      e.append("userid",this.userid)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "创建成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/students/gonzuorizhi")
      }).catch(err => {
        this.$message({
          message: "创建失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    }
  },
}
</script>

<style scoped>

</style>
