package inventory

type inventoryUsecaseImpl struct {
	inventoryRepo InventoryRepository
}

func NewInventoryUsecase(inventoryRepo InventoryRepository) InventoryUsecase {
	return &inventoryUsecaseImpl{
		inventoryRepo: inventoryRepo,
	}
}

func (u *inventoryUsecaseImpl) Listing(userId string) ([]*InventoryListing, error) {
	inventoryList, err := u.inventoryRepo.Listing(userId)
	if err != nil {
		return nil, err
	}

	mapInventory := make(map[int]*InventoryListing)
	for _, inventory := range inventoryList {
		if _, exists := mapInventory[inventory.ItemId]; exists {
			mapInventory[inventory.ItemId].Quantity++
		} else {
			mapInventory[inventory.ItemId] = &InventoryListing{
				ItemId:          inventory.ItemId,
				ItemName:        inventory.ItemName,
				ItemDescription: inventory.ItemDescription,
				ItemPicture:     inventory.ItemPicture,
				CreatedAt:       inventory.CreatedAt,
				Quantity:        1,
			}
		}
	}

	list := make([]*InventoryListing, len(mapInventory))
	idx := 0
	for _, inventory := range mapInventory {
		list[idx] = inventory
		idx++
	}

	return list, nil
}
