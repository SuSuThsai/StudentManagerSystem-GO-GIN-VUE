<template>
  <div>
    <span>添加学生考核成绩</span>
    <el-card>
      <el-main>
<!--        <span>学期(yyyy-mm-dd)：</span>-->
<!--        <el-input-->
<!--          type="Title"-->
<!--          autosize-->
<!--          placeholder="请输入学期"-->
<!--          v-model="xueqi">-->
<!--        </el-input>-->
<!--        <div style="margin: 20px 0;"></div>-->
        <span>学号：</span>
        <el-input
          type="userid"
          autosize
          placeholder="请输入学号"
          v-model="userid">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>工作能力：</span>
        <el-input
          type="working_ability"
          autosize
          placeholder="请输入成绩"
          v-model="working_ability">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>沟通能力：</span>
        <el-input
          type="communication"
          autosize
          placeholder="请输入成绩"
          v-model="communication">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>工作态度：</span>
        <el-input
          type="working_attitude"
          autosize
          placeholder="请输入成绩"
          v-model="working_attitude">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <el-button @click="create_baogao">
          确认
        </el-button>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "kaohe_create",
  data(){
    return{
      userid:"",
      working_ability: "",
      working_attitude: "",
      communication:"",
      id:""
    }
  },
  methods:{
    create_baogao(){
      const api="http://localhost:3000/api/v1/admin/assessmentupdate"

      let e=new URLSearchParams()
      e.append("userid",this.userid)
      e.append("id",this.ID)
      e.append("working_ability",this.working_ability)
      e.append("working_ability",this.working_ability)
      e.append("communication",this.communication)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "创建成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/teachers/kaohe")
      }).catch(err => {
        this.$message({
          message: "创建失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    }
  }
}
</script>

<style scoped>

</style>
