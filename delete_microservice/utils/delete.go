package utils

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Node struct {
	ID   int64 `json:"id"`
	Name string `json:"name"`
}



func List_nodes_with_relation_direct(driver neo4j.Driver, to_change int64 ,  relation_name string , to_change_test *int ) (*[]Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	var result neo4j.Result
	var err error
	if to_change_test != nil {  fmt.Println("relation_name" , relation_name  )
	fmt.Println("to_change" , to_change  )
	
	
	
		fmt.Println("relation_name" , relation_name  )
		result, err = session.Run("MATCH (child)-[r1:direct]->(node) WHERE node.IdAd = $to_change RETURN child.IdAd AS id , child.name AS name ", map[string]interface{}{
			"to_change": to_change_test ,
		})
		if err != nil {
			return nil, err
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
fmt.Println("list_nodes_with_relation : " , &nodes) 
return &nodes, nil

	}else {
	fmt.Println("relation_name" , relation_name  )
	fmt.Println("to_change" , to_change  )
	
	
	
		fmt.Println("relation_name" , relation_name  )
		result, err = session.Run("MATCH (child)-[r1:direct]->(node) WHERE id(node) = $to_change RETURN id(child) AS id , child.name AS name ", map[string]interface{}{
			"to_change": to_change,
		})
		if err != nil {
			return nil, err
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
fmt.Println("list_nodes_with_relation : " , &nodes)   
	return &nodes, nil } 
}

func List_nodes_indirect(driver neo4j.Driver, father int , father_test *int) (*[]Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	if father_test != nil {result, err := session.Run("MATCH (n2)-[:indirect]->(n1) WHERE n1.IdAd  = $father   RETURN n2.IdAd  AS id, n2.name AS name", map[string]interface{}{

		"father": father_test ,
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
		fmt.Println("list_nodes_indirect" , &nodes)
	return &nodes, nil
		
		}else{result, err := session.Run("MATCH (n2)-[:indirect]->(n1) WHERE id(n1) = $father   RETURN id(n2) AS id, n2.name AS name", map[string]interface{}{

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
	fmt.Println("list_nodes_indirect" , &nodes)
	return &nodes, nil
	}


}

func Parent_node(driver neo4j.Driver, from_node int , from_node_test *int) (*Node, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	if from_node_test != nil {
		result, err := session.Run("MATCH (node)-[:direct]->(node1) WHERE node.IdAd  = $from_node RETURN node1.IdAd  AS id, node1.name AS name", map[string]interface{}{
			"from_node": from_node_test ,
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
			fmt.Println("parent_node function : " , node)
			return node, nil
		}
		
	}else {
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
		fmt.Println("parent_node function : " , node)
		return node, nil
	}
	}
	return nil, fmt.Errorf("node not found")
}







func DeleteNode(driver neo4j.Driver, ID int64) (int64, error) {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
direct_kids , err := List_nodes_with_relation_direct(driver , ID  , "direct" , nil )
if err != nil {
		return 0, err
	}
	fmt.Println("l id a supprimer  " , ID)
fmt.Println("le direct avec l utilisateur a supprimer " , direct_kids)
indirect_kids , err := List_nodes_indirect(driver , int(ID) , nil )
fmt.Println("indirect kids ",indirect_kids)

parent , err := Parent_node(driver , int(ID) , nil )
	
	parents := []*Node{
		{
			ID:   parent.ID,
			Name: parent.Name,
		},
	}
fmt.Println("le pere de la tete est " , parent )
err = Delete_relation(driver , parents , direct_kids)
err = Create_direct_relations(driver , parents , direct_kids)

	result, err := session.Run(`
		MATCH (n)
		WHERE ID(n) = $nodeId
		DETACH DELETE n
		RETURN $nodeId AS id
	`, map[string]interface{}{
		"nodeId": ID,
	})
	if err != nil {
		return 0, err
	}

	if result.Next() {
		record := result.Record()
		idValue, exists := record.Get("id")
		if !exists {
			return 0, fmt.Errorf("failed to retrieve ID")
		}

		id, ok := idValue.(int64)
		if !ok {
			return 0, fmt.Errorf("failed to convert ID to int64")
		}
fmt.Println("Delete node function" , id)
		return id, nil
	}

	return 0, fmt.Errorf("failed to delete node")
}



func Delete_relation(driver neo4j.Driver, start_node []*Node, end_nodes *[]Node) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()

	for _, node := range *end_nodes {
		for _, node2 := range start_node {
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
func Create_direct_relations(driver neo4j.Driver, target_nodes []*Node, end_nodes *[]Node) error {
	session := driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	for _, node := range *end_nodes {
		for _, p_nodes := range target_nodes {
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
	