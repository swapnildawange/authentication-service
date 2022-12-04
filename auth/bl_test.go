package auth

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/authentication-service/auth/mocks"
	"github.com/authentication-service/spec"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/mock"
)

func GetTestBL() (BL, *mocks.Repository) {
	logger := log.NewLogfmtLogger(os.Stderr)
	mockRepo := &mocks.Repository{}
	mockBL := NewBL(logger, mockRepo)
	return mockBL, mockRepo
}

func TestBL_Login(t *testing.T) {

	type args struct {
		ctx          context.Context
		loginRequest spec.LoginRequest
	}
	tests := []struct {
		name        string
		args        args
		want        spec.AuthUserResponse
		prepareTest func(args, *mocks.Repository)
		wantErr     bool
	}{
		{
			name: "Positive",
			args: args{
				ctx: context.TODO(),
				loginRequest: spec.LoginRequest{
					Email:    "test@email.com",
					Password: "testPassword",
				},
			},
			prepareTest: func(args args, repo *mocks.Repository) {
				repo.Mock.On("GetUserFromAuth", mock.Anything, mock.Anything).Return(spec.AuthUserResponse{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bl, mockRepo := GetTestBL()
			tt.prepareTest(tt.args, mockRepo)

			got, err := bl.Login(tt.args.ctx, tt.args.loginRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("bl.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bl.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
