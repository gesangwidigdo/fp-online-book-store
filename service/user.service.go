package service

import (
	"github.com/Djuanzz/go-template/dto"
	"github.com/Djuanzz/go-template/model"
	"github.com/Djuanzz/go-template/repository"
	"github.com/Djuanzz/go-template/utils"
)

type UserService interface {
	Register(userReq dto.UserRegisterReq) (dto.UserRegisterRes, error)
	Login(userReq dto.UserLoginReq) (dto.UserLoginRes, error)
	Me(id string) (dto.UserMeRes, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (us *userService) Register(userReq dto.UserRegisterReq) (dto.UserRegisterRes, error) {
	isUsername, err := us.userRepo.IsUsernameExist(userReq.Name)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}
	if isUsername {
		return dto.UserRegisterRes{}, dto.ERR_USERNAME_ALREADY_EXISTS
	}

	isEmail, err := us.userRepo.IsEmailExist(userReq.Email)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	if isEmail {
		return dto.UserRegisterRes{}, dto.ERR_EMAIL_ALREADY_EXISTS
	}

	password, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	user := model.User{
		Email:    userReq.Email,
		Name: userReq.Name,
		Password: password,
		Address: userReq.Address,
		Gender: userReq.Gender,
		PhoneNumber: userReq.PhoneNumber,
		BirthDate: userReq.BirthDate,
	}

	usr, err := us.userRepo.Register(user)
	if err != nil {
		return dto.UserRegisterRes{}, err
	}

	return dto.UserRegisterRes{
		Email:    usr.Email,
		Name: usr.Name,
	}, nil
}

func (us *userService) Login(userReq dto.UserLoginReq) (dto.UserLoginRes, error) {
	// fmt.Println(userReq)
	// --- CARI EMAIL
	isUser, err := us.userRepo.FindUserByEmmail(userReq.Email)
	if err != nil {
		return dto.UserLoginRes{}, dto.ERR_LOGIN
	}
	if isUser.Email == "" {
		return dto.UserLoginRes{}, dto.ERR_LOGIN
	}

	// --- KALO NEMU LANJUT CEK PASSWORD
	passHash, _ := utils.HashPassword(isUser.Password)
	isPasswordMatch := utils.CheckPasswordHash(isUser.Password, passHash)
	if !isPasswordMatch {
		return dto.UserLoginRes{}, dto.ERR_LOGIN
	}

	// --- KALO BENER GENERATE TOKEN
	token, err := utils.GenerateToken(isUser.ID)
	if err != nil {
		return dto.UserLoginRes{}, err
	}

	user := dto.UserLoginRes{
		Email: isUser.Email,
		Token: token,
	}

	return user, nil
}

func (us *userService) Me(id string) (dto.UserMeRes, error) {
	user, err := us.userRepo.FindUserById(id)
	if err != nil {
		return dto.UserMeRes{}, dto.ERR_USER_NOT_FOUND
	}

	return dto.UserMeRes{
		Email:    user.Email,
		Name: user.Name,
	}, nil
}
