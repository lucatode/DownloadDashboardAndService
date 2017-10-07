<template></template>

<script>
import {eventBus} from './../../../main'
import { getMin, getMax, normalizeValue, getColor } from './util'

export default {
  props: ['geojson', 'options'],
  data(){
    return {
      geoJson: undefined,
      mapOptions: undefined,
      geoJSON: undefined
    }
  },
  mounted() {
    
    this.geoJson = this.geojson;
    this.mapOptions = this.options;
    this.geoJSON = L.geoJSON(this.geoJson, this.mapOptions);

    this.$geoJSON = this.geoJSON 

    if (this.$parent._isMounted) {
      this.deferredMountedTo(this.$parent.$mapObject);
    }
  },
  beforeUpdate(){
    //console.log('before update')
    this.$geoJSON = L.geoJSON(this.geojson, this.options);

    if (this.$parent._isMounted) {
      this.deferredMountedTo(this.$parent.$mapObject);
    }
  },
  beforeDestroy() {
    this.setVisible(false);
  },
  watch: {
    geojson: {
      handler(newVal) {
        this.$geoJSON.clearLayers()
        this.addGeoJSONData(newVal);
      },
      deep: true,
    },
  },
  methods: {
    deferredMountedTo(parent) {
      this.parent = parent;
      this.$geoJSON.addTo(parent);
      for (var i = 0; i < this.$children.length; i++) {
        this.$children[i].deferredMountedTo(parent);
      }
    },
    addGeoJSONData(geojsonData) {
      this.$geoJSON.addData(geojsonData);
    },
    getGeoJSONData() {
      return this.$geoJSON.toGeoJSON();
    },
    getBounds() {
      return this.$geoJSON.getBounds();
    },
    setVisible(newVal, oldVal) {
      if (newVal === oldVal) return;
      if (newVal) {
        this.$geoJSON.addTo(this.parent);
      } else {
        this.parent.removeLayer(this.$geoJSON);
      }
    },
  },
  created(){ //on component created
    eventBus.$on('refreshMapData', (worldMap, updatedData, colorScale) =>{
      //console.log('updated Data', updatedData)
      this.mapOptions.style = 
        feature => {
                    let itemGeoJSONID = feature.properties["ISO_A3"]
                    let color = "NONE"
                    if(updatedData === undefined){
                                        return {                            
                            color: "white",
                            weight:2
                        }
                    }

                    let item = updatedData.find(x => x["code"] === itemGeoJSONID)
                    if (!item) {
                        console.log('item not found');
                        return {
                            color: "white",
                            weight: 2
                        }
                    }
                    
                    //Get quantity
                    let valueParam = item["value"]
                    if (!Number(valueParam)) {
                        console.log('quantity not found');
                        return {
                            color: "white",
                            weight: 2
                        }
                    }
                    //import funcs
                    const { min, max } = this
                    
                    var countryColor = getColor(valueParam, colorScale, min, max)
                    console.log('color:', countryColor);
                    //build feature object
                    return {
                        weight: 2,
                        opacity: 10,
                        color: "white",
                        dashArray: "3",
                        fillOpacity: 0.7,
                        fillColor: "black"
                    }

          
      }


        var newLayer = L.geoJSON(worldMap, this.mapOptions);
        this.parent.removeLayer(this.$geoJSON);
        this.$geoJSON = newLayer
        this.$geoJSON.addTo(this.parent);
    })
  }
};
</script>
