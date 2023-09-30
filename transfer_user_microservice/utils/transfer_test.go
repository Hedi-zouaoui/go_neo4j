package utils 



import (
	"testing"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)


func TestTest_head(t *testing.T) {
	// Create an in-memory Neo4j instance for testing
	mockDriver, _ := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	defer mockDriver.Close()

	id := 3 
	nodes , err := List_nodes_indirect(mockDriver, 0 , &id )

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Define expected nodes based on your test setup
	expectedNodes := []Node{
		{ID: 44, Name: "salim"},
	
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
}
