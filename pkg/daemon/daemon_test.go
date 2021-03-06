// Copyright 2018 Dmitry Kargashin <dkargashin3@gmail.com>
// Use of this source code is governed by GNU LGPL
// license that can be found in the LICENSE file.

package daemon_test

import (
	"database/sql"
	"os"
	"syscall"
	"testing"
	"time"

	"bmstu.codes/developers34/SBWeb/pkg/s3"

	"bmstu.codes/developers34/SBWeb/pkg/api"
	"bmstu.codes/developers34/SBWeb/pkg/daemon"
	"bmstu.codes/developers34/SBWeb/pkg/db"
	"bmstu.codes/developers34/SBWeb/pkg/sessionmanager"
)

// must be executed from docker container linked with postgres and redis
func TestRunService(t *testing.T) {
	cfg := &daemon.Config{
		DB: db.Config{
			DBAddress:    "postgresql://runner:@postgres/data?sslmode=disable",
			MaxOpenConns: 10,
		},
		SM: sessionmanager.Config{
			DBAddress:      "redis://redis:6379/0",
			TockenLength:   32,
			ExpirationTime: 86400,
		},
		API: api.Config{
			Address:      ":54000",
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
			IdleTimeout:  "10s",
		},
		IM: s3.Config{
			Region: "eu-central-1",
		},
	}

	database, _ := sql.Open("postgres", "postgresql://runner:@postgres/data?sslmode=disable")
	database.Exec(`
	CREATE TABLE IF NOT EXISTS users
(
    id                SERIAL      PRIMARY KEY,
    first_name        varchar(80) NOT NULL,
    last_name         varchar(80) NOT NULL,
    email             varchar(80) UNIQUE NOT NULL,
    password_hash     text        NOT NULL,
    telephone         varchar(80),
    about             text,
    avatar_address    text,
    reg_time          timestamp   DEFAULT CURRENT_TIMESTAMP NOT NULL
);`)
	database.Exec(`
	CREATE TABLE IF NOT EXISTS ads
(
    id             SERIAL       PRIMARY KEY,
    title          varchar(80)  NOT NULL,
    price          integer      CONSTRAINT positive_price CHECK (price > 0),
    country        varchar(80),
    city           varchar(80),
    subway_station varchar(80),
    ad_images      varchar(256)[],
    -- when deleting user we should delete his ads
    owner_ad       integer      REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    description_ad text,
    creation_time  timestamp    DEFAULT CURRENT_TIMESTAMP NOT NULL
);`)

	database.Close()

	ch := make(chan error)
	go func() {
		ch <- daemon.RunService(cfg)
	}()
	time.Sleep(time.Millisecond * 1000)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	time.Sleep(time.Millisecond * 2000)

	if err := <-ch; err != nil {
		t.Error("Unexpected error: ", err.Error())
	}

	cfg = &daemon.Config{
		DB: db.Config{
			DBAddress:    "postgresql://runner:@badhost/data?sslmode=disable",
			MaxOpenConns: 10,
		},
		SM: sessionmanager.Config{
			DBAddress:      "redis://redis:6379/0",
			TockenLength:   32,
			ExpirationTime: 86400,
		},
		API: api.Config{
			Address:      ":54000",
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
			IdleTimeout:  "10s",
		},
		IM: s3.Config{
			Region: "eu-central-1",
		},
	}

	go func() {
		ch <- daemon.RunService(cfg)
	}()

	if err := <-ch; err == nil {
		t.Error("Error must be not nil")
	}

	cfg = &daemon.Config{
		DB: db.Config{
			DBAddress:    "postgresql://runner:@postgres/data?sslmode=disable",
			MaxOpenConns: 10,
		},
		SM: sessionmanager.Config{
			DBAddress:      "redis:/dwedfewfewfd:6379/0",
			TockenLength:   32,
			ExpirationTime: 86400,
		},
		API: api.Config{
			Address:      ":54000",
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
			IdleTimeout:  "10s",
		},
		IM: s3.Config{
			Region: "eu-central-1",
		},
	}

	go func() {
		ch <- daemon.RunService(cfg)
	}()

	if err := <-ch; err == nil {
		t.Error("Error must be not nil")
	}

	cfg = &daemon.Config{
		DB: db.Config{
			DBAddress:    "postgresql://runner:@postgres/data?sslmode=disable",
			MaxOpenConns: 10,
		},
		SM: sessionmanager.Config{
			DBAddress:      "redis://redis:6379/0",
			TockenLength:   32,
			ExpirationTime: 86400,
		},
		API: api.Config{
			Address:      "wefew",
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
			IdleTimeout:  "10s",
		},
		IM: s3.Config{
			Region: "eu-central-1",
		},
	}

	go func() {
		ch <- daemon.RunService(cfg)
	}()

	if err := <-ch; err != nil {
		t.Error("Expected error")
	}

	cfg = &daemon.Config{
		DB: db.Config{
			DBAddress:    "postgresql://runner:@badhost/data?sslmode=disable",
			MaxOpenConns: 10,
		},
		SM: sessionmanager.Config{
			DBAddress:      "redis://redis:6379/0",
			TockenLength:   32,
			ExpirationTime: 86400,
		},
		API: api.Config{
			Address:      ":54000",
			ReadTimeout:  "10s",
			WriteTimeout: "10s",
			IdleTimeout:  "10s",
		},
		IM: s3.Config{
			Region: "eu-central-1",
		},
	}

	os.Setenv("AWS_ACCESS_KEY_ID", "bad and strange id")

	go func() {
		ch <- daemon.RunService(cfg)
	}()

	if err := <-ch; err == nil {
		t.Error("Error must be not nil")
	}
}
