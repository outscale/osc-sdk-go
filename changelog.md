# v1.5.1

- fixed missing files

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
