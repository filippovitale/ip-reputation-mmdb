package fileutil

import (
	"fmt"
	"path/filepath"
	"strings"
)

func CreateOutputFilename(fileIn string, extensionOut string) string {
	var extensionIn = filepath.Ext(fileIn)
	var baseFilename = strings.TrimRight(fileIn, extensionIn)
	var fileOut = fmt.Sprintf("%s.%s", baseFilename, extensionOut)
	return fileOut
}
