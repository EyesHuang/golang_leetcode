package leetcode_580

import (
	"reflect"
	"testing"

	"leetcode/util"
)

var tests572 = []struct {
	name     string
	root     *util.TreeNode
	subRoot  *util.TreeNode
	expected bool
}{
	{
		name: "Test 1",
		root: &util.TreeNode{
			Val: 3,
			Left: &util.TreeNode{
				Val: 4,
				Left: &util.TreeNode{
					Val: 1,
				},
				Right: &util.TreeNode{
					Val: 2,
				},
			},
			Right: &util.TreeNode{
				Val: 5,
			},
		},
		subRoot: &util.TreeNode{
			Val: 4,
			Left: &util.TreeNode{
				Val: 1,
			},
			Right: &util.TreeNode{
				Val: 2,
			},
		},
		expected: true,
	},
	{
		name: "Test 2",
		root: &util.TreeNode{
			Val: 3,
			Left: &util.TreeNode{
				Val: 4,
				Left: &util.TreeNode{
					Val: 1,
				},
				Right: &util.TreeNode{
					Val: 2,
					Left: &util.TreeNode{
						Val: 0,
					},
				},
			},
			Right: &util.TreeNode{
				Val: 5,
			},
		},
		subRoot: &util.TreeNode{
			Val: 4,
			Left: &util.TreeNode{
				Val: 1,
			},
			Right: &util.TreeNode{
				Val: 2,
			},
		},
		expected: false,
	},
}

func TestIsSubtree(t *testing.T) {
	for _, tt := range tests572 {
		t.Run(tt.name, func(t *testing.T) {
			res := isSubtree(tt.root, tt.subRoot)

			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("got %v, want %v", res, tt.expected)
			}
		})
	}
}
