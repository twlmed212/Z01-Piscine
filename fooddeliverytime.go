package piscine

type food struct {
	burgerTime  int
	chipsTime   int
	nuggetsTime int
}

func FoodDeliveryTime(order string) int {
	setTime := food{15, 10, 12}
	if order == "burger" {
		return setTime.burgerTime
	} else if order == "chips" {
		return setTime.chipsTime
	} else if order == "nuggets" {
		return setTime.nuggetsTime
	} else {
		return 404
	}
}
