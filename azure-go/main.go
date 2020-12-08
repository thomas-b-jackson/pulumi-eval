package main

import (
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/network"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type infrastructure struct {
	server *compute.LinuxVirtualMachine
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		_, err := createInfrastructure(ctx)

		return err
	})
}
func createInfrastructure(ctx *pulumi.Context) (*infrastructure, error) {

	_, err := network.NewVirtualNetwork(ctx, "exampleVirtualNetwork", &network.VirtualNetworkArgs{
		AddressSpaces: pulumi.StringArray{
			pulumi.String("10.0.0.0/16"),
		},
		Location:          pulumi.String("westus2"),
		ResourceGroupName: pulumi.String("npeccpdevwu2RgGolduck"),
	})
	if err != nil {
		return nil, err
	}

	subnetNetwork, err := network.LookupSubnet(ctx, &network.LookupSubnetArgs{
		Name:               "clusternodes",
		ResourceGroupName:  "npeccpdevwu2RgGolduck",
		VirtualNetworkName: "npeccpdevwu2RgGolduckVPC",
	})
	if err != nil {
		return nil, err
	}
	exampleNetworkInterface, err := network.NewNetworkInterface(ctx, "exampleNetworkInterface", &network.NetworkInterfaceArgs{
		Location:          pulumi.String("westus2"),
		ResourceGroupName: pulumi.String("npeccpdevwu2RgGolduck"),
		IpConfigurations: network.NetworkInterfaceIpConfigurationArray{
			&network.NetworkInterfaceIpConfigurationArgs{
				Name:                       pulumi.String("internal"),
				SubnetId:                   pulumi.String(subnetNetwork.Id),
				PrivateIpAddressAllocation: pulumi.String("Dynamic"),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	linuxMachine, err := compute.NewLinuxVirtualMachine(ctx, "exampleLinuxVirtualMachine", &compute.LinuxVirtualMachineArgs{
		ResourceGroupName:             pulumi.String("npeccpdevwu2RgGolduck"),
		Location:                      pulumi.String("westus2"),
		Size:                          pulumi.String("Standard_F2"),
		AdminUsername:                 pulumi.String("adminuser"),
		DisablePasswordAuthentication: pulumi.Bool(false),
		AdminPassword:                 pulumi.String("x8NL6w86jeBt"),
		NetworkInterfaceIds: pulumi.StringArray{
			exampleNetworkInterface.ID(),
		},
		OsDisk: &compute.LinuxVirtualMachineOsDiskArgs{
			Caching:            pulumi.String("ReadWrite"),
			StorageAccountType: pulumi.String("Standard_LRS"),
		},
		SourceImageReference: &compute.LinuxVirtualMachineSourceImageReferenceArgs{
			Publisher: pulumi.String("Canonical"),
			Offer:     pulumi.String("UbuntuServer"),
			Sku:       pulumi.String("16.04-LTS"),
			Version:   pulumi.String("latest"),
		},
	})
	if err != nil {
		return nil, err
	}

	return &infrastructure{
		server: linuxMachine,
	}, nil
}
