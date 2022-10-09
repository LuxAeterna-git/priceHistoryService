package model

type PriceItem struct {
	Time           string  `json:"time" bson:"time"`
	GasPrice       float64 `json:"gasPrice" bson:"gasPrice"`
	GasValue       float64 `json:"gasValue" bson:"gasValue"`
	Average        float64 `json:"average" bson:"average"`
	MaxGasPrice    float64 `json:"maxGasPrice" bson:"maxGasPrice"`
	MedianGasPrice float64 `json:"medianGasPrice" bson:"medianGasPrice"`
}

type RawPriceResponse struct {
	Ethereum struct {
		Transactions []PriceItem `json:"transactions"`
	} `json:"ethereum"`
}

type StoreRawData struct {
	Items []PriceItem `bson:"items"`
}
