
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName struct {
	XMLName xml.Name `xml:"configuration"`
	V_security__zone  struct {
		XMLName xml.Name `xml:"security-zone"`
		V_name  string  `xml:"name"`
		V_interfaces  struct {
			XMLName xml.Name `xml:"interfaces"`
			V_name__1  string  `xml:"name"`
			V_system__services  struct {
				XMLName xml.Name `xml:"system-services"`
				V_name__2  string  `xml:"name"`
			} `xml:"host-inbound-traffic>system-services"`
		} `xml:"interfaces"`
	} `xml:"security>zones>security-zone"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameCreate(d *schema.ResourceData, m interface{}) error {

    pcfg := m.(*ProviderConfig)
    client, err := pcfg.Client()
    check(err)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)
	commit := true

	config := xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName{}
	config.V_security__zone.V_name = V_name
	config.V_security__zone.V_interfaces.V_name__1 = V_name__1
	config.V_security__zone.V_interfaces.V_system__services.V_name__2 = V_name__2

    err = client.SendTransaction("", config, commit)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", pcfg.Cfg.Host, id))
    
    err = client.Close()
    check(err)
    
	return junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead(d,m)
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead(d *schema.ResourceData, m interface{}) error {

    pcfg := m.(*ProviderConfig)
    client, err := pcfg.Client()
    check(err)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.V_security__zone.V_name)
	d.Set("name__1", config.V_security__zone.V_interfaces.V_name__1)
	d.Set("name__2", config.V_security__zone.V_interfaces.V_system__services.V_name__2)

    err = client.Close()
    check(err)
    
	return nil
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameUpdate(d *schema.ResourceData, m interface{}) error {

    pcfg := m.(*ProviderConfig)
    client, err := pcfg.Client()
    check(err)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)
	commit := true

	config := xmlSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName{}
	config.V_security__zone.V_name = V_name
	config.V_security__zone.V_interfaces.V_name__1 = V_name__1
	config.V_security__zone.V_interfaces.V_system__services.V_name__2 = V_name__2

    err = client.SendTransaction(id, config, commit)
    check(err)
    
    err = client.Close()
    check(err)
    
	return junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead(d,m)
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameDelete(d *schema.ResourceData, m interface{}) error {

    pcfg := m.(*ProviderConfig)
    client, err := pcfg.Client()
    check(err)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfig(id)
    check(err)

    d.SetId("")
    
    err = client.Close()
    check(err)
    
	return nil
}

func junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesName() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameCreate,
		Read: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameRead,
		Update: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameUpdate,
		Delete: junosSecurityZonesSecurity__ZoneInterfacesHost__Inbound__TrafficSystem__ServicesNameDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.V_security__zone",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.V_security__zone.V_interfaces",
			},
			"name__2": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.V_security__zone.V_interfaces.V_system__services. ",
			},
		},
	}
}