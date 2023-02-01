resource "ftd_access_rule" "tf_first_rule" {
  accesspolicyid = ftd_access_policy.defaul_access_rule.id
  name = "tf_test_1"
  ruleaction = "PERMIT"
  eventlogaction = "LOG_BOTH"

  ruleposition = -1

  sourcezones {
    name = ftd_security_zone.ft_sz.name
    id = ftd_security_zone.ft_sz.id
  }

}

resource "ftd_access_rule" "tf_test_rule" {
  accesspolicyid = ftd_access_policy.defaul_access_rule.id
  name = "tf_test_rule"
  ruleaction = "PERMIT"
  eventlogaction = "LOG_BOTH"

  sourcezones {
    name = ftd_security_zone.ft_sz_outside.name
    id = ftd_security_zone.ft_sz_outside.id
  }

  sourcezones {
    name = ftd_security_zone.ft_sz.name
    id = ftd_security_zone.ft_sz.id
  }

  sourceports {
    name = data.ftd_tcp_udp_port.ssh.name
    id = data.ftd_tcp_udp_port.ssh.id
    type = data.ftd_tcp_udp_port.ssh.type
  }
  sourceports {
    name = data.ftd_tcp_udp_port.imap.name
    id = data.ftd_tcp_udp_port.imap.id
    type = data.ftd_tcp_udp_port.imap.type
  }
  
  destinationnetworks {
    name = ftd_network_object.tf_ip_address.name
    id = ftd_network_object.tf_ip_address.id
    type = ftd_network_object.tf_ip_address.type
  }

  embeddedappfilter {
    applications {
      name = data.ftd_application.rdp.name
      id = data.ftd_application.rdp.id
      type = data.ftd_application.rdp.type
    }
    applicationfilters {
      name = ftd_application_filter.remote.name
      id = ftd_application_filter.remote.id
      type = ftd_application_filter.remote.type
    }
    conditions {
      risks {
        risk = "VERY_LOW"
      }
      categories {
        name = data.ftd_application_category.ad.name
        id = data.ftd_application_category.ad.id
        type = data.ftd_application_category.ad.type
      }
    }
  }
  
}