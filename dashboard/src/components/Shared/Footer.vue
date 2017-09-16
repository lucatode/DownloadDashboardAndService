<template>
    <div>
        <button @click='refresh'>Refresh Dashboard</button>
        <footer>
            Made with Love
        </footer>
        
    </div>
</template>

<script>
import {eventBus} from './../../main'

export default{
    methods:{
        refresh(){
            eventBus.refresh()
            eventBus.refreshView()
        }
    }
}


</script>

<style></style>