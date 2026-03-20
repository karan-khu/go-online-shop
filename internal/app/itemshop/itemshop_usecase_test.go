package itemshop

import (
	"errors"
	"log/slog"
	"os"
	"testing"

	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/karan-khu/go-online-shop/internal/app/balance"
	"github.com/karan-khu/go-online-shop/internal/app/inventory"
	"github.com/karan-khu/go-online-shop/internal/app/item"
	"github.com/karan-khu/go-online-shop/internal/app/purchase"
)

// ========================================================================
// Mock: ใช้ testify/mock — แค่ embed mock.Mock แล้ว implement interface
// ========================================================================

type MockItemShopRepo struct{ mock.Mock }

func (m *MockItemShopRepo) Listing(limit int, page int, searchText string) ([]*ItemShopEntity, int64, error) {
	args := m.Called(limit, page, searchText)
	return args.Get(0).([]*ItemShopEntity), args.Get(1).(int64), args.Error(2)
}
func (m *MockItemShopRepo) Begin() *gorm.DB            { return m.Called().Get(0).(*gorm.DB) }
func (m *MockItemShopRepo) Commit(tx *gorm.DB) error   { return m.Called(tx).Error(0) }
func (m *MockItemShopRepo) Rollback(tx *gorm.DB) error { return m.Called(tx).Error(0) }

type MockImageBuilder struct{ mock.Mock }

func (m *MockImageBuilder) GetHost() string                              { return "" }
func (m *MockImageBuilder) Build(path string) string                     { return m.Called(path).String(0) }
func (m *MockImageBuilder) SaveImage(pctx *echo.Context) (string, error) { return "", nil }

// mock ที่ยังไม่ได้ใช้ แต่ต้องมีเพราะ NewItemShopUsecase ต้องการ
type MockItemRepo struct{ mock.Mock }

func (m *MockItemRepo) FindById(itemId int) (*item.ItemEntity, error) {
	args := m.Called(itemId)
	return args.Get(0).(*item.ItemEntity), args.Error(1)
}
func (m *MockItemRepo) FindExists(itemId int) bool                          { return false }
func (m *MockItemRepo) Create(i *item.ItemEntity) (*item.ItemEntity, error) { return nil, nil }
func (m *MockItemRepo) Edit(i *item.ItemEntity) (*item.ItemEntity, error)   { return nil, nil }
func (m *MockItemRepo) Archive(itemId int) error                            { return nil }

type MockBalanceRepo struct{ mock.Mock }

func (m *MockBalanceRepo) CoinAdd(tx *gorm.DB, req *balance.UserBalanceEntity) (*balance.UserBalanceEntity, error) {
	args := m.Called(tx, req)
	return args.Get(0).(*balance.UserBalanceEntity), args.Error(1)
}
func (m *MockBalanceRepo) CoinShow(userId string) (*balance.UserCoin, error) {
	args := m.Called(userId)
	return args.Get(0).(*balance.UserCoin), args.Error(1)
}

type MockInventoryRepo struct{ mock.Mock }

func (m *MockInventoryRepo) Listing(userId string) ([]*inventory.InventoryEntity, error) {
	return nil, nil
}
func (m *MockInventoryRepo) Filling(tx *gorm.DB, userId string, itemId int, qty int) ([]*inventory.InventoryEntity, error) {
	args := m.Called(tx, userId, itemId, qty)
	return args.Get(0).([]*inventory.InventoryEntity), args.Error(1)
}
func (m *MockInventoryRepo) Removing(tx *gorm.DB, userId string, itemId int, limit int) error {
	return m.Called(tx, userId, itemId, limit).Error(0)
}
func (m *MockInventoryRepo) UserItemCount(userId string, itemId int) int {
	return m.Called(userId, itemId).Int(0)
}

type MockPurchaseRepo struct{ mock.Mock }

func (m *MockPurchaseRepo) Listing(userId string) ([]*purchase.PurchaseHistoryEntity, error) {
	return nil, nil
}
func (m *MockPurchaseRepo) Create(tx *gorm.DB, p *purchase.PurchaseHistoryEntity) (*purchase.PurchaseHistoryEntity, error) {
	args := m.Called(tx, p)
	return args.Get(0).(*purchase.PurchaseHistoryEntity), args.Error(1)
}

// ========================================================================
// Helper
// ========================================================================

func testLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
}

// ========================================================================
// Test: ItemList
// ========================================================================

func TestItemList_Success(t *testing.T) {
	// Arrange
	shopRepo := new(MockItemShopRepo)
	imgBuilder := new(MockImageBuilder)

	shopRepo.On("Listing", 10, 1, "").Return([]*ItemShopEntity{
		{ItemId: 1, Name: "Sword", Description: "A sharp sword", Picture: "uploads/sword.png", Price: 100},
		{ItemId: 2, Name: "Shield", Description: "A strong shield", Picture: "uploads/shield.png", Price: 200},
	}, int64(2), nil)

	imgBuilder.On("Build", "uploads/sword.png").Return("http://localhost/uploads/sword.png")
	imgBuilder.On("Build", "uploads/shield.png").Return("http://localhost/uploads/shield.png")

	usecase := NewItemShopUsecase(testLogger(), shopRepo, nil, nil, nil, nil, imgBuilder)

	// Act
	result, err := usecase.ItemList(&RequestItemFilter{Limit: 10, Page: 1, SearchText: ""})

	// Assert
	assert.NoError(t, err)
	assert.Len(t, result.Items, 2)
	assert.Equal(t, "Sword", result.Items[0].Name)
	assert.Equal(t, "http://localhost/uploads/sword.png", result.Items[0].Picture)
	assert.Equal(t, 2, result.Pagination.Total)
	assert.Equal(t, 1, result.Pagination.TotalPages)

	shopRepo.AssertExpectations(t)
	imgBuilder.AssertExpectations(t)
}

func TestItemList_Empty(t *testing.T) {
	shopRepo := new(MockItemShopRepo)
	shopRepo.On("Listing", 10, 1, "").Return([]*ItemShopEntity{}, int64(0), nil)

	usecase := NewItemShopUsecase(testLogger(), shopRepo, nil, nil, nil, nil, new(MockImageBuilder))

	result, err := usecase.ItemList(&RequestItemFilter{Limit: 10, Page: 1})

	assert.NoError(t, err)
	assert.Empty(t, result.Items)
	assert.Equal(t, 0, result.Pagination.Total)
}

func TestItemList_RepoError(t *testing.T) {
	shopRepo := new(MockItemShopRepo)
	shopRepo.On("Listing", 10, 1, "").Return(([]*ItemShopEntity)(nil), int64(0), errors.New("db connection failed"))

	usecase := NewItemShopUsecase(testLogger(), shopRepo, nil, nil, nil, nil, nil)

	result, err := usecase.ItemList(&RequestItemFilter{Limit: 10, Page: 1})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "db connection failed")
}
