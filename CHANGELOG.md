# Changelog

All notable changes to this project will be documented in this file.

- [1.3.0 (2025-12-31)](#130-2025-12-31)
- [1.2.0 (2025-03-05)](#120-2025-03-05)
- [1.1.1 (2022-07-31)](#111-2022-07-31)
- [1.1.0 (2022-07-31)](#110-2022-07-31)
- [1.0.0 (2022-07-29)](#100-2022-07-29)
- [0.1.1 (2019-04-05)](#011-2019-04-05)
- [0.1.0 (2016-07-10)](#010-2016-07-10)

---

<a name="1.3.0"></a>
## [1.3.0](https://github.com/aisbergg/go-bruh/compare/v1.2.0...v1.3.0) (2025-12-31)

### Features

- update transliteration tables from github.com/avian2/unidecode

### Performance Improvements

- improve performance of unidecode writer


<a name="1.2.0"></a>
## [1.2.0](https://github.com/aisbergg/go-bruh/compare/v1.1.1...v1.2.0) (2025-03-05)

### Bug Fixes

- include changes from python-unidecode

### Documentation

- fix readme-top link
- update documentation

### Features

- update benchmarks
- new public API; performance improvements


<a name="1.1.1"></a>
## [1.1.1](https://github.com/aisbergg/go-bruh/compare/v1.1.0...v1.1.1) (2022-07-31)


<a name="1.1.0"></a>
## [1.1.0](https://github.com/aisbergg/go-bruh/compare/v1.0.0...v1.1.0) (2022-07-31)

### Features

- use empty string as unsuccessful transliteration

### Performance Improvements

- prevent allocating memory when converting rune to string


<a name="1.0.0"></a>
## [1.0.0](https://github.com/aisbergg/go-bruh/compare/v0.1.1...v1.0.0) (2022-07-29)

### Code Refactoring

- rename module and move source code into dedicated dir

### Documentation

- update documentation

### Features

- add configurable error handling
- put converted files directly into pkg dir
- update tables with latest changes from avian2/unidecode
- add tool to convert avian2/unidecode to Go


<a name="0.1.1"></a>
## [0.1.1](https://github.com/aisbergg/go-bruh/compare/v0.1.0...v0.1.1) (2019-04-05)


<a name="0.1.0"></a>
## [0.1.0]() (2016-07-10)

Initial Release
