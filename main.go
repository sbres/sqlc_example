package main

import (
	"context"

	"github.com/sbres/sqlc_example/sqlc"
)

func main() {

}

func doSomethingInterface(q sqlc.Querier, ctx context.Context) {
	q.CheckUserExist(ctx, "email string")
}

func doSomethingStruct(q sqlc.Queries, ctx context.Context) {
	q.CheckUserExist(ctx, "email string")
}
