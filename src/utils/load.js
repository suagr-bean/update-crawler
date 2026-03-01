import {ref} from'vue'
function loading(targetRef,callback){
    const page=rer(1)
    const load=ref(false) 
    const nomore=ref(false)
   const handle=async => {
    if (load.value|| nomore.value)return;
    load.value=true;
   }
}