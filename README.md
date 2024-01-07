# parcel

HTTP rendering and binding library for Go

## Getting Started

```
go get -u github.com/jaredhughes1012/parcel
```

## Examples

### HTTP clients

```
func example(client *http.Client) error {
  data := map[string]any {
    "someObject": "goesHere",
  }

  // No package-level default for new requests, simply use the type of content writer you want
  req, _ := httpwrite.NewJsonReader().NewRequest(context.Background(), http.MethodPost, "/a/path", data)
  res, err := client.Do(req)
  if err != nil {
    return err
  }

  var result someStruct
  if err = httpread.ReadRequest(req, &result); err != nil {
    return err
  }

  // ...
}

```
### HTTP handler

```
func handleSomeRequest(w http.ResponseWriter, r *http.Request) {
  var data SomeStruct
  if err := httpread.Request(r, &data); err != nil {
    httpwrite.Error(w, err)
    return
  }

  _ = httpwrite.Response(w, &map[string]any {
    "someObject": "goesHere",
  })
}
```

