package edit_test

import (
	"errors"
	"testing"

	"github.com/TristanSch1/flow/cmd/edit"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/internal/infra/filesystem"
	"github.com/TristanSch1/flow/test"
	"github.com/matryer/is"
)

func TestEditCommand(t *testing.T) {
	is := is.New(t)

	sessionRepository := filesystem.NewFileSystemSessionRepository("/tmp/flow")
	dateProvider := infra.NewStubDateProvider()
	app := test.InitializeApp(&sessionRepository, dateProvider)

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
		{
			name: "Edit session",
			args: []string{"1234567"},
			givenSessions: []session.Session{
				{
					Id:        "1234567",
					Project:   "project",
					StartTime: dateProvider.GetNow(),
				},
			},
			want: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			for _, s := range tc.givenSessions {
				sessionRepository.Save(s)
			}

			c := edit.Command(app, "/tmp/flow")

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(tc.error, err)

			if tc.error == nil {
				is.Equal(tc.want, got)
			}
		})
	}
}
