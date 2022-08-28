package transaction

import (
	"encoding/json"
	"log"
	"strings"
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(recipientAddress string, senderAddress string, value float32) *Transaction {
	return &Transaction{
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
	}
}

func (transaction *Transaction) Print() {
	log.Printf("%s\n", strings.Repeat("-", 64))
	log.Printf("sender address			%s\n", transaction.senderAddress)
	log.Printf("recipient address		%s\n", transaction.recipientAddress)
	log.Printf("value				%.2f\n", transaction.value)
}

func (transaction *Transaction) MashalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddress    string  `json:"sender_address"`
		RecipientAddress string  `json:"recipient_address"`
		Value            float32 `json:"value"`
	}{
		SenderAddress:    transaction.senderAddress,
		RecipientAddress: transaction.recipientAddress,
		Value:            transaction.value,
	})
}
