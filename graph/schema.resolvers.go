package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/bketelsen/gqlops/graph/generated"
	"github.com/bketelsen/gqlops/graph/model"
)

func (r *queryResolver) Profiles(ctx context.Context) ([]*model.Profile, error) {

	return r.ProfileData, nil
}

func (r *queryResolver) Search(ctx context.Context, name string) (*model.Profile, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
