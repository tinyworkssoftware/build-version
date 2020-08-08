//+build unit

package repository

import (
	"build-version/model/data"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestUnitCreateOrganisation(t *testing.T) {
	t.Parallel()
	targetData := data.OrganisationData{
		Id: uuid.New().String(),
		Name: "testdata",
		PlanType: uuid.New().String(),
	}
	db, mock, _ := sqlmock.New()
	defer db.Close()
	mock.ExpectExec("INSERT INTO tbl_organisation").
		WithArgs(targetData.Id, targetData.Name, targetData.PlanType).
		WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectCommit()
	sqlxDB := sqlx.NewDb(db, "mysql")
	if err := CreateOrganisation(sqlxDB, &targetData); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
}

