package controller

import (
	"github.com/gofiber/fiber/v2"
	"programmerzamannow/belajar-golang-restful-api/model/web"
	"programmerzamannow/belajar-golang-restful-api/service"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(c *fiber.Ctx) error {
	categoryCreateRequest := new(web.CategoryCreateRequest)
	if err := c.BodyParser(categoryCreateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	categoryResponse := controller.CategoryService.Create(c.Context(), *categoryCreateRequest)
	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Update(c *fiber.Ctx) error {
	categoryUpdateRequest := new(web.CategoryUpdateRequest)
	if err := c.BodyParser(categoryUpdateRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("categoryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Category ID",
			Data:   err.Error(),
		})
	}
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(c.Context(), *categoryUpdateRequest)
	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	})
}

func (controller *CategoryControllerImpl) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("categoryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Category ID",
			Data:   err.Error(),
		})
	}

	controller.CategoryService.Delete(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
	})
}

func (controller *CategoryControllerImpl) FindById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("categoryId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Category ID",
			Data:   err.Error(),
		})
	}

	categoryResponse := controller.CategoryService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	})
}

func (controller *CategoryControllerImpl) FindAll(c *fiber.Ctx) error {
	categoryResponses := controller.CategoryService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	})
}
