
# Redo

Redo is an command line application to easily create reusable functions in your own shell. Think of redo like an interactive way combine multiple commands from history in a single command.

![demo](https://github.com/barthr/redo/blob/master/docs/demo.gif)
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

## Configuration

Redo can mostly run without requiring any specific configuration, however it is possible to customize this configuration by setting the following environment variables:

`REDO_ALIAS_PATH`: The path where the alias file of redo is stored (defaults to aliases file in user config dir)

`REDO_CONFIG_PATH`: The config path for redo (defaults to user config dir)

`REDO_HISTORY_PATH`: The location of the history file which redo uses to source commands (defaults to HISTFILE)

`REDO_EDITOR`: The editor you wan't to use when running commands like `redo edit` (defaults to EDITOR)





## Roadmap

- Reordering of selected tasks
- Easy listing/deletion of functions
- Inline editting of shell functions
- Prebuilt binaries published as .deb .rpm .yum etc.


## License

[MIT](https://choosealicense.com/licenses/mit/)

