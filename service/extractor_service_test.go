package service_test

import (
	"testing"

	"github.com/agungcandra/gojira/entity"
	"github.com/agungcandra/gojira/service"
	"github.com/stretchr/testify/assert"
)

func TestExtractorService_ExtractFromMergeRequest(t *testing.T) {
	testCases := map[string]struct {
		input    *entity.MergeRequestAttribute
		expected []string
	}{
		"source branch issue": {
			input: &entity.MergeRequestAttribute{
				SourceBranch: "XX-1812",
			},
			expected: []string{"XX-1812"},
		},
		"source and title": {
			input: &entity.MergeRequestAttribute{
				SourceBranch: "TR-1812",
				Title:        "[TR-5456] Somethings",
			},
			expected: []string{"TR-1812", "TR-5456"},
		},
		"double issue in title": {
			input: &entity.MergeRequestAttribute{
				SourceBranch: "RT-1812",
				Title:        "[RT-5456][BYS-123] Somethings",
			},
			expected: []string{"RT-1812", "RT-5456", "BYS-123"},
		},
		"empty": {
			input:    &entity.MergeRequestAttribute{},
			expected: []string{},
		},
		"duplicate title and source": {
			input: &entity.MergeRequestAttribute{
				SourceBranch: "SX-10",
				Title:        "[SX-10] Some title here",
			},
			expected: []string{"SX-10"},
		},
		"duplicate issue in title": {
			input: &entity.MergeRequestAttribute{
				SourceBranch: "SX-101",
				Title:        "[SX-10] [SX-10] Some title here",
			},
			expected: []string{"SX-101", "SX-10"},
		},
	}

	serv := service.ExtractorService{}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := serv.ExtractFromMergeRequest(tc.input)
			assert.ElementsMatch(t, result, tc.expected)
		})
	}
}

func TestExtractorService_ExtractFromPushRequest(t *testing.T) {
	testCases := map[string]struct {
		input    *entity.PushRequest
		expected []string
	}{
		"single commit": {
			input: &entity.PushRequest{
				Commits: []entity.Commit{
					{
						Title: "[ABC-101]",
					},
				},
			},
			expected: []string{"ABC-101"},
		},
		"multiple commit": {
			input: &entity.PushRequest{
				Commits: []entity.Commit{
					{
						Title: "[ABC-101]",
					},
					{
						Title: "[ABD-101]",
					},
				},
			},
			expected: []string{"ABC-101", "ABD-101"},
		},
		"multiple commit duplicate": {
			input: &entity.PushRequest{
				Commits: []entity.Commit{
					{
						Title: "[ABC-101]",
					},
					{
						Title: "[ABC-101]",
					},
				},
			},
			expected: []string{"ABC-101"},
		},
		"multiple commit duplicate some": {
			input: &entity.PushRequest{
				Commits: []entity.Commit{
					{
						Title: "[ABC-101]",
					},
					{
						Title: "[ABC-101]",
					},
					{
						Title: "[ABD-101]",
					},
				},
			},
			expected: []string{"ABC-101", "ABD-101"},
		},
		"with refs": {
			input: &entity.PushRequest{
				Ref: "refs/heads/AB-1",
				Commits: []entity.Commit{
					{
						Title: "[ABC-101]",
					},
					{
						Title: "[ABD-101]",
					},
				},
			},
			expected: []string{"AB-1", "ABC-101", "ABD-101"},
		},
		"with refs duplicate": {
			input: &entity.PushRequest{
				Ref: "refs/heads/AB-1",
				Commits: []entity.Commit{
					{
						Title: "[AB-1]",
					},
					{
						Title: "[ABD-101]",
					},
				},
			},
			expected: []string{"AB-1", "ABD-101"},
		},
	}

	serv := service.ExtractorService{}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			result := serv.ExtractFromPushRequest(tc.input)
			assert.ElementsMatch(t, result, tc.expected)
		})
	}
}
