package svc

import (
	"work/gingo/entity"
)

type Manipulate interface {
	Save(mn entity.Manipulate) entity.Manipulate
}

type ManipulateService struct {
	manipulate entity.Manipulate
}

func NewManuPulate() Manipulate {
	return &ManipulateService{
		manipulate: entity.Manipulate{},
	}
}

func (sv *ManipulateService) Save(mn entity.Manipulate) entity.Manipulate {
	sv.manipulate = mn
	return sv.manipulate
}
