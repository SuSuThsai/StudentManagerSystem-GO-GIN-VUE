<template>
  <div>
    <span>创建企业实习需求</span>
    <el-card>
      <el-main>
        <span>企业名：</span>
        <el-input
          type="company_name"
          autosize
          placeholder="请输入内容"
          v-model="company_name">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>职位名称：</span>
        <el-input
          type="station"
          autosize
          placeholder="请输入内容"
          v-model="station">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>截止日期：</span>
        <el-input
          type="end_day"
          autosize
          placeholder="请输入内容"
          v-model="end_day">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>职位要求：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="position_description">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <el-button @click="create_baogao">
          确认
        </el-button>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "qiyexuqiu_create",
  data(){
    return{
      position_description:"",
      company_name:"",
      creat_day:"",
      station:"",
      end_day:"",
      update_day:"",
      id:"",
      userid:""
    }
  },
  methods:{
    create_baogao(){
      const api="http://localhost:3000/api/v1/admin/demandpost"

      let e=new URLSearchParams()
      e.append("company_name",this.company_name)
      e.append("station",this.station)
      e.append("end_day",this.end_day)
      e.append("position_description",this.position_description)
      this.axios.post(api, e).then(res => {
        this.$message({
          message: "创建成功:" + res.data.message,
          type: 'success'
        })
        this.$router.push("/teachers/qiyexuqiu")
      }).catch(err => {
        this.$message({
          message: "创建失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    }
  }
}
</script>

<style scoped>

</style>
