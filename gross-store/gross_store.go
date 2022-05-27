package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	m := make(map[string]int)
	m["quarter_of_a_dozen"] = 3
	m["half_of_a_dozen"] = 6
	m["dozen"] = 12
	m["small_gross"] = 120
	m["gross"] = 144
	m["great_gross"] = 1728

	return m
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	n, prs := units[unit]

	if !prs {
		return prs
	}

	amount, prs := bill[item]

	if prs {
		bill[item] = amount + n
	} else {
		bill[item] = n
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitN, presentUnit := units[unit]

	if !presentUnit {
		return false
	}

	billAmount, presentItem := bill[item]

	if !presentItem {
		return false
	}

	reducedAmount := billAmount - unitN

	if reducedAmount == 0 {
		delete(bill, item)
		return true
	} else if reducedAmount < 0 {
		return false
	} else {
		bill[item] = reducedAmount
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billAmount, presentItem := bill[item]

	if presentItem {
		return billAmount, presentItem
	}

	return 0, false
}
