package ydna

// YDNA 描述基因样本的信息
type YDNA struct {
	Title       string
	Description string
	Type        string
	// 地点-经纬度
	Location [2]float64
	// 地点-名称
	Address string

	// 智能分组, 智能分类, 智能文件夹
	Groups []string

	// 国家
	Country string
}

//
