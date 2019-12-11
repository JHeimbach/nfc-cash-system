package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/JHeimbach/nfc-cash-system/server/api"
	"github.com/JHeimbach/nfc-cash-system/server/models"
	isPkg "github.com/matryer/is"
)

type accountMockStorager struct {
	create func(name, description string, startSaldo float64, groupId int32, nfcChipId string) (*api.Account, error)
	list   func() (*api.Accounts, error)
	read   func(id int32) (*api.Account, error)
	delete func(id int32) error
	update func(m *api.Account) error
}

func (a accountMockStorager) Create(name, description string, startSaldo float64, groupId int32, nfcChipId string) (*api.Account, error) {
	return a.create(name, description, startSaldo, groupId, nfcChipId)
}

func (a accountMockStorager) GetAll() (*api.Accounts, error) {
	return a.list()
}

func (a accountMockStorager) Read(id int32) (*api.Account, error) {
	return a.read(id)
}

func (a accountMockStorager) Delete(id int32) error {
	return a.delete(id)
}

func (a accountMockStorager) Update(m *api.Account) error {
	return a.update(m)
}

func TestAccountserver_List(t *testing.T) {
	tests := []struct {
		name    string
		input   *api.AccountListRequest
		want    *api.Accounts
		wantErr error
	}{
		{
			name:  "get simple list of accounts",
			input: &api.AccountListRequest{},
			want: &api.Accounts{
				Accounts: getAccountModels(2),
			},
		},
		{
			name:    "has error",
			input:   &api.AccountListRequest{},
			want:    nil,
			wantErr: ErrGetAll,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &accountserver{
				storage: accountMockStorager{list: func() (*api.Accounts, error) {
					if tt.wantErr != nil {
						return nil, sql.ErrNoRows
					}
					return &api.Accounts{
						Accounts: getAccountModels(2),
					}, nil
				},
				},
			}
			got, err := a.List(context.Background(), tt.input)

			if err != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountserver_Get(t *testing.T) {
	is := isPkg.New(t)
	db := genAccountMap(3)

	tests := []struct {
		name    string
		input   *api.IdRequest
		want    *api.Account
		wantErr error
	}{
		{
			name:  "get account with id 1",
			input: &api.IdRequest{Id: 1},
			want:  db[1],
		},
		{
			name:  "get account with id 2",
			input: &api.IdRequest{Id: 2},
			want:  db[2],
		},
		{
			name:    "get account that does not exist",
			input:   &api.IdRequest{Id: 100},
			wantErr: ErrNotFound,
		},
	}

	server := accountserver{storage: accountMockStorager{
		read: func(id int32) (account *api.Account, err error) {
			if acc, ok := db[id]; ok {
				return acc, nil
			}
			return nil, models.ErrNotFound
		},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got, err := server.Get(context.Background(), tt.input)

			if tt.wantErr != nil {
				is.Equal(err, tt.wantErr) //expected error
				return
			}

			is.NoErr(err)
			is.Equal(got, tt.want)
		})
	}
}

func TestAccountserver_Create(t *testing.T) {

	is := isPkg.New(t)
	tests := []struct {
		name      string
		input     *api.AccountCreate
		want      *api.Account
		returnErr error
		wantErr   error
	}{
		{
			name: "create account",
			input: &api.AccountCreate{
				Name:        "test",
				Description: "",
				Saldo:       120,
				NfcChipId:   "nfcchip",
				GroupId:     1,
			},
			want: &api.Account{
				Id:          1,
				Description: "",
				Saldo:       120,
				NfcChipId:   "nfcchip",
				Group: &api.Group{
					Id: 1,
				},
			},
		},
		{
			name: "create account with same nfcchip",
			input: &api.AccountCreate{
				Name:        "test",
				Description: "",
				Saldo:       120,
				NfcChipId:   "nfcchip",
				GroupId:     1,
			},
			returnErr: models.ErrDuplicateNfcChipId,
			wantErr:   ErrCouldNotCreateAccount,
		},
		{
			name: "create account with unknown group",
			input: &api.AccountCreate{
				Name:        "test",
				Description: "",
				Saldo:       120,
				NfcChipId:   "nfcchip",
				GroupId:     100,
			},
			returnErr: models.ErrGroupNotFound,
			wantErr:   ErrCouldNotCreateAccount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			server := accountserver{
				storage: accountMockStorager{
					create: func(name, description string, startSaldo float64, groupId int32, nfcChipId string) (account *api.Account, err error) {
						if tt.returnErr != nil {
							return nil, tt.returnErr
						}
						return tt.want, nil
					},
				},
			}

			got, err := server.Create(context.Background(), tt.input)

			if tt.wantErr != nil {
				is.Equal(err, tt.wantErr)
				return
			}

			is.NoErr(err)
			is.Equal(got, tt.want) // got wrong account back
		})
	}
}

func TestAccountserver_Update(t *testing.T) {
	mockStorage := genAccountMap(3)
	is := isPkg.New(t)

	tests := []struct {
		name      string
		input     *api.Account
		returnErr error
		wantErr   error
	}{
		{
			name: "update account",
			input: &api.Account{
				Id:          1,
				Name:        "test",
				Description: "test",
				Saldo:       145,
				NfcChipId:   "nfc_chip_1",
				Group: &api.Group{
					Id: 1,
				},
			},
		},
		{
			name: "update returns error",
			input: &api.Account{
				Id:          1,
				Name:        "test",
				Description: "test",
				Saldo:       145,
				NfcChipId:   "nfc_chip_1",
				Group: &api.Group{
					Id: 1,
				},
			},
			returnErr: models.ErrGroupNotFound,
			wantErr:   ErrSomethingWentWrong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			server := accountserver{
				storage: accountMockStorager{
					update: func(m *api.Account) error {
						if tt.returnErr != nil {
							return tt.returnErr
						}
						mockStorage[m.Id] = m
						return nil
					},
				},
			}

			got, err := server.Update(context.Background(), tt.input)

			if tt.wantErr != nil {
				is.Equal(err, tt.wantErr) // error is not the expected
				return
			}

			is.NoErr(err)
			is.Equal(got, tt.input)            // returned account is not the same
			is.Equal(got, mockStorage[got.Id]) // account was not updated
		})
	}
}

func TestAccountserver_Delete(t *testing.T) {
	mockStorage := genAccountMap(3)
	is := isPkg.New(t)

	tests := []struct {
		name      string
		input     *api.IdRequest
		want      *api.Status
		returnErr error
	}{
		{
			name:  "delete account",
			input: &api.IdRequest{Id: 1},
			want: &api.Status{
				Success:      true,
				ErrorMessage: "",
			},
		},
		{
			name:      "delete account that does not exist",
			input:     &api.IdRequest{Id: 1},
			returnErr: models.ErrNotFound,
			want: &api.Status{
				Success:      false,
				ErrorMessage: ErrNotFound.Error(),
			},
		},
		{
			name:      "delete returns other than models.ErrNotFound",
			input:     &api.IdRequest{Id: 1},
			returnErr: errors.New("this is a test"),
			want: &api.Status{
				Success:      false,
				ErrorMessage: ErrSomethingWentWrong.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := accountserver{
				storage: accountMockStorager{
					delete: func(id int32) error {
						if tt.returnErr != nil {
							return tt.returnErr
						}
						delete(mockStorage, id)
						return nil
					},
				},
			}

			got, err := server.Delete(context.Background(), tt.input)
			is.NoErr(err)
			is.Equal(got, tt.want) // status is not correct
		})
	}

}

func getAccountModels(num int) []*api.Account {
	accounts := make([]*api.Account, 0, num)

	for i := 1; i <= num; i++ {
		accounts = append(accounts, &api.Account{
			Id:          int32(i),
			Name:        "test",
			Description: "test",
			Saldo:       0,
			NfcChipId:   fmt.Sprintf("ncf_chip_%d", i),
			Group: &api.Group{
				Id: 1,
			},
		})
	}
	return accounts
}

func genAccountMap(num int) map[int32]*api.Account {
	accounts := getAccountModels(num)
	m := make(map[int32]*api.Account)
	for _, account := range accounts {
		m[account.Id] = account
	}
	return m
}
