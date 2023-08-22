package service_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rzldimam28/auth-grpc/src/entity"
	"github.com/rzldimam28/auth-grpc/src/model"
	mockRepository "github.com/rzldimam28/auth-grpc/src/repository/mocks"
	"github.com/rzldimam28/auth-grpc/src/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func newMockDB(t *testing.T) (sqlmock.Sqlmock, *sql.DB) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("failed create new mock db: %s", err.Error())
	}
	return mock, db
}

func TestLogin(t *testing.T) {

	m, mockDB := newMockDB(t)
	mockValidate := validator.New()
	mockRepo := new(mockRepository.Repository)

	t.Run("Success Login", func(t *testing.T) {
		initPass := "dummy-password"
		hashPass, err := bcrypt.GenerateFromPassword([]byte(initPass), 14)
		hashPassStr := string(hashPass)

		assert.Nil(t, err)

		ret := entity.User{
			ID: "dummy-uuid",
			Username: "dummy-username",
			Email: "dummy-email@gmail.com",
			Password: hashPassStr,
		}

		m.ExpectBegin()

		mockRepo.On("FindByEmail", mock.Anything, mock.Anything, ret.Email).Return(&ret, nil).Once()

		service := service.New(mockDB, mockValidate, mockRepo)

		req := model.LoginRequest{
			Email: ret.Email,
			Password: initPass,
		}
		expectedRes := &model.UserResponse{
			ID: ret.ID,
			Username: ret.Username,
			Email: ret.Email,
			Password: ret.Password,
		}

		userResp, err := service.Login(context.Background(), req)

		assert.Nil(t, err)
		assert.Equal(t, expectedRes, userResp)
	})

	t.Run("Failed Login | Wrong Password", func(t *testing.T) {
		initPass := "dummy-password"
		hashPass, err := bcrypt.GenerateFromPassword([]byte(initPass), 14)
		hashPassStr := string(hashPass)

		assert.Nil(t, err)

		ret := entity.User{
			ID: "dummy-uuid",
			Username: "dummy-username",
			Email: "dummy-email@gmail.com",
			Password: hashPassStr,
		}

		m.ExpectBegin()

		mockRepo.On("FindByEmail", mock.Anything, mock.Anything, ret.Email).Return(&ret, nil).Once()

		service := service.New(mockDB, mockValidate, mockRepo)

		req := model.LoginRequest{
			Email: ret.Email,
			Password: "wrongpassword",
		}

		userResp, err := service.Login(context.Background(), req)

		assert.Nil(t, userResp)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("Failed Login | No Username Found", func(t *testing.T) {
		m.ExpectBegin()

		mockRepo.On("FindByEmail", mock.Anything, mock.Anything, "wrongemail@gmail.com").Return(nil, errors.New("email not found")).Once()

		service := service.New(mockDB, mockValidate, mockRepo)

		req := model.LoginRequest{
			Email: "wrongemail@gmail.com",
			Password: "randompassword",
		}

		userResp, err := service.Login(context.Background(), req)

		assert.Nil(t, userResp)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

	t.Run("Failed Login | Validate Error", func(t *testing.T) {
		initPass := "dummy-password"
		hashPass, err := bcrypt.GenerateFromPassword([]byte(initPass), 14)
		hashPassStr := string(hashPass)

		assert.Nil(t, err)

		ret := entity.User{
			ID: "dummy-uuid",
			Username: "dummy-username",
			Email: "dummy-email@gmail.com",
			Password: hashPassStr,
		}

		m.ExpectBegin()

		mockRepo.On("FindByEmail", mock.Anything, mock.Anything, ret.Email).Return(&ret, nil).Once()

		service := service.New(mockDB, mockValidate, mockRepo)

		req := model.LoginRequest{
			Email: "unvalidateemailaddress",
			Password: "",
		}

		userResp, err := service.Login(context.Background(), req)

		assert.Nil(t, userResp)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})
	
}

func TestRegister(t *testing.T) {

	m, mockDB := newMockDB(t)
	mockValidate := validator.New()
	mockRepo := new(mockRepository.Repository)

	t.Run("Success Register", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "dummy-username",
			Email: "dummy-email@gmail.com",
			Password: "dummy-password",
		}
	
		hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		hashPassStr := string(hashPass)
		assert.Nil(t, err)
	
		id := uuid.NewString()
	
		userToCreate := entity.User{
			ID: id,
			Username: req.Username,
			Email: req.Email,
			Password: hashPassStr,
		}
	
		m.ExpectBegin()
	
		mockRepo.On("InsertUser", mock.Anything, mock.Anything, mock.AnythingOfType("entity.User")).Return(&userToCreate, nil).Once()
	
		service := service.New(mockDB, mockValidate, mockRepo)
	
		expectedRes := &model.UserResponse{
			ID: userToCreate.ID,
			Username: userToCreate.Username,
			Email: userToCreate.Email,
			Password: userToCreate.Password,
		}
	
		userResp, err := service.Register(context.Background(), req)
	
		assert.Nil(t, err)
		assert.Equal(t, expectedRes, userResp)
	})

	t.Run("Failed Register | Validate Error", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "",
			Email: "wrong-email#email",
			Password: "wrpswd",
		}
	
		hashPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		hashPassStr := string(hashPass)
		assert.Nil(t, err)
	
		id := uuid.NewString()
	
		userToCreate := entity.User{
			ID: id,
			Username: req.Username,
			Email: req.Email,
			Password: hashPassStr,
		}
	
		m.ExpectBegin()
	
		mockRepo.On("InsertUser", mock.Anything, mock.Anything, mock.AnythingOfType("entity.User")).Return(&userToCreate, nil).Once()
	
		service := service.New(mockDB, mockValidate, mockRepo)
	
		userResp, err := service.Register(context.Background(), req)
	
		assert.Nil(t, userResp)
		assert.NotNil(t, err)
		assert.Error(t, err)
	})

}