package prisma

import (
	"database/sql"
	"fmt"
	"strings"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLmode  string
}

func NewDB(driverName string, c DBConfig) (*sql.DB, error) {
	conn, err := sql.Open(driverName, fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", c.Username, c.DBName, c.Password, c.Port))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

type Property func(*Query)

type Query struct {
	Fields []string
	Table  []string
}

// func DefaultQuery() Query {
// 	return Query{
// 		Fields: []string{"*"},
// 	}
// }

func Select(query ...Property) string {
	//defaultQuery = DefaultQuery()
	defaultQuery := new(Query) //mnogo mi fora ovo new(&)

	for _, fn := range query {
		fn(defaultQuery)
	}

	return defaultQuery.Build()
}

func Field(fields ...string) Property {
	return func(q *Query) {
		q.Fields = fields
	}
}

func Table(table ...string) Property {
	return func(q *Query) {
		q.Table = table
	}
}

func (q Query) Build() string {
	//another func
	fields := fmt.Sprint("*")

	if len(q.Fields) > 0 {
		fields = strings.Join(q.Fields, ",")
	}

	//another func
	table := q.Table

	query := fmt.Sprintf("SELECT %s FROM %s", fields, table)

	return query
}
