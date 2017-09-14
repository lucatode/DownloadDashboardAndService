<template>
    <div>
        <h2>This will be dashboard header</h2>
        <map-widget></map-widget>
        <chart-widget></chart-widget>        
    </div>
</template>
<script>
    import MapWidget from './Widgets/MapWidget.vue'
    import ChartWidget from './Widgets/ChartWidget.vue'
    import {eventBus} from './../main'

    export default {
        data(){ 
            return{
            }
        },
        components:{
            'map-widget': MapWidget,
            'chart-widget': ChartWidget
        },
        methods:{
            // updateMapWidget(){
            //     this.name = "Ila"
            // },
            // resetFunction(){
            //     this.name = "Luca"
            // }
        }
    }
</script>
<style>

</style>