<template>
  <div>
    <span>实习报告</span>
    <el-card>
      <el-main>
        <span>标题：</span>
        <el-input
          type="title"
          autosize
          placeholder="请输入内容"
          v-model="title"
          :readonly="true">
        </el-input>
        <span>描述：</span>
        <el-input
          type="desc"
          autosize
          placeholder="请输入内容"
          v-model="desc"
          :readonly="true">
        </el-input>
        <div style="margin: 20px 0;"></div>
        <span>报告内容：</span>
        <el-input
          type="textarea"
          :autosize="{ minRows: 20 }"
          placeholder="请输入内容"
          v-model="content"
          :readonly="true">
        </el-input>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "shixibaogao_look",
  data(){
    return{
      title: "",
      desc: "",
      content:"",
    }
  },
  methods:{
  },
  mounted() {
    this.title=this.$route.params.title
    this.desc=this.$route.params.desc
    this.content=this.$route.params.content
  }
}
</script>

<style scoped>

</style>
