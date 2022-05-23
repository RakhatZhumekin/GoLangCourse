package service

import (
	"log"

	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/dto"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/entity"
	"github.com/RakhatZhumekin/GoLangCourse/tree/main/Final/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService {
		userRepository,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}

	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}