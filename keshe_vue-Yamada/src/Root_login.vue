<template>
  <el-main>
    <div class="register" style="width: 100%;">
      <el-card title="登陆" style="opacity: 0.7">
        <label>管理员登陆（本地）</label>
        <el-form label-width="50%" >
          <el-form-item label="账号:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="username"></el-input>
          </el-form-item>
          <el-form-item label="密码:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="password"></el-input>
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
import qs from 'qs' // axios自带的工具不需要安装npm依赖
export default {
  name: "Root_login",
  data () {
    return {
        username: 'Yamada',
        password: 'cc123',
    }
  },
  methods: {
    open(){
      this.GLOBAL.daohan=this.$route.path
      console.log(this.GLOBAL.daohan)
    },
    login () {
      var params=new URLSearchParams();
      params.append("username",this.username);
      params.append("password",this.password)
      //请求
      const api='http://localhost:3000/api/v1/Locallogin'

      this.axios.post(api,params).then(res=>{
        //保存token
        console.log(res.data)
        // storagService.set(storagService.USER_TOKEN,res.data.data.token)
        // console.log(storagService.USER_TOKEN)
        console.log("IsYamada")
        localStorage.setItem('token_root',res.data.token)

        //跳转主页
        this.$router.push('/api/v1/user/update')
        location.reload()
        console.log("IsYamada")
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

<style scoped>

</style>
<!--{ headers: {-->
<!--'Content-Type': "application/x-www-form-urlencoded"-->
<!--}-->
