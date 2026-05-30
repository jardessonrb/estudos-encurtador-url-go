package postgres

import (
	"context"
	"encurtador-url-go/internal/database"
	"encurtador-url-go/internal/domain"

	"gorm.io/gorm"
)

type UrlRepository struct{}

func NewUrlRepository() *UrlRepository {
	return &UrlRepository{}
}

func (r *UrlRepository) Salvar(url *domain.URL) error {
	return database.DB.Create(url).Error
}

func (r *UrlRepository) BuscarPorCodigo(codigo string) (*domain.URL, error) {
	url, err := gorm.G[domain.URL](database.DB).Where("codigo = ?", codigo).First(context.Background())

	if err != nil {
		return nil, err
	}

	return &url, nil
}
