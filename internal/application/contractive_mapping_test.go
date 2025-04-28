package application_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/application"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
)

func TestContractiveMapping(t *testing.T) {
	// Prepare a coefficient instance.
	coeff := &domain.Coeff{}

	// Call the function.
	application.ContractiveMapping(coeff)

	// Validate the results.
	sumOfSquares := coeff.AC*coeff.AC + coeff.BC*coeff.BC + coeff.DC*coeff.DC + coeff.EC*coeff.EC
	otherPart := (coeff.AC*coeff.EC - coeff.BC*coeff.DC) * (coeff.AC*coeff.EC - coeff.BC*coeff.DC)

	assert.LessOrEqual(t, sumOfSquares-otherPart, 1.0, "Sum of squares of AC, BC, DC, and EC should be <= 1")

	// Check the specific range constraints.
	assert.NotZero(t, coeff.AC, "AC should not be zero")
	assert.NotZero(t, coeff.BC, "BC should not be zero")
	assert.NotZero(t, coeff.DC, "DC should not be zero")
	assert.NotZero(t, coeff.EC, "EC should not be zero")

	assert.GreaterOrEqual(t, coeff.CC, -2.0, "CC should be >= -2")
	assert.LessOrEqual(t, coeff.CC, 2.0, "CC should be <= 2")
	assert.GreaterOrEqual(t, coeff.FC, -2.0, "FC should be >= -2")
	assert.LessOrEqual(t, coeff.FC, 2.0, "FC should be <= 2")
}
