package main

import (
	"testing"
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

func TestInsertDataIntoDatabase(t *testing.T) {
	myData := []string{"data1", "data2"}
	insertIntoDynamoDB(myData, "file1")
}
