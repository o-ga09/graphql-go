package testutil

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

func GetRequestBody(t *testing.T, testdata, name string) io.Reader {
	t.Helper()

	queryBody, err := os.ReadFile(testdata + name + ".golden")
	if err != nil {
		t.Fatal(err)
	}
	query := struct{ Query string }{
		string(queryBody),
	}
	reqBody := bytes.Buffer{}
	if err := json.NewEncoder(&reqBody).Encode(&query); err != nil {
		t.Fatal("error encode", err)
	}
	return &reqBody
}

func GetResponseBody(t *testing.T, res *http.Response) string {
	t.Helper()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("error read body", err)
	}
	var got bytes.Buffer
	if err := json.Indent(&got, raw, "", "  "); err != nil {
		t.Fatal("json.Indent", err)
	}
	return got.String()
}

func SetUpMySQL(t *testing.T) {
	t.Helper()
	pwd, _ := os.Getwd()
	ddlpath := pwd + "/testdata/init-script/"
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("failed to create new pool: %v", err)
	}

	if err := pool.Client.Ping(); err != nil {
		t.Fatalf("failed to ping: %v", err)
	}

	opts := dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=pass",
			"MYSQL_DATABASE=testdb",
			"MYSQL_USER=user",
			"MYSQL_PASSWORD=pass",
		},
		Mounts: []string{
			ddlpath + ":/docker-entrypoint-initdb.d",
		},
	}

	hcopt := func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	}

	hcopts := []func(*docker.HostConfig){
		hcopt,
	}

	resource, err := pool.RunWithOptions(&opts, hcopts...)
	if err != nil {
		t.Fatalf("failed to run with options: %v", err)
	}
	port := resource.GetPort("3306/tcp")
	fmt.Println(port)
	t.Setenv("DATABASE_URL", fmt.Sprintf("user:pass@tcp(localhost:%s)/testdb?parseTime=true", port))
	if err := pool.Retry(func() error {
		db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
		db.SetConnMaxLifetime(5 * time.Second)
		if err != nil {
			t.Logf("failed to connect to db: %v", err)
			return err

		}
		if err := db.Ping(); err != nil {
			t.Log("waiting for db to be available", err)
			return err
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	t.Logf("completed to connect db")
	t.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			t.Fatal(err)
		}
	})
}
