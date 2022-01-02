// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import axios from 'axios'
import VueAxios from 'vue-axios'
import global_ from './Global'


import VueAMap from 'vue-amap';
Vue.use(VueAMap);
// 初始化vue-amap
VueAMap.initAMapApiLoader({
  key: '0efdc2cde0fd075c86ce0939020d62a7'
});

Vue.use(ElementUI);
Vue.use(VueAxios,axios);
Vue.prototype.GLOBAL=global_

/* eslint-disable no-new */
new Vue({
  router,
  el: '#app',
  components: {App},
  template: '<App/>',
  render :h => h(App)
});


