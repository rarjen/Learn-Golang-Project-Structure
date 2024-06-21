package datasource

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"template-ulamm-backend-go/utils"

	"go.uber.org/zap"
)

// temporarily open the connection and close the connecton
func NewSQLConnection(sqlStringConn string) error {
	db, err := sql.Open("sqlserver", sqlStringConn)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func ChangeSQLConnection(connString string) {
	var err error
	// if current conection is the second DB, try to change to first DB
	if !dataSource.isUseSqlServerPrimaryDb {
		err = NewSQLConnection(utils.GetConfig().SqlServer.Hosts[0])
		if err != nil {
			utils.GetLogger().Error("failed when create a connection to primary DB", zap.Error(err))
			return
		}

		// create new connection to primary DB
		if err = NewGormDb(utils.GetConfig().SqlServer.Hosts[0]); err != nil {
			// this panic wouldn't be invoked (normally)
			log.Panic(err)
		}

		dataSource.isUseSqlServerPrimaryDb = true
		utils.GetLogger().Info(fmt.Sprintf("connection to %s DB sucessfully established\n", strings.Split(connString, ";")[0]))
		return
	}

	if connString == utils.GetConfig().SqlServer.Hosts[0] {
		if err = dataSource.PingDB(); err != nil {
			// if failed, recursive to change the con
			utils.GetLogger().Error("error when ping database on cron job", zap.Error(err))
			ChangeSQLConnection(utils.GetConfig().SqlServer.Hosts[1])
			return
		}

		dataSource.isUseSqlServerPrimaryDb = true
	} else {
		err = NewGormDb(connString)
		if err != nil {
			// change to the second DB
			log.Fatalln("failed to connect to DB Warehouse")
		}

		dataSource.isUseSqlServerPrimaryDb = false
		utils.GetLogger().Info(fmt.Sprintf("connection to %s DB sucessfully established\n", strings.Split(connString, ";")[0]))
	}
}
