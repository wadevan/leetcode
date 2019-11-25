package golang

/**
给定一个 每个结点的值互不相同 的二叉树，和一个目标值 k，找出树中与目标值 k 最近的叶结点。

这里，与叶结点 最近 表示在二叉树中到达该叶节点需要行进的边数与到达其它叶结点相比最少。而且，当一个结点没有孩子结点时称其为叶结点。

在下面的例子中，输入的树以逐行的平铺形式表示。实际上的有根树 root 将以TreeNode对象的形式给出。

示例 1：

输入：
root = [1, 3, 2], k = 1
二叉树图示：
          1
         / \
        3   2

输出： 2 (或 3)

解释： 2 和 3 都是距离目标 1 最近的叶节点。


示例 2：

输入：
root = [1], k = 1
输出：1

解释： 最近的叶节点是根结点自身。


示例 3：

输入：
root = [1,2,3,4,null,null,null,5,null,6], k = 2
二叉树图示：
             1
            / \
           2   3
          /
         4
        /
       5
      /
     6

输出：3
解释： 值为 3（而不是值为 6）的叶节点是距离结点 2 的最近结点。


注：

root 表示的二叉树最少有 1 个结点且最多有 1000 个结点。
每个结点都有一个唯一的 node.val ，范围为 [1, 1000]。
给定的二叉树中有某个结点使得 node.val == k。

*/
func findClosestLeaf(root *TreeNode, k int) int {
	graph := map[int]map[int]struct{}{}
	treeToGraph(root, graph)
	visited := map[int]struct{}{k: {}}
	q := []int{k}
	for len(q) != 0 {
		obj := q[0]
		q = q[1:]

		if graph[obj] == nil {
			return obj
		}
		for k, _ := range graph[obj] {
			print(k)
			if _, ok := visited[k]; !ok {
				if graph[k] == nil {
					return k
				}
				q = append(q, k)
				visited[k] = struct{}{}
			}
		}
	}
	return -1
}
func treeToGraph(root *TreeNode, graph map[int]map[int]struct{}) {
	addEdge := func(x int, edge int) {
		list, ok := graph[x]
		if !ok {
			list = map[int]struct{}{}
			graph[x] = list
		}
		list[edge] = struct{}{}
	}
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		graph[root.Val] = nil
		return
	}
	if root.Left != nil {
		addEdge(root.Val, root.Left.Val)
		addEdge(root.Left.Val, root.Val)
		treeToGraph(root.Left, graph)
	}
	if root.Right != nil {
		addEdge(root.Val, root.Right.Val)
		addEdge(root.Right.Val, root.Val)
		treeToGraph(root.Right, graph)
	}
}
