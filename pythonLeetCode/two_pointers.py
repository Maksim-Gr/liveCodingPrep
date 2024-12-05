from typing import List


class Solution:
    def twoSum(self, nums: List[int], target:int) -> List[int]:
        left, right = 0, len(nums)-1
        while left < right:
            current_sum  = nums[left] + nums[right]
            if current_sum == target:
                return [left, right]
            
            
            if target > current_sum:
                left +=1
            else:
                right -= 1
        return [-1,-1]
    
    