terraform {
  required_providers {
    ftd = {
      version = "~>0.2"
      source  = "github.com/mr-olenoid/ftd"
    }

  }
}

provider "ftd" {
    username = "admin"
    password = "Cisco_1234"
    url = "https://10.100.16.210"
}

resource "ftd_security_zone" "ft_sz" {
  name = "tf_zone_22"
  type = "securityzone"
  description = "Hello from terraform (edited)"
  mode = "ROUTED"
}

resource "ftd_security_zone" "ft_sz_outside" {
  name = "imported"
  mode = "ROUTED"
  /*
  interfaces {
    id = ftd_interface.outside.id
    name = ftd_interface.outside.name
    type = ftd_interface.outside.type
  }
  */
}

resource "ftd_network_object" "tf_ip_address" {
  name = "ip_from_terraform"
  subtype = "HOST"
  value = "10.0.0.1"
  type = "networkobject"
}

resource "ftd_security_zone" "ft_sz_default_outside" {
  name = "outside_zone"
  mode = "ROUTED"
  
  interfaces {
    id = ftd_interface.outside.id
    name = ftd_interface.outside.name
    type = ftd_interface.outside.type
  }
  
}


resource "ftd_interface" "outside" {
  name = "outside"
  mode = "ROUTED"
  monitorinterface = true
  description = "Hello from terraform"
  ctsenabled = true
  gigabitinterface = true
  present = true
  ipv4 {
      addressnull = false
      defaultrouteusingdhcp = true
      dhcproutemetric = 0
      iptype = "STATIC"
      type = "interfaceipv4"
      ipaddress {
        ipaddress = "192.168.33.11"
        netmask = "255.255.255.0"
      }
    }
}

resource "ftd_access_rule" "tf_test_rule"{
  accesspolicyid = "c78e66bc-cb57-43fe-bcbf-96b79b3475b3"
  name = "tf_test_rule"
  ruleaction = "PERMIT"
  eventlogaction = "LOG_BOTH"
  sourcezones  {
    name = ftd_security_zone.ft_sz_default_outside.name
    id = ftd_security_zone.ft_sz_default_outside.id
  }
  sourcezones {
    name = ftd_security_zone.ft_sz_outside.name
    id = ftd_security_zone.ft_sz_outside.id
  }
  
  destinationnetworks {
    name = ftd_network_object.tf_ip_address.name
    id = ftd_network_object.tf_ip_address.id
    type = ftd_network_object.tf_ip_address.type
  }
  
}