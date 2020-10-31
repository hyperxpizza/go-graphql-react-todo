package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gosimple/slug"
	"github.com/hyperxpizza/go-react-gql-todo/graph/generated"
	"github.com/hyperxpizza/go-react-gql-todo/graph/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, name string, description string) (*model.Task, error) {
	// create uuid
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// get current timestamp and convert it to strings
	ts := time.Now().Format("02-Jan-2006 15:04:05")

	// create slug
	taskSlug := slug.Make(name)

	task := model.Task{
		ID:          id.String(),
		Name:        name,
		Description: description,
		Done:        false,
		CreatedAt:   ts,
		UpdatedAt:   ts,
		Slug:        taskSlug,
	}

	stmt, err := r.Database.Prepare(`INSERT INTO tasks VALUES($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id, name, description, false, ts, ts)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTask(ctx context.Context, name string, description string, done bool) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTask(ctx context.Context, slug string) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
