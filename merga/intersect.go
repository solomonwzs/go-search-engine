package merga

type DataLQueue interface {
	Pop() (interface{}, bool)
	Head() (interface{}, bool)
	Len() int
}

type Intersect struct {
	funcComp FuncComp
	queues   []DataLQueue
	si       int
	end      bool
}

func NewIntersect(queues []DataLQueue, f FuncComp) *Intersect {
	in := &Intersect{
		funcComp: f,
		queues:   queues,
		si:       0,
		end:      false,
	}

	l := -1
	for i, q := range queues {
		ql := q.Len()
		if l == -1 || l < ql {
			l = ql
			in.si = i
		}
	}

	return in
}

func (in *Intersect) Extract() (ret interface{}, ok bool) {
	if in.end {
		return nil, false
	}

	for true {
		data, ok := in.queues[in.si].Pop()
		if !ok {
			in.end = true
			break
		}

		c := true
		for i := 0; i < len(in.queues) && c; i++ {
			if i == in.si {
				continue
			}

			for true {
				data0, ok0 := in.queues[i].Head()
				if !ok0 {
					in.end = true
					break
				}

				cmp := in.funcComp(data, data0)
				if cmp == 0 {
					in.queues[i].Pop()
					break
				} else if cmp > 0 {
					in.queues[i].Pop()
				} else {
					c = false
					break
				}
			}
		}

		if c {
			return data, true
		}
	}

	return
}
