package repository

import (
	"CRUD/models"
)

type TagsRepository interface{
    Save(tags models.Tag)
    Update(tags models.Tag)
    Delete(tagId int)
    FindById(tagId int) (tags models.Tag, error error)
    FindAll() []models.Tag
}