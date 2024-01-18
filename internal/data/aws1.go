// Logging: Printfm vs log 

package data

import (
	"context"
	"errors"
	"log"

	"github.com/RomanTykhyi/students-api/internal/di"
	"github.com/RomanTykhyi/students-api/internal/models"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type StudentsStore interface {
	PutStudent(student *models.Student) error
	QueryStudents() ([]models.Student, error)
	GetStudent(id string) (*models.Student, error)
	UpdateStudent(student *models.Student) error
	DeleteStudent(id string) error
}

type StudentsRepository struct {
	tableName   string
	partitionID string
	dbClient    *dynamodb.Client
}

func NewStudentsRepository(tableName, partitionID string) (*StudentsRepository, error) {
	client, err := di.GetAppContainer().Resolve("dynamo-client")
	if err != nil {
		return nil, err
	}
	return &StudentsRepository{
		tableName:   tableName,
		partitionID: partitionID,
		dbClient:    client.(*dynamodb.Client),
	}, nil
}

func (repo *StudentsRepository) PutStudent(student *models.Student) error {
	// Perform error checks for attributevalue.MarshalMap and dbClient.PutItem

	// Implementation
	return nil
}

func (repo *StudentsRepository) QueryStudents() ([]models.Student, error) {
	// Perform error checks for dbClient.Query and attributevalue.UnmarshalListOfMaps

	// Implementation
	return nil, nil
}

func (repo *StudentsRepository) GetStudent(id string) (*models.Student, error) {
	// Perform error checks for dbClient.GetItem, attributevalue.UnmarshalMap, and nil check

	// Implementation
	return nil, nil
}

func (repo *StudentsRepository) UpdateStudent(student *models.Student) error {
	// Perform error checks for expression.NewBuilder, dbClient.UpdateItem, and attributevalue.UnmarshalMap

	// Implementation
	return nil
}

func (repo *StudentsRepository) DeleteStudent(id string) error {
	// Perform error checks for dbClient.DeleteItem

	// Implementation
	return nil
}
