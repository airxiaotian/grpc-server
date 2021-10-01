//go:generate ./protoc.sh
//go:generate ./mockgen.sh
package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"git.paylabo.com/c002/harp/backend-purchase/common/constant"
	"git.paylabo.com/c002/harp/backend-purchase/common/logging"
	"git.paylabo.com/c002/harp/backend-purchase/infra/repository"
	"git.paylabo.com/c002/harp/backend-purchase/interfaces"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	logger *zap.SugaredLogger
)

const (
	bindAddr               = "0.0.0.0"
	listeningPort          = "50051"
	network                = "tcp"
	dbDriver               = "mysql"
	dbConnectRetryMaxCount = 10
	retryWaitTime          = 3000 // ミリ秒
)

func main() {
	loadEnv()
	logging.Configure(getLoggingConfig())
	logger = logging.NewLogger()

	if err := validateEnv(); err != nil {
		logger.Fatalw("failed to validate env", "domain_error", err)
	}

	db, err := connectDatabaseWithGorm()
	if err != nil {
		logger.Fatalw("failed to connect database", "domain_error", err)
	}
	defer db.Close()

	server := interfaces.NewServer(interfaces.ServerParams{
		AcceptanceDetailsRepository:     repository.NewAcceptanceDetailsRepository(db),
		OrderTypeRepository:             repository.NewOrderTypeRepository(db),
		QuotationRepository:             repository.NewQuotationRepository(db),
		OrderRepository:                 repository.NewOrderRepository(db),
		OrderDetailsRepository:          repository.NewOrderDetailsRepository(db),
		OrderItemsRepository:            repository.NewOrderItemsRepository(db),
		OrderStatesRepository:           repository.NewOrderStatesRepository(db),
		OrderHistoryRepository:          repository.NewOrderHistoryRepository(db),
		ItemUnitsRepository:             repository.NewItemUnitsRepository(db),
		QuotationsRepository:            repository.NewQuotationRepository(db),
		QuotationItemsRepository:        repository.NewQuotationItemsRepository(db),
		QuotationDetailRepository:       repository.NewQuotationDetailRepository(db),
		QuotationHistoryRepository:      repository.NewQuotationHistoryRepository(db),
		ContractManagerDetailRepository: repository.NewContractManagerDetailRepository(db),
		HarpSysSequenceRepository:       repository.NewHarpSysSequenceRepository(db),
		ProjectCostDetailRepository:     repository.NewProjectCostDetailRepository(db),
	})

	var stop = make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)

	go func() {
		sig := <-stop
		logger.Infof("shutdown signal is called, signal=[%v], shutting down...", sig)
		server.GracefulStop()
	}()

	listener, err := net.Listen(network, fmt.Sprintf("%v:%v", bindAddr, listeningPort))
	if err != nil {
		logger.Fatalw("failed to listen server", "domain_error", err)
	}
	logger.Infof("server listen start, addr=[%v:%v]", bindAddr, listeningPort)
	if err = server.Serve(listener); err != nil {
		logger.Fatalw("unknown domain_error occurred, when shutting down", "domain_error", err)
		return
	}

	logger.Infof("shutdown success, bye!")
}

func connectDatabaseWithGorm() (*gorm.DB, error) {
	dbHost := viper.GetString(constant.DbHostEnv)
	dbPort := viper.GetString(constant.DbPortEnv)
	dbName := viper.GetString(constant.DbNameEnv)
	dbUser := viper.GetString(constant.DbUserEnv)
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
		logger.Warnw("db access error retry...", "error", err)
		time.Sleep(retryWaitTime * time.Millisecond)
	}
	if err != nil {
		return nil, err
	}

	idle := viper.GetInt(constant.DbIdleConnectionEnv)
	max := viper.GetInt(constant.DbMaxConnectionEnv)
	lifeTime := viper.GetInt(constant.DbConnMaxLifeTimeEnv)

	db.DB().SetMaxIdleConns(idle)
	db.DB().SetMaxOpenConns(max)

	db.DB().SetConnMaxLifetime(time.Duration(lifeTime) * time.Second)

	db = db.LogMode(dbLogMode)
	return db, nil
}

func loadEnv() {

	viper.AutomaticEnv()
	viper.SetDefault(constant.DbHostEnv, "127.0.0.1")
	viper.SetDefault(constant.DbPortEnv, "3306")
	viper.SetDefault(constant.DbUserEnv, "root")
	viper.SetDefault(constant.DbNameEnv, "harp")
	viper.SetDefault(constant.DbPasswordEnv, "")
	viper.SetDefault(constant.DbIdleConnectionEnv, 20)
	viper.SetDefault(constant.DbMaxConnectionEnv, 40)
	viper.SetDefault(constant.DbConnMaxLifeTimeEnv, 40)
	viper.SetDefault(constant.AppRevisionEnv, "NotSetting")
	viper.SetDefault(constant.AppVersionEnv, "NotSetting")
	viper.SetDefault(constant.DbLogModeEnv, true)
}

func validateEnv() error {
	e := make([]string, 0)

	if dbPort, err := strconv.Atoi(viper.GetString(constant.DbPortEnv)); err != nil || dbPort < 0 || dbPort > 65535 {
		e = append(e, fmt.Sprintf("%v=%v", constant.DbPortEnv, viper.GetString(constant.DbPortEnv)))
	}

	if idle := viper.GetInt(constant.DbIdleConnectionEnv); idle < 0 {
		e = append(e, fmt.Sprintf("%v=%v", constant.DbIdleConnectionEnv, idle))
	}

	if max := viper.GetInt(constant.DbMaxConnectionEnv); max < 0 {
		e = append(e, fmt.Sprintf("%v=%v", constant.DbMaxConnectionEnv, max))
	}

	if maxLife := viper.GetInt(constant.DbConnMaxLifeTimeEnv); maxLife < 0 {
		e = append(e, fmt.Sprintf("%v=%v", constant.DbConnMaxLifeTimeEnv, maxLife))
	}

	if len(e) != 0 {
		return errors.New(strings.Join(e, ", "))
	}
	return nil
}

func getLoggingConfig() logging.Config {
	return logging.Config{
		LogLevel:    viper.GetString(constant.LogLevelEnv),
		AppVersion:  viper.GetString(constant.AppVersionEnv),
		AppRevision: viper.GetString(constant.AppRevisionEnv),
	}
}
