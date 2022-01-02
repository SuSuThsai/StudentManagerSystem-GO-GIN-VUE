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
        width="200">
      </el-table-column>
      <el-table-column
        prop="username"
        label="姓名"
        width="150">
      </el-table-column>
      <el-table-column
        sortable
        prop="working_ability"
        label="工作能力分值"
        width="150"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        sortable
        prop="working_attitude"
        label="工作态度分值"
        width="150"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        sortable
        prop="communication"
        label="沟通能力分值"
        width="150"
        show-overflow-tooltip>
      </el-table-column>
      <el-table-column
        sortable
        prop="total_score"
        label="总成绩(工作能力*0.5+工作态度*0.25+沟通能力*0.25)"
        show-overflow-tooltip>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "teachers_kaohe",
  data() {
    return {
      data: [],
      userid:""
    }
  },
  methods: {
    create(){
      this.$router.push({
        name: 'teachers_kaohe_create',
      })
    },
    cha() {
      const api = "http://localhost:3000/api/v1/admin/assessmentcheckuserid"
      let e = new URLSearchParams()
      e.append("userid", this.userid)
      this.axios.get(api,{
          headers:{
            'userid': this.userid,
          }
        }).then(res => {
        this.data = res.data.data
        // let m_xueqi = new Set()
        // for (let i = 0; i < this.data.length; i++) {
        //   m_xueqi.add(this.data[i].Xueqi)
        // }
        // this.xueqi_=[]
        // for (var x of m_xueqi){
        //   this.xueqi_.push({text:x,value:x})
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
    const api = 'http://localhost:3000/api/v1/admin/assessmentcreate'
    this.axios.post(api).then(res => {
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
