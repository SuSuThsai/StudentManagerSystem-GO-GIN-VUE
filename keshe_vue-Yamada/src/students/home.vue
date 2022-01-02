<template>

</template>

<script>
export default {
  name: "students_home",
  methods:{
    open(){
      this.GLOBAL.daohan=this.$route.path
      console.log(this.GLOBAL.daohan)
    }
  },
  mounted() {
    this.open()
    this.$forceUpdate()
  }
}
</script>

<style scoped>

</style>
