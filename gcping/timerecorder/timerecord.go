package timerecorder

import (
	"time"
)

type TimeRecorder interface {
	// 获取花费时间
	Cost() time.Duration

	//设置花费时间
	SetCost(time.Duration)
}

func TimeCostDecorator(rec TimeRecorder, f func() error) func() error {
	return func() error {
		startTime := time.Now()
		err := f()
		if err != nil {
			// fmt.Println("启动连接错误，错误信息为：", err)
			return err
		} else {
			endTime := time.Now()
			timeCost := endTime.Sub(startTime)
			// fmt.Println(reflect.TypeOf(timeCost))
			rec.SetCost(timeCost)
		}
		return nil
	}
}

// 结构体存储花费时间并实现TimeRecoder接口
type timeRecorder struct {
	cost time.Duration
}

// * 设置花费的时间
func (tr *timeRecorder) SetCost(cost time.Duration) {
	tr.cost = cost
}

// * 获取花费的时间
func (tr *timeRecorder) Cost() time.Duration {
	return tr.cost
}

func NewTimeRecorder() TimeRecorder {
	/*
			  ? 这里返回类型是TimeRecorder
			  ? 而是实现TimeRecorder接口的实际是*timeRecorder类型
		      ? 故return 需要传递&timeRecorder
	*/
	return &timeRecorder{}
}
