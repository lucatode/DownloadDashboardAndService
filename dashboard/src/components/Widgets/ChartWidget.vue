<template>
    <div>
        <p>
            This will be chart widget - {{this.chartData.name}}
        </p>
        <ul class="list-group">
          <li class="list-group-item" v-for="d in downloads"> Country: {{d.Country}} - Count: {{d.Count}} </li>
        </ul>
        <button class="btn btn-primary" @click="refresh">Refresh</button>
    </div>
</template>
<script>
    import {eventBus} from './../../main'
    export default{
        data(){ 
            return{
                'chartData': {},
                downloads: [],
            }
        },
        methods:{
            refresh(){
                this.resource.getDownloadsByCountry()
                .then(
                response => { 
                    ////console.log(response)
                    return response.json()
                }, 
                error =>{
                    ////console.log(error)
                }
                ).then(data => { 
                const resultArray= [];

                ////console.log(data)
                for (let key in data){
                    resultArray.push(data[key])
                }
                this.downloads = resultArray
                ////console.log(resultArray)

                })
            }
        },
        //on component created
        created(){ 
            //custom actions for vue-resource
            const customActions = {
                getDownloadsByCountry: {method: 'GET', url:'countDownloadsByCountry'}
            }
            this.resource = this.$resource('countDownloadsByCountry', {}, customActions);

            //register on event
            eventBus.$on('refreshChart',(data)=>{ 
                this.chartData = data
            })
        }
    }
</script>
<style></style>