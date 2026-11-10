# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.9.0] - 2026-11-10

### Changed

- Reuse a single standard input scanner and return input errors, including end-of-file
- Handle input errors in the main menu and return normally when exiting the application
- Display invalid main menu selections on the next render without a timed delay
- Remove timed messages and the invalid-selection message from the Git configuration prompt
- Shorten the main menu exit label and update the farewell message

## [2.8.2] - 2026-11-10

### Added

- Add a new alias for cloning Git repositories

## [2.8.1] - 2026-11-09

### Added

- Add a custom dark theme for Lazygit
- Download the Lazygit plugin configuration when applying Neovim configurations

### Changed

- Disable Snacks automatic Lazygit configuration to preserve the YAML theme, including disabling its automatic Neovim editing integration

## [2.8.0] - 2026-11-09

### Added

- Add `lazygit` to the CLI tools installed with Homebrew
- Add Lazygit configuration with interface and Git preferences in `config/lazygit/config.yml`
- Download Lazygit configuration to `~/.config/lazygit/config.yml` when applying configurations
- Export `XDG_CONFIG_HOME` as `$HOME/.config` in the Zsh environment

## [2.7.0] - 2026-11-09

### Added

- Add zoxide to the CLI tools installed with Homebrew
- Add `tools.zsh` to initialize zoxide and source it at the end of `.zshrc`
- Download `tools.zsh` when applying Zsh configurations

## [2.6.2] - 2026-11-09

### Added

- Add Ctrl + G to toggle the Neovim terminal in normal and terminal modes while keeping its session running
- Download the Neovim keymap configuration when applying configurations

## [2.6.1] - 2026-11-09

### Added

- Add Ctrl + L to advance through Neovim snippet placeholders while keeping Tab for completion

## [2.6.0] - 2026-11-09

### Added

- Enable the Go and Markdown LazyVim extras through `lazyvim.json`
- Enable spell checking in Spanish and English in Neovim
- Configure Blink completion to accept suggestions with Tab and keep Enter for new lines
- Configure Tab and Shift-Tab to move between snippet placeholders
- Download the Neovim extras, options, and completion configuration when applying configurations

## [2.5.2] - 2026-11-08

### Added

- Add the Rosé Pine theme with transparency to Neovim
- Configure the Snacks file explorer to open on the right
- Download the Neovim theme and explorer configuration when applying configurations

### Changed

- Remove the previous Neovim configuration before cloning the LazyVim starter
- Remove the starter's `.git` directory after cloning

## [2.5.1] - 2026-11-07

### Changed

- Centralize download URL generation in `GetDownloadURL` and reuse it across configurations and installation steps

## [2.5.0] - 2026-11-06

### Added

- Add `just` to the CLI tools installed by the installer
- Configure Neovim by cloning the LazyVim starter into `~/.config/nvim` when applying configurations

### Changed

- Update the error message displayed when cloning the LazyVim starter fails

## [2.4.0] - 2026-11-06

### Added

- Add Neovim installer

## [2.3.2] - 2026-11-06

### Added

- Add `exports.zsh` for environment variables, downloaded to `~/.config/zsh/exports.zsh` and sourced from `.zshrc`

### Fixed

- Set `EZA_CONFIG_DIR` to `~/.config/eza` so eza loads the installed Rosé Pine theme on macOS

## [2.3.1] - 2026-11-06

### Added

- Add Ghostty configuration file with the Rosé Pine theme
- Download Ghostty configuration to `~/.config/ghostty/config.ghostty` when applying configurations

## [2.3.0] - 2026-11-06

### Added

- Add Ghostty installer

## [2.2.0] - 2026-11-06

### Added

- Add tree-sitter-cli installer

## [2.1.0] - 2026-11-06

### Added

- Add FigCraft installer

## [2.0.2] - 2026-11-05

### Fixed

- Build the download URLs with `baseURL` (`url.Parse` + `JoinPath`) instead of `path.Join`
- Pass the download URL as string to `DownloadFiles`

## [2.0.1] - 2026-11-05

### Fixed

- Fix the path for download configuration files

## [2.0.0] - 2026-11-05

### Added

- Add GUI Installer
- Add Helium browser and KeePassXC to GUI installer

## [1.5.1] - 2026-11-04

### Changed

- Update the download URL

## [1.5.0] - 2026-11-03

### Added

- Add rose-pine theme for eza
- Add new alias for eza
- Add eza config

## [1.4.0] - 2026-11-02

### Added

- Add rose-pine theme for fzf
- Add config file for bat
- Add rose-pine theme for bat
- Add rose-pine theme for micro
- Add cofig file for micro

### Changed

- Update parameter name in the download function

## [1.3.0] - 2026-10-31

### Added

- Add micro

## [1.2.0] - 2026-10-30

### Added

- Add fd and bat installer
- Add dir function
- Add .zshrc config file
- Add alias.zsh file
- Add configuration to zshrc
- Add new executor function
- Add the configuration applier to the existing steps

### Fixed

- Fix an invisible error that occurs when downloading files

### Update

- Update homebrew and ssh to use the new dir function

## [1.1.0] - 2026-10-27

### Added

- Add ripgrep installer

### Changed

- Update justfile to improve the build on mac and add build process on Linux

## [1.0.0] - 2026-10-21

### Changed

- Update messages

## [0.8.3] - 2026-10-21

### Fixed

- Fix the path cannot be found when installing homebrew

## [0.8.2] - 2026-10-21

### Fixed

- Fix homebrew installer

## [0.8.1] - 2026-10-21

### Fixed

- Try to fix homebrew installer

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
