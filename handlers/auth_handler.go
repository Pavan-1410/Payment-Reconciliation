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

// Welcome godoc
// @Summary Welcome to Payment Reconciliation API
// @Description Returns a welcome message for the API
// @Tags General
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/auth/ [get]
func Welcome(c *gin.Context){
	c.JSON(http.StatusOK, "Welcome to Reconciliation API")
}


// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Registration details"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/auth/register [post]
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

// Login godoc
// @Summary Login user
// @Description Authenticates a user and returns a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /api/auth/login [post]
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