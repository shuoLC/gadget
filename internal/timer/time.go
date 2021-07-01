package timer

import "time"

// GetNowTime 返回当前本地时间的 Time 对象
func GetNowTime() time.Time {
	return time.Now()
}

// GetCalculateTime 时间推算
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration,err := time.ParseDuration(d)
	if err != nil {
		return time.Time{},err
	}
	return currentTimer.Add(duration),nil
}