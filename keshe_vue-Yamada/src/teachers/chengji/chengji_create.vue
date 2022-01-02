<template>
  <div>
    <span>添加学生成绩</span>
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
<!--        <span>科目：</span>-->
<!--        <el-input-->
<!--          type="Title"-->
<!--          autosize-->
<!--          placeholder="请输入科目"-->
<!--          v-model="kemu">-->
<!--        </el-input>-->
<!--        <div style="margin: 20px 0;"></div>-->
        <span>学业成绩：</span>
        <el-input
          type="academic_score"
          autosize
          placeholder="请输入成绩"
          v-model="academic_score">
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
  name: "chengji_create",
  data(){
    return{
      academic_score:"",
      userid:"",
      id:""
    }
  },
  methods:{
    create_baogao(){
      const api="http://localhost:3000/api/v1/admin/gpaupdate"

      let e=new URLSearchParams()
      e.append("id",this.ID)
      e.append("userid",this.userid)
      e.append("academic_score",this.academic_score)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "创建成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/teachers/chengji")
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
