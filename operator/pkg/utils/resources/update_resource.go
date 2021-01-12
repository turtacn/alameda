package resources

import (
	"context"

	autuscaling "github.com/turtacn/alameda/operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// UpdateResource define resource update functions
type UpdateResource struct {
	client.Client
}

// NewUpdateResource return UpdateResource instance
func NewUpdateResource(client client.Client) *UpdateResource {
	return &UpdateResource{
		client,
	}
}

// UpdateAlamedaScaler updates AlamedaScaler
func (updateResource *UpdateResource) UpdateAlamedaScaler(alamedaScaler *autuscaling.AlamedaScaler) error {
	err := updateResource.updateResource(alamedaScaler)
	return err
}

// UpdateResource updates resource
func (updateResource *UpdateResource) UpdateResource(resource runtime.Object) error {
	err := updateResource.updateResource(resource)
	return err
}

func (updateResource *UpdateResource) updateResource(resource runtime.Object) error {
	if err := updateResource.Update(context.TODO(),
		resource); err != nil {
		scope.Debug(err.Error())
		return err
	}
	return nil
}
