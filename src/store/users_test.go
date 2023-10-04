package store_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/joho/godotenv/autoload"
	"github.com/soramon0/portfolio/src/internal/database"
	"github.com/soramon0/portfolio/src/lib"
	"github.com/soramon0/portfolio/src/store"
)

type checkErrFn func(error) (stopTest bool, wantErr error)

func TestCreateUsers(t *testing.T) {
	setup := func(t *testing.T) store.Store {
		t.Helper()

		db, err := store.NewStore(lib.GetTestDatabaseURL())
		if err != nil {
			t.Fatal("failed to connect to db: ", err)
		}

		if err := db.Migrate("../sql/schema/", "up"); err != nil {
			t.Fatal("failed to migrate db up: ", err)
		}

		if _, err := db.Exec("DELETE FROM users"); err != nil {
			t.Fatal("failed to clean users table: ", err)
		}

		return db
	}

	teardown := func(t *testing.T, db store.Store) {
		t.Helper()

		defer func() {
			if err := db.Close(); err != nil {
				t.Fatal("failed to close db connection: ", err)
			}
		}()

		if err := db.Migrate("../sql/schema/", "down"); err != nil {
			t.Fatal("failed to migrate db down: ", err)
		}
	}

	tests := map[string]struct {
		user    *database.User
		checkFn checkErrFn
	}{
		"user type": {
			user: &database.User{
				ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
				CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				Email:     "test1@email.com",
				Password:  "password",
				Username:  "username1",
				UserType:  database.UserTypeUser,
				FirstName: "",
				LastName:  "",
			},
			checkFn: func(err error) (bool, error) {
				if err != nil {
					return true, fmt.Errorf("CreateUser() err = %v, want nil", err)
				}
				return false, nil
			},
		},
		"admin type": {
			user: &database.User{
				ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
				CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				Email:     "test2@email.com",
				Password:  "password",
				Username:  "username2",
				UserType:  database.UserTypeAdmin,
				FirstName: "",
				LastName:  "",
			},
			checkFn: func(err error) (bool, error) {
				if err != nil {
					return true, fmt.Errorf("CreateUser() err = %v, want nil", err)
				}
				return false, nil
			},
		},
		"wrong type": {
			user: &database.User{
				ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
				CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				UpdatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
				Email:     "test2@email.com",
				Password:  "password",
				Username:  "username2",
				UserType:  "wrong type",
				FirstName: "",
				LastName:  "",
			},
			checkFn: func(err error) (bool, error) {
				if err == nil {
					return true, fmt.Errorf("CreateUser() err = nil, want constraint violation error")
				}

				pgErr, ok := err.(*pgconn.PgError)
				if !ok {
					return true, fmt.Errorf("CreateUser() err.(type) = %v; want err.(*pgconn.PgError)", err)
				}
				if !strings.Contains(pgErr.Message, `invalid input value for enum user_type`) {
					return true, fmt.Errorf("CreateUser() err.Error = %v; want enum type error", pgErr.Error())
				}

				return true, nil
			},
		},
	}

	for name, want := range tests {
		t.Run(name, func(t *testing.T) {
			db := setup(t)
			defer teardown(t, db)
			testCreateUser(t, db, want.user, want.checkFn)
		})
	}
}

func testCreateUser(t *testing.T, db store.Store, want *database.User, checkFn checkErrFn) {
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

	stopTest, wantErr := checkFn(err)
	if wantErr != nil {
		t.Fatal(wantErr)
	}
	if stopTest {
		return
	}

	expectUserEq(t, &user, want)
	expectRowsCount(t, db, "users", 1)

	userById, err := db.GetUserById(context.TODO(), user.ID)
	if err != nil {
		t.Fatalf("GetUserById() err = %v, want nil", err)
	}

	// if user.ID.String() != userById.ID.String() {
	// 	t.Fatalf("user.ID != userById.ID; got %v; want %v", user.ID, userById.ID)
	// }
	if !user.CreatedAt.Time.Equal(userById.CreatedAt.Time) {
		t.Fatalf("user.CreatedAt != userById.CreatedAt; got %v; want %v", user.CreatedAt, userById.CreatedAt)
	}
	if !user.UpdatedAt.Time.Equal(userById.UpdatedAt.Time) {
		t.Fatalf("user.UpdatedAt != userById.UpdatedAt; got %v; want %v", user.UpdatedAt, userById.UpdatedAt)
	}
}

func expectUserEq(t *testing.T, got, want *database.User) {
	if got == want {
		return
	}

	// if got.ID.String() != want.ID.String() {
	// 	t.Fatalf("user.ID = %v; want %v", got.ID, want.ID)
	// }
	// time.Equal fails because of nanosecond difference
	// using time.Sub make sure the difference is less
	// than 1 second between the two dates
	if got.CreatedAt.Time.Sub(want.CreatedAt.Time) > time.Second {
		t.Fatalf("user.CreatedAt = %v; want %v", got.CreatedAt, want.CreatedAt)
	}
	if got.UpdatedAt.Time.Sub(want.UpdatedAt.Time) > time.Second {
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
