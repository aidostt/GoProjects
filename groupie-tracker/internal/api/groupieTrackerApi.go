package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/internal/models"
	"io"
	"net/http"
	"strconv"
)

const (
	artistsApi   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsApi = "https://groupietrackers.herokuapp.com/api/locations"
	datesApi     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsApi = "https://groupietrackers.herokuapp.com/api/relation"
)

type Api struct {
}

func New() *Api {
	return &Api{}
}

func (a *Api) GetGroups() (*[]models.Group, error) {
	var groups []models.Group

	jsonFile, err := http.Get(artistsApi)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Body.Close()
	jsonData, err := io.ReadAll(jsonFile.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &groups)
	if err != nil {
		return nil, err
	}
	return &groups, nil
}

func (a *Api) GetGroupById(id int) (*models.Group, error) {
	var group models.Group

	artistApi := artistsApi + "/" + strconv.Itoa(id)
	relationApi := relationsApi + "/" + strconv.Itoa(id)

	jsonFile, err := http.Get(artistApi)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Body.Close()
	jsonData, err := io.ReadAll(jsonFile.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonData, &group)
	if err != nil {
		return nil, err
	}

	// Set Relations
	jsonRelationFile, err := http.Get(relationApi)
	if err != nil {
		return nil, err
	}
	defer jsonRelationFile.Body.Close()
	jsonRelationData, err := io.ReadAll(jsonRelationFile.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonRelationData, &group.Relations)
	if err != nil {
		return nil, err
	}

	if group.Id == 0 {
		return nil, fmt.Errorf("Not Found")
	}
	return &group, nil
}
