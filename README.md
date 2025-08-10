go >= 1.24.4

# da

**da** (aka disk analyzer) a project created for educational purposes to analyze disk space in the form of a file or directory.

<img src="./docs/assets/demo-install.gif" alt="Video instruction" width="600" />

# Building
```
git clone https://github.com/under-tension/go-project-242.git da
cd ./da
make build
```

# Running
```
./bin/hexlet-path-size -r -a -H  ./testdata
```

<table>
    <tr>
        <th>Option name</th>
        <th>Alias</th>
        <th>Description</th>
    </tr>
    <tr>
        <td>--human</td>
        <td>-H</td>
        <td>human-readable sizes</td>
    </tr>
    <tr>
        <td>--all</td>
        <td>-a</td>
        <td>include hidden files and directories</td>
    </tr>
    <tr>
        <td>--recursive</td>
        <td>-r</td>
        <td>recursive size of directories</td>
    </tr>
</table>

# Testing

run all unit tests
```
make test
```

### Tests and linter status when latest merge in master:
[![Actions Status](https://github.com/under-tension/go-project-242/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/under-tension/go-project-242/actions)

[![on-push](https://github.com/under-tension/go-project-242/actions/workflows/on-push.yml/badge.svg)](https://github.com/under-tension/go-project-242/actions/workflows/on-push.yml)