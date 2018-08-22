// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package cli

import (
	context "context"
	build "github.com/windmilleng/tilt/internal/build"
	engine "github.com/windmilleng/tilt/internal/engine"
	k8s "github.com/windmilleng/tilt/internal/k8s"
	model "github.com/windmilleng/tilt/internal/model"
	service "github.com/windmilleng/tilt/internal/service"
	dirs "github.com/windmilleng/wmclient/pkg/dirs"
)

// Injectors from wire.go:

func wireServiceCreator(ctx context.Context) (model.ServiceCreator, error) {
	client, err := build.DefaultDockerClient()
	if err != nil {
		return nil, err
	}
	windmillDir, err := dirs.UseWindmillDir()
	if err != nil {
		return nil, err
	}
	manager := service.ProvideMemoryManager()
	env, err := k8s.DetectEnv()
	if err != nil {
		return nil, err
	}
	buildAndDeployer, err := engine.NewLocalBuildAndDeployer(ctx, client, windmillDir, manager, env)
	if err != nil {
		return nil, err
	}
	upper, err := engine.NewUpper(ctx, buildAndDeployer)
	if err != nil {
		return nil, err
	}
	serviceCreator := provideServiceCreator(upper, manager)
	return serviceCreator, nil
}