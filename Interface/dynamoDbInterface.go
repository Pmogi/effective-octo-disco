package Interface

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBAPI interface {
	BatchGetItemRequest(*dynamodb.BatchGetItemInput) dynamodb.BatchGetItemRequest

	BatchGetItemPages(*dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool) error
	BatchGetItemPagesWithContext(aws.Context, *dynamodb.BatchGetItemInput, func(*dynamodb.BatchGetItemOutput, bool) bool, ...aws.Option) error

	BatchWriteItemRequest(*dynamodb.BatchWriteItemInput) dynamodb.BatchWriteItemRequest

	CreateTableRequest(*dynamodb.CreateTableInput) dynamodb.CreateTableRequest

	DeleteItemRequest(*dynamodb.DeleteItemInput) dynamodb.DeleteItemRequest

	DeleteTableRequest(*dynamodb.DeleteTableInput) dynamodb.DeleteTableRequest

	DescribeLimitsRequest(*dynamodb.DescribeLimitsInput) dynamodb.DescribeLimitsRequest

	DescribeTableRequest(*dynamodb.DescribeTableInput) dynamodb.DescribeTableRequest

	DescribeTimeToLiveRequest(*dynamodb.DescribeTimeToLiveInput) dynamodb.DescribeTimeToLiveRequest

	GetItemRequest(*dynamodb.GetItemInput) dynamodb.GetItemRequest

	ListTablesRequest(*dynamodb.ListTablesInput) dynamodb.ListTablesRequest

	ListTablesPages(*dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool) error
	ListTablesPagesWithContext(aws.Context, *dynamodb.ListTablesInput, func(*dynamodb.ListTablesOutput, bool) bool, ...aws.Option) error

	ListTagsOfResourceRequest(*dynamodb.ListTagsOfResourceInput) dynamodb.ListTagsOfResourceRequest

	PutItemRequest(*dynamodb.PutItemInput) dynamodb.PutItemRequest

	QueryRequest(*dynamodb.QueryInput) dynamodb.QueryRequest

	QueryPages(*dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool) error
	QueryPagesWithContext(aws.Context, *dynamodb.QueryInput, func(*dynamodb.QueryOutput, bool) bool, ...aws.Option) error

	ScanRequest(*dynamodb.ScanInput) dynamodb.ScanRequest

	ScanPages(*dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool) error
	ScanPagesWithContext(aws.Context, *dynamodb.ScanInput, func(*dynamodb.ScanOutput, bool) bool, ...aws.Option) error

	TagResourceRequest(*dynamodb.TagResourceInput) dynamodb.TagResourceRequest

	UntagResourceRequest(*dynamodb.UntagResourceInput) dynamodb.UntagResourceRequest

	UpdateItemRequest(*dynamodb.UpdateItemInput) dynamodb.UpdateItemRequest

	UpdateTableRequest(*dynamodb.UpdateTableInput) dynamodb.UpdateTableRequest

	UpdateTimeToLiveRequest(*dynamodb.UpdateTimeToLiveInput) dynamodb.UpdateTimeToLiveRequest

	WaitUntilTableExists(*dynamodb.DescribeTableInput) error
	WaitUntilTableExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...aws.WaiterOption) error

	WaitUntilTableNotExists(*dynamodb.DescribeTableInput) error
	WaitUntilTableNotExistsWithContext(aws.Context, *dynamodb.DescribeTableInput, ...aws.WaiterOption) error
}
