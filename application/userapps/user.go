package userapps

import (
	"errors"
	"fmt"
	"strings"

	"github.com/enablefzm/mserver/application/dboperate"
	"github.com/enablefzm/mserver/dbs"
	"github.com/enablefzm/mserver/dtos/commdtos"
	"github.com/enablefzm/mserver/dtos/userdtos"
	"github.com/enablefzm/mserver/model"
	"github.com/enablefzm/mserver/model/user"

	"github.com/enablefzm/gotools/guid"
	"github.com/enablefzm/gotools/vatools"
	"gorm.io/gorm"
)

var AppUser = &User{}

type User struct {
}

// 创建一个新用户
func (u *User) Create(input userdtos.CreateUpdateUserDto) error {
	// 判断是否有同名的uid
	if err := dboperate.CheckRepeat(input.GetUid(), &user.User{UID: input.GetUid()}); err != nil {
		return err
	}
	newPassword := input.GetPassword()
	if len(newPassword) < 1 {
		newPassword = "123456"
	}
	obUser := &user.User{
		BaseModel:       model.NewModelBase(),
		UID:             input.GetUid(),
		Name:            input.Name,
		PassWord:        u.OperatePassword(newPassword),
		IdentityCard:    input.IdentityCard,
		PhoneNumber:     input.PhoneNumber,
		TelephoneNumber: input.TelephoneNumber,
		IsOnJob:         input.IsOnJob,
		Sex:             input.Sex,
	}
	obUser.Guid = guid.NewString()
	tx := dbs.ObDB.Create(obUser)
	return tx.Error
}

// 修改用户信息
func (u *User) Edit(guid string, input userdtos.CreateUpdateUserDto) error {
	// 获取对象
	var obUser user.User
	tx := dboperate.GetGuid(guid, &obUser)
	if tx.Error != nil {
		return tx.Error
	}
	// 判断是否需要更新密码
	mpUpdate := map[string]interface{}{
		"uid":              input.UID,
		"name":             input.Name,
		"sex":              input.Sex,
		"identity_card":    input.IdentityCard,
		"phone_number":     input.PhoneNumber,
		"telephone_number": input.TelephoneNumber,
		"is_on_job":        input.IsOnJob,
	}
	newPassword := input.GetPassword()
	if len(newPassword) > 0 {
		mpUpdate["pass_word"] = u.OperatePassword(newPassword)
	}
	// 更新数据
	tx = dbs.ObDB.Model(&obUser).Updates(mpUpdate)
	return tx.Error
}

// 通过UID检索数据
func (u *User) GetOnUid(uid string) (*user.User, error) {
	obUser := &user.User{}
	result := dbs.ObDB.First(obUser, "uid = ?", strings.ToLower(uid))
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return obUser, errors.New("用户名不存在")
	}
	return obUser, result.Error
}

// 登入并判断用是否在存
func (u *User) Login(uid, password string) (*user.User, error) {
	pUser, err := u.GetOnUid(uid)
	if err != nil {
		return nil, err
	}
	// 判断密码是否相符
	if pUser.PassWord != u.OperatePassword(password) {
		return nil, fmt.Errorf("用户密码错误")
	}
	// 判断用户是否是已经离职人员
	if !pUser.IsOnJob {
		return nil, fmt.Errorf("这名用户已从海天离职")
	}
	return pUser, nil
}

// 通过Guid来获取用户信息
func (u *User) InfoOnGuid(id string) (userdtos.UserInfoDto, error) {
	obUser := &user.User{}
	tx := dbs.ObDB.First(obUser, "guid = ?", id)
	if tx.Error != nil {
		return userdtos.UserInfoDto{}, tx.Error
	}
	return u.Convert(obUser), nil
}

// 列表方式获取用户信息
func (u *User) ListInfo(commInput commdtos.CommInputGet) (*commdtos.CommResultPage, error) {
	var userInfos []*user.User
	// 获取总数据
	var count int64
	qry := dbs.ObDB.Scopes(dboperate.Pginate(commInput)).Not(user.User{UID: "admin"})
	qryCount := dbs.ObDB.Model(user.User{}).Not(user.User{UID: "admin"})
	// 判断是否是要查询在职人员或着不在职人员
	var mpQryIsOnJob = make(map[string]interface{}, 1)
	switch commInput.Params {
	// 查询离职人员
	case "0":
		mpQryIsOnJob["is_on_job"] = false
	// 查询在职人员
	case "1":
		mpQryIsOnJob["is_on_job"] = true
	}
	if len(mpQryIsOnJob) > 0 {
		qry.Where(mpQryIsOnJob)
		qryCount.Where(mpQryIsOnJob)
	}
	strFilter := commInput.GetFilter()
	count = dboperate.CountFilter(
		qryCount,
		strFilter,
		"name LIKE ? OR phone_number LIKE ? OR  uid LIKE ?", strFilter, strFilter, strFilter)
	if len(strFilter) > 0 {
		qry.Where("name LIKE ? OR phone_number LIKE ? OR  uid LIKE ?", strFilter, strFilter, strFilter)
	}
	result := qry.Order("uid asc").Find(&userInfos)

	// 将对象转换为用户列表dto数据
	userInfosDto := make([]userdtos.UserInfoDto, 0, len(userInfos))
	for _, ob := range userInfos {
		userInfosDto = append(userInfosDto, u.Convert(ob))
	}
	return commdtos.NewCommResultPage(count, commInput.Page, userInfosDto), result.Error
}

// 用户对象数据转换为Dto传输数据
func (u *User) Convert(obUser *user.User) userdtos.UserInfoDto {
	return userdtos.UserInfoDto{
		GUID:            obUser.Guid,
		UID:             obUser.UID,
		Name:            obUser.Name,
		Sex:             obUser.Sex,
		IdentityCard:    obUser.IdentityCard,
		PhoneNumber:     obUser.PhoneNumber,
		TelephoneNumber: obUser.TelephoneNumber,
		IsOnJob:         obUser.IsOnJob,
	}
}

func (u *User) OperatePassword(sourcePasswd string) string {
	return vatools.MD5(sourcePasswd + "_jimmyFanZhiMing2200")
}
