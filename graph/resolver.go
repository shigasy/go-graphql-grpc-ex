package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "github.com/shigasy/go-graphql-grpc-ex/article/client"

type Resolver struct {
	ArticleClient *client.Client
}
