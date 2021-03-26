# {{.Info.Name}} 

{{.Info.Description}}


# 目录 
{{range $content:=.Item}}
* [{{$content.Name}}](#{{$content.Name}}) 
{{end}}



# 接口列表

{{range .Item}}
## {{ .Name }} 

{{.Request.Description}}

### Request

- Method : {{.Request.Method}}
- URL : {{.Request.URL.Raw}}
- Query :
    
    | 参数名  | 值  | 说明 |
    |---------|-----|------ |
    {{range $query:=.Request.URL.Query}} | {{$query.Key}} |{{$query.Value}}|{{$query.Description}}| {{end}} 


- Header:
    
    | 参数名  | 值  | 说明 |
    |---------|-----|------ |
    {{range $query:=.Request.Header}} | {{$query.Key}} |{{$query.Value}}|{{$query.Description}}| {{end}} 
- Variable:
    
    | 参数名  | 值  | 说明 |
    |---------|-----|------ |{{range $query:=.Request.URL.Variable}} 
    | {{$query.Key}} |{{$query.Value}}|{{$query.Description}}| {{end}} 


- Body:
    + mode: {{.Request.Body.Mode}}

    {{if eq .Request.Body.Mode "raw"}}
    + Language : {{.Request.Body.Options}}
    + Content:

``` json 
{{.Request.Body.Raw}}

```

    {{end}}

    {{if eq .Request.Body.Mode "urlencoded"}}
    | 参数名  | 值  | 说明 |
    |---------|-----|------ |{{range $query:=.Request.Body.Urlencoded}} 
    | {{$query.Key}} |{{$query.Value}}|{{$query.Description}}| {{end}} 
    {{end}}

    {{if eq .Request.Body.Mode "formdata"}}
    | 参数名  | 值  | 说明 |
    |---------|-----|------ |{{range $query:=.Request.Body.Formdata}} 
    | {{$query.Key}} |{{$query.Value}}|{{$query.Description}}| {{end}} 
    {{end}}



### Response

{{range $index,$response:=.Response}}
    
#### Example{{$index}}

- StatusText: {{$response.Status}}
- StatusCode: {{$response.Code}}
- Body:

``` {{$response.PostmanPreviewlanguage}} 
{{$response.Body}}
```

{{end}}


{{end}}
