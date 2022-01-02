<template>
  <div>
    <div>
      <el-button @click="bg_create" type="success">
        新建
      </el-button>
    </div>
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
        prop="update_day"
        label="最近更新"
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
              <el-button @click="update(scope.row)" size="mini">编辑</el-button>
            </el-col>
            <el-col>
              <el-button @click="shanchu(scope.row)" type="danger" size="mini">删除</el-button>
            </el-col>
          </el-row>
        </template>
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
  name: "students_shixibaogao",
  data() {
    return {
      data: [],
    }
  },
  methods: {
    shanchu(bg) {
      console.log(bg)
      console.log(bg.ID)
      const api = "http://localhost:3000/api/v1/student/reportdelete"
      let e = new URLSearchParams()
      e.append("id", bg.ID)
      e.append("userid", bg.userid)
      e.append("class", bg.class)
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
    update(bg) {
      this.$router.push({
        name: 'students_shixibaogao_update',
        params: {title: bg.title, text: bg.text,content:bg.content,userid:bg.userid,id: bg.ID,desc:bg.desc}
      })
    },
    bg_create() {
      this.$router.push("/students/shixibaogao/create")
    }
  },
  mounted() {
    if (this.data.length != 0) {
      return
    }
    const api = 'http://localhost:3000/api/v1/student/reportcheckitself'
    this.axios.get(api).then(res => {
      this.data = res.data.data
    }).catch(err => {
      this.$message({
        message: "获取个人实习报告失败:" + err.response.data.message,
        type: 'error'
      })
      console.log('err:', err.response.data)
    })

  }
}
</script>

<style scoped>

</style>
