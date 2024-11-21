package controller

import (
	"database/sql"
	"search_api/model"
)

var(
	Limit = 100
	Offset = 0
) 

type Dependency struct {
	DB *sql.DB
}

type Resp struct {
	Total       int           `json:"total"`
	Results     []model.Users `json:"results"`
	TimeTakenMS int64         `json:"time_taken_ms"`
}
