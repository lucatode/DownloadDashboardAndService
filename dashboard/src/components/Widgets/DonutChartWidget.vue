<template>
    <div>
        <h3>Donut Chart</h3>
        <hr>
        <donut-chart 
            id="donut" 
            :data="donutData" 
            v-if="haveData"
            resize="true">
        </donut-chart>
        <p v-else>No Data</p>
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
                    //console.logresponse)
                    return response.json()
                }, 
                error =>{
                    //console.logerror)
                }
                ).then(data => { 
                const resultArray= [];

                //console.logdata)
                for (let key in data){
                    resultArray.push(data[key])
                }
                this.downloads = resultArray
                //console.logresultArray)
                this.prepareData()
                })
            }, 
            prepareData(){
                var array =[]
                for (let key in this.downloads){
                    array.push({ label: this.downloads[key].CountryDetails.alpha3, value: this.downloads[key].Count })
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
                getDownloadsByCountry: {method: 'GET', url:'downloadsByCountryDetail'}
            }
            this.resource = this.$resource('downloadsByCountryDetail', {}, customActions);

            this.refresh();

            //register on event
            eventBus.$on('refreshChart',(data)=>{ 
                this.chartData = data
            })
        }
    }
</script>
<style></style>