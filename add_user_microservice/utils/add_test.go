package utils

import (
	"testing"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func TestParentNode(t *testing.T) {
	mockDriver, _ := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	defer mockDriver.Close()

	// Define constants for test values
	expectedNodeID := int64(6)
	expectedNodeName := "amine"

	t.Run("existing_node", func(t *testing.T) {
		idAd := 4
		got, err := Parent_node(mockDriver, 0 , &idAd)
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
    // Set up the Neo4j driver
    mockDriver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
    if err != nil {
        t.Fatalf("Failed to create Neo4j driver: %v", err)
    }
    defer mockDriver.Close()

    // Run the test
    t.Run("existing_nodes", func(t *testing.T) {
        testListNodesIndirect(t, mockDriver)
    })
}

func testListNodesIndirect(t *testing.T, driver neo4j.Driver) {
    fatherNodeID := 30 // Replace with the actual node ID you want to test
    nodes, err := List_nodes_indirect(driver, 0, &fatherNodeID)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    // Define expected nodes based on your test setup
    expectedNodes := []Node{
        {ID: 45, Name: "mariem"},
		{ID: 98, Name: "nacef"},
        
        // Add more expected nodes as needed
    }

    // Compare the expected nodes with the actual nodes
    if len(*nodes) != len(expectedNodes) {
        t.Errorf("Expected %d nodes, but got %d nodes", len(expectedNodes), len(*nodes))
        switch {
        case len(*nodes) < len(expectedNodes):
            t.Errorf("Missing nodes in actual result")
        case len(*nodes) > len(expectedNodes):
            t.Errorf("Extra nodes in actual result")
        }
        return
    }

    for i, expected := range expectedNodes {
        actual := (*nodes)[i]
        if actual.ID != expected.ID || actual.Name != expected.Name {
            t.Errorf("Node at index %d - got %+v, wanted %+v", i, actual, expected)
            switch {
            case actual.ID != expected.ID:
                t.Errorf("Mismatch in Node ID")
            case actual.Name != expected.Name:
                t.Errorf("Mismatch in Node Name")
            }
        }
    }
}




func TestTestHead(t *testing.T) {
	// Mock the Neo4j driver here (you need to define a mock driver)
	mockDriver, _ := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	defer mockDriver.Close()

	// Test case 1: Relationship exists
	t.Run("relationship_exists", func(t *testing.T) {
		headNodeID := int64(0) // Replace with the actual node ID you want to test
		head := 38
		relationshipExists, err := Test_head(mockDriver, headNodeID , &head )
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
		head := 1
		relationshipExists, err := Test_head(mockDriver, headNodeID , &head)
		if err != nil {
			t.Fatalf("expected error: %v", err)
		}

		if relationshipExists {
			t.Errorf("not excepted to be the head")
		}
	})
}