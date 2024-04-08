package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/personal/user-manager-backend/app/dto"
	"github.com/personal/user-manager-backend/mocks/app/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createControllerAndServices(t *testing.T, mockCtrl *gomock.Controller) (*UserController, *service.MockUserService) {
	t.Helper()

	userService := service.NewMockUserService(mockCtrl)
	userController := NewUserController(userService)

	return userController, userService
}

func TestUserController_GetUserSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	gin.SetMode(gin.TestMode)

	userController, userService := createControllerAndServices(t, mockCtrl)

	userService.EXPECT().GetUser(gomock.Any(), 1).Return(&dto.UserDto{}, nil)

	r := gin.Default()
	r.GET("/api/v1/users/:id", userController.GetUser)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/users/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_GetUsersSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	gin.SetMode(gin.TestMode)

	userController, userService := createControllerAndServices(t, mockCtrl)

	userService.EXPECT().GetUsers(gomock.Any(), 1, 10).Return(&dto.UserDtoPaginated{}, nil)

	r := gin.Default()
	r.GET("/api/v1/users", userController.GetUsers)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserController_GetUsersInternalServerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	gin.SetMode(gin.TestMode)

	userController, userService := createControllerAndServices(t, mockCtrl)

	userService.EXPECT().GetUsers(gomock.Any(), 1, 10).Return(nil, errors.New("failed to fetch users"))

	r := gin.Default()
	r.GET("/api/v1/users", userController.GetUsers)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
