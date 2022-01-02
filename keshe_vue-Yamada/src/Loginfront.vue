<template>
  <el-main>
    <div class="register" style="width: 100%;">
      <el-card title="登陆" style="opacity: 0.7">
        <el-button @click="login2">教师登录</el-button>
        <label>学生登录</label>
        <el-form label-width="50%" >
          <el-form-item label="学号:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="userid"></el-input>
          </el-form-item>
          <el-form-item label="密码:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="password"></el-input>
          </el-form-item>
          <el-form-item style="width: 20%; position:relative;
    left:52%;top:35px">
            <el-button @click="login">确认</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </el-main>
</template>

<script>
export default {
  name: 'Loginfront',
  data () {
    return {
      userid: '1906100045',
      password: 'cc12345',
    }
  },
  methods: {
    login2(){
      this.$router.push("/api/v1/Login")
    },
    open(){
      this.GLOBAL.daohan=this.$route.path
      console.log(this.GLOBAL.daohan)
    },
    login () {
      var params=new URLSearchParams();
      params.append("userid",this.userid);
      params.append("password",this.password)

      //请求
      const api='http://localhost:3000/api/v1/Loginfront'

      this.axios.post(api,params).then(res=>{
        //保存token
        console.log(res.data)
        // storagService.set(storagService.USER_TOKEN,res.data.data.token)
        // console.log(storagService.USER_TOKEN)
        console.log("IsYamada")
        sessionStorage.setItem('token',res.data.token)
        sessionStorage.setItem('role',res.data.role)
        console.log("token:",res.data.token)
        console.log("role:",res.data.role)
        //跳转主页
        if (res.data.role==2){
          this.$router.push('/student/home')
          location.reload()
          this.$message({
            message:"登陆成功，当前为学生账号",
            type:'success'
          })
          return
        }
      }).catch(err=>{
        this.$message({
          message:"登陆失败"+err.response.data.message,
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

<style >

</style>
