package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/anudevSaraswat/cutshort-url/database"
	"github.com/anudevSaraswat/cutshort-url/helpers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"custom_short"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"custom_short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"x_rate_remaining"`
	XRateLimitReset time.Duration `json:"x_rate_limit_reset"`
}

func APIShortenURL(c *gin.Context) {

	var payload Request
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO: implement rate limiting

	// check if the input is an actual URL
	urlToShort, isURL := helpers.ValidateURL(payload.URL)
	if !isURL {
		c.JSON(http.StatusBadRequest, "Please enter a valid URL")
		return
	}

	// check if input URL's domain isn't the domain of our app
	isSelfDomain := helpers.CheckForSelfDomain(urlToShort)
	if isSelfDomain {
		c.JSON(http.StatusBadRequest, "Invalid Domain")
		return
	}

	// produce a short URL
	shortURLInfo, err := helpers.GenerateShortURL(urlToShort, false)
	if err != nil {
		log.Default().Println("(ShortenURL) err in helpers.GenerateShortURL:", err.Error())
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	err = database.InsertDocument(shortURLInfo)
	if err != nil {
		log.Default().Println("(ShortenURL) err in database.InsertDocument:", err.Error())
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, Response{
		URL: shortURLInfo.ShortURL,
	})

}

func APIResolveURL(c *gin.Context) {

	shortURL := c.Param("short_url")

	filter := bson.D{{"short_url", shortURL}}

	var (
		originalURL string
		isExist     bool
	)

	// check for the url in cache, if cache miss then request the db
	originalURL, isExist = database.GetKey(shortURL)

	if !isExist {
		result, err := database.FindDocument(filter, 1)
		if err != nil {
			log.Default().Println("(APIResolveURL) err in database.FindDocument:", err.Error())
			c.JSON(http.StatusInternalServerError, "Something went wrong")
		}

		urlInfo := helpers.URLInfo{}
		if result.Next(context.TODO()) {
			err = result.Decode(&urlInfo)
			if err != nil {
				log.Default().Println("(APIResolveURL) err in result.Decode:", err.Error())
				c.JSON(http.StatusInternalServerError, "Something went wrong")
			}
		}

		originalURL = urlInfo.OriginalURL
	}

	if originalURL != "" {
		// set the data in the cache
		database.SetKey(shortURL, originalURL)
		c.Redirect(http.StatusPermanentRedirect, originalURL)
	} else {
		c.JSON(http.StatusBadRequest, "invalid short url!")
	}

}
