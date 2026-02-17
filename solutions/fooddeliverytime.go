package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var item food
	switch order {
	case "burger":
		item.preptime = 15
	case "chips":
		item.preptime = 10
	case "nuggets":
		item.preptime = 12
	default:
		item.preptime = 404
	}

	return item.preptime
}
