package linkedList

import (
	"testing"
)

func TestLinkedList_IsEmpty(t *testing.T) {
	testList := NewLinkedList[int]()

	if !testList.IsEmpty() {
		t.Errorf("Expected linked list to be empty, but it is not")
	}

	testList.AddNode(5)
	if testList.IsEmpty() {
		t.Errorf("Expected linked list to not be empty after adding a node, but it is")
	}
}

func TestLinkedList_AddNode(t *testing.T) {
	testList := NewLinkedList[int]()
	testList.AddNode(1)
	testList.AddNode(2)
	testList.AddNode(3)

	if testList.Length != 3 {
		t.Errorf("Expected linked list length to be 3, got %d", testList.Length)
	}

	if testList.First.Value != 1 {
		t.Errorf("Expected first node value to be 1, got %d", testList.First.Value)
	}

	if testList.First.GetNext().Value != 2 {
		t.Errorf("Expected second node value to be 2, got %d", testList.First.GetNext().Value)
	}
}

func TestLinkedList_AddAllNodes(t *testing.T) {
	testList := NewLinkedList[int]()
	testList.AddAllNodes([]int{1, 2, 3})

	if testList.Length != 3 {
		t.Errorf("Expected linked list length to be 3 after adding all nodes, got %d", testList.Length)
	}

	if testList.First.Value != 1 {
		t.Errorf("Expected first node value to be 1 after adding all nodes, got %d", testList.First.Value)
	}

	if testList.First.GetNext().Value != 2 {
		t.Errorf("Expected second node value to be 2 after adding all nodes, got %d", testList.First.GetNext().Value)
	}

	if testList.First.GetNext().GetNext().Value != 3 {
		t.Errorf("Expected third node value to be 3 after adding all nodes, got %d", testList.First.GetNext().GetNext().Value)
	}
}

func TestLinkedList_AppendList(t *testing.T) {
	testList := NewLinkedList[int]()
	otherList := NewLinkedList[int]()

	testList.AddAllNodes([]int{1, 2, 3})
	otherList.AddAllNodes([]int{4, 5, 6})

	testList.AppendList(&otherList)

	if testList.Length != 6 {
		t.Errorf("Expected linked list length to be 6 after appending, got %d", testList.Length)
	}

	if testList.First.Value != 1 {
		t.Errorf("Expected first node value to be 1 after appending, got %d", testList.First.Value)
	}

	if testList.First.GetNext().Value != 2 {
		t.Errorf("Expected second node value to be 2 after appending, got %d", testList.First.GetNext().Value)
	}

	if testList.First.GetNext().GetNext().Value != 3 {
		t.Errorf("Expected third node value to be 3 after appending, got %d", testList.First.GetNext().GetNext().Value)
	}

	if testList.First.GetNext().GetNext().GetNext().Value != 4 {
		t.Errorf("Expected fourth node value to be 4 after appending, got %d", testList.First.GetNext().GetNext().GetNext().Value)
	}
}

func TestLinkedList_GetNode(t *testing.T) {
	testList := NewLinkedList[int]()
	testList.AddNode(1)
	testList.AddNode(2)
	testList.AddNode(3)

	node := testList.GetNode(2)
	if node == nil || node.Value != 2 {
		t.Errorf("Expected to find node with value 2, got %v", node)
	}

	node = testList.GetNode(4)
	if node != nil {
		t.Errorf("Expected to not find node with value 4, got %v", node)
	}
}

func TestLinkedList_DropFirstNode(t *testing.T) {
	testList := NewLinkedList[int]()
	testList.AddNode(1)
	testList.AddNode(2)
	testList.AddNode(3)

	if testList.Length != 3 {
		t.Errorf("Expected linked list length to be 3, got %d", testList.Length)
	}

	testList.DropFirstNode()

	if testList.Length != 2 {
		t.Errorf("Expected linked list length to be 2 after dropping first node, got %d", testList.Length)
	}

	if testList.First.Value != 2 {
		t.Errorf("Expected first node value to be 2 after dropping first node, got %d", testList.First.Value)
	}

	if testList.First.GetNext().Value != 3 {
		t.Errorf("Expected second node value to be 3 after dropping first node, got %d", testList.First.GetNext().Value)
	}
}

func TestLinkedList_RemoveNode(t *testing.T) {
	testList := NewLinkedList[int]()
	testList.AddNode(1)
	testList.AddNode(2)
	testList.AddNode(3)

	node := testList.GetNode(2)
	if node == nil || node.Value != 2 {
		t.Errorf("Expected to find node with value 2, got %v", node)
	}

	// Assuming a RemoveNode method exists
	testList.RemoveNode(2)

	node = testList.GetNode(2)
	if node != nil {
		t.Errorf("Expected to not find node with value 2 after removal, got %v", node)
	}

	if testList.Length != 2 {
		t.Errorf("Expected linked list length to be 2 after removal, got %d", testList.Length)
	}
}

func TestLinkedList_GetNodeByIndex(t *testing.T) {
	testList := NewLinkedList[int]()
	testList.AddNode(1)
	testList.AddNode(2)
	testList.AddNode(3)

	node := testList.GetNodeByIndex(0)
	if node == nil || node.Value != 1 {
		t.Errorf("Expected to find node at index 0 with value 1, got %v", node)
	}

	node = testList.GetNodeByIndex(1)
	if node == nil || node.Value != 2 {
		t.Errorf("Expected to find node at index 1 with value 2, got %v", node)
	}

	node = testList.GetNodeByIndex(2)
	if node == nil || node.Value != 3 {
		t.Errorf("Expected to find node at index 2 with value 3, got %v", node)
	}

	node = testList.GetNodeByIndex(3)
	if node != nil {
		t.Errorf("Expected to not find a node at index 3, got %v", node)
	}
}
