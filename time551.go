package time551

import (
	"strconv"
	"time"
)

const (
	termNow  time.Duration = 10 * time.Second
	termSec  time.Duration = 60 * time.Second
	termMin  time.Duration = 60 * time.Minute
	termHour time.Duration = 24 * time.Hour
)

func Elapsed(u time.Time) string {
	fzJst := time.FixedZone("Asia/Tokyo", 9*60*60)

	now := time.Now().In(fzJst)
	u = u.In(fzJst)

	d := now.Sub(u)

	switch {
	case u.Format("2006") != now.Format("2006"):
		return u.Format("2006.01.02")
	case d < termNow:
		return "今"
	case d < termSec:
		return strconv.Itoa(int(d.Seconds())) + "秒前"
	case d < termMin:
		return strconv.Itoa(int(d.Minutes())) + "分前"
	case d < termHour:
		return u.Format("15:04")
	default:
		return u.Format("01月02日")
	}

}
