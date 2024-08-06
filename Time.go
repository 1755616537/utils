package utils

import (
	"errors"
	"strconv"
	"time"
)

// 10位时间戳
func Get10Timestamp(date time.Time) string {
	return strconv.FormatInt(date.Unix(), 10)
}

func TimeToFormat(date time.Time) string {
	return date.Format("2006/01/02 15:04:05.000000")

}

func StringTotime(timeLayout, toBeCharge string) (*time.Time, error) {
	loc, err := time.LoadLocation("Local") //重要：获取时区
	if err != nil {
		return nil, errors.New("时区获取失败")
	}
	theTime, err := time.ParseInLocation(toBeCharge, timeLayout, loc) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return nil, errors.New("时间转换失败")
	}
	return &theTime, nil
}

// 时间转换 【20060102】uint32类型
func DateTyUint32(date time.Time) (uint32, error) {
	return TyUint32(date.Format("20060102"))
}

// 获取前n天工作日
func DateAoArr(date time.Time, daysBack int, dt bool) ([]uint32, error) {
	var workDates []uint32

	// 向前迭代直到找到
	for len(workDates) < daysBack {
		if !dt {
			// 减去一天
			date = date.AddDate(0, 0, -1)
		}

		// 检查是否为工作日（周一至周五）
		if !TimeIsWeekend(date) {
			u32date, err := DateTyUint32(date)
			if err != nil {
				return nil, err
			}

			// 添加到结果列表
			workDates = append(workDates, u32date)
		}

		if dt {
			// 减去一天
			date = date.AddDate(0, 0, -1)
		}
	}

	return workDates, nil
}

// 判断给定的时间是否是周末（星期六或星期天）
func TimeIsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
