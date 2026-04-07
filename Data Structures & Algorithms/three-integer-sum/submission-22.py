class Solution:
    def kSum(self, nums, start, target, k):
        res = []
        n = len(nums)

        if k == 2:
            # 2 sum
            left, right = start, n - 1

            while left < right:
                curr = nums[left] + nums[right]
                if curr == target:
                    res.append([nums[left], nums[right]])
                    left += 1
                    right -= 1

                    while left < right and nums[left] == nums[left-1]:
                        left += 1
                elif curr < target:
                    left +=1 
                else:
                    right -= 1
        else:
            for i in range(start, n - k + 1):
                if i > start and nums[i] == nums[i - 1]:
                    continue

                subsets = self.kSum(nums, i + 1, target - nums[i], k - 1)

                for subset in subsets:
                    res.append([nums[i]] + subset)

        return res


    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        target = 0
        k = 3
        start = 0

        return self.kSum(nums, start, target, k)