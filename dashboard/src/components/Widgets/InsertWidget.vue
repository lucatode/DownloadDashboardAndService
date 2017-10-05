<template>
    <div>
        <h3>Input form</h3>
        <hr>
        <form class="form-horizontal">
            <div class="form-group">
                <label class="col-xs-3 control-label" for="appIdSelect">AppId</label>
                <div class="col-xs-8">
                    <!-- <input type="text" class="form-control" id="appIdSelect" v-model="download.appId"> -->
                    <select class="form-control" id="appIdSelect" v-model="download.appId">
                        <option selected value="iOS_App#1">iOS_App#1</option>
                        <option value="iOS_App#2">iOS_App#2</option>
                        <option value="Droid_App#1">Droid_App#1</option>
                        <option value="Droid_App#2">Droid_App#2</option>
                    </select>
                </div>
                <div class="col-xs-1">
                    <p v-bind:class="appIdValidated"></p>
                </div>
            </div>
            <div class="form-group">
                <label class="col-xs-3 control-label" for="longitudeInput">Longitude</label>
                <div class="col-xs-8">
                    <input type="text"  class="form-control" id="longitudeInput" placeholder="[-180 <-> 180]" v-model="download.longitude">
                </div>
                <div class="col-xs-1">
                    <p v-bind:class="longitudeValidated"></p>
                </div>
            </div>
            <div class="form-group">
                <label class="col-sm-3 control-label" for="latitudeInput">Latitude</label>
                <div class="col-sm-8">
                    <input type="text" class="form-control" id="latitudeInput" placeholder="[-90 <-> 90]" v-model="download.latitude">
                </div>
                <div class="col-xs-1">
                    <p v-bind:class="latitudeValidated"></p>
                </div>
            </div>
            <div class="form-group">
                <label class="col-sm-3 control-label" for="dateTimeInput">Datetime</label>
                <div class="col-sm-8">
                    <select class="form-control" id="dateTimeInput" v-model="download.downloadedAt">
                        <option value='10:00'>10:00</option>
                        <option value='17:00'>17:00</option>
                        <option value='22:00'>22:00</option>
                        <option value='03:00'>03:00</option>
                    </select>
                </div>
                <div class="col-xs-1">
                    <span class="glyphicon glyphicon-question-sign"></span>
                </div>
            </div>
            <button class="btn btn-primary" :enabled="formValid" @click="submit">Submit</button>
        </form>
        <hr>
        <!-- <p>Insert status: {{'ok'}}</p> -->
        <!-- <button @click='editAge()'>Edit Age</button> -->
    </div>
</template>
<script>
    import {eventBus} from './../../main'
    export default{
        data(){ 
            return{
                'chartData': {},
                download:{
                    appId: '',
                    longitude: '',
                    latitude: '',
                    downloadedAt: '',
                    country: ''
                },
                downloads: [],
                resource: {}
            }
        },
        methods:{
            submit(){
                this.resource.postDownload(this.download).then(response =>{}, error=>{
                    response => { 
                    //console.logresponse)
                    return response.json()
                    }, 
                    error =>{
                        //console.logerror)
                    }
                }) //TODO add promise
            },
            appIdValid(){
                return (this.download !== null || this.download.localeCompare('') !== 0 || this.download.localeCompare("")!==0)
            },
            longitudeValid(){
                return (!isNaN(this.download.longitude)
                &&(this.download.longitude <= 180 && this.download.longitude>= -180)
                && this.download.longitude !== '')
            },
            latitudeValid(){
                return (!isNaN(this.download.longitude)
                &&(this.download.latitude <= 90 && this.download.latitude>= -90)
                && this.download.longitude !== '')
            },

            logSubmit(){
                //console.log"Logging post"+this.download)
            }
        },
        computed:{
            
            
            appIdValidated(){
                if(this.appIdValid()){
                    return "glyphicon glyphicon-ok-circle green"
                }else{
                    return "glyphicon glyphicon-ban-circle red"
                }
            },

            longitudeValidated(){
                if(this.longitudeValid()){
                    return "glyphicon glyphicon-ok-circle green"
                }else{
                    return "glyphicon glyphicon-ban-circle red"
                }
            },

            latitudeValidated(){
                if(this.latitudeValid()){
                    return "glyphicon glyphicon-ok-circle green"
                }else{
                    return "glyphicon glyphicon-ban-circle red"
                }
            },
            formValid(){
                return this.appIdValid() && this.longitudeValid() && this.latitudeValid()
            }

        },
        created(){ //on component created

            //custom action of vue-resource
            const customActions = {
                postDownload: {method:'POST', url:'downloads'},
                getDownloads: {method: 'GET', url:'downloads'}
            }
            this.resource = this.$resource('downloads', {}, customActions);
            
            //register on event
            eventBus.$on('refreshChart',(data)=>{ 
                this.chartData = data
            })
        }
    }
</script>
<style></style>