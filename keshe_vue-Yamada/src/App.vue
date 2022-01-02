<template>
  <div id="app">
    <h5 >ManagerSystem</h5>
    <el-container>
      <el-aside width="15%" v-if="this.GLOBAL.daohan!='/api/v1/Locallogin'&&this.GLOBAL.daohan!='/api/v1/user/update'&&this.GLOBAL.daohan!='/Login'&&this.GLOBAL.daohan!='/Register'">
        <Navbar >
        </Navbar>
      </el-aside>
      <el-main>
        <router-view/>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import Navbar from "./layout/Navbar";

export default {
  name: 'App',
  components: {Navbar},
  data() {
    return {
    }
  },
  methods:{

    open(){
      this.GLOBAL.daohan=this.$route.path
      console.log(this.GLOBAL.daohan)
    }
  },
  mounted() {
    this.open()
    this.$forceUpdate()
  }
}
</script>

<!--<style>-->
<!--/*#app {*/-->
<!--/*  font-family: 'Avenir', Helvetica, Arial, sans-serif;*/-->
<!--/*  -webkit-font-smoothing: antialiased;*/-->
<!--/*  -moz-osx-font-smoothing: grayscale;*/-->
<!--/*  text-align: center;*/-->
<!--/*  color: #2c3e50;*/-->
<!--/*  margin-top: 60px;*/-->
<!--/*}*/-->
<!--</style>-->
