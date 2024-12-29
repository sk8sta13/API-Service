package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/sk8sta13/API-Service/internal/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDBProduct(t *testing.T) (*gorm.DB, *Product) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}

	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		t.Fatalf("Error migrating the database: %v", err)
	}

	return db, NewProduct(db)
}

func TestFindAllProducts(t *testing.T) {
	db, ProductDB := setupDBProduct(t)

	for i := 1; i < 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product #%d", i), rand.Float64())
		assert.NoError(t, err)
		db.Create(product)
	}

	products, err := ProductDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product #1", products[0].Name)
	assert.Equal(t, "Product #10", products[9].Name)

	products, err = ProductDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product #11", products[0].Name)
	assert.Equal(t, "Product #20", products[9].Name)

	products, err = ProductDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product #21", products[0].Name)
	assert.Equal(t, "Product #24", products[3].Name)
}

func TestFindByIdProduct(t *testing.T) {
	_, ProductDB := setupDBProduct(t)

	product, _ := entity.NewProduct("Quebra Cabeça Marvel Vingadores 60 Peças Toyster", 21.49)

	_ = ProductDB.Create(product)

	NewProduct, err := ProductDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, NewProduct.Name)
	assert.Equal(t, product.Price, NewProduct.Price)
}

func TestCreateNewProduct(t *testing.T) {
	_, ProductDB := setupDBProduct(t)

	product, _ := entity.NewProduct("Raspbarry Pi", 395.00)
	err := ProductDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestUpdateProduct(t *testing.T) {
	db, ProductDB := setupDBProduct(t)

	product, _ := entity.NewProduct("Raspbarry Pi", 395.00)
	db.Create(product)

	product.Name = "Livro: Rápido e devagar: Duas formas de pensar"
	product.Price = 58.67
	err := ProductDB.Update(product)
	assert.NoError(t, err)

	var UpdatedProduct entity.Product
	db.First(&UpdatedProduct, "id = ?", product.ID.String())
	assert.Equal(t, UpdatedProduct.Name, product.Name)
	assert.Equal(t, UpdatedProduct.Price, product.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, ProductDB := setupDBProduct(t)

	product, _ := entity.NewProduct("Livro: Laranja Mecânica", 43.16)
	db.Create(product)

	err := ProductDB.Delete(product.ID.String())
	assert.NoError(t, err)

	var DeletedProduct entity.Product
	err = db.First(&DeletedProduct, "id = ?", product.ID.String()).Error
	assert.Error(t, err)
}
