package bigquery

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
)

func setupRowsTest(t testing.TB) func(t testing.TB) {
	cfg, err := ConfigFromConnString(testConnectionString)
	if err != nil {
		t.Fatal(err)
	}
	testConn, err = NewConn(context.TODO(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)
	return func(t testing.TB) {

	}
}
