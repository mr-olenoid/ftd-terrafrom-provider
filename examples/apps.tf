data "ftd_application" "rdp" {
    name = "rdp"
}

data "ftd_application" "ssh" {
  name = "ssh"
}

data "ftd_application_category" "ad" {
  name = "active directory"
}