package utils

import (
	"math/rand"
	"time"
)

// math/rand包中的函数默认使用了一个全局锁，在并发加高频率调用的情况下，全局锁损耗非常大。
// 解决方法：对每个并发的goroutine，单独生成一个rand
func generator() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 生成指定范围的随机数
func RandomInt(min, max int) int {
	if max < min {
		return 0
	}

	max = max - min + 1
	return generator().Intn(max) + min
}

//生成随机字符串
func RandomString(size int) string {
	if size < 1 {
		return ""
	}

	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, 0, size)
	for i := 0; i < size; i++ {
		result = append(result, str[generator().Intn(len(str))])
	}

	return string(result)
}
