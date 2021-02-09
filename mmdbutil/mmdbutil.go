package mmdbutil

import (
	"fmt"
	"github.com/maxmind/mmdbwriter"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func NewMMDB(dbType string, dbDesc string) *mmdbwriter.Tree {
	recordSize := 28
	mmdb, err := mmdbwriter.New(
		mmdbwriter.Options{
			DatabaseType:            dbType,
			Description:             map[string]string{"en": dbDesc},
			DisableIPv4Aliasing:     true,
			IncludeReservedNetworks: true,
			RecordSize:              recordSize,
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	return mmdb
}

func CreateFilename(fileIn string) string {
	extensionOut := "mmdb"
	var extensionIn = filepath.Ext(fileIn)
	var baseFilename = strings.TrimRight(fileIn, extensionIn)
	var fileOut = fmt.Sprintf("%s.%s", baseFilename, extensionOut)
	return fileOut
}

func DumpMMDBToFile(mmdbFilename string, mmdb *mmdbwriter.Tree) {
	fh, err := os.Create(mmdbFilename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = mmdb.WriteTo(fh)
	if err != nil {
		log.Fatal(err)
	}
}
