<template>
    <div>
        <h3>DonutChart</h3>
        <hr>
        <donut-chart 
            id="donut" 
            :data="donutData" 
            v-if="haveData"
            resize="true">
        </donut-chart>
        <p v-else>No Data</p>
        <button class="btn btn-primary" @click="refresh">Refresh</button>
    </div>
</template>
<script>
    import {eventBus} from './../../main'
    import { DonutChart } from 'vue-morris'

    export default{
        data(){ 
            return{
                donutData: [],
                downloads: [],
            }
        },
        components:{
            'donut-chart': DonutChart
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
                    array.push({ label: this.downloads[key].Country, value: this.downloads[key].Count })
                }
                this.donutData = array

            }
        },
        computed:{
            haveData(){
                return this.donutData.length > 0
            }
        },
        //on component created
        created(){ 
            //custom actions for vue-resource
            const customActions = {
                getDownloadsByCountry: {method: 'GET', url:'countDownloadsByCountry'}
            }
            this.resource = this.$resource('countDownloadsByCountry', {}, customActions);

            this.refresh();

            //register on event
            eventBus.$on('refreshChart',(data)=>{ 
                this.chartData = data
            })
        }
    }
</script>
<style></style>