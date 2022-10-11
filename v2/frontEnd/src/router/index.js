// /$$
//  $ @Author: WuShaobei
//  $ @Date: 2022-10-09 14:55:13
//  $ @LastEditTime: 2022-10-09 16:54:08
//  $ @FilePath: /Screen-Display/v2/frontEnd/src/router/index.js
//  $ @Description: 路由设置
//  $/
import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/views/login'
import Display from '@/views/display'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/display',
      name: 'Display',
      component: Display
    }
  ]
})
