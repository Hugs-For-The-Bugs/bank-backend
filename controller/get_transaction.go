package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hugsforthebugs/bank-backend/util"
)

type TransactionResponse struct {
	ID                       uint64 `gorm:"primaryKey" json:"id"`
	FromSocialSecurityNumber string `json:"from_social_security_number"`
	ToSocialSecurityNumber   string `json:"to_social_security_number"`
	Amount                   string `json:"amount"`
	State                    string `json:"state"`
}

func GetTransactions(c *gin.Context) {
	var transactionResponses []TransactionResponse

	session := sessions.Default(c)

	id := session.Get("id")

	//result := util.DB.Model(&model.Transaction{}).Preload("Accounts").Where("from_account_id = ? OR to_account_id = ?", id, id).Find(&transactionResponses)
	result := util.DB.Raw(`
	SELECT transactions.*, fromAccounts.social_security_number as from_social_security_number, fromAccounts.id, toAccounts.social_security_number as to_social_security_number, toAccounts.id FROM transactions
	JOIN accounts AS fromAccounts ON fromAccounts.id = transactions.from_account_id
	JOIN accounts AS toAccounts ON toAccounts.id = transactions.to_account_id
	WHERE fromAccounts.id = ? OR toAccounts.id = ?
	`, id, id).Scan(&transactionResponses)
	//result := util.DB.Where("SELECT * FROM accounts JOIN transactions ON accounts.id = transactions.from_account_id OR accounts.id = transactions.to_account_id;").Find(&transactions)

	if result.Error == nil {
		util.SuccessResponse(c, transactionResponses)
	} else {
		util.BadRequestResponse(c, "bad request")
	}
}
