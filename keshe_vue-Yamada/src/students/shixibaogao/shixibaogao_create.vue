<template>
  <div>
    <span>创建实习报告</span>
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
        <span>描述：</span>
        <el-input
          type="desc"
          autosize
          placeholder="请输入内容"
          v-model="desc">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>报告内容：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="content">
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
  name: "shixibaogao_create",
  data(){
    return{
      title: "",
      desc:"",
      content: ""
    }
  },
  methods:{
    create_baogao(){
      const api="http://localhost:3000/api/v1/student/reportpost"

      let e=new URLSearchParams()
      e.append("title",this.title)
      e.append("desc",this.desc)
      e.append("content",this.content)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "创建成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/students/shixibaogao")
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
