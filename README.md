# serializeGen

As a yaml template tool
helm is a good way to orginaze application, but when we need multi-region, `serializeGen`` can simple generate values.yaml with template and values.

```
go run main.go -template example/temp -output example/output -values example/values.yaml
```

Todo:
json template tool
yaml <-> json
