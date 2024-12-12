from typing import List


class Solution:
    @staticmethod
    def two_sum(nums: List[int], target:int) -> List[int]:
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
    @staticmethod
    def find_sum(nums: List[ int], k: int) -> List[int]:
        nums.sort()

        left = 0
        right = len(nums)-1

        while left < right:
            sum_val = nums[left] + nums[right]
            if sum_val == k:
                return [left, right]
            elif sum_val < k:
                left += 1
            else:
                right -= 1


            return [nums[left], nums[right]]
    