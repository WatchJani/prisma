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

type Postgres struct {
	*sql.DB
	PostgresQuery
}

func NewDB(driverName string, c DBConfig) (*Postgres, error) {
	conn, err := sql.Open(driverName, fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", c.Username, c.DBName, c.Password, c.SSLmode))

	if err != nil {
		return nil, err
	}

	return &Postgres{
		conn, PostgresQuery{},
	}, nil
}

type PostgresQuery struct {
	Fields   []string
	Tables   []string
	Function string
}

func (p *Postgres) Field(fields ...string) QueryDB {
	p.Fields = fields
	return p
}

func (p *Postgres) Table(table ...string) QueryDB {
	p.Tables = table
	return p
}

func (q *Postgres) Select() {
	//another func
	fields := fmt.Sprint("*")

	if len(q.Fields) > 0 {
		fields = strings.Join(q.Fields, ",")
	}

	//another func
	table := q.Tables

	fmt.Println(fmt.Sprintf("SELECT %s FROM %s", fields, table))
}
