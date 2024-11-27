package service

import (
	"dto/dao"
	"dto/dto"
	"dto/models"
)

type UserService struct {
	UserDAO *dao.UserDAO
}

func NewUserService(UserDAO *dao.UserDAO) *UserService {
	return &UserService{UserDAO: UserDAO}
}

func (s *UserService) Create(userRequest *dto.UserRequest) error {
	user := &models.User{
		Username: userRequest.Username,
		Password: userRequest.Password,
	}

	return s.UserDAO.Create(user)
}

func (s *UserService) GetAll() ([]dto.UserResponse, error) {
	users, err := s.UserDAO.GetAll()
	if err != nil {
		return nil, err
	}

	var userResponses []dto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dto.UserResponse{
			ID:       user.ID,
			Username: user.Username,
		})
	}
	return userResponses, nil
}

func (s *UserService) GetByID(id int) (*dto.UserResponse, error) {
	user, err := s.UserDAO.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *UserService) DeleteByID(id int) error {
	return s.UserDAO.Delete(id)
}
