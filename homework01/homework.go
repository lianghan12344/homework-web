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
	// 定义括号匹配的映射：key为右括号，value为对应的左括号
	pair := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	// 用切片模拟栈，存储左括号
	stack := []rune{}

	// 遍历字符串的每个字符
	for _, char := range s {
		// 情况1：左括号，入栈
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			// 情况2：右括号，检查栈是否为空（无左括号匹配）
			if len(stack) == 0 {
				return false
			}
			// 取出栈顶元素，判断是否匹配
			top := stack[len(stack)-1]
			if top != pair[char] {
				return false
			}
			// 匹配成功，弹出栈顶
			stack = stack[:len(stack)-1]
		}
	}

	// 遍历结束后，栈必须为空（所有左括号都匹配闭合）
	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 遍历第一个字符串的每个字符位置
	for i := 0; i < len(strs[0]); i++ {
		// 取当前位置的字符作为基准
		char := strs[0][i]
		// 对比其他字符串的同一位置字符
		for j := 1; j < len(strs); j++ {
			// 若某字符串长度不足，或字符不匹配，返回前i个字符
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i]
			}
		}
	}

	// 第一个字符串完全匹配所有字符串
	return strs[0]
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	// 从最后一位（个位）开始遍历
	for i := len(digits) - 1; i >= 0; i-- {
		// 当前位加一
		digits[i]++
		// 取模10：若≤9则无进位，直接退出循环；若=10则变为0（进位）
		digits[i] %= 10
		// 无进位，直接返回
		if digits[i] != 0 {
			return digits
		}
	}

	// 执行到这里说明所有位都是9（如[9,9]），需要在头部插入1
	newDigits := make([]int, len(digits)+1)
	newDigits[0] = 1
	// 其余位默认是0，无需修改
	return newDigits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// 边界条件：空数组直接返回0
	if len(nums) == 0 {
		return 0
	}

	// 慢指针初始化为0，指向去重后数组的最后一个元素
	slow := 0

	// 快指针遍历数组（从1开始，因为第一个元素必然保留）
	for fast := 1; fast < len(nums); fast++ {
		// 找到与慢指针位置不同的元素（新的唯一元素）
		if nums[fast] != nums[slow] {
			slow++ // 慢指针后移，准备存储新元素
			nums[slow] = nums[fast] // 覆盖重复位置，完成原地修改
		}
		// 若元素重复，快指针继续前进，慢指针不动
	}

	// 新长度为慢指针+1（索引从0开始）
	return slow + 1
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	// 边界条件：空数组或单区间直接返回
	if len(intervals) <= 1 {
		return intervals
	}

	// 步骤1：手动实现冒泡排序，按区间起始值升序排列（替换sort包）
	n := len(intervals)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			// 比较相邻区间的起始值，前大后小则交换
			if intervals[j][0] > intervals[j+1][0] {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
				swapped = true
			}
		}
		// 无交换说明已排序完成，提前退出
		if !swapped {
			break
		}
	}

	// 步骤2：合并重叠区间（逻辑与原解法一致）
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		curr := intervals[i]

		// 重叠则合并，否则加入结果
		if curr[0] <= last[1] {
			last[1] = max(last[1], curr[1])
		} else {
			result = append(result, curr)
		}
	}

	return result
}

// 辅助函数：取两个整数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	// 定义map：key为数组值，value为对应下标
	numMap := make(map[int]int)

	// 遍历数组，同时记录已遍历的数
	for idx, num := range nums {
		// 计算需要匹配的补数
		complement := target - num
		// 检查补数是否已在map中（且不是当前下标）
		if complementIdx, ok := numMap[complement]; ok {
			// 找到结果，返回两个下标
			return []int{complementIdx, idx}
		}
		// 补数不存在，将当前数和下标存入map
		numMap[num] = idx
	}

	// 题目保证有解，此处仅为语法兜底
	return nil
}