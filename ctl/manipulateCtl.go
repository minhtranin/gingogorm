package ctl

import (
	"work/gingo/entity"
	"work/gingo/svc"

	"github.com/gin-gonic/gin"
)

type Manipulate interface {
	Save(ctx *gin.Context) error
}

type ManiputlateCtl struct {
	service svc.Manipulate
}

func NewManuCtl(s svc.Manipulate) Manipulate {
	return &ManiputlateCtl{
		service: s,
	}
}

func (c *ManiputlateCtl) Save(ctx *gin.Context) error {
	var manipl entity.Manipulate
	err := ctx.ShouldBindJSON(&manipl)
	if err != nil {
		return err
	}
	c.service.Save(manipl)
	return nil
}
