package maxmind

import (
	"net"

	"github.com/oschwald/geoip2-golang"
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

func CreateReader(cfg ReaderConfig) (*Reader, error) {
	reader, err := geoip2.Open(cfg.Path)
	if err != nil {
		return nil, err
	}
	return &Reader{reader: reader}, nil
}

func (r *Reader) Lookup(ip string) (ReaderResult, error) {
	record, err := r.reader.City(net.ParseIP(ip))
	result := ReaderResult{}
	if err != nil {
		return result, err
	}
	result.Country = record.Country.IsoCode
	if len(record.Subdivisions) > 0 {
		result.Region = record.Subdivisions[len(record.Subdivisions)-1].IsoCode
	}
	return result, nil
}

func (r *Reader) Drop() error {
	return r.reader.Close()
}

func (r *Reader) DropMsg() string {
	return "close geo reader"
}
