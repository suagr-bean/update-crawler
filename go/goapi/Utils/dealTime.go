package Utils

import "time"
//工具类
func DealTime(times int)time.Time {
	deal := time.Duration(times)*time.Minute
    return time.Now().Add(deal)
}
