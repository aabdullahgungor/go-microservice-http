package product

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IProductRepository interface {
	GetAllProducts(ctx context.Context) ([]Product, error)
	GetProductById(ctx context.Context, id string) (Product, error)
	CreateProduct(ctx context.Context, product *Product) error
	EditProduct(ctx context.Context, product *Product) error
	DeleteProduct(ctx context.Context, id string) error
}

var (
	ErrProductNotFound = errors.New("FromRepository - product not found")
)

type MongoDbProductRepository struct {
	connectionPool *mongo.Database
	logger         log.Logger
}

func NewMongoDbProductRepository(db *mongo.Database, logger log.Logger) *MongoDbProductRepository {
	return &MongoDbProductRepository{
		connectionPool: db,
		logger:         log.With(logger, "repo", "mongodb"),
	}
}

func (m *MongoDbProductRepository) GetAllProducts(ctx context.Context) ([]Product, error) {

	productCollection := m.connectionPool.Collection("product")

	var products []Product
	productCursor, err := productCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}
	if err = productCursor.All(ctx, &products); err != nil {
		panic(err)
	}

	return products, err
}

func (m *MongoDbProductRepository) GetProductById(ctx context.Context, id string) (Product, error) {

	productCollection := m.connectionPool.Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)

	var product Product
	err := productCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&product)
	if err != nil {
		return Product{}, ErrProductNotFound
	}

	return product, nil

}

func (m *MongoDbProductRepository) CreateProduct(ctx context.Context, product *Product) error {

	productCollection := m.connectionPool.Collection("product")

	result, err := productCollection.InsertOne(ctx, product)

	if err != nil {
		panic(err)
	}

	fmt.Printf("\ndisplay the ids of the newly inserted objects: %v", result.InsertedID)

	return err
}

func (m *MongoDbProductRepository) EditProduct(ctx context.Context, product *Product) error {

	productCollection := m.connectionPool.Collection("product")

	bsonBytes, err := bson.Marshal(&product)

	if err != nil {
		panic(fmt.Errorf("can't marshal:%s", err))
	}

	var upt bson.M
	bson.Unmarshal(bsonBytes, &upt)

	update := bson.M{"$set": upt}

	result, err := productCollection.UpdateOne(ctx, bson.M{"_id": product.Id}, update)

	if err != nil {
		panic(err)
	}

	fmt.Println("Number of documents updated:" + strconv.Itoa(int(result.ModifiedCount)))

	return err
}

func (m *MongoDbProductRepository) DeleteProduct(ctx context.Context, id string) error {

	productCollection := m.connectionPool.Collection("product")

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := productCollection.DeleteOne(ctx, bson.M{"_id": objId})

	// check for errors in the deleting
	if err != nil {
		panic(err)
	}

	// display the number of documents deleted
	fmt.Println("deleting the first result from the search filter\n" + "Number of documents deleted:" + strconv.Itoa(int(result.DeletedCount)))

	return err
}
