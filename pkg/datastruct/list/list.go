package list

type List struct {
	value string
	next  *List
}

// Create a way to print the list
func (list *List) String() string {
	if list == nil {
		return "(empty)"
	}

	result := ""

	for list != nil {
		if len(result) > 0 {
			result += " > "
		}
		result += list.value
		list = list.next
	}

	return result
}

// search a list and return the pointer to the value if it's found
func (list *List) Find(value string) *List {
	for list != nil {
		if list.value == value {
			return list
		}
	}
	return nil
}

// insert a node into the list
func (list *List) Insert(node *List) *List {
	if list == nil { // new list
		return node
	}
	if list.next == nil {
		list.next = node
	} else {
		node.next = list.next
		list.next = node
	}
	return list
}

// Assumes a shallow list
func (list *List) Reverse() *List {

	// forward declaration necessary for recursion
	var revRecursiveInPlace func(items *List, result *List) *List

	revRecursiveInPlace = func(list *List, result *List) *List {
		if list == nil {
			return result
		}
		temp := list.next
		list.next = result
		return revRecursiveInPlace(temp, list)
	}

	result := (*List)(nil)
	return revRecursiveInPlace(list, result)
}
