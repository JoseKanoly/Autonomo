package services

import (
	"go_project/internal/models"
	"go_project/internal/repository"
)

type ListaPreciosService struct {
	repo *repository.ListaPreciosRepository
}

func NewListaPreciosService(repo *repository.ListaPreciosRepository) *ListaPreciosService {
	return &ListaPreciosService{repo: repo}
}

func (s *ListaPreciosService) GetAll() ([]models.ListaPrecios, error) {
	return s.repo.GetAll()
}

func (s *ListaPreciosService) GetByID(id int) (*models.ListaPrecios, error) {
	return s.repo.GetByID(id)
}

func (s *ListaPreciosService) Create(lp *models.ListaPrecios) error {
	return s.repo.Create(lp)
}

func (s *ListaPreciosService) Update(lp *models.ListaPrecios) error {
	return s.repo.Update(lp)
}

func (s *ListaPreciosService) Delete(id int) error {
	return s.repo.Delete(id)
}
