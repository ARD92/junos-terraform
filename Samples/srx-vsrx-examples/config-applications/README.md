# Configuring Applications
The provider generated in this path, would be associated for configuring the below 

```
set applications application <NAME> protocol <PROTOCOL> source-port <SOURCE_PORT> destination-port <DESTINATION_PORT>
```
## Compiling the provider 
In order to compile, generate the yin and xml files from the yang file using cmd/processYang as below. The 
config.toml file is present in the same directory. 
 
```
./processYang -config /var/tmp/srx-vsrx-examples/config-applications/config.toml
```
Post generating the xpaths, choose the exact xpath for which the application hierarchy can be compiled. In this resource, we are specifically choosing protocol, source-port and destination port 

Compile the provider using cmd/processProviders 

```
./processProviders -config /var/tmp/srx-vsrx-examples/config-applications/config.toml
``` 

Generate the binary using the below once the terraform APIs are generated.
```
go build -o terraform-junos-device
```

## Sample example 
The example file is written in main.tf to program applications 

```
resource "junos-qfx_ApplicationsApplicationProtocol" "appprotocol" {
    resource_name = "RESOURCE_NAME"
    name = "APPLICATION_NAME"
    protocol = "tcp"
}

# configure application source port
resource "junos-qfx_ApplicationsApplicationSource__Port" "appsp" {
    resource_name = "RESOURCE_NAME"
    name = "APPLICATION_NAME"
    source__port = "270"
}

# configure application destination port
resource "junos-qfx_ApplicationsApplicationDestination__Port" "appdp" {
    resource_name = "RESOURCE_NAME"
    name = "APPLICATION_NAME"
    destination__port = "270"
}
```
## Loading the Junos provider

Create a directory in the same working directory where your main.tf file is placed as below 

```
mkdir -p terraform.d/plugins/<HOST_NAME>/<NAMESPACE_NAME>/<PROVIDER_NAME>/<VERSION>/linux_amd64
```
Copy the generated binary into the above created directory. Here below are the variables

In the below example
VERSION = 1.0.3
HOST_NAME = 10.85.47.161
NAMESPACE_NAME = ns
PROVIDER_NAME = junos-qfx 

## Init procedure 

```
terraform init 
```

In case a newer version is present you can do the below

```
terraform init --upgrade
```

## Plan and apply accordingly

```
terraform plan

terraform apply
```

