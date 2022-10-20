package signedTransaction

import (
	"powblockchain/utils"

	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
)

type SignedTransaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewSignedTransaction(senderPrivateKey *ecdsa.PrivateKey, senderPublicKey *ecdsa.PublicKey,
	senderAddress string, recipientAddress string, value float32) *SignedTransaction {
	return &SignedTransaction{
		senderPrivateKey: senderPrivateKey,
		senderPublicKey:  senderPublicKey,
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		value:            value,
	}
}

func (signedTransaction *SignedTransaction) MashalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddress    string  `json:"sender_address"`
		RecipientAddress string  `json:"recipient_address"`
		Value            float32 `json:"value"`
	}{
		SenderAddress:    signedTransaction.senderAddress,
		RecipientAddress: signedTransaction.recipientAddress,
		Value:            signedTransaction.value,
	})
}

func (signedTransaction *SignedTransaction) SenderAddress() string {
	return signedTransaction.senderAddress
}

func (signedTransaction *SignedTransaction) RecipientAddress() string {
	return signedTransaction.recipientAddress
}

func (signedTransaction *SignedTransaction) Value() float32 {
	return signedTransaction.value
}

func (signedTransaction *SignedTransaction) Signature() *utils.Signature {
	serializedTransaction, _ := json.Marshal(signedTransaction)
	transactionHash := sha256.Sum256(serializedTransaction)
	r, s, _ := ecdsa.Sign(rand.Reader, signedTransaction.senderPrivateKey, transactionHash[:])
	return &utils.Signature{
		R: r,
		S: s,
	}
}

func (transaction *SignedTransaction) MarshalJSON() ([]byte, error) {
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
