import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'

Vue.use(VueResource);

Vue.http.options.root = "http://localhost:6060"


export const eventBus = new Vue({
  methods:{
    //add global event methods
    refresh(){
      var chartData = {name:"chart"}
      var mapData = {name:"map"}
      this.$emit('refreshChart', chartData)
      this.$emit('refreshMap', mapData)
    },
    callRefreshMapData(map,data){
      this.$emit('refreshMapData', map, data)
    }    
  }
});

new Vue({
  el: '#app',
  render: h => h(App)
})


