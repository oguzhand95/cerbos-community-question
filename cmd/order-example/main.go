package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
)

func main() {
	c, err := cerbos.New("127.0.0.1:3593", cerbos.WithTLSInsecure(), cerbos.WithPlaintext())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	alex := "alex"
	john := "john"
	action := "update"

	principal := cerbos.NewPrincipal(alex, "user")
	resource := cerbos.NewResource("order", "order_123").WithAttr("user_id", alex)
	isAllowed(ctx, c, principal, resource, action)

	principal = cerbos.NewPrincipal(john, "supervisor")
	resource = cerbos.NewResource("order", "order_123")
	isAllowed(ctx, c, principal, resource, action)
}

func isAllowed(
	ctx context.Context,
	c *cerbos.GRPCClient,
	principal *cerbos.Principal,
	resource *cerbos.Resource,
	action string,
) {
	isAllowed, err := c.IsAllowed(
		ctx,
		principal,
		resource,
		action,
	)
	if err != nil {
		fmt.Printf("Failed to check is allowed: %v\n", err)
		return
	}

	if !isAllowed {
		fmt.Printf("action '%s' is not allowed for principal '%s'\n", action, principal.ID())
		return
	}

	fmt.Printf("action '%s' is allowed for principal '%s'\n", action, principal.ID())
}
