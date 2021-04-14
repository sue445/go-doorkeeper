## Unreleased
[full changelog](http://github.com/sue445/go-doorkeeper/compare/v0.1.7...master)

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
