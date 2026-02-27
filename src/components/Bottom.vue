<template>
  <div class="bar">
    <button class="home" @click="$emit('show-home')">首页</button>
    <button @click="$emit('show-sub')">订阅</button>
  </div>
</template>
<style scoped>
.bar {
  display: flex;
  color:black;
  justify-content: space-around; /* 让按钮左右撑开 */
  align-items: center; /* 核心：让按钮在 45px 高度内垂直居中 */
  width: 80%;
  height: 60px; /* 稍微加高一点点，避开手机边缘 */
  background-color:white; 
  border:1px solid black;
  border-radius: 30px;
}

button {
  width:40px; /* 让两个按钮平分宽度 */
  margin: 0 15px; /* 左右间距 */
  height: 40px;
  background-color:rgb(255,255,255,0.8)
}


</style>
<script setup>
defineEmits(["show-home", "show-sub"]);
</script>
