package testdb

import (
	"database/sql/driver"
	"fmt"
	"io"
)

type rows struct {
	closed  bool
	columns []string
	rows    [][]driver.Value
	pos     int
}

func (rs *rows) clone() *rows {
	if rs == nil {
		return nil
	}

	return &rows{closed: false, columns: rs.columns, rows: rs.rows, pos: 0}
}

func (rs *rows) Next(dest []driver.Value) error {
	rs.pos++
	if rs.pos > len(rs.rows) {
		rs.closed = true

		return io.EOF // per interface spec
	}

	if len(rs.rows[rs.pos-1]) > len(dest) {
		return fmt.Errorf("row too long (%d) for dest columns (%d)", len(rs.rows[rs.pos-1]), len(dest))
	}
	for i, col := range rs.rows[rs.pos-1] {
		dest[i] = col
	}

	return nil
}

func (rs *rows) Err() error {
	return nil
}

func (rs *rows) Columns() []string {
	return rs.columns
}

func (rs *rows) Close() error {
	return nil
}
