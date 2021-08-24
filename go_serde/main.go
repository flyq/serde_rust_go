package main

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

type signatureContent struct {
	TokenType string `json:"token_type"`
	Price     uint64 `json:"price"`
	TimeStamp uint64 `json:"time_stamp"`
}

func main() {
	content := signatureContent{
		TokenType: "ETH",
		Price:     10000,
		TimeStamp: 0,
	}
	msg, _ := json.Marshal(content)
	fmt.Println("msg bytes:", msg)
	fmt.Println("msg in hex:", hexutil.Encode(msg))
}
