from typing import List


class Solution:
    def move_zeros_to_the_beginning(self, nums: List[int]) -> None:
        if len(nums) < 1:
            return

        length_nums = len(nums)

        write_index = length_nums - 1
        read_index = length_nums - 1


        while read_index >= 0:
            if nums[read_index] != 0:
                nums[write_index] = nums[read_index]
                write_index -= 1
            read_index -= 1

        while write_index >= 0:
            nums[write_index] = 0
            write_index -= 1
