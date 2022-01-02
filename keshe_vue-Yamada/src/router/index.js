import Vue from 'vue'
import VueRouter from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Login from "../Login";
import Loginfront from "../Loginfront";
import Register from "../Register";
import Root_login from "../Root_login";
import Root_set from "../Root_set";


////学生
import Students_home from "../students/home"
//打卡
import Students_kaoqin from "../students/kaoqin"
//实习报告
import Students_shixibaogao from "../students/shixibaogao/shixibaogao"
import Students_shixibaogao_create from "../students/shixibaogao/shixibaogao_create"
import Students_shixibaogao_update from "../students/shixibaogao/shixibaogao_update"
//请假
import Students_qinjia from "../students/qinjia/qinjia"
import Students_qinjia_create from "../students/qinjia/qinjia_create"
import Students_qinjia_look from "../students/qinjia/qinjia_look"
//工作日志
import Students_gonzuorizhi from "../students/gonzuorizhi/gonzuorizhi"
import Students_gonzuorizhi_create from "../students/gonzuorizhi/gonzuorizhi_create"
import Students_gonzuorizhi_look from "../students/gonzuorizhi/gonzuorizhi_look"
//实习通知
import Students_shixitonzhi from "../students/shixitonzhi/shixitonzhi"
import Students_shixitonzhi_look from "../students/shixitonzhi/shixitonzhi_look"
//企业需求
import Students_qiyexuqiu from "../students/qiyexuqiu/qiyexuqiu"
import Students_qiyexuqiu_look from "../students/qiyexuqiu/qiyexuqiu_look"
//学业成绩
import Students_chengji from "../students/chengji/chengji"
//考核成绩
import Students_kaohe from "../students/kaohe/kaohe"
//实时考勤
import Students_dinwei from "../Students/dinwei/dinwei"
////

////教师
import Teachers_home from "../teachers/home"
//打卡
import Teachers_kaoqin from "../teachers/kaoqin"
//实习报告
import Teachers_shixibaogao from "../teachers/shixibaogao/shixibaogao"
import Teachers_shixibaogao_look from "../teachers/shixibaogao/shixibaogao_look"
//请假
import Teachers_qinjia from "../teachers/qinjia/qinjia"
import Teachers_qinjia_look from "../teachers/qinjia/qinjia_look"
//工作日志
import Teachers_gonzuorizhi from "../teachers/gonzuorizhi/gonzuorizhi"
import Teachers_gonzuorizhi_look from "../teachers/gonzuorizhi/gonzuorizhi_look"
//实习通知
import Teachers_shixitonzhi from "../teachers/shixitonzhi/shixitonzhi"
import Teachers_shixitonzhi_create from "../teachers/shixitonzhi/shixitonzhi_create"
import Teachers_shixitonzhi_update from "../teachers/shixitonzhi/shixitonzhi_update"
//企业需求
import Teachers_qiyexuqiu from "../teachers/qiyexuqiu/qiyexuqiu"
import Teachers_qiyexuqiu_update from "../teachers/qiyexuqiu/qiyexuqiu_update"
import Teachers_qiyexuqiu_create from "../teachers/qiyexuqiu/qiyexuqiu_create"
//学业成绩
import Teachers_chengji from "../teachers/chengji/chengji"
import Teachers_chengji_create from "../teachers/chengji/chengji_create"
//考核成绩
import Teachers_kaohe from "../teachers/kaohe/kaohe"
import Teachers_kaohe_create from "../teachers/kaohe/kaohe_create"
//实时考勤
import Teachers_dinwei from "../teachers/dinwei/dinwei"
import axios from "axios";
////

Vue.use(VueRouter)
const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

axios.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
  return config
})

export default new VueRouter({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld,
    },
    //管理员
    {
      path: '/api/v1/Locallogin',
      name: 'Root_login',
      component: Root_login,
    }, {
      path: '/api/v1/user/update',
      name: 'Root_set',
      component: Root_set,
    }, //注册
    {
      path: '/api/v1/user/add',
      name: 'Register',
      component: Register,
    },
    //登陆
    {
      path: '/api/v1/Login',
      name: 'Login',
      component: Login,
    },{
      path: '/api/v1/Loginfront',
      name: 'Loginfront',
      component: Loginfront,
    },
    //home
    {
      path: '/students/home',
      name: 'students_home',
      component: Students_home,
    },{
      path: '/teachers/home',
      name: 'teachers_home',
      component: Teachers_home,
    },
    //考勤打卡
    {
      path: '/students/kaoqin',
      name: 'students_kaoqin',
      component: Students_kaoqin,
    }, {
      path: '/teachers/kaoqin',
      name: 'teachers_kaoqin',
      component: Teachers_kaoqin,
    },
    //实习报告
    {
      path: '/students/shixibaogao',
      name: 'students_shixibaogao',
      component: Students_shixibaogao,
    },{
      path: '/students/shixibaogao/create',
      name: 'students_shixibaogao_create',
      component: Students_shixibaogao_create,
    },{
      path: '/students/shixibaogao/update',
      name: 'students_shixibaogao_update',
      component:Students_shixibaogao_update
    },{
      path: '/teachers/shixibaogao',
      name: 'teachers_shixibaogao',
      component:Teachers_shixibaogao
    },{
      path: '/teachers/shixibaogao/look',
      name: 'teachers_shixibaogao_look',
      component:Teachers_shixibaogao_look
    },
    //请假
    {
      path: '/students/qinjia',
      name: 'students_qinjia',
      component: Students_qinjia,
    },{
      path: '/students/qinjia/create',
      name: 'students_qinjia_create',
      component: Students_qinjia_create,
    },{
      path: '/students/qinjia/look',
      name: 'students_qinjia_look',
      component: Students_qinjia_look,
    },
    {
      path: '/teachers/qinjia',
      name: 'teachers_qinjia',
      component: Teachers_qinjia,
    },
    {
      path: '/teachers/qinjia/look',
      name: 'teachers_qinjia_look',
      component: Teachers_qinjia_look,
    },
    //工作日志Students_gonzuorizhi
    {
      path: '/students/gonzuorizhi',
      name: 'students_gonzuorizhi',
      component: Students_gonzuorizhi,
    },
    {
      path: '/students/gonzuorizhi/create',
      name: 'students_gonzuorizhi_create',
      component: Students_gonzuorizhi_create,
    },{
      path: '/students/gonzuorizhi/look',
      name: 'students_gonzuorizhi_look',
      component: Students_gonzuorizhi_look,
    },
    {
      path: '/teachers/gonzuorizhi',
      name: 'teachers_gonzuorizhi',
      component: Teachers_gonzuorizhi,
    },{
      path: '/teachers/gonzuorizhi/look',
      name: 'teachers_gonzuorizhi_look',
      component: Teachers_gonzuorizhi_look,
    },
    //实习通知
    {
      path: '/students/shixitonzhi',
      name: 'students_shixitonzhi',
      component: Students_shixitonzhi,
    },{
      path: '/students/shixitonzhi/look',
      name: 'students_shixitonzhi_look',
      component: Students_shixitonzhi_look,
    },{
      path: '/teachers/shixitonzhi',
      name: 'teachers_shixitonzhi',
      component: Teachers_shixitonzhi,
    },{
      path: '/teachers/shixitonzhi/create',
      name: 'teachers_gonzuorizhi_create',
      component: Teachers_shixitonzhi_create,
    },{
      path: '/teachers/shixitonzhi/update',
      name: 'teachers_shixitonzhi_update',
      component: Teachers_shixitonzhi_update,
    },
    //企业需求
    {
      path: '/students/qiyexuqiu',
      name: 'students_qiyexuqiu',
      component: Students_qiyexuqiu,
    },{
      path: '/students/qiyexuqiu/look',
      name: 'students_qiyexuqiu_look',
      component: Students_qiyexuqiu_look,
    }, {
      path: '/teachers/qiyexuqiu',
      name: 'teachers_qiyexuqiu',
      component: Teachers_qiyexuqiu,
    },{
      path: '/teachers/qiyexuqiu/create',
      name: 'teachers_qiyexuqiu_create',
      component: Teachers_qiyexuqiu_create,
    },{
      path: '/teachers/qiyexuqiu/update',
      name: 'teachers_qiyexuqiu_update',
      component: Teachers_qiyexuqiu_update,
    },
    //学业成绩
    {
      path: '/students/chengji',
      name: 'students_chengji',
      component: Students_chengji,
    },{
      path: '/teachers/chengji',
      name: 'teachers_chengji',
      component: Teachers_chengji,
    },{
      path: '/teachers/chengji/create',
      name: 'teachers_chengji_create',
      component: Teachers_chengji_create,
    },
    //考核成绩
    {
      path: '/students/kaohe',
      name: 'students_kaohe',
      component: Students_kaohe,
    },{
      path: '/teachers/kaohe',
      name: 'teachers_kaohe',
      component: Teachers_kaohe,
    },{
      path: '/teachers/kaohe/create',
      name: 'teachers_kaohe_create',
      component: Teachers_kaohe_create,
    },
    //实时考勤
    {
      path: '/teachers/dinwei',
      name: 'teachers_dinwei',
      component: Teachers_dinwei,
    },{
      path: '/students/dinwei',
      name: 'students_dinwei',
      component: Students_dinwei,
    },
  ]
})
