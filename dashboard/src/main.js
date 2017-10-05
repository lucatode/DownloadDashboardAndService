import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'
import Raphael from 'raphael/raphael'
global.Raphael = Raphael

Vue.use(VueResource);

Vue.http.options.root = "https://gomicroservice.herokuapp.com"


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
    },
    refreshView(){
      this.$forceUpdate()
    },
    sendCountryDataToWidget(name,value){
      this.$emit('sendCountryDataToWidget',name,value);
    } 
  }
});

new Vue({
  el: '#app',
  render: h => h(App)
})


