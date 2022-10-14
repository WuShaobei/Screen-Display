<!--
 * @Date: 2022-10-09 14:55:13
 * @LastEditTime: 2022-10-14 17:04:54
 * @FilePath: /frontEnd/src/App.vue
 * @Description: 
-->
<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
  export default {
    name: 'App',
    watch: {
    $route(to, from) {
      //监测到路由发生跳转时刷新一次页面
      window.location.reload(); 
    },
  },
}
</script>

<style>
#app {
  text-align: center;
  color: rgb(255, 100, 0);
  /* font-size:20px */
}
</style>
