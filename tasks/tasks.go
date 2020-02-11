package main

import "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"

// GenerateTasks will return a slice of tasks
func GenerateTasks() []v1alpha1.Task {
	tasks := []v1alpha1.Task{}
	tasks = append(tasks,
		GenerateGithubStatusTask(),
		GenerateDeployFromSourceTask())
	return tasks
}
