package bitManipulation



func reverseBits(num uint32) uint32 {
	//  num向右移动，消去最右边低位的1，将1给res

	//var res uint32
	//for i := 0; i < 32; i++ {
	//	res = res << 1 | num & 1
	//	res >>= 1
	//}
	//
	//return res

	//注意前缀0,
	var res uint32
	for i := 0; i < 32; i++ {
		res = res * 2 + num % 2
		num = num / 2
	}
	return res

}
