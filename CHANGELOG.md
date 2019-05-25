# Change Log

## 2.1.0 - 2019-05-25

### Added

- A new data source to convert `yaml` to `json` - `yaml_to_json` - that is meant to be used with `jsondecode` in Terraform 0.12

## 2.0.2 - 2019-04-20

### Fixed

- Error on non-string values in `yaml_list_of_strings` - serialize them to flow-style YAML

## 2.0.1 - 2019-04-12

### Fixed

- Crash on empty values in `yaml_map_of_strings` - convert them to empty strings instead

## 2.0.0 - 2018-06-28

### Added

- A new data source to de-serialize a list of strings - `yaml_list_of_strings`

### Changed

- `yaml` renamed to `yaml_map_of_strings`


## 1.0.0 - 2018-06-20

### Added

- Initial implementation for `yaml` data source
