<template>
  <div>
    <div class="block">
      <el-date-picker
        v-model="value1"
        value-format="yyyy-MM-dd"
        align="right"
        type="date"
        placeholder="选择日期"
        :picker-options="pickerOptions">
      </el-date-picker>
      <el-button @click="daka">
        确定
      </el-button>
    </div>
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
        sortable
        label="学号"
        width="200">
      </el-table-column>
      <el-table-column
        prop="username"
        sortable
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
        sortable
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
  name: "teachers_kaoqin",
  data() {
    return {
      data: [],
      pickerOptions: {
        disabledDate(time) {
          return time.getTime() > Date.now();
        },
        shortcuts: [{
          text: '今天',
          onClick(picker) {
            picker.$emit('pick', new Date());
          }
        }, {
          text: '昨天',
          onClick(picker) {
            const date = new Date();
            date.setTime(date.getTime() - 3600 * 1000 * 24);
            picker.$emit('pick', date);
          }
        }, {
          text: '一周前',
          onClick(picker) {
            const date = new Date();
            date.setTime(date.getTime() - 3600 * 1000 * 24 * 7);
            picker.$emit('pick', date);
          }
        }]
      },
      value1: '',
      value2: '',
    }
  },
  methods: {
    daka() {
      console.log("value1:",this.value1)
      let e = new URLSearchParams()
      e.append("sign_day", this.value1)
      const api = 'http://localhost:3000/api/v1/admin/siginchecksignday'
      this.axios.get(api, {
        headers:{
        'sign_day': this.value1,
      }
      }).then(res => {
        this.data = res.data.data
        for (var i = 0; i < this.data.length; i++) {
          if (this.data[i].is_value) {
            this.data[i].is_value = "已打卡"
          } else {
            this.data[i].is_value = "未打卡"
          }
        }
      }).catch(err => {
        this.$message({
          message: "获取学生打卡情况失败:" + err.response.data.message,
          type: 'error'
        })
        console.log('err:', err.response.data)
      })
    }
  },
  mounted() {
    const api = 'http://localhost:3000/api/v1/admin/sigincheckclass'
    this.axios.get(api, null).then(res => {
      this.data = res.data.data
      for (var i = 0; i < this.data.length; i++) {
        if (this.data[i].is_value) {
          this.data[i].is_value = "已打卡"
        } else {
          this.data[i].is_value = "未打卡"
        }
      }
      console.log("data.data:", res.data.data)
      console.log("data:", this.data)
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
