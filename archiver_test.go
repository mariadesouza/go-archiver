package archiver

import (
	"os"
	"path/filepath"
	"testing"
)

const testPath = "/tmp/atest/"
const testFile = "a.txt"


func TestAddFileToArchive(t *testing.T){

	_, err := os.Stat(testPath + testFile)
	sourcefile := filepath.Join(testPath, testFile)
	if os.IsNotExist(err) {
		err := os.MkdirAll(testPath, 0777)
		if err != nil {
			t.Errorf("Failed to create temp directory %s: %s", testPath, err.Error())
			return
		}

		dstFile, err := os.Create(sourcefile)
		if err != nil {
			t.Errorf("Failed to write file %s%s: %s", testPath, testFile, err.Error())
			return
		}
		defer dstFile.Close()
	}

	err = AddFileToArchive("a.tar", sourcefile, testPath)
	if err != nil {
		t.Errorf("Failed to create archive with file  %s: %s", sourcefile, err.Error())
			return
	}


}