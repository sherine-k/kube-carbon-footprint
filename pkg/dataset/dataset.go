package dataset

type Dataset struct {
	instances []*Instance
	regions   []*Region
}

func Load() (*Dataset, error) {
	inst, err := loadInstances()
	if err != nil {
		return nil, err
	}
	regs, err := loadRegions()
	if err != nil {
		return nil, err
	}
	return &Dataset{
		instances: inst,
		regions:   regs,
	}, nil
}
