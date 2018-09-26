package main

import (
	"fmt"

	"context"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	)

type ConnectionString struct {
	name  string
	IPAddress string
	Ports []string
}

func main(){

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.38"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList( context.Background(), types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.Names)

		inspection, err:= cli.ContainerInspect(ctx,container.ID)

		if err != nil{
			panic(err)
		}

		networksettings :=inspection.NetworkSettings

		fmt.Print(networksettings.IPAddress)

		for port, _ := range networksettings.Ports{
			fmt.Println(port)
		}


		for _, network := range inspection.NetworkSettings.Networks {
			fmt.Println(network.IPAddress)
			fmt.Println(network.Gateway)
			//spew.Dump(network)
		}
	}
}
