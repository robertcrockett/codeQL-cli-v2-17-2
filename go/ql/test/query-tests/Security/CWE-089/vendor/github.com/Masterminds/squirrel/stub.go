// Code generated by depstubber. DO NOT EDIT.
// This is a simple stub for github.com/Masterminds/squirrel, strictly for use in testing.

// See the LICENSE file for information about the licensing of the original library.
// Source: github.com/Masterminds/squirrel (exports: ; functions: Expr,StatementBuilder)

// Package squirrel is a stub of github.com/Masterminds/squirrel, generated by depstubber.
package squirrel

import (
	context "context"
	sql "database/sql"
)

type BaseRunner interface {
	Exec(_ string, _ ...interface{}) (sql.Result, error)
	Query(_ string, _ ...interface{}) (*sql.Rows, error)
}

type DeleteBuilder struct{}

func (_ DeleteBuilder) Exec() (sql.Result, error) {
	return nil, nil
}

func (_ DeleteBuilder) ExecContext(_ context.Context) (sql.Result, error) {
	return nil, nil
}

func (_ DeleteBuilder) From(_ string) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) Limit(_ uint64) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) Offset(_ uint64) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) OrderBy(_ ...string) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) PlaceholderFormat(_ PlaceholderFormat) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) Prefix(_ string, _ ...interface{}) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) Query() (*sql.Rows, error) {
	return nil, nil
}

func (_ DeleteBuilder) RunWith(_ BaseRunner) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) Suffix(_ string, _ ...interface{}) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ DeleteBuilder) ToSql() (string, []interface{}, error) {
	return "", nil, nil
}

func (_ DeleteBuilder) Where(_ interface{}, _ ...interface{}) DeleteBuilder {
	return DeleteBuilder{}
}

func Expr(_ string, _ ...interface{}) interface{} {
	return nil
}

type InsertBuilder struct{}

func (_ InsertBuilder) Columns(_ ...string) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) Exec() (sql.Result, error) {
	return nil, nil
}

func (_ InsertBuilder) ExecContext(_ context.Context) (sql.Result, error) {
	return nil, nil
}

func (_ InsertBuilder) Into(_ string) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) Options(_ ...string) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) PlaceholderFormat(_ PlaceholderFormat) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) Prefix(_ string, _ ...interface{}) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) Query() (*sql.Rows, error) {
	return nil, nil
}

func (_ InsertBuilder) QueryContext(_ context.Context) (*sql.Rows, error) {
	return nil, nil
}

func (_ InsertBuilder) QueryRow() RowScanner {
	return nil
}

func (_ InsertBuilder) QueryRowContext(_ context.Context) RowScanner {
	return nil
}

func (_ InsertBuilder) RunWith(_ BaseRunner) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) Scan(_ ...interface{}) error {
	return nil
}

func (_ InsertBuilder) ScanContext(_ context.Context, _ ...interface{}) error {
	return nil
}

func (_ InsertBuilder) Select(_ SelectBuilder) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) SetMap(_ map[string]interface{}) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) Suffix(_ string, _ ...interface{}) InsertBuilder {
	return InsertBuilder{}
}

func (_ InsertBuilder) ToSql() (string, []interface{}, error) {
	return "", nil, nil
}

func (_ InsertBuilder) Values(_ ...interface{}) InsertBuilder {
	return InsertBuilder{}
}

type PlaceholderFormat interface {
	ReplacePlaceholders(_ string) (string, error)
}

type RowScanner interface {
	Scan(_ ...interface{}) error
}

type SelectBuilder struct{}

func (_ SelectBuilder) Column(_ interface{}, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Columns(_ ...string) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Distinct() SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Exec() (sql.Result, error) {
	return nil, nil
}

func (_ SelectBuilder) ExecContext(_ context.Context) (sql.Result, error) {
	return nil, nil
}

func (_ SelectBuilder) From(_ string) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) FromSelect(_ SelectBuilder, _ string) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) GroupBy(_ ...string) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Having(_ interface{}, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Join(_ string, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) JoinClause(_ interface{}, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) LeftJoin(_ string, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Limit(_ uint64) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) MustSql() (string, []interface{}) {
	return "", nil
}

func (_ SelectBuilder) Offset(_ uint64) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Options(_ ...string) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) OrderBy(_ ...string) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) PlaceholderFormat(_ PlaceholderFormat) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Prefix(_ string, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Query() (*sql.Rows, error) {
	return nil, nil
}

func (_ SelectBuilder) QueryContext(_ context.Context) (*sql.Rows, error) {
	return nil, nil
}

func (_ SelectBuilder) QueryRow() RowScanner {
	return nil
}

func (_ SelectBuilder) QueryRowContext(_ context.Context) RowScanner {
	return nil
}

func (_ SelectBuilder) RemoveLimit() SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) RightJoin(_ string, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) RunWith(_ BaseRunner) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) Scan(_ ...interface{}) error {
	return nil
}

func (_ SelectBuilder) ScanContext(_ context.Context, _ ...interface{}) error {
	return nil
}

func (_ SelectBuilder) Suffix(_ string, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

func (_ SelectBuilder) ToSql() (string, []interface{}, error) {
	return "", nil, nil
}

func (_ SelectBuilder) Where(_ interface{}, _ ...interface{}) SelectBuilder {
	return SelectBuilder{}
}

var StatementBuilder StatementBuilderType = StatementBuilderType{}

type StatementBuilderType struct{}

func (_ StatementBuilderType) Delete(_ string) DeleteBuilder {
	return DeleteBuilder{}
}

func (_ StatementBuilderType) Insert(_ string) InsertBuilder {
	return InsertBuilder{}
}

func (_ StatementBuilderType) PlaceholderFormat(_ PlaceholderFormat) StatementBuilderType {
	return StatementBuilderType{}
}

func (_ StatementBuilderType) RunWith(_ BaseRunner) StatementBuilderType {
	return StatementBuilderType{}
}

func (_ StatementBuilderType) Select(_ ...string) SelectBuilder {
	return SelectBuilder{}
}

func (_ StatementBuilderType) Update(_ string) UpdateBuilder {
	return UpdateBuilder{}
}

type UpdateBuilder struct{}

func (_ UpdateBuilder) Exec() (sql.Result, error) {
	return nil, nil
}

func (_ UpdateBuilder) ExecContext(_ context.Context) (sql.Result, error) {
	return nil, nil
}

func (_ UpdateBuilder) Limit(_ uint64) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) Offset(_ uint64) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) OrderBy(_ ...string) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) PlaceholderFormat(_ PlaceholderFormat) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) Prefix(_ string, _ ...interface{}) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) Query() (*sql.Rows, error) {
	return nil, nil
}

func (_ UpdateBuilder) QueryContext(_ context.Context) (*sql.Rows, error) {
	return nil, nil
}

func (_ UpdateBuilder) QueryRow() RowScanner {
	return nil
}

func (_ UpdateBuilder) QueryRowContext(_ context.Context) RowScanner {
	return nil
}

func (_ UpdateBuilder) RunWith(_ BaseRunner) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) Scan(_ ...interface{}) error {
	return nil
}

func (_ UpdateBuilder) ScanContext(_ context.Context, _ ...interface{}) error {
	return nil
}

func (_ UpdateBuilder) Set(_ string, _ interface{}) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) SetMap(_ map[string]interface{}) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) Suffix(_ string, _ ...interface{}) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) Table(_ string) UpdateBuilder {
	return UpdateBuilder{}
}

func (_ UpdateBuilder) ToSql() (string, []interface{}, error) {
	return "", nil, nil
}

func (_ UpdateBuilder) Where(_ interface{}, _ ...interface{}) UpdateBuilder {
	return UpdateBuilder{}
}
