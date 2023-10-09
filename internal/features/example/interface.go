package example

import "boilerplate/internal/models"

type ExampleRepository interface {
	GetAllExamples() (*[]models.Example, *models.AppError)
	GetExampleById(id int64) (*models.Example, *models.AppError)
	CreateExample(example *models.Example) (*models.Example, *models.AppError)
	UpdateExample(example *models.Example) (*models.Example, *models.AppError)
	DeleteExampleById(id int64) *models.AppError
}

type ExampleService interface {
	GetAllExamples() (*[]models.Example, *models.AppError)
	GetExampleById(id int64) (*models.Example, *models.AppError)
	CreateExample(example *models.Example) (*models.Example, *models.AppError)
	UpdateExample(example *models.Example) (*models.Example, *models.AppError)
	DeleteExampleById(id int64) *models.AppError
}
