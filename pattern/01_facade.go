package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Фасад — это структурный паттерн, который предоставляет простой (но урезанный) интерфейс к сложной системе объектов, библиотеке или фреймворку.

Применение:
    Нужно представить простой или урезанный интерфейс к сложной подсистеме.
    Нужно уменьшить количество зависимостей между клиентом и сложной системой.
    Нужно разложить подсистему на отдельные слои.

Плюсы:
	Изолирует клиентов от компонентов сложной подсистемы.

Минусы:
	Фасад рискует стать божественным объектом, привязанным ко всем структурам программы.

Примеры:
	Функция Login() в системе авторизации
	Функция Run() в любом веб сервере
*/

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting account creation")
	walletFacacde := &WalletFacade{
		account:      newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *WalletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Printf("Adding %v to wallet...\n", amount)
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.addToBalance(amount)
	return nil
}

func (w *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Printf("Deducting %v from wallet...\n", amount)
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.deductFromBalance(amount)
	if err != nil {
		return err
	}
	return nil
}

type Account struct {
	name string
}

func newAccount(accountName string) *Account {
	return &Account{
		name: accountName,
	}
}

func (a *Account) checkAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("Invalid account name")
	}
	fmt.Println("Account verified")
	return nil
}

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) checkCode(code int) error {
	if s.code != code {
		return fmt.Errorf("Invalid security code")
	}
	fmt.Println("Security code verified")
	return nil
}

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (w *Wallet) addToBalance(amount int) {
	w.balance += amount
	fmt.Printf("Added %v successfully\n", amount)
	return
}

func (w *Wallet) deductFromBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("Insufficient balance")
	}
	w.balance = w.balance - amount
	fmt.Printf("Deducted %v successfully\n", amount)
	return nil
}

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("user", 1234)
	fmt.Println()

	err := walletFacade.addMoneyToWallet("user", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.deductMoneyFromWallet("user", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
