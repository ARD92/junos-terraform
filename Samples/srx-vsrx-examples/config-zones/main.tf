# Below is required to ensure custom provider is used
terraform {
  required_providers {
    junos-device = {
      source = "aprabh-mbp/ns/junos-device"
      version = "1.0.0"
    }
  }
}
 
# Make connection to the junos device to configure using the resources used
provider "junos-device" {
    host = "52.170.83.35"
    port = 830
    username = "aprabh"
    password = "Juniper@1234567"
    sshkey = ""
}
 
resource "junos-device_SecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficProtocolsName" "seczonesprotocol" {
    name = "zone-test"
    name__1 = "ge-0/0/0.0"
    name__2 = "all"
}

resource "junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName" "seczonesservices" {
    name = "zone-test"
    name__1 = "ge-0/0/0.0"
    name__2 = "all"
}
 
resource "junos-device_commit" "commit" {
    resource_name = "commit"
    depends_on = [
        junos-device_SecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficProtocolsName.seczonesprotocol,
        junos-device_SecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName.seczonesservices
    ]
}
