<template>
  <div class="all">
    <p v-if="list.length>0">{{list[0].count}}条</p>
    <van-list v-model:loading="loading" :finished="finished" @load="onLoad" finished-test="到底了">
    <li class="card" v-for="(item,index) in list":key="index">
        <p class="title">{{ item.title }}</p>
        <p class="time">{{ formatTime(item.published_time) }}</p>
        <div class="control">
        <a :href="item.link"rel="noreferrer">原文</a>
        <button @click="Deal(item)"class="play">▶</button>
        </div>
    </li>
    </van-list>
  </div>
</template>
<script setup>
import {ref,defineEmits, defineProps}from 'vue'
import axios from 'axios'
import {List as VanList} from'vant'
const list=ref([])
const loading=ref(false )
const finished=ref(false )
const page =ref(1)
const props =defineProps(['apiurl'])
const emit=defineEmits(['control'])
const onLoad =async()=>{
  try{
   const response =await axios.get("/api/showdetail",{params:{
    url:props.apiurl,
    start:page.value,
    size:50,
   }
  })
  const newdata=response.data.data
  list.value.push(...newdata)
  loading.value=false
   page.value++
  if (!newdata||newdata.length<50){
    finished.value=true
 
  }
}catch(error){
    loading.value=true
    finished.value=true
  }
}

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
  display:flex;
  flex-direction:column;
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
   width:95%;
  border:2px solid black;
  border-radius: 15px;
  margin-bottom:10px;
}
</style>