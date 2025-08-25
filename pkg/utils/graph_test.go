package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphHasCycle(t *testing.T) {
	testcases := []struct {
		name   string
		graph  KahnGraph
		except bool
	}{
		{
			// a -> b
			name: "单头单深度无环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"b"},
				},
				indegree: map[string]int{
					"a": 0,
					"b": 1,
				},
			},
			except: false,
		},
		{
			// a -> a
			name: "单头单深度有环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"a"},
				},
				indegree: map[string]int{
					"a": 1,
				},
			},
			except: true,
		},
		{
			// a -> b
			// a -> c
			name: "多头单深度无环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"b", "c"},
				},
				indegree: map[string]int{
					"a": 0,
					"b": 1,
					"c": 1,
				},
			},
			except: false,
		},
		{
			// a -> b
			// a -> a
			name: "多头单深度有环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"b", "a"},
				},
				indegree: map[string]int{
					"a": 1,
					"b": 1,
				},
			},
			except: true,
		},
		{
			// a -> b -> c
			name: "单头多深度无环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"b"},
					"b": []string{"c"},
				},
				indegree: map[string]int{
					"a": 0,
					"b": 1,
					"c": 1,
				},
			},
			except: false,
		},
		{
			// a -> b -> a
			name: "单头多深度有环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"b"},
					"b": []string{"a"},
				},
				indegree: map[string]int{
					"a": 1,
					"b": 1,
				},
			},
			except: true,
		},
		{
			// a -> b
			// a -> c -> d
			// a -> d -> b
			name: "多头多深度无环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"b", "c", "d"},
					"c": []string{"d"},
					"d": []string{"b"},
				},
				indegree: map[string]int{
					"a": 0,
					"b": 2,
					"c": 1,
					"d": 2,
				},
			},
			except: false,
		},
		{
			// a -> b
			// a -> c -> d
			// a -> d -> c
			name: "多头多深度有环",
			graph: KahnGraph{
				edges: map[string][]string{
					"a": []string{"b", "c", "d"},
					"c": []string{"d"},
					"d": []string{"c"},
				},
				indegree: map[string]int{
					"a": 0,
					"b": 1,
					"c": 2,
					"d": 2,
				},
			},
			except: false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.except, tc.graph.hasCycle())
		})
	}
}
