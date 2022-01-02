<template>
  <div>
    <span>请假申请</span>
    <el-card>
      <el-main>
        <span>请假原因：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          :readonly="true"
          v-model="reason">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <el-row type="flex" v-if="status==1">
          <el-button type="success" @click="tijiao(true)">批准</el-button>
          <el-button type="danger" @click="tijiao(false)">拒绝</el-button>
        </el-row>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "qinjia_look",
  data() {
    return {
      userid: "",
      username: "",
      reason: "",
      id: "",
      status: 0,
    }
  },
  methods: {
    tijiao(t) {
      const api = "http://localhost:3000/api/v1/admin/leaveupdate"
      let e = new URLSearchParams()
      e.append("reason", this.reason)
      e.append("id", this.id)
      if (t) {
        e.append("status", 2)
      } else {
        e.append("status", 3)
      }
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/teachers/qinjia")
      }).catch(err => {
        this.$message({
          message: "失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    }
  },
  mounted() {
    this.userid = this.$route.params.userid
    this.username = this.$route.params.username
    this.reason = this.$route.params.reason
    this.id = this.$route.params.id
    this.status = this.$route.params.status
  }
}
</script>

<style scoped>

</style>
