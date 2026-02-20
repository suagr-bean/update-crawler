<template>
  <div class="play-button-container">
    <p class="title">{{ info.title }}</p>
    <audio autoplay @timeupdate="updateTime" @loadedmetadata="setDuration" @pause="isPlaying=true" @play="isPlaying=false"ref="audioRef"  :src="info.auto_link"  style="width: 100%;"></audio>
    <input type="range" :value="now" :max="end">
    <span>{{ now }}/{{end }}</span>
    <div class="control">
    <button class="play"@click="togglePlay"> {{ isPlaying?'暂停' :'▶'}}</button>
    <button class="stop"@click="stop">定时</button>
   </div>
  </div>
</template>

<script setup>
import {ref}from'vue'
const props = defineProps({
  info: {
    type: Object,
    default: () => ({})
  }
})
const isPlaying=ref(false)
const audioRef=ref(null)
const togglePlay=()=>{
  if (!audioRef.value){
    return;
  }
  if (isPlaying.value){
    audioRef.value.pause();
  }else{
    audioRef.value.play();
  }
  isPlaying.value=!isPlaying.value;
};
//总时长
const end=ref(0)
//当前时间
const now=ref(0)
//更新当前时间
const updateTime=(event)=>{
  now.value=event.target.currentTime
};
//总时长
const setDuration=(event)=>{
  end.value=event.target.duration;
};
//定时
let timer=null
const remaining=ref(0)
const stop=()=>{
   remaining.value=10*60
   timer=setInterval(()=>{
    remaining.value--;
    if (remaining.value<=0){
     audioRef.value.pause();
     clearInterval(timer);
    }
   },1000)
}
</script>

<style scoped>
.play-button-container {
  display: flex;
  flex-direction: column; /* Arrange title and player vertically */
  align-items: center; /* Center them horizontally */
  background-color: #f0f0f0;
  padding: 10px 5px; /* Add some padding */
  border-top: 1px 
}
.control{
  display:flex;
  flex-direction:row;
  gap:50px;
  margin-left:0;
}
.play{
  width:60px;
  height:60px;
}
.stop{
  width:60px;
}

.title {
  font-weight: 500;
  margin-bottom: 8px; /* Add space between title and player */
  text-align: center;
  width: 95%;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis; /* Use '...' for very long titles */
}
</style>
