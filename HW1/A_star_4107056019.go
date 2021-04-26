package main

import (
	"container/heap"
	"fmt"
	"math"
)

const far = 999999
const maxRow = 6
const maxCol = 6

var rowShift = [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
var colShift = [8]int{-1, -1, -1, 0, 0, 1, 1, 1}

type Node struct {
	row                              int
	col                              int
	previous                         *Node
	distFromStartToThis              float64
	predictDistFromStartToEndViaThis float64
}

type MinHeap []*Node

//--------------------------------------------------------------------------------------
// Implement interface for heap
//
// type Interface interface {
//     sort.Interface
//     Push(x interface{}) // add x as element Len()
//     Pop() interface{}   // remove and return element Len() - 1.
// }

func (this MinHeap) Len() int { return len(this) }
func (this MinHeap) Less(i, j int) bool {
	return this[i].predictDistFromStartToEndViaThis < this[j].predictDistFromStartToEndViaThis
}
func (this MinHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func (this *MinHeap) Push(x interface{}) {
	new := x.(*Node)
	*this = append(*this, new)
}
func (this *MinHeap) Pop() interface{} {
	old := *this
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*this = old[0 : n-1]
	return item
}

//--------------------------------------------------------------------------------------

func main() {
	rawMap := [maxCol][maxRow]int{
		{1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 0, 0},
		{1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 1},
		{0, 1, 1, 1, 1, 0},
		{1, 1, 1, 1, 1, 0},
	}

	// key of map is row*maxRow + col
	nodeSet := map[int]*Node{}
	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			if rawMap[i][j] == 1 {
				newNode := Node{
					row:                              i,
					col:                              j,
					distFromStartToThis:              far,
					predictDistFromStartToEndViaThis: far,
				}
				nodeSet[i*maxRow+j] = &newNode
			}
		}
	}

	start := nodeSet[0]
	end := nodeSet[5*maxRow+2]
	start.distFromStartToThis = 0
	start.predictDistFromStartToEndViaThis = start.huristic(end)

	discoveredNode := MinHeap{}
	discoveredNode = append(discoveredNode, start)
	heap.Init(&discoveredNode)

	current := start
	success := false
	for discoveredNode.Len() > 0 {
		current = heap.Pop(&discoveredNode).(*Node)
		if current == end {
			success = true
			break
		}

		// find neighbor of current
		// 0 1 2
		// 3   4
		// 5 6 7
		for i := 0; i < 8; i++ {
			if !(current.row+rowShift[i] < maxRow && current.row+rowShift[i] >= 0 &&
				current.col+colShift[i] < maxCol && current.col+colShift[i] >= 0) {
				continue
			}
			neighbor, ok := nodeSet[(current.row+rowShift[i])*maxRow+current.col+colShift[i]]
			if ok {
				// tentative distance from start to neighbor (via current or not)
				var tentative float64
				if i == 0 || i == 2 || i == 5 || i == 7 {
					tentative = current.distFromStartToThis + 1.4
				} else {
					tentative = current.distFromStartToThis + 1
				}

				//fmt.Printf("(%d, %d) t: %f o: %f\n", neighbor.row, neighbor.col, tentative, neighbor.distFromStartToThis)
				if tentative < neighbor.distFromStartToThis {
					neighbor.distFromStartToThis = tentative
					neighbor.previous = current
					neighbor.predictDistFromStartToEndViaThis = neighbor.distFromStartToThis +
						neighbor.huristic(end)
					heap.Push(&discoveredNode, neighbor)
				}
			}
		}
	}
	if !success {
		fmt.Print("Can not find a path from start to end")
		return
	}
	reverse := []*Node{}
	for pointer := end; pointer != nil; pointer = pointer.previous {
		reverse = append(reverse, pointer)
	}
	for i := 0; i < (len(reverse) >> 1); i++ {
		reverse[i], reverse[len(reverse)-1-i] = reverse[len(reverse)-1-i], reverse[i]
	}
	for _, v := range reverse {
		fmt.Printf("(%d, %d)", v.row, v.col)
	}
}

func (from *Node) huristic(to *Node) float64 {
	return math.Sqrt(float64((from.row-to.row)*(from.row-to.row) + (from.col-to.col)*(from.col-to.col)))
}
