package report

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"cv-url/internal/utils"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Create(dto CreateDto) (string, error) {
	rep := s.BuildWithHash(dto)

	jsonFile, err := s.SaveToJson(rep)
	if err != nil {
		return "", err
	}

	zipFile := "source_code.zip"
	if err := utils.ZipSource(zipFile); err != nil {
		return "", err
	}

	if err := utils.SendEmail(
		dto.Email,
		dto.SMTPLogin,
		dto.SMTPPassword,
		dto.SMTPServer,
		dto.SMTPPort,
		rep.UserID,
		jsonFile,
		zipFile,
	); err != nil {
		return "", err
	}

	return rep.UserID, nil
}

func (s *Service) BuildWithHash(dto CreateDto) *Report {
	hash := sha256.Sum256([]byte(dto.CVURL))
	hashStr := hex.EncodeToString(hash[:])
	userID := fmt.Sprintf("%s-%s", hashStr[:8], utils.RandomString(4))

	return &Report{
		CVURL:     dto.CVURL,
		Hash:      hashStr,
		UserID:    userID,
		Email:     dto.Email,
		Timestamp: time.Now().UTC(),
	}
}

func (s *Service) SaveToJson(rep *Report) (string, error) {
	fileName := fmt.Sprintf("report_%s.json", rep.UserID)
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(rep); err != nil {
		return "", err
	}
	return fileName, nil
}