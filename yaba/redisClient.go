package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
)

/********************************
*
* 		DATABASE REDIS
*
********************************/

// Book struct
type Book struct {
	ImagePath       string
	Title           string
	Author          string
	Rating          string
	NumberRating    string
	Price           string
	Length          string
	Genre           string
	FileSize        string
	Country         string
	DatePublication string
	Publisher       string
	Resume          string
	FilePath        string
}

func redisSetNewBook(client *redis.Client, book *Book) error {
	bookJSON, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
		return err
	}
	id := uuid.NewV5(uuid.NamespaceOID, book.Title)
	err = client.HMSet(RedisTableBooks, id, bookJSON).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func redisGetAllBooks(client *redis.Client) []Book {
	var bookObj Book
	var booksSlice []Book

	booksString, err := client.HGetAll(RedisTableBooks).Result()
	if err != nil {
		fmt.Println("error", err)
		return nil
	}
	for _, bookStr := range booksString {
		bookBytes := []byte(bookStr)
		err = json.Unmarshal(bookBytes, &bookObj)
		booksSlice = append(booksSlice, bookObj)
	}
	//printBookSlice(booksSlice)
	return booksSlice
}

func redisStartClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: RedisHostAddr,
	})
	return client
}

func redisIsClientConnected(client *redis.Client) error {
	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Redis error")
	} else if pong == RedisPong {
		fmt.Println("Redis connected")
	}
	return err
}