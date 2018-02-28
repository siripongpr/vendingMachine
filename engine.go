package vendingMachine

type VendingMachine struct {
	insertedMoney int
	coins         map[string]int
	items         map[string]int
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

func (m VendingMachine) change(c int) string {
	var str string
	values := [...]int{10, 5, 2, 1}
	coins := [...]string{"T", "F", "TW", "O"}
	for i := 0; i < len(values); i++ {
		if c >= values[i] {
			str += ", " + coins[i]
			c -= values[i]
			i--
		}
	}
	return str
}

func (m *VendingMachine) CoinReturn() string {
	coins := m.change(m.insertedMoney)
	m.insertedMoney = 0
	return coins[2:len(coins)]
}

