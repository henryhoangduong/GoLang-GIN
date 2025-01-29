package repository

import (
	"CRUD/helper"
	"CRUD/models"

	"github.com/jinzhu/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImpl) FindAll() []models.Tag {
	panic("unimplemented")
}

// FindById implements TagsRepository.
func (t *TagsRepositoryImpl) FindById(tagId int) (tags models.Tag, error error) {
	panic("unimplemented")
}

// Save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags models.Tag) {
	panic("unimplemented")
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tags models.Tag) {
	panic("unimplemented")
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags models.Tag
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.error)
	panic("unimplmented")
}
