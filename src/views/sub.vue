<template>
  <h3>订阅</h3>
  <input  v-model="data.url"placeholder="请输入URL">
  <button @click="Submit">添加</button>
  
</template>
<script setup>
import {reactive}from 'vue'
import axios from 'axios'
const data=reactive({
  url:'',
})
const Submit=async()=>{
try{
 const response=await axios.post('/api/add',data)
  alert("发送成功")
  
}catch(error){

}
}
</script>