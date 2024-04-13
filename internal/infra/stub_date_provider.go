package infra

import "time"

type StubDateProvider struct {
	Now time.Time
}

func (s *StubDateProvider) GetNow() time.Time {
	return s.Now
}
