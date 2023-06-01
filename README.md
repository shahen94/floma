# Floma CLI tool [![Build](https://github.com/shahen94/floma/actions/workflows/build.yml/badge.svg)](https://github.com/shahen94/floma/actions/workflows/build.yml)
Run multiple processes using one command


### Help
```sh
$ floma --help
```


### Config File Example

```yaml
commands:
  - name: "Run v1"
    exec: "npm start"
    args:
      - development
    environment:
      - NODE_ENV=production
      - NODE_GYP=25
      - FFI_V=1.3
  - name: "Run v2"
    exec: "npm start"
  - name: "Run Backoffice"
    exec: "npm start"

```

### Start floma using default config file - floma.yml
```sh
floma
```

### Specify config file

```sh
floma --config floma.example.yml
```


