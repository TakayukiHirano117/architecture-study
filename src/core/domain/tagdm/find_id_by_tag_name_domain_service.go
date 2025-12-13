//go:generate mockgen -source=$GOFILE -destination=../../../support/mock/domain/tagdm/find_id_by_tag_name_domain_service_mock.go -package=tagdm_mock
package tagdm

import (
	"context"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/shared"
)

type FindIDByTagNameDomainService interface {
	Exec(ctx context.Context, tagName TagName) (*shared.UUID, error)
}

type findIDByTagNameDomainService struct {
	tagRepo TagRepository
}

func NewFindIDByTagNameDomainService(tr TagRepository) FindIDByTagNameDomainService {
	return &findIDByTagNameDomainService{
		tagRepo: tr,
	}
}

func (ftds *findIDByTagNameDomainService) Exec(ctx context.Context, tagName TagName) (*shared.UUID, error) {
	tagID, err := ftds.tagRepo.FindIdByTagName(ctx, tagName)

	if err != nil {
		return nil, err
	}

	return tagID, nil
}
