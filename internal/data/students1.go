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

const (
	tableName   = "Students"
	partitionID = "students"
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

package data

import (
	"context"
	"fmt"

	"github.com/RomanTykhyi/students-api/internal/di"
	"github.com/RomanTykhyi/students-api/internal/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const TABLE_NAME = "Students"
const PARTITION = "students"

func retrieveDynamoClient() (*dynamodb.Client, error) {
	client, err := di.GetAppContainer().Resolve("dynamo-client")
	if err != nil {
		return nil, fmt.Errorf("cannot get dynamo client: %w", err)
	}
	dbClient, ok := client.(*dynamodb.Client)
	if !ok {
		return nil, fmt.Errorf("could not assert dynamo client type")
	}
	return dbClient, nil
}

func getStudentKeyMap(studentId string) (map[string]types.AttributeValue, error) {
	partitionId, err := attributevalue.Marshal(PARTITION)
	if err != nil {
		return nil, fmt.Errorf("error marshalling partition id: %w", err)
	}

	id, err := attributevalue.Marshal(studentId)
	if err != nil {
		return nil, fmt.Errorf("error marshalling student id: %w", err)
	}

	return map[string]types.AttributeValue{"PartitionId": partitionId, "Id": id}, nil
}

type StudentsStore interface {
	// [unchanged interface definitions]
}

type StudentsRepository struct{}

func (repo StudentsRepository) PutStudent(student *models.Student) error {
	// [unchanged PutStudent method, replace panic/log.Fatal with error return]
	// ...
}

func (repo StudentsRepository) QueryStudents() ([]models.Student, error) {
	// [update function signature to return error]
	// ...
}

func (repo StudentsRepository) GetStudent(id string) (*models.Student, error) {
	// [update function signature to return error]
	// ...
}

func (repo StudentsRepository) UpdateStudent(student *models.Student) (*models.Student, error) {
	// [update function signature to return error]
	// ...
}

func (repo StudentsRepository) DeleteStudent(id string) (bool, error) {
	// [update function signature to return error]
	// ...
}

// [Implement the rest of CRUD operations in a similar fashion]
