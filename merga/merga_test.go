package merga

import (
	"fmt"
	"testing"
)

type intList struct {
	offset int
	arr    []int
}

func newIntList(arr []int) *intList {
	return &intList{
		offset: 0,
		arr:    arr,
	}
}

func (l *intList) Pop() (interface{}, bool) {
	if l.offset < len(l.arr) {
		l.offset += 1
		return l.arr[l.offset-1], true
	} else {
		return nil, false
	}
}

func (l *intList) Len() int {
	return len(l.arr) - l.offset
}

func (l *intList) Head() (interface{}, bool) {
	if l.offset < len(l.arr) {
		return l.arr[l.offset], true
	} else {
		return nil, false
	}
}

func win(a, b interface{}) int8 {
	if a.(int) < b.(int) {
		return -1
	} else if a.(int) == b.(int) {
		return 0
	} else {
		return 1
	}
}

var (
	listArray  []DataQueue
	listArray2 []DataLQueue
)

func init() {
	listArray = make([]DataQueue, 16, 16)
	listArray[0] = newIntList([]int{87, 99, 104, 119})
	listArray[1] = newIntList([]int{48, 56, 88, 97})
	listArray[2] = newIntList([]int{98, 104, 128, 151})
	listArray[3] = newIntList([]int{58, 70, 76, 100})
	listArray[4] = newIntList([]int{33, 91, 156, 205})
	listArray[5] = newIntList([]int{48, 55, 60, 68})
	listArray[6] = newIntList([]int{44, 55, 66, 77})
	listArray[7] = newIntList([]int{80, 96, 106, 113})
	listArray[8] = newIntList([]int{87, 99, 104, 119})
	listArray[9] = newIntList([]int{48, 56, 88, 97})
	listArray[10] = newIntList([]int{98, 104, 128, 151})
	listArray[11] = newIntList([]int{58, 70, 76, 100})
	listArray[12] = newIntList([]int{33, 91, 156, 205})
	listArray[13] = newIntList([]int{48, 55, 60, 68})
	listArray[14] = newIntList([]int{44, 55, 66, 77})
	listArray[15] = newIntList([]int{80, 96, 106, 113})

	listArray2 = make([]DataLQueue, 4, 4)
	listArray2[0] = newIntList([]int{1, 3, 4, 5, 7})
	listArray2[1] = newIntList([]int{3, 4, 6, 7})
	listArray2[2] = newIntList([]int{1, 3, 4, 7})
	listArray2[3] = newIntList([]int{1, 4, 7})
}

func TestWinerTree(t *testing.T) {
	for i, _ := range listArray {
		listArray[i].(*intList).offset = 0
	}

	tree := NewWinerTree(listArray[:], win)
	n := 0
	m := 0
	for true {
		if i, ok := tree.Extract(); !ok {
			break
		} else if i.(int) < n {
			t.Fatal("Error")
		} else {
			n = i.(int)
			m += 1
		}
	}
	if m != 16*4 {
		t.Fatal("Error")
	}
}

func TestLoserTree(t *testing.T) {
	for i, _ := range listArray {
		listArray[i].(*intList).offset = 0
	}

	tree := NewLoserTree(listArray[:], win)
	n := 0
	m := 0
	for true {
		if i, ok := tree.Extract(); !ok {
			break
		} else if i.(int) < n {
			t.Fatal("Error")
		} else {
			n = i.(int)
			m += 1
		}
	}
	if m != 16*4 {
		t.Fatal("Error")
	}
}

func TestIntersect(t *testing.T) {
	for i, _ := range listArray2 {
		listArray2[i].(*intList).offset = 0
	}

	in := NewIntersect(listArray2, win)
	for true {
		if i, ok := in.Extract(); !ok {
			break
		} else {
			fmt.Println(i)
		}
	}
}

func BenchmarkWinerTree(b *testing.B) {
	for k := 0; k < b.N; k++ {
		for i, _ := range listArray {
			listArray[i].(*intList).offset = 0
		}

		tree := NewWinerTree(listArray, win)
		ok := true
		for ok {
			_, ok = tree.Extract()
		}
	}
}

func BenchmarkLoserTree(b *testing.B) {
	for k := 0; k < b.N; k++ {
		for i, _ := range listArray {
			listArray[i].(*intList).offset = 0
		}

		tree := NewLoserTree(listArray, win)
		ok := true
		for ok {
			_, ok = tree.Extract()
		}
	}
}
