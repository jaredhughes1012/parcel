# parcel

HTTP rendering/binding library for Go

## Getting Started

Add to your project

```
go get github.com/jaredhughes1012/parcel
```

## Usage

Parcel uses renderers/binders globally based on the type of data received. When binding from sources,
this detection is done using the "Content-Type" header of the source request or response. When rendering
to a target, this uses the type of data passed in for rendering (e.g. text/plain for strings, application/json for structs)

All renderers/binders for a given data type can be modified globally using the corresponding `Set` method. For
example, to change from a json renderer to another type

```
func Example() {
  parcel.SetObjectRenderer(anothertype.Renderer{})
}
```

Custom binders/renderers can be created simply by implementing the `Renderer` or `Binder` interfaces.
Simply pass your custom renderer or binder into one of these global setters to use it globally.

Parcel also supports adding HTTP response data to errors and automatically rendering those errors.
By default parcel formats error bodies as JSON objects but plain text can be used as well. These
wrapped errors include a status which is automatically set on the response.

## Examples

### Bind JSON data from Request

```
type Example struct {
  Field string `json:"field"`
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
  var data Example
  if err := parcel.BindRequest(r, &data); err != nil {
    parcel.RenderErrorResponse(w, err) // Parcel will return correct http status based on error type
  }
}
```

### Bind text data from Request

```
func sampleHandler(w http.ResponseWriter, r *http.Request) {
  var data string

  // Parcel will automatically validate text content by using a string receiver
  if err := parcel.BindRequest(r, &data); err != nil {
    parcel.RenderErrorResponse(w, err)
  }
}
```

### Render JSON to response

```
type Example struct {
  Field string `json:"field"`
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
  data := Example {
    Field: "test"
  }

  if err := parcel.RenderResponse(w, &data); err != nil {
    parcel.RenderErrorResponse(w, err)
  }
}
```

### Render text to response

```

func sampleHandler(w http.ResponseWriter, r *http.Request) {
  data := "This will be rendered as text"

  // no need to use a pointer when rendering text
  if err := parcel.RenderResponse(w, data); err != nil {
    parcel.RenderErrorResponse(w, err)
  }
}
```

### Create HTTP error and respond with it

```

func sampleHandler(w http.ResponseWriter, r *http.Request) {
  err := httperror.New(http.StatusNotFound, "Resource not found")

  // This will set the HTTP response to NotFound(404)
  parcel.RenderErrorResponse(w, err)
}
```

### Wrap existing error with HTTP response

```

func sampleHandler(w http.ResponseWriter, r *http.Request) {
  err := functionThatFailed()

  // This will set the HTTP response to Conflict(409)
  parcel.RenderErrorResponse(w, httperror.Wrap(http.StatusConflict, err))
}
```

### Render error not wrapped by parcel

```

func sampleHandler(w http.ResponseWriter, r *http.Request) {
  err := functionThatFailed()

  // This will set the HTTP response to InternalServerError(500)
  parcel.RenderErrorResponse(w, err)
}
```

### Set Global Renderers and Binders

```
func setHandlers() {
  parcel.SetTextBinder(somepackage.Binder{})
  // This binder will only be used for requests with the mime type "application/sometype"
  parcel.SetObjectBinder("application/sometype", somepackage.Binder{})

  parcel.SetTextRenderer(somepackage.Renderer{})
  parcel.SetObjectRenderer(somepackage.Renderer{})
  parcel.SetErrorRenderer(somepackage.Renderer{})
}
```