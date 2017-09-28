<template>
    <div class="container">
        <div class="row">
            <div class="col-xs-12">
                <h1>App Download Dashboard</h1>
            </div>
        </div>
        <hr>
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
            <div class="col-xs-12">
                <bar-chart-widget></bar-chart-widget>
            </div>
        </div>
        <div class="row">
            <div class="col-xs-6">
                <insert-widget></insert-widget>
            </div>
            <div class="col-xs-6">
                <donut-chart-widget></donut-chart-widget>
            </div>
        </div>
        <hr>
        <div class="row">
            <div class="col-xs-12">
                <table-widget></table-widget>
            </div>
        </div>
    </div>
</template>
<script>
    import MapWidget from './Widgets/MapWidget.vue'
    import ChartWidget from './Widgets/ChartWidget.vue'
    import InsertWidget from './Widgets/InsertWidget.vue'
    import DonutChartWidget from './Widgets/DonutChartWidget.vue'
    import BarChartWidget from './Widgets/BarChartWidget.vue'
    import TableWidget from './Widgets/TableWidget.vue'
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
            'donut-chart-widget': DonutChartWidget,
            'bar-chart-widget': BarChartWidget,
            'table-widget': TableWidget
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