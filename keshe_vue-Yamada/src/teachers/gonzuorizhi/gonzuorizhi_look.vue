<template>
  <div>
    <span>工作日志</span>
    <el-card>
      <el-main>
        <el-card>
          <span>日期：{{ this.creat_day }}</span>
          <div style="margin: 20px 0;"></div>
          <el-row type="flex">
            <span >标题：</span>
            <el-row>
              {{title}}
            </el-row>
          </el-row>
          <div style="margin: 20px 0;"></div>

          <el-row type="flex">
            <span >负责人：</span>
            <el-row>
              {{manager}}
            </el-row>
          </el-row>
          <div style="margin: 20px 0;"></div>

          <el-row type="flex">
            <span >成员：</span>
            <el-row>
              {{team_member}}
            </el-row>
          </el-row>
          <div style="margin: 20px 0;"></div>

        </el-card>


        <div style="margin: 20px 0;"></div>
        <span>目的：</span>
        <el-input
          type="goal"
          autosize
          placeholder="请输入内容"
          v-model="goal"
          :readonly="true">
        </el-input>
        <span>日志内容：</span>
        <el-input
          type="textarea"
          :readonly="true"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="results">
        </el-input>
        <div style="margin: 20px 0;"></div>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "gonzuorizhi_look",
  data() {
    return {
      title:"",
      manager:"",
      team_member:"",
      goal:"",
      results:"",
      creat_day:""
    }
  },
  methods: {},
  mounted() {
    this.title = this.$route.params.title
    this.manager = this.$route.params.manager
    this.team_member = this.$route.params.team_member
    this.goal = this.$route.params.goal
    this.results = this.$route.params.results
    this.creat_day = this.$route.params.creat_day
  }
}
</script>

<style scoped>

</style>
