package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kiyuu10/2fa-sys/config"
	"github.com/kiyuu10/2fa-sys/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var (
		input models.User
		db    = config.DB
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	err := db.
		Where("email = ?", input.Email).
		First(&existingUser).
		Error
	if err == nil {
		c.JSON(http.StatusConflict,
			gin.H{"error": "Email already registered"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to hash password"})
		return
	}
	input.Password = string(hashedPassword)

	err = db.
		Create(&input).
		Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var (
		input models.User
		db    = config.DB
	)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err = db.
		Where("email = ?", input.Email).
		First(&user).
		Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}
