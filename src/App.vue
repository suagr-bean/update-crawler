<script setup>
import { ref } from "vue";
import Top from "./components/Top.vue";
import Bottom from "./components/Bottom.vue";
import Home from "./views/home.vue";
import Sub from "./views/sub.vue";
import Detail from './views/detail.vue'
import playbutton from './components/playbutton.vue'
const currentView = ref("Home");

const url=ref('')
const ToDetail=(load)=>{
 currentView.value="detail";
 url.value=load.url;
}
const receive =ref({})
const Send=(item)=>{
   receive.value=item;
}
</script>

<template>
  <div class="all">
    <Top class="top" />
    <main class="show">
      <Home v-if="currentView === 'Home'" @detail="ToDetail"/>
      <Sub v-if="currentView === 'Sub'" />
      <Detail v-if="currentView=='detail'" :apiurl="url" @control="Send"></Detail>
     
    </main>
    <footer class="fixed-footer">
  
      <playbutton :info="receive"/>
      <Bottom
        @show-home="currentView = 'Home'"
        @show-sub="currentView = 'Sub'"
      />
    </footer>
  </div>
</template>

<style scoped>
.show {
  overflow-y: auto;
  flex: 1;
  padding: 8px;
  /* Increased padding to accommodate taller footer */
  padding-bottom: 150px; 
}
.fixed-footer {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  background-color: white; /* Or any color that matches your design */
}
</style>
