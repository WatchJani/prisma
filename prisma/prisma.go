package prisma

type QueryDB interface {
	Field(fields ...string) QueryDB
	Table(table ...string) QueryDB
	Select()
}
