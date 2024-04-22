package viewsessionsreport

import "time"

type Command struct {
	From    time.Time
	To      time.Time
	Project string
	Format  string
}
