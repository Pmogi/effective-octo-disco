package Service

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/google/uuid"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Use a factory for making client so I can mock it
type MessageStoreService struct {
	DynamoDbClient *dynamodb.Client
}

// Message For marshaling, entries need to be tagged
type Message struct {
	Id        string `dynamodbav:"id"`
	Timestamp string `dynamodbav:"timestamp"`
	Message   string `dynamodbav:"message"`
	UserId    string `dynamodbav:"userId"`
}

var tableName = "message-store"

func (messageService MessageStoreService) StoreMessage(messageStr string, timestamp time.Time, userId string) error {

	id, _ := uuid.NewUUID()
	fmt.Printf("Adding message: %v %v %v %v\n", id.String(), messageStr, userId, timestamp.String())

	newMsg := Message{Message: messageStr, Timestamp: timestamp.String(), UserId: userId, Id: id.String()}
	msgItem, err := attributevalue.MarshalMap(newMsg)

	_, err = messageService.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("message-store"),
		Item:      msgItem,
	})

	if err != nil {
		println("Couldn't add to table, error: " + err.Error())
	}
	return err
}

func (messageService MessageStoreService) GetMessages() ([]Message, error) {
	var messages []Message
	var err error
	var response *dynamodb.ScanOutput

	response, err = messageService.DynamoDbClient.Scan(
		context.TODO(),
		&dynamodb.ScanInput{
			TableName: aws.String(tableName)})
	if err != nil {
		fmt.Println("Error when scanning table: " + err.Error())
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &messages)
	if err != nil {
		fmt.Println("Error when unmarshalling returning data: " + err.Error())
	}

	return messages, err

}
