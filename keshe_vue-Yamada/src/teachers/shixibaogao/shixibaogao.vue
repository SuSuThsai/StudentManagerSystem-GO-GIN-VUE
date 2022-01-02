<template>
  <div>
    <el-main>

      <el-row type="flex">
        <span>学号：</span>
        <el-col :span="10">
          <el-input label="学号" type="content" v-model="userid" maxlength="11" size="small"></el-input>
        </el-col>
        <el-col :span="5" :offset="1">
          <el-button @click="cha" size="small">查询</el-button>
        </el-col>
      </el-row>
      <el-table
        :data="data"
        style="width: 100%">
        <el-table-column
          sortable
          prop="creat_day"
          label="日期"
          width="150">
        </el-table-column>
        <el-table-column
          sortable
          prop="userid"
          label="学号"
          width="150">
        </el-table-column>
        <el-table-column
          sortable
          prop="username"
          label="姓名"
          width="150">
        </el-table-column>
        <el-table-column
          prop="title"
          label="标题"
          show-overflow-tooltip>
        </el-table-column>
        <el-table-column
          prop="desc"
          label="描述"
          show-overflow-tooltip>
        </el-table-column>
        <el-table-column
          fixed="right"
          label="操作"
          width="200">
          <template slot-scope="scope">
            <el-row type="flex">
              <el-col :offset="1">
                <el-button @click="look(scope.row)" size="mini" >查看</el-button>
              </el-col>
              <el-col>
                <el-button @click="delete1(scope.row)" type="danger" size="mini">删除</el-button>
              </el-col>
            </el-row>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </div>
</template>

<script>
export default {
  name: "teachers_shixibaogao",
  data() {
    return {
      data: [],
      userid: ""
    }
  },
  methods: {
    delete1(bg) {
      console.log(bg)
      console.log(bg.ID)
      const api = "http://localhost:3000/api/v1/admin/reportdelete"
      let e = new URLSearchParams()
      e.append("id", bg.ID)
      e.append("class", bg.class)
      e.append("userid", bg.userid)
      this.axios.post(api, e).then(res => {
        for (let i = 0; i < this.data.length; i++) {
          if (this.data[i].ID == bg.ID) {
            this.data.splice(i, i)
            break
          }
        }
        console.log(this.data)
        this.$forceUpdate()
        this.$message({
          message: "删除成功:" + res.data.message,
          type: 'success'
        })
      }).catch(err => {
        this.$message({
          message: "删除失败:" + err.response.data.message,
          type: 'error'
        })
      })
    },
    look(bg) {
      this.$router.push({
        name: 'teachers_shixibaogao_look',
        params: {title: bg.title,desc:bg.desc, content: bg.content}
      })
    },
    cha() {
      const api = "http://localhost:3000/api/v1/admin/reportcheckuserid"
      let e = new URLSearchParams()
      e.append("userid", this.userid)
      this.axios.get(api,  {
        headers:{
          'userid': this.userid,
        }
      }).then(res => {
        this.data = res.data.data
        this.$message({
          message: "获取学号：" + this.userid + "的实习报告成功:" + res.data.message,
          type: 'success'
        })
      }).catch(err => {
        this.$message({
          message: "获取学号：" + this.userid + "的实习报告失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    }
  },
  mounted() {
    console.log(this.data)
    if (this.data.length != 0) {
      return
    }
    const api = 'http://localhost:3000/api/v1/admin/reportcheckclass'
    this.axios.get(api).then(res => {
      this.data = res.data.data
    }).catch(err => {
      this.$message({
        message: "获取全班实习报告失败:" + err.response.data.message,
        type: 'error'
      })
      console.log('err:', err.response.data)
    })
  }

}
</script>

<style scoped>

</style>
