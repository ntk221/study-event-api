package handlers_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"study-event-api/config"
	"study-event-api/internal/db"
	"study-event-api/internal/handlers"
	"testing"
	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	testConfig *config.Config
	fixtures   *testfixtures.Loader
	testDB     *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testConfig, err = config.Load()
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	testDB, err = sql.Open("postgres", testConfig.GetDSN())
	if err != nil {
		fmt.Printf("Failed to open database connection: %v\n", err)
		os.Exit(1)
	}
	defer testDB.Close()

	if err := prepare(testDB); err != nil {
		fmt.Printf("Failed to prepare test database: %v\n", err)
		os.Exit(1)
	}

	txdb.Register("txdb", "postgres", testConfig.GetDSN())

	code := m.Run()
	os.Exit(code)
}

func prepare(db *sql.DB) error {
	var err error
	fixtures, err = testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("postgres"),
		testfixtures.Directory(fixturesPath),
	)
	if err != nil {
		return fmt.Errorf("failed to create fixtures loader: %v", err)
	}
	return fixtures.Load()
}

func TestEventHandler(t *testing.T) {
	txDB, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}
	defer txDB.Close()

	queries := db.New(txDB)
	handler := handlers.EventHandler(queries)

	req, err := http.NewRequest("GET", "/events", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var events []db.Event
	err = json.Unmarshal(rr.Body.Bytes(), &events)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	assert.Len(t, events, 2)

	expectedEvents := []db.Event{
		{ID: 1, Name: "Test Event 1", Date: time.Date(2023, 8, 5, 15, 0, 0, 0, time.UTC)},
		{ID: 2, Name: "Test Event 2", Date: time.Date(2023, 8, 6, 10, 0, 0, 0, time.UTC)},
	}

	for i, expectedEvent := range expectedEvents {
		assert.Equal(t, expectedEvent.ID, events[i].ID)
		assert.Equal(t, expectedEvent.Name, events[i].Name)
		assert.Equal(t, expectedEvent.Date, events[i].Date)
	}
}
