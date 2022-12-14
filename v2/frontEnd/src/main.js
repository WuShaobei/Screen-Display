///
 // @Date: 2022-10-09 14:55:13
 // @LastEditTime: 2022-10-14 17:57:40
 // @FilePath: /frontEnd/src/main.js
 // @Description: 
 ///
// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

// Cookie
import VueCookies from 'vue-cookies'
Vue.use(VueCookies)
// 三天面登录
VueCookies.config('3d')

// echarts
import 'echarts'
import ECharts from 'vue-echarts'
Vue.component('ECharts', ECharts)

import 'default-passive-events'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
