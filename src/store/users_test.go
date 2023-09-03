package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/store"
)

func TestCreateUsers(t *testing.T) {
	setup := func(t *testing.T) store.Store {
		t.Helper()

		db, err := store.NewStore("postgres://postgres:example@127.0.0.1:5433/test_db?sslmode=disable")
		if err != nil {
			t.Fatal("failed to connect to db: ", err)
		}

		if _, err := db.Exec("DELETE FROM users"); err != nil {
			t.Fatal("failed to clean users table: ", err)
		}

		return db
	}

	teardown := func(t *testing.T, db store.Store) {
		t.Helper()

		if _, err := db.Exec("DELETE FROM users"); err != nil {
			t.Fatal("failed to clean users table: ", err)
		}

		if err := db.Close(); err != nil {
			t.Fatal("failed to close db connection: ", err)
		}
	}

	tests := map[string]*database.User{
		"user type": {
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Email:     "test@email.com",
			Password:  "password",
			Username:  "username",
			UserType:  "user",
			FirstName: "",
			LastName:  "",
		},
		"admin type": {
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Email:     "test@email.com",
			Password:  "password",
			Username:  "username",
			UserType:  "admin",
			FirstName: "",
			LastName:  "",
		},
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			db := setup(t)
			defer teardown(t, db)
			testCreateUser(t, db, want)
		})
	}
}

func testCreateUser(t *testing.T, db store.Store, want *database.User) {
	expectRowsCount(t, db, "users", 0)

	user, err := db.CreateUser(context.TODO(), database.CreateUserParams{
		ID:        want.ID,
		CreatedAt: want.CreatedAt,
		UpdatedAt: want.CreatedAt,
		Email:     want.Email,
		Password:  want.Password,
		Username:  want.Username,
		UserType:  want.UserType,
		FirstName: want.FirstName,
		LastName:  want.LastName,
	})
	if err != nil {
		t.Fatalf("CreateUser() err = %v, want nil", err)
	}

	expectUserEq(t, &user, want)
	expectRowsCount(t, db, "users", 1)

	userById, err := db.GetUserById(context.TODO(), user.ID)
	if err != nil {
		t.Fatalf("GetUserById() err = %v, want nil", err)
	}

	if user.ID.String() != userById.ID.String() {
		t.Fatalf("user.ID != userById.ID; got %v; want %v", user.ID, userById.ID)
	}
	if !user.CreatedAt.Equal(userById.CreatedAt) {
		t.Fatalf("user.CreatedAt != userById.CreatedAt; got %v; want %v", user.CreatedAt, userById.CreatedAt)
	}
	if !user.UpdatedAt.Equal(userById.UpdatedAt) {
		t.Fatalf("user.UpdatedAt != userById.UpdatedAt; got %v; want %v", user.UpdatedAt, userById.UpdatedAt)
	}
}

func expectUserEq(t *testing.T, got, want *database.User) {
	if got == want {
		return
	}

	if got.ID.String() != want.ID.String() {
		t.Fatalf("user.ID = %v; want %v", got.ID, want.ID)
	}
	// time.Equal fails because of nanosecond difference
	// using time.Sub make sure the difference is less
	// than 1 second between the two dates
	if got.CreatedAt.Sub(want.CreatedAt) > time.Second {
		t.Fatalf("user.CreatedAt = %v; want %v", got.CreatedAt, want.CreatedAt)
	}
	if got.UpdatedAt.Sub(want.UpdatedAt) > time.Second {
		t.Fatalf("user.UpdatedAt = %v; want %v", got.UpdatedAt, want.UpdatedAt)
	}
	if got.Email != want.Email {
		t.Fatalf("user.Email = %v; want %v", got.Email, want.Email)
	}
	if got.Username != want.Username {
		t.Fatalf("user.Username = %v; want %v", got.Username, want.Username)
	}
	if got.UserType != want.UserType {
		t.Fatalf("user.UserType = %v; want %v", got.UserType, want.UserType)
	}
	if got.Password != want.Password {
		t.Fatalf("user.Password = %v; want %v", got.Password, want.Password)
	}
	if got.FirstName != want.FirstName {
		t.Fatalf("user.FirstName = %v; want %v", got.FirstName, want.FirstName)
	}
	if got.LastName != want.LastName {
		t.Fatalf("user.LastName = %v; want %v", got.LastName, want.LastName)
	}
}

func expectRowsCount(t *testing.T, db store.Store, table string, count int) {
	t.Helper()

	n := countRows(t, db, table)
	if n != count {
		t.Fatalf("%s count() = %d; want %d", table, n, count)
	}
}

func countRows(t *testing.T, db store.Store, table string) int {
	t.Helper()

	var n int
	if err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&n); err != nil {
		t.Fatalf("%s count() err = %v, want nil", table, err)
	}
	return n
}
