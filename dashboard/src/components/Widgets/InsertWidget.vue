<template>
    <div>
        <p>
            <div class="form-group">
                <label>AppId</label>
                <input type="text" class="form-control" v-model="download.appId">
            </div>
            <div class="form-group">
                <label>Longitude</label>
                <input type="text" class="form-control" v-model="download.longitude">
            </div>
            <div class="form-group">
                <label>Latitude</label>
                <input type="text" class="form-control" v-model="download.latitude">
            </div>
            <div class="form-group">
                <label>DownloadedAt</label>
                <input type="text" class="form-control" v-model="download.downloadedAt">
            </div>
            <button class="btn btn-primary" @click="submit">Submit</button>
        </p>
        <!-- <p v-show="correct">{{}}</p> -->
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