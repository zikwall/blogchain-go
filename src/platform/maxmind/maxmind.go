package maxmind

import (
	"github.com/oschwald/geoip2-golang"
	"net"
)

type Reader struct {
	reader *geoip2.Reader
}
type ReaderConfig struct {
	Path string
}
type ReaderResult struct {
	Country string
	Region  string
}

func CreateReader(conf ReaderConfig) (*Reader, error) {
	reader, err := openDatabase(conf.Path)

	if err != nil {
		return nil, err
	}

	return &Reader{reader: reader}, nil
}

func (rd *Reader) Lookup(ip string) (ReaderResult, error) {
	record, err := rd.reader.City(net.ParseIP(ip))
	result := ReaderResult{}

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

func (rd *Reader) Close() error {
	return rd.reader.Close()
}

func (rd Reader) CloseMessage() string {
	return "close Maxmind City database"
}

func openDatabase(path string) (*geoip2.Reader, error) {
	return geoip2.Open(path)
}
