package mysql

import (
	"context"
	"database/sql"
	jetmysql "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/tests/dbconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"runtime"

	"github.com/pkg/profile"
	"os"
	"testing"
)

var db *jetmysql.DB

var source string

const MariaDB = "MariaDB"

func init() {
	source = os.Getenv("MY_SQL_SOURCE")
}

func sourceIsMariaDB() bool {
	return source == MariaDB
}

func TestMain(m *testing.M) {
	defer profile.Start().Stop()

	var err error
	sqlDB, err := sql.Open("mysql", dbconfig.MySQLConnectionString(sourceIsMariaDB(), ""))
	if err != nil {
		panic("Failed to connect to test db" + err.Error())
	}

	db = jetmysql.NewDB(sqlDB).WithStatementsCaching(true)
	defer db.Close()

	for i := 0; i < 2; i++ {
		ret := m.Run()
		if ret != 0 {
			os.Exit(ret)
		}
	}
}

var loggedSQL string
var loggedSQLArgs []interface{}
var loggedDebugSQL string

var queryInfo jetmysql.QueryInfo
var callerFile string
var callerLine int
var callerFunction string

func init() {
	jetmysql.SetLogger(func(ctx context.Context, statement jetmysql.PrintableStatement) {
		loggedSQL, loggedSQLArgs = statement.Sql()
		loggedDebugSQL = statement.DebugSql()
	})

	jetmysql.SetQueryLogger(func(ctx context.Context, info jetmysql.QueryInfo) {
		queryInfo = info
		callerFile, callerLine, callerFunction = info.Caller()
	})
}

func requireLogged(t *testing.T, statement postgres.Statement) {
	query, args := statement.Sql()
	require.Equal(t, loggedSQL, query)
	require.Equal(t, loggedSQLArgs, args)
	require.Equal(t, loggedDebugSQL, statement.DebugSql())
}

func requireQueryLogged(t *testing.T, statement postgres.Statement, rowsProcessed int64) {
	query, args := statement.Sql()
	queryLogged, argsLogged := queryInfo.Statement.Sql()

	require.Equal(t, query, queryLogged)
	require.Equal(t, args, argsLogged)
	require.Equal(t, queryInfo.RowsProcessed, rowsProcessed)

	pc, file, _, _ := runtime.Caller(1)
	funcDetails := runtime.FuncForPC(pc)
	require.Equal(t, file, callerFile)
	require.NotEmpty(t, callerLine)
	require.Equal(t, funcDetails.Name(), callerFunction)
}

func skipForMariaDB(t *testing.T) {
	if sourceIsMariaDB() {
		t.SkipNow()
	}
}

func onlyMariaDB(t *testing.T) {
	if !sourceIsMariaDB() {
		t.SkipNow()
	}
}
