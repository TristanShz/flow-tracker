package viewsessionsreport

import "time"

type Command struct {
	Since   time.Time
	Until   time.Time
	Project string
	Format  string
}
