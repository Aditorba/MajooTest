package service

import (
	"errors"
	"majooTest/config"
	"majooTest/domain"
	"majooTest/dto"
	"majooTest/log"
	"majooTest/repository"
	"majooTest/util"
	"strings"
	"time"
)

type AuthService interface {
	Login(loginDTO dto.LoginDTO) (bool, error)
	GenerateToken(username string) (dto.LoginDTO, error)
	ValidateToken(tokenRequest string) (bool, dto.Users, string)
}

type authService struct {
	configData     *config.ConfigData
	userRepository repository.UsersRepository
}

func NewAuthService(
	configData *config.ConfigData,
	userRepository repository.UsersRepository,
) AuthService {
	return &authService{
		configData:     configData,
		userRepository: userRepository,
	}
}

func (service *authService) Login(loginDTO dto.LoginDTO) (bool, error) {
	log.Info("Login request data : ", loginDTO)
	userData, err := service.userRepository.GetDetailByUsername(loginDTO.Username)
	log.Info("Users data on DB : ", userData)
	if err != nil {
		log.Error("Error get user data from DB", err)
		return false, err
	}

	/* TO DO
		Seharusnya disini bisa ditambahkan logic untuk penambahan SALT password. disini saya tidak menambahkan
	karena saya mengikuti contoh data yang sudah diberikan yaitu hanya md5 sederhana saja. Contoh data password
	yang disimpan di database tidak hanya hasil md5 melainkan ditambahkan prefix SALT tertentu ex: md5(md5(value)+"PREFIX")).
	*/
	if loginDTO.Password != userData.Password {
		return false, errors.New("Invalid password")
	}

	return true, nil
}

func (service *authService) GenerateToken(username string) (dto.LoginDTO, error) {
	log.Info("username request : ", username)
	var userDataDTO dto.LoginDTO

	userDataFromDB, err := service.userRepository.GetDetailByUsername(username)
	log.Info("userDataFromDB : ", userDataFromDB)
	if err != nil {
		return userDataDTO, err
	}

	generatedToken, err := util.GenerateToken(userDataFromDB.UserName, service.configData.Secret.KeyGenerate)
	if err != nil {
		return userDataDTO, err
	}

	userDataDTO.Username = userDataFromDB.UserName
	userDataDTO.Password = util.MaskingString(2, 2, userDataFromDB.Password)
	userDataDTO.Token = generatedToken

	userDataFromDB.UpdateAt = time.Now()
	userDataFromDB.Token = generatedToken

	savedUserData, err := service.userRepository.SaveUser(userDataFromDB)
	if err != nil {
		log.Error("error save generate token", err)
		return userDataDTO, err
	}
	log.Info("response save generate token", savedUserData)

	return userDataDTO, nil
}

func (service *authService) ValidateToken(tokenRequest string) (bool, dto.Users, string) {
	splittedAccessKey := strings.Split(tokenRequest, " ")
	bearerString := splittedAccessKey[0]
	token := splittedAccessKey[1]
	var userDTO = dto.Users{}
	log.Info("bearer string ", bearerString)
	log.Info("jwt token request ", token)

	if bearerString != util.Bearer {
		err := errors.New("Invalid Token")
		log.Error("Invalid Token ", err)
		return false, userDTO, err.Error()
	}

	isValid, claims, message := util.ValidateToken(token, service.configData.Secret.KeyGenerate)
	userDataOnDB, err := service.userRepository.GetDetailByUsername(claims.Username)
	if err != nil {
		log.Error("Error while get data user on DB ", err)
		return false, userDTO, err.Error()
	}
	log.Info("user data from DB ", userDataOnDB)

	if isValid {
		if token != userDataOnDB.Token {
			err = errors.New("Invalid Token")
			log.Error("Invalid Token ", err)

			return false, userDTO, err.Error()
		}
	} else {
		err = errors.New("Invalid Token")
		log.Error("Invalid Token ", err)

		return false, userDTO, err.Error()
	}

	return isValid, MappingUsersDomainToDTO(userDataOnDB, userDTO), message
}

func MappingUsersDTOtoDomain(userDTO dto.Users, userDomain domain.Users) domain.Users {

	if userDTO.Id != 0 {
		userDomain.Id = userDTO.Id
	}

	if userDTO.UserName != "" {
		userDomain.UserName = userDTO.UserName
	}

	if userDTO.Password != "" {
		userDomain.Password = userDTO.Password
	}

	if userDTO.Name != "" {
		userDomain.Name = userDTO.Name
	}

	if userDTO.Token != "" {
		userDomain.Token = userDTO.Token
	}

	return userDomain
}

func MappingUsersDomainToDTO(userDomain domain.Users, userDTO dto.Users) dto.Users {

	if userDomain.Id != 0 {
		userDTO.Id = userDomain.Id
	}

	if userDomain.UserName != "" {
		userDTO.UserName = userDomain.UserName
	}

	if userDomain.Password != "" {
		userDTO.Password = userDomain.Password
	}

	if userDomain.Name != "" {
		userDTO.Name = userDomain.Name
	}

	if userDomain.Token != "" {
		userDTO.Token = userDomain.Token
	}

	return userDTO
}
