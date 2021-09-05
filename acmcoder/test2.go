package main

import "fmt"

func main() {
	fmt.Print(byte(127))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param bills int整型一维数组 顾客队列
 * @return int整型
 */
func billsChange( bills []int ) int {
	// write code here
	res := 0
	five, ten := 2, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
			res++
		} else if bill == 10 {
			if five == 0 {
				return res
			}
			five--
			ten++
			res++
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
				res++
			} else if five >= 3 {
				res++
			} else {
				return res
			}
		}
	}
	return res



}


/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param e int整型一维数组 充电桩最大充电量
 * @param c int整型一维数组 达到下一个充电桩需消耗的电量
 * @return int整型
 */

// 1,2,3,4,5   3,4,5,1,2
//第三号位置开始

//func canCompleteRace( e []int ,  c []int ) int {
//	// write code here
//
//	var start int
//	for i := 0; i < len(c); i++ {
//		if e[i] > c[i] {
//			//从i开始遍历
//			start = i
//		}
//
//		energy := e[start]
//		end := 0
//		for {
//			if
//		}
//
//
//	}
//
//
//
//
//
//}



/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 输入多个城市之间的距离，返回到其他城市最近的城市的标号，如果无法从这几个城市之间选出满足条件的城市，那么返回-1。
 * @param distancePairs int整型二维数组 二维数组，每个子数组都是三个整数组成的，前两个整数代表城市的标号，第三个整数代表两个城市的距离
 * @param CityNum int整型 城市的数量，每个城市的标号分别是0~CityNum-1
 * @return int整型
 */

// [0, 1, 3] [1, 2, 2]

// [0, 1, ] ...  [0, 2, 1]

// 选择一个离其他城市最近的仓库

//func GetBestWarehouseLocate( distancePairs [][]int ,  CityNum int ) int {
//	// write code here
//}





/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 该方法接收一个二叉树root和一条输入路径path，返回一个数组，数组包含二叉树中的最优质路径，每个数组元素的值代表二叉树中节点的val
 * @param root TreeNode类 表示一颗二叉树的根节点root，不会为空
 * @param path int整型一维数组 输入路径path，不会为空
 * @return int整型一维数组
 */
//func solution( root *TreeNode ,  path []int ) []int {
//	// write code here
//
//
//}

