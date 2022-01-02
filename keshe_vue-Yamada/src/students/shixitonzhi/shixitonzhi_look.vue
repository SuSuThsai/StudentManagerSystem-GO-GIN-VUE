<template>
  <div>
    <span>实习通知</span>
    <el-card>
      <el-main>
        <el-card>
          <span>日期：{{ this.creat_day }}</span>
          <div style="margin: 20px 0;"></div>
          <span>最近更新：{{ this.update_day }}</span>
          <div style="margin: 20px 0;"></div>
          <el-row type="flex">
            <span >标题：</span>
            <el-row>
              {{title}}
            </el-row>
          </el-row>

        </el-card>


        <div style="margin: 20px 0;"></div>
        <span>通知内容：</span>
        <el-input
          type="textarea"
          :readonly="true"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="text">
        </el-input>
        <div style="margin: 20px 0;"></div>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "shixitonzhi_look",
  data() {
    return {
      title: "",
      text: "",
      creat_day: "",
      update_day: ""
    }
  },
  methods: {},
  mounted() {
    this.text = this.$route.params.text
    this.title = this.$route.params.title
    this.creat_day = this.$route.params.creat_day
    this.update_day = this.$route.params.update_day
  }
}
</script>

<style scoped>

</style>
