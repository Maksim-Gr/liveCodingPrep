from typing import List


class Solution:
    @staticmethod
    def rotate(nums: List[int], n: int) -> List[int]:
        length = len(nums)
        n = n % length
        if n < 0:
            n = length - n
        # reverse the whole array
        Solution.reverse_array(nums, 0, len(nums) - 1)
        # fixing the order of rotated elements
        Solution.reverse_array(nums, 0, n - 1)
        # fixing the order of the remaining elements
        Solution.reverse_array(nums, 0, length - 1)

        return nums

    @classmethod
    def reverse_array(cls, nums: List[int], start: int, end: int) -> None:
        while start < end:
            nums[start], nums[end] = nums[end], nums[start]
            start += 1
            end -= 1
