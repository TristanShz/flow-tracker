package infra

import "github.com/TristanShz/flow/utils"

type RealIDProvider struct{}

func (s RealIDProvider) Provide() string {
	return utils.GenerateID(7)
}
