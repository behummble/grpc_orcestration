package orcestr

import (
	"context"

	orcestrationv1 "github.com/behummble/grpc_orcestration/protos/gen/go/orcestration"
)

type Orcestr interface {
	StartContainer(
		ctx context.Context,
		image string,
		version string,
		volumes string,
		dockerFilePath string,
	) (bool, []string, string)
	StopContainer(
		ctx context.Context,
		container_id string,
	) (bool, []string)
	RestartContainer(
		ctx context.Context,
		container_id string,
	) (bool, []string)
}
