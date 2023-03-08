/**
* @Author: caoyongfei
* @Date: 2022/11/4
* @Description:
**/

package time

import "time"

// 时间格式化 go中使用具体的数字表示对应的 星期、年、月、日、时、分、秒   Monday 2006 01 02 15:04:05
// 如 format="2006/01/02 15:04:05" 表示 "年/月/日 时:分:秒"

/**
* @FuncName: NewDate
* @Description: time包下构造时间的方法， 传入任意的年、月、日、时、分、秒、纳秒、时区
* @Params <t time.Time>:
* @Return <time.Time>:
**/
func NewDate(t time.Time) time.Time {
	newT := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	return newT
}

/**
* @FuncName: LocationTimeNow
* @Description: 获取本地当前时间
* @Return <time.Time>:
**/
func LocationTimeNow() time.Time {
	return time.Now()
}

/**
* @FuncName: UnixTimestampLocation
* @Description: 获取本地时区的当前时间戳
* @Return <int64>:
**/
func UnixTimestampLocation() int64 {
	return time.Now().Unix()
}

/**
* @FuncName: UnixNanoTimestampLocation
* @Description: 获取本地时区的当前时间戳(纳秒级)
* @Return <int64>:
**/
func UnixNanoTimestampLocation() int64 {
	return time.Now().UnixNano()
}

/**
* @FuncName: UTCTimeNow
* @Description: 获取utc当前时间
* @Return <time.Time>:
**/
func UTCTimeNow() time.Time {
	return time.Now().UTC()
}

/**
* @FuncName: UnixTimestampUTC
* @Description: 获取utc的当前时间戳
* @Return <int64>:
**/
func UnixTimestampUTC() int64 {
	return time.Now().UTC().Unix()
}

/**
* @FuncName: UnixNanoTimestampUTC
* @Description: 获取utc的当前时间戳(纳秒级)
* @Return <int64>:
**/
func UnixNanoTimestampUTC() int64 {
	return time.Now().UTC().UnixNano()
}

/**
* @FuncName: Timestamp2Time
* @Description: 时间戳转时间
* @Params <timestamp int64>:
* @Return <time.Time>:
**/
func Timestamp2Time(timestamp int64) time.Time {
	return time.Unix(timestamp, 0)
}


/**
* @FuncName: FormatTime
* @Description: 将时间格式化为 "2006-01-02 15:04:05"
* @Params <t time.Time>:
* @Return <string>:
**/
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

/**
* @FuncName: String2UTCTime
* @Description: 格式化时间转为UTC时区的时间类型， 注意这里是UTC时间
* @Params <layout string>: 格式化时间的格式布局 "2006-01-02 15:04:05" 表示  format="年-月-日 时:分:秒"
* @Params <value string>: 格式化时间的值 如 "2022-10-11 18:23:12" 应按format="年-月-日 时:分:秒"解析
* @Return <time.Time>:
* @Return <error>:
**/
func String2UTCTime(layout string, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

/**
* @FuncName: String2LocationTime
* @Description: 格式化时间转为某时区的时间类型， 需要传入时区，如中国是东8区 即UTC+8
* @Params <layout string>: 格式化时间的格式布局 "2006-01-02 15:04:05" 表示  format="年-月-日 时:分:秒"
* @Params <value string>: 格式化时间的值 如 "2022-10-11 18:23:12" 应按format="年-月-日 时:分:秒"解析
* @Params <loc *time.Location>: 想要转的时区,  比如中国北京属于东8区，time.Local就是输出当前的时区
* @Return <time.Time>:
* @Return <error>:
**/
func String2LocationTime(layout string, value string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, value, loc)
}

/**
* @FuncName: SubBetweenTime
* @Description: 求两个时间的差值
* @Params <t1 time.Time>:
* @Params <t2 time.Time>:
* @Return <time.Duration>:
**/
func SubBetweenTime(t1 time.Time, t2 time.Time) time.Duration {
	return t1.Sub(t2)
}

/**
* @FuncName: AddTime
* @Description: 在time的基础上调整时间， d < 0 表示将时间往前推  d > 0 表示将时间往后推
* @Params <t1 time.Time>:
* @Params <d time.Duration>:
* @Return <time.Time>:
**/
func AddTime(t1 time.Time, d time.Duration) time.Time {
	return t1.Add(d)
}
