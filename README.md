# httperror

A small Go package for HTTP errors that separates internal error details from the message exposed to the client.

## The problem

When an error occurs in an HTTP handler you often want to log the full internal error but only return a safe, user-friendly message in the response body. `httperror` encodes that distinction in a single type.

## Usage

```go
// Internal error with a separate public message
err := httperror.New(err, "something went wrong")

// Log the internal detail
log.Println(err.Error())         // original error message

// Send the public message in the response
json.NewEncoder(w).Encode(err)   // {"errorMessage": "something went wrong"}
```

## Constructors

### `New(err error, message string) *HttpError`
The standard constructor. If `err` is `nil`, the internal error is set to `message` so you always get a valid error.

```go
httperror.New(db.ErrNoRows, "user not found")
```

### `OptNew(err error, message string) *HttpError`
Returns `nil` when `err` is `nil`. Useful when you only want to wrap an error if one actually occurred, without an explicit nil check at the call site.

```go
if httpErr := httperror.OptNew(err, "failed to save record"); httpErr != nil {
    return httpErr
}
```

### `NewStr(message string) *HttpError`
Sets both the internal and public error to the same string. Handy for validation errors where there is no underlying error.

```go
httperror.NewStr("email is required")
```

## JSON serialisation

`HttpError` marshals to JSON with only the public message exposed. The internal error is tagged `json:"-"` and never leaves the server.

```json
{"errorMessage": "user not found"}
```

## Interfaces

`HttpError` implements two interfaces:

- `error` — via `Error() string`, returns the internal error message
- `pubinterr` — via `Public() string`, returns the public message
