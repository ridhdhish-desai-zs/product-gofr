package service

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/ridhdhish-desai-zs/product-gofr/models"
)

type Product interface {
	GetById(ctx *gofr.Context, id string) (*models.Product, error)
}
