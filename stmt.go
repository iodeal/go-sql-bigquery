package bigquery

import (
	"context"
	"database/sql/driver"

	"github.com/sirupsen/logrus"
)

//Stmt A query statement
type Stmt struct {
	query string
	c     *Conn
}

//NewStmt New Statement
func NewStmt(query string, c *Conn) *Stmt {
	return &Stmt{query: query, c: c}
}

//Close close Stmt
func (s *Stmt) Close() error {
	return nil
}

//NumInput number of input
func (s *Stmt) NumInput() int {
	return -1
}

//Exec Deprecated: Drivers should implement StmtExecContext instead (or additionally).
func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	logrus.Debugf("Got stmt.Exec: %s", s.query)
	return s.c.Exec(s.query, args)
}

//ExecContext Exec with Context parameter
func (s *Stmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	return s.c.ExecContext(ctx, s.query, args)
}

//Query Deprecated: Drivers should implement StmtQueryContext instead (or additionally).
func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	logrus.Debugf("Got stmt.Query: %s", s.query)
	return s.c.Query(s.query, args)
}

//QueryContext Query with Context parameter
func (s *Stmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	return s.c.QueryContext(ctx, s.query, args)
}
