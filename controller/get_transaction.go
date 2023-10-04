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
	UserBalance              string `json:"user_balance"`

	FromPhone string `json:"from_phone"`
	ToPhone   string `json:"to_phone"`
}

func GetTransactions(c *gin.Context) {
	var transactionResponses []TransactionResponse

	session := sessions.Default(c)

	id := session.Get("id")

	result := util.DB.Raw(`SELECT transactions.*,
		CASE
	WHEN fromAccounts.id = 1 THEN fromAccounts.balance
	WHEN toAccounts.id = 1 THEN toAccounts.balance
	ELSE NULL
	END AS UserBalance,
		fromAccounts.social_security_number as from_social_security_number,
		toAccounts.social_security_number as to_social_security_number,
		fromAccounts.phone as FromPhone,
		toAccounts.phone as ToPhone
	FROM transactions
	JOIN accounts AS fromAccounts ON fromAccounts.id = transactions.from_account_id
	JOIN accounts AS toAccounts ON toAccounts.id = transactions.to_account_id
	WHERE fromAccounts.id = ? OR toAccounts.id = ?`, id, id).Scan(&transactionResponses)

	if result.Error == nil {
		util.SuccessResponse(c, transactionResponses)
	} else {
		util.BadRequestResponse(c, "bad request")
	}
}
