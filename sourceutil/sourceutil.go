package sourceutil

import (
	"encoding/csv"
	"github.com/jszwec/csvutil"
	"log"
	"net"
	"os"
)

func NewCSVDecoder(csvFilename string) *csvutil.Decoder {
	inputReader, err := os.Open(csvFilename)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(inputReader)
	csvDecoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	csvDecoder.Register(func(data []byte, ip *net.IP) error {
		*ip = net.ParseIP(string(data))
		return nil
	})

	return csvDecoder
}
