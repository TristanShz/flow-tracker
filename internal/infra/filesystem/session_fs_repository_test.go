package filesystem_test

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra/filesystem"
	"github.com/TristanSch1/flow/pkg/timerange"
	"github.com/matryer/is"
)

const (
	TestFolderPath = "./.flow"
)

func setup() {
	os.RemoveAll("./.flow")
}

func TestConstructorCreateFolder_Success(t *testing.T) {
	setup()

	filesystem.NewFileSystemSessionRepository(TestFolderPath)

	path := filepath.Join(TestFolderPath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("File not found at location %v", path)
	}
}

func TestFileSystemSessionRepository_Save(t *testing.T) {
	is := is.New(t)
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	tt := []struct {
		name    string
		session session.Session
		error   error
	}{
		{
			name: "Success",
			session: session.Session{
				Id:        "1",
				StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
				Project:   "Flow",
				Tags:      []string{"test-save"},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := repository.Save(tc.session)
			is.NoErr(err)
		})
	}
}

func TestFileSystemSessionRepository_FindAllSessions(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"test-save"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		Project:   "Flow",
	})

	got := repository.FindAllSessions()

	want := []session.Session{
		{
			Id:        "1",
			StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
			Project:   "Flow",
			Tags:      []string{"test-save"},
		},
		{
			Id:        "2",
			StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
			Project:   "Flow",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FileSystemSessionRepository.FindAll() = %v, want %v", got, want)
	}
}

func TestFindAllSessions_NoSessions_Success(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	got := repository.FindAllSessions()

	want := []session.Session{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FileSystemSessionRepository.FindAll() = %v, want %v", got, want)
	}
}

func TestFileSystemSessionRepository_FindLastSession(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"test-save"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		Project:   "Flow",
	})

	got := repository.FindLastSession()

	want := session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		Project:   "Flow",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("FileSystemSessionRepository.FindLastSession() = %v, want %v", *got, want)
	}
}

func TestFileSystemSessionRepository_FindAllProjects(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"test"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 22, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
	})

	repository.Save(session.Session{
		Id:        "3",
		StartTime: time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
	})

	got := repository.FindAllProjects()

	want := []string{"Flow", "MyTodo"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FileSystemSessionRepository.FindAllProjects() = %v, want %v", got, want)
	}
}

func TestFileSystemSessionRepository_FindAllProjectTags(t *testing.T) {
	is := is.New(t)
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"tests", "integration"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 22, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"add-todo", "update-todo"},
	})

	repository.Save(session.Session{
		Id:        "3",
		StartTime: time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"update-todo", "delete-todo"},
	})

	tt := []struct {
		name string
		want []string
	}{
		{
			name: "Flow",
			want: []string{"tests", "integration"},
		},
		{
			name: "MyTodo",
			want: []string{"add-todo", "update-todo", "delete-todo"},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := repository.FindAllProjectTags(tc.name)
			sort.Strings(got)
			sort.Strings(tc.want)
			is.Equal(got, tc.want)
		})
	}
}

func TestFileSystemSessionRepository_FindInTimeRange(t *testing.T) {
	setup()
	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)
	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"tests", "integration"},
	})
	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"add-todo", "update-todo"},
	})
	repository.Save(session.Session{
		Id:        "3",
		StartTime: time.Date(2024, 4, 18, 21, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"delete-todo"},
	})

	tt := []struct {
		name string
		args timerange.TimeRange
		want []session.Session
	}{
		{
			name: "All",
			args: timerange.TimeRange{},
			want: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"tests", "integration"},
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo", "update-todo"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, 4, 18, 21, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"delete-todo"},
				},
			},
		},
		{
			name: "Since",
			args: timerange.TimeRange{
				Since: time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
			},
			want: []session.Session{
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo", "update-todo"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, 4, 18, 21, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"delete-todo"},
				},
			},
		},
		{
			name: "Until",
			args: timerange.TimeRange{
				Until: time.Date(2024, 4, 17, 20, 1, 0, 0, time.UTC),
			},
			want: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"tests", "integration"},
				},
			},
		},
		{
			name: "Since and Until",
			args: timerange.TimeRange{
				Since: time.Date(2024, 4, 17, 17, 0, 0, 0, time.UTC),
				Until: time.Date(2024, 4, 17, 22, 0, 0, 0, time.UTC),
			},
			want: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"tests", "integration"},
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo", "update-todo"},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := repository.FindInTimeRange(tc.args); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FileSystemSessionRepository.FindInTimeRange() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFileSystemSessionRepository_FindById(t *testing.T) {
	setup()
	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)
	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"test-save"},
	})
	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		Project:   "Flow",
	})

	tt := []struct {
		name string
		id   string
		want *session.Session
	}{
		{
			name: "Existing session",
			id:   "1",
			want: &session.Session{
				Id:        "1",
				StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
				EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
				Project:   "Flow",
				Tags:      []string{"test-save"},
			},
		},
		{
			name: "Non-existing session",
			id:   "3",
			want: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := repository.FindById(tc.id); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("FileSystemSessionRepository.FindById() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestFileSystemSessionRepository_Delete(t *testing.T) {
	is := is.New(t)
	setup()
	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	tt := []struct {
		error         error
		name          string
		id            string
		want          []session.Session
		givenSessions []session.Session
	}{
		{
			name: "Existing session",
			id:   "1",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					Project:   "Flow",
				},
			},
			want: []session.Session{
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					Project:   "Flow",
				},
			},
		},
		{
			name: "Non-existing session",
			id:   "3",
			givenSessions: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					Project:   "Flow",
				},
			},
			want: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					Project:   "Flow",
				},
			},
			error: errors.New("session with id 3 not found"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			for _, s := range tc.givenSessions {
				repository.Save(s)
			}
			err := repository.Delete(tc.id)

			is.Equal(err, tc.error)

			if tc.error != nil {
				got := repository.FindAllSessions()

				is.Equal(got, tc.want)
			}
		})
	}
}
