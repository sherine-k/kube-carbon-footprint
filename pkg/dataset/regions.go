package dataset

import (
	"os"

	"github.com/gocarina/gocsv"
)

const regFilename = "data/regions.csv"

type Region struct {
	Name    string  `csv:"Region"`
	Country string  `csv:"Country"`
	CO2e    float32 `csv:"CO2e (metric gram/kWh)"`
	PUE     float32 `csv:"PUE"`
}

func loadRegions() ([]*Region, error) {
	in, err := os.Open(regFilename)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	regions := []*Region{}

	if err := gocsv.UnmarshalFile(in, &regions); err != nil {
		return nil, err
	}

	return regions, nil
}

func (ds *Dataset) FindRegion(name string) *Region {
	for _, reg := range ds.regions {
		if reg.Name == name {
			return reg
		}
	}
	return nil
}
