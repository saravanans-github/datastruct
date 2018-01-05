package main

import (
	"testing"
)

func TestAppendElement(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])

	if (*start).value != testData[0].value || (*(*start).next).value != testData[1].value {
		t.FailNow()
	}

	if length != len(testData) {
		t.FailNow()
	}
}

func TestRemoveElementAtIndex(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}, Element{"Good day!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])
	AppendElement(testData[2])

	RemoveElementAtIndex(2)

	if start.value != testData[0].value || start.next.value != testData[1].value || start.next.next != nil {
		t.FailNow()
	}

	if length != (len(testData) - 1) {
		t.FailNow()
	}
}
