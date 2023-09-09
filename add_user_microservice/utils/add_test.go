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
	expectedNodeName := "bayrem"

	t.Run("existing_node", func(t *testing.T) {
		got, err := Parent_node(mockDriver, 3)
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


}

func TestListNodesIndirect(t *testing.T) {
	mockDriver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	if err != nil {
		t.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer mockDriver.Close()

	t.Run("existing_nodes", func(t *testing.T) {
		fatherNodeID := 13 // Replace with the actual node ID you want to test
		nodes, err := List_nodes_indirect(mockDriver, fatherNodeID)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Define expected nodes based on your test setup
		expectedNodes := []Node{
			{ID: 17, Name: "racem"},
	
			{ID: 14, Name: "amira"},
			
			
			{ID: 12, Name: "anas"},
			{ID: 10, Name: "aymen"},
			// Add more expected nodes as needed
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



func TestTestHead(t *testing.T) {
	// Mock the Neo4j driver here (you need to define a mock driver)
	mockDriver, _ := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	defer mockDriver.Close()

	// Test case 1: Relationship exists
	t.Run("relationship_exists", func(t *testing.T) {
		headNodeID := int64(1) // Replace with the actual node ID you want to test
		relationshipExists, err := Test_head(mockDriver, headNodeID)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if !relationshipExists {
			t.Errorf("unExpected to be the head")
		}
	})

	// Test case 2: Relationship does not exist
	t.Run("relationship_does_not_exist", func(t *testing.T) {
		headNodeID := int64(0) // Replace with the actual node ID you want to test
		relationshipExists, err := Test_head(mockDriver, headNodeID)
		if err != nil {
			t.Fatalf("expected error: %v", err)
		}

		if relationshipExists {
			t.Errorf("not excepted to be the head")
		}
	})
}