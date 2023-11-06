# i18ngen

This project is still work in progress.

Inspired by [Flutter localizations](https://docs.flutter.dev/ui/accessibility-and-localization/internationalization) generator.

## What is i18ngen?

- Generate typed language translations in Go

## Getting started

1. Create a tools folder `./tools` in your project
1. Add `tools.go` file into  `./tools`
1. Add i18ngen as a tool dependency in `./tools/tools.go`
    ```go
    //go:build tools
    // +build tools

    package tools

    import (
        _ "github.com/jonsaw/i18ngen"
    )
    ```
1. Install dependency
    ```
    go get github.com/jonsaw/i18ngen
    ```
1. In your project folder, create a subfolder to house your translations (ie., `./i18n`).
1. Create a file name `i18n.go` in your `./i18n` subfolder (ie., `./i18n/i18n.go`)
1. Insert the following contents into `i18n.go`.
    ```go
    //go:generate go run github.com/jonsaw/i18ngen generate
    package i18n
    ```
1. Create a file name i18ngen.yml in your subfolder (ie., `./i18n/i18ngen.yml`).
1. Insert the following contents into `i18ngen.yml`

    ```yml
    base: translations/en.json

    translations:
      - translations/*.json

    exec:
      path: generated/
      package: generated
    ```
1. Create a subfolder called `translations` (ie., `./i18n/translations`).
1. Create your translations in the `translations` subfolder (ie., `./i18n/translations/en.json`). Click [here](./_example/simple/) to see examples.
1. Run `go generate ./...`
