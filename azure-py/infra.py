import pulumi
import pulumi_azure as azure

subnet_network = azure.network.get_subnet(name="clusternodes",
    resource_group_name="npeccpdevwu2RgGolduck",
    virtual_network_name="npeccpdevwu2RgGolduckVPC")

