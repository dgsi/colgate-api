package models

import ( 
	"encoding/json"
)

type ItemReward struct {
	ItemName  string `json:"item_name"`
	Qty json.Number `json:"qty,Number"`
}