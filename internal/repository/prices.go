package repository

import (
	"context"
	"github.com/LuxAeterna-git/priceHistoryService/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func (m *MongoDB) StoreRawPrices(data []model.PriceItem) error {
	var d model.StoreRawData
	d.Items = data
	res, err := m.db.InsertOne(context.Background(), d)
	if err != nil {
		log.Println("Error while StoreRawPrices func ", err)
		return err
	}
	// Delete old data
	m.db.DeleteOne(context.Background(), bson.M{"_id": m.lastPricesReq})

	id := res.InsertedID.(primitive.ObjectID)
	m.lastPricesReq = id

	return nil
}
