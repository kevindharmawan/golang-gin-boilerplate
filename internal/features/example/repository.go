package example

import (
	"boilerplate/internal/models"
	"errors"

	"gorm.io/gorm"
)

type exampleRepositoryImpl struct {
	db *gorm.DB
}

func NewExampleRepository(
	db *gorm.DB,
) ExampleRepository {
	return &exampleRepositoryImpl{
		db: db,
	}
}

type ExampleGorm struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"not null"`
	Desc string
}

func (ExampleGorm) TableName() string {
	return "examples"
}

func (r *exampleRepositoryImpl) makeExampleGormFrom(example *models.Example) *ExampleGorm {
	return &ExampleGorm{
		ID:   example.ID,
		Name: example.Name,
		Desc: example.Desc,
	}
}

func (r *exampleRepositoryImpl) makeExampleFrom(exampleGorm *ExampleGorm) *models.Example {
	return &models.Example{
		ID:   exampleGorm.ID,
		Name: exampleGorm.Name,
		Desc: exampleGorm.Desc,
	}
}

func (r *exampleRepositoryImpl) GetAllExamples() (*[]models.Example, *models.AppError) {
	var examplesGorm []ExampleGorm
	if result := r.db.Find(&examplesGorm); result.Error != nil {
		return nil, models.NewInternalServerError()
	}

	var examples []models.Example
	for _, exampleGorm := range examplesGorm {
		examples = append(examples, *r.makeExampleFrom(&exampleGorm))
	}
	return &examples, nil
}

func (r *exampleRepositoryImpl) GetExampleById(id int64) (*models.Example, *models.AppError) {
	var exampleGorm ExampleGorm
	if err := r.db.First(&exampleGorm, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.NewNotFoundError("Example")
		}

		return nil, models.NewInternalServerError()
	}

	return r.makeExampleFrom(&exampleGorm), nil
}

func (r *exampleRepositoryImpl) CreateExample(example *models.Example) (*models.Example, *models.AppError) {
	exampleGorm := r.makeExampleGormFrom(example)

	if err := r.db.Create(&exampleGorm).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, models.NewConflictError("Example")
		}

		return nil, models.NewInternalServerError()
	}

	return r.makeExampleFrom(exampleGorm), nil
}

func (r *exampleRepositoryImpl) UpdateExample(example *models.Example) (*models.Example, *models.AppError) {
	exampleGorm := r.makeExampleGormFrom(example)

	if err := r.db.Model(&exampleGorm).Updates(exampleGorm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.NewNotFoundError("Example")
		}

		return nil, models.NewInternalServerError()
	}

	return r.makeExampleFrom(exampleGorm), nil
}

func (r *exampleRepositoryImpl) DeleteExampleById(id int64) *models.AppError {
	if err := r.db.Delete(&ExampleGorm{}, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.NewNotFoundError("Example")
		}

		return models.NewInternalServerError()
	}

	return nil
}
