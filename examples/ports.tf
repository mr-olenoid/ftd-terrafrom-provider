data "ftd_tcp_udp_port" "ssh" {
    name = "SSH"
    type = "tcpportobject"
}

data "ftd_tcp_udp_port" "imap" {
  name = "IMAP"
  type = "tcpportobject"
}

#resource "ftd_tcp_udp_port_user" "custom_port"{
#    name = "8443"
#    port = "8443"
#    type = "udpportobject"
#}