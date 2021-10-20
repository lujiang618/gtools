package utils

// 获取分页的offset
func GetPageOffset(pageSize uint, page uint) uint {
	// 默认从0开始
	if page < 1 {
		return 0
	}

	return (page - 1) * pageSize
}
