package main

import (
	"fmt"
)

// BouquetBuilder вместе с его методами - конкретный конструктор букетов.
// Каждое поле в нём - конкретный товар (например, SKU или ссылка на сайте).
type BouquetBuilder struct {
	tasty    string
	prikol   string
	utility  string
	surprise string
	filler   string
}

func (t BouquetBuilder) AddTasty() BouquetBuilder {
	t.tasty = GetRandS(SKU, Tasties)
	return t
}
func (t BouquetBuilder) AddPrikol() BouquetBuilder {
	t.prikol = GetRandS(SKU, Prikols)
	return t
}
func (t BouquetBuilder) AddUtility() BouquetBuilder {
	t.utility = GetRandS(SKU, Utils)
	return t
}
func (t BouquetBuilder) AddSurprise() BouquetBuilder {
	t.surprise = GetRandS(SKU, Surprises)
	return t
}
func (t BouquetBuilder) AddFiller() BouquetBuilder {
	t.filler = GetRandS(SKU, Fillers)
	return t
}

func (t BouquetBuilder) Build() Bouquet {
	return Bouquet{
		tasty:    t.tasty,
		prikol:   t.prikol,
		utility:  t.utility,
		surprise: t.surprise,
		filler:   t.filler,
	}
}

func NewBouquetBuilder1() BouquetBuilder {
	return BouquetBuilder{}
}

// Bouquet - структура готового букета
type Bouquet struct {
	tasty    string
	prikol   string
	utility  string
	surprise string
	filler   string
}

func (b Bouquet) String() string {
	return fmt.Sprintf("nice bouquet made from %s, %s, %s, %s",
		b.tasty, b.prikol, b.utility, b.surprise)
}
