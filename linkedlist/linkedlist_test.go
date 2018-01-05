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

	if start.value != testData[0].value || start.next.value != testData[1].value || start.next.next != nil {
		t.FailNow()
	}

	if length != (len(testData) - 1) {
		t.FailNow()
	}
}

func TestElementAtIndex_First(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}, Element{"Good day!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])
	AppendElement(testData[2])

	e := ElementAtIndex(0)

	if e.value != testData[0].value {
		t.FailNow()
	}
}

func TestElementAtIndex_Middle(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}, Element{"Good day!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])
	AppendElement(testData[2])

	e := ElementAtIndex(1)

	if e.value != testData[1].value {
		t.FailNow()
	}
}

func TestElementAtIndex_Last(t *testing.T) {
	testData := []Element{Element{"hello", nil}, Element{"everyone!", nil}, Element{"Good day!", nil}}
	AppendElement(testData[0])
	AppendElement(testData[1])
	AppendElement(testData[2])

	e := ElementAtIndex(2)

	if e.value != testData[2].value {
		t.FailNow()
	}
}
