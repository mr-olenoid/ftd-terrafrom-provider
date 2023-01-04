resource "ftd_application_filter" "remote" {
  name = "remote_connections"
  applications {
    name = data.ftd_application.rdp.name
    id = data.ftd_application.rdp.id
    type = data.ftd_application.rdp.type
  }
  applications {
    name = data.ftd_application.ssh.name
    id = data.ftd_application.ssh.id
    type = data.ftd_application.ssh.type
  }

}