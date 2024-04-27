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
	"strings"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/pkg/timerange"
)

type Sessions []session.Session

func (s Sessions) Len() int {
	return len(s)
}

func (s Sessions) Less(i, j int) bool {
	return s[i].StartTime.Before(s[j].StartTime)
}

func (s Sessions) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type FileSystemSessionRepository struct {
	FlowFolderPath string
}

func NewFileSystemSessionRepository(flowFolderPath string) FileSystemSessionRepository {
	if _, err := os.Stat(flowFolderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(flowFolderPath, 0777); err != nil {
			log.Fatal("Error while creating .flow folder : ", err)
		}
	}

	return FileSystemSessionRepository{
		FlowFolderPath: flowFolderPath,
	}
}

func (r *FileSystemSessionRepository) getSessionFileName(s session.Session) string {
	return strconv.FormatInt(s.StartTime.Unix(), 10) + ".json"
}

func (r *FileSystemSessionRepository) readFlowFolder() ([]fs.FileInfo, error) {
	dir, err := os.Open(r.FlowFolderPath)
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

func (r *FileSystemSessionRepository) FindById(id string) *session.Session {
	fileInfos, err := r.readFlowFolder()
	if err != nil {
		log.Fatal(err)
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		filePath := filepath.Join(r.FlowFolderPath, fileInfo.Name())
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error while reading file %v : '%v'", fileInfo.Name(), err)
		}

		session, convertErr := r.rawFileToSession(file)
		if convertErr != nil {
			log.Fatalf("Invalid session data for file : %v", fileInfo.Name())
		}

		if session.Id == id {
			return session
		}
	}

	return nil
}

func (r *FileSystemSessionRepository) Save(sessionToSave session.Session) error {
	marshaled, marshaledErr := json.MarshalIndent(sessionToSave, "", "  ")

	if marshaledErr != nil {
		return marshaledErr
	}

	fullPath := filepath.Join(r.FlowFolderPath, r.getSessionFileName(sessionToSave))
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

func (r *FileSystemSessionRepository) FindAllSessions() []session.Session {
	fileInfos, err := r.readFlowFolder()
	if err != nil {
		log.Fatal(err)
	}

	sessions := Sessions{}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			continue
		}

		filePath := filepath.Join(r.FlowFolderPath, fileInfo.Name())
		file, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error while reading file %v : '%v'", fileInfo.Name(), err)
		}

		session, convertErr := r.rawFileToSession(file)
		if convertErr != nil {
			log.Fatalf("Invalid session data for file : %v", fileInfo.Name())
		}
		sessions = append(sessions, *session)
	}

	sort.Sort(sessions)

	return sessions
}

func (r *FileSystemSessionRepository) FindAllByProject(project string) []session.Session {
	allSessions := r.FindAllSessions()

	sessions := []session.Session{}

	for _, session := range allSessions {
		if session.Project == project {
			sessions = append(sessions, session)
		}
	}

	return sessions
}

func (r *FileSystemSessionRepository) FindLastSession() *session.Session {
	fileInfos, err := r.readFlowFolder()
	if err != nil {
		log.Fatal(err)
	}

	fileNames := []string{}

	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			before, _ := strings.CutSuffix(fileInfo.Name(), ".json")
			fileNames = append(fileNames, before)
		}
	}

	if len(fileNames) == 0 {
		return nil
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
	lastSessionFile := fileNames[0] + ".json"

	lastSessionFilePath := filepath.Join(r.FlowFolderPath, lastSessionFile)

	fileData, err := os.ReadFile(lastSessionFilePath)
	if err != nil {
		log.Fatalf("Error while reading file %v", lastSessionFilePath)
	}

	session, convertErr := r.rawFileToSession(fileData)

	if convertErr != nil {
		log.Fatalf("Invalid session data for file : %v", lastSessionFilePath)
	}

	return session
}

func (r *FileSystemSessionRepository) FindAllProjects() []string {
	sessions := r.FindAllSessions()

	projects := []string{}

	for _, session := range sessions {
		if slices.Contains(projects, session.Project) {
			continue
		}

		projects = append(projects, session.Project)
	}

	return projects
}

func (r *FileSystemSessionRepository) FindAllProjectTags(project string) []string {
	sessionsForProject := r.FindAllByProject(project)

	tags := []string{}

	for _, session := range sessionsForProject {
		for _, tag := range session.Tags {
			if slices.Contains(tags, tag) {
				continue
			}

			tags = append(tags, tag)
		}
	}

	return tags
}

func (r *FileSystemSessionRepository) FindInTimeRange(timeRange timerange.TimeRange) []session.Session {
	// TODO: Optimize this function by reading only the files that are in the time range
	allSessions := r.FindAllSessions()

	sessions := []session.Session{}

	for _, session := range allSessions {
		if timeRange.JustUntil() {
			if session.StartTime.Before(timeRange.Until) {
				sessions = append(sessions, session)
			}
		} else if timeRange.JustSince() {
			if session.StartTime.After(timeRange.Since) {
				sessions = append(sessions, session)
			}
		} else if timeRange.SinceAndUntil() {
			if session.StartTime.After(timeRange.Since) && session.StartTime.Before(timeRange.Until) {
				sessions = append(sessions, session)
			}
		} else {
			sessions = append(sessions, session)
		}
	}
	return sessions
}
