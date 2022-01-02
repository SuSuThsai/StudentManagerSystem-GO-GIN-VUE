<template>
  <div>
    <el-button size="medium" type="success" @click="daka">
      打卡
    </el-button>
    <el-table
      ref="multipleTable"
      :data="data"
      style="width: 100%">
<!--      <el-table-column-->
<!--        type="selection"-->
<!--        width="55">-->
<!--      </el-table-column>-->
      <el-table-column
        sortable
        prop="sign_day"
        label="日期"
        width="200">
<!--        <template slot-scope="scope">{{ scope.row.date }}</template>-->
      </el-table-column>
      <el-table-column
        prop="userid"
        label="学号"
        width="200">
      </el-table-column>
      <el-table-column
        prop="username"
        label="姓名"
        width="200">
      </el-table-column>
      <el-table-column
        prop="class"
        label="班级"
        width="300">
      </el-table-column>
      <el-table-column
        prop="is_value"
        label="打卡记录"
        show-overflow-tooltip>
      </el-table-column>
    </el-table>
    <!--  <div style="margin-top: 20px">-->
    <!--&lt;!&ndash;    <el-button @click="toggleSelection([tableData[1], tableData[2]])">切换第二、第三行的选中状态</el-button>&ndash;&gt;-->
    <!--&lt;!&ndash;    <el-button @click="toggleSelection()">取消选择</el-button>&ndash;&gt;-->
    <!--  </div>-->
  </div>
</template>

<script>
export default {
  name: "students_kaoqin",
  data() {
    return {
      data: [],
    }
  },
  methods: {
    daka(){
      const api = 'http://localhost:3000/api/v1/student/sigin'
      this.axios.post(api).then(res => {
        this.data = res.data.data
        for(var i=0;i<this.data.length;i++){
          if(this.data[i].is_value) {
            this.data[i].is_value = "已打卡"
          }else {
            this.data[i].is_value=  "未打卡"
          }
        }
          this.$message({
            message: "打卡成功:" + res.data.message,
            type: 'success'
          })
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
    if (this.data.length!=0){
      return
    }
    const api = 'http://localhost:3000/api/v1/student/signcheckitself'
    this.axios.get(api).then(res => {
      this.data = res.data.data
      for(var i=0;i<this.data.length;i++){
        if(this.data[i].is_value) {
          this.data[i].is_value = "已打卡"
        }else {
          this.data[i].is_value=  "未打卡"
        }
      }
      if (res.data.is_value==true){
        this.daka_today_state=true
        this.$message({
          message: "已打卡:" + res.data.message,
          type: 'success'
        })
      }
      console.log("data.data:",res.data.message)
      console.log("data:",this.data)
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
