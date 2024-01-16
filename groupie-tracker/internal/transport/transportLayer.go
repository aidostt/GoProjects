package transport

import "groupie-tracker/internal/models"

type Repo interface {
	GetGroups() (*[]models.Group, error)
	GetGroupById(int) (*models.Group, error)
}

type Transport struct {
	r Repo
}

func New(r Repo) *Transport {
	return &Transport{r}
}
