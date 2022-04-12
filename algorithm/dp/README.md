# dynamic programming
dynamic programming，在查找有很多**重叠子问题**情况的**最优解**时很有效。避免多次重复解决子问题，将计算结果保存（精髓即在dp表）。

动态规划只应用于有最优子结构的问题，意思**局部最优解能决定全局最优解**。

|q|name|difficulty|desc|
|:-|:-|:-|:-|
|70.|climb stairs|easy|1-dimension dp, compressed to prev2, prev1|
|198.|rob|easy|1-dimension dp, compressed to prev2, prev1|
|413.|number of arithmetic slices|medium|1-dimension dp|
|64.|min path sum|medium|2-dimensions dp|
|300.|longest increasing subsequence|medium|subsequence, 1-dimension dp|
|1143.|longest common subsequence|medium|subsequence, 2-dimensions dp|
|416.|partition equal subset sum|medium|bag, 2-dimensions dp|


*两道子序列的题目时间复杂度都是大于O(n^2)的, 可以记住*
