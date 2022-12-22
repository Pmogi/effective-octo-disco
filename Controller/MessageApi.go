package Controller

import (
	"GoTweet/Model"
	"GoTweet/Service"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type MessageApi struct {
	Engine *gin.Engine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (MessageApi MessageApi) Init() {
	// Should allow CORs
	MessageApi.Engine.Use(CORSMiddleware())

	var cfg aws.Config
	var err error

	fmt.Println("Current environment: ", os.Getenv("ENV"))

	if os.Getenv("ENV") == "local" {
		cfg, err = config.LoadDefaultConfig(
			context.TODO(),
			config.WithSharedConfigProfile("default"))
		cfg.Region = "us-east-1"

		if err != nil {
			fmt.Println("Error with local config: " + err.Error())
			return
		}

	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
		if err != nil {
			fmt.Println("Error with default config: " + err.Error())
			return
		}
	}

	s3Client := s3.NewFromConfig(cfg)
	messageStoreService := Service.MessageStoreService{
		DynamoDbClient: dynamodb.NewFromConfig(cfg)}

	MessageApi.initEndpoints(s3Client, &messageStoreService)
}

func (MessageApi MessageApi) initEndpoints(s3Client *s3.Client, messageStoreService *Service.MessageStoreService) {
	MessageApi.Engine.GET("/ping", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	MessageApi.Engine.GET("/bucket/data", func(c *gin.Context) {
		v2, err := s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
			Bucket: aws.String("misc-bucket-z123-4321"),
		})
		if err != nil {
			print(err.Error())
			return
		}

		data := ""

		for _, s := range v2.Contents {
			// Append keys in bucket to string to return
			data += *s.Key + "\n"
		}

		c.String(http.StatusOK, "Keys in bucket:\n"+data)
	})

	MessageApi.Engine.POST("/message", func(c *gin.Context) {
		var msgDto Model.MessageDTO
		var err error

		err = c.BindJSON(&msgDto)
		if err != nil {
			fmt.Println("Error adding new message: " + err.Error())
			c.String(http.StatusBadRequest, "Error with input")
		}

		err = messageStoreService.StoreMessage(msgDto.Message, time.Now(), msgDto.UserId)
		if err != nil {
			fmt.Println("Error adding new message: " + err.Error())
			c.String(http.StatusInternalServerError, "Error with input")
		}

		c.String(http.StatusOK, "Stored message.")
	})

	MessageApi.Engine.GET("/message", func(c *gin.Context) {
		var messages []Service.Message
		var err error

		// call ddb for messages, should probably paginate
		messages, err = messageStoreService.GetMessages()

		if err != nil {
			c.String(500, "Error retrieving messages from dynamoDB.")
		}

		c.JSON(http.StatusOK, messages)
	})

	MessageApi.Engine.GET("/health", func(c *gin.Context) {
		c.String(200, "Healthy")
	})
}
