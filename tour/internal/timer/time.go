package timer

import "time"

func GetNewTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location) //获得时间的时候也需要注意时区
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(duration), nil
}
