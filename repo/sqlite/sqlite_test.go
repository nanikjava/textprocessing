package sqlite

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"rockt/model"
	"rockt/repo"
	"testing"
)

var repository repo.Repository

func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	repository, err = NewRepository()
	repository.Create()
	return m.Run(), nil

}

func TestBulkInsert(t *testing.T) {
	m := []model.Datarecord{model.Datarecord{
		DateISO8601:  "2000-10-10T00:58:42Z",
		EmailAddress: "1@1.com",
		SessionID:    "1",
	}, model.Datarecord{
		DateISO8601:  "2000-10-10T00:58:42Z",
		EmailAddress: "2@2.com",
		SessionID:    "2",
	}}

	err := repository.BulkInsert(m)

	require.NoError(t, err)
}
