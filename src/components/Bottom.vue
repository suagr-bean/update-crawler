<template>
  <div class="bar">
    <button class="home" @click="$emit('show-home')">首页</button>
    <button @click="$emit('show-sub')">订阅</button>
  </div>
</template>
<style scoped>
.bar {
  display: flex;
  color:sepia;
  justify-content: space-around; /* 让按钮左右撑开 */
  align-items: center; /* 核心：让按钮在 45px 高度内垂直居中 */
  width: 100%;
  height: 80px; /* 稍微加高一点点，避开手机边缘 */
  background-color:#FAF0DC; /* 给个背景色，防止透出底层内容 */
}

button {
  width:70px; /* 让两个按钮平分宽度 */
  margin: 0 15px; /* 左右间距 */
  height: 40px;
  border-radius: 8px;
}
</style>
<script setup>
defineEmits(["show-home", "show-sub"]);
</script>
