package main

// Element is the structure of each node in the linked list.
type Element struct {
	value string
	next  *Element
}

var start *Element
var end *Element
var length int

func main() {

}

func Init() {

}

// AppendElement appends the element passed into the end of the linked list.
func AppendElement(e Element) {
	if length == 0 {
		start = &e
		end = start
	} else {
		end.next = &e
		end = &e
	}

	length++
}

// RemoveElementAtIndex removes the element at index i. NOTE index starts from 0 (similar to an array).
func RemoveElementAtIndex(i int) {
	if !(i < length) {
		return
	}

	length--

	// if it is the first element no need to loop
	if i == 0 {
		start = start.next

		return
	}

	var prev, curr *Element
	curr = start

	var j int
	for j = 1; j <= i; j++ {
		prev = curr
		curr = curr.next
	}

	prev.next = curr.next
}

// ElementAtIndex returns the element at the specified index. NOTE: Index starts from 0
func ElementAtIndex(i int) *Element {
	if !(i < length) {
		return nil
	}

	// if it's the first element
	if i == 0 {
		return start
	}

	curr := start

	for j := 1; j <= i; j++ {
		curr = curr.next
	}

	return curr
}

// Length returns the length of the linked list
func Length() int {
	return length
}
