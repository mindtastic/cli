# cli

Generated using `openapi-generator`.

# Generate

Assuming you have the [kopfsachen-dev/api](github.com/kopfsachen-dev/api) repository checked out two levels above this one like this:
```bash
.
├── kopfsache-dev
│   └── api
└── mindtastic
    └── cli
```

execute the following command:
```bash
openapi-generator generate -g go -i ../../kopfsache-dev/api/openapi.yaml -o client/ --git-host github.com --git-repo-id cli --git-user-id mindtastic --package-name client
```
