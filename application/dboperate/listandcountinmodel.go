package dboperate

import (
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/commdtos"
)

// 用Struct来查询，这种方法如果查询条件是空值的话则不会构造查询
// 		dest   用来接收要被返回的查询对象
//		pModel 查询的模型对象
func ListAndCountInModel(
	commInput commdtos.CommInputGet,
	dest interface{},
	pModel interface{},
	order string,
	whereKey string,
	searchs ...interface{}) (*commdtos.CommResultPage, error) {

	strFilter := commInput.GetFilter()
	var count int64
	qry := dbs.ObDB.Scopes(Pginate(commInput))
	count = CountFilterOnModel(pModel, strFilter, whereKey, searchs...)
	if len(whereKey) > 0 && len(strFilter) > 0 {
		qry.Where(whereKey, searchs...)
	}
	if len(order) > 0 {
		qry.Order(order)
	}
	result := qry.Find(&dest)
	return commdtos.NewCommResultPage(count, commInput.Page, dest), result.Error
}
