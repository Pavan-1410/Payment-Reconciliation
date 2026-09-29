package handlers
import (
	"net/http"

	"github.com/gin-gonic/gin"

	"payment_reconciliation/dto"
	"payment_reconciliation/services"
)

type AuthHandler struct {
	AuthService *services.AuthService		// indirectly connected to DB
}

func (h *AuthHandler) Register(c *gin.Context) {

	var req dto.RegisterRequest

	// Read and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request data",
		})
		return
	}
	// Call service
	user, err := h.AuthService.Register(req)		// calling the method

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Send response
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user":    user,
	})
}

func (h *AuthHandler) Login(c *gin.Context){
	var req dto.LoginRequest

	err := c.ShouldBindJSON(&req)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid request data",
		})
		return
	}

	// call service
	user, err := h.AuthService.Login(req);

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//send responce

	c.JSON(http.StatusOK, gin.H{
		"message" : "User Login",
		"User" : user,
	})
}