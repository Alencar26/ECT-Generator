package files

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ECTFile struct {
	ECTType   string
	Author    string
	TestSuite string
	Evicences []string
}

func NewECTFile() (*ECTFile, error) {
	path := GetExecutablePath()
	fileName := filepath.Base(path)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

	parts := strings.Split(fileName, "_")

	if len(parts) < 3 {
		err := errors.New("ATENÇÃO!! Formato inválido. Esperado: ECT_<TestSuite>_<Author>")
		return nil, err
	}

	dir := GetExecutableDir(path)

	return &ECTFile{
		ECTType:   parts[0],
		TestSuite: parts[1],
		Author:    parts[2],
		Evicences: GetEvidences(dir),
	}, nil
}

func GetEvidences(dir string) []string {

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var evidences []string

	for _, file := range files {
		if !file.IsDir() && IsImage(file.Name()) {
			evidences = append(evidences, file.Name())
		}
	}

	//ordenando pela numeração natural (não pela comparação de string)
	sort.Slice(evidences, func(i, j int) bool {
		return extractNumber(evidences[i]) < extractNumber(evidences[j])
	})

	return evidences
}

func GetExecutablePath() string {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return path
}

func GetExecutableDir(executablePath string) string {
	return filepath.Dir(executablePath)
}

func IsImage(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))
	switch ext {
	case ".png", ".jpg", ".jpeg", ".gif", ".webp":
		return true
	default:
		return false
	}
}

func extractNumber(fileName string) int {
	var numRegex = regexp.MustCompile(`\d+`)

	numStr := numRegex.FindString(fileName)
	if numStr == "" {
		return 0
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}
	return num
}
