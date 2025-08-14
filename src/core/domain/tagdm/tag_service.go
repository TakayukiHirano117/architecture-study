package tagdm

import (
	"context"
)

type TagDomainService interface {
	IsExistByTagId(ctx context.Context, tagId TagId) (bool, error)
	FindIdByTagName(ctx context.Context, tagName TagName) (*TagId, error)
}

type tagDomainService struct {
	tagRepo TagRepository
}

func NewTagDomainService(tr TagRepository) TagDomainService {
	return &tagDomainService{
		tagRepo: tr,
	}
}

func (tds *tagDomainService) FindIdByTagName(ctx context.Context, tagName TagName) (*TagId, error) {
	tagID, err := tds.tagRepo.FindIdByTagName(ctx, tagName)

	if err != nil {
		return nil, err
	}

	return tagID, nil
}

func (tds *tagDomainService) IsExistByTagId(ctx context.Context, tagId TagId) (bool, error) {
	tag, err := tds.tagRepo.FindById(ctx, tagId)

	if err != nil {
		return false, err
	}

	return tag != nil, nil
}
