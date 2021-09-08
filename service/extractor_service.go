package service

import (
	"regexp"
	"strings"

	"github.com/agungcandra/gojira/entity"
)

type ExtractorService struct{}
type Checker map[string]bool

func NewExtractorService() *ExtractorService {
	return &ExtractorService{}
}

func (es *ExtractorService) ExtractFromMergeRequest(mr *entity.MergeRequestAttribute) []string {
	issues := make([]string, 0)
	checker := make(Checker)
	branch := mr.SourceBranch

	if branch != "" {
		checker[branch] = true
		issues = append(issues, branch)
	}

	issues = es.extractIssue(mr.Title, checker, issues)
	return issues
}

func (es *ExtractorService) ExtractFromPushRequest(push *entity.PushRequest) []string {
	issues := make([]string, 0)
	checker := make(Checker)

	refs := strings.SplitN(push.Ref, "/", 3)
	if len(refs) > 2 && refs[2] != "" {
		checker[refs[2]] = true
		issues = append(issues, refs[2])
	}

	for _, commit := range push.Commits {
		issues = es.extractIssue(commit.Title, checker, issues)
	}
	return issues
}

func (es *ExtractorService) extractIssue(issueable string, checker Checker, issues []string) []string {
	re := regexp.MustCompile(`\[([A-Za-z\-0-9]*)\]`)
	matches := re.FindAllStringSubmatch(issueable, -1)
	for _, match := range matches {
		issue := match[1]
		if !checker[issue] && issue != "" {
			issues = append(issues, issue)
			checker[issue] = true
		}
	}

	return issues
}
