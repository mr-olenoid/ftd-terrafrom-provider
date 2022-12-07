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
}

resource "ftd_network_object" "tf_ip_address" {
  name = "ip_from_terraform"
  subtype = "HOST"
  value = "10.0.0.1"
  type = "networkobject"
}
/*
resource "ftd_security_zone" "ft_sz_default_outside" {
  name = "outside_zone"
  mode = "ROUTED"
}
*/

resource "ftd_interface" "outside" {
  name = "outside"
  mode = "ROUTED"
  type = "physicalinterface"
  monitorinterface = true
  description = "Hello from"
  ctsenabled = true
  gigabitinterface = true
  present = true
  ipv4 {
      addressnull = false
      defaultrouteusingdhcp = true
      dhcp = true
      dhcproutemetric = 0
      iptype = "DHCP"
      type = "interfaceipv4"
      ipaddress {
          ipaddress = "10.100.16.181"
          netmask = "255.255.255.0"
          type = "haipv4address"
        }
    }
}