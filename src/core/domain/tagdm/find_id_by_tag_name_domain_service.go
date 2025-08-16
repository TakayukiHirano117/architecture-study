package tagdm

import "context"

type FindIDByTagNameDomainService interface {
	Exec(ctx context.Context, tagName TagName) (*TagID, error)
}

type findIDByTagNameDomainService struct {
	tagRepo TagRepository
}

func NewFindIDByTagNameDomainService(tr TagRepository) FindIDByTagNameDomainService {
	return &findIDByTagNameDomainService{
		tagRepo: tr,
	}
}

func (ftds *findIDByTagNameDomainService) Exec(ctx context.Context, tagName TagName) (*TagID, error) {
	tagID, err := ftds.tagRepo.FindIdByTagName(ctx, tagName)

	if err != nil {
		return nil, err
	}

	return tagID, nil
}
