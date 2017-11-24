package merga

type DataLQueue interface {
	DataQueue
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

	return
}
