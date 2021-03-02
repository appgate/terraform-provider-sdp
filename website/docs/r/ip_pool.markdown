---
layout: "appgate"
page_title: "APPGATE: appgate_ip_pool"
sidebar_current: "docs-appgate-resource-ip_pool"
description: |-
   Create a new IP Pool.
---

# appgate_ip_pool

Create a new IP Pool.

## Example Usage

```hcl

resource "appgate_ip_pool" "example_ip_pool" {
  name            = "ip range example"
  lease_time_days = 5
  ranges {
    first = "10.0.0.1"
    last  = "10.0.0.254"
  }
}

```

## Argument Reference

The following arguments are supported:


* `ip_version6`: (Optional) Whether the IP pool is for v4 or v6.
* `ranges`: (Optional) List of (non-conflicting) IP address ranges to allocate IPs in order.
* `lease_time_days`: (Optional) Number of days Allocated IPs will be reserved for device&users before they are reclaimable by others.
* `total`: (Optional) The total size of the IP Pool.
* `currently_used`: (Optional) Number of IPs in the pool are currently in use by device&users.
* `reserved`: (Optional) Number of IPs in the pool are not currently in use but reserved for device&users according to the "leaseTimeDays" setting.
* `id`: (Required) ID of the object.
* `name`: (Required) Name of the object.
* `notes`: (Optional) Notes for the object. Used for documentation purposes.
* `tags`: (Optional) Array of tags.


### ranges
List of (non-conflicting) IP address ranges to allocate IPs in order.

* `first`:  (Required) The beginning of the IP range.
* `last`:  (Required) The end of the IP range.

### tags
Array of tags.




## Import

Instances can be imported using the `id`, e.g.

```
$ terraform import appgate_ip_pool d3131f83-10d1-4abc-ac0b-7349538e8300
```
