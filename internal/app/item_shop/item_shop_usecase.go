package itemshop

import (
	"errors"
	"log/slog"
	"math"

	"github.com/karan-khu/go-online-shop/internal/app/balance"
	"github.com/karan-khu/go-online-shop/internal/app/inventory"
	"github.com/karan-khu/go-online-shop/internal/app/item"
	"github.com/karan-khu/go-online-shop/internal/app/purchase"
	"github.com/karan-khu/go-online-shop/pkg/upload"
)

type itemShopUsecaseImpl struct {
	logger        *slog.Logger
	itemShopRepo  ItemShopRepository
	itemRepo      item.ItemRepository
	balanceRepo   balance.BalanceRepository
	inventoryRepo inventory.InventoryRepository
	purchaseRepo  purchase.PurchaseRepository
	imageBuilder  upload.ImageBuilder
}

func NewItemShopUsecase(logger *slog.Logger, itemShopRepo ItemShopRepository, itemRepo item.ItemRepository, balanceRepo balance.BalanceRepository, inventoryRepo inventory.InventoryRepository, purchaseRepo purchase.PurchaseRepository, imageBuilder upload.ImageBuilder) ItemShopUsecase {
	return &itemShopUsecaseImpl{
		logger:        logger,
		itemShopRepo:  itemShopRepo,
		itemRepo:      itemRepo,
		balanceRepo:   balanceRepo,
		inventoryRepo: inventoryRepo,
		purchaseRepo:  purchaseRepo,
		imageBuilder:  imageBuilder,
	}
}

func (u *itemShopUsecaseImpl) ItemList(req *RequestItemFilter) (*ResponseItemList, error) {
	list, total, err := u.itemShopRepo.Listing(req.Limit, req.Page, req.SearchText)
	if err != nil {
		u.logger.Error("failed to get item list", "error", err)
		return nil, err
	}

	itemModels := make([]*ItemShopEntity, len(list))
	for i, item := range list {
		itemModels[i] = &ItemShopEntity{
			ItemId:      item.ItemId,
			Name:        item.Name,
			Description: item.Description,
			Picture:     u.imageBuilder.Build(item.Picture),
			Price:       item.Price,
		}
	}

	return &ResponseItemList{
		Items: itemModels,
		Pagination: &Pagination{
			Total:      int(total),
			TotalPages: int(math.Ceil(float64(total) / float64(req.Limit))),
		},
	}, nil
}

// 1. Check item exists
// 2. Total price item.Price * req.Quantity
// 3. Check coin balance enough
// 4. Adding balance record Buy
// 5. Adding inventory record
// 6. Adding purchase record
// 7. Returning item entity
func (u *itemShopUsecaseImpl) Buying(req *RequestItemBuying) error {
	targetItem, itemErr := u.itemRepo.FindById(req.ItemId)
	if itemErr != nil {
		u.logger.Error("failed to find item by id", "error", itemErr)
		return itemErr
	}

	totalPrice := targetItem.Price * req.Quantity

	userBalance, err := u.balanceRepo.CoinShow(req.UserId)
	if err != nil {
		u.logger.Error("failed to show user balance", "error", err)
		return err
	} else if userBalance.Coin < totalPrice {
		u.logger.Error("coin balance not enough", "error", err)
		return errors.New("Coin balance not enough")
	}

	tx := u.itemShopRepo.Begin()

	newBalance := &balance.UserBalanceEntity{
		UserId: req.UserId,
		Amount: -totalPrice,
		Status: "ITEM-BUY",
	}
	if _, err := u.balanceRepo.CoinAdd(tx, newBalance); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	if _, err := u.inventoryRepo.Filling(tx, req.UserId, targetItem.ItemId, req.Quantity); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	newPurchase := &purchase.PurchaseHistoryEntity{
		BuyerId:         req.UserId,
		ItemId:          targetItem.ItemId,
		ItemName:        targetItem.Name,
		ItemDescription: targetItem.Description,
		ItemPrice:       targetItem.Price,
		Quantity:        req.Quantity,
		Type:            "BUY",
	}
	if _, err := u.purchaseRepo.Create(tx, newPurchase); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	if err := u.itemShopRepo.Commit(tx); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	return nil
}

// 1. Item is exists
// 2. Check item inventory item count
// 3. Check item quantity enough
// 4. Total price selling / 2
// 5. Adding balance record Sell
// 6. Remove inventory record
// 7. Adding purchase record
// 8. Returning item entity
func (u *itemShopUsecaseImpl) Selling(req *RequestItemSelling) error {
	targetItem, err := u.itemRepo.FindById(req.ItemId)
	if err != nil {
		u.logger.Error("failed to find item by id", "error", err)
		return err
	}

	userItemCount := u.inventoryRepo.UserItemCount(req.UserId, req.ItemId)
	if req.Quantity > userItemCount {
		u.logger.Error("item quantity not enough", "error", errors.New("Item quantity not enough"))
		return errors.New("Item quantity not enough")
	}

	totalPrice := (targetItem.Price * req.Quantity) / 2

	tx := u.itemShopRepo.Begin()

	newBalance := &balance.UserBalanceEntity{
		UserId: req.UserId,
		Amount: totalPrice,
		Status: "ITEM-SELL",
	}
	if _, err := u.balanceRepo.CoinAdd(tx, newBalance); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	if err := u.inventoryRepo.Removing(tx, req.UserId, targetItem.ItemId, req.Quantity); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	newPurchase := &purchase.PurchaseHistoryEntity{
		BuyerId:         req.UserId,
		ItemId:          targetItem.ItemId,
		ItemName:        targetItem.Name,
		ItemDescription: targetItem.Description,
		ItemPrice:       targetItem.Price,
		Quantity:        req.Quantity,
		Type:            "SELL",
	}
	if _, err := u.purchaseRepo.Create(tx, newPurchase); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	if err := u.itemShopRepo.Commit(tx); err != nil {
		u.itemShopRepo.Rollback(tx)
		return err
	}

	return nil
}
