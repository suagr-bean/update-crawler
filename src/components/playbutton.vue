<template>
  <timing @stopTime="dealStop" class="show" v-if="showTimer"></timing>
  <div v-if="info.auto_link" class="play-button-container">
    <p class="title">{{ info.title }}</p>
    <audio
      autoplay
      ref="audioRef"
      :src="info.auto_link"
      style="width: 100%;"
      @timeupdate="updateTime"
      @loadedmetadata="setDuration"
      @play="isPlaying = true"
      @pause="isPlaying = false"
    ></audio>
    <input type="range" min="0" :max="end" v-model="now" step="0.01" @input="onSeek"></input>
    <div class="control">
    <span>{{ FormatTime(now) }} / {{ FormatTime(end ) }}</span>
      <button class="play" @click="togglePlay">{{ isPlaying ? '暂停' : '▶' }}</button>
      <button class="stop" @click="stop" >定时</button>
      
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { FormatTime } from '@/utils/formatTime'
import timing from './timing.vue'
const props = defineProps({
  info: {
    type: Object,
    default: () => ({})
  }
})

const audioRef = ref(null)
const isPlaying = ref(false)
const now = ref(0)
const end = ref(0)
const remaining = ref(0)
//控制定时
const showTimer=ref(false)
//定时操作
const sleepTime=ref(0)
const dealStop=(val)=>{
    showTimer.value=false
    sleepTime.value=val*60*1000
    setTimeout(()=>{
      audioRef.value.pause()
    },sleepTime.value)
}
const togglePlay = () => {
  if (!audioRef.value) return;
  if (isPlaying.value) {
    audioRef.value.pause()
  } else {
    audioRef.value.play()
  }
}

const updateTime = (e) => {
  now.value = e.target.currentTime
}

const setDuration = (e) => {
  end.value = e.target.duration
}
const onSeek =() =>{
   if (audioRef.value){
    audioRef.value.currentTime=now.value
   }
}

const stop= () => {
  showTimer.value=!showTimer.value
}
</script>

<style scoped>
.play-button-container {
  height: 130px;
  width: 100%;
  background-color:white;
}

.title {
  text-align: center;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
.control {
  display: flex;
  justify-content: space-around;
  height: 50px;
  align-items: center;
}
input{
  margin-left:20px;
  width:90%;
  height:20%;
}

.show {
  height:150px;
}
</style>
