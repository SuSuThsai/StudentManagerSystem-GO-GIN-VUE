<template>
  <div>
    <span>企业实习需求</span>
    <el-card>
      <el-main>
        <el-card>
          <span>日期：{{ this.creat_day }}</span>
          <div style="margin: 20px 0;"></div>
          <span>截至日期：{{ this.end_day }}</span>
          <div style="margin: 20px 0;"></div>
          <span>日期：{{ this.update_day }}</span>
          <div style="margin: 20px 0;"></div>
          <el-row type="flex">
            <span >企业：</span>
            <el-row>
              {{company_name}}
            </el-row>
          </el-row>
          <div style="margin: 20px 0;"></div>
          <el-row type="flex">
            <span >职位：</span>
            <el-row>
              {{station}}
            </el-row>
          </el-row>
        </el-card>


        <div style="margin: 20px 0;"></div>
        <span>职位要求：</span>
        <el-input
          type="textarea"
          :readonly="true"
          :autosize="{ minRows: 20 }"
          placeholder="无"
          v-model="position_description">
        </el-input>
        <div style="margin: 20px 0;"></div>
      </el-main>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "qiyexuqiu_look",
  data() {
    return {
      position_description:"",
      company_name:"",
      creat_day:"",
      station:"",
      end_day:"",
      update_day:"",
    }
  },
  methods: {},
  mounted() {
    this.position_description = this.$route.params.position_description
    this.company_name = this.$route.params.company_name
    this.creat_day = this.$route.params.creat_day
    this.station = this.$route.params.station
    this.end_day = this.$route.params.end_day
    this.update_day = this.$route.params.update_day
  }
}
</script>

<style scoped>

</style>
