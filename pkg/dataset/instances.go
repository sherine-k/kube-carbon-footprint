package dataset

import (
	"os"

	"github.com/gocarina/gocsv"
)

const filename = "data/instances.csv"

type Instance struct {
	Name     string  `csv:"Instance type"`
	LoadIdle float32 `csv:"Instance @ Idle"`
	Load10   float32 `csv:"Instance @ 10%"`
	Load50   float32 `csv:"Instance @ 50%"`
	Load100  float32 `csv:"Instance @ 100%"`
}

func loadInstances() ([]*Instance, error) {
	in, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	instances := []*Instance{}

	if err := gocsv.UnmarshalFile(in, &instances); err != nil {
		return nil, err
	}

	return instances, nil
}

func (ds *Dataset) FindInstance(name string) *Instance {
	for _, inst := range ds.instances {
		if inst.Name == name {
			return inst
		}
	}
	return nil
}
