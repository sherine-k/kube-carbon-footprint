package dataset

type Dataset struct {
	instances []*Instance
}

func Load() (*Dataset, error) {
	inst, err := loadInstances()
	if err != nil {
		return nil, err
	}
	return &Dataset{
		instances: inst,
	}, nil
}
