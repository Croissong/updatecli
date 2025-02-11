package yaml

import (
	"testing"

	"gotest.tools/assert"
)

func TestJoinPathWithWorkingDirectoryPath(t *testing.T) {

	testData := []struct {
		name           string
		fileName       string
		workingDir     string
		expectedResult string
	}{
		{
			name:           "scenario 1",
			fileName:       "test.yaml",
			workingDir:     "/tmp",
			expectedResult: "/tmp/test.yaml",
		},
		{
			name:           "scenario 2",
			fileName:       "/tmp/test.yaml",
			workingDir:     "/opt",
			expectedResult: "/tmp/test.yaml",
		},
		{
			name:           "scenario 3",
			fileName:       "https://test.yaml",
			workingDir:     "/opt",
			expectedResult: "https://test.yaml",
		},
		{
			name:           "scenario 4",
			fileName:       "http://test.yaml",
			workingDir:     "/opt",
			expectedResult: "http://test.yaml",
		},
		{
			name:           "scenario 5",
			fileName:       "test.yaml",
			workingDir:     "",
			expectedResult: "test.yaml",
		},
		{
			name:           "scenario 6",
			fileName:       "./test.yaml",
			workingDir:     "/opt",
			expectedResult: "/opt/test.yaml",
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := joinPathWithWorkingDirectoryPath(tt.fileName, tt.workingDir)
			assert.Equal(t, tt.expectedResult, gotResult)
		})
	}
}

func TestParseKey(t *testing.T) {

	testData := []struct {
		key            string
		expectedResult []string
	}{
		{
			key: `image`,
			expectedResult: []string{
				`image`,
			},
		},
		{
			key: `image.tag`,
			expectedResult: []string{
				`image`,
				`tag`,
			},
		},
		{
			key: `image\.tag`,
			expectedResult: []string{
				`image.tag`,
			},
		},
		{
			key: `image.`,
			expectedResult: []string{
				`image`,
			},
		},
		{
			key: `image\.`,
			expectedResult: []string{
				`image.`,
			},
		},
		{
			key: `image*`,
			expectedResult: []string{
				`image*`,
			},
		},
		{
			key: `image\`,
			expectedResult: []string{
				`image`,
			},
		},
		{
			key: `image\\`,
			expectedResult: []string{
				`image\`,
			},
		},
		{
			key:            ``,
			expectedResult: []string{},
		},
	}

	for _, tt := range testData {
		t.Run(tt.key, func(t *testing.T) {
			gotResult := parseKey(tt.key)

			assert.Equal(t, len(tt.expectedResult), len(gotResult))
			for i := range tt.expectedResult {
				assert.Equal(t, tt.expectedResult[i], gotResult[i])
			}
		})
	}
}
