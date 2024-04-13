package application

import "github.com/tristansch1/flow/internal/domain/entry"

type EntryRepository interface {
	save(entry entry.Entry) error
}
