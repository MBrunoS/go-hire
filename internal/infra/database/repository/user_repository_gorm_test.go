package repository_test

import (
	"testing"

	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/internal/infra/database/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	repo, user, db := setupUserRepo()

	err := repo.Create(user)
	assert.Nil(t, err)

	var u entity.User
	err = db.First(&u, user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, u.ID)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.NotEmpty(t, u.Password)
}

func TestFindUserByID(t *testing.T) {
	repo, user, _ := setupUserRepo()

	err := repo.Create(user)
	assert.Nil(t, err)

	u, err := repo.FindByID(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, u.ID)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.NotEmpty(t, u.Password)
}

func TestFindByEmail(t *testing.T) {
	repo, user, _ := setupUserRepo()

	err := repo.Create(user)
	assert.Nil(t, err)

	u, err := repo.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, u.ID)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.NotEmpty(t, u.Password)
}

func TestUpdateUser(t *testing.T) {
	repo, user, db := setupUserRepo()

	err := repo.Create(user)
	assert.Nil(t, err)

	user.Name = "New Name"
	user.Email = "john2@doe.com"
	user.Password = "newpassword"

	err = repo.Update(user)
	assert.Nil(t, err)

	var u entity.User
	err = db.First(&u, user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, u.ID)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Email, u.Email)
	assert.NotEmpty(t, u.Password)
}

func TestDeleteUser(t *testing.T) {
	repo, user, db := setupUserRepo()

	err := repo.Create(user)
	assert.Nil(t, err)

	err = repo.Delete(user.ID)
	assert.Nil(t, err)

	var u entity.User
	err = db.First(&u, user.ID).Error
	assert.NotNil(t, err)
}

func setupUserRepo() (*repository.UserRepository, *entity.User, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john@doe.com", "password")
	repo := repository.NewUserRepository(db)

	return repo, user, db
}
