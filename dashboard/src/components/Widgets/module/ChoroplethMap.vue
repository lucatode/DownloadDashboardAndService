<template>
    <v-map id="map" :zoom="zoom" :center="center" :style="mapStyle" :options="mapOptions">
        <v-geojson-layer :geojson="geojson" :options="getOptions"></v-geojson-layer>
        <!--slot :currentItem="currentItem" :unit="value.metric" :min="min" :max="max"></slot-->
    </v-map>
</template>

<script>
import {eventBus} from './../../../main'
import InfoControl from './InfoControl.vue'
import ReferenceChart from './ReferenceChart.vue'
import Vue2Leaflet from 'vue2-leaflet'
import GeoJSON from './GeoJSON.vue'
import { getMin, getMax, normalizeValue, getColor } from './util'

function mouseover({ target }) {
    target.setStyle({
        weight: 5,
        color: "#0B62A4",
        dashArray: ""
    })

    if (!L.Browser.ie && !L.Browser.opera) {
        target.bringToFront()
    }

    let geojsonItem = target.feature


    if(this.getData === undefined){
        this.currentItem = { name: "", value: 0 } 
        return;
    }

    //let item = this.getData.find(x => x[this.idKey] === Number(geojsonItem[this.geojsonIdKey]))


    let itemGeoJSONID = geojsonItem[this.geojsonIdKey]
    if(this.getData === undefined){
        console.log(this.getData)
    }

    let item = this.getData.find(x => x[this.idKey] === itemGeoJSONID)
    let itemName = geojsonItem.properties["name"];
    
    if(!itemName){
        this.currentItem = { name: "", value: 0 }
        console.log('no item:'+item)
    }else if (!item ) {
    this.currentItem = { name: itemName, value: 0 }
    eventBus.sendCountryDataToWidget(this.currentItem.name, this.currentItem.value);
    console.log('itemName: '+itemName+', no item:'+item)
    }else{
        let tempItem = { name: itemName, value: item.value }

        if (this.extraValues) {
            let tempValues = [];
            for (let x of this.extraValues) {
                tempValues.push({
                    value: item[x.key],
                    metric: x.metric
                })
            }
            //tempItem = { ...tempItem, extraValues: tempValues }
            tempItem = {name: tempItem.name, value:tempItem.value, extraValues: tempValues}
        }
            this.currentItem = tempItem
            eventBus.sendCountryDataToWidget(this.currentItem.name, this.currentItem.value);
    }


}




function mouseout({ target }) {
    target.setStyle({
        weight: 2,
        color: "#FFF",
        dashArray: ""
    })
    
}

export default {
    props: [
        "geojson",
        "data",
        "center",
        "colorScale",
        "titleKey",
        "idKey",
        "value",
        "extraValues",
        "geojsonIdKey",
        "mapStyle",
        "zoom",
        "mapOptions"
    ],
    data() {
        return {
            updatedData:undefined,
            updatedMap: undefined,
            currentItem: { name: "", value: 0 }
        }
    },
    computed: {
        min() {
            return getMin(this.data, this.value.key)
        },
        max() {
            return getMax(this.data, this.value.key)
        },
        getGeojsonOptions(){
            return this.geojsonOptions
        },
        getGeoJson(){
            return this.updatedMap === undefined ? this.geojson : this.updatedMap
        },
        updatedMapValue(){
            return this.updatedMap
        },
        getData(){
            if(this.updatedData===undefined){
                return this.data
            }else{
                return this.updatedData
            }
        },
        getOptions(){
            return {
                style: feature => {
                    let itemGeoJSONID = Number(feature[this.geojsonIdKey])
                    let color = "NONE"
                    if(this.getData === undefined){
                        console.log('no data');
                        return {                            
                            color: "white",
                            weight:2
                        }
                    }

                    let item = this.getData.find(x => x[this.idKey] === itemGeoJSONID)
                    if (!item) {
                        console.log('item not found');
                        return {
                            color: "white",
                            weight: 2
                        }
                    }
                    
                    //Get quantity
                    let valueParam = item[this.value.key]
                    if (!Number(valueParam)) {
                        console.log('quantity not found');
                        return {
                            color: "white",
                            weight: 2
                        }
                    }
                    //import funcs
                    const { min, max } = this
                    
                    console.log('value param:', valueParam);
                    //build feature object
                    return {
                        weight: 2,
                        opacity: 10,
                        color: "white",
                        dashArray: "3",
                        fillOpacity: 0.7,
                        fillColor: "black"
                    }

                    //getColor(valueParam, this.colorScale, min, max)
                },
                onEachFeature: (feature, layer) => {
                    layer.on({
                        mouseover: mouseover.bind(this),
                        mouseout: mouseout.bind(this)
                    })
                },
            }
        },
        getMap(){
            if(updatedMap===undefined){
                return this.geojson;
            }
            return this.updatedMap;

        }

    },
    components: {
        "v-map": Vue2Leaflet.Map,
        "v-geojson-layer": GeoJSON,
        'v-tilelayer': Vue2Leaflet.TileLayer,
        InfoControl,
        ReferenceChart
    },
    methods:{
        
    },
    mounted() {

    },
    created(){ //on component created
            eventBus.$on('refreshMapData',(worldMap, updatedData) =>{
                this.updatedMap = worldMap
                console.log('updated data', updatedData)
                this.updatedData = updatedData
            })
            eventBus.refresh();
        }
}
</script>