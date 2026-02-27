<template>
  
  <div class="item">
    <p v-if="list.length>0">{{ list[0].count}}条</p>
    
    <li @click="detail(data)"  v-for="(data,index) in list" :key="index"class="card">  
     <div class="context_type">
     <span v-if="data.context_type=='podcast'">音频</span> 
     <span v-if="data.context_type=='article'">文章</span>
     </div>
      <p class="name">{{ data.name }}</p>
      <p class="last"> {{ data.last }}</p>
    </li>
   <div class="space"><button>加载更多</button></div>
  </div>
</template>
<script setup>
import {ref,onMounted,defineEmits} from'vue'
import axios from'axios'
const list =ref([])
const GetShow=async()=>{
try{
  const response=await axios.get('/api/home')
  list.value=response.data.info
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
  height:auto;
  overflow-y:auto;
}
.name{
  text-align:center;
  font-size:large;
  font-weight:800;
}
.last{
  font-size:medium;
  /* Prevent text from overflowing */
  white-space: normal;
  word-wrap: break-word;
}
.card{
  /* Make card a flex container */
 /* Stack items vertically */
  word-break:break-all;
  width:95%;
  border:2px solid black;
  padding:6px;
  border-radius:15px;
}
li{
  list-style:none;
  padding:0;
  margin:0;
}
.space{
height:100px;
}
.context_type{
  background-color:white;
  text-align:center;
}
</style>
