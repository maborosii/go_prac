package timerecorder

import (
	"testing"
	"time"
)

func TestTimeCostDecorator(t *testing.T) {
	testFunc := func() error {
		time.Sleep(time.Duration(1) * time.Second)
		return nil
	}

	type args struct {
		rec TimeRecorder
		f   func() error
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"test time cost decorator",
			args{
				NewTimeRecorder(),
				testFunc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeCostDecorator(tt.args.rec, tt.args.f)
			got()
			if tt.args.rec.Cost().Round(time.Second) != time.Duration(1)*time.Second.Round(time.Second) {
				t.Errorf("Record time cost abnormal, recorded cost: %s, real cost: %s",
					tt.args.rec.Cost().String(), time.Duration(1)*time.Second)
			}
		})
	}
}
