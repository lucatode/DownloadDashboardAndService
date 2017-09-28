<template>
    <div>       
        <ChoroplethMap 
            :data="pyDepartmentsDataReal" 
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
                center: [45, 15],
                notUpdating: true,
                colorScale: ["b0cce1", "3980b5", "0b62a4"],
                value: {
                    key: "amount_w",
                    metric: " downloads"
                },
                extraValues: [{
                    key: "amount_m",
                    metric: " downloads"
                }],
                mapOptions: {
                    attributionControl: false
                },
                zero: false
            }
        },
        computed:{
            pyDepartmentsDataReal: function(){
                return this.zero ? pyDepartmentsDataZero: pyDepartmentsData
            },
            inverseDepData: function(){
                return !this.zero ? pyDepartmentsDataZero: pyDepartmentsData
            },
            worldGeojson: function(){
                return worldGeojson
            },
            mapDataComputed: ()=>{
                return this.mapData
            },
            haveData(){
                return this.mapData.length > 0
            }

        },
        methods:{  
            refresh(){
                this.resource.getDownloadsByCountry()
                .then(
                response => { 
                    console.log(response)
                    return response.json()
                }, 
                error =>{
                    console.log(error)
                }
                ).then(data => { 
                const resultArray= [];

                console.log(data)
                for (let key in data){
                    resultArray.push(data[key])
                }
                this.downloads = resultArray
                console.log(resultArray)
                this.prepareData()
                })
            },
            prepareData(){
                var array =[]
                for (let key in this.downloads){
                    array.push({ label: this.downloads[key].CountryDetails.alpha3, value: this.downloads[key].Count })
                }
                this.mapData = array

            }
        },
        created(){ //on component created
            const customActions = {
                getDownloadsByCountry: {method: 'GET', url:'downloadsByCountryDetail'}
            }
            this.resource = this.$resource('downloadsByCountryDetail', {}, customActions);

            this.refresh();
            


            eventBus.$on('refreshMap',(data)=>{ //register on event
                if(this.zero === false){
                    eventBus.callRefreshMapData(this.worldGeojson,this.pyDepartmentsDataReal )
                    console.log('zero',this.pyDepartmentsDataReal)
                }else{
                    eventBus.callRefreshMapData(this.worldGeojson,this.pyDepartmentsData )
                    console.log('normal')
                }
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