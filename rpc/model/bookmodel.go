package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	bookFieldNames          = builderx.RawFieldNames(&Book{})
	bookRows                = strings.Join(bookFieldNames, ",")
	bookRowsExpectAutoSet   = strings.Join(stringx.Remove(bookFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	bookRowsWithPlaceHolder = strings.Join(stringx.Remove(bookFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	BookModel interface {
		Insert(data Book) (sql.Result, error)
		FindOne(id int64) (*Book, error)
		Update(data Book) error
		Delete(id int64) error
		FindByName(name string) (*Book, error)
		FindAll() ([]Book, error)
	}

	defaultBookModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Book struct {
		Name      string         `db:"name"`  // book name
		Price     int64          `db:"price"` // book price
		CreatedAt sql.NullString `db:"created_at"`
		Id        int64          `db:"id"`
	}
)

func NewBookModel(conn sqlx.SqlConn) BookModel {
	return &defaultBookModel{
		conn:  conn,
		table: "`book`",
	}
}

func (m *defaultBookModel) Insert(data Book) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, bookRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Name, data.Price, data.CreatedAt)
	return ret, err
}

func (m *defaultBookModel) FindOne(id int64) (*Book, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", bookRows, m.table)
	var resp Book
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) FindByName(name string) (*Book, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", bookRows, m.table)
	var resp Book
	err := m.conn.QueryRow(&resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) FindAll() ([]Book, error) {
	query := fmt.Sprintf("select %s from %s ", bookRows, m.table)
	var resp []Book
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBookModel) Update(data Book) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, bookRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Name, data.Price, data.CreatedAt, data.Id)
	return err
}

func (m *defaultBookModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
