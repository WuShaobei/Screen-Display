// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

// Cookie
import VueCookies from 'vue-cookies'
Vue.use(VueCookies)


// axios
// import axios from 'axios'
// // 全局设置 axios 发送请求带上cookie
// axios.defaults.withCredentials = true

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
