package service

import (
	"bank-test/dto"
	"bank-test/entity"
	"bank-test/repo"
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

func NewTransactionService(transactionRepo *repo.TransactionRepo) TransactionService {
	accountRepo := repo.NewAccountRepo()
	accountService := NewAccountService(&accountRepo)

	merchantRepo := repo.NewMerchantRepo()
	merchantService := NewMerchantService(&merchantRepo)

	logRepo := repo.NewLogRepo()
	logService := NewLogService(&logRepo)

	return &TransactionServiceImpl{
		AccountService:  accountService,
		TransactionRepo: *transactionRepo,
		MerchantService: merchantService,
		LogService:      logService,
	}
}

type TransactionServiceImpl struct {
	TransactionRepo repo.TransactionRepo
	AccountService  AccountService
	MerchantService MerchantService
	LogService      LogService
}

func (service *TransactionServiceImpl) Create(transaction entity.Transaction) error {
	// service.AccountService.GetAccountId()
	var account entity.Account
	var accountId string

	var merchant entity.Merchant
	var merchantId string

	// getAccountById
	accountId = transaction.GetAccountId()
	account = service.AccountService.GetAccountId(accountId)

	// getMerchantById
	merchantId = transaction.GetMerchantId()
	merchant = service.MerchantService.GetMerchantId(merchantId)

	// payment merchant
	bill := account.GetBalance() - transaction.GetAmount()
	err := service.TransactionRepo.Create(transaction)

	if account.GetId() == "" {
		err := errors.New("access ilegal, data not found")
		return err

	} else if merchant.GetId() == "" {
		err := errors.New("access ilegal, data not found")
		return err

	} else {
		transaction.SetId("transA")
		account.SetBalance(bill)

		//get list account
		var accounts []entity.Account
		byteValue, _ := ioutil.ReadFile("./jsonDb/account.json")
		err := json.Unmarshal(byteValue, &accounts)
		if err != nil {
			return err
		}
		//update account
		for idx, val := range accounts {
			if val.GetId() == accountId {
				accounts = append(accounts[:idx], accounts[idx+1:]...)
				accounts = append(accounts, account)
				tokensByte, _ := json.Marshal(accounts)
				err = ioutil.WriteFile("./jsonDb/account.json", tokensByte, 0644)
				return err
			}
		}
	}

	var log dto.Log
	currentTime := time.Now()
	log.SetId("Transaction ID: " + transaction.GetId())
	log.SetTime(currentTime.Format("2006-01-02 15:04:05"))
	log.SetName("Transaction Name: " + transaction.GetId())
	service.LogService.Create(log)
	return err
}

func (service *TransactionServiceImpl) GetAll() []entity.Transaction {
	transactions := service.TransactionRepo.GetAll()
	return transactions
}
