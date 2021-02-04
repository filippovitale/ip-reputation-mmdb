package main

import (
	"encoding/csv"
	"fmt"
	"github.com/filippovitale/ip-reputation-mmdb/fileutil"
	"github.com/filippovitale/ip-reputation-mmdb/iputil"
	"github.com/jszwec/csvutil"
	"io"
	"log"
	"net"
	"os"
)

const CityAsReservedIPNet = "reserved"

type NA04Subset struct {
	IPv4Begin net.IP `csv:"ipv4-begin-address"`
	IPv4End   net.IP `csv:"ipv4-end-address"`
	Country2  string `csv:"edge-two-letter-country"`
	City      string `csv:"edge-city"`
	Latitude  string `csv:"edge-latitude"`
	Longitude string `csv:"edge-longitude"`
	TZName    string `csv:"edge-timezone-name"`
	Speed     string `csv:"edge-conn-speed"`
}

type IPNetGeo struct {
	IPv4     net.IPNet
	Country2 string
	City     string
	LatLon   string
	TZName   string
	Speed    string
}

var inputReader io.Reader

func init() {
	if len(os.Args) != 2 {
		log.Fatal("one argument needed: the input fileIn")
	}
	fileIn := os.Args[1]
	fileOut := fileutil.CreateOutputFilename(fileIn, "mmdb")

	r, err := os.Open(fileIn)
	if err != nil {
		log.Fatal(err)
	}
	inputReader = r

	// TODO
	fmt.Println(fileOut)
	fmt.Println("---------------------------------")
}

func main() {

	csvReader := csv.NewReader(inputReader)
	csvDecoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	csvDecoder.Register(func(data []byte, ip *net.IP) error {
		*ip = net.ParseIP(string(data))
		return nil
	})

	var ipnets []IPNetGeo
	for {
		u := NA04Subset{}
		if err := csvDecoder.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if u.City == CityAsReservedIPNet {
			continue
		}

		ipnets = append(ipnets, IPNetGeo{
			IPv4:     *iputil.FromIPBeginAndEndToIPNet(u.IPv4Begin, u.IPv4End),
			Country2: u.Country2,
			City:     u.City,
			LatLon:   fmt.Sprintf("%s,%s", u.Latitude, u.Longitude),
			TZName:   u.TZName,
			Speed:    u.Speed,
		})
	}

	for _, e := range ipnets {
		fmt.Println(e.LatLon)
	}
}
