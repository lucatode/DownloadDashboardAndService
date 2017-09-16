<template>
    <div class="container">
        <h2>This will be dashboard header</h2>
        <div class="row">
            <div class="col-xs-12">
                <map-widget></map-widget>
            </div>
        </div>
        <hr>
        <!-- <div class="row">
            <div class="col-xs-6">
                <insert-widget></insert-widget>
            </div>
            <div class="col-xs-6">
                <chart-widget></chart-widget>
            </div>
        </div> -->
        <div class="row">
            <div class="col-xs-4">
                <insert-widget></insert-widget>
            </div>
            <div class="col-xs-4">
                <donut-chart-widget></donut-chart-widget>
            </div>
        </div>
        <hr>
    </div>
</template>
<script>
    import MapWidget from './Widgets/MapWidget.vue'
    import ChartWidget from './Widgets/ChartWidget.vue'
    import InsertWidget from './Widgets/InsertWidget.vue'
    import DonutChartWidget from './Widgets/DonutChartWidget.vue'
    import {eventBus} from './../main'

    export default {
        data(){ 
            return{
            }
        },
        components:{
            'map-widget': MapWidget,
            'chart-widget': ChartWidget,
            'insert-widget': InsertWidget,
            'donut-chart-widget': DonutChartWidget
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