package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const min = true
const max = false

type state struct {
	graph []int8
	val   int8
	alpha int8
	beta  int8
}

func newState() state {
	return state{
		alpha: -128,
		beta:  127,
	}
}

func main() {
	// 0  : empty
	// -1 : O AI
	// 1  : X human

	t := []int8{0, 0, 0, 0, 0, 0, 0, 0, 0}
	halt := false
	round := 0
	rand.Seed(time.Now().UnixNano())

	for !halt {
		fmt.Printf("Round %d\n", round)
		showTicTacToe(t)
		if round&1 == 0 { // human round
			var where int
			fmt.Printf("Your move: ")
			_, err := fmt.Scanf("%d\n", &where)
			for err != nil {
				fmt.Fprintf(os.Stderr, "input error, %v try again:", err)
				_, err = fmt.Scanf("%d", &where)
			}
			if(where == 0){
				fmt.Println("You give up. The program is terminated")
				os.Exit(0)
			}
			for where < 1 || where > 9 || t[where-1] != 0 {
				fmt.Fprintf(os.Stderr, "input error, try again:")
				_, err = fmt.Scanf("%d", &where)
			}
			t[where-1] = 1
		} else { // AI round
			// find the best step
			current := newState()
			current.graph = t
			next := minimax(current, min)
			var move int
			for i, v := range next.graph {
				if v != t[i] {
					move = i
					break
				}
			}
			t = next.graph
			fmt.Println("AI's move:", move+1)
		}

		fmt.Println()
		round++

		if done, res := goalTest(t); done {
			showTicTacToe(t)
			if res < 0 {
				fmt.Println("AI win, 回家再練10年吧")
			} else if res == 0 {
				fmt.Println("tie, 讓你一把")
			} else {
				fmt.Println("AI 是不會輸的 ^_^")
			}
			halt = true
		}
	}
}

func showTicTacToe(this []int8) {
	for k, v := range this {
		switch v {
		case 0:
			fmt.Printf("|%d", k+1)
		case 1:
			fmt.Printf("|X")
		case -1:
			fmt.Printf("|O")
		}
		if k%3 == 2 {
			fmt.Println("|")
		}
	}
}

func minimax(parent state, minMax bool) state {
	var ret state
	initRet := true
	for i, v := range parent.graph {
		// Generate all possible chlidren nodes
		if v == 0 {
			child := newState()
			child.graph = make([]int8, len(parent.graph))
			// copy
			for i, v := range parent.graph {
				child.graph[i] = v
			}

			// max -> fill with 1
			// min -> fill with -1
			if minMax == max {
				child.graph[i] = 1
			} else {
				child.graph[i] = -1
			}

			done, val := goalTest(child.graph)
			child.alpha = parent.alpha // copy alpha value to child's alpha
			child.beta = parent.beta   // copy beta value to child's beta
			if !done {
				child.val = minimax(child, !minMax).val
			} else {
				child.val = val
			}

			if initRet {
				ret = child
				initRet = false
			}

			// min level 更新 parent beta
			// max level 更新 parent alpha
			if minMax == min && child.val < ret.val {
				ret = child
				parent.beta = child.val
			}

			if minMax == max && child.val > ret.val {
				ret = child
				parent.alpha = child.val
			}

			if child.val == ret.val && rand.Float32() > .5 {
				ret = child
			}

			// min level 的 val 值若比 min 的 alpha 小則忽略其他節點
			// max level 的 val 值若比 max 的 beta 大則忽略其他節點
			if minMax == min && child.val < parent.alpha {
				break
			}
			if minMax == max && child.val > parent.beta {
				break
			}
		}
	}
	parent.val = ret.val

	return ret
}

func innerProduct(a []int8, b []int8) int8 {
	var sum int8 = 0
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum
}

// @return[0] if the state is at destination
// @return[1] -1 when AI `O` win
//             1 when human `X` win
//             0 when tie
func goalTest(state []int8) (bool, int8) {
	winState := [][]int8{
		{1, 1, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 1},
		{1, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 0, 1, 0, 0},
	}

	for _, oneOfGoalState := range winState {
		res := innerProduct(oneOfGoalState, state)
		if res == -3 {
			return true, -3
		} else if res == 3 {
			return true, 3
		}
	}

	var product int8 = 1
	for _, k := range state {
		product *= k
	}
	if product == 0 { // have not done
		return false, 0
	}
	return true, 0 // tie
}
