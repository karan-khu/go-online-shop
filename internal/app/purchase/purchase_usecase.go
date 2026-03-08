package purchase

type purchaseUsecaseImpl struct {
	purchaseRepo PurchaseRepository
}

func NewPurchaseUsecase(purchaseRepo PurchaseRepository) PurchaseUsecase {
	return &purchaseUsecaseImpl{
		purchaseRepo: purchaseRepo,
	}
}

func (u *purchaseUsecaseImpl) Listing(userId string) ([]*PurchaseHistoryEntity, error) {
	return u.purchaseRepo.Listing(userId)
}
