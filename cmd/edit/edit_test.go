package edit_test

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/TristanShz/flow/cmd/edit"
	"github.com/TristanShz/flow/internal/application"
	"github.com/TristanShz/flow/internal/domain/session"
	"github.com/TristanShz/flow/internal/infra"
	"github.com/TristanShz/flow/internal/infra/filesystem"
	"github.com/TristanShz/flow/test"
	"github.com/matryer/is"
)

type mockSessionRepository struct {
	FindByIdFn        func(id string) *session.Session
	FindLastSessionFn func() *session.Session
}

func (m *mockSessionRepository) Save(session session.Session) error {
	return nil
}

func (m *mockSessionRepository) Delete(id string) error {
	return nil
}

func (m *mockSessionRepository) FindAllSessions(filters *application.SessionsFilters) []session.Session {
	return []session.Session{}
}

func (m *mockSessionRepository) FindAllProjects() []string {
	return []string{}
}

func (m *mockSessionRepository) FindAllProjectTags(project string) []string {
	return []string{}
}

func (m *mockSessionRepository) FindById(id string) *session.Session {
	return m.FindByIdFn(id)
}

func (m *mockSessionRepository) FindLastSession() *session.Session {
	return m.FindLastSessionFn()
}

func TestEditCommand(t *testing.T) {
	is := is.New(t)

	dateProvider := infra.NewStubDateProvider()

	dateProvider.Now = time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)

	tmpDir := t.TempDir()

	testSessions := []session.Session{
		{
			Id:        "1234567",
			Project:   "project",
			StartTime: time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2021, 1, 1, 10, 0, 0, 0, time.UTC),
		},
		{
			Id:        "7654321",
			Project:   "project",
			StartTime: time.Date(2021, 1, 1, 10, 30, 0, 0, time.UTC),
			EndTime:   time.Date(2021, 1, 1, 12, 0, 0, 0, time.UTC),
		},
	}

	for _, testSession := range testSessions {
		marshaled, _ := json.Marshal(testSession)

		filename := filesystem.SessionFilename{
			Id:        testSession.Id,
			Project:   testSession.Project,
			StartTime: testSession.StartTime,
		}

		os.WriteFile(filepath.Join(tmpDir, filename.String()), marshaled, 0644)
	}

	sessionRepository := &mockSessionRepository{
		FindByIdFn: func(id string) *session.Session {
			if id == "1234567" {
				return &testSessions[0]
			}

			return nil
		},
		FindLastSessionFn: func() *session.Session {
			return &testSessions[1]
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
