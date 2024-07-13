package pkg

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
)

func Load(array []Financial) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "NAME"},
			{Align: simpletable.AlignCenter, Text: "Price"},
			{Align: simpletable.AlignCenter, Text: "Budget"},
			{Align: simpletable.AlignCenter, Text: "Created Time"},
		},
	}

	var subTotal = float64(array[0].Budget)
	var priceTotal = float64(0)

	for index, row := range array {
		budgetStr := ""

		if row.Price == 0.0 && index == 0 {
			budgetStr = Green(fmt.Sprintf("$%v↑", row.Budget))
		} else if row.Price == 0.0 && index != 0 {
			budgetStr = Green(fmt.Sprintf("$%v↑", row.Budget))
			subTotal += row.Budget
		} else {
			budgetStr = Red(fmt.Sprintf("$%v↓", row.Price))
			subTotal -= row.Price
		}

		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", index+1)},
			{Text: row.Name},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("$%v", row.Price)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%v", budgetStr)},
			{Text: row.CreateTime.Format(time.RFC822)},
		}

		priceTotal += row.Price
		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{},
			{},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("Total: $%.2f", priceTotal)},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("SubTotal: $%.2f", subTotal)},
			{},
		},
	}
	fmt.Println(table.String())
}
