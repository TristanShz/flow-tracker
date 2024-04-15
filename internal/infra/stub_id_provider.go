package infra

type StubIDProvider struct {
	Id string
}

func (s *StubIDProvider) Provide() string {
	return s.Id
}

func NewStubIDProvider() StubIDProvider {
	return StubIDProvider{
		Id: "stub-id",
	}
}
