terraform {
  required_providers {
    junos-vsrx = {
      source = "juniper/providers/junos-vsrx"
      version = "19.41.101"
    }
  }
}

resource "junos-vsrx_InterfacesInterfaceDescription" "vsrx_1" {
    resource_name = "vsrx_1"
    name = "ge-0/0/0"
    description = "Test description"
}

resource "junos-vsrx_InterfacesInterfaceUnitFamilyInetAddressName" "vsrx_2" {
    resource_name = "vsrx_2"
    name = "ge-0/0/0"
    name__1 = "0"
    name__2 = "10.0.0.2/24"
}
