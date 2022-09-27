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
  font-size:20px
}
</style>
