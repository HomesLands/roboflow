// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: step.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type BulkInsertStepsParams struct {
	ID                  uuid.UUID
	WorkflowExecutionID uuid.UUID
	Env                 map[string]string
	Node                []byte
	Status              string
}

const getStepByID = `-- name: GetStepByID :one
SELECT id, workflow_execution_id, env, node, inputs, outputs, status, started_at, completed_at FROM steps
WHERE id = $1
`

func (q *Queries) GetStepByID(ctx context.Context, id uuid.UUID) (*Step, error) {
	row := q.db.QueryRow(ctx, getStepByID, id)
	var i Step
	err := row.Scan(
		&i.ID,
		&i.WorkflowExecutionID,
		&i.Env,
		&i.Node,
		&i.Inputs,
		&i.Outputs,
		&i.Status,
		&i.StartedAt,
		&i.CompletedAt,
	)
	return &i, err
}

const listSteps = `-- name: ListSteps :many
SELECT id, workflow_execution_id, env, node, inputs, outputs, status, started_at, completed_at FROM steps
WHERE workflow_execution_id = $1
`

func (q *Queries) ListSteps(ctx context.Context, workflowExecutionID uuid.UUID) ([]*Step, error) {
	rows, err := q.db.Query(ctx, listSteps, workflowExecutionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*Step{}
	for rows.Next() {
		var i Step
		if err := rows.Scan(
			&i.ID,
			&i.WorkflowExecutionID,
			&i.Env,
			&i.Node,
			&i.Inputs,
			&i.Outputs,
			&i.Status,
			&i.StartedAt,
			&i.CompletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStep = `-- name: UpdateStep :exec
UPDATE steps
SET
    status = $2,
    started_at = $3,
    completed_at = $4
WHERE id = $1
`

type UpdateStepParams struct {
	ID          uuid.UUID
	Status      string
	StartedAt   pgtype.Timestamptz
	CompletedAt pgtype.Timestamptz
}

func (q *Queries) UpdateStep(ctx context.Context, arg UpdateStepParams) error {
	_, err := q.db.Exec(ctx, updateStep,
		arg.ID,
		arg.Status,
		arg.StartedAt,
		arg.CompletedAt,
	)
	return err
}