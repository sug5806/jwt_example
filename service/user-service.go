package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"jwt-example/model"
	"jwt-example/model/db_model"
	"jwt-example/utils/db"
	"jwt-example/utils/errors"
	"net/http"
	"time"
)

var (
	UserService userServiceInterface
)

type userService struct {
	DB *gorm.DB
}

type userServiceInterface interface {
	SignUp(model model.SignUpRequest) (*model.SignUpResponse, *errors.ApiError)
	UserCheck(userId string) *errors.ApiError
}

func init() {
	database, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	UserService = &userService{
		DB: database,
	}
}

func (u userService) SignUp(requestModel model.SignUpRequest) (*model.SignUpResponse, *errors.ApiError) {
	dbUserModel := db_model.User{
		Email: requestModel.Email,
		Pass:  requestModel.Password,
		Uuid:  uuid.New().String(),
	}

	jwtToken, err := u.CreateToken(dbUserModel.Uuid)
	if err != nil {
		return nil, &errors.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "token create fail",
		}
	}

	dbResp := u.DB.Table(dbUserModel.TableName()).Where(dbUserModel).FirstOrCreate(&dbUserModel)
	if dbResp.Error != nil {
		return nil, &errors.ApiError{
			Code:    http.StatusInternalServerError,
			Message: "user create Error",
		}
	}

	return &model.SignUpResponse{
		JwtToken: jwtToken,
	}, nil
}

func (u userService) CreateToken(userId string) (string, error) {
	claims := &model.Claims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(time.Minute * 60).Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte("jwt secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u userService) UserCheck(userId string) *errors.ApiError {
	var dbUserModel db_model.User

	dbResp := u.DB.Table(dbUserModel.TableName()).
		Where("uuid = ?", userId).
		First(&dbUserModel)

	if dbResp.Error != nil {
		return &errors.ApiError{
			Code:    http.StatusInternalServerError,
			Message: dbResp.Error.Error(),
		}
	}

	return nil
}
