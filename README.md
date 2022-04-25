# Redo

![Release](https://github.com/barthr/redo/actions/workflows/release.yml/badge.svg)
![CI](https://github.com/barthr/redo/actions/workflows/golangci-lint.yml/badge.svg)
![Test](https://github.com/barthr/redo/actions/workflows/test.yml/badge.svg)

Redo is a command line application to easily create reusable functions in your own shell. Think of redo like an
interactive way combine multiple commands from your shell history in a single command. This can be handy for quickly
re-doing multiple commands for example deleting and starting a new docker container.

<p align="center">
  <img src="https://github.com/barthr/redo/blob/master/docs/demo.gif" width="90%" height="90%" />
</p>

* [Features](#features)
* [Installation](#installation)
    * [Prebuilt binaries](#prebuilt-binaries)
    * [Install from source](#install-from-source)
* [Quickstart](#quickstart)
* [Configuration](#configuration)
* [Shortcuts](#shortcuts)
* [Roadmap](#roadmap)
* [Acknowledgements](#acknowledgements)
* [License](#license)

## Features

- Easily create reusable functions from shell history
- Shell agnostic, can be used with ZSH, Bash etc.
- Aliases are stored in a single file which can be put in version control

## Installation

### Prebuilt binaries

Download one of the prebuilt binaries from: https://github.com/barthr/redo/releases and run the following command

```bash
tar -xf <downloaded_archive> redo && sudo mv redo /usr/local/bin
```

### Install from source

```bash
go install github.com/barthr/redo@latest
```

*After downloading add the following line to your* `~/.bashrc` or `~/.zshrc`

```bash
source "$(redo alias-file)"
```

This will make sure that the aliases from redo are loaded on every shell session.

## Quickstart

redo contains a couple of commands, which can be used to create reusable functions.

1. `redo` - Opens up the interactive window to create a new function
2. `redo alias-file` - Prints the path to the functions file
3. `redo edit` - Opens the functions file in your configured editor
4. `redo help` - Prints a help message which includes information about all the commands

## Configuration

Redo can mostly run without requiring any specific configuration, however it is possible to customize this configuration
by setting the following environment variables:

`REDO_ALIAS_PATH`: The path where the alias file of redo is stored (defaults to aliases file in user config dir)

`REDO_CONFIG_PATH`: The config path for redo (defaults to user config dir)

`REDO_HISTORY_PATH`: The location of the history file which redo uses to source commands (*defaults to HISTFILE **if it is
exported**)

`REDO_EDITOR`: The editor you want to use when running commands like `redo edit` (defaults to EDITOR **if it is exported**)

## Shortcuts

Redo can be bind to a shortcut, so you can easily summon it without calling it directly.

**zsh CTRL+e summons redo**:
Put the following line in your zshrc file

```zsh
bindkey -s '^e' 'redo^M'
```

**bash CTRL+e summons redo**:
Put the following line in your bashrc file or bash_profile

```bash
bind '"\C-e":"redo\n"'
```

## Roadmap

- Reordering of selected tasks
- Easy listing/deletion of functions
- Inline editing of shell functions
- Prebuilt binaries published as .deb .rpm .yum etc.

## Acknowledgements

- [Bubbletea TUI framework](https://github.com/charmbracelet/bubbletea)
- [Sh](https://github.com/mvdan/sh)

## License

[MIT](https://choosealicense.com/licenses/mit/)

