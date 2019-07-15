<template>
  <div id="app">
    <top-nav></top-nav>
    <router-view/>
  </div>
</template>

<script>
import TopNav from './components/Topnav.vue'

export default {
  name: 'App',
  components: {
    'top-nav': TopNav,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
