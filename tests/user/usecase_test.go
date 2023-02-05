package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	bcryptmocks "github.com/Risuii/config/bcrypt/mocks"
	"github.com/Risuii/helpers/exception"
	"github.com/Risuii/internal/user"
	"github.com/Risuii/models/token"
	"github.com/Risuii/models/users"
	"github.com/Risuii/tests/user/mocks"
)

func TestRegister(t *testing.T) {
	t.Run("Register Success", func(t *testing.T) {
		employeeRepository := new(mocks.UserRepository)
		bcrypt := new(bcryptmocks.Bcrypt)

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Employee{}, exception.ErrNotFound)
		employeeRepository.On("Create", mock.Anything, mock.AnythingOfType("users.Employee")).Return(int64(1), nil)
		bcrypt.On("HashPassword", mock.AnythingOfType("string")).Return("hashed password", nil)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()

		params := users.Employee{
			ID:       1,
			Name:     "test",
			Password: "test",
			Email:    "test@test.com",
		}

		resp := employeeUseCase.Register(ctx, params)

		assert.NoError(t, resp.Err())

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})

	t.Run("Register Error Conflict", func(t *testing.T) {
		employeeRepository := new(mocks.UserRepository)
		bcrypt := new(bcryptmocks.Bcrypt)

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Employee{}, nil)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()

		params := users.Employee{
			ID:       1,
			Name:     "test",
			Password: "test",
			Email:    "test@test.com",
		}

		resp := employeeUseCase.Register(ctx, params)

		assert.Error(t, resp.Err())

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})

	t.Run("Register Error Internal Server Bcrypt", func(t *testing.T) {
		employeeRepository := new(mocks.UserRepository)
		bcrypt := new(bcryptmocks.Bcrypt)

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Employee{}, exception.ErrNotFound)
		bcrypt.On("HashPassword", mock.AnythingOfType("string")).Return("hashed password", exception.ErrInternalServer)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()

		params := users.Employee{
			ID:       1,
			Name:     "test",
			Password: "test",
			Email:    "test@test.com",
		}

		resp := employeeUseCase.Register(ctx, params)

		assert.Error(t, resp.Err())

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})

	t.Run("Register Erorr Internal Server", func(t *testing.T) {
		employeeRepository := new(mocks.UserRepository)
		bcrypt := new(bcryptmocks.Bcrypt)

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Employee{}, exception.ErrNotFound)
		employeeRepository.On("Create", mock.Anything, mock.AnythingOfType("users.Employee")).Return(int64(0), exception.ErrInternalServer)
		bcrypt.On("HashPassword", mock.AnythingOfType("string")).Return("hashed password", nil)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()

		params := users.Employee{
			ID:       1,
			Name:     "test",
			Password: "test",
			Email:    "test@test.com",
		}

		resp := employeeUseCase.Register(ctx, params)

		assert.Error(t, resp.Err())

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login Success", func(t *testing.T) {
		bcrypt := new(bcryptmocks.Bcrypt)
		employeeRepository := new(mocks.UserRepository)

		password := "hashed"

		mockAccount := users.Employee{
			Password: password,
		}

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockAccount, nil)
		bcrypt.On("ComparePasswordHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()
		params := users.EmployeeLogin{
			Email: "test@test.com",
		}

		resp, tokens := employeeUseCase.Login(ctx, params)
		tokens = token.Token{
			Token: "jwt-token-test",
		}

		assert.NoError(t, resp.Err())
		assert.NotEmpty(t, tokens)

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})

	t.Run("Login Error Not Found", func(t *testing.T) {
		bcrypt := new(bcryptmocks.Bcrypt)
		employeeRepository := new(mocks.UserRepository)

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Employee{}, exception.ErrNotFound)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()
		params := users.EmployeeLogin{
			Email: "test@test.com",
		}

		resp, _ := employeeUseCase.Login(ctx, params)

		assert.Error(t, resp.Err())

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})

	t.Run("Login Error Internal Server", func(t *testing.T) {
		bcrypt := new(bcryptmocks.Bcrypt)
		employeeRepository := new(mocks.UserRepository)

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(users.Employee{}, exception.ErrInternalServer)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()
		params := users.EmployeeLogin{
			Email: "test@test.com",
		}

		resp, _ := employeeUseCase.Login(ctx, params)

		assert.Error(t, resp.Err())

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})

	t.Run("Login Error Password Not Valid", func(t *testing.T) {
		bcrypt := new(bcryptmocks.Bcrypt)
		employeeRepository := new(mocks.UserRepository)

		password := "hashed"

		mockAccount := users.Employee{
			Password: password,
		}

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockAccount, nil)
		bcrypt.On("ComparePasswordHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(false)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()
		params := users.EmployeeLogin{
			Email: "test@test.com",
		}

		resp, _ := employeeUseCase.Login(ctx, params)

		assert.NoError(t, resp.Err())

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})

	t.Run("Login Error Token Empty", func(t *testing.T) {
		bcrypt := new(bcryptmocks.Bcrypt)
		employeeRepository := new(mocks.UserRepository)

		password := "hashed"

		mockAccount := users.Employee{
			Password: password,
		}

		employeeRepository.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(mockAccount, nil)
		bcrypt.On("ComparePasswordHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(true)

		employeeUseCase := user.NewUserUseCase(
			employeeRepository,
			bcrypt,
		)

		ctx := context.TODO()
		params := users.EmployeeLogin{
			Email: "test@test.com",
		}

		resp, tokens := employeeUseCase.Login(ctx, params)
		tokens = token.Token{
			Token: "",
		}

		assert.NoError(t, resp.Err())
		assert.Empty(t, tokens)

		employeeRepository.AssertExpectations(t)
		bcrypt.AssertExpectations(t)
	})
}
