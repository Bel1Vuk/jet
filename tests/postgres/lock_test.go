package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/Bel1Vuk/jet/v2/internal/testutils"
	"github.com/stretchr/testify/require"

	. "github.com/Bel1Vuk/jettArrays/v2/postgres"
	. "github.com/Bel1Vuk/jettArrays/v2/tests/.gentestdata/jetdb/dvds/table"
)

func TestLockTable(t *testing.T) {
	skipForCockroachDB(t) // doesn't support

	expectedSQL := `
LOCK TABLE dvds.address IN`

	var testData = []TableLockMode{
		LOCK_ACCESS_SHARE,
		LOCK_ROW_SHARE,
		LOCK_ROW_EXCLUSIVE,
		LOCK_SHARE_UPDATE_EXCLUSIVE,
		LOCK_SHARE,
		LOCK_SHARE_ROW_EXCLUSIVE,
		LOCK_EXCLUSIVE,
		LOCK_ACCESS_EXCLUSIVE,
	}

	for _, lockMode := range testData {
		query := Address.LOCK().IN(lockMode)

		testutils.AssertDebugStatementSql(t, query, expectedSQL+" "+string(lockMode)+" MODE;\n")
		testutils.AssertExecAndRollback(t, query, db)
	}

	for _, lockMode := range testData {
		query := Address.LOCK().IN(lockMode).NOWAIT()

		testutils.AssertDebugStatementSql(t, query, expectedSQL+" "+string(lockMode)+" MODE NOWAIT;\n")
		testutils.AssertExecAndRollback(t, query, db)
	}
}

func TestLockExecContext(t *testing.T) {
	skipForCockroachDB(t)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()

	time.Sleep(10 * time.Millisecond)

	tx, _ := db.Begin()
	defer tx.Rollback()

	_, err := Address.LOCK().IN(LOCK_ACCESS_SHARE).ExecContext(ctx, tx)

	require.Error(t, err, "context deadline exceeded")
}
