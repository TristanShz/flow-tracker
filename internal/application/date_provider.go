package application

import "time"

type DateProvider interface {
	GetNow() time.Time
}
