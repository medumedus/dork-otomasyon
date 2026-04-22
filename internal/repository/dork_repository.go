package repository

import (
	"encoding/json"
	"os"

	"dork-otomasyon/internal/models"
)

type DorkRepository struct {
	FilePath string
}

func NewDorkRepository(path string) *DorkRepository {
	return &DorkRepository{
		FilePath: path,
	}
}

func (r *DorkRepository) GetAll() ([]models.Dork, error) {
	file, err := os.ReadFile(r.FilePath)
	if err != nil {
		return nil, err
	}

	var dorks []models.Dork

	err = json.Unmarshal(file, &dorks)
	if err != nil {
		return nil, err
	}

	return dorks, nil
}
