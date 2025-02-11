package database

import (
    "context"
    "log"
    "testing"
    "time"

    "github.com/testcontainers/testcontainers-go"
    "github.com/testcontainers/testcontainers-go/modules/postgres"
    "github.com/testcontainers/testcontainers-go/wait"
)

func mustStartPostgresContainer() (func(context.Context) error, error) {
    var (
        dbName   = "todos"
        dbPwd    = "password"
        dbUser   = "user"
        dbSchema = "public" // Schema hinzufügen
    )

    dbContainer, err := postgres.RunContainer(
        context.Background(),
        testcontainers.WithImage("postgres:latest"),
        postgres.WithDatabase(dbName),
        postgres.WithUsername(dbUser),
        postgres.WithPassword(dbPwd),
        testcontainers.WithWaitStrategy(
            wait.ForLog("database system is ready to accept connections").
                WithOccurrence(2).
                WithStartupTimeout(5*time.Second)),
        )
    if err != nil {
        return nil, err
    }

    database = dbName
    password = dbPwd
    username = dbUser
    schema = dbSchema // Schema setzen

    dbHost, err := dbContainer.Host(context.Background())
    if err != nil {
        return dbContainer.Terminate, err
    }

    dbPort, err := dbContainer.MappedPort(context.Background(), "5432")
    if err != nil {
        return dbContainer.Terminate, err
    }

    host = dbHost
    port = dbPort.Port()

    return dbContainer.Terminate, nil
}

func TestMain(m *testing.M) {
    teardown, err := mustStartPostgresContainer()
    if err != nil {
        log.Fatalf("could not start postgres container: %v", err)
    }

    m.Run()

    if teardown != nil {
        if err := teardown(context.Background()); err != nil {
            log.Fatalf("could not teardown postgres container: %v", err)
        }
    }
}

func TestNew(t *testing.T) {
    srv, err := New()
    if err != nil {
        t.Fatalf("New() returned error: %v", err)
    }
    if srv == nil {
        t.Fatal("New() returned nil")
    }
}

func TestHealth(t *testing.T) {
    srv, err := New()
    if err != nil {
        t.Fatalf("New() returned error: %v", err)
    }

    stats := srv.Health()

    if stats["status"] != "up" {
        t.Fatalf("expected status to be up, got %s", stats["status"])
    }

    if _, ok := stats["error"]; ok {
        t.Fatalf("expected error not to be present")
    }

    if stats["message"] != "It's healthy" {
        t.Fatalf("expected message to be 'It's healthy', got %s", stats["message"])
    }
}

func TestClose(t *testing.T) {
    srv, err := New()
    if err != nil {
        t.Fatalf("New() returned error: %v", err)
    }

    if err := srv.Close(); err != nil {
        t.Fatalf("expected Close() to return nil, got %v", err)
    }
}
