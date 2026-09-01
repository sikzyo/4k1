# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.5] - 2026-09-01

### Added

- Se agrego la plantilla para el script de instalación completa `full_install.go`
- Se agrego el sistema para poder navegar desde el menu principal
- Se agrego una forma de manejar los errores del menu principal

### Changed

- Se modifico `input.go` para que mejorar el manejo de errores y los valores que retorna

## [0.0.4] - 2026-08-31

### Added

- Se agrega `command.go` para ejecutar comandos de sistema
- Se utiliza `command.go` en `menu.go` para limpiar la pantalla al momento de iniciar el programa, ademas de manejar los errores en dado caso de que el script falle

### Changed

- Se mejoran los mensajes de entrada de texto para mayor claridad

## [0.0.3] - 2026-08-31

### Added

- Agregar sistema de inputs `input/input.go`
- Usar el sistema de inputs como prueba en `main.go`

## [0.0.2] - 2026-08-31

### Changed

- Agregar cambios al CHANGELOG

## [0.0.1] - 2026-08-29

- Actualizar el `CHANGELOG.md`
- Generar archivo `CHANGELOG.md`
- Agregar licencia MIT
- Realizar el primer commit
- Agregar archivo `main.go` que inicializa el menu principal de la aplicación
- Agregar `logo.go` para mostrar el logo de la aplicación
- Agregar `divider.go` para las separaciones de el menu
- Agregar `menu.go` como punto de partida para mostrar de manera dinámica diferentes menus
- Agregar `models/menu.go` para estructurar los datos para crear menus
- Agregar parámetro para opciones en el modelo `models/menu.go`
- Imprimir las opciones en `menu.go`
- Eliminar espacio innecesario en `logo.go`
- Definir opciones del menu principal en `main.go`
- Agregar campo de salida en `models/menu.go` y su respectiva implementación en `menu.go`
- Corregir error en los indices de las opciones