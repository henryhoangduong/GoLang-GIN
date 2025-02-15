package repository

import (
	"errors"
	"henry/data/request"
	"henry/helper"
	"henry/model"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImpl) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return result
}

// FindById implements TagsRepository.
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	var tag model.Tags
	result := t.Db.Find(&tags, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TagsRepository.
func (t *TagsRepositoryImpl) Save(tags model.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// Update implements TagsRepository.
func (t *TagsRepositoryImpl) Update(tags model.Tags) {
	var updateTag = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}

func NewTagsREpositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}
