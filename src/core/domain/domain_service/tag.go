package domain_service

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/tagdm"
)

type TagDomainService interface {
	IsExistByTagId(ctx context.Context, tagId tagdm.TagId) (bool, error)
	IsExistByTagName(ctx context.Context, tagName string) (bool, error)
}

type tagDomainService struct {
	tagRepo tagdm.TagRepository
}

func NewTagDomainService(tr tagdm.TagRepository) TagDomainService {
	return &tagDomainService{
		tagRepo: tr,
	}
}

func (tds *tagDomainService) IsExistByTagName(ctx context.Context, tagName string) (bool, error) {
	tag, err := tds.tagRepo.FindByTagName(ctx, tagName)

	if err != nil {
		return false, err
	}

	return tag != nil, nil
}

func (tds *tagDomainService) IsExistByTagId(ctx context.Context, tagId tagdm.TagId) (bool, error) {
	tag, err := tds.tagRepo.FindById(ctx, tagId)

	if err != nil {
		return false, err
	}

	return tag != nil, nil
}
