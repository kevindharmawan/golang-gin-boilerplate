package example

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"boilerplate/internal/models"
	"boilerplate/internal/shared/dto"
)

type ExampleHandler struct {
	exampleService ExampleService
}

func NewExampleHandler(exampleService ExampleService) *ExampleHandler {
	return &ExampleHandler{
		exampleService: exampleService,
	}
}

type exampleResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}

type exampleCreate struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"description" binding:"-"`
}

type exampleUpdate struct {
	Name string `json:"name" binding:"-"`
	Desc string `json:"description" binding:"-"`
}

func (h *ExampleHandler) makeExampleResponseFrom(example *models.Example) *exampleResponse {
	return &exampleResponse{
		ID:   example.ID,
		Name: example.Name,
		Desc: example.Desc,
	}
}

func (h *ExampleHandler) GetAllExample(c *gin.Context) {
	examples, err := h.exampleService.GetAllExamples()

	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	var output []exampleResponse
	for _, example := range *examples {
		output = append(output, *h.makeExampleResponseFrom(&example))
	}

	dto.SendSuccessResponse(c, output)
}

func (h *ExampleHandler) GetExampleById(c *gin.Context) {
	id, convErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if convErr != nil {
		dto.SendErrorResponse(c, models.NewBadRequestError("ID is not in the correct format"))
		return
	}

	output, err := h.exampleService.GetExampleById(id)

	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, h.makeExampleResponseFrom(output))
}

func (h *ExampleHandler) CreateExample(c *gin.Context) {
	var input exampleCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.SendErrorResponse(c, models.NewBadRequestError(err.Error()))
		return
	}

	example := models.Example{
		Name: input.Name,
		Desc: input.Desc,
	}
	output, err := h.exampleService.CreateExample(&example)

	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, h.makeExampleResponseFrom(output))
}

func (h *ExampleHandler) UpdateExample(c *gin.Context) {
	var input exampleUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		dto.SendErrorResponse(c, models.NewBadRequestError(err.Error()))
		return
	}

	id, convErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if convErr != nil {
		dto.SendErrorResponse(c, models.NewBadRequestError("ID is not in the correct format"))
		return
	}

	example := models.Example{
		ID:   id,
		Name: input.Name,
		Desc: input.Desc,
	}
	output, err := h.exampleService.UpdateExample(&example)

	if err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, h.makeExampleResponseFrom(output))
}

func (h *ExampleHandler) DeleteExample(c *gin.Context) {
	id, convErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if convErr != nil {
		dto.SendErrorResponse(c, models.NewBadRequestError("ID is not in the correct format"))
		return
	}

	if err := h.exampleService.DeleteExampleById(id); err != nil {
		dto.SendErrorResponse(c, err)
		return
	}

	dto.SendSuccessResponse(c, "Example deleted.")
}
