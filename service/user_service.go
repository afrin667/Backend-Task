package service

import "task/entity"

type UserService interface {
	Save(entity.User) entity.User
	FindAll() []entity.User
}

type userService struct {
	users  []entity.User
	userId int
}

func NewUserService() UserService {
	return &userService{
		userId: 1,
	}
}

func (s *userService) Save(user entity.User) entity.User {
	newUser := &entity.User{
		User_Id:  s.userId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	s.users = append(s.users, *newUser)
	s.userId++
	return user
}

func (s *userService) FindAll() []entity.User {
	return s.users
}
