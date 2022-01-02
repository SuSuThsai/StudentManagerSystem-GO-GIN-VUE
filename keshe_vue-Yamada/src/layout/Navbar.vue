<template>
  <div>
    <el-col :span="12">
      <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        @open=""
        @close="">
        <!--        学生导航栏-->
        <el-menu-item index="1" v-if="!this.GLOBAL.root" @click="$router.push('/students/kaoqin')">
          <i class="el-icon-menu"></i>
          <span slot="title">打卡</span>
        </el-menu-item>
        <el-menu-item index="2" v-if="!this.GLOBAL.root" @click="$router.push('/students/shixibaogao')">
          <i class="el-icon-menu"></i>
          <span slot="title">实习报告</span>
        </el-menu-item>
        <el-menu-item index="3" v-if="!this.GLOBAL.root" @click="$router.push('/students/qinjia')">
          <i class="el-icon-menu"></i>
          <span slot="title">请假申请</span>
        </el-menu-item>
        <el-menu-item index="4" v-if="!this.GLOBAL.root" @click="$router.push('/students/gonzuorizhi')">
          <i class="el-icon-menu"></i>
          <span slot="title">工作日志</span>
        </el-menu-item>
        <el-menu-item index="5" v-if="!this.GLOBAL.root" @click="$router.push('/students/shixitonzhi')">
          <i class="el-icon-menu"></i>
          <span slot="title">实习通知</span>
        </el-menu-item>
        <el-menu-item index="6" v-if="!this.GLOBAL.root" @click="$router.push('/students/qiyexuqiu')">
          <i class="el-icon-menu"></i>
          <span slot="title">企业实习需求</span>
        </el-menu-item>
        <el-menu-item index="7" v-if="!this.GLOBAL.root" @click="$router.push('/students/chengji')">
          <i class="el-icon-menu"></i>
          <span slot="title">学业成绩</span>
        </el-menu-item>
        <el-menu-item index="8" v-if="!this.GLOBAL.root" @click="$router.push('/students/kaohe')">
          <i class="el-icon-menu"></i>
          <span slot="title">考核成绩</span>
        </el-menu-item>
        <el-menu-item index="18" v-if="!this.GLOBAL.root" @click="$router.push('/students/dinwei')">
          <i class="el-icon-menu"></i>
          <span slot="title">实时考勤</span>
        </el-menu-item>
        <!--        教师导航栏-->
        <el-menu-item index="9" v-if="this.GLOBAL.root" @click="$router.push('/teachers/kaoqin')">
          <i class="el-icon-menu"></i>
          <span slot="title">打卡统计</span>
        </el-menu-item>
        <el-menu-item index="10" v-if="this.GLOBAL.root" @click="$router.push('/teachers/shixibaogao')">
          <i class="el-icon-menu"></i>
          <span slot="title">实习报告</span>
        </el-menu-item>
        <el-menu-item index="11" v-if="this.GLOBAL.root" @click="$router.push('/teachers/qinjia')">
          <i class="el-icon-menu"></i>
          <span slot="title">请假申请</span>
        </el-menu-item>
        <el-menu-item index="12" v-if="this.GLOBAL.root" @click="$router.push('/teachers/gonzuorizhi')">
          <i class="el-icon-menu"></i>
          <span slot="title">工作日志</span>
        </el-menu-item>
        <el-menu-item index="13" v-if="this.GLOBAL.root" @click="$router.push('/teachers/shixitonzhi')">
          <i class="el-icon-menu"></i>
          <span slot="title">实习通知</span>
        </el-menu-item>
        <el-menu-item index="14" v-if="this.GLOBAL.root" @click="$router.push('/teachers/qiyexuqiu')">
          <i class="el-icon-menu"></i>
          <span slot="title">企业实习需求</span>
        </el-menu-item>
        <el-menu-item index="15" v-if="this.GLOBAL.root" @click="$router.push('/teachers/chengji')">
          <i class="el-icon-menu"></i>
          <span slot="title">学业成绩</span>
        </el-menu-item>
        <el-menu-item index="16" v-if="this.GLOBAL.root" @click="$router.push('/teachers/kaohe')">
          <i class="el-icon-menu"></i>
          <span slot="title">考核成绩</span>
        </el-menu-item>
        <el-menu-item index="17" v-if="this.GLOBAL.root" @click="$router.push('/teachers/dinwei')">
          <i class="el-icon-menu"></i>
          <span slot="title">实时考勤</span>
        </el-menu-item>
      </el-menu>
    </el-col>
  </div>
</template>

<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    open() {
      let ll = this.$route.path
      if (ll[1] == 't') {
        this.GLOBAL.root = true
      } else {
        this.GLOBAL.root = false
      }
    }
  },
  mounted() {
    this.open()
    this.$forceUpdate()
  }
}
</script>
