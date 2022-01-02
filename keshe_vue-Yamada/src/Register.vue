<template>
  <el-main>
    <div class="register" style="width: 100%;">
      <el-card title="登陆" style="opacity: 0.7">
        <el-button @click="set1">改变权限</el-button>
        <label>注册</label>
        <el-form label-width="50%">
          <el-form-item label="学号:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="userid"></el-input>
          </el-form-item>
          <el-form-item label="姓名:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="username"></el-input>
          </el-form-item>
          <el-form-item label="班级:" style="width: 35%; position:relative;
    left:30%;top:35px">
            <el-input type="text" v-model="clas"></el-input>
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
  name: 'Register',
  data() {
    return {
      userid: '1906100150',
      username: 'lalala7',
      clas: '计科191',
      password: '1906120150',
    }
  },
  methods: {
    set1(){
      this.$router.push("/api/v1/user/update")
    },
    open() {
      this.GLOBAL.daohan = this.$route.path
      console.log(this.GLOBAL.daohan)
    },
    login() {
      var params = new URLSearchParams();
      params.append("userid", this.userid);
      params.append("username", this.username);
      params.append("password", this.password)
      params.append("class", this.clas)

      //请求
      const api = 'http://localhost:3000/api/v1/user/add'

      this.axios.post(api, params).then(res => {
        //保存token
        console.log(res.data)
        // storagService.set(storagService.USER_TOKEN,res.data.data.token)
        // console.log(storagService.USER_TOKEN)
        // localStorage.setItem('token', res.data.token)
        sessionStorage.setItem('token', res.data.token)
        console.log("root:", res.data.root)
        this.$message({
          message: "注册成功，跳转登陆页" ,
          type: 'success'
        })
        //跳转登陆页
        this.$router.push("/Login")
        console.log("IsYamada")
      }).catch(err => {
        this.$message({
          message: "注册失败：" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    },
  },
  mounted: function () {
    this.open()
    this.$forceUpdate()
  }
}
</script>

<style>

</style>
