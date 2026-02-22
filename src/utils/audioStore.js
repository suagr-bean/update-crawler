const Storage_key="audio_record"
export const AudioStore ={
    /**
     * 保存进度条
     * @param {string} url  哪个音频
     * @param {int} time  时间
     */
    save(url,time){
        try{
         const history=JSON.parse(localStorage.getItem(Storage_key)||'{}')
         history[url]=time
         localStorage.setItem(Storage_key,JSON.stringify(history))
        }catch(error){
           console.log("存储失败",error)
        }
    },
    /**
     *  获取音频的时间
     * @param {string} url 
     * @returns {time} 时间
     */
    get(url){
        try{
         const history=JSON.parse(localStorage.getItem(Storage_key)||'{}');
         return history[url]?parseFloat(history [url]):0;
        }catch(error){
            console.log("获取失败",error)
        }
    }
}