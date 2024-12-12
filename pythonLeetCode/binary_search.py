from typing import List



class Solution:
    @staticmethod
    def binary_search(nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:

            mid = int((left + right) / 2)
            if nums[mid] == target:
                return mid
            elif target  < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        return -1


sorted_nums = [-1,0,3,5,9,12]
target = 9
s = Solution()
print(s.binary_search(sorted_nums, target))