# Changelog

## [ 2024/09/05 - v1.2.2]

### Add/Updated
* Updated go modules
* new metadata prefix

### Deprecated
* All flags have been deprecated except a select few. The config file should be used from now on.

## [ 2024/04/26 - v1.2.1 ]

### Add/Updated
* Added ability to override default contained version
* Updated go modules

## [ 2024/03/15 - v1.2.0 ]

### BREAKING CHANGES

### Remove
* Baski server has been removed - it was extra work that was sporadically used. It's up to the user to provide a way to
  interact with the results. If you wish to keep server, then look at a previous release and fork that code by all means.

## [ 2024/03/11 - v1.1.1 ]

### Changed/Added

* Setting images to private by default and will only set to public once a scan has passed successfully - this can still
  be overridden at the config level.
  
## [ 2024/02/29 - v1.1.0 ]

### BREAKING CHANGES

* changed `cloud` prefix to `infra` in flags and config.
* changed `build.nvidia` prefix to `build.gpu` in flags and config.
* changed `build.nvidia.enable-nvidia-support` prefix to `build.gpu.enable-gpu-support` in flags and config.

### Changed/Added

* Added KubeVirt as a build option.
* Supports AMD GPUs

## [ 2024/02/15 - v1.0.0 ]

### Changed/Added

First release with:

* Functioning support for OpenStack build, scan and signing.
* Baski Server 
