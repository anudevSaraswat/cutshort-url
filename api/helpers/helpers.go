package helpers

import (
	"context"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/anudevSaraswat/cutshort-url/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var COUNTER = 1000000000000

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type URLInfo struct {
	OriginalURL string    `bson:"original_url"`
	ShortURL    string    `bson:"short_url"`
	Counter     int       `bson:"counter"`
	CreatedOn   time.Time `bson:"created_on"`
	IsCustom    bool      `bson:"is_custom"`
}

func getCounter() (int, error) {

	// read the last counter value from db
	// if there is a value, return value + 1
	// else return counter defined above

	db := database.ConnectToDB()

	dbName := os.Getenv("OBJECT_DB_NAME")

	urlCollection := db.Database(dbName).Collection("urls")

	counterResult := urlCollection.FindOne(context.TODO(), bson.D{{}}, &options.FindOneOptions{
		Sort: map[string]interface{}{
			"counter": -1,
		},
	})

	urlInfo := URLInfo{}
	err := counterResult.Decode(&urlInfo)
	if err != nil && err != mongo.ErrNoDocuments {
		log.Default().Println("(getCounter) err in counterResult.Decode:", err)
		return COUNTER, err
	}

	if urlInfo.Counter > 0 {
		return urlInfo.Counter + 1, nil
	}

	return COUNTER, nil

}

func ValidateURL(urlToValidate string) (*url.URL, bool) {

	parsedURL, err := url.ParseRequestURI(urlToValidate)
	if err != nil {
		log.Default().Println("(validateURL) err in url.Parse:", err)
		return nil, false
	}

	return parsedURL, true

}

func CheckForSelfDomain(urlToCheck *url.URL) bool {

	domain := os.Getenv("DOMAIN")

	return urlToCheck.Host == domain

}

func GenerateShortURL(originalURL *url.URL, isCustom bool) (*URLInfo, error) {

	var shortURL string

	currentCounterValue, err := getCounter()
	if err != nil {
		log.Default().Println("(GenerateShortURL) err in getCounter:", err)
		return nil, err
	}

	shortURL = toBase62(currentCounterValue)

	urlInfo := &URLInfo{
		OriginalURL: originalURL.String(),
		ShortURL:    shortURL,
		Counter:     currentCounterValue,
		CreatedOn:   time.Now(),
		IsCustom:    isCustom,
	}

	return urlInfo, nil

}

func toBase62(value int) string {

	indexes := []int{}
	for {
		if value < 62 {
			indexes = append(indexes, value)
			break
		}
		indexes = append(indexes, value%62)
		value /= 62
	}

	// read the array in reverse to get base62 string

	var str strings.Builder
	for i := len(indexes) - 1; i >= 0; i-- {
		str.WriteString(string(chars[indexes[i]]))
	}

	shortURL := str.String()

	return shortURL

}
