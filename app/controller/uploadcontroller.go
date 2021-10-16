package controller

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/middleware"
	"github.com/directoryxx/fiber-clean-template/app/utils/response"
	"github.com/gofiber/fiber/v2"
)

var pageUploadPermission string = "upload"

// A UserController belong to the interface layer.
type UploadController struct {
	Logger   interfaces.Logger
	Fiber    *fiber.App
	Enforcer *casbin.Enforcer
}

func NewUploadController(logger interfaces.Logger, fiber *fiber.App, enforcer *casbin.Enforcer) *UploadController {
	return &UploadController{
		Logger:   logger,
		Fiber:    fiber,
		Enforcer: enforcer,
	}
}

func (controller UploadController) UploadRouter() {
	controller.Fiber.Post("/upload", middleware.CheckPermission(controller.Enforcer, pageUploadPermission), controller.uploadTemp)
}

// Upload Temp
// @Summary Upload Temp
// @Description Upload Temp
// @Tags Role
// @Param Authorization header string true "With the bearer started"
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /upload [post]
func (controller UploadController) uploadTemp(c *fiber.Ctx) error {
	controller.Logger.LogAccess("%s %s %s\n", c.IP(), c.Method(), c.OriginalURL())

	file, err := c.FormFile("document")
	if err != nil {
		return err
	}
	// Save file to root directory:
	c.SaveFile(file, fmt.Sprintf("/app/public/%s", file.Filename))

	return c.JSON(&response.SuccessResponse{
		Success: true,
		Data:    "dataRole",
		Message: "Role berhasil ditampilkan",
	})

}
