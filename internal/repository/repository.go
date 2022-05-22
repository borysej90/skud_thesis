package repository

import (
	"context"

	"skud"
)

type Repository interface {
	// GetEmployeeIDByCode returns person ID by pass code.
	GetEmployeeIDByCode(ctx context.Context, code string) (int64, error)

	// GetCurrentAccessNode returns the latest access node that employee has accessed with all direct
	// children that can be physically accessed.
	GetCurrentAccessNode(ctx context.Context, employeeID int64) (*skud.AccessNode, error)

	// GetAccessNodeChecks returns required checks that have to be passed by employee
	// before entering node with ID equal to nodeID and actual checks results.
	GetAccessNodeChecks(ctx context.Context, employeeID, nodeID int64) (skud.Checks, error)
}
