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
                    <p v-bind:class="appIdValid"></p>
                </div>
            </div>
            <div class="form-group">
                <label class="col-xs-3 control-label" for="longitudeInput">Longitude</label>
                <div class="col-xs-8">
                    <input type="text"  class="form-control" id="longitudeInput" v-model="download.longitude">
                </div>
                <div class="col-xs-1">
                    <p v-bind:class="longitudeValid"></p>
                </div>
            </div>
            <div class="form-group">
                <label class="col-sm-3 control-label" for="latitudeInput">Longitude</label>
                <div class="col-sm-8">
                    <input type="text" class="form-control" id="latitudeInput" v-model="download.latitude">
                </div>
                <div class="col-xs-1">
                    <p v-bind:class="latitudeValid"></p>
                </div>
            </div>
            <div class="form-group">
                <label class="col-sm-3 control-label" for="dateTimeInput">Datetime</label>
                <div class="col-sm-8">
                    <select class="form-control" id="dateTimeInput" v-model="download.downloadedAt">
                        <option v-for="n in 24" value="n">{{n}}:00</option>
                    </select>
                </div>
                <div class="col-xs-1">
                    <span class="glyphicon glyphicon-question-sign"></span>
                </div>
            </div>
            <button class="btn btn-primary" @click="submit">Submit</button>
        </form>
        <hr>
        <p>Insert status: {{'ok'}}</p>
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
                this.resource.postDownload(this.download) //TODO add promise
            },
            logSubmit(){
                console.log("Logging post"+this.download)
            }
        },
        computed:{
            appIdValid(){
                if(this.download !== null || this.download.localeCompare('') !== 0 || this.download.localeCompare("")!==0){
                    return "glyphicon glyphicon-ok-circle green"
                }else{
                    return "glyphicon glyphicon-ban-circle red"
                }
            },
            longitudeValid(){
                if(!isNaN(this.download.longitude)
                &&(this.download.longitude < 180 && this.download.longitude> -180)
                && this.download.longitude !== ''){
                    return "glyphicon glyphicon-ok-circle green"
                }else{
                    return "glyphicon glyphicon-ban-circle red"
                }
            },
            latitudeValid(){
                if(!isNaN(this.download.longitude)
                &&(this.download.latitude < 90 && this.download.latitude> -90)
                && this.download.longitude !== ''){
                    return "glyphicon glyphicon-ok-circle green"
                }else{
                    return "glyphicon glyphicon-ban-circle red"
                }
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