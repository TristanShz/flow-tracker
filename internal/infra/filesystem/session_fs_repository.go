package filesystem

import (
	"encoding/json"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"

	"github.com/TristanSch1/flow/internal/domain/session"
)

const (
	FlowFolderName = ".flow"
)

type FileSystemSessionRepository struct {
	FlowFolderPath string
}

func NewFileSystemSessionRepository(flowFolderPath string) (FileSystemSessionRepository, error) {
	path := filepath.Join(flowFolderPath, FlowFolderName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0777); err != nil {
			return FileSystemSessionRepository{}, err
		}
	}

	return FileSystemSessionRepository{
		FlowFolderPath: flowFolderPath,
	}, nil
}

func (r *FileSystemSessionRepository) getFlowPath() string {
	return filepath.Join(r.FlowFolderPath, FlowFolderName)
}

func (r *FileSystemSessionRepository) getSessionFileName(s session.Session) string {
	return s.Id + ".json"
}

func (r *FileSystemSessionRepository) readFlowFolder() ([]fs.FileInfo, error) {
	dir, err := os.Open(r.getFlowPath())
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}

func (r *FileSystemSessionRepository) Save(sessionToSave session.Session) error {
	marshaled, marshaledErr := json.Marshal(sessionToSave)

	if marshaledErr != nil {
		return marshaledErr
	}

	fullPath := filepath.Join(r.getFlowPath(), r.getSessionFileName(sessionToSave))
	saveErr := os.WriteFile(fullPath, marshaled, 0666)

	if saveErr != nil {
		return saveErr
	}

	return nil
}

func (r *FileSystemSessionRepository) rawFileToSession(raw []byte) (*session.Session, error) {
	var sessionData session.Session
	if err := json.Unmarshal(raw, &sessionData); err != nil {
		return nil, err
	}

	return &sessionData, nil
}

func (r *FileSystemSessionRepository) FindAllSessions() ([]session.Session, error) {
	fileInfos, err := r.readFlowFolder()
	if err != nil {
		log.Fatal(err)
	}

	sessions := []session.Session{}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		file, err := os.ReadFile(fileInfo.Name())
		if err != nil {
			log.Fatalf("Error while reading file %v", fileInfo.Name())
		}

		session, convertErr := r.rawFileToSession(file)
		if convertErr != nil {
			log.Fatalf("Invalid session data for file : %v", fileInfo.Name())
		}
		sessions = append(sessions, *session)
	}

	return sessions, nil
}

func (r *FileSystemSessionRepository) FindAllByProject(project string) ([]session.Session, error) {
	allSessions, err := r.FindAllSessions()
	if err != nil {
		log.Fatalf("Error while reading all session files : %v", err)
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
	fileInfos, err := r.readFlowFolder()
	if err != nil {
		log.Fatal(err)
	}

	fileNames := []string{}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}

	if len(fileNames) == 0 {
		return nil, nil
	}

	sort.Slice(fileNames, func(i, j int) bool {
		numI, err := strconv.ParseInt(fileNames[i], 10, 64)
		if err != nil {
			log.Fatal("Error when converting file name to integer:", err)
		}
		numJ, err := strconv.ParseInt(fileNames[j], 10, 64)
		if err != nil {
			log.Fatal("Error when converting file name to integer:", err)
		}
		return numI > numJ
	})
	lastSessionFileName := fileNames[0]

	lastSessionFilePath := filepath.Join(r.getFlowPath(), lastSessionFileName)

	fileData, err := os.ReadFile(lastSessionFilePath)
	if err != nil {
		log.Fatalf("Error while reading file %v", lastSessionFilePath)
	}

	session, convertErr := r.rawFileToSession(fileData)

	if convertErr != nil {
		log.Fatalf("Invalid session data for file : %v", lastSessionFilePath)
	}

	return session, nil
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
