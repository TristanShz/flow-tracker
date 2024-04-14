package infra

import "time"

type StubDateProvider struct {
	Now time.Time
}

func NewStubDateProvider() *StubDateProvider {
	return &StubDateProvider{
		Now: time.Now(),
	}
}

func (s *StubDateProvider) GetNow() time.Time {
	return s.Now
}
