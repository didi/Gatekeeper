package testsqlhandler

import (
	"fmt"
	"github.com/didi/gatekeeper/golang_common/lib"
	"github.com/didi/gatekeeper/golang_common/zerolog/log"
	"github.com/didi/gatekeeper/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/e421083458/gorm"
)

var Db *gorm.DB

func InitGORMHandler() {
	DbConfMap := &lib.MysqlMapConf{}
	fmt.Println("TEST", lib.GetConfPath("mysql_map"))
	err := lib.ParseConfig(lib.GetConfPath("mysql_map"), DbConfMap)
	if err != nil {
		return
	}
	fmt.Println("TEST", DbConfMap)
	mysqlConf := DbConfMap.List["default"]
	dsn := mysqlConf.DataSourceName
	fmt.Println("TEST", dsn)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//tmp, err := lib.GetGormPool("default")
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	//Db = tmp
	log.Info().Msg("InitConfig GORMHandler Success.")
}

func UpdateServiceStripPrefix(serviceName string, stripPrefix int) {
	Db.Model(&model.ServiceInfo{}).Where("service_name = ?", serviceName).Update("http_strip_prefix", stripPrefix)
}

func GetServiceStripPrefix(serviceName string) int {
	task := model.ServiceInfo{}
	Db.Where("service_name = ?", serviceName).First(&task)
	return task.HttpStripPrefix
}

func GetServiceLoadBalanceStrategy(serviceName string) string {
	task := model.ServiceInfo{}
	Db.Where("service_name = ?", serviceName).First(&task)
	return task.LoadBalanceStrategy
}

func DeleteServiceInfo(serviceName string) {
	Db.Where("service_name = ?", serviceName).Delete(model.ServiceInfo{})
}

func AddServiceInfo(serviceInfo *model.ServiceInfo) {
	Db.Create(serviceInfo)
}

//func Save(serviceInfo *model.ServiceInfo) {
//	tmp, err := lib.GetGormPool("default")
//	if err != nil {
//		panic("连接数据库失败, error=" + err.Error())
//	}
//
//	err = tmp.Save(serviceInfo).Error
//	if err != nil {
//		fmt.Println("SAVE:", err)
//	}
//}
