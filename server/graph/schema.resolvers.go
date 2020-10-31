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
	// check if name is not taken
	//err := r.Database.QueryRow(`SELECT name FROM tasks WHERE task=$1`, name).Scan(&name)
	//if err == sql.ErrNoRows {

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

	stmt, err := r.Database.Prepare(`INSERT INTO tasks VALUES($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(id, name, description, false, ts, ts, taskSlug)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTask(ctx context.Context, id string, name string, description string, done bool) (*model.Task, error) {
	// check if task exists
	var task model.Task
	err := r.Database.QueryRow(`SELECT * FROM tasks WHERE id=$1`, id).Scan(&task.ID, &task.Name, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt, &task.Slug)
	if err != nil {
		return nil, err
	}

	// check if name is not taken
	//err = r.Database.QueryRow(`SELECT name FROM tasks WHERE task=$1`, name).Scan(&name)
	//if err != sql.ErrNoRows {
	//	return nil, fmt.Errorf("This name is already taken")
	//}

	// get update time
	ts := time.Now().Format("02-Jan-2006 15:04:05")

	// create slug
	newSlug := slug.Make(name)

	stmt, err := r.Database.Prepare(`UPDATE tasks SET taskName=$1, taskDescription=$2, done=$3, updated_at=$4, slug=$5 WHERE id=$6`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(name, description, done, ts, newSlug, id)
	if err != nil {
		return nil, err
	}

	task.Name = name
	task.Description = description
	task.Done = done
	task.UpdatedAt = ts
	task.Slug = newSlug

	return &task, nil
}

func (r *queryResolver) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	var tasks []*model.Task

	rows, err := r.Database.Query(`SELECT * FROM TASKS`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt, &task.Slug)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *queryResolver) GetTaskBySlug(ctx context.Context, slug string) (*model.Task, error) {
	var task model.Task
	err := r.Database.QueryRow(`SELECT * FROM tasks WHERE slug=$1`, slug).Scan(&task.ID, &task.Name, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt, &task.Slug)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *queryResolver) GetTaskByID(ctx context.Context, id string) (*model.Task, error) {
	var task model.Task
	err := r.Database.QueryRow(`SELECT * FROM tasks WHERE id=$1`, id).Scan(&task.ID, &task.Name, &task.Description, &task.Done, &task.CreatedAt, &task.UpdatedAt, &task.Slug)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) GetTask(ctx context.Context, slug string) (*model.Task, error) {
	var task model.Task
	err := r.Database.QueryRow(`SELECT * FROM tasks WHERE slug=$1`, slug).Scan(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
