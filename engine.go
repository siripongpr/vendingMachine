package vendingMachine

type VendingMachine struct {
	insertedMoney int
	coins map[string]int
	items map[string]int
}

func NewVendingMachine(coins, items map[string]int) *VendingMachine {
	
	return &VendingMachine{coins: coins, items: items}
}

func (m VendingMachine) InsertedMoney() int {
	return m.insertedMoney
}

func (m *VendingMachine) InsertCoin(coin string) {
	m.insertedMoney += m.coins[coin]
}

func (m *VendingMachine) SelectSD() string {
	price := m.items["SD"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "SD" + m.change(change)
}

func (m *VendingMachine) SelectCC() string {
	price := m.items["CC"]
	change := m.insertedMoney - price
	m.insertedMoney = 0
	return "CC" + m.change(change)
}
func (m VendingMachine) GetInsertedMoney() int {
	return ((vm.t * 10) + (vm.f * 5) + (vm.tw * 2) + vm.o)
}

