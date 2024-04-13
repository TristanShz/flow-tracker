package fs

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"slices"

	"github.com/TristanSch1/flow/internal/domain/session"
)

type FileSystemSessionRepository struct {
	FlowPath string
}

func (r *FileSystemSessionRepository) Save(s session.Session) error {
	sessions, err := r.FindAll()
	if err != nil {
		return err
	}

	startedSessionIndex := slices.IndexFunc(sessions, func(s session.Session) bool {
		return s.StartTime.Equal(s.StartTime)
	})

	if startedSessionIndex == -1 {
		sessions = append(sessions, s)
	} else {
		sessions[startedSessionIndex] = s
	}

	marshaled, marshaledErr := json.Marshal(sessions)

	if marshaledErr != nil {
		return marshaledErr
	}

	fullPath := filepath.Join(r.FlowPath, ".flow")
	saveErr := os.WriteFile(fullPath, marshaled, 0666)

	if saveErr != nil {
		return saveErr
	}

	return nil
}

func (r *FileSystemSessionRepository) FindAll() ([]session.Session, error) {
	file, err := fs.ReadFile(os.DirFS(r.FlowPath), ".flow")
	if os.IsNotExist(err) {
		return []session.Session{}, nil
	}
	if err != nil {
		return nil, err
	}

	var sessions []session.Session
	unmarshalErr := json.Unmarshal(file, &sessions)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return sessions, nil
}

func (r *FileSystemSessionRepository) FindAllByProject(project string) ([]session.Session, error) {
	allSessions, err := r.FindAll()
	if err != nil {
		return nil, err
	}

	sessions := []session.Session{}

	for _, session := range allSessions {
		if session.Project == project {
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

func (r *FileSystemSessionRepository) FindLastSession() (*session.Session, error) {
	sessions, err := r.FindAll()
	if err != nil {
		return nil, err
	}

	if len(sessions) == 0 {
		return nil, nil
	}
	return &sessions[len(sessions)-1], nil
}
