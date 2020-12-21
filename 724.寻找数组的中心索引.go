/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 *
 * https://leetcode-cn.com/problems/find-pivot-index/description/
 *
 * algorithms
 * Easy (36.26%)
 * Likes:    239
 * Dislikes: 0
 * Total Accepted:    71K
 * Total Submissions: 180K
 * Testcase Example:  '[1,7,3,6,5,6]'
 *
 * 给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。
 * 
 * 我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
 * 
 * 如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：
 * nums = [1, 7, 3, 6, 5, 6]
 * 输出：3
 * 解释：
 * 索引 3 (nums[3] = 6) 的左侧数之和 (1 + 7 + 3 = 11)，与右侧数之和 (5 + 6 = 11) 相等。
 * 同时, 3 也是第一个符合要求的中心索引。
 * 
 * 
 * 示例 2：
 * 
 * 输入：
 * nums = [1, 2, 3]
 * 输出：-1
 * 解释：
 * 数组中不存在满足此条件的中心索引。
 * 
 * 
 * 
 * 说明：
 * 
 * 
 * nums 的长度范围为 [0, 10000]。
 * 任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。
 * 
 * 
 */

// @lc code=start
func pivotIndex(nums []int) int {
	 // 总的一句, 就是左边和 + 索引值 + 右边和 == 数组和
	RSum, LSum := 0, 0

	for i := range nums {
		RSum += nums[i]
	}

	// 速度始终无法100%
	// 左边和 * 2 + 索引值 == 数组和, 经过两次提价, 发现速度都不太理想, 所以改了, 好点, 没全好
	for i := range nums {
		if RSum -= nums[i] ; RSum == LSum {
			return i
		}

		LSum += nums[i]
	}

	return -1
}
// @lc code=end

