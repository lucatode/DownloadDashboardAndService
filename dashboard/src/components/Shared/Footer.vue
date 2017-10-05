<template>
    <div>
        <button class="btn btn-primary" @click='refresh'>Clear Map Data</button>
        <footer>
            
        </footer>
        
    </div>
</template>

<script>
import {eventBus} from './../../main'

export default{
    methods:{
        refresh(){
            eventBus.refresh();
            //console.log('footer');
            //eventBus.refreshView()
        }
    }
}


</script>

<style></style>