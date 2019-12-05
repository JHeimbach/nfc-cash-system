package mysql

import (
	"database/sql"

	"github.com/JHeimbach/nfc-cash-system/server/api"
	"github.com/JHeimbach/nfc-cash-system/server/models"
	"github.com/go-sql-driver/mysql"
)

// AccountModel provides API for the accounts table
type AccountModel struct {
	db *sql.DB
}

// Create inserts new account it returns error models.ErrGroupNotFound if the groupId is not associated with a group
// it returns models.ErrDuplicateNfcChipId if the provided nfcchipid is already in the database present
func (a *AccountModel) Create(name, description string, startSaldo float64, groupId int, nfcChipId string) (int, error) {
	nullDescription := createNullableString(description)

	createStmt := `INSERT INTO accounts (name, description, saldo, group_id, nfc_chip_uid) VALUES (?,?,?,?,?)`

	res, err := a.db.Exec(createStmt, name, nullDescription, startSaldo, groupId, nfcChipId)

	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1452 {
				return 0, models.ErrGroupNotFound
			}
			if err.Number == 1062 {
				return 0, models.ErrDuplicateNfcChipId
			}
		}
		return 0, err
	}

	// mysql result returns no error, so we can ignore it
	lastId, _ := res.LastInsertId()

	return int(lastId), nil
}

// Read returns account struct for given id
func (a *AccountModel) Read(id int) (*api.Account, error) {
	readStmt := `SELECT id, name, description, saldo, group_id, nfc_chip_uid FROM accounts WHERE id=?`

	m := &api.Account{}
	row := a.db.QueryRow(readStmt, id)
	var nullDesc sql.NullString
	err := row.Scan(&m.Id, &m.Name, &nullDesc, &m.Saldo, &m.GroupId, &m.NfcChipId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	m.Description = decodeNullableString(nullDesc)

	return m, nil
}

// Update saves the (changed) model in the database will return models.ErrGroupNotFound if group id is not associated with a group
func (a *AccountModel) Update(m api.Account) error {
	updateStmt := `UPDATE accounts SET name=?, description=?, saldo=?, group_id=?, nfc_chip_uid=? WHERE id=?`

	_, err := a.db.Exec(updateStmt, m.Name, m.Description, m.Saldo, m.GroupId, m.NfcChipId, m.Id)

	if err != nil {
		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1452 {
				return models.ErrGroupNotFound
			}
		}
		return err
	}

	return nil
}

// Delete deletes a account
func (a *AccountModel) Delete(id int32) error {

	deleteStmt := `DELETE FROM accounts WHERE id=?`

	_, err := a.db.Exec(deleteStmt, id)

	if err == sql.ErrNoRows {
		return models.ErrNotFound
	}

	return err
}

// UpdateSaldo provides a simpler update method for the saldo field
func (a *AccountModel) UpdateSaldo(id int32, newSaldo float64) error {
	_, err := a.db.Exec("UPDATE accounts SET saldo=? WHERE id=?", newSaldo, id)

	return err
}

// GetAll returns slice with all accounts in the database
func (a *AccountModel) GetAll() ([]*api.Account, error) {
	rows, err := a.db.Query("SELECT id, name, description, saldo, group_id FROM accounts")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	return scanRowsToAccounts(rows)
}

// GetAllByGroup returns slice with all accounts with given group id
func (a *AccountModel) GetAllByGroup(groupId int) ([]*api.Account, error) {
	rows, err := a.db.Query("SELECT id, name, description, saldo, group_id FROM accounts WHERE group_id=?", groupId)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	return scanRowsToAccounts(rows)
}

// scanRowsToAccounts returns slice of Accounts from given sql.Rows
func scanRowsToAccounts(rows *sql.Rows) ([]*api.Account, error) {
	var accounts []*api.Account

	for rows.Next() {
		s := &api.Account{}

		var nullDesc sql.NullString

		err := rows.Scan(&s.Id, &s.Name, &nullDesc, &s.Saldo, &s.GroupId)
		if err != nil {
			return nil, err
		}

		s.Description = decodeNullableString(nullDesc)

		accounts = append(accounts, s)
	}

	return accounts, nil
}
