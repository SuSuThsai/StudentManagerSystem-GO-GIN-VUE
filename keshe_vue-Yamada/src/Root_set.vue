<template>
  <el-main >
    <div class="register" style="width: 100%;" >
      <el-card title="登陆" style="opacity: 0.7">
        <label>改变权限</label>
        <el-button @click="creat">注册</el-button>
        <el-form label-width="50%" >
          <el-form-item label="学号:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="userid"></el-input>
          </el-form-item>
          <el-form-item label="role:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-radio v-model="role" label=2>不给</el-radio>
            <el-radio v-model="role" label=1>给</el-radio>
          </el-form-item>
          <el-form-item style="width: 20%; position:relative;
    left:52%;top:35px">
            <el-button @click="login()">确认</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </el-main>
</template>

<script>
export default {
  name: "Root_set",
  data () {
    return {
        userid: '1906100045',
        role: 2,
    }
  },
  methods: {
    creat(){
      this.$router.push("/api/v1/root/add")
    },
    open(){
      this.GLOBAL.daohan=this.$route.path
      console.log(this.GLOBAL.daohan)
      const t=localStorage.getItem("token_root")
      console.log("token_root:",t)
      if (t==""||t==null){
        this.$router.replace({name:"Root_login"})
      }
    },
    login () {
      //请求
      var params=new URLSearchParams()
      params.append("userid",this.userid)
      params.append("role",this.role)

      const api='http://localhost:3000/api/v1/user/update'
      // this.param.token=localStorage.getItem("token")
      // this.headers.set("token",localStorage.getItem("token"))
      this.axios.post(api,params,{
        headers: {
          'token':`Bearer ${localStorage.getItem('token_root')}`,
        }
      }).then(res=>{
        // console.log(res.data)
          this.$message({
            message:"成功:"+res.data.message,
            type:'success'
          })
      }).catch(err=>{
        this.$message({
          message:"失败:"+err.response.data.message,
          type:'error'
        })
        console.log('err:',err.response.data)
      })
    },
  },
  mounted:function () {
    this.open()
    this.$forceUpdate()
  }
}
</script>

<style scoped>

</style>
