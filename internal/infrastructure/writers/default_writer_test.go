package writers_test

import (
	"image/color"
	"os"
	"testing"

	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/domain"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/internal/infrastructure/writers"
	"github.com/central-university-dev/backend_academy_2024_project_4-go-z0tedd/pkg/mocks"

	"github.com/stretchr/testify/require"
)

func TestDefaultWriter_Write(t *testing.T) {
	tempFile, err := os.CreateTemp("", "output.*.png")
	require.NoError(t, err)

	tempFilePath := tempFile.Name()
	defer os.Remove(tempFilePath)

	config := &domain.Config{
		OutputPath: tempFilePath,
		Invert:     false,
	}

	mockCanvas := new(mocks.Canvas)
	mockCanvas.On("Canvas").Return([][]domain.Pixel{
		{
			{Color: color.RGBA{R: 255, G: 0, B: 0}},
			{Color: color.RGBA{R: 0, G: 255, B: 0}},
		},
		{
			{Color: color.RGBA{R: 0, G: 0, B: 255}},
			{Color: color.RGBA{R: 255, G: 255, B: 0}},
		},
	})

	writer := writers.NewDefaultWriter(mockCanvas, config, 80)

	err = writer.Write()
	require.NoError(t, err)

	// Validate that the output file exists
	fileInfo, err := os.Stat(tempFilePath)
	require.NoError(t, err)
	require.False(t, fileInfo.IsDir())

	// Validate the file content is non-empty
	content, err := os.ReadFile(tempFilePath)
	require.NoError(t, err)
	require.True(t, len(content) > 0)

	mockCanvas.AssertExpectations(t)
}

func TestDefaultWriter_Write_WithInvert(t *testing.T) {
	tempFile, err := os.CreateTemp("", "output.*.bmp")
	require.NoError(t, err)

	tempFilePath := tempFile.Name()
	defer os.Remove(tempFilePath)

	config := &domain.Config{
		OutputPath: tempFilePath,
		Invert:     true,
	}

	mockCanvas := new(mocks.Canvas)
	mockCanvas.On("Canvas").Return([][]domain.Pixel{
		{
			{Color: color.RGBA{R: 255, G: 255, B: 255}},
			{Color: color.RGBA{R: 0, G: 0, B: 0}},
		},
		{
			{Color: color.RGBA{R: 128, G: 128, B: 128}},
			{Color: color.RGBA{R: 64, G: 64, B: 64}},
		},
	})

	writer := writers.NewDefaultWriter(mockCanvas, config, 80)

	err = writer.Write()
	require.NoError(t, err)

	// Validate the file is non-empty
	fileContent, err := os.ReadFile(tempFilePath)
	require.NoError(t, err)
	require.True(t, len(fileContent) > 0)

	mockCanvas.AssertExpectations(t)
}
