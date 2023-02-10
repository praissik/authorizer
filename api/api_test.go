package api

import (
	errors "authorizer/pkg/error"
	"authorizer/pkg/proto/pb"
	"authorizer/test"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestServer_Register(t *testing.T) {
	e := test.GetEngine(t, NewServer())

	e.CreateAccount()

	tests := []struct {
		name    string
		in      *pb.AuthRequest
		want    *pb.AuthReply
		wantErr bool
	}{
		{
			name: "TestOk",
			in: &pb.AuthRequest{
				CorrelationID: uuid.New().String(),
				Email:         e.Email,
				Password:      e.Password,
			},
			want: &pb.AuthReply{
				Status:  200,
				Message: "",
			},
			wantErr: false,
		},
		{
			name: "TestNoData",
			in:   &pb.AuthRequest{},
			want: &pb.AuthReply{
				Status:  400,
				Message: errors.EmptyEmail,
			},
			wantErr: false,
		},
		{
			name: "TestBusyEmail",
			in: &pb.AuthRequest{
				CorrelationID: uuid.New().String(),
				Email:         e.Email,
				Password:      e.Password,
			},
			want: &pb.AuthReply{
				Status:  400,
				Message: fmt.Sprintf(errors.BusyEmail, e.Email),
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
			if got.Message != tt.want.Message ||
				got.Status != tt.want.Status {
				t.Errorf("Register() got = %v, want %v", got, tt.want)
			}
		})
	}
}
