package main

type SqlType string

type Sql struct {
	SqlType   SqlType // sql 类型
	TableName string
	Fields    []string
}