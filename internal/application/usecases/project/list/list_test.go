package list_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestListProjects_Success(t *testing.T) {
	f := tests.GetSessionFixture(t)
	tt := []struct {
		name          string
		givenSessions []session.Session
		want          []string
	}{
		{
			name: "Two projects",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 10, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo"},
				}, {
					Id:        "2",
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: []string{"MyTodo", "Flow"},
		},
		{
			name:          "No sessions",
			givenSessions: []session.Session{},
			want:          []string{},
		},
		{
			name: "Two sessions with the same project",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 13, 10, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				}, {
					Id:        "2",
					StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
					EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"start-usecase"},
				},
			},
			want: []string{"Flow"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			f.GivenSomeSessions(tc.givenSessions)

			f.WhenGettingListOfProjects()

			f.ThenProjectsShouldBe(tc.want)
		})
	}
}
