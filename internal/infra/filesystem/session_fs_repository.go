package filesystem

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
)

const (
	FlowFolderName = ".flow"
)

type FileSystemSessionRepository struct {
	FlowFolderPath string
}

func (r *FileSystemSessionRepository) getFlowPath() string {
	return filepath.Join(r.FlowFolderPath, FlowFolderName)
}

func (r *FileSystemSessionRepository) getSessionFolderPath() string {
	todaysDate := time.Now().Format("02012006")
	return filepath.Join(r.getFlowPath(), todaysDate)
}

func (r *FileSystemSessionRepository) getSessionFileName(s session.Session) string {
	return s.Id + ".json"
}

func (r *FileSystemSessionRepository) checkForExistingFlowFolderOrCreate() error {
	if _, err := os.Stat(r.getFlowPath()); os.IsNotExist(err) {
		if err := os.MkdirAll(r.getFlowPath(), 0777); err != nil {
			return err
		}
	}

	return nil
}

func (r *FileSystemSessionRepository) Save(sessionToSave session.Session) error {
	r.checkForExistingFlowFolderOrCreate()

	sessions, err := r.FindAllSessions()
	if err != nil {
		return err
	}

	startedSessionIndex := slices.IndexFunc(sessions, func(s session.Session) bool {
		return s.Equals(sessionToSave)
	})

	if startedSessionIndex == -1 {
		sessions = append(sessions, sessionToSave)
	} else {
		sessions[startedSessionIndex] = sessionToSave
	}

	marshaled, marshaledErr := json.Marshal(sessions)

	if marshaledErr != nil {
		return marshaledErr
	}

	fullPath := filepath.Join(r.getSessionFolderPath(), r.getSessionFileName(sessionToSave))
	saveErr := os.WriteFile(fullPath, marshaled, 0666)

	if saveErr != nil {
		return saveErr
	}

	return nil
}

func (r *FileSystemSessionRepository) FindAllSessions() ([]session.Session, error) {
	sessions := []session.Session{}

	dir, err := os.Open(r.getFlowPath())
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			// TODO: Ici on arrive dans le dossier correspodant a un jour, il faut lire tout les fichiers dedans
			continue
		}

		filePath := filepath.Join(r.getFlowPath(), fileInfo.Name())
		file, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		var sessionData session.Session
		if err := json.Unmarshal(file, &sessionData); err != nil {
			return nil, err
		}

		sessions = append(sessions, sessionData)
	}

	return sessions, nil
}

func (r *FileSystemSessionRepository) FindAllByProject(project string) ([]session.Session, error) {
	allSessions, err := r.FindAllSessions()
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
	sessions, err := r.FindAllSessions()
	if err != nil {
		return nil, err
	}

	if len(sessions) == 0 {
		return nil, nil
	}
	return &sessions[len(sessions)-1], nil
}

func (r *FileSystemSessionRepository) FindAllProjects() ([]string, error) {
	sessions, _ := r.FindAllSessions()

	projects := []string{}

	for _, session := range sessions {
		if slices.Contains(projects, session.Project) {
			continue
		}

		projects = append(projects, session.Project)
	}

	return projects, nil
}

func (r *FileSystemSessionRepository) FindAllProjectTags(project string) ([]string, error) {
	sessionsForProject, _ := r.FindAllByProject(project)

	tags := []string{}

	for _, session := range sessionsForProject {
		for _, tag := range session.Tags {
			if slices.Contains(tags, tag) {
				continue
			}

			tags = append(tags, tag)
		}
	}

	return tags, nil
}
