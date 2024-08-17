# go-pkgs [![CICD](https://github.com/neurocode-io/go-pkgs/actions/workflows/main.yaml/badge.svg)](https://github.com/neurocode-io/go-pkgs/actions/workflows/main.yaml) ![cod cov](https://pub-822896f0c4d14a8b85304d7a32e484f1.r2.dev/go-pkgs.svg)


Collection of Go packages that I use in my projects. Implemented with go generics. Thus they are compatible with Go 1.18+.

## Installation

```bash
go get github.com/neurocode-io/go-pkgs@latest
```


## Packages

### slice

Slice is a package that provides a set of functions to work with slices.

### map

Map is a package that provides a set of functions to work with maps.

### set

Set is a package that provides a generic set implementation.

### async

```
group := result.Group[ResultType]{}

ctx := context.Background()
threshold := 1

group, ctx := resultgroup.WithErrorsThreshold[ResultType](ctx, threshold)

group.Go(func() ([]ResultType, error) {
    return []ResultType{}, nil
})

results, err := group.Wait()
if err != nil {
		fmt.Println("Error:", err)
    fmt.Println("Wrapped errors", err.Unwrap())
	}
```

### streams

Coming soon