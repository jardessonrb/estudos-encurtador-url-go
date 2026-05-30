package repository

import "encurtador-url-go/internal/domain"

type UrlRepository interface {
	Salvar(url *domain.URL) error
}
