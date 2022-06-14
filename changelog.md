# 1.17.0

 - SDK update for Outscale API v1.20.0

# 1.16.0

 - SDK update for Outscale API v1.20.0

# 1.15.0

 - SDK update for Outscale API v1.19.0

# 1.14.0

 - SDK update for Outscale API v1.18.0

# 1.13.0

 - SDK update for Outscale API v1.17.0

# 1.12.0

- SDK update for Outscale API v1.16.0

# 1.11.2

 - Add workaround for missing BsuToCreate.DeleteOnVmDeletion value

# 1.11.1

 - Go module fix

# 1.11.0

 - SDK update for Outscale API v1.15.0

# 1.10.0

- SDK update for Outscale API v1.14.4

# 1.9.0

 - SDK update for Outscale API v1.10.3

# 1.8.0

 - SDK update for Outscale API v1.8

# 1.7.0

 - SDK update for Outscale API v1.7.4

# v1.6.0

- Upgraded to API v1.7.2

# v1.5.5

- autobuild: keep workflow errors for buggy builds only

# v1.5.4

- fix Github CI for autobuild-v1

# v1.5.3

- use Github CI to detect osc-api change and create new release

# v1.5.2

- make sure sdk generation can be reproduced in through testing
- add CI integration tests

# v1.5.1

- fixed missing files
- rework examples to prepare CI/CD

# v1.5.0

- Upgraded to API v1.6

# v1.4.0

- Based on Outscale API v1.4
- Fixed SDK package versioning
- Warning: SDK versioning may now differ from Outscale API versioning

# v1.3.0

- Based on Outscale API v1.3
- AccessKeys: can now have an expiration date
- Flexible GPU support filtering on Generation
- Tag filtering added for Subnets and VPN connections

# v1.2.0

- Based on Outscale API v1.2
- SnapshotSize: fixed yaml type to avoid size limitation (as it is expressed in bytes)
- CreateAccessKey: has it's own object description
- New filters on ressources:
  - InternetService: LinkNetIds, LinkNetIds
  - Net: IsDefault
  - RouteTable: LinkRouteTableIds, LinkRouteTableMain, RouteDestinationServiceIds
  - SecurityGroup: NetIds
  - FlexibleGpuCatalog: Generations
- Documentation fixes

# v1.1.0

- Based on Outscale API v1.1
- New region: cn-southeast-1

# v1.0.0

- first stable version of Outscale API !
- Based on Outscale API v1.0
- breaking change if you come from v0.0.9 version
- more examples
- does not use a custom version of openapi-generator
- major new API calls support:
  - AccessKey actions
  - CheckAuthentication
  - Account actions
  - ReadConsumptionAccount
  - ListenerRule actions
  - ReadRegions
  - SecretAccessKey actions
  - ServerCertificate actions

# v0.15 to v0.17

- same API features
- ... but breaking change with previous version
- SDK generated with openapi-generator
- add examples


# v0.0.1 to v0.0.9

- first (beta) versions of Outscale API
