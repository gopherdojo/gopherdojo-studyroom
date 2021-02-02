package utilities

// 配列の一部を別配列に置き換える
func FillByteArr(arr []byte, startAt int, partArr []byte) {
	for i := 0; i < len(partArr); i++ {
		globalIndex := i + startAt
		arr[globalIndex] = partArr[i]
	}
}
