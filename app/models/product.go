package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID         string `gorm:"size:36;index"`
	User             User
	UserID           string `gorm:"size:36;index"`
	ProductImages    []ProductImage
	Categories       []Category      `gorm:"many2many:product_categories;"`
	Sku              string          `gorm:"size:100;index"`
	Name             string          `gorm:"size:255"`
	Slug             string          `gorm:"size:255"`
	Price            decimal.Decimal `gorm:"type:decimal(16,2);"`
	Stock            int
	Weight           decimal.Decimal `gorm:"type:decimal(10,2);"`
	ShortDescription string          `gorm:"type:text"`
	Description      string          `gorm:"type:text"`
	Status           int             `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (p *Product) GetProducts(db *gorm.DB, perPage int, page int) (*[]Product, int64, error) {
	var err error
	var products []Product
	var count int64

	// query mengambil data produk
	err = db.Debug().Model(&Product{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offSet := (page - 1) * perPage

	// query menampilkan data produk
	err = db.Debug().Model(&Product{}).Order("Created_at desc").Limit(perPage).Offset(offSet).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return &products, count, nil
}

func (p *Product) FindBySlug(db *gorm.DB, slug string) (*Product, error) {
	var err error
	var product Product

	err = db.Debug().Preload("ProductImages").Model(&Product{}).Where("slug = ?", slug).First(&product).Error

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) FindByID(db *gorm.DB, productID string) (*Product, error) {
	var err error
	var product Product

	err = db.Debug().Preload("ProductImages").Model(&Product{}).Where("id = ?", productID).First(&product).Error

	if err != nil {
		return nil, err
	}
	return &product, nil
}


