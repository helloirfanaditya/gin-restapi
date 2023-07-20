package repositories

import (
	"trawlcode/database"
	"trawlcode/models"
)

type ResultPost struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	User        struct {
		UserID uint   `json:"user_id"`
		Name   string `json:"name"`
	} `json:"user"`
}

func ListPost(payload []models.Post) interface{} {
	var result []ResultPost
	database.Db.Preload("User").Find(&payload)
	for _, data := range payload {
		r := ResultPost{
			ID:          data.ID,
			Title:       data.Title,
			Description: data.Description,
			CreatedAt:   data.CreatedAt.Format("January 2, 2006"),
			UpdatedAt:   data.UpdatedAt.Format("January 2, 2006"),
			User: struct {
				UserID uint   `json:"user_id"`
				Name   string `json:"name"`
			}{
				UserID: data.User.ID,
				Name:   data.User.Name,
			},
		}
		result = append(result, r)
	}

	if len(payload) == 0 {
		return []string{}
	}

	return result
}
