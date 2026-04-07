class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # can sort first
        res = []
        n = len(nums)
        target = 0
        nums.sort()

        for i in range(n-3+1):
            if i > 0 and nums[i] == nums[i-1]:
                continue

            first = nums[i]

            left, right = i + 1, n - 1

            while left < right:
                curr = first + nums[left] + nums[right]

                if curr == target:
                    res.append([first, nums[left], nums[right]])

                    left += 1
                    right -= 1

                    while left < right and nums[left] == nums[left-1]:
                        left += 1
                elif curr < target:
                    left += 1
                else:
                    right -= 1

        return res

