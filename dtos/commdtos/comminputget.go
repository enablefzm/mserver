package commdtos

import (
	"fmt"
	"strings"
)

type CommInputGet struct {
	Params   string `json:"params"`
	PageSize int    `json:"page_size"`
	Page     int    `json:"page"`
	Filter   string `json:"filter"`
}

// 获取公共查询Dto里的查询关键字,如果这个Filter关健字不为空,则返回 %Filter% 字符
func (c *CommInputGet) GetFilter() string {
	s := strings.Trim(c.Filter, " ")
	if len(s) > 0 {
		return fmt.Sprint("%", s, "%")
	}
	return s
}
