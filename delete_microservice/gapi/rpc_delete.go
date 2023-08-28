package gapi

import (
	"context"
	"fmt"
	"log"

	"github.com/delete_user/pb"
	"github.com/delete_user/utils"
	"github.com/streadway/amqp"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func (server *Server) Deleteuser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// Create a new Neo4j driver
	driver, err := neo4j.NewDriver("bolt://neo4j_db:7687", neo4j.BasicAuth("neo4j", "12345678", ""))
	// docker : bolt://neo4j_db:7687
	if err != nil {
		log.Fatalf("Failed to create Neo4j driver: %v", err)
	}
	defer driver.Close()
	node, err := utils.DeleteNode(driver, req.Id)
	if err != nil {
		response := &pb.DeleteUserResponse{
			Id: 0,
			Err:   err.Error(),
		}
		return response, nil
	}
	response := &pb.DeleteUserResponse{
		Id: node,
		
		Err:   "",
	}


	// Send a RabbitMQ notification
	err = sendRabbitMQNotification(req.Id)
	if err != nil {
		log.Printf("Failed to send RabbitMQ notification: %v", err)
	}

	
	return response, nil
}


func sendRabbitMQNotification(userId int64) error {
	// Establish a connection to RabbitMQ server
	conn, err := amqp.Dial("amqp://guest:guest@message-broker:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare a queue
	queueName := "user_deletion_notifications"
	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Publish a message
	message := []byte(fmt.Sprintf("User %d has been deleted", userId))
	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}

	return nil
}