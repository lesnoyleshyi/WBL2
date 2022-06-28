package main

type TicketBuilder struct {
	Tasty    int
	Prikol   int
	Utility  int
	Surprise int
	Filler   int
}

func (t TicketBuilder) AddTasty() TicketBuilder {
	t.Tasty = GetRandD(Price, Tasties)
	return t
}
func (t TicketBuilder) AddPrikol() TicketBuilder {
	t.Prikol = GetRandD(Price, Prikols)
	return t
}
func (t TicketBuilder) AddUtility() TicketBuilder {
	t.Utility = GetRandD(Price, Utils)
	return t
}
func (t TicketBuilder) AddSurprise() TicketBuilder {
	t.Surprise = GetRandD(Price, Surprises)
	return t
}
func (t TicketBuilder) AddFiller() TicketBuilder {
	t.Filler = GetRandD(Price, Fillers)
	return t
}

func (t TicketBuilder) Build() Ticket {
	return Ticket{
		Tasty:    t.Tasty,
		Prikol:   t.Prikol,
		Utility:  t.Utility,
		Surprise: t.Surprise,
		Filler:   t.Filler,
	}
}

func NewTicketBuilder1() TicketBuilder {
	return TicketBuilder{}
}

// Ticket - структура готового чека
type Ticket struct {
	Tasty    int
	Prikol   int
	Utility  int
	Surprise int
	Filler   int
}
