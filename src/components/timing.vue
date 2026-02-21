<template>
    <div class="show">
    <button @click="choose(time)" v-for="time in timeOption" :key="time">{{ time }}</button>
    </div>
</template>
<style scoped>
.show{
    display:flex;
    flex-direction:column;
    width:95%;
    padding:10px;
    height:auto; 

}
button{
   flex:0 0 40px; 
}
</style>
<script setup>
import {defineEmits} from 'vue'
const timeOption=[10,20,30,60]
//给父类传时间
const emit =defineEmits(['stopTime'])
const choose =(time)=>{
emit('stopTime',time)
}
</script>