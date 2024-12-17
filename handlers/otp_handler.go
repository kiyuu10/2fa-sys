package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kiyuu10/2fa-sys/config"
	"github.com/kiyuu10/2fa-sys/models"
	"github.com/kiyuu10/2fa-sys/utils"
)

func GenerateAndSendOTP(c *gin.Context) {
	var (
		input OTPEmailReq
		db    = config.DB
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := db.Where("email = ?", input.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Generate OTP
	otp := utils.GenerateOTP()
	otpHash := utils.HashOTP(otp)

	user.OTPHash = otpHash
	user.OTPExpire = time.Now().Add(5 * time.Minute).Unix()
	db.Save(&user)

	err = utils.SendEmail(user.Email, otp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func VerifyOTP(c *gin.Context) {
	var (
		input OTPEmailVerifyReq
		db    = config.DB
	)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := db.Where("email =?", input.Email).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.OTPHash == "" || user.OTPExpire < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP or expired"})
		return
	}

	otpHash := utils.HashOTP(input.OTP)
	if user.OTPHash != otpHash {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified"})
}
