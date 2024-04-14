package list_test

import (
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/tests"
)

func TestListProjects_Success(t *testing.T) {
	f := tests.GetSessionFixture(t)

	f.GivenSomeSessions([]session.Session{{
		StartTime: time.Date(2024, time.April, 14, 10, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 13, 10, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"add-todo"},
	}, {
		StartTime: time.Date(2024, time.April, 14, 14, 12, 0, 0, time.UTC),
		EndTime:   time.Date(2024, time.April, 14, 15, 12, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"start-usecase"},
	}})

	f.WhenGettingListOfProjects()

	f.ThenProjectsShouldBe([]string{"MyTodo", "Flow"})
}
