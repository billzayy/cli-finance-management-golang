package pkg

import (
	"flag"
	"fmt"
	"os"
)

const fileName = "./budget.json"

func InitFlag(list *List) {
	add := flag.String("add", "", "Add Budget")
	price := flag.Float64("price", 0.0, "Price Budget")

	addBudget := flag.Float64("addBudget", 0.0, "Add Budget")

	delete := flag.Int("delete", 0, "Delete Budget")

	updateName := flag.String("updateName", "", "Update Name")
	updatePrice := flag.Float64("updatePrice", 0.0, "Update Price")
	index := flag.Int("index", 0, "Index of Financial")

	listArr := flag.Bool("list", false, "list array")

	flag.Parse()

	if err := LoadFile(fileName, &list.Array); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add != "" && *price != 0:
		resultArr := list.Add(*add, *price, 0.0)

		err := CreateFile(fileName, resultArr)

		if err != nil {
			panic(err)
		}

	case *addBudget != 0:
		resultArr := list.Add("Add Budget", 0.0, *addBudget)

		err := CreateFile(fileName, resultArr)

		if err != nil {
			panic(err)
		}

	case *delete != 0:
		resultArr := list.Delete(*delete - 1)

		err := CreateFile(fileName, resultArr)

		if err != nil {
			panic(err)
		}

	case *updateName != "" && *index != 0:
		resultArr := list.UpdateName(*index-1, *updateName)

		err := CreateFile(fileName, resultArr)

		if err != nil {
			panic(err)
		}

	case *updatePrice != 0.0 && *index != 0:
		resultArr := list.UpdatePrice(*index-1, *updatePrice)

		err := CreateFile(fileName, resultArr)

		if err != nil {
			panic(err)
		}

	case *listArr:
		Load(list.Array)
		return

	default:
		fmt.Fprint(os.Stdout, "invalid command")
		os.Exit(1)
	}
}
