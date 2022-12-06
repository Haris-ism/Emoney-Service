package controller

import (
	"encoding/json"
	"fmt"
	"gin/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *Repo) GetProfile(c *gin.Context) {
	email,_:=c.Get("email")
	var accounts model.Users

	if err := t.DB.Find(&accounts, "email = ?", email).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return
	}
	if accounts.ID == 0 {
		log.Println("data not Found")
		c.JSON(400, gin.H{
			"message": "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get profile success",
		"data":    accounts,
	})
}

func (t *Repo) GetBalance(c *gin.Context) {
	email,_:=c.Get("email")
	var accounts model.Users
	var balance model.Balances
	if err := t.DB.Find(&accounts, "email = ?", email).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return
	}
	
	if err := t.DB.Find(&balance, "users_id = ?", accounts.ID).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to get balance",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "get balance success",
		"data":    balance,
	})
}

func (t *Repo) TopUpBalance(c *gin.Context) {
	var body model.TopUp

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}
	email,_:=c.Get("email")
	var accounts model.Users
	var balance model.Balances
	if err := t.DB.Find(&accounts, "email = ?", email).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return
	}
	if err := t.DB.Find(&balance, "users_id = ?", accounts.ID).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to get balance",
		})
		return
	}
	balance.Balance = balance.Balance+body.TopUp
	balance.UsersID=int(accounts.ID)

	if err := t.DB.Save(&balance).Error; err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "top-up failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "top-up success",
		"data":    balance,
	})
}

func (t *Repo) Inquiry(c *gin.Context) {
	response, err := http.Get("https://phoenix-imkas.ottodigital.id/interview/biller/v1/list")
    if err != nil {
        c.Status(http.StatusServiceUnavailable)
        return
    }
    
    defer response.Body.Close()    
     
    if response.StatusCode != http.StatusOK {
        c.Status(http.StatusServiceUnavailable)
        return
    }
    
    var req model.Requests
    json.NewDecoder(response.Body).Decode(&req)
	c.JSON(http.StatusOK, gin.H{
		"message": "Please Select Your Inquiry",
		"data":req.Data,
	})
}

func (t *Repo) TransactionConfirmation(c *gin.Context) {
	var body model.Confirm

	if err := c.BindJSON(&body); err != nil {
		log.Println(err)
		return
	}

	email,_:=c.Get("email")

	var accounts model.Users
	var balance model.Balances
	var history model.History

	if err := t.DB.Find(&accounts, "email = ?", email).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return
	}
	if err := t.DB.Find(&balance, "users_id = ?", accounts.ID).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to get balance",
		})
		return
	}

	url:=fmt.Sprintf("https://phoenix-imkas.ottodigital.id/interview/biller/v1/detail?billerId=%v",body.InquiryID)
	response, err := http.Get(url)
    if err != nil {
        c.Status(http.StatusServiceUnavailable)
        return
    }
    
    defer response.Body.Close()    
     
    if response.StatusCode != http.StatusOK {
        c.Status(http.StatusServiceUnavailable)
        return
    }
    
    var req model.Requests1
    json.NewDecoder(response.Body).Decode(&req)	
	
	balance.Balance = balance.Balance-req.Data.Price-req.Data.Fee
	balance.UsersID=int(accounts.ID)

	if balance.Balance<0{
		c.JSON(http.StatusOK, gin.H{
			"message": "Insufficient Balance",
		})
		return
	}

	if err := t.DB.Save(&balance).Error; err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "update balance failed",
		})
		return
	}

	history.UsersID=int(accounts.ID)
	history.Price=req.Data.Price
	history.Fee=req.Data.Fee
	history.Product=req.Data.Product
	history.Description=req.Data.Description

	if err := t.DB.Save(&history).Error; err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "update history failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Confirmation Success",
		"data":req.Data,
	})
}

func (t *Repo) History(c *gin.Context) {
	email,_:=c.Get("email")
	var accounts model.Users
	var history []model.History
	if err := t.DB.Find(&accounts, "email = ?", email).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to find email",
		})
		return
	}
	if err := t.DB.Find(&history, "users_id = ?", accounts.ID).Error; err != nil {
		log.Println("failed to get data")
		c.JSON(500, gin.H{
			"message": "failed to get history",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get History Success",
		"data":history,
	})
}