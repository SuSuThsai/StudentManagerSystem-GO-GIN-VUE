<template>
  <div>
    <span>创建请假申请</span>
    <el-card>
      <el-main>
        <span>请假天数：</span>
        <el-input
          type="last_for"
          placeholder="请输入内容"
          v-model="last_for">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>请假原因：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="reason">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <el-button @click="create_qinjia">
          确认
        </el-button>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "qinjia_create",
  data(){
    return{
      last_for:"",
      reason: ""
    }
  },
  methods:{
    create_qinjia(){
      const api="http://localhost:3000/api/v1/student/leavepost"

      let e=new URLSearchParams()
      e.append("reason",this.reason)
      e.append("last_for",this.last_for)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "创建成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/students/qinjia")
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
