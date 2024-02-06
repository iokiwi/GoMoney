package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aclindsa/ofxgo"
)

func main() {

	arg := os.Args[1]

	f, err := os.Open(arg)
	if err != nil {
		fmt.Printf("can't open file: %v\n", err)
		return
	}
	defer f.Close()

	resp, err := ofxgo.ParseResponse(f)
	if err != nil {
		fmt.Printf("can't parse response: %v\n", err)
		return
	}

	if stmt, ok := resp.CreditCard[0].(*ofxgo.CCStatementResponse); ok {
		fmt.Printf("Balance: %s\n", stmt.BalAmt)
		fmt.Println("Transactions:")
		for _, tran := range stmt.BankTranList.Transactions {
			amount, _ := tran.TrnAmt.Float64()
			fmt.Printf("%s %10.2f\t%s\n", tran.DtPosted.Format(time.DateOnly), amount, tran.Name)
			// fmt.Printf("%s %-15s %-11s %s%s%s\n", tran.DtPosted, tran.TrnAmt.String()+" "+currency.String(), tran.TrnType, tran.Name, tran.Payee.Name, tran.Memo)
		}
	}
}