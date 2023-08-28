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
		got, err := Parent_node(mockDriver, 9)
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
