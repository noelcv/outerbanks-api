# outerbanks-api
A graphql api where we can store and get information on characters in Outerbanks.

Fetch library using `go get`
https://github.com/99designs/gqlgen

```go get github.com/99designs/gqlgen```

Add gqlgen to your project's tools.go
```
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
go mod tidy
```

Start the graphql server
```
go run server.go
```


## 1. Create your schema at schema.graphqls

## 2. Generate the schema.resolvers / model_gen.go  /generated.go by running 
`go run github.com/99designs/gqlgen generate`

PS: When regenerating the code on VS Code, you might have a bunch of linting errors.
Reloading the Page should fix them.

## 3. Confirm the API can still work
`go run server.go`