package testutil

import (
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"

	"git.paylabo.com/c002/harp/backend-purchase/common/constant"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

const (
	dbDriver               = "mysql"
	dbConnectRetryMaxCount = 10
	retryWaitTime          = 3000 // ミリ秒
)

func init() {

	db, err := ConnectDatabaseWithGorm()
	if err != nil {
		panic(err)
	}
	if err := loadSchema(db); err != nil {
		panic(err)
	}
	if err := initData(db); err != nil {
		panic(err)
	}

}

func ConnectDatabaseWithGorm() (*gorm.DB, error) {
	loadEnv()
	dbHost := viper.GetString(constant.DbHostEnv)
	dbPort := viper.GetString(constant.DbPortEnv)
	dbUser := viper.GetString(constant.DbUserEnv)
	dbName := viper.GetString(constant.DbNameEnv)
	dbPassword := viper.GetString(constant.DbPasswordEnv)
	dbLogMode, _ := strconv.ParseBool(viper.GetString(constant.DbLogModeEnv))

	dbUserAndPassword := dbUser
	if len(dbPassword) > 0 {
		dbUserAndPassword += ":" + dbPassword
	}
	dbURL := fmt.Sprintf("%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", dbUserAndPassword, dbHost, dbPort, dbName)
	var db *gorm.DB
	var err error
	for i := 1; i <= dbConnectRetryMaxCount; i++ {
		db, err = gorm.Open(dbDriver, dbURL)
		if err == nil {
			break
		}
		time.Sleep(retryWaitTime * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}
	db = db.LogMode(dbLogMode)
	return db, nil
}

func TruncateTable(db *gorm.DB, table string) error {
	_, err := db.DB().Exec("TRUNCATE " + table)
	if err != nil {
		return err
	}
	return nil
}

func initData(db *gorm.DB) error {
	tables := []string{
		"acceptance_details",
		"contract_manager_details",
		"harp_sys_sequences",
		"item_units",
		"order_details",
		"order_items",
		"order_states",
		"order_types",
		"orders",
		"project_cost_details",
		"quotation_details",
		"quotation_histories",
		"quotation_items",
		"quotations",
		"order_histories",
	}
	for _, table := range tables {
		// _, err := db.DB().Exec("TRUNCATE " + table)
		// if err != nil {
		// 	return err
		// }
		_, p, _, _ := runtime.Caller(0)
		initDataFile := path.Join(path.Dir(p), "init_data", table+".sql")
		if content, err := ioutil.ReadFile(initDataFile); err == nil {
			ddl := string(content)
			for _, q := range strings.Split(ddl, ";") {
				if len(q) == 0 {
					continue
				}
				if _, err := db.DB().Exec(q); err != nil {
					return err
				}

			}
		}
	}
	return nil
}

func loadSchema(db *gorm.DB) error {
	_, filename, _, _ := runtime.Caller(0)
	schema := path.Join(path.Dir(filename), "schema.sql")

	content, err := ioutil.ReadFile(schema)
	if err != nil {
		return err
	}
	ddl := string(content)
	for _, q := range strings.Split(ddl, ";") {
		if _, err := db.DB().Exec(q); err != nil {
			return err
		}
	}

	return nil
}

func loadEnv() {
	viper.AutomaticEnv()
	viper.SetDefault(constant.DbHostEnv, "127.0.0.1")
	viper.SetDefault(constant.DbPortEnv, "3306")
	viper.SetDefault(constant.DbUserEnv, "root")
	viper.SetDefault(constant.DbNameEnv, "harp_test")
	viper.SetDefault(constant.DbPasswordEnv, "")
	// export DB_LOG_MODE=true then run test case.
	viper.SetDefault(constant.DbLogModeEnv, false)
}
