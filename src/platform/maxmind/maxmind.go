package maxmind

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"net"
)

type (
	Finder struct {
		reader *geoip2.Reader
	}
	FinderConfig struct {
		Path string
	}
	FindResult struct {
		Country string
		Region  string
	}
)

func CreateFinder(c FinderConfig) (*Finder, error) {
	f := new(Finder)
	reader, err := f.openDatabase(c.Path)

	if err != nil {
		return nil, err
	}

	f.reader = reader

	fmt.Printf("Read MaxMind database from %s \n", c.Path)

	return f, nil
}

func (f *Finder) Lookup(ip string) (FindResult, error) {
	record, err := f.reader.City(net.ParseIP(ip))
	result := FindResult{}

	if err != nil {
		return result, err
	}

	result.Country = record.Country.IsoCode

	if len(record.Subdivisions) > 0 {
		// use last subdivision
		result.Region = record.Subdivisions[len(record.Subdivisions)-1].IsoCode
	}

	return result, nil
}

func (g Finder) openDatabase(mmdbPath string) (*geoip2.Reader, error) {
	return geoip2.Open(mmdbPath)
}

func (f *Finder) Close() error {
	return f.reader.Close()
}

func (f Finder) CloseMessage() string {
	return "Close Maxmind City database"
}
