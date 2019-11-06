package mysql

import (
	"database/sql"
	"github.com/JHeimbach/nfc-cash-system/server/models"
	isPkg "github.com/matryer/is"
	"testing"
)

func TestGroupModel_Create(t *testing.T) {
	isIntegrationTest(t)
	db, teardown := getTestDb(t)
	defer teardown()

	model := GroupModel{
		db: db,
	}

	t.Run("create group without description", func(t *testing.T) {
		is := isPkg.New(t)
		err := model.Create("testgroup", "")
		is.NoErr(err)

		want := models.Group{
			ID:          1,
			Name:        "testgroup",
			Description: "",
		}

		var got models.Group

		row := db.QueryRow("SELECT id, name, description FROM `account_groups` WHERE id = ?", want.ID)
		err = row.Scan(&got.ID, &got.Name, &got.Description)
		is.NoErr(err)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("create group with description", func(t *testing.T) {
		is := isPkg.New(t)

		want := models.Group{
			ID:          2,
			Name:        "testgroup",
			Description: "with description",
		}
		err := model.Create(want.Name, want.Description)
		is.NoErr(err)

		var got models.Group

		row := db.QueryRow("SELECT id, name, description FROM `account_groups` WHERE id = ?", want.ID)
		err = row.Scan(&got.ID, &got.Name, &got.Description)
		is.NoErr(err)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestGroupModel_Read(t *testing.T) {
	isIntegrationTest(t)
	is := isPkg.New(t)
	tests := []struct {
		name        string
		want        *models.Group
		insertGroup bool
		wantErr     bool
		expectedErr error
	}{
		{
			name: "insert group with no description",
			want: &models.Group{
				ID:   1,
				Name: "testgroup",
			},
			insertGroup: true,
		},
		{
			name: "insert group with description",
			want: &models.Group{
				ID:          2,
				Name:        "testgroup",
				Description: "with description",
			},
			insertGroup: true,
		},
		{
			name: "if group id does not exist, return models.ErrNotFound",
			want: &models.Group{
				ID: 3,
			},
			insertGroup: false,
			wantErr:     true,
			expectedErr: models.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, teardown := getTestDb(t)
			defer teardown()

			model := GroupModel{
				db: db,
			}

			is := is.New(t)
			if tt.insertGroup {
				var desc sql.NullString
				if tt.want.Description != "" {
					err := desc.Scan(tt.want.Description)
					if err != nil {
						t.Fatal(err)
					}
				}
				_, err := db.Exec("INSERT INTO `account_groups` (id, name, description) VALUES (?,?,?)", tt.want.ID, tt.want.Name, desc)
				is.NoErr(err)
			}

			got, err := model.Read(tt.want.ID)

			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("got err %v but expected %v", err, tt.expectedErr)
				}
				return
			}
			is.NoErr(err)

			if *got != *tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestGroupModel_Update(t *testing.T) {
	isIntegrationTest(t)
	db, teardown := getTestDb(t)
	defer teardown()

	model := GroupModel{
		db: db,
	}

	tests := []struct {
		name        string
		insert      models.Group
		want        models.Group
		wantErr     bool
		expectedErr error
	}{
		{
			name: "insert group with no description, add description and return group",
			insert: models.Group{
				Name: "testgroup",
			},
			want: models.Group{
				ID:          1,
				Name:        "testgroup",
				Description: "test description",
			},
		},
		{
			name: "insert group and change name",
			insert: models.Group{
				Name:        "testgroup",
				Description: "non empty",
			},
			want: models.Group{
				ID:          2,
				Name:        "test",
				Description: "non empty",
			},
		},
		{
			name:        "empty group will not be updated",
			want:        models.Group{},
			wantErr:     true,
			expectedErr: models.ErrModelNotSaved,
		},
		{
			name: "group with non existend id returns models.ErrNotFound",
			want: models.Group{
				ID: 12,
			},
			wantErr:     true,
			expectedErr: models.ErrNotFound,
		},
		{
			name: "group without a id will not be updated",
			want: models.Group{
				Name:        "testgroup",
				Description: "test description",
			},
			wantErr:     true,
			expectedErr: models.ErrModelNotSaved,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := isPkg.New(t)
			if tt.insert.Name != "" {
				_, err := db.Exec("INSERT INTO `account_groups` (name, description) VALUES (?,?)", tt.insert.Name, tt.insert.Description)
				is.NoErr(err)
			}

			got, err := model.Update(tt.want)

			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("got err %v but expected %v", err, tt.expectedErr)
				}
				return
			}

			is.NoErr(err)

			if *got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestGroupModel_Delete(t *testing.T) {
	isIntegrationTest(t)
	t.Run("empty group delete", func(t *testing.T) {
		is := isPkg.New(t)
		db, teardown := getTestDb(t)
		defer teardown()

		model := GroupModel{
			db: db,
		}
		res, err := db.Exec("INSERT INTO `account_groups` (name) VALUES (?)", "test")
		is.NoErr(err)

		groupId, err := res.LastInsertId()
		is.NoErr(err)

		err = model.Delete(int(groupId))
		is.NoErr(err)

		var groupName string
		err = db.QueryRow("SELECT name from `account_groups` WHERE id=?", int(groupId)).Scan(&groupName)
		if err == nil {
			t.Errorf("wanted err, got none")
		}
		if err != sql.ErrNoRows {
			t.Errorf("got %v but wanted %v err", sql.ErrNoRows, err)
		}
	})
	t.Run("trying to delete nonempty group, return err", func(t *testing.T) {
		is := isPkg.New(t)
		db, teardown := getTestDb(t)
		defer teardown()

		model := GroupModel{
			db: db,
		}
		res, err := db.Exec("INSERT INTO `account_groups` (name) VALUES (?)", "test")
		is.NoErr(err)

		groupId, err := res.LastInsertId()
		is.NoErr(err)

		_, err = db.Exec("INSERT INTO `accounts` (name, group_id) VALUES (?,?)", "test", int(groupId))
		is.NoErr(err)

		err = model.Delete(int(groupId))
		if err == nil {
			t.Errorf("expected error, got none")
		}
		if err != models.ErrNonEmptyDelete {
			t.Errorf("got %v, wanted %v", err, models.ErrNonEmptyDelete)
		}
	})
}
