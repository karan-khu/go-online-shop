package adapter

type BuySellAdapter interface{}

type buySellAdapterImpl struct {
}

func NewBuySellAdapter() BuySellAdapter {
	return &buySellAdapterImpl{}
}
