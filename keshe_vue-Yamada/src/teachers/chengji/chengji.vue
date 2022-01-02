<template>
  <div>
    <div>
      <el-button @click="create" type="success">
        添加
      </el-button>
    </div>
    <el-card>
      <el-row type="flex">
        <span>学号：</span>
        <el-col :span="6">
          <el-input label="学号" type="text" v-model="userid" maxlength="11" size="small"></el-input>
        </el-col>
        <el-col :span="5" :offset="1">
          <el-button @click="cha" size="small">查询</el-button>
          <el-button @click="gengxin" size="small">更新成绩数据</el-button>
        </el-col>
      </el-row>
    </el-card>
    <el-table
      :data="data"
      style="width: 100%"
    >
<!--      <el-table-column-->
<!--        sortable-->
<!--        prop="Xueqi"-->
<!--        label="学期"-->
<!--        :filters="xueqi_"-->
<!--        :filter-method="xueqi_sai"-->
<!--        width="200">-->
<!--      </el-table-column>-->
      <el-table-column
        sortable
        prop="userid"
        label="学号"
        width="150">
      </el-table-column>
      <el-table-column
        prop="username"
        label="姓名"
        width="200">
      </el-table-column>
      <el-table-column
        sortable
        prop="academic_score"
        label="学业成绩"
        width="200">
      </el-table-column>
      <el-table-column
        sortable
        prop="internship_score"
        label="实习成绩"
        width="200">
      </el-table-column>
      <el-table-column
        sortable
        prop="total_score"
        label="总成绩(学业*0.75+实习*0.25)"
        show-overflow-tooltip>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "teachers_chengji",
  data() {
    return {
      data: [],
      userid:""
    }
  },
  methods: {
    create(){
      this.$router.push({
        name: 'teachers_chengji_create',
      })
    },
    cha() {
      const api = "http://localhost:3000/api/v1/admin/gpacheckuserid"
      let e = new URLSearchParams()
      e.append("userid", this.userid)
      this.axios.get(api, {
        headers:{
          'userid': this.userid,
        }
      }).then(res => {
        this.data = res.data.data
        // let m_xueqi = new Set()
        // let m_kemu = new Set()
        // for (let i = 0; i < this.data.length; i++) {
        //   m_xueqi.add(this.data[i].Xueqi)
        //   m_kemu.add(this.data[i].Kemu)
        // }
        // this.xueqi_=[]
        // this.kemu_=[]
        // for (var x of m_xueqi){
        //   this.xueqi_.push({text:x,value:x})
        // }
        // for (var x of m_kemu){
        //   this.kemu_.push({text:x,value:x})
        // }
        this.$message({
          message: "成功:" + res.data.message,
          type: 'success'
        })
      }).catch(err => {
        this.$message({
          message: "失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    },
    gengxin() {
      const api = "http://localhost:3000/api/v1/admin/gpacheckupdate"
      let e = new URLSearchParams()
      this.axios.post(api,null).then(res => {
        this.data = res.data.data
        this.$message({
          message: "成功:" + res.data.message,
          type: 'success'
        })
      }).catch(err => {
        this.$message({
          message: "失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    },
    // xueqi_sai(value, row) {
    //   return row.Xueqi === value;
    // },
    // kemu_sai(value, row) {
    //   return row.Kemu === value;
    // }
  },
  mounted() {
    if (this.data.length != 0) {
      return
    }
    const api = 'http://localhost:3000/api/v1/admin/gpacheckclass'
    this.axios.get(api).then(res => {
      this.data = res.data.data
      // let m_xueqi = new Set()
      // let m_kemu = new Set()
      // for (let i = 0; i < this.data.length; i++) {
      //   m_xueqi.add(this.data[i].Xueqi)
      //   m_kemu.add(this.data[i].Kemu)
      // }
      // for (var x of m_xueqi){
      //   this.xueqi_.push({text:x,value:x})
      // }
      // for (var x of m_kemu){
      //   this.kemu_.push({text:x,value:x})
      // }
    }).catch(err => {
      this.$message({
        message: "失败:" + err.response.data.message,
        type: 'error'
      })
      console.log('err:', err.response.data)
    })

  }
}
</script>

<style scoped>

</style>
