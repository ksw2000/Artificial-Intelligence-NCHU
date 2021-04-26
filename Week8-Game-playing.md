# Game Playing

## Adversarial Search

Adversarial search (對抗搜尋) is search when there is an "enemy" or "opponent" changing the state of the problem every step in direction you do not want.

## Minimax: Recursive implementation

The minimax algorithm tries to predict the opponent's behaviour. It predicts the opponent will take the *worst* action from our viewpoint.

For every two-person, zero-sum game with finitely many strategies, there exists a value **V** and a mixed strategy for each player, s.t.

+ Given player 2's strategy, the best payoff possible for player is V
+ Given player 2's strategy, the best payoff possible for player is -V

Equivalently, Player 1's strategy guarantees them a payoff of V regardless of Player 2's strategy, and similarly Player  2 can guarantee themselves a payoff of -V.

## Alpha-beta pruning

The problem with minimax search is that the number of game states it has to examine is exponential in the depth of the tree. Unfortunately, we can’t eliminate the exponent, but it turns out we can effectively cut it in half.

α: Best choice so far for MAX

β: Best choice so far for MIN

Alpha-Beta 剪枝是 Minimax 對局搜尋法的一個修改版，主要是在 Minimax 當中加入了 α 與 β 兩個紀錄值，用來做為是否要修剪的參考標準。兩個參數以交錯的方式傳遞給下層的子樹。

- 在**最大層**取最大值的時候，若發現了一個大於等於 β 的值，就不用再對其它分枝進行搜尋，這就是所謂的 **β 剪枝**。
- 在**最小層**取最小值的時候，發現了一個小於等於 α 的值，也不用再對其它分枝進行搜尋，這就是所謂的 **α 剪枝**。

*reference*

1. [Adversarial Search - computing.dcu.ie](https://computing.dcu.ie/~humphrys/Notes/AI/adversarial.search.html#:~:text=Adversarial%20search%20is%20search%20when,unpredictable)
2. [Minimax Algorithm and Alpha-beta Pruning | Mr. Opengate](https://mropengate.blogspot.com/2015/04/ai-ch4-minimax-alpha-beta-pruning.html)

![](./src/alpha-beta-pruing.png)

green: max orange: min

左邊為 β 剪枝右邊為  α 剪枝