package _164_maximum_gap

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 找到桶的范围
	var maxNum, minNum = nums[0], nums[0]
	for i := range nums {
		maxNum = max(maxNum, nums[i])
		minNum = min(minNum, nums[i])
	}
	
	// 桶的大小，n个元素有n-1个间距，间距大小即为桶大小，故用最大值和最小值的差，除以间距个数，得到平均间距大小
	bucketSize := max(1, (maxNum-minNum)/(len(nums)-1))
	// 桶个数
	bucketNum := (maxNum-minNum)/bucketSize + 1
	// 桶排序
	var buckets = make([]*Bucket, bucketNum)
	for i := range nums {
		index := (nums[i] - minNum) / bucketSize
		if buckets[index] == nil {
			buckets[index] = &Bucket{
				max: nums[i],
				min: nums[i],
			}
		} else {
			buckets[index].max = max(buckets[index].max, nums[i])
			buckets[index].min = min(buckets[index].min, nums[i])
		}
	}
	
	var preBucketMax = minNum
	// 间隔 = 当前桶的最小值 - 前个桶中最大值
	var maxGap = 0
	for _, bucket := range buckets {
		if bucket == nil {
			continue
		}
		maxGap = max(maxGap, bucket.min-preBucketMax)
		preBucketMax = bucket.max
	}
	return maxGap
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type Bucket struct {
	max int
	min int
}
