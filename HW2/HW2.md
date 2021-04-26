# HW2

1. For each of the following assertions, say whether it is true or false and support your answer with examples or counterexamples where appropriate. 

   + (a) An agent that senses only partial information about the state cannot be perfectly rational. 
     > 錯, perfectly rational 是指在收到傳感器訊息的情況下，能夠做出正確的決策。
     >
     
   + (b) There exist task environments in which no pure reflex agent can behave rationally.

     > 對，pure reflex agent 忽略了以前的 percept，所以在有些情況下無法達到 rational。  

    + (c) The input to an agent program is the same as the input to the agent function. 

      > 錯，agent function 會把 percept 到的所有資料做為輸入；而 agent program 只會把當前的 percept 做為輸入

    + (d) Every agent is rational in an unobservable environment. 

      > 錯，we can find an agent is not rational in an unobservable environment. 比如說一個什麼都不做的代理人

    + (e) A perfectly rational poker-playing agent never loses

      > 錯，撲克牌有些是跟運氣成份(機率)有關的，比如抽鬼牌，10點半

2. Give a complete problem formulation for each of the following. Choose a formulation that is precise enough to be implemented.

    + (a) Using only four colors, you have to color a planar map in such a way that no two adjacent regions have the same color. 
      + initial state: 空平面
      + actions: 對其中一個色塊著色
      + transition model: 有新色塊被塗色，預計不塗與相鄰色塊相同的顏色
      + goal test: 檢查是否有相鄰色塊同色
      + path cost: 著色的次數
    + (b) A 3-foot-tall monkey is in a room where some bananas are suspended from the 8-foot ceiling. He would like to get the bananas. The room contains two stackable, movable, climbable 3-foot-high crates.
      + initial state: X-Y 平面上有兩個不重疊的箱子
      + actions: 移動箱子、疊箱子、爬箱子
      + transition model: 移動箱子使箱子在 X-Y 平面上做移動，疊箱子使箱子在 Z  軸上移動，只有在兩箱子於 Z 軸上重疊時才做爬箱子的動作
      +  goal test: 猴子是否有拿到香蕉
      +  path cost: 計算移動箱子的距離、疊箱子、爬箱子的總次數
    + (c) You have a program that outputs the message “illegal input record” when fed a certain file of input records. You know that processing of each record is independent of the other records. You want to discover what record is illegal.
      + initial state: 將欲檢索之檔案餵給該程式
      + actions: 觀察
      + transition model: 觀察程式是否噴出錯誤訊息
      + goal test: 當程式檢索完時
      + path cost: 取決於掃描記錄的數量
    + (d) You have three jugs, measuring 12 gallons, 8 gallons, and 3 gallons, and a water faucet. You can fill the jugs up or empty them out from one to another or onto the ground. You need to measure out exactly one gallon.
      + initial state: 三個水壺皆為空
      + actions: 注水、倒水
      + transition model: 注水可將一壺水由空壺注滿，倒水則可將非空之壺的水倒入一個未滿的壺或倒出至地面
      + goal test: 其中一壺水恰好有 1 加侖
      + path cost: 注水與倒水的次數 

3. Define in your own words the following terms: state, state space, search tree, search node, goal, action, transition model, and branching factor.

    + state: 環境中物件的狀態，由解題者定義
    + state space: 環境中所有物件可能的狀態的集合 (可由state+action+transition model推得)
    + search tree: 以 initial state 為 root，接著將下一個可能的 state 當做 node，一直重複往下擴張
    + search node: 每個可能發生的 state 形成一個 node
    + goal: 解題最後所呈現的狀態
    + action: 狀態與狀態間的動作
    + transition model: 經由一個動作後會造成前一個狀態如何改變成下一個狀態
    + branching factor: 一個已知狀態經過一個動作會變成其他狀態的所有可能性的數量；亦即搜尋樹上某一節點的子節點數量

4. Which of the following are true and which are false? Explain your answers. 

    + (a) Depth-first search always expands at least as many nodes as A∗ search with an admissible heuristic.

      > False, DFS 有時可以剛好找到答案，而 A* 要考慮其他條路徑來確保正確性

    + (b) h(n) = 0 is an admissible heuristic for the 8-puzzle. 

      > True, h(n) = 0 時仍可以走到 goal state，只是這樣一來就沒有 heuristic 的意義了

    + (c) A* is of no use in robotics because percepts, states, and actions are continuous. 

      > False, 天生我材必有用, 雖然 A\* 是離散的搜尋，但我們仍可以將連續路徑離散化

    + (d) Breadth-first search is complete even if zero step costs are allowed.

      > True, but it's not really true. 只要 depth 是有限的當然能走到終點，但如果深度無限則永遠也走不到

    + (e) Assume that a rook can move on a chessboard any number of squares in a straight line, vertically or horizontally, but cannot jump over other pieces. Manhattan distance is an admissible heuristic for the problem of moving the rook from square A to square B in the smallest number of moves.

      > False, 因為「車」可以一次移動好幾個方格，用曼哈頓距離評估會高估他所要的移動次數

5. Which of the following are true and which are false? Give brief explanations. 

    + (a) In a fully observable, turn-taking, zero-sum game between two perfectly rational players, it does not help the first player to know what strategy the second player is using— that is, what move the second player will make, given the first player’s move. 

      > True，因為是 fully observable，第一個玩家與第二個玩家所知道的事情完全相同，並且每個玩家都是 perfectly rational players，那麼其實第一個玩家早就可以事先知道第二個玩家能做什麼事了。

    + (b) In a partially observable, turn-taking, zero-sum game between two perfectly rational players, it does not help the first player to know what move the second player will make, given the first player’s move. 

      > False，第一個玩家若能提前知道第二個玩家下一步動作，勢必能減少計算上的複雜度。比如「吹牛」、「九九」若能先知道對方下一步出什麼牌則對計算上一定有幫助。

    + (c) A perfectly rational backgammon agent never loses.

      > False，backgammon 這個遊戲是一種 stochastic 遊戲，不可能永遠都不會輸
