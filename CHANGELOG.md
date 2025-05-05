## Unreleased
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.12...master)

## [v0.1.12](https://github.com/sue445/go-doorkeeper/releases/tag/v0.1.12)
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.11...v0.1.12)

* Requires Go 1.23+
  * https://github.com/sue445/go-doorkeeper/pull/103
* Update dependencies

## [v0.1.11](https://github.com/sue445/go-doorkeeper/releases/tag/v0.1.11)
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.10...v0.1.11)

* Add golangci-lint
  * https://github.com/sue445/go-doorkeeper/pull/92

## v0.1.10
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.9...v0.1.10)

* Drop Go 1.17
  * https://github.com/sue445/go-doorkeeper/pull/91
* Update dependencies

## v0.1.9
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.8...v0.1.9)

* Drop Go 1.16
  * https://github.com/sue445/go-doorkeeper/pull/77
* Migrate to github.com/cockroachdb/errors
  * https://github.com/sue445/go-doorkeeper/pull/80

## v0.1.8
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.7...v0.1.8)

* Drop Go 1.15
  * https://github.com/sue445/go-doorkeeper/pull/76
* Fix deprecation warning
  * https://github.com/sue445/go-doorkeeper/pull/75
* Update dependencies

## v0.1.7
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.6...v0.1.7)

* Wrap all errors with `errors.WithStack`
  * https://github.com/sue445/go-doorkeeper/pull/49

## v0.1.6
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.5...v0.1.6)

* Resolved `unexpected end of JSON input` when `X-Ratelimit` header isn't returned from Doorkeeper API
  * https://github.com/sue445/go-doorkeeper/pull/48
* Migrate to GitHub Actions
  * https://github.com/sue445/go-doorkeeper/pull/19
* Update dependencies

## v0.1.5
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.4...v0.1.5)

* Resolved. json cannot unmarshal number into Go struct field rawEvent.lat of type string
  * https://github.com/sue445/go-doorkeeper/pull/18
* Update dependencies

## v0.1.4
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.3...v0.1.4)

* Changed. returns `SortEnum` function instead of `var`
  * https://github.com/sue445/go-doorkeeper/pull/14

## v0.1.3
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.2...v0.1.3)

* Fixed. `unexpected end of JSON input` when API is failed
  * https://github.com/sue445/go-doorkeeper/pull/12

## v0.1.2
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.1...v0.1.2)

* Bugfix: invalid query string
  * https://github.com/sue445/go-doorkeeper/pull/11

## v0.1.1
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.0...v0.1.1)

* Extract to Sort enum
  * https://github.com/sue445/go-doorkeeper/pull/10

## v0.1.0
* First release
