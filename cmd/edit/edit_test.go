package edit_test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/TristanSch1/flow/cmd/edit"
	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra"
	"github.com/TristanSch1/flow/internal/infra/filesystem"
	"github.com/TristanSch1/flow/pkg/timerange"
	"github.com/TristanSch1/flow/test"
	"github.com/matryer/is"
)

type mockSessionRepository struct {
	FindByIdFn func(id string) *session.Session
}

func (m *mockSessionRepository) Save(session session.Session) error {
	return nil
}

func (m *mockSessionRepository) Delete(id string) error {
	return nil
}

func (m *mockSessionRepository) FindLastSession() *session.Session {
	return nil
}

func (m *mockSessionRepository) FindAllSessions() []session.Session {
	return []session.Session{}
}

func (m *mockSessionRepository) FindAllByProject(project string) []session.Session {
	return []session.Session{}
}

func (m *mockSessionRepository) FindAllProjects() []string {
	return []string{}
}

func (m *mockSessionRepository) FindAllProjectTags(project string) []string {
	return []string{}
}

func (m *mockSessionRepository) FindInTimeRange(timeRange timerange.TimeRange) []session.Session {
	return []session.Session{}
}

func (m *mockSessionRepository) FindById(id string) *session.Session {
	return m.FindByIdFn(id)
}

func TestEditCommand(t *testing.T) {
	is := is.New(t)

	dateProvider := infra.NewStubDateProvider()

	dateProvider.Now = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	tmpDir := t.TempDir()

	testSession := session.Session{
		Id:        "1234567",
		Project:   "project",
		StartTime: dateProvider.GetNow(),
	}

	marshaled, _ := json.Marshal(testSession)

	filename := filesystem.SessionFilename{
		Id:        testSession.Id,
		Project:   testSession.Project,
		StartTime: testSession.StartTime,
	}

	os.WriteFile(filepath.Join(tmpDir, filename.String()), marshaled, 0644)

	sessionRepository := &mockSessionRepository{
		FindByIdFn: func(id string) *session.Session {
			if id == "1234567" {
				return &testSession
			}

			return nil
		},
	}
	app := test.InitializeApp(sessionRepository, dateProvider)
	c := edit.Command(app, tmpDir)
	tt := []struct {
		name  string
		args  []string
		want  string
		error error
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
			args: []string{"abcdefg"},
			want: "Session not found",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			is := is.New(t)

			got, err := test.ExecuteCmd(t, c, tc.args...)

			is.Equal(tc.error, err)

			if tc.error == nil {
				is.Equal(tc.want, got)
			}
		})
	}
}
