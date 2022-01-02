<template>
  <div>
    <span>修改实习通知</span>
    <el-card>
      <el-main>
        <span>标题：</span>
        <el-input
          type="title"
          autosize
          placeholder="请输入内容"
          v-model="title">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>通知内容：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="text">
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
  name: "shixitonzhi_update",
  data(){
    return{
      title: "",
      text: "",
      id:"",
      userid:""
    }
  },
  methods:{
    create_baogao(){
      const api="http://localhost:3000/api/v1/admin/noticeupdate"

      let e=new URLSearchParams()
      e.append("title",this.title)
      e.append("text",this.text)
      e.append("id",this.id)
      e.append("userid",this.userid)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "修改成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/teachers/shixitonzhi")
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
    this.title=this.$route.params.title
    this.text=this.$route.params.text
    this.id=this.$route.params.id
    this.userid=this.$route.params.userid
    console.log(this.id)
  }
}
</script>

<style scoped>

</style>
