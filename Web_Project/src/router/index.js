import Vue from 'vue'
import Router from 'vue-router'

import Login from '@/views/login/Login'
import Data from '@/views/data/Data'
import Data1 from '@/views/data/Data1'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login,
    },
    {
      path: '/data/Data',
      name: 'Data',
      component: Data,
    },
    {
      path: '/data/Data1',
      name: 'Data1',
      component: Data1,
    },
  ]
})
