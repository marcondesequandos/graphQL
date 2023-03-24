//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen" //por estar com underline quer dizer que não é obrigatório
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
