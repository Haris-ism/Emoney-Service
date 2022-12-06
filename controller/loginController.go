package controller

import (
	"gin/model"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func (t *Repo) SignUp(c *gin.Context) {
	var body model.Sign

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}

	var accounts model.Users

	if err := t.DB.Find(&accounts, "email = ?", body.Email).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return
	}
	if accounts.ID != 0 {
		log.Println("data not Found")
		c.JSON(400, gin.H{
			"message": "email is already used",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		log.Println("failed to hash password")
	}
	accounts.Email = body.Email
	accounts.Name = body.Name
	accounts.Password = string(hash)

	if err1 := t.DB.Create(&accounts).Error; err1 != nil {
		log.Println(err1)
		return
	}

	var balance model.Balances
	balance.UsersID=int(accounts.ID)

	if err1 := t.DB.Create(&balance).Error; err1 != nil {
		log.Println(err1)
		return
	}
	c.JSON(201, gin.H{
		"message": "write success",
		"data":    accounts,
	})
}

func (t *Repo) SignIn(c *gin.Context) {
	var body model.Sign
	var accounts model.Users

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}

	if err := t.DB.Find(&accounts, "email = ?", body.Email).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return
	}
	if accounts.ID == 0 {
		log.Println("data not Found")
		c.JSON(400, gin.H{
			"message": "invalid email",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(accounts.Password), []byte(body.Password)); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": body.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT")))
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "/", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign In Success",
		"token":   tokenString,
	})
}
