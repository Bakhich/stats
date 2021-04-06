package stats

import (
	"fmt"

	"github.com/Bakhich/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   53_00,
			Category: "Cat",
			Status:   "OK",
		},
		{
			ID:       2,
			Amount:   51_00,
			Category: "Cat",
			Status:   "OK",
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cat",
			Status:   "FAIL",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 3466
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10_000_00,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   20_000_00,
			Category: "pharmacy",
		},
		{
			ID:       3,
			Amount:   30_000_00,
			Category: "restaurant",
		},
	}

	inCategory := types.Category("auto")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)
	//Output:  1000000

}
