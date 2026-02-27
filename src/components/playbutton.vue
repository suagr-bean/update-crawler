
<template>
   
    
    <div class="play"> 
        <audio></audio>
        <h3 class="title">{{ info.title }}</h3>
           
    <button class="control">▶️</button>
    
</div>

</template>
<script setup>
import {ref,defineProps}from 'vue'
const audioRef=ref(null)
const isPlay=ref(false)
//接收URL
const props =defineProps( {
info:{
    type:Object,
    default:()=>({})
}
})

</script>
<style scoped>
.all{
    display:flex;
    flex-direction:column;
    border :2px solid black;
    border-radius:10px;
    
}
 .play{
    display:flex;
    flex-direction:row;
    width:100%;
   align-items:center;
 }
input{
    width:70%;
    margin:0 auto;
}
 .title{
    flex:0.9;
    
    font-size:13px;
    white-space:nowrap;
     overflow:hidden;
     text-overflow:ellipsis;
     line-height:1.3;
}

.control{
    margin:0;
}
.more{
    width:40px;
    height:20px;
    margin-right:5px;
}
</style>