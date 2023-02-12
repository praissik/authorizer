package api

import (
	"authorizer/pkg/account"
	errors "authorizer/pkg/error"
	"authorizer/pkg/proto/pb"
	"authorizer/test"
	"fmt"
	"net/http"
	"testing"
)

const (
	ID              = "63e8de6816129322bdc04c63"
	email           = "new@email.com"
	busyEmail       = "busy@email.com"
	invalidEmail    = "invalidemail.com"
	password        = "P@ssw0rd"
	shortPassword   = "P@ssw0"
	invalidPassword = "P@ssw0 rd"
	correlationID   = "c0rr3l@tion-id"
)

func TestServer_Register(t *testing.T) {
	e, err := test.GetEngine(t, NewServer())
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	//objectID, err := primitive.ObjectIDFromHex(ID)
	//if err != nil {
	//	t.Errorf(err.Error())
	//	return
	//}

	accountEntity := &account.Entity{
		//ID:       objectID,
		Email:    busyEmail,
		Password: password,
	}

	_, err = accountEntity.Create()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	tests := []struct {
		name    string
		in      *pb.AuthRequest
		want    *pb.AuthReply
		wantErr bool
	}{
		{
			name: "TestOK",
			in: &pb.AuthRequest{
				CorrelationID: correlationID,
				Email:         email,
				Password:      password,
			},
			want: &pb.AuthReply{
				Status:  http.StatusOK,
				Message: "",
			},
			wantErr: false,
		},
		{
			name: "TestEmptyEmail",
			in: &pb.AuthRequest{
				CorrelationID: correlationID,
				Password:      password,
			},
			want: &pb.AuthReply{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf(errors.EmptyEmail),
			},
			wantErr: false,
		},
		{
			name: "TestEmptyPassword",
			in: &pb.AuthRequest{
				CorrelationID: correlationID,
				Email:         email,
			},
			want: &pb.AuthReply{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf(errors.EmptyPassword),
			},
			wantErr: false,
		},
		{
			name: "TestBusyEmail",
			in: &pb.AuthRequest{
				CorrelationID: correlationID,
				Email:         busyEmail,
				Password:      password,
			},
			want: &pb.AuthReply{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf(errors.BusyEmail, busyEmail),
			},
			wantErr: false,
		},
		{
			name: "TestInvalidEmail",
			in: &pb.AuthRequest{
				CorrelationID: correlationID,
				Email:         invalidEmail,
				Password:      password,
			},
			want: &pb.AuthReply{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf(errors.InvalidEmail, invalidEmail),
			},
			wantErr: false,
		},
		{
			name: "TestInvalidPassword",
			in: &pb.AuthRequest{
				CorrelationID: correlationID,
				Email:         email,
				Password:      invalidPassword,
			},
			want: &pb.AuthReply{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf(errors.InvalidPassword),
			},
			wantErr: false,
		},
		{
			name: "TestToShortPassword",
			in: &pb.AuthRequest{
				CorrelationID: correlationID,
				Email:         email,
				Password:      shortPassword,
			},
			want: &pb.AuthReply{
				Status:  http.StatusBadRequest,
				Message: fmt.Sprintf(errors.ToShortPassword),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := e.Client.Register(e.Ctx, tt.in)
			if err != nil != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Status == http.StatusOK && got.Status == tt.want.Status {
				if len(got.Message) == 0 {
					t.Errorf("Register() got = %v, want %v", got, tt.want)
					return
				}
			} else if got.Message != tt.want.Message ||
				got.Status != tt.want.Status {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
