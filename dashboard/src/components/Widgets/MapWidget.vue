<template>
    <div>       
        <ChoroplethMap 
            :data="pyDepartmentsData" 
            titleKey="department_name" 
            idKey="id" :value="value" 
            :extraValues="extraValues" 
            geojsonIdKey="id" 
            :geojson="worldGeojson" 
            :center="center" 
            :colorScale="colorScale" 
            mapStyle="height: 400px;" 
            :zoom="3" 
            :mapOptions="mapOptions">
        <template scope="props">
            <InfoControl :item="props.currentItem" :unit="props.unit" title="Country" placeholder="Hover over a country"></InfoControl>
        </template>
        </ChoroplethMap>
    </div>
</template>
<script>
    import {eventBus} from './../../main'
    import  ChoroplethMap from './module/ChoroplethMap.vue'
    import InfoControl from './module/InfoControl.vue'
    import ReferenceChart from './module/ReferenceChart.vue'
    import { geojson } from './data/py-departments-geojson'
    import worldGeojson from './data/world.json'
    import { pyDepartmentsData } from './data/py-departments-data'
    import { pyDepartmentsDataZero } from './data/py-departments-data-zero'

    export default{
        components:{
            ChoroplethMap,
            InfoControl, 
            ReferenceChart
        },
        data(){ 
            return{
                center: [0, 0],
                notUpdating: true,
                colorScale: ["000000", "cccccc", "ffffff"],
                value: {
                    key: "amount_w",
                    metric: " downloads"
                },
                extraValues: [{
                    key: "amount_m",
                    metric: " downloads"
                }],
                mapOptions: {
                    attributionControl: true
                },
                zero: false
            }
        },
        computed:{
            pyDepartmentsData: function(){
                return this.zero ? pyDepartmentsDataZero: pyDepartmentsData
            },
            inverseDepData: function(){
                return !this.zero ? pyDepartmentsDataZero: pyDepartmentsData
            },
            worldGeojson: function(){
                return worldGeojson
            }

        },
        methods:{  
        },
        created(){ //on component created
            eventBus.$on('refreshMap',(data)=>{ //register on event
                if(this.zero === false){
                    eventBus.callRefreshMapData(this.worldGeojson,this.pyDepartmentsDataZero )
                    console.log('zero')
                }else{
                    eventBus.callRefreshMapData(this.worldGeojson,this.pyDepartmentsData )
                    console.log('normal')
                }
                this.$forceUpdate()
                this.zero = !this.zero 
         
            })
        }
    }
</script>
<style>
@import "../../../node_modules/leaflet/dist/leaflet.css";
#map {
  background-color: #eee;
}
</style>