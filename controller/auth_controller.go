package controller

import (
	"adriandidimqttgate/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

var RegisterRequest struct {
    Email         string `json:"email" validate:"required,email"`
    Password      string `json:"password" validate:"required,min=6"`
    HousingAreaId int64  `json:"housingAreaId" validate:"required"`
}

var LoginRequest struct {
    Email         string `json:"email" validate:"required,email"`
    Password      string `json:"password" validate:"required,min=6"`
}

func Login(c *gin.Context) {
	if err := c.ShouldBindJSON(&LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": err.Error(),
		})
		return
	}
		
	if err := validate.Struct(LoginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	var user model.User
	    if err := model.DB.Where("email = ?", LoginRequest.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  false,
            "message": "Pengguna tidak ditemukan",
        })
        return
    }

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginRequest.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "status":  false,
            "message": "Password salah",
        })
        return
    }

	hashedToken, err := bcrypt.GenerateFromPassword([]byte("mqttgatedidiadrian"), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
			"status": false,
			"message": err,
		})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "status":  true,
        "message": "Berhasil login",
        "token": hashedToken,
    })

}

func Register(c *gin.Context) {
	if err := c.ShouldBindJSON(&RegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": err.Error(),
		})
		return
	}
	
	if err := validate.Struct(RegisterRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	

	var existingUser model.User
    if err := model.DB.Where("email = ?", RegisterRequest.Email).First(&existingUser).Error; err == nil {
       	c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "Email sudah digunakan",
		})
        return
    }
	var housingArea model.HousingArea
	err := model.DB.First(&housingArea, RegisterRequest.HousingAreaId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H {
			"status": false,
			"message": "Perumahan tidak ada",
		})
        return
	}

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(RegisterRequest.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
			"status": false,
			"message": err,
		})
        return
    }

	newUser := model.User{
        Email: RegisterRequest.Email,
        Password: string(hashedPassword),
		HousingAreaID: RegisterRequest.HousingAreaId,
		HousingArea: housingArea,
    }
	
	if err := model.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"status": false,
			"message": err,
		})
        return
	}

	model.DB.Save(&newUser)

	c.JSON(http.StatusCreated, gin.H{
        "status":  true,
        "message": "Pengguna berhasil terdaftar",
		"data": &newUser,
    })
}