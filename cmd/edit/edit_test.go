package edit_test

import (
	"errors"
	"testing"

	"github.com/TristanSch1/flow/cmd/edit"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/test"
	"github.com/matryer/is"
)

func TestEditCommand(t *testing.T) {
	is := is.New(t)

	sessionRepository := &infra.InMemorySessionRepository{}
	dateProvider := infra.NewStubDateProvider()
	app := test.InitializeApp(sessionRepository, dateProvider)

	tt := []struct {
		name          string
		args          []string
		givenSessions []session.Session
		want          string
		error         error
	}{
		{
			name:  "No args",
			error: errors.New("missing session ID"),
		},
		{
			name:  "Invalid ID",
			args:  []string{"hello"},
			error: errors.New("invalid ID hello"),
		},
		{
			name:  "Too many arguments",
			args:  []string{"1", "2"},
			error: errors.New("too many arguments"),
		},
		{
			name: "Session not found",
			args: []string{"1234567"},
			want: "Session not found",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			c := edit.Command(app, "")

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(tc.error, err)

			if tc.error == nil {
				is.Equal(tc.want, got)
			}
		})
	}
}
