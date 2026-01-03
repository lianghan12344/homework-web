package homework01


// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	// 定义map，key为数组元素，value为出现次数
	countMap := make(map[int]int)

	// 第一次遍历：统计每个元素的出现次数
	for _, num := range nums {
		countMap[num]++ // 元素出现一次，计数+1
	}

	// 第二次遍历：找到出现次数为1的元素
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}

	// 题目保证非空且必有唯一元素，此处仅为语法兜底
	return -1
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(num int) bool {
// 边界条件：负数/末尾为0且非0的数，直接返回false
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}

	reversedHalf := 0 // 存储反转的后半部分数字
	// 反转后半部分，直到原数 <= 反转后的数（说明已处理一半）
	for num > reversedHalf {
		// 取原数最后一位，拼接到反转数末尾
		reversedHalf = reversedHalf*10 + num%10
		num = num / 10 // 原数去掉最后一位
	}

	// 偶数位：num == reversedHalf；奇数位：num == reversedHalf/10（去掉中间位）
	return num == reversedHalf || num == reversedHalf/10
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// TODO: implement
	return false
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	// TODO: implement
	return ""
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// TODO: implement
	return nil
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// TODO: implement
	return 0
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// TODO: implement
	return nil
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// TODO: implement
	return nil
}