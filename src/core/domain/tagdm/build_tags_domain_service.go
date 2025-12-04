//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/tagdm/build_tags_domain_service_mock.go -package=tagdm_mock
package tagdm

import (
	"context"
	"errors"
)

type TagRequest struct {
	ID   string
	Name string
}

type BuildTagsDomainService interface {
	Exec(ctx context.Context, reqTags []TagRequest) ([]Tag, error)
}

type buildTagsDomainService struct {
	tagRepo TagRepository
}

func NewBuildTagsDomainService(tr TagRepository) BuildTagsDomainService {
	return &buildTagsDomainService{
		tagRepo: tr,
	}
}

func (btds *buildTagsDomainService) Exec(ctx context.Context, reqTags []TagRequest) ([]Tag, error) {
	var existingTagIDs []TagID
	newTags := make([]Tag, len(reqTags))

	for i, reqTag := range reqTags {
		if reqTag.ID == "" {
			tagName, err := NewTagNameByVal(reqTag.Name)
			if err != nil {
				return nil, err
			}
			newTag, err := NewTag(NewTagID(), tagName)
			if err != nil {
				return nil, err
			}
			newTags[i] = *newTag
		} else {
			tagID, err := NewTagIDByVal(reqTag.ID)
			if err != nil {
				return nil, err
			}
			existingTagIDs = append(existingTagIDs, tagID)
		}
	}

	if len(newTags) > 0 {
		if err := btds.tagRepo.BulkInsert(ctx, newTags); err != nil {
			return nil, err
		}
	}

	existingTags := make([]Tag, 0)

	if len(existingTagIDs) > 0 {
		var err error
		existingTags, err = btds.tagRepo.FindByIDs(ctx, existingTagIDs)
		if err != nil {
			return nil, err
		}

		if len(existingTags) != len(existingTagIDs) {
			return nil, errors.New("some tags not found")
		}
	}

	tags := append(existingTags, newTags...)
	return tags, nil
}
