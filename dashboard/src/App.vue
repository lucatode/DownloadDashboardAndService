<template>
  <div id="app">
    <!-- <app-header></app-header> -->
    <dashboard></dashboard>
    <app-footer></app-footer>
  </div>
</template>

<script>

import Header from './components/Shared/Header.vue'
import Dashboard from './components/Dashboard.vue'
import Footer from './components/Shared/Footer.vue'

export default {
  name: 'app',
  components:{
   'app-header': Header,
   'dashboard': Dashboard,
   'app-footer': Footer,
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
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

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
