<template>
    <div class="container">       
<!-- 
     idKey="alpha3"  //alpha3 - code retrived from server
-->
        <div class="row">
            <ChoroplethMap 
                class="col-xs-12"
                :data="mapDataComputed" 
                titleKey="department_name" 
                idKey="code"  
                :value="value" 
                :extraValues="extraValues" 
                geojsonIdKey="ISO_A3" 
                :geojson="worldGeojson" 
                :center="center" 
                :colorScale="colorScale" 
                mapStyle="height: 400px;" 
                :zoom="3" 
                :mapOptions="mapOptions">
                <!--template scope="props"-->
                    <!--InfoControl :item="props.currentItem" :unit="props.unit" title="Country" placeholder="Hover over a country"></InfoControl-->
                <!--/template-->
            </ChoroplethMap>
        </div>
        <br>
        <div class="row">
            <div class="col-xs-4">
                <ul class="list-group">
                    <li class="list-group-item">Country: {{selectedCountry.name}}</li>
                    <li class="list-group-item">Stat: {{selectedCountry.value}}</li>
                </ul>
            </div>
        </div>  
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
                    key: "value",
                    metric: " downloads"
                },
                extraValues: [{
                    key: "amount_m",
                    metric: " downloads"
                }],
                mapOptions: {
                    attributionControl: false
                },
                mapData: [],
                selectedCountry: {
                    name: "",
                    value: -1
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
                console.log('get data, dowloads: ',resultArray)
                this.prepareData()
                })
            },
            prepareData(){
                console.log('prepareData', this.downloads)
                var array =[]
                for (let key in this.downloads){
                    array.push({ code: this.downloads[key].CountryDetails.alpha3, value: this.downloads[key].Count })
                }
                
                this.mapData = array
                console.log('prepared array',this.mapData)
                eventBus.callRefreshMapData(this.worldGeojson,this.mapData )
            }
        },
        created(){ //on component created
            //custom actions for vue-resource
            const customActions = {
                getDownloadsByCountry: {method: 'GET', url:'downloadsByCountryDetail'}
            }
            this.resource = this.$resource('downloadsByCountryDetail', {}, customActions);
            
            eventBus.$on('sendCountryDataToWidget',(name, value)=>{ //register on event sendCountryDataToWidget

                this.selectedCountry.name = name;
                this.selectedCountry.value = value;
                console.log('name', name)
         
            })

            eventBus.$on('refreshMap',(data)=>{ //register on event sendCountryDataToWidget

                this.refresh()

                console.log('zero',this.mapData)
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