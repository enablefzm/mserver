package ukclientapps

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/enablefzm/mserver/dbs"
	ukclientdtos "github.com/enablefzm/mserver/dtos/ukclientDtos"
	"github.com/enablefzm/mserver/model"
	"github.com/enablefzm/mserver/model/ukclient"

	"gorm.io/gorm"
)

func init() {
	PAppUkClient.load()
}

var PAppUkClient = &AppUkClient{
	mpClients: make(map[string]*UkClientInfo, 10),
	lkmp:      new(sync.RWMutex),
}

type AppUkClient struct {
	mpClients map[string]*UkClientInfo
	lkmp      *sync.RWMutex
}

func (uk *AppUkClient) load() {
	// 加载数据
	var clients []ukclient.UkClient
	tx := dbs.ObDB.Find(clients)
	if tx.Error != nil {
		log.Println("加载数据库里存有的客户端信息出错")
		return
	}
	uk.lkmp.Lock()
	defer uk.lkmp.Unlock()
	for i := 0; i < len(clients); i++ {
		ob := clients[i]
		uk.mpClients[ob.Tax] = NewUkClientInfo(&ob)
	}
}

func (uk *AppUkClient) CheckUpdate(ipAddress, version string, pCompanyInfo *ukclientdtos.BaseCompanyInfoDto, modClient *ukclient.UkClient) {
	// 判断哪些数据需要更新
	// TODO 判断哪些数据需要去更新...
}

// 获取组件客户端的信息
func (uk *AppUkClient) BaseInfo(input *ukclientdtos.BaseInfoDto) error {
	uk.lkmp.Lock()
	defer uk.lkmp.Unlock()
	// 查询本地是否有这个对象信息
	for i := 0; i < len(input.CompanyInfos); i++ {
		ob := input.CompanyInfos[i]
		p, ok := uk.mpClients[ob.Tax]
		if ok {
			uk.CheckUpdate(input.IpAddress, input.Version, &ob, &p.UkClient)
		} else {
			pClient, err := ukclient.NewUkClientOnTax(ob.Tax)
			if err != nil {
				// 数据库里没有
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// 新增对象
					pNewUkClient := &ukclient.UkClient{
						SoftDelModel: model.NewSoftDelModelOnCreateGuid(),
						Name:         ob.Name,
						Tax:          ob.Tax,
						IpAddress:    input.IpAddress,
						Version:      input.Version,
					}
					tx := dbs.ObDB.Create(pNewUkClient)
					if tx.Error != nil {
						log.Println(ob.Tax, "同步客户端信息时,新增客户端信息到数据库里发生错误:", err.Error())
						continue
					}
					// 生成一个新的对内存对象信息
					p = NewUkClientInfo(pNewUkClient)
				} else {
					log.Println(ob.Tax, "同步客户端信息时,内存里没有这个对在和数据库里检索该对象信息时出错:", err.Error())
					continue
				}
			} else {
				// 有这个对象则开始重新加载到当前对象
				p = NewUkClientInfo(pClient)
			}
			// 放入内存的map里
			uk.mpClients[ob.Tax] = p
		}
		p.lastTime = time.Now()
	}
	return nil
}

// 收到客户端心跳包检测
func (uk *AppUkClient) KeepAlive(input *ukclientdtos.BaseInfoDto) error {
	return nil
}

func NewUkClientInfo(obUkClient *ukclient.UkClient) *UkClientInfo {
	return &UkClientInfo{
		UkClient: *obUkClient,
		lastTime: time.Now().Add(1 * time.Minute),
	}
}

type UkClientInfo struct {
	ukclient.UkClient
	lastTime time.Time
}
