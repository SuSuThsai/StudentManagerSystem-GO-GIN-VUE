<template>
  <div>
    <span>工作日志</span>
    <el-card>
      <el-main>
        <el-card>
          <span>日期：{{ this.creat_day }}</span>
          <div style="margin: 20px 0;"></div>
          <el-row type="flex">
            <span >标题：</span>
            <el-row>
              {{title}}
            </el-row>
          </el-row>
          <div style="margin: 20px 0;"></div>

          <el-row type="flex">
            <span >负责人：</span>
            <el-row>
              {{manager}}
            </el-row>
          </el-row>
          <div style="margin: 20px 0;"></div>

          <el-row type="flex">
            <span >成员：</span>
            <el-row>
              <el-input
                type="team_member"
                v-model="team_member">
              </el-input>
            </el-row>
          </el-row>
          <div style="margin: 20px 0;"></div>

        </el-card>


        <div style="margin: 20px 0;"></div>
        <span>目的：</span>
        <el-input
          type="goal"
          autosize
          placeholder="请输入内容"
          v-model="goal">
        </el-input>
        <span>日志内容：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="results">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <el-button @click="xiugai_rizhi">
          确认
        </el-button>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "gonzuorizhi_look",
  data() {
    return {
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
  methods: {
    xiugai_rizhi(){
      const api="http://localhost:3000/api/v1/student/worklogdupdate"

      let e=new URLSearchParams()
      e.append("team_member",this.team_member)
      e.append("goal",this.goal)
      e.append("results",this.results)
      e.append("title",this.title)
      e.append("userid",this.userid)
      e.append("id",this.id)
      e.append("class",this.class)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "修改成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/students/gonzuorizhi")
      }).catch(err => {
        this.$message({
          message: "修改失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    }
  },
  mounted() {
    this.title = this.$route.params.title
    this.manager = this.$route.params.manager
    this.team_member = this.$route.params.team_member
    this.goal = this.$route.params.goal
    this.results = this.$route.params.results
    this.creat_day = this.$route.params.creat_day
    this.id = this.$route.params.id
    this.userid = this.$route.params.userid
  }
}
</script>

<style scoped>

</style>
