package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func TestReturnsRightPositionOfLine(t *testing.T) {
	const name1 = "MyName"
	const name2 = "AnotherName"

	myData := extractData("Line1," + name1 + "\nAnotherLine," + name2)

	if len(myData) == 0 {
		t.Error("extracted data is empty")
	}

	if myData[0] != name1 {
		t.Errorf("Expected %s, received %s", name1, myData[0])
	}

	if myData[1] != name2 {
		t.Errorf("Expected %s, received %s", name2, myData[1])
	}
}

type mockDynamoDBAPI struct {
	dynamodbiface.DynamoDBAPI
}

var fakeDynamoDb []MyDataFromS3

func (m mockDynamoDBAPI) PutItem(putItem *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	var item MyDataFromS3
	dynamodbattribute.Unmarshal(putItem.Item["FileKey"], &item.FileKey)
	dynamodbattribute.Unmarshal(putItem.Item["Name"], &item.Name)
	fakeDynamoDb = append(fakeDynamoDb, item)
	return nil, nil
}

func TestReturnsFormattedData(t *testing.T) {
	const name1 = "Joseph"
	const name2 = "Lucas"
	const fileName = "myFile"
	const secondFileKey = "myFile-1"

	mockData := &mockDynamoDBAPI{}
	myData := []string{name1, name2}

	settings := &dynamoSettings{
		svc: mockData,
	}

	settings.insertIntoDynamoDB(myData, fileName)

	if len(fakeDynamoDb) != 2 {
		t.Error("Could not insert")
	}

	if fakeDynamoDb[1].FileKey != secondFileKey {
		t.Error("Invalid file key")
	}

}
