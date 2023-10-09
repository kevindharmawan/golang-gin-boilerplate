package example

import "boilerplate/internal/models"

type exampleServiceImpl struct {
	exampleRepository ExampleRepository
}

func NewExampleService(
	exampleRepository ExampleRepository,
) ExampleService {
	return &exampleServiceImpl{
		exampleRepository: exampleRepository,
	}
}

func (s *exampleServiceImpl) GetAllExamples() (*[]models.Example, *models.AppError) {
	return s.exampleRepository.GetAllExamples()
}

func (s *exampleServiceImpl) GetExampleById(id int64) (*models.Example, *models.AppError) {
	return s.exampleRepository.GetExampleById(id)
}

func (s *exampleServiceImpl) CreateExample(example *models.Example) (*models.Example, *models.AppError) {
	return s.exampleRepository.CreateExample(example)
}

func (s *exampleServiceImpl) UpdateExample(example *models.Example) (*models.Example, *models.AppError) {
	return s.exampleRepository.UpdateExample(example)
}

func (s *exampleServiceImpl) DeleteExampleById(id int64) *models.AppError {
	return s.exampleRepository.DeleteExampleById(id)
}
