package db

impot (
	"database/sql"
	"log"
	_"github.com/lib/pq"
	"gitgub.co,/pressly/goose/v3"
)

func RunMigrations(db *sql.DB, migrationsDir string) {
	if err := goose.SetDialect("postgress"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	if err := goose.Up(db, migrationsDir); err != nil {
		log.Fatalf("Failed to apply migrations")
	}
}