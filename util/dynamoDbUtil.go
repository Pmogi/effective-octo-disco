package util

import "GoTweet/Service"

type DynamoDbApi interface {
	init() error
	GetAllItems(item Service.Message, tableName string) ([]Service.Message, error)
	PostItem(tableName string) error
}
