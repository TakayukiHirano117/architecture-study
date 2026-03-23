package testhelper

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/TakayukiHirano117/architecture-study/config"
)

type TestDB struct {
	conn *sqlx.DB
	tx   *sqlx.Tx
}

// CheckDBTestEnv はDBテストの環境をチェックする（TestMain用）
// 条件を満たさない場合はエラーを返す
func CheckDBTestEnv() error {
	_ = godotenv.Load("../../../.env")

	goEnv := os.Getenv("GO_ENV")
	if goEnv != "test" {
		return fmt.Errorf("GO_ENV=%s（test以外）。'make test-docker' を使用してください", goEnv)
	}

	if os.Getenv("TEST_DB_NAME") == "" {
		return fmt.Errorf("TEST_DB_NAME未設定。.envにTEST_DB_NAMEを追加してください")
	}

	return nil
}

func NewTestDB(t *testing.T) *TestDB {
	t.Helper()

	_ = godotenv.Load("../../../.env")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("TEST_DB_NAME")

	t.Logf("Using test database: %s", dbName)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	return &TestDB{conn: conn}
}

func (db *TestDB) BeginTx(t *testing.T) {
	t.Helper()

	tx, err := db.conn.Beginx()
	if err != nil {
		t.Fatalf("failed to begin transaction: %v", err)
	}
	db.tx = tx
}

func (db *TestDB) Rollback(t *testing.T) {
	t.Helper()

	if db.tx != nil {
		if err := db.tx.Rollback(); err != nil {
			t.Errorf("failed to rollback transaction: %v", err)
		}
	}
}

func (db *TestDB) Close(t *testing.T) {
	t.Helper()

	if db.conn != nil {
		if err := db.conn.Close(); err != nil {
			t.Errorf("failed to close connection: %v", err)
		}
	}
}

func (db *TestDB) Ctx() context.Context {
	return context.WithValue(context.Background(), config.DBKey, db.tx)
}

func (db *TestDB) Tx() *sqlx.Tx {
	return db.tx
}

func SetupTestWithTx(t *testing.T) (*TestDB, context.Context) {
	t.Helper()

	testDB := NewTestDB(t)
	testDB.BeginTx(t)

	return testDB, testDB.Ctx()
}

func (db *TestDB) Cleanup(t *testing.T) {
	t.Helper()

	db.Rollback(t)
	db.Close(t)
}

func (db *TestDB) InsertTestUser(t *testing.T, id, name, email, password string) {
	t.Helper()

	query := `
		INSERT INTO users (id, name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`
	_, err := db.tx.Exec(query, id, name, email, password)
	if err != nil {
		t.Fatalf("failed to insert test user: %v", err)
	}
}

func (db *TestDB) InsertTestCategory(t *testing.T, id, name string) {
	t.Helper()

	query := `
		INSERT INTO categories (id, name, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
	`
	_, err := db.tx.Exec(query, id, name)
	if err != nil {
		t.Fatalf("failed to insert test category: %v", err)
	}
}

func (db *TestDB) InsertTestTag(t *testing.T, id, name string) {
	t.Helper()

	query := `
		INSERT INTO tags (id, name, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
	`
	_, err := db.tx.Exec(query, id, name)
	if err != nil {
		t.Fatalf("failed to insert test tag: %v", err)
	}
}

func (db *TestDB) GetPlanByID(t *testing.T, planID string) map[string]interface{} {
	t.Helper()

	query := `SELECT id, mentor_id, title, category_id, description, status, consultation_type, price FROM plans WHERE id = $1`
	row := db.tx.QueryRowx(query, planID)

	result := make(map[string]interface{})
	err := row.MapScan(result)
	if err != nil {
		t.Fatalf("failed to get plan: %v", err)
	}

	for key, val := range result {
		if b, ok := val.([]byte); ok {
			result[key] = string(b)
		}
	}

	return result
}

func (db *TestDB) GetPlanTagsByPlanID(t *testing.T, planID string) []string {
	t.Helper()

	query := `SELECT tag_id FROM plan_tags WHERE plan_id = $1`
	rows, err := db.tx.Queryx(query, planID)
	if err != nil {
		t.Fatalf("failed to get plan tags: %v", err)
	}
	defer rows.Close()

	var tagIDs []string
	for rows.Next() {
		var tagID string
		if err := rows.Scan(&tagID); err != nil {
			t.Fatalf("failed to scan tag id: %v", err)
		}
		tagIDs = append(tagIDs, tagID)
	}

	return tagIDs
}
