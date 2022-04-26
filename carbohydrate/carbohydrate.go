package carbohydrate

import (
	"encoding/json"
	"fmt"
	"math"
)

type Carbo struct {
	Name     string
	Calories int
	Price    int
}

type Carbohydrates []*Carbo

func (carbos *Carbohydrates) ChangePriceDay(ratio, day int) {
	for _, carbo := range *carbos {
		carbo.Price += carbo.Price * ratio / 100

		fmt.Printf("%s Day %d : Rp%d\n", carbo.Name, day, carbo.Price)
	}
	fmt.Println("")
}

type Price struct {
	Name string
	TotalPrice int
	TotalCalories int
}
type DetailPrice struct {
	Day string
	Lists []*Price
}

func Carbohydrate() {

	carbos := &Carbohydrates{
		&Carbo{
			Name:     "Rice",
			Calories: 28,
			Price:    1000, // 100gr
		},
		&Carbo{
			Name:     "Corn",
			Calories: 21,
			Price:    600, // 100gr
		},
		&Carbo{
			Name:     "Potato",
			Calories: 17,
			Price:    400, // 100gr
		},
	}

	ratioChangePerDay := 5 // percent
	gramsPerDay := 400
	gramsPerEach := int(math.Floor(float64(gramsPerDay / 3))) + 1

	detailPrices := []*DetailPrice{}
	
	for i := 1; i <= 3; i++ {
		prices := []*Price{}

		carbos.ChangePriceDay(ratioChangePerDay, i)

		priceThisDay := 0

		for _, c := range *carbos {
			
			gramsThisDay := 0
			priceThisCarbo := 0

			for {
				if gramsThisDay <= gramsPerEach {
					gramsThisDay += c.Calories

					priceThisCarbo += c.Price
					priceThisDay += c.Price
				} else {
					break
				}
			}

			prices = append(prices, &Price{
				Name: c.Name,
				TotalPrice: priceThisCarbo,
				TotalCalories: gramsThisDay,
			})
		}

		detailPrices = append(detailPrices, &DetailPrice{
			Day: fmt.Sprintf("Day %d", i),
			Lists: prices,
		})
	}

	output, _ := json.Marshal(detailPrices)
	outputMap := []map[string]interface{}{}

	json.Unmarshal(output, &outputMap)
	fmt.Println(outputMap)
}