from collections import deque
from typing import List, Optional


class TreeNode:
    def __init__(self, x):
        self.x = x
        self.left = None
        self.right = None
        
        
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        result = []
        if root is None:
            return result
        
        queue= deque([root])
        while queue:
            level_size = len(queue)
            current_level = []
            for _ in range(level_size):
                current_node = queue.popleft()
                current_level.append(current_node.val)
                
                if current_node.left:
                    queue.append(current_node.left)
                if current_node.right:
                    queue.append(current_node.right)
                    
            result.append(current_level)
        return result