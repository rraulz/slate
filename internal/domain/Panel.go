package domain

type Panel struct {
	Id   int
	User int
}

func NewPanel(id, user int) *Panel {

	return &Panel{id, user}
}
