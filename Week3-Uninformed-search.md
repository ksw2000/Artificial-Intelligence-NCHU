# Uninformed search

[TOC]

+ Uninformed search:
  
  **Uninformed search** is a searching technique that has no additional information about the distance from the current state to the goal. e.g. *breadth-first*, *uniform-cost*, *depth-first*, *bi-directional*, ...

+ Informed search
  
  **Informed search** is another technique that has additional information about the estimate distance from the current state to the goal. e.g. best-first, A\* Heuristic

```go
func Gegeral-Search(problem, queueing_func) (a solution or fail){
    nodes := make_quere(make_node(initial_state[problem]));
    for{
        if nodes is tmpey{
            return fail
        }
        nodes = remove_front(nodes);
        if goal_test() applied to state(node) succeeds{
            return ndoe
        }
        nodes = queueing_func(nodes, expand(node, operators[problem]))
    }
}
```

## BFS

<img src="./src/BFS.png" style="zoom:50%;" />

Time and space complexity are measured in the following terms:

+ b: max branching factor of the search tree
+ d: depth of the least-cost solution
+ m: max depth of the search tree (may be infinity)

Completeness(改演算法是否能達成找到最短路徑的目標): yes, if b is finite

Time complexity: $1+b+b^2+...+b^d=O(b^d)$

Space: $O(b^d)$ (由queue儲存的資料推算，通常比 DFS 差)

## UCS

Uniform Cost Search

從最小的路徑開始做走訪

> + Initialize **Queue** with root node (build from start state)
> 
> + Repeat until *Queue empty* or *first node has Goal state*
>   
>   + Remove first node form front of Queue
>   + Expand node (find its children)
>   + Reject those children that have already been considered, to avoid loop
>   + Add remaining children to Queue, in a way that keeps entire queue sorted by increasing path cost
> 
> + If Goal was reached, return success, otherwise failure

+ g(n) is the path cost to node n
+ b is branching factor
+ d is the depth of least-cost solution

Completeness: yes, if step cost is non-negative (不可有負環)

Time complexity: $O(b^d)$

Space complexity: $O(b^d)$ 

## DFS

<img src="./src/DFS.png" style="zoom:50%;" />

+ b: max branching factor of the search tree
+ d: depth of the least-cost solution
+ m: max depth of the search tree (may be infinity)

Completeness: yes, if d is finite

Time complexity: $b^d+b^{d+1}+...+1 = O(b^d)$

Space: $O(d\times b)$

> In our complexity analysis, we do not take the built-in loop-detection into account.

## Interactive deepening search

結合 BFS 及 DFS

+ Restrict a depth-first search to a fixed depth.

+ If no path was found, increase the depth and restart the search.

## Bidirectional search

除了從起始點開始找終點，同時從終點開始尋找起始點

---

## Comparing of Uniform search

|           | BFS | Uniform cost | DFS | Depth-limited              | Interactive<br> deepening | Biderctional |
| --------- | --- | ------------ | --- | -------------------------- | ------------------------- | ------------ |
| Time      | b^d | b^d          | b^m | b^l                        | b^d                       | b^(d/2)      |
| Space     | b^d | b^d          | bm  | bl                         | bd                        | b^(d/2)      |
| Optimal?  | Yes | Yes          | No  | No                         | Yes                       | Yes          |
| Complete? | Yes | Yes          | No  | Yes, if l is bigger than d | Yes                       | Yes          |

where,

+ b is max branching factor of the search tree 

+ d is depth of the least-cost solution

+ m is max depth of the state-space (may be infinity)

+ l (limit) is depth cutoff