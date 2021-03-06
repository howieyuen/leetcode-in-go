package _287_find_the_duplicate_number

func findDuplicate(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if i == nums[i] {
			continue
		}
		for nums[i] != i {
			if nums[nums[i]] != nums[i] {
				nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
			} else {
				return nums[i]
			}
		}
	}
	return nums[0]
}

/*
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

相当于带环的链表，找到环的入口
*/
func findDuplicate1(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
