# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.8.1] - 2026-10-21

### Fixed

- Fix homebrew installation command

## [0.8.0] - 2026-10-21

### Added

- Add cli installer

## [0.7.0] - 2026-10-20

### Added

- Add Go installer

## [0.6.0] - 2026-10-20

### Added

- Add Node and pnpm installer

## [0.5.0] - 2026-10-19

### Added

- Add font installation

## [0.4.0] - 2026-10-19

### Added

- Add git configuration file
- Add SSH configuration

### Changed

- Separate git configuration into independent functions
- Rename `validateXcode` function
- Use new `userDir` variable

### Fixed

- Fix infinite loop when configuring git

## [0.3.0] - 2026-10-16

### Added

- Add git installation

### Changed

- Refactor error handling

## [0.2.2.] - 2026-10-14

### Fixed

- Fix the installation command in the README

## [0.2.1] - 2026-10-14

### Changed

- Update the executable name

## [0.2.0] - 2026-10-14

### Added

- Add `run` command to Justfile
- Add `.zprofile` configuration file
- Add `homebrew` installation
- Add `downloader` function

### Changed

- Update go to 1.27.2

## [0.1.1] - 2026-10-09

### Added

- Add Justfile
- Add install script for Mac

### Changed

- Update gitignore
- Update README

## [0.1.0] - 2026-10-08

### Added

- Add installer for Xcode

## [0.0.5] - 2026-10-08

### Added

- Add a function to execute the steps
- Add the steps

### Changed

- The responsibility for exiting the app has been separated, and a feature to
  start the full install has been added

## [0.0.4] - 2026-10-07

### Added

- Add input package
- Add interactivity to the menu

### Fixed

- Fix `TimeMessage` for use seconds
- Fix style for the menu

## [0.0.3] - 2026-10-07

### Added

- Add `message` package to display different messages
- `ErrorMessage` is added to `ShowMenu`

## [0.0.2] - 2026-10-07

### Added

- Add package `shell` for executing commands in the terminal
- Add the 'clear' command to the `ShowMenu` function

## [0.0.1] - 2026-09-26

### Added

- Add types and package Menu
- Add types for menu options
- Add menu items for rendering
- Add main menu to the app
