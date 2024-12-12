package services

import (
	"go_project/internal/models"
	"go_project/internal/repository"
)

type ResiduosService struct {
	repo *repository.ResiduosRepository
}

func NewResiduosService(repo *repository.ResiduosRepository) *ResiduosService {
	return &ResiduosService{repo: repo}
}

func (s *ResiduosService) GetAll() ([]models.Residuo, error) {
	return s.repo.GetAll()
}

func (s *ResiduosService) GetByID(id int) (*models.Residuo, error) {
	return s.repo.GetByID(id)
}

func (s *ResiduosService) Create(residuo *models.Residuo) error {
	return s.repo.Create(residuo)
}

func (s *ResiduosService) Update(residuo *models.Residuo) error {
	return s.repo.Update(residuo)
}

func (s *ResiduosService) Delete(id int) error {
	return s.repo.Delete(id)
}
