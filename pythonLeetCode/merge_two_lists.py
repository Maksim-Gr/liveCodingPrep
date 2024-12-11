from typing import List


class Solution:
    @staticmethod
    def merge_two_sorted_lists(nums1: List[int], nums2: List[int]) -> List[int]:
        result = [None] * (len(nums1) + len(nums2))
        p1 = 0
        p2 = 0
        p3 = 0

        while p1 < len(nums1) and p2 < len(nums2):
            if nums1[p1] < nums2[p2]:
                result[p3] = nums1[p1]
                p1 += 1
                p3 += 1
            else:
                result[p3] = nums2[p2]
                p2 += 1
                p3 += 1
        while p1 < len(nums1):
            result[p3] = nums1[p1]
            p1 += 1
            p3 += 3
        while p2 < len(nums2):
            result[p3] = nums2[p2]
            p2 += 1
            p3 += 1
        return result


nums1  = [23, 33, 35, 41, 44, 47, 56, 91, 105]
nums2 = [12, 34, 45, 56, 67, 78, 89, 99]


s = Solution()
print(s.merge_two_sorted_lists(nums1, nums2))