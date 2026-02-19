<template>
  <div class="play-button-container">
    <p class="title">{{ info.title }}</p>
    <audio autoplay controls :src="info.auto_link" style="width: 100%;"></audio>
  </div>
</template>

<script setup>
const props = defineProps({
  info: {
    type: Object,
    default: () => ({})
  }
})
</script>

<style scoped>
.play-button-container {
  display: flex;
  flex-direction: column; /* Arrange title and player vertically */
  align-items: center; /* Center them horizontally */
  background-color: #f0f0f0;
  padding: 10px 5px; /* Add some padding */
}

.title {
  font-weight: 500;
  margin-bottom: 8px; /* Add space between title and player */
  text-align: center;
  width: 95%;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis; /* Use '...' for very long titles */
}
</style>
