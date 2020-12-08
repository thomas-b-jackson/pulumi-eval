resource "azurerm_virtual_network" "example" {
    name = "example-network"
    address_space = ["10.0.0.0/16"]
    location = "westus2"
    resource_group_name = "npeccpdevwu2RgGolduck"
}

data "azurerm_subnet" "subnet_network" {
  name                   = "clusternodes"
  resource_group_name    = "npeccpdevwu2RgGolduck"
  virtual_network_name   = "npeccpdevwu2RgGolduckVPC"
}

resource "azurerm_network_interface" "example" {
    name = "example-nic"
    location = "westus2"
    resource_group_name = "npeccpdevwu2RgGolduck"

    ip_configuration {
        name = "internal"
        subnet_id = data.azurerm_subnet.subnet_network.id
        private_ip_address_allocation = "Dynamic"
    }
}

resource "azurerm_linux_virtual_machine" "example" {
    name = "example-machine"
    resource_group_name = "npeccpdevwu2RgGolduck"
    location = "westus2"
    size = "Standard_F2"
    admin_username = "adminuser"

    disable_password_authentication = false
    admin_password = "x8NL6w86jeBt"

    network_interface_ids = [
        azurerm_network_interface.example.id,
    ]

    os_disk {
        caching = "ReadWrite"
        storage_account_type = "Standard_LRS"
    }

    source_image_reference {
        publisher = "Canonical"
        offer = "UbuntuServer"
        sku = "16.04-LTS"
        version = "latest"
    }
}
