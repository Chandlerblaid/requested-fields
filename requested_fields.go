package fields

import (
	"context"
	"strings"
)

// RequestedFor returns all requested fields for some Resolver.
func RequestedFor(ctx context.Context, resolver interface{}) []string {
	return RequestedForAt(ctx, resolver, "")
}

// RequestedForAt returns all requested fields for
//some path from a reference Resolver.
func RequestedForAt(ctx context.Context, resolver interface{}, path_to_append string) []string {
	tree := ctx.Value("graphqlRequestTree").(map[string][]string)

	name := nameFromResolver(resolver)
	field := fromResolver(resolver)

	path := append(field.ParentTree, name)

	// Remove the first "query" path
	_, path = path[0], path[1:]

	pathTree := strings.Join(path, ".")

	if path_to_append != "" {
		pathTree += "." + path_to_append
	}

	return tree[pathTree]
}
