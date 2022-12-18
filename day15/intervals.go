package day15

type Interval struct {
	from, to int
}

type Intervals []Interval

func (is *Intervals) Add(from, to int) {
	// not the optimal implementation, but concise and fast enough for this case
	var l, r Intervals
	for _, interval := range *is {
		if interval.to < from {
			l = append(l, interval)
		} else if interval.from > to {
			r = append(r, interval)
		} else {
			if interval.from < from {
				from = interval.from
			}
			if interval.to > to {
				to = interval.to
			}
		}
	}
	*is = make(Intervals, len(l)+len(r)+1)
	copy(*is, l)
	(*is)[len(l)] = Interval{from, to}
	copy((*is)[len(l)+1:], r)
}
