#--------------------------------------#
# Procedure to load custom provider 
# mkdir -p terraform.d/plugins/<HOST_NAME>/<NAMESPACE_NAME>/<PROVIDER_NAME>/<VERSION>/linux_amd64

# copy the generated binary into the above path  
# In the below example 
# VERSION = 1.0.3
# HOST_NAME = 10.85.47.161
# NAMESPACE_NAME = ns
# PROVIDER_NAME = junos-qfx 
#--------------------------------------#


# Below is required to ensure custom provider is used 
terraform {
  required_providers {
    junos-qfx = {
      source = "10.85.47.161/ns/junos-qfx"
      version = "1.0.3"
    }
  }
}

# Make connection to the junos device to configure using the resources used
provider "junos-qfx" {
    host = "10.102.144.2"
    port = 830
    username = "root"
    password = "Juniper123"
    sshkey = ""
}

# configure application protocol 
resource "junos-qfx_ApplicationsApplicationProtocol" "appprotocol" {
    resource_name = "aprabh"
    name = "app1"
    protocol = "tcp"
}

# configure application source port 
resource "junos-qfx_ApplicationsApplicationSource__Port" "appsp" {
    resource_name = "aprabh"
    name = "app1"
    source__port = "270"
}

# configure application destination port 
resource "junos-qfx_ApplicationsApplicationDestination__Port" "appdp" {
    resource_name = "aprabh"
    name = "app1"
    destination__port = "270"
}

# commit the configuration
resource "junos-qfx_commit" "commit" {
    resource_name = "commit"
    depends_on = [
        junos-qfx_ApplicationsApplicationProtocol.appprotocol,
        junos-qfx_ApplicationsApplicationSource__Port.appsp,
        junos-qfx_ApplicationsApplicationDestination__Port.appdp
    ]
}
