package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unit_maps := make(map[string]int)
    unit_maps["quarter_of_a_dozen"] = 3
    unit_maps["half_of_a_dozen"] = 6
    unit_maps["dozen"] = 12
    unit_maps["small_gross"] = 120 
    unit_maps["gross"] = 144
    unit_maps["great_gross"] = 1728
    return unit_maps
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exist := units[unit]
    if !exist {
        return false
    }
    
    bill[item] = bill[item] + value
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unit_value, exists := units[unit]
    if !exists {
        return false
    }
    bill_value, exist := bill[item]
    if !exist {
        return false
    }
	if bill_value - unit_value < 0 {
        return false
    } else if bill_value - unit_value == 0 {
        delete(bill,item)
    } else {
        bill[item] = bill[item] - unit_value
    }
    return true
    
    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	item_value, exist := bill[item]
    if !exist {
        return 0,false
    }
    return item_value, true
}
