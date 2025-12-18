package controllers

import (
	"database/sql"
	"net/http"

	"github.com/HenryKristofani/GoFutsal/auth"
	"github.com/HenryKristofani/GoFutsal/config"
	"github.com/HenryKristofani/GoFutsal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest represents the login credentials
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	Success      bool        `json:"success" example:"true"`
	Message      string      `json:"message" example:"Login successful"`
	AccessToken  string      `json:"access_token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	RefreshToken string      `json:"refresh_token,omitempty" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User         models.User `json:"user"`
}

// Login godoc
// @Summary      Login user
// @Description  Login dengan username dan password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        credentials  body  LoginRequest  true  "Login Credentials"
// @Success      200  {object}  LoginResponse
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/auth/login [post]
func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username dan password harus diisi"})
		return
	}

	// Query user dari database berdasarkan username
	var user models.User
	var hashedPassword string

	query := `SELECT id, username, email, password, role FROM users WHERE username = $1`
	err := config.DB.QueryRow(query, loginReq.Username).Scan(
		&user.ID, &user.Username, &user.Email, &hashedPassword, &user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Terjadi kesalahan pada server"})
		return
	}

	// Verifikasi password dengan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginReq.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	// Generate JWT tokens
	accessToken, err := auth.GenerateJWT(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to generate access token",
		})
		return
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to generate refresh token",
		})
		return
	}

	// Login berhasil - tidak mengembalikan password
	user.Password = "" // Kosongkan password untuk keamanan

	c.JSON(http.StatusOK, LoginResponse{
		Success:      true,
		Message:      "Login berhasil",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	})
}

// GetUsers godoc
// @Summary      Get all users
// @Description  Menampilkan semua user
// @Tags         Users
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /api/users [get]
func GetUsers(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, username, email, role FROM users ORDER BY id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, u)
	}
	c.JSON(http.StatusOK, users)
}

// RefreshToken godoc
// @Summary      Refresh access token
// @Description  Generate new access token using refresh token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  LoginResponse
// @Failure      401  {object}  map[string]interface{}
// @Router       /api/auth/refresh [post]
func RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Refresh token required",
		})
		return
	}

	// Remove "Bearer " prefix if exists
	refreshToken := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		refreshToken = authHeader[7:]
	}

	claims, err := auth.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid refresh token: " + err.Error(),
		})
		return
	}

	// Get user data from database
	var user models.User
	var hashedPassword string
	query := `SELECT id, username, email, password, role FROM users WHERE id = $1`
	err = config.DB.QueryRow(query, claims.UserID).Scan(
		&user.ID, &user.Username, &user.Email, &hashedPassword, &user.Role,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	// Generate new access token
	newAccessToken, err := auth.GenerateJWT(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to generate new token",
		})
		return
	}

	user.Password = "" // Don't return password

	c.JSON(http.StatusOK, LoginResponse{
		Success:     true,
		Message:     "Token refreshed successfully",
		AccessToken: newAccessToken,
		User:        user,
	})
}

// GetProfile godoc
// @Summary      Get current user profile
// @Description  Get authenticated user profile information
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  models.User
// @Failure      401  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Router       /api/profile [get]
func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "User not authenticated",
		})
		return
	}

	var user models.User
	query := `SELECT id, username, email, role FROM users WHERE id = $1`
	err := config.DB.QueryRow(query, userID).Scan(
		&user.ID, &user.Username, &user.Email, &user.Role,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Database error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Profile retrieved successfully",
		"data":    user,
	})
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Menampilkan user berdasarkan ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  models.User
// @Failure      404  {object}  map[string]string
// @Router       /api/users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	err := config.DB.QueryRow("SELECT id, username, email, role FROM users WHERE id = $1", id).Scan(&u.ID, &u.Username, &u.Email, &u.Role)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, u)
}

// RegisterUser godoc
// @Summary      Register akun client
// @Description  Membuat akun baru untuk client (role otomatis client)
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User Data (tanpa role, role otomatis client)"
// @Success      201  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/users/register [post]
func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Set role otomatis ke client
	user.Role = "client"

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	query := `INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id`
	err = config.DB.QueryRow(query, user.Username, user.Email, string(hashedPassword), user.Role).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Memperbarui data user berdasarkan ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "User ID"
// @Param        user  body      models.User true  "User Data"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /api/users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query := `UPDATE users SET username=$1, email=$2, password=$3, role=$4 WHERE id=$5`
	res, err := config.DB.Exec(query, u.Username, u.Email, u.Password, u.Role, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser godoc
// @Summary      Delete user
// @Description  Menghapus user berdasarkan ID
// @Tags         Users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	res, err := config.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
