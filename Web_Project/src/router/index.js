import Vue from 'vue'
import Router from 'vue-router'

import Login from '@/views/login/Login'
import SalesVolume from '@/views/data/SalesVolume'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
    },
    {
      path: '/data/salesVolume',
      name: 'SalesVolume',
      component: SalesVolume,
    },
  ]
})
