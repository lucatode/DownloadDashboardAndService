<template>
    <div>
        <h3>Bar Chart</h3>
        <hr>
        <bar-chart 
            id="bar" 
            :data="donutData"
            xkey="label"
            ykeys='["value"]'
            grid="true"
            v-if="haveData"
            resize="true">
        </bar-chart>
        <p v-else>No Data</p>
    </div>
</template>
<script>
    import {eventBus} from './../../main'
    import { BarChart } from 'vue-morris'

    export default{
        data(){ 
            return{
                donutData: [],
                downloads: [],
            }
        },
        components:{
            'bar-chart': BarChart
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