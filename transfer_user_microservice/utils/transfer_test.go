package utils 



import (
	"testing"

)


func TestTest_head(t *testing.T) {
	// Create an in-memory Neo4j instance for testing
	server := dbtest.NewDatabase()

	// Connect to the test database
	driver, err := neo4j.NewDriver(server.URI(), neo4j.BasicAuth("neo4j", "12345678", ""))
	if err != nil {
		t.Fatalf("Failed to connect to Neo4j: %v", err)
	}
	defer driver.Close()

	// Create a node with a direct relationship in the test d
	session := driver.NewSession(neo4j.SessionConfig{})
	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		_, err := tx.Run("CREATE (node1)-[:direct]->(node2) RETURN node1, node2", nil)
		return nil, err
	})
	session.Close()

	// Test the Test_head function
	headNodeID := int64(0) // Use the node ID you 
	relationshipExists, err := Test_head(driver, headNodeID)

	// Check for errors
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// Check if the relationship exists as expected
	if !relationshipExists {
		t.Error("Expected relationship to exist, but it does not")
	}
}