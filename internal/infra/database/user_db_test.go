package database

import (
	"testing"

	"github.com/sk8sta13/API-Service/internal/entity"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupDBUser(t *testing.T) (*gorm.DB, *User) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening database: %v", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		t.Fatalf("Error migrating the database: %v", err)
	}

	return db, NewUser(db)
}

func TestCreateUser(t *testing.T) {
	db, userDB := setupDBUser(t)

	user, _ := entity.NewUser("Marcelo Soto Campos", "marcelo@gmail.com", "123321")
	err := userDB.Create(user)
	assert.NoError(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, userDB := setupDBUser(t)

	user, _ := entity.NewUser("Marcelo Soto Campos", "sk8sta13@gmail.com", "123321")
	db.Create(user)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, userFound.ID, user.ID)
	assert.Equal(t, userFound.Name, user.Name)
	assert.Equal(t, userFound.Email, user.Email)
	assert.NotNil(t, userFound.Password)
}
