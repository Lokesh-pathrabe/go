package models

import (
	"context"
	"fmt"

	"github.com/lokesh/mongo_go/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Collection

type Book struct {
	Name        string `bson:"name" json:"name"`
	Author      string `bson:"author" json:"author"`
	Publication string `bson:"publication"json:"publication"`
}

func init() {
	db = config.Connect()
	fmt.Println(db)
}

func (b *Book) CreateBook() *Book {
	db.InsertOne(context.TODO(), &b)
	return b
}

func GetAllBooks() []Book {
	var book []Book
	b, _ := db.Find(context.TODO(), bson.M{})
	b.All(context.TODO(), &book)
	fmt.Println(book)
	return book
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db.FindOne(context.TODO(), &getBook)
	return &getBook, db
}

// func DeleteBook(ID int64) Book {
// 	var book Book
// 	db.Where("ID=?", ID).Delete(&book)
// 	return book
// }
