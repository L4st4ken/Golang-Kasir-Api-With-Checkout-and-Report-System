package services

import (
	"kasir-api/internal/models"
	"kasir-api/internal/repositories"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetTodaySummary() (*models.SalesSummary, error) {
	return s.repo.GetTodaySummary()
}
