package infra

import "time"

type RealDateProvider struct{}

func (d *RealDateProvider) GetNow() time.Time {
	return time.Now()
}
