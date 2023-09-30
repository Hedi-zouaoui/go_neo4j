package utils

import (
	"testing"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func TestParentNode(t *testing.T) {
	mockDriver, _ := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	defer mockDriver.Close()

	// Define constants for test values
	expectedNodeID := int64(1)
	expectedNodeName := "ahmed"

	t.Run("existing_node", func(t *testing.T) {
		id := 11
		got, err := Parent_node(mockDriver, 0 , &id )
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		want := &Node{
			ID:   expectedNodeID,
			Name: expectedNodeName,
		}

		if got.ID != want.ID || got.Name != want.Name {
			t.Errorf("Test case 'existing_node' failed. Got %+v, wanted %+v", got, want)
		}
	})

	t.Run("non_existing_parent", func(t *testing.T) {
		// Test case for non-existing node
		
		got, err := Parent_node(mockDriver, 0 , nil ) // Assuming node with ID 99 doesn't exist
		if err == nil {
			t.Fatalf("Expected error, but got nil")
		}
		if got != nil {
			t.Errorf("Expected nil node, but got %+v", got)
		}
	})
}

func TestListNodesIndirect(t *testing.T) {
	mockDriver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	defer mockDriver.Close()

	if err != nil {
		t.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer mockDriver.Close()

	// Set up your Neo4j database with appropriate nodes and relationships for testing

	// Run the function under test
	fatherNodeID := 1 // Replace with the actual node ID you want to test
	id := 30 
	nodes, err := List_nodes_indirect(mockDriver, fatherNodeID , &id )
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Perform assertions on the result
	// Define your expected nodes based on your test setup
	expectedNodes := []Node{
		{ID: 45, Name: "mariem"},
	
		{ID: 98, Name: "nacef"},
		
		
		
	}

	if len(*nodes) != len(expectedNodes) {
		t.Errorf("Expected %d nodes, but got %d", len(expectedNodes), len(*nodes))
		return
	}

	for i, expectedNode := range expectedNodes {
		if (*nodes)[i].ID != expectedNode.ID || (*nodes)[i].Name != expectedNode.Name {
			t.Errorf("Mismatched node at index %d. Got %+v, wanted %+v", i, (*nodes)[i], expectedNode)
		}
	}
}


func TestListNodesWithRelationDirect(t *testing.T) {
	mockDriver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	if err != nil {
		t.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer mockDriver.Close()

	t.Run("direct_relation", func(t *testing.T) {
		relationName := "direct"
num := 14
		nodes, err := List_nodes_with_relation_direct(mockDriver, 0 , relationName , &num )
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		expectedNodes := []Node{
			{ID: 45, Name: "mariem"},
			{ID: 98, Name: "nacef"},
			
		}

		// Compare the expected nodes with the actual nodes
		if len(*nodes) != len(expectedNodes) {
			t.Errorf("Expected %d nodes, but got %d nodes", len(expectedNodes), len(*nodes))
			return
		}

		for i, expected := range expectedNodes {
			actual := (*nodes)[i]
			if actual.ID != expected.ID || actual.Name != expected.Name {
				t.Errorf("Node at index %d - got %+v, wanted %+v", i, actual, expected)
			}
		}
	})
}