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

func TestRemoveElementAtIndex_Middle(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}, Element{"Good day!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])
	AppendElement(testData[2])

	RemoveElementAtIndex(1)

	if start.value != testData[0].value || start.next.value != testData[2].value || start.next.next != nil {
		t.FailNow()
	}

	if length != (len(testData) - 1) {
		t.FailNow()
	}
}

func TestRemoveElementAtIndex_First(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}, Element{"Good day!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])
	AppendElement(testData[2])

	RemoveElementAtIndex(0)
	t.Log(start.value)
	t.Log(start.next.value)
	t.Log(start.next.next)

	if start.value != testData[1].value || start.next.value != testData[2].value || start.next.next != nil {
		t.FailNow()
	}

	if length != (len(testData) - 1) {
		t.FailNow()
	}
}

func TestRemoveElementAtIndex_Last(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}, Element{"Good day!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])
	AppendElement(testData[2])

	RemoveElementAtIndex(2)
	t.Log(start.value)
	t.Log(start.next.value)
	t.Log(start.next.next)

	if start.value != testData[0].value || start.next.value != testData[1].value || start.next.next != nil {
		t.FailNow()
	}

	if length != (len(testData) - 1) {
		t.FailNow()
	}
}
