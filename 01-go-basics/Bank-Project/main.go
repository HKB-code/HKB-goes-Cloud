package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const accountFile = "balance.txt"
const TransactionFile = "Transaction.txt"

func writeToFile(data float64){
	d:= fmt.Sprint(data)
	os.WriteFile(accountFile,[]byte(d),0644)
}

func getBalanceFile()(float64,error){

	data,err:=os.ReadFile(accountFile)
if err != nil {
	return 1000,errors.New("couldn't read the file")
}
balance:= 	string(data)

balanceTxt,err := strconv.ParseFloat(balance,64)
if err != nil {
	return 1000, errors.New("Failed to parse stored balance value.")
}

return  balanceTxt,nil
}

func ReadTransaction(){
	d,err:=os.ReadFile(TransactionFile)
	if  err!= nil {
		fmt.Println("couldn't read the transaction file")
	}
	k:=string(d)
	fmt.Println(k)
}

func writeTransaction(data string){

f1,err:=os.OpenFile(TransactionFile,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0600)
if  err!= nil {
	fmt.Println(err)
}
defer f1.Close()
_,err=f1.WriteString(data+"\n")
if  err!= nil {
	fmt.Println(err)
}
}


func main() {

accountBalance,err:= getBalanceFile()
if _,err=os.Stat(TransactionFile);os.IsNotExist(err){
writeTransaction(fmt.Sprintf("Default Balance: %f",accountBalance))
}





if  err!= nil {
	fmt.Println(err)
}



	fmt.Println("Welcome To The Go Bank")
	for{
		fmt.Println("What do you want to do?")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Deposit")
	fmt.Println("3.Withdraw")
	fmt.Println("4.Transaction History")
	fmt.Println("5.Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice{
	case 1:
		fmt.Println("Here is your Balance: ", accountBalance)
	
	case 2:
		var deposit float64
		fmt.Scan(&deposit)
		if deposit<0{
			fmt.Println("invalid deposit")
			
		}
		accountBalance+=deposit
			fmt.Println("Account Balance Upgraded: ", accountBalance)
            writeToFile(accountBalance)
			str := fmt.Sprintf("Deposit | Amount: %.2f | Balance: %.2f | Time: %s", deposit,accountBalance,time.Now().Format(time.RFC3339))
			writeTransaction(str)
		
		continue
	case 3:
		var withdrawal float64
		fmt.Scan(&withdrawal)
		if withdrawal<=0 && withdrawal>accountBalance{
			fmt.Println("Invalid Withdrawal")
			continue
		}
		accountBalance-=withdrawal
			fmt.Println("Account Balance Upgraded: ", accountBalance)
			writeToFile(accountBalance)
			str := fmt.Sprintf("Withdrawal | Amount: %.2f | Balance: %.2f | Time: %s", withdrawal,accountBalance,time.Now().Format(time.RFC3339))
			writeTransaction(str)
		
		
	case 4:
		ReadTransaction()
	case 5:
		fmt.Println("GoodBye")
		return
	default:
		fmt.Println("enter a valid choice")
	}
	}
	
}