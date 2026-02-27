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
  
      <playbutton  class="playbutton":info="receive"/>
      <Bottom class="navbottom"
        @show-home="currentView = 'Home'"
        @show-sub="currentView = 'Sub'"
      />
    </footer>
  </div>
</template>

<style scoped >
.show {
  overflow-y: auto;
  flex: 0.8;
  overflow-x:hidden;
  padding: 5px;
  /* Increased padding to accommodate taller footer */
  padding-bottom: 150px; 
}
.fixed-footer {
  position: fixed;
  bottom: 3vh;
  left: 0;
  right: 0;
  z-index: 1000;
   /* Or any color that matches your design */
}
.navbottom{
   width:80%;
   margin:0 auto; 
}
.playbutton{
  width:93%;
  margin:0 auto 1px auto;
  height:50px;
  background-color:orange;
}
</style>
