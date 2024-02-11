package data

import "log"

type DefaultRepository struct {
	FileRepository
}

func NewDefaultRepository() *DefaultRepository {
	return &DefaultRepository{FileRepository: *NewFileRepository("")}
}

func (dr *DefaultRepository) Connect() error {
	dr.groups["default"] = *NewGroup()

	log.Printf("INFO: finished initializing defaultRepository")
	return nil
}
