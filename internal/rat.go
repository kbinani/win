package internal

type Rat struct {
	Num, Denom int
}

func NewRat(num, denom int) *Rat {
	ret := new(Rat)
	ret.Num = num
	ret.Denom = denom
	return ret
}

func (r *Rat) Float64() float64 {
	return float64(r.Num) / float64(r.Denom)
}
