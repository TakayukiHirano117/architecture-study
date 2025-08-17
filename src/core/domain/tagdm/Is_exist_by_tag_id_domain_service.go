package tagdm

import "context"

type IsExistByTagIDDomainService interface {
	Exec(ctx context.Context, tagID TagID) (bool, error)
}

type isExistByTagIDDomainService struct {
	tagRepo TagRepository
}

func NewIsExistByTagIDDomainService(tr TagRepository) IsExistByTagIDDomainService {
	return &isExistByTagIDDomainService{
		tagRepo: tr,
	}
}

func (ibt *isExistByTagIDDomainService) Exec(ctx context.Context, tagID TagID) (bool, error) {
	tag, err := ibt.tagRepo.FindByID(ctx, tagID)

	if err != nil {
		return false, err
	}

	return tag != nil, nil
}