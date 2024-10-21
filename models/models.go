package models

type P2PServerMessage struct {
	PeerAddress 	string
	Transaction 	Transaction
}

// Transaction represents a single transaction in the blockchain
type Transaction struct {
	Sender    string	`json:"sender"			validate:"required"`
	Recipient string	`json:"recipient" 		validate:"required"`
	Amount    float64	`json:"amount"			validate:"required"`
}
