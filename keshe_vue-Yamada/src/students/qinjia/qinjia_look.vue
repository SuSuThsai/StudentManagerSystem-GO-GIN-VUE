<template>
  <div>
    <span>请假申请</span>
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
        <el-row type="flex" v-if="status==1">
          <el-button type="xiugai_qingjia" @click="xiugai_qingjia">确定</el-button>
        </el-row>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "qinjia_look",
  data(){
    return{
      last_for: "",
      reason:"",
      status:""
    }
  },
  methods:{
    xiugai_qingjia(){
      const api="http://localhost:3000/api/v1/student/leaveupdate"

      let e=new URLSearchParams()
      e.append("last_for",this.last_for)
      e.append("reason",this.reason)
      e.append("userid",this.userid)
      e.append("id",this.id)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "修改成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/students/qinjia")
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
    this.last_for=this.$route.params.last_for
    this.reason=this.$route.params.reason
    this.id=this.$route.params.id
    this.status=this.$route.params.status
    this.userid=this.$route.params.userid
  }
}
</script>

<style scoped>

</style>
