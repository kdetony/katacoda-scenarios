/*
Copyright 2019 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package problem helps deliver actionable feedback to a user based on an error message.
package problem

import (
	"fmt"
	"regexp"

	"k8s.io/minikube/pkg/minikube/out"
	"k8s.io/minikube/pkg/minikube/translate"
)

const issueBase = "https://github.com/kubernetes/minikube/issues"

// Problem represents a known problem in minikube.
type Problem struct {
	// ID is an arbitrary unique and stable string describing this issue
	ID string
	// Err is the original error
	Err error
	// Advice is actionable text that the user should follow
	Advice string
	// URL is a reference URL for more information
	URL string
	// Issues are a list of related issues to this problem
	Issues []int
	// Hide the new issue link: it isn't our problem, and we won't be able to suggest additional assistance.
	HideCreateLink bool
}

// match maps a regular expression to problem metadata.
type match struct {
	Regexp *regexp.Regexp
	Advice string
	URL    string
	Issues []int
	// GOOS is what platforms this problem may be specific to, when disambiguation is necessary.
	GOOS string
	// Hide the new issue link: it isn't our problem, and we won't be able to suggest additional assistance.
	HideCreateLink bool
}

// Display problem metadata to the console
func (p *Problem) Display() {
	out.ErrT(out.FailureType, "Error: [{{.id}}] {{.error}}", out.V{"id": p.ID, "error": p.Err})
	out.ErrT(out.Tip, "Suggestion: {{.advice}}", out.V{"advice": translate.T(p.Advice)})
	if p.URL != "" {
		out.ErrT(out.Documentation, "Documentation: {{.url}}", out.V{"url": p.URL})
	}
	if len(p.Issues) == 0 {
		return
	}
	out.ErrT(out.Issues, "Related issues:")
	issues := p.Issues
	if len(issues) > 3 {
		issues = issues[0:3]
	}
	for _, i := range issues {
		out.ErrT(out.Issue, "{{.url}}", out.V{"url": fmt.Sprintf("%s/%d", issueBase, i)})
	}
}

// FromError returns a known problem from an error on an OS
func FromError(err error, os string) *Problem {
	maps := []map[string]match{
		osProblems,
		vmProblems,
		netProblems,
		deployProblems,
		stateProblems,
	}
	for _, m := range maps {
		for k, v := range m {
			if v.GOOS != "" && v.GOOS != os {
				continue
			}
			if v.Regexp.MatchString(err.Error()) {
				return &Problem{
					Err:            err,
					Advice:         v.Advice,
					URL:            v.URL,
					ID:             k,
					Issues:         v.Issues,
					HideCreateLink: v.HideCreateLink,
				}
			}
		}
	}
	return nil
}
