package keycodes

import (
	Keycodes "github.com/containers-ai/api/datahub/keycodes"
	"github.com/golang/protobuf/ptypes/empty"
	KeycodeMgt "github.com/turtacn/alameda/datahub/pkg/account-mgt/keycodes"
	AlamedaUtils "github.com/turtacn/alameda/pkg/utils"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

func (c *ServiceKeycodes) GenerateRegistrationData(ctx context.Context, in *empty.Empty) (*Keycodes.GenerateRegistrationDataResponse, error) {
	scope.Debug("Request received from GenerateRegistrationData grpc function: " + AlamedaUtils.InterfaceToString(in))

	keycodeMgt := KeycodeMgt.NewKeycodeMgt()

	// Generate registration data
	registrationData, err := keycodeMgt.GetRegistrationData()
	if err != nil {
		scope.Error(err.Error())
		return &Keycodes.GenerateRegistrationDataResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	}

	return &Keycodes.GenerateRegistrationDataResponse{
		Status: &status.Status{
			Code: int32(code.Code_OK),
		},
		Data: registrationData,
	}, nil
}
