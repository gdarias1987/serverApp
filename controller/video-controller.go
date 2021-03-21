package controller

import (
	"github.com/gdarias1987/serverApp/customValidators"
	"github.com/gdarias1987/serverApp/entity"
	"github.com/gdarias1987/serverApp/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	validate := validator.New()
	validate.RegisterValidation("is-willi3", customValidators.ValidateWilli3)

	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video) //ShouldBindJSON for validation

	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	c.service.Save(video)
	return nil
}
