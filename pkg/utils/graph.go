package utils

// KahnGraph 表示一个有向图，并通过 Kahn 算法高效检测环路。
// Kahn 算法不断移除入度为 0 的节点（即没有依赖的节点）。
// 如果存在环路，则环上的节点永远不会变为入度为 0，算法无法处理所有节点，从而检测出环的存在。
type KahnGraph struct {
	// edges 是邻接表：每个节点存储其所有指向的目标节点（出边）。
	// key: 源节点，value: 目标节点切片
	edges map[string][]string

	// indegree 记录每个节点的入度（被指向的次数）。
	// key: 节点名，value: 入度计数
	indegree map[string]int
}

// NewKahnGraph 初始化并返回一个空的 KahnGraph。
func NewKahnGraph() *KahnGraph {
	return &KahnGraph{
		edges:    make(map[string][]string),
		indegree: make(map[string]int),
	}
}

// AddEdgeAndCheckCycle 向图中添加一条从节点 a 到节点 b 的有向边，并立即检测是否形成环路。
// 参数：
//   - a: 源节点
//   - b: 目标节点
//
// 返回值：
//   - 如果添加该边后形成环路，返回 true；否则返回 false
func (g *KahnGraph) AddEdgeAndCheckCycle(a, b string) bool {
	// 添加从 a 到 b 的边
	g.edges[a] = append(g.edges[a], b)

	// b 的入度加 1（多了一个指向 b 的边）
	g.indegree[b]++

	// 确保 a 在入度表中（如果没有则初始化为 0）
	if _, exists := g.indegree[a]; !exists {
		g.indegree[a] = 0
	}

	// 添加新边后检测是否有环
	return g.hasCycle()
}

// hasCycle 判断当前图中是否存在环路。
// 返回值：
//   - 存在环路返回 true，否则返回 false
func (g *KahnGraph) hasCycle() bool {
	// 拷贝一份入度表，避免修改原始数据
	indegreeCopy := make(map[string]int, len(g.indegree))
	for node, degree := range g.indegree {
		indegreeCopy[node] = degree
	}

	// 初始化队列，收集所有入度为 0 的节点
	queue := make([]string, 0)
	for node, degree := range indegreeCopy {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	processed := 0 // 记录被拓扑排序处理的节点数

	// 处理所有入度为 0 的节点，并更新其邻居的入度
	for len(queue) > 0 {
		// 出队第一个节点
		node := queue[0]
		queue = queue[1:]
		processed++

		// 遍历所有邻居，入度减 1
		for _, neighbor := range g.edges[node] {
			indegreeCopy[neighbor]--
			// 如果邻居入度变为 0，加入队列
			if indegreeCopy[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 如果所有节点都被处理，则无环；否则存在环路
	return processed != len(indegreeCopy)
}
