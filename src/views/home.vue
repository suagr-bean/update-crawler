<template>
  
  <div class="item">
    <li @click="detail(data)"  v-for="(data,index) in list" :key="index" class="card">   
      <p class="name">{{ data.name }}</p></br>
      <p class="last"> {{ data.last }}</p>
      
    </li>
  </div>
</template>
<script setup>
import {ref,onMounted,defineEmits} from'vue'
import axios from'axios'
const list =ref([])
const GetShow=async()=>{
try{
  const response=await axios.get('/api/home')
  list.value=response.data
   console.log(response.data)
}catch(error){
  console.log("error",error)
}

}
onMounted(()=>{
  GetShow()
})
const emit=defineEmits(['detail'])
const detail=(data)=>{
    emit('detail',{url:data.url})
}

</script>
<style scoped>

.item{
  padding:10px;
  display:flex;
  flex-direction:column;
  gap:25px;
  width:95%;
}
.name{
  text-align:center;
  font-size:large;
  font-weight:800;
}
.last{
  font-size:medium;
}
.card{
  word-break:break-all;
  width:100%;
  height:auto;
  border:2px solid black;
  padding:6px;
  border-radius:15px;
}
li{
  list-style:none;
  padding:0;
  margin:0;
}
</style>
