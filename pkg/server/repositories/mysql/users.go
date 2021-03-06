package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jheimbach/nfc-cash-system/api"
	"github.com/jheimbach/nfc-cash-system/pkg/server/repositories"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository handles the users from the database
type UserRepository struct {
	db *sql.DB
}

func NewUserModel(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new userId in the database.
// if a userId with the same email already exists, Create will return a models.ErrDuplicateEmail
func (u *UserRepository) Create(ctx context.Context, name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	insertSql := `INSERT INTO users (name,email,hashed_password,created) VALUES(?,?,?,UTC_TIMESTAMP())`
	_, err = u.db.ExecContext(ctx, insertSql, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return repositories.ErrDuplicateEmail
			}
		}
	}

	return err
}

// Get returns a models.User from given id, if id does not exists Get will return a models.ErrNotFound
func (u *UserRepository) Get(ctx context.Context, id int) (*api.User, error) {
	m := &api.User{}
	var t time.Time
	getSql := `SELECT id,name,email,created FROM users WHERE id = ?`
	err := u.db.QueryRowContext(ctx, getSql, id).Scan(&m.Id, &m.Name, &m.Email, &t)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repositories.ErrNotFound
		}
		return nil, err
	}

	m.Created, err = ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Authenticate returns the id for a userId if it exists with given email and password
// if email does not exists or the password is wrong Authenticate will return a models.ErrInvalidCredentials
func (u *UserRepository) Authenticate(ctx context.Context, email, password string) (*api.User, error) {
	var user = &api.User{}
	var hashedPassword []byte
	var created time.Time
	row := u.db.QueryRowContext(ctx, "SELECT id, name, email, hashed_password, created FROM users WHERE email = ?", email)
	err := row.Scan(&user.Id, &user.Name, &user.Email, &hashedPassword, &created)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repositories.ErrInvalidCredentials
		}
		return nil, err
	}

	user.Created, _ = ptypes.TimestampProto(created)

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, repositories.ErrInvalidCredentials
		}
		return nil, err
	}

	return user, nil
}
