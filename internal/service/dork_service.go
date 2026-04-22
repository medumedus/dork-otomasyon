package service

import (
	"net/url"
	"strings"

	"dork-otomasyon/internal/models"
	"dork-otomasyon/internal/repository"
)

type DorkService struct {
	repo *repository.DorkRepository
}

func NewDorkService(r *repository.DorkRepository) *DorkService {
	return &DorkService{
		repo: r,
	}
}

func (s *DorkService) Generate(target string, category string) ([]models.GeneratedDork, error) {

	dorks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var results []models.GeneratedDork

	for _, d := range dorks {

		if category != "" && category != "all" && d.Category != category {
			continue
		}
		query := strings.ReplaceAll(d.Template, "{target}", target)
		googleURL := "https://www.google.com/search?q=" + url.QueryEscape(query)

		result := models.GeneratedDork{
			ID:        d.ID,
			Category:  d.Category,
			Title:     d.Title,
			Desc:      d.Desc,
			Query:     query,
			GoogleURL: googleURL,
		}

		results = append(results, result)
	}

	return results, nil
}
