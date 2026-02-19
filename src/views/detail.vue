<template>

  
  <div class="all">
    <li class="card" v-for="(item,index) in list":key="index">
        <p class="title">{{ item.title }}</p>
        <div class="control">
        <a :href="item.link">原文</a>
        <button @click="Deal(item)"class="play">▶</button>
        </div>
    </li>
  </div>
</template>
<script setup>
import {ref,onMounted,defineEmits}from 'vue'
import axios from 'axios'
const list=ref([])

const props=defineProps(['apiurl'])
const GetDetail=async()=>{
    try{
     const response =await axios.get("/api/showdetail",{params:{
        url:props.apiurl,
     }})
     
     list.value=response.data.data
     
    }catch{

    }

}
onMounted (()=>{
    GetDetail()
})
const emit =defineEmits (['control'])
const Deal=(item)=>{
  emit('control',item);
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
  
}
.card{
   padding:10px;
   height:auto;
   width:100%;
  border:2px solid black;
  border-radius: 15px;
}
</style>