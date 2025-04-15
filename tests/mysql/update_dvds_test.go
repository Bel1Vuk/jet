package mysql

import (
	"testing"

	"github.com/Bel1Vuk/jetArrays/v2/internal/testutils"
	. "github.com/Bel1Vuk/jetArraystArrays/v2/mysql"
	. "github.com/Bel1Vuk/jetArraystArrays/v2/tests/.gentestdata/mysql/dvds/table"
)

func TestUpdateWithJoin(t *testing.T) {
	statement := Staff.INNER_JOIN(Address, Address.AddressID.EQ(Staff.AddressID)).
		UPDATE(Staff.LastName).
		SET(String("New staff name")).
		WHERE(Staff.StaffID.EQ(Int(1)))

	testutils.AssertStatementSql(t, statement, `
UPDATE dvds.staff
INNER JOIN dvds.address ON (address.address_id = staff.address_id)
SET last_name = ?
WHERE staff.staff_id = ?;
`, "New staff name", int64(1))

	testutils.AssertExecAndRollback(t, statement, db)
}
