package bitManipulation

func hammingWeight(num uint32) int {
	//求无符号int32的位数
	//return bits.OnesCount(uint(num))

	//使用 x & (x - 1) 清除最低位的
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}