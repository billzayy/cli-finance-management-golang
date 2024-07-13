package pkg

import (
	"errors"
	"time"
)

type Financial struct {
	Name       string
	Price      float64
	Budget     float64
	CreateTime time.Time
}

type List struct {
	Array []Financial
}

func (l *List) Add(name string, price float64, budget float64) []Financial {
	if price == 0.0 && budget == 0.0 {
		panic(errors.New("empty input"))
	} else {
		l.Array = append(l.Array, Financial{
			Name:       name,
			Price:      price,
			Budget:     budget,
			CreateTime: time.Now(),
		})

		return l.Array
	}
}

func (l *List) Delete(index int) []Financial {
	l.Array = append(l.Array[:index-1], l.Array[index:]...)

	return l.Array
}

func (l *List) UpdateName(index int, name string) []Financial {
	l.Array[index].Name = name

	return l.Array
}
func (l *List) UpdatePrice(index int, price float64) []Financial {
	l.Array[index].Price = price

	return l.Array
}
