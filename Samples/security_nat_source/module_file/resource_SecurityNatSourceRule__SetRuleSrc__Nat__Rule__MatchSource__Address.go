
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__Address struct {
	XMLName xml.Name `xml:"configuration"`
	V_rule__set  struct {
		XMLName xml.Name `xml:"rule-set"`
		V_name  string  `xml:"name"`
		V_rule  struct {
			XMLName xml.Name `xml:"rule"`
			V_name__1  string  `xml:"name"`
			V_src__nat__rule__match  struct {
				XMLName xml.Name `xml:"src-nat-rule-match"`
				V_source__address  string  `xml:"source-address"`
			} `xml:"src-nat-rule-match"`
		} `xml:"rule"`
	} `xml:"security>nat>source>rule-set"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressCreate(d *schema.ResourceData, m interface{}) error {

    pcfg := m.(*ProviderConfig)
    client, err := pcfg.Client()
    check(err)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_source__address := d.Get("source__address").(string)
	commit := true

	config := xmlSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__Address{}
	config.V_rule__set.V_name = V_name
	config.V_rule__set.V_rule.V_name__1 = V_name__1
	config.V_rule__set.V_rule.V_src__nat__rule__match.V_source__address = V_source__address

    err = client.SendTransaction("", config, commit)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", pcfg.Cfg.Host, id))
    
    err = client.Close()
    check(err)
    
	return junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressRead(d,m)
}

func junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressRead(d *schema.ResourceData, m interface{}) error {

    pcfg := m.(*ProviderConfig)
    client, err := pcfg.Client()
    check(err)

    id := d.Get("resource_name").(string)
    
	config := &xmlSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__Address{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.V_rule__set.V_name)
	d.Set("name__1", config.V_rule__set.V_rule.V_name__1)
	d.Set("source__address", config.V_rule__set.V_rule.V_src__nat__rule__match.V_source__address)

    err = client.Close()
    check(err)
    
	return nil
}

func junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressUpdate(d *schema.ResourceData, m interface{}) error {

    pcfg := m.(*ProviderConfig)
    client, err := pcfg.Client()
    check(err)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_name__1 := d.Get("name__1").(string)
	V_source__address := d.Get("source__address").(string)
	commit := true

	config := xmlSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__Address{}
	config.V_rule__set.V_name = V_name
	config.V_rule__set.V_rule.V_name__1 = V_name__1
	config.V_rule__set.V_rule.V_src__nat__rule__match.V_source__address = V_source__address

    err = client.SendTransaction(id, config, commit)
    check(err)
    
    err = client.Close()
    check(err)
    
	return junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressRead(d,m)
}

func junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressDelete(d *schema.ResourceData, m interface{}) error {

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

func junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__Address() *schema.Resource {
	return &schema.Resource{
		Create: junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressCreate,
		Read: junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressRead,
		Update: junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressUpdate,
		Delete: junosSecurityNatSourceRule__SetRuleSrc__Nat__Rule__MatchSource__AddressDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.V_rule__set",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.V_rule__set.V_rule",
			},
			"source__address": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.V_rule__set.V_rule.V_src__nat__rule__match. Source address",
			},
		},
	}
}