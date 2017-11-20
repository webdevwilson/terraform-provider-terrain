# Terrain - Configuration Manager for Terraform

## Provider

```hcl
provider "terrain" {
  version         = "~> 0.1"
  install_command = "yum install -y %s"
}
```

## Resources

### Install Packages

Install packages using package manager.

```hcl
// a single package
resource "terrain_package" "ruby" {
  name    = "ruby"
  version = "2.3.0"
}

// multiple packages
resource "terrain_package" "pkgs" {
  count   = "${length(var.pkgs)"
  name    = "${element(var.pkgs, count.index)}"
}
```

### Run Commands

Run a command on the machine.

```hcl
resource "terrain_command" "create_dir" {
  command = "mkdir -p /opt/app"
  unless  = "-d /opt/app"
}
```

### Create Files

Copy a file from a remote location.

```hcl
resource "terrain_file" "config" {
  source      = "http://domain.com/config.json"
  destination = "/etc/config"
}
```
