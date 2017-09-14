<template>
    <div>
        <p>
            This will be chart widget - {{this.chartData.name}}
        </p>
        <!-- <button @click='editAge()'>Edit Age</button> -->
    </div>
</template>
<script>
    import {eventBus} from './../../main'
    export default{
        data(){ 
            return{
                'chartData': {},
            }
        },
        methods:{
            // editAge() {
            //    this.userAge = 30;
            //    //this.$emit('ageWasEdited', this.userAge )
            //    eventBus.changeAge(this.userAge) <--- comunicate between components
            // }
        },
        created(){ //on component created
            eventBus.$on('refreshChart',(data)=>{ //register on event
                this.chartData = data
            })
        }
    }
</script>
<style></style>