package mysql

import (
	"testing"

	"github.com/Bel1Vuk/jet/v2/internal/testutils"
	. "github.com/Bel1Vuk/jettArrays/v2/mysql"
	. "github.com/Bel1Vuk/jettArrays/v2/tests/.gentestdata/mysql/dvds/table"
	"github.com/stretchr/testify/require"
)

func TestLockRead(t *testing.T) {
	query := Customer.LOCK().READ()

	testutils.AssertStatementSql(t, query, `
LOCK TABLES dvds.customer READ;
`)
	tx, err := db.DB.Begin() // can't prepare LOCK statement
	require.NoError(t, err)
	defer func() {
		err := tx.Rollback()
		require.NoError(t, err)
	}()

	testutils.AssertExec(t, query, tx)
}

func TestLockWrite(t *testing.T) {
	query := Customer.LOCK().WRITE()

	testutils.AssertStatementSql(t, query, `
LOCK TABLES dvds.customer WRITE;
`)

	tx, err := db.DB.Begin() // can't prepare LOCK statement
	require.NoError(t, err)
	defer func() {
		err := tx.Rollback()
		require.NoError(t, err)
	}()

	testutils.AssertExec(t, query, tx)
}

func TestUnlockTables(t *testing.T) {
	query := UNLOCK_TABLES()

	testutils.AssertStatementSql(t, query, `
UNLOCK TABLES;
`)

	tx, err := db.DB.Begin() // can't prepare LOCK statement
	require.NoError(t, err)
	defer func() {
		err := tx.Rollback()
		require.NoError(t, err)
	}()

	testutils.AssertExec(t, query, tx)
}
