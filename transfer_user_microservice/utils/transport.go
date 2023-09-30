package utils

import (
	"fmt"

"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)





type Node struct {
	ID   int64 `json:"id"`
	Name string `json:"name"`
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

func List_nodes_indirect(driver neo4j.Driver, father int , father_test *int) (*[]Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	if father_test != nil {
		result, err := session.Run("MATCH (n2)-[:indirect]->(n1) WHERE n1.IdAd = $father   RETURN n2.IdAd AS id, n2.name AS name", map[string]interface{}{

			"father": father_test,
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
			return &nodes, nil
	}}else{
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
		nodes = append(nodes, node)}
	
	
	
	return &nodes, nil
	}
	return nil , nil 
}
func List_nodes_with_relation(driver neo4j.Driver, to_change int64 , father *int64, relation_name string) (*[]Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	var result neo4j.Result
	var err error
	if relation_name == "indirect" {
		result, err = session.Run("MATCH (n2)-[:indirect]->(n1), (n2)-[:indirect]->(n3) WHERE id(n1) = $to_change AND id(n3) = $father RETURN id(n2) AS id, n2.name AS name", map[string]interface{}{
			"to_change":     to_change,
			"father":        *father,
			"relation_name": relation_name,
		})

		if err != nil {
			return nil, err
		}
	} else {
		result, err = session.Run("MATCH (child)-[r1:direct]->(node) WHERE id(node) = $to_change RETURN id(child) AS id , child.name AS name ", map[string]interface{}{
			"to_change": to_change,
		})
		if err != nil {
			return nil, err
		}

	}
	nodes := []Node{}
	for result.Next() {
		record := result.Record()
		 ID , ok := record.Get("id")
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
				"start_node": node2.ID		,
				"end_node":   node.ID		,
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
