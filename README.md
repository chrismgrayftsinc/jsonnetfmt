jsonnetfmt
===============

A hook for formatting jsonnet files

### Using this hook

Add this to your `.pre-commit-config.yaml`

```
- repo: https://github.com/chrismgrayftsinc/jsonnetfmt.git
  rev: main
  hooks:
    - id: jsonnetfmt
```
