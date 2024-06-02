package handler

import (
	"demo/internal/application/service"
	"demo/internal/infrastructure/cache"
	"demo/internal/infrastructure/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Test(c *gin.Context) {
	cache.MyCachePool.Set("test", "你好golang")
	cacheVal, _ := cache.MyCachePool.Get("test")
	log.Println("获取的key test ", cacheVal)
	c.JSON(http.StatusOK, nil)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	appConfig := utils.GetConfig(c)
	appConfigStr, _ := utils.JsonToString(appConfig)
	log.Println("全局配置: ", appConfigStr)
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
