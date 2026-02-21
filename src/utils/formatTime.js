export const FormatTime=(seconds)=>{
    if (isNaN(seconds))return "00:00";
    const min=Math.floor(seconds/60);
    const sec=Math.floor(seconds%60);
    const hour=Math.floor(min/60);
    
    const mm=String(min%60).padStart(2,'0');
    const ss=String(sec).padStart(2,'0');
    const hh=String(hour).padStart(2,'0');
    return `${hh}:${mm}:${ss}`
}