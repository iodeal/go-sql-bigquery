package bigquery

import (
	"database/sql/driver"
	"io"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type bqRows struct {
	columns []string
	types   []string
	ri      *bigquery.RowIterator
	c       *Conn
}

func (b *bqRows) Columns() []string {
	return b.columns
}

func (b *bqRows) Close() error {
	return nil
}

func (b *bqRows) Next(dest []driver.Value) error {
	var values []bigquery.Value
	if err := b.ri.Next(&values); err != nil {
		if err == iterator.Done {
			return io.EOF
		}
		return err
	}
	for i, val := range values {
		if len(dest) > i {
			dest[i] = val
		} else {
			break
		}
	}
	return nil
}

func (b *bqRows) ColumnTypeDatabaseTypeName(index int) string {
	return b.types[index]
}
