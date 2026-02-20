<template>
  <div class="all">
    <li class="card" v-for="(item,index) in list":key="index">
        <p class="title">{{ item.title }}</p>
        <p class="time">{{ formatTime(item.published_time) }}</p>
        <div class="control">
        <a :href="item.link">原文</a>
        <button @click="Deal(item)"class="play">▶</button>
        </div>
    </li>
    <div class="space">
      <button>加载更多</button>
    </div>
  </div>
</template>
<script setup>
import {ref,onMounted,defineEmits, defineProps}from 'vue'
import axios from 'axios'
const list=ref([])

const props=defineProps(['apiurl'])
const GetDetail=async()=>{
    try{
     const response =await axios.get("/api/showdetail",{params:{
        url:props.apiurl,
     }})
     
     list.value=response.data.data
     
    }catch(e){
      console.error(e)
    }

}
onMounted (()=>{
    GetDetail()
})
const emit =defineEmits (['control'])
const Deal=(item)=>{
  emit('control',item);
};

const formatTime = (unixTimestamp) => {
  if (!unixTimestamp) {
    return '';
  }
 
  const date = new Date(unixTimestamp * 1000);
  const year = date.getFullYear();
  const month = `0${date.getMonth() + 1}`.slice(-2);
  const day = `0${date.getDate()}`.slice(-2);
  const hours = `0${date.getHours()}`.slice(-2);
  const minutes = `0${date.getMinutes()}`.slice(-2);
  return `${year}-${month}-${day} ${hours}:${minutes}`;
};
</script>
<style scoped>
.play{
  border-radius:4px;
  width:100px;
  height:40px;
}
.control{
  display:flex;
  gap:150px;
}
.space{
  height:100px;
}
a{
   width:100px;
   height:30px;
  margin-left:20px;
}
.title{
  text-align:center;
  font-size:x-large;
  font-weight:800;
}
.time {
  text-align: center;
  color: #888;
  font-size: 0.9em;
  margin-bottom: 15px;
}
li{
  list-style:none;
  padding:0;
}
.all{
  padding:10px;
  display:flex;
  flex-direction:column;
  gap:50px;
  margin-left:0;
  width:90%;
  height:auto;
}
.card{
   padding:10px;
   height:auto;
   width:100%;
  border:2px solid black;
  border-radius: 15px;
}
</style>