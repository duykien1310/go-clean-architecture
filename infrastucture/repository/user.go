package repository

import (
	"auth_service/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := entity.User{}

	result := r.db.Where("email = ?", email).
		First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}

	return &user, nil
}

func (r UserRepository) FindById(id int) (*entity.User, error) {
	user := entity.User{}

	result := r.db.Where("id = ?", id).
		Preload("Post").
		Preload("FollowingUser").
		Preload("FollowerUser").
		First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, result.Error
		}
	}

	return &user, nil
}

func (r UserRepository) VerifyUserNameExist(userName string) (bool, error) {
	user := entity.User{}
	err := r.db.Model(&entity.User{}).
		Where("user_name = ?", userName).
		Find(&user).Error
	if err != nil || user.Id != 0 {
		return true, err
	}

	return false, nil
}

func (r UserRepository) Register(user *entity.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r UserRepository) GetListFolowingUserIdByUserId(userId int) ([]int, error) {
	var userIds []int

	err := r.db.Model(&entity.User{}).
		Joins(`join following on "user"."id" = following.user_id`).
		Where(`"user"."id" = ?`, userId).
		Pluck(`following.follow_user_id`, &userIds).Error

	if err != nil {
		return nil, err
	}
	return userIds, nil
}

func (r UserRepository) GetListFolowerIdByUserId(userId int) ([]int, error) {
	var userIds []int

	err := r.db.Model(&entity.User{}).
		Joins(`join following on "user"."id" = following.follow_user_id`).
		Where(`"user"."id" = ?`, userId).
		Pluck(`following.user_id`, &userIds).Error

	if err != nil {
		return nil, err
	}
	return userIds, nil
}
