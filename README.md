
# stt

Serialization Template Tool (STT)

## As a YAML Template Tool

Helm is a great way to organize your application, but when dealing with multiple regions or chart templates, `stt` simplifies the process of generating/modifying values using templates and values.

```shell
go run main.go -t example/temp -o example/output -v example/values.yaml
```

### Example

 `-t example/temp/example_yaml.yaml`

```
templatelist:
{{- range .list }}
  - name: "{{.name}}"
    age: {{.age}}
    city: "{{.city}}"
    status: {{if ge .age 30}}Senior Citizen{{else}}Youngster{{end}}
{{- end }}
```
`-v example/values.yaml`
```
name: John Doe
age: 30
email: john.doe@example.com

list:
  - name: Alice
    age: 25
    city: New York
  - name: Bob
    age: 30
    city: London
  - name: Carol
    age: 28
    city: Paris
```
`-o example/output`
> This will generate/overwrite files in the `example/output` directory.

```
templatelist:
  - name: "Alice"
    age: 25
    city: "New York"
    status: Youngster
  - name: "Bob"
    age: 30
    city: "London"
    status: Senior Citizen
  - name: "Carol"
    age: 28
    city: "Paris"
    status: Youngster

```


## option

```
  -o string
        Output Directory path, only [Dir]
  -t string
        Template file or directory path
  -v string
        Values file
```

Todo:

> Implement JSON template tool

> YAML <-> JSON conversion

> Read from JSON

> Check the validity of the value file and output file

> Map for template to output, also support multi-input/output