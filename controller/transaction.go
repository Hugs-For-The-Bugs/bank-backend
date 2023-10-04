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
	TargetPhoneNumber string `json:"targetPhoneNumber"`
	Amount            int    `json:"amount"`
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
	Amount        int       `json:"amount"`
	State         stateType `gorm:"type:ENUM('Successful', 'Failed')"`
}

// CreateTransaction
func CreateTransaction(c *gin.Context) {
	var request TransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		util.BadRequestResponse(c, "Account not active")
		c.JSON(http.StatusBadRequest, gin.H{"error": "From account is not active"})
		return
	}

	if fromAccount.Balance <= request.Amount {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}
	var toAccount model.Account
	result = tx.Where("phone = ?", request.TargetPhoneNumber).First(&toAccount)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if !toAccount.Active {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "To account is not active"})
		return
	}
	//Check whether the account is active

	// update the fromAccount
	fromAccount.Balance -= request.Amount
	result = tx.Save(&fromAccount)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// update the toAccount

	result = tx.Where("phone = ?", request.TargetPhoneNumber).First(&toAccount)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	toAccount.Balance += request.Amount
	result = tx.Save(&toAccount)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	//create transaction history

	transaction := Transaction{
		FromAccountID: int(fromAccount.ID),
		ToAccountID:   int(toAccount.ID),
		Amount:        request.Amount,
		State:         Successful,
	}
	result = tx.Create(&transaction)
	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully"})
}
