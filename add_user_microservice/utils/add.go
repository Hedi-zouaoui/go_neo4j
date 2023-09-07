package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Node struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type Transfer_node struct {
	To      int64
	From    int64
	Methode string
}

func CreateNode(driver neo4j.Driver, name string) (*Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run(`CREATE (node:Node {name: $name }) RETURN id(node) as id, node.name as name`, map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return nil, err
	}

	if result.Next() {
		record := result.Record()
		idValue, exists := record.Get("id")
		if !exists {
			return nil, fmt.Errorf("failed to retrieve ID")
		}
		nameValue, exists := record.Get("name")
		if !exists {
			return nil, fmt.Errorf("failed to retrieve name")
		}

		id, ok := idValue.(int64)
		if !ok {
			return nil, fmt.Errorf("failed to convert ID to int64")
		}
		name, ok := nameValue.(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert name to string")
		}

		nodeObj := &Node{
			ID:   id,
			Name: name,
		}
		fmt.Println(nodeObj)
		return nodeObj, nil
	}

	return nil, fmt.Errorf("failed to create node")
}

func TransferNode(driver neo4j.Driver, req Transfer_node) Transfer_node {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	from := []Node{
		{ID: req.From, Name: "x"},
	}

	to := []Node{
		{ID: req.To, Name: "x"},
	}

	t4, _ := Test_head(driver, int64(to[0].ID))
	fmt.Println("test t4", t4)
	if !t4 {
		fmt.Println("le to est le head de l arbre ")
		err := Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		return req

	}
	p2, _ := Parent_node(driver, int(to[0].ID))
	t2, _ := Test_head(driver, int64(p2.ID))
	fmt.Println("test t2", p2)
	if !t2 { /* le parent de TO est le HEAD de l arbre  */
		fmt.Println("le parent de destination est le HEAD de l arbre")

		err := Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}

		p2 := []Node{
			{ID: p2.ID, Name: p2.Name},
		}
		err = Create_indirect_relations(driver, &p2, &from)

		if err != nil {
			log.Fatal(err)
		}
		return req
	}

	if t4 && t2 {

		fmt.Println("normal case ")

		err := Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		parent, err := Parent_node(driver, int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("le pere", *parent)
		nodes_to, err := Recursion_getNodeWithDicRelation(driver, int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut de nouveau parent  ", nodes_to)

		err = Create_indirect_relations(driver, nodes_to, &from)

		if err != nil {
			log.Fatal(err)
		}

	}
	return req

}

func Test_head(driver neo4j.Driver, head int64) (bool, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run(`
		MATCH (node1)-[r:direct]->()
		WHERE id(node1) = $nodeID
		RETURN COUNT(r) > 0 AS relationship_exists
	`, map[string]interface{}{
		"nodeID": head,
	})
	if err != nil {
		return false, err
	}

	if result.Next() {
		record := result.Record()
		relationshipExists, ok := record.Get("relationship_exists")
		if !ok {
			return false, fmt.Errorf("failed to retrieve relationship existence")
		}

		return relationshipExists.(bool), nil
	}

	return false, nil // Relationship not found

}

func List_nodes_indirect(driver neo4j.Driver, father int) (*[]Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.Run("MATCH (n2)-[:indirect]->(n1) WHERE id(n1) = $father   RETURN id(n2) AS id, n2.name AS name", map[string]interface{}{

		"father": father,
	})

	if err != nil {
		return nil, err
	}
	nodes := []Node{}
	for result.Next() {
		record := result.Record()
		ID, ok := record.Get("id")
		if !ok {
			return nil, fmt.Errorf("failed to retrieve node id")
		}

		Name, ok := record.Get("name")
		if !ok {
			return nil, fmt.Errorf("failed to retrieve node name")
		}

		node := Node{
			ID:   ID.(int64),
			Name: Name.(string),
		}
		nodes = append(nodes, node)
	}

	return &nodes, nil

}
func Delete_relation(driver neo4j.Driver, start_node *[]Node, end_nodes *[]Node) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	for _, node := range *end_nodes {
		for _, node2 := range *start_node {
			_, err := session.Run("MATCH (node1) WHERE id(node1) = $end_node MATCH (node2) WHERE id(node2) = $start_node MATCH (node1)-[r]-(node2) DELETE r ", map[string]interface{}{
				"start_node": node2.ID,
				"end_node":   node.ID,
			})
			if err != nil {
				return err
			}

		}
	}
	return nil

}
func Create_indirect_relations(driver neo4j.Driver, target_nodes *[]Node, end_nodes *[]Node) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	for _, node := range *end_nodes {
		for _, p_nodes := range *target_nodes {
			_, err := session.Run("MATCH (node1) WHERE id(node1) = $end_node MATCH (node2) WHERE id(node2) = $target_node CREATE (node1)-[:indirect]->(node2) RETURN node1, node2  ", map[string]interface{}{
				"target_node": p_nodes.ID,
				"end_node":    node.ID,
			})

			if err != nil {
				return err
			}
		}
	}

	return nil

}
func Create_direct_relations(driver neo4j.Driver, target_nodes *[]Node, end_nodes *[]Node) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	for _, node := range *end_nodes {
		for _, p_nodes := range *target_nodes {
			_, err := session.Run("MATCH (node1) WHERE id(node1) = $end_node MATCH (node2) WHERE id(node2) = $target_node CREATE (node1)-[:direct]->(node2) RETURN node1, node2  ", map[string]interface{}{
				"target_node": p_nodes.ID,
				"end_node":    node.ID,
			})

			if err != nil {
				return err
			}
		}
	}

	return nil

}

func Create_indirect_relations_single_to_many(driver neo4j.Driver, target_nodes *[]Node, end_node int) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	for _, p_nodes := range *target_nodes {
		_, err := session.Run("MATCH (node1) WHERE id(node1) = $end_node MATCH (node2) WHERE id(node2) = $target_node CREATE (node1)-[:indirect]->(node2) RETURN node1, node2  ", map[string]interface{}{
			"target_node": p_nodes.ID,
			"end_node":    end_node,
		})

		if err != nil {
			return err
		}
	}

	return nil

}

func Move_from_to_direct(driver neo4j.Driver, from_node int, to_node int) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.Run("MATCH (node1), (node2) WHERE id(node1) = $from_node AND id(node2) = $to_node CREATE (node1)-[:direct]->(node2) ", map[string]interface{}{
		"from_node": from_node,
		"to_node":   to_node,
	})
	if err != nil {
		return err
	}
	return nil
}

func Parent_node(driver neo4j.Driver, from_node int) (*Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.Run("MATCH (node)-[:direct]->(node1) WHERE id(node) = $from_node RETURN id(node1) AS id, node1.name AS name", map[string]interface{}{
		"from_node": from_node,
	})
	if err != nil {
		return nil, err
	}

	if result.Next() {
		record := result.Record()

		ID, ok := record.Get("id")
		if !ok {
			return nil, fmt.Errorf("failed to retrieve node id")
		}

		Name, ok := record.Get("name")
		if !ok {
			return nil, fmt.Errorf("failed to retrieve node name")
		}

		node := &Node{
			ID:   ID.(int64),
			Name: Name.(string),
		}
		return node, nil
	}

	return nil, fmt.Errorf("node not found")
}

func GetNodeWithDirectRelationship(driver neo4j.Driver, from_node int) (*Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := session.Run("MATCH (node)-[:direct]->(node1) WHERE id(node) = $from_node  RETURN id(node1) AS id, node1.name AS name ", map[string]interface{}{
		"from_node": from_node,
	})

	if err != nil {
		return nil, err
	}

	if result.Next() {
		record := result.Record()
		ID, ok := record.Get("id")
		if !ok {
			return nil, fmt.Errorf("failed to retrieve node id")
		}

		Name, ok := record.Get("name")
		if !ok {
			return nil, fmt.Errorf("failed to retrieve node name")
		}

		node := &Node{
			ID:   ID.(int64),
			Name: Name.(string),
		}

		return node, nil
	}
	p, _ := Parent_node(driver, from_node)
	return p, nil // No node found with direct relationship
}

func Recursion_getNodeWithDicRelation(driver neo4j.Driver, from_node int) (*[]Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	result, err := GetNodeWithDirectRelationship(driver, from_node)
	fmt.Println(result)
	if err != nil {
		return nil, err
	}

	if result != nil {
		nodes := []Node{*result}
		res2, err := Recursion_getNodeWithDicRelation(driver, int(result.ID))
		if err != nil {
			return nil, err
		}

		if res2 != nil {
			nodes = append(nodes, *res2...)
		}
		return &nodes, nil
	}

	return nil, nil

}

type node struct {
	Name     string `json:"name"`
	Value    Value  `json:"value"`
	Children []node `json:"children"`
}

type Value struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func Add_json(ParentID int, name string, id int) {
	// Read the existing JSON data from the 
	jsonFile := "/home/test.json"
	data, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Unmarshal the JSON data into a slice of node structs
	var nodes []node
	if err := json.Unmarshal(data, &nodes); err != nil {
		log.Fatalf("Error unmarshaling JSON data: %v", err)
	}

	// Example: Create a new child node and add it to the parent with ID 3
	parentID := ParentID
	newChild := node{
		Name:     name,
		Value:    Value{Name: name, ID: id},
		Children: []node{
			// Add more children here if needed
		},
	}

	if err := AddChildToNode(&nodes, parentID, newChild); err != nil {
		log.Fatalf("Error adding child to node: %v", err)
	}

	// Marshal the updated data back to JSON
	updatedData, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON data: %v", err)
	}

	// Write the updated JSON data back to the file
	if err := ioutil.WriteFile(jsonFile, updatedData, 0644); err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	fmt.Println("Child added and JSON file updated successfully.")
}

// AddChildToNode adds a new child node to the node with the specified ID.
func AddChildToNode(nodes *[]node, nodeID int, child node) error {
	for i, node := range *nodes {
		if node.Value.ID == nodeID {
			(*nodes)[i].Children = append(node.Children, child)
			return nil
		}
		// Recursively search for the parent node within children
		if err := AddChildToNode(&((*nodes)[i].Children), nodeID, child); err == nil {
			return nil
		}
	}
	return fmt.Errorf("node with ID %d not found", nodeID)
}
