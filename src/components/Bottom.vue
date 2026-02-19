<template>
  <div class="bar">
    <button class="home" @click="$emit('show-home')">首页</button>
    <button @click="$emit('show-sub')">订阅</button>
  </div>
</template>
<style scoped>
.bar {
  display: flex;
  justify-content: space-around; /* 让按钮左右撑开 */
  align-items: center; /* 核心：让按钮在 45px 高度内垂直居中 */
  width: 100%;
  height: 50px; /* 稍微加高一点点，避开手机边缘 */
  background-color: white; /* 给个背景色，防止透出底层内容 */
  border-top: 1px solid #eee;
}

button {
  flex: 1; /* 让两个按钮平分宽度 */
  margin: 0 15px; /* 左右间距 */
  height: 35px;
  border-radius: 8px;
}
</style>
<script setup>
defineEmits(["show-home", "show-sub"]);
</script>
