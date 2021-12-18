package entity

type Manipulate struct {
	Start string `json:"start" binding:"required"`
	Move  int32  `json:"move"`
}
