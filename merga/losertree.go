package merga

type LoserTree struct {
	funcWin FuncComp
	nodes   []*mergaTreeNode
	queues  []DataQueue
	winer   *mergaTreeNode
}

func NewLoserTree(queues []DataQueue, f FuncComp) *LoserTree {
	n := len(queues)*2 - 1

	t := &LoserTree{
		funcWin: f,
		nodes:   make([]*mergaTreeNode, n, n),
		queues:  queues,
	}

	for i := n - 1; i >= n/2; i-- {
		which := i + 1 - len(queues)
		data, ok := queues[which].Pop()

		t.nodes[i] = new(mergaTreeNode)
		if ok {
			t.nodes[i].which = which
			t.nodes[i].data = data
		} else {
			t.nodes[i].setNone()
		}
	}
	t.winer = t.value(0)

	return t
}

func (t *LoserTree) value(i int) (n *mergaTreeNode) {
	if t.isLeaf(i) {
		return t.nodes[i]
	} else {
		left := nLeft(i)
		right := nRight(i)

		ln := t.value(left)
		rn := t.value(right)

		if nodeWin(t.funcWin, ln, rn) {
			n = ln
			t.nodes[i] = rn
		} else {
			n = rn
			t.nodes[i] = ln
		}

		return
	}
}

func (t *LoserTree) Extract() (ret interface{}, ok bool) {
	if t.winer.isNone() {
		return nil, false
	}
	ret, ok = t.winer.data, true

	which := t.winer.which
	ni := which + len(t.queues) - 1
	data, ok0 := t.queues[which].Pop()

	if ok0 {
		t.nodes[ni].data = data
	} else {
		t.nodes[ni].setNone()
	}

	for ni > 0 {
		parent := nParent(ni)
		if !nodeWin(t.funcWin, t.winer, t.nodes[parent]) {
			t.winer, t.nodes[parent] = t.nodes[parent], t.winer
		}
		ni = parent
	}

	return
}

func (t *LoserTree) isLeaf(i int) bool {
	return i >= len(t.nodes)/2
}
