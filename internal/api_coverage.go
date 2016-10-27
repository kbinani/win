package internal

type APICoverage struct {
	Dll      string
	Coverage *Rat
}

type APICoverages []APICoverage

func (this APICoverages) Len() int {
	return len(this)
}

func (this APICoverages) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this APICoverages) Less(i, j int) bool {
	covI := this[i].Coverage.Float64()
	covJ := this[j].Coverage.Float64()
	if covI == covJ {
		return this[i].Coverage.Denom > this[j].Coverage.Denom
	} else {
		return covI > covJ
	}
}
