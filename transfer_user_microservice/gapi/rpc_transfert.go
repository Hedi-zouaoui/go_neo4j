package gapi

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j_golang/pb"
	
	utils "github.com/neo4j_golang/utils"
)


func (server *Server)Transfertuser(ctx context.Context , req *pb.TransferNode )(  *pb.TransferNode ,error ) {

	from := []utils.Node{
		{ID: req.From , Name: "x"},
	}

	to := []utils.Node{
		{ID: req.To  , Name: "x"},
	}

	// create a new Neo4j driver
	driver, err := neo4j.NewDriver("bolt://neo4j_db:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	// docker : bolt://neo4j_db:7687
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer driver.Close()

	 Ch := req.Methode
if Ch == "with" {
	/*
	t4 , _ :=  utils.Test_head(driver, int64(from[0].ID))  
	fmt.Println("test t4 from ", t4)*/
t3, _ := utils.Test_head(driver, int64(to[0].ID))
	fmt.Println("test t3", t3)
if t3  {

t, _ := utils.Parent_node(driver, int(from[0].ID))
	t1, _ := utils.Test_head(driver, int64(t.ID))

	fmt.Println("le test t1", t1)
	p2, _ := utils.Parent_node(driver, int(to[0].ID))
	t2, _ := utils.Test_head(driver, int64(p2.ID))
	fmt.Println("test t2", p2)
	if !t1 { /* le parent from d origine est le HEAD */
		fmt.Println("la tete est le HEAD PRINCIPAL de l arbre ")
		var nodes2 *[]utils.Node

		nodes2, err := utils.List_nodes_indirect(driver, int(from[0].ID ) , nil)
		if err != nil {
			log.Fatal(err)
		}

		first_direct, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
		fmt.Println(*first_direct)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("la famille bas d origine ", nodes2)
		fmt.Println("le premier direct avec la tete ", *first_direct)


		
		total := append(*first_direct, *nodes2...)
		fmt.Println("total ", total)

		head := []utils.Node{
			{ID: t.ID, Name: "x"},
		}
		err = utils.Delete_relation(driver, &total, &head)
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Delete_relation(driver, &from, &head )
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		
		parent, err := utils.Parent_node(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("le nouveau pere", *parent)
		nodes_to, err := utils.Recursion_getNodeWithDicRelation(driver, int(parent.ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut de nouveau parent  ", *nodes_to)

	
		err = utils.Create_indirect_relations(driver, nodes_to, &from)
		if err != nil {
			log.Fatal(err)
		}
			err = utils.Create_indirect_relations(driver, nodes_to, &total)
		if err != nil {
			log.Fatal(err)
		}


	}

	if !t2  { /* le parent de destination est le HEAD de l arbre  */
		fmt.Println("le parent de destination est le HEAD de l arbre")

		nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}

		first_direct, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
		fmt.Println(*first_direct)
		if err != nil {
			log.Fatal(err)
		}
		err = utils.Create_indirect_relations(driver, nodes, first_direct)

		if err != nil {
			log.Fatal(err)
		}

		nodes, err = utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut d origine ", *nodes) // la famille haut d origine

		parent, err := utils.Parent_node(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("le pere", *parent)
		nodes2, err :=utils.List_nodes_indirect(driver, int(from[0].ID) , nil)
		if err != nil {
			log.Fatal(err)
		}
		par := &[]utils.Node{}
		*par = append(*first_direct, *nodes2...)

		fmt.Println("la famille bas d origine ", *par) // la famille bas d origine

		err = utils.Create_indirect_relations(driver, &to, par)
		

		if err != nil {
			log.Fatal(err)
		}

		err =utils.Delete_relation(driver, par, nodes)
		if err != nil {
			log.Fatal(err)
		}
		err =utils.Delete_relation(driver, &from, nodes)
		if err != nil {
			log.Fatal(err)
		}

		nodeH := []utils.Node{*p2}

		nodes = &[]utils.Node{}
		*nodes = append(to, nodeH...)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut de parent ", *nodes)

		
		err = utils.Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}

		p2 := []utils.Node{
			{ID: p2.ID, Name: "x"},
		}
/*
		
		   err = create_indirect_relations(driver , &to , par )

		    if err != nil {
		   	log.Fatal(err)
		   }*/
		err = utils.Create_indirect_relations(driver, &p2, &from)

		if err != nil {
			log.Fatal(err)
		}
		err = utils.Create_indirect_relations(driver, &p2, par)
			if err != nil {
			log.Fatal(err)
		}

	}

	if t1  && t2  && t3 {
		fmt.Println("normal case ")

		fmt.Println("else ")
		if err != nil {
			log.Fatal(err)
		}

		first_direct, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
		fmt.Println(*first_direct)
		if err != nil {
			log.Fatal(err)
		}

		nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut d origine ", *nodes) // la famille haut d origine

		parent, err := utils.Parent_node(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("le pere", *parent)
		nodes2, err := utils.List_nodes_indirect(driver, int(from[0].ID) , nil)
		if err != nil {
			log.Fatal(err)
		}
		par := &[]utils.Node{}
		*par = append(*first_direct, *nodes2...)

		fmt.Println("la famille bas d origine ", *par) // la famille bas d origine

		err = utils.Delete_relation(driver, par, nodes)
		if err != nil {
			log.Fatal(err)
		}

		nodes, err = utils.Recursion_getNodeWithDicRelation(driver, int(parent.ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut de parent ", *nodes)

		err = utils.Delete_relation(driver, nodes, &from)
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Delete_relation(driver, nodes, par)
		if err != nil {
			log.Fatal(err)
		}
		/*
		
		t3h , _ := parent_node(driver ,int( to[0].id))*/
		t3H := []utils.Node{
			{ID: parent.ID , Name: "x"},
		}
			err = utils.Delete_relation(driver, &from, &t3H)
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		parent, err = utils.Parent_node(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("le nouveau pere", *parent)
		nodes_to, err := utils.Recursion_getNodeWithDicRelation(driver, int(parent.ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut de nouveau parent  ", *nodes)

		err = utils.Create_indirect_relations_single_to_many(driver, nodes_to, int(from[0].ID))

		if err != nil {
			log.Fatal(err)
		}

		nodes, err = utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la nouvelle famille haut d origine ", *nodes)

		err = utils.Create_indirect_relations(driver, nodes, par)

		if err != nil {
			log.Fatal(err)
		}

	}

}else  { /* to est le head de l arbre */
		fmt.Println(" to est le head de l arbre  ")
		nodes2, err := utils.List_nodes_indirect(driver, int(from[0].ID) , nil)
		if err != nil {
			log.Fatal(err)
		}

		first_direct, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
		fmt.Println(*first_direct)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("la famille bas d origine ", *nodes2)
		fmt.Println("le premier direct avec la tete ", *first_direct)

		total := &[]utils.Node{}
		*total = append(*first_direct, *nodes2...)
		fmt.Println("total de la famille bas d origine ", *total)
		nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut d origine ", *nodes) // la famille haut d origine

		err = utils.Delete_relation(driver, total, nodes)
		if err != nil {
			log.Fatal(err)
		}
		err =utils.Delete_relation(driver, nodes, &from)
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Create_indirect_relations(driver, &to, total)

		if err != nil {
			log.Fatal(err)
		}

	}

}
if Ch == "without"{
t3, _ := utils.Test_head(driver, int64(to[0].ID))
	fmt.Println("test t3", t3)
if t3 {

t, _ := utils.Parent_node(driver, int(from[0].ID))
	t1, _ := utils.Test_head(driver, int64(t.ID))

	fmt.Println("le test t1", t1)
	p2, _ := utils.Parent_node(driver, int(to[0].ID))
	t2, _ := utils.Test_head(driver, int64(p2.ID))
	fmt.Println("test t2", p2)
	if !t1  { /* le parent from d origine est le HEAD */
			fmt.Println("la tete est le HEAD PRINCIPAL de l arbre WITHOUT CHILDREN ")
		nodes2, err := utils.List_nodes_indirect(driver, int(from[0].ID) , nil)
		if err != nil {
			log.Fatal(err)
		}

		first_direct, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
		fmt.Println(*first_direct)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("la famille bas d origine ", *nodes2)
		fmt.Println("le premier direct avec la tete ", *first_direct)

		total := &[]utils.Node{}
		*total = append(*first_direct, *nodes2...)
		fmt.Println("total ", *total)

			err = utils.Delete_relation(driver, total, &from)
		if err != nil {
			log.Fatal(err)
		}
				nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
			err = utils.Delete_relation(driver, nodes , &from)
		if err != nil {
			log.Fatal(err)
		}

			head := []utils.Node{
			{ID: t.ID, Name: "x"},
		}
				err = utils.Delete_relation(driver, &head , &from)
		if err != nil {
			log.Fatal(err)
		}
			err = utils.Delete_relation(driver, &head, first_direct)
		if err != nil {
			log.Fatal(err)
		}
err = utils.Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		
		err = utils.Create_direct_relations(driver , &head , first_direct)
		if err != nil {
			log.Fatal(err)
		}
	
		f_to, err := utils.Recursion_getNodeWithDicRelation(driver , int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}

		err = utils.Create_indirect_relations(driver , f_to , &from)
if err != nil {
			log.Fatal(err)
		}

	}



if !t2 { /* le parent de destination est le HEAD de l arbre  */
		fmt.Println("le parent de destination est le HEAD de l arbre without children ")
		nodes2, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
fmt.Println(nodes2)
		first_direct, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
		fmt.Println(*first_direct)
		if err != nil {
			log.Fatal(err)
		}
nodes3, err := utils.List_nodes_indirect(driver, int(from[0].ID) , nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("la famille haut d origine ", *nodes2)
		fmt.Println("le premier direct avec la tete ", *first_direct)

		total := &[]utils.Node{}
		*total = append(*first_direct, *nodes3...)
		*total = append(*total, *nodes2...)

		fmt.Println("total ", *total)

			err = utils.Delete_relation(driver, total, &from)
		if err != nil {
			log.Fatal(err)
		}
			head := []utils.Node{
			{ID: t.ID, Name: "x"},
		}
				err = utils.Delete_relation(driver, &head , &from)
		if err != nil {
			log.Fatal(err)
		}
			err = utils.Delete_relation(driver, &head, first_direct)
		if err != nil {
			log.Fatal(err)
		}
err = utils.Move_from_to_direct(driver, int(from[0].ID), int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}

err = utils.Create_direct_relations(driver , &head , first_direct)
	if err != nil {
			log.Fatal(err)
		}


  
		f_to := []utils.Node{
			{ID: p2.ID, Name: "x"},
	
		}
err = utils.Create_indirect_relations(driver , &f_to , &from)
	if err != nil {
			log.Fatal(err)
		}




}
if t1  && t2  && t3  { 
fmt.Println("normal case without  ")
/*
nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		fmt.Println("else ")
		if err != nil {
			log.Fatal(err)
		}

*/


		firstDirect, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
fmt.Println(*firstDirect)
if err != nil {
    log.Fatal(err)
}
parent, err := utils.Parent_node(driver, int(from[0].ID))
if err != nil {
    log.Fatal(err)
}
fmt.Println("le pere", *parent)




		nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("recursion" , nodes  )
		fmt.Println("la famille haut d origine ", *nodes) // la famille haut d origine

		
		nodes2, err := utils.List_nodes_indirect(driver, int(from[0].ID) , nil)
		if err != nil {
			log.Fatal(err)
		}
		par := &[]utils.Node{}
		*par = append(*firstDirect, *nodes2...)

		fmt.Println("la famille bas d origine ", *par) // la famille bas d origine

		if len(*par) > 0 {
			fmt.Println("la famille bas d'origine ", *par) // Print the family
		
			err := utils.Delete_relation(driver, par, &from)
			if err != nil {
				log.Fatal(err)
			}
		}
		err = utils.Delete_relation(driver, nodes , &from)
		if err != nil {
			log.Fatal(err)
		}
		//var parentSlice *[]utils.Node
parentSlice := &[]utils.Node{*parent}
err = utils.Delete_relation(driver, firstDirect, parentSlice)
		if err != nil {
			log.Fatal(err)
		}
err = utils.Move_from_to_direct(driver  , int(from[0].ID), int(to[0].ID))
	if err != nil {
			log.Fatal(err)
		}
	head := []utils.Node{
			{ID: (*firstDirect)[0].ID , Name: "x"},
		}
fmt.Println(head)

if len(*firstDirect) > 0 {
    toNodeID := int((*firstDirect)[0].ID)
    err = utils.Move_from_to_direct(driver, toNodeID, int(parent.ID))
    if err != nil {
        log.Fatal(err)
    }
}
destination , _  := utils.Parent_node(driver , int(to[0].ID) )
nodes, err = utils.Recursion_getNodeWithDicRelation(driver, int(to[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("recursion" , nodes  )
		fmt.Println("la famille haut de dierect ", *nodes) // la famille haut d to
	
		fmt.Println("le nouveau pere" , destination)
		
err = utils.Create_indirect_relations(driver, nodes , &from )
		if err != nil {
			log.Fatal(err)
		}
		
		nodes3, err := utils.List_nodes_indirect(driver, int(from[0].ID) , nil)
		if err != nil {
			log.Fatal(err)
		}
		firstDirect, err = utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
			if err != nil {
			log.Fatal(err)
		}
			par = &[]utils.Node{}
		*par = append(*firstDirect, *nodes3...)
		utils.Create_indirect_relations(driver , nodes , par)
}




}



if !t3 {
	// to est le head de l arbre 
	fmt.Println(" to est le head de l arbre  ")

	/*
nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		fmt.Println("else ")
		if err != nil {
			log.Fatal(err)
		}
*/



		firstDirect, err := utils.List_nodes_with_relation(driver, int64(from[0].ID), nil, "direct")
fmt.Println(*firstDirect)
if err != nil {
    log.Fatal(err)
}
parent, err :=utils.Parent_node(driver, int(from[0].ID))
if err != nil {
    log.Fatal(err)
}
fmt.Println("le pere", *parent)




		nodes, err := utils.Recursion_getNodeWithDicRelation(driver, int(from[0].ID))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("recursion" , nodes  )
		fmt.Println("la famille haut de tous ", *nodes) // la famille haut d origine

		
		nodes2, err := utils.List_nodes_indirect(driver, int(from[0].ID) , nil )
		if err != nil {
			log.Fatal(err)
		}
		par := &[]utils.Node{}
		*par = append(*firstDirect, *nodes2...)

		fmt.Println("la famille bas d origine ", *par) // la famille bas d origine

		err = utils.Delete_relation(driver , &from , par)
			if err != nil {
			log.Fatal(err)
		}
		err = utils.Delete_relation(driver , &from , nodes)
			if err != nil {
			log.Fatal(err)
		}
		head := []utils.Node{
			{ID: parent.ID, Name: "x"},
		}
		err = utils.Create_direct_relations(driver , &head , firstDirect)
			if err != nil {
			log.Fatal(err)
		}
		err = utils.Create_direct_relations(driver , &from, &to )
			if err != nil {
			log.Fatal(err)
		}
		

}

}


return  req , err
}


