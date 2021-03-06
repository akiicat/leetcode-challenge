package main
import . "main/pkg/tree_node"

// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

// leetcode 104.
// T: O(n)
// M: O(h)
// -- start --

func diameterOfBinaryTree(root *TreeNode) int {
  max := 0
  maxDepth(root, &max)
  return max
}

// 0104-maximum-depth-of-binary-tree
func maxDepth(root *TreeNode, max *int) int {
  if root == nil {
    return 0
  }
  l := maxDepth(root.Left, max)
  r := maxDepth(root.Right, max)
  *max = Max(*max, l + r)
  return Max(l, r) + 1
}

func Max(a, b int) int {
  if a > b {
    return a
  }
  return b
}

// -- end --

