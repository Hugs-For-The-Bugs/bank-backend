package controller

import (
	"hugsforthebugs/bank-backend/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"hugsforthebugs/bank-backend/util"
)

// TransactionRequest
type TransactionRequest struct {
	TargetPhoneNumber string  `json:"targetPhoneNumber"`
	Amount            float64 `json:"amount"`
}
type stateType string

const (
	Successful stateType = "Successful"
	Failed     stateType = "Failed"
)

type Transaction struct {
	ID            uint64    `gorm:"primary_key" json:"id"`
	FromAccountID int       `json:"from_account_id"`
	ToAccountID   int       `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	Fee           float64   `json:"fee"`
	State         stateType `gorm:"type:ENUM('Successful', 'Failed')"`
}

// CreateTransaction
func CreateTransaction(c *gin.Context) {
	var request TransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		util.ServerErrorResponse(c, err.Error())
		return
	}
	if request.Amount <= 0 {
		util.BadRequestResponse(c, "Invalid amount")
		return
	}

	// Validation logic is performed here,
	// such as checking whether the account balance is sufficient

	// Start database transaction
	tx := util.DB.Begin()

	// Check whether the balance in the transferred account is sufficient
	var fromAccount model.Account
	session := sessions.Default(c)
	id := session.Get("id")
	result := util.DB.First(&fromAccount, "id =?", id)
	if result.Error != nil {
		tx.Rollback()
		util.BadRequestResponse(c, "User not found")
		return
	}
	if !fromAccount.Active {
		tx.Rollback()
		util.BadRequestResponse(c, "Account is not active")
		return
	}

	if float64(fromAccount.Balance) < float64(request.Amount)*1.01 {
		tx.Rollback()
		util.BadRequestResponse(c, "Insufficient balance")
		return
	}
	var toAccount model.Account
	result = tx.Where("phone = ?", request.TargetPhoneNumber).First(&toAccount)
	if result.Error != nil {
		tx.Rollback()
		util.BadRequestResponse(c, "User not found")
		return
	}

	if !toAccount.Active {
		tx.Rollback()
		util.BadRequestResponse(c, "to Account is not active")
		return
	}
	//Check whether the account is active

	// update the Account
	fromAccount.Balance -= request.Amount * 1.01
	result = tx.Save(&fromAccount)
	if result.Error != nil {
		tx.Rollback()
		util.ServerErrorResponse(c, result.Error.Error())
		return
	}

	// update the toAccount

	result = tx.Where("phone = ?", request.TargetPhoneNumber).First(&toAccount)
	if result.Error != nil {
		tx.Rollback()
		util.ServerErrorResponse(c, result.Error.Error())
		return
	}

	toAccount.Balance += request.Amount
	result = tx.Save(&toAccount)
	if result.Error != nil {
		tx.Rollback()
		util.ServerErrorResponse(c, result.Error.Error())
		return
	}

	//create transaction history

	transaction := Transaction{
		FromAccountID: int(fromAccount.ID),
		ToAccountID:   int(toAccount.ID),
		Amount:        request.Amount,
		Fee:           request.Amount * 0.01,
		State:         Successful,
	}
	result = tx.Create(&transaction)
	if result.Error != nil {
		tx.Rollback()
		util.ServerErrorResponse(c, result.Error.Error())
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"msg": "Transaction created successfully"})
}
