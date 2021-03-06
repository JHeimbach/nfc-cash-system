package handlers

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jheimbach/nfc-cash-system/api"
	"github.com/jheimbach/nfc-cash-system/pkg/server/repositories"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type accountserver struct {
	storage  repositories.AccountStorager
	tStorage repositories.TransactionStorager // only used to delete accounts
}

func RegisterAccountServer(s *grpc.Server, storage repositories.AccountStorager, tStorage repositories.TransactionStorager) {
	api.RegisterAccountServiceServer(s, &accountserver{storage: storage, tStorage: tStorage})
}

func (a *accountserver) ListAccounts(ctx context.Context, req *api.ListAccountsRequest) (*api.ListAccountsResponse, error) {
	var limit int32 = 0
	var offset int32 = 0

	if req.Paging != nil {
		limit = req.Paging.Limit
		offset = req.Paging.Offset
	}
	accounts, totalCount, err := a.storage.GetAll(ctx, req.GroupId, limit, offset)

	if err != nil {
		return nil, ErrGetAll
	}

	return &api.ListAccountsResponse{
		Accounts:   accounts,
		TotalCount: int32(totalCount),
	}, nil
}

func (a *accountserver) CreateAccount(ctx context.Context, req *api.CreateAccountRequest) (*api.Account, error) {
	account, err := a.storage.Create(ctx, req.Name, req.Description, req.Saldo, req.GroupId, req.NfcChipId)
	if err != nil {
		if err == repositories.ErrDuplicateNfcChipId {
			return nil, status.Error(codes.AlreadyExists, "nfc chip is already in use")
		}
		if err == repositories.ErrGroupNotFound {
			return nil, ErrGroupNotFound
		}
		return nil, ErrCouldNotCreateAccount
	}

	return account, nil
}

func (a *accountserver) GetAccount(ctx context.Context, req *api.GetAccountRequest) (*api.Account, error) {
	account, err := a.storage.Read(ctx, req.Id)

	if err != nil {
		return nil, ErrAccountNotFound
	}

	return account, nil
}

func (a *accountserver) UpdateAccount(ctx context.Context, req *api.Account) (*api.Account, error) {
	acc, err := a.storage.Update(ctx, req)

	if err != nil {
		if err == repositories.ErrUpdateSaldo {
			return nil, status.Error(codes.PermissionDenied, "can not update account saldo trough update")
		}
		if err == repositories.ErrGroupNotFound {
			return nil, ErrGroupNotFound
		}
		return nil, ErrSomethingWentWrong
	}

	return acc, nil
}

func (a *accountserver) DeleteAccount(ctx context.Context, req *api.DeleteAccountRequest) (*empty.Empty, error) {
	err := a.tStorage.DeleteAllByAccount(ctx, req.Id)
	if err != nil {
		return &empty.Empty{}, status.Errorf(codes.Internal, "could not delete transactions for account %d", req.Id)
	}

	err = a.storage.Delete(ctx, req.Id)

	if err != nil {
		return &empty.Empty{}, status.Errorf(codes.Internal, "could not delete account %d", req.Id)
	}

	return &empty.Empty{}, nil
}
