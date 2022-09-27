import Vue from 'vue'
import Router from 'vue-router'

import Login from '@/views/login/Login'
import Screen from '@/views/data/Screen'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
    },
    {
      path: '/data/Screen',
      name: 'Screen',
      component: Screen,
    },
  ]
})
