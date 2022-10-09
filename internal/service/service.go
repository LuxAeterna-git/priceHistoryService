package service

import (
	"encoding/json"
	"fmt"
	"github.com/LuxAeterna-git/priceHistoryService/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Repository interface {
	StoreRawPrices(data []model.PriceItem) error
	GetRawData() ([]model.PriceItem, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetRawData() error {
	r, err := http.NewRequest(http.MethodGet, "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json", nil)
	if err != nil {
		log.Println("Failed to build request for raw price data: ", err)
		return err
	}
	client := &http.Client{
		Timeout: time.Duration(time.Minute),
	}
	res, err := client.Do(r)
	if err != nil {
		log.Println("Failed to get raw price data: ", err)
		return err
	}
	b, _ := ioutil.ReadAll(res.Body)
	var result model.RawPriceResponse
	json.Unmarshal(b, &result)
	err = s.repo.StoreRawPrices(result.Ethereum.Transactions)
	if err != nil {
		log.Println("Failed to store raw price data: ", err)
		return err
	}
	return nil
}

func (s *Service) CalculateValues() error {
	rawData, err := s.repo.GetRawData()
	if err != nil {
		return err
	}
	results := s.GasPerMonth(rawData)
	fmt.Println(results)
	return nil
}

func (s *Service) GasPerMonth(data []model.PriceItem) []float64 {
	results := make([]float64, 12)
	for _, item := range data {
		t, err := time.Parse("06-01-02", item.Time[:8])
		if err != nil {
			log.Println("Failed to parse data: ", err)
		}
		results[t.Month()-1] += item.GasValue * item.GasPrice
	}
	return results
}
