package Service

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"testing"
)

type EndpointCache struct {
}

type ClientStub struct {
	stub          string
	endpointCache *EndpointCache
}

type PutItemInput struct {
}

type Options struct {
}

type PutItemOutput struct {
}

func (cs *ClientStub) PutItem(ctx context.Context, params *PutItemInput, optFns ...func(*Options)) (*PutItemOutput, error) {
	return nil, errors.New("mock aws error")
}

type ScanInput struct {
}

type ScanOutput struct {
	items []map[string]types.AttributeValue
}

func (cs *ClientStub) GetMessages(ctx context.Context, params *ScanInput, optFns ...func(*Options)) (*ScanOutput, error) {

	testMsg := Message{Id: "id", Timestamp: "01-01-01", Message: "testmsg", UserId: "user"}

	// need to unravel the shenanigans on this type
	testItem, err := attributevalue.MarshalMap(testMsg)

	testItemList := []map[string]types.AttributeValue{testItem}

	if err != nil {
		panic(err)
	}

	return &ScanOutput{items: testItemList}, err
}

func TestMessageStoreService_StoreMessages(t *testing.T) {

}

func TestMessageStoreService_GetMessage(t *testing.T) {
	cs := ClientStub{stub: "I'm a stub!"}
}
