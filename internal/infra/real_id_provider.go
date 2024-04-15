package infra

import "github.com/TristanSch1/flow/utils"

type RealIDProvider struct{}

func (s RealIDProvider) Provide() string {
	return utils.GenerateID(7)
}
