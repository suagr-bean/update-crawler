<template>   
    <div class="all">
        <div class="progress">
        <span>{{FormatTime(now)}}/{{ FormatTime (end)}}</span>
        </div>
       <div class="play">
        <audio ref="audiourl" @play="isPlay = true" @pause="isPlay = false" @loadedmetadata="getload" @timeupdate="ontimeupdate"></audio>
        <h3 class="title">{{ info.title }}</h3> 
        <button @click="togglePlay" class="control">{{ isPlay ? '⏸️' : '▶️' }}</button> 
</div>
    </div>
</template>
<script setup>
import { ref, defineProps, watch } from 'vue'
import {FormatTime} from'../utils/formatTime'
const audiourl = ref(null) 
const isPlay = ref(false)  
const now =ref(0)
const end =ref(0)
const props = defineProps({
  info: {
    type: Object,
    default: () => ({})
  }
})
const getload=()=>{
  end.value=audiourl.value.duration;
}
const ontimeupdate=()=>{
  now.value=audiourl.value.currentTime;
}
watch(()=>props.info.auto_link,(newlink)=>{
   if (newlink){
    audiourl.value.src=newlink
    audiourl.value.load();
    audiourl.value.play();
    isPlay.value=true;
   }
})
const togglePlay=()=>{
  if (!audiourl.value) return;
  if (audiourl.value.paused){
    audiourl.value.play();
  }else {
    audiourl.value.pause();
  }

}

</script>
<style scoped>
.all{
  display:flex;
  flex-direction:column;
  width:100%;
  height:100%;
}
.progress{
  margin-top:0px;
  height:3px;
  padding:0;
  width:100%;
}
progress{
  width:100%;
   height:3px;
  border:none;
  color:white;
  background-color:white;
}
 .play{
   flex:1;
   margin-top:6px;
    display:flex;
    flex-direction:row;
    width:100%;
   align-items:center;
 }
 .title{
   margin-left:10px;
    flex:0.8;
    font-size:13px;
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
