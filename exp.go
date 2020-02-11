package main

import (
	"fmt"
	"log"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha2"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func GenerateCreateGithubStatusTask() {
	var task v1alpha1.Task
	// task.SetName("")
	// task.SetDefaults()
	task = v1alpha1.Task{
		TypeMeta: v1.TypeMeta{
			Kind:       "Task",
			APIVersion: "tekton.dev/v1alpha1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "create-github-status-task",
		},
		Spec: v1alpha1.TaskSpec{
			Inputs: &v1alpha1.Inputs{
				Params: []v1alpha1.ParamSpec{
					v1alpha1.ParamSpec{
						Name:        "REPO",
						Type:        v1alpha1.ParamTypeString,
						Description: "The repo to publish the status update for e.g. tektoncd/triggers",
					},
					v1alpha1.ParamSpec{
						Name:        "COMMIT_SHA",
						Type:        v1alpha1.ParamTypeString,
						Description: "The specific commit to report a status for.",
					},
					v1alpha1.ParamSpec{
						Name:        "STATE",
						Type:        v1alpha1.ParamTypeString,
						Description: "The state to report error, failure, pending, or success.",
					},
					v1alpha1.ParamSpec{
						Name:        "TARGET_URL",
						Type:        v1alpha1.ParamTypeString,
						Description: "The target URL to associate with this status.",
						Default: &v1alpha2.ArrayOrString{
							StringVal: "",
						},
					},
					v1alpha1.ParamSpec{
						Name:        "DESCRIPTION",
						Type:        v1alpha1.ParamTypeString,
						Description: "A short description of the status.",
					},
					v1alpha1.ParamSpec{
						Name:        "CONTEXT",
						Type:        v1alpha1.ParamTypeString,
						Description: "A string label to differentiate this status from the status of other systems.",
					},
				},
			},
			Steps: []v1alpha1.Step{
				v1alpha1.Step{
					Container: corev1.Container{
						Name:       "start-status",
						Image:      "quay.io/kmcdermo/github-tool:latest",
						WorkingDir: "/workspace/source",
						Env: []corev1.EnvVar{
							corev1.EnvVar{
								Name: "GITHUB_TOKEN",
								ValueFrom: &corev1.EnvVarSource{
									SecretKeyRef: &corev1.SecretKeySelector{
										LocalObjectReference: corev1.LocalObjectReference{
											Name: "github-auth",
										},
										Key: "token",
									},
								},
							},
						},
						Command: []string{"github-tools"},
						Args: []string{
							"create-status",
							"--repo",
							"$(inputs.params.REPO)",
							"--sha",
							"$(inputs.params.COMMIT_SHA)",
							"--state",
							"$(inputs.params.STATE)",
							"--target-url",
							"$(inputs.params.TARGET_URL)",
							"--description",
							"$(inputs.params.DESCRIPTION)",
							"--context",
							"$(inputs.params.CONTEXT)",
						},
					},
				},
			},
		},
	}
	log.Println(task)
}

func GenerateCreateGithubStatusTaskFromUnstrcutured() {
	task := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tekton.dev/v1alpha1",
			"kind":       "Task",
			"metadata": map[string]interface{}{
				"name": "create-github-status-task",
			},
			"spec": map[string]interface{}{
				"inputs": map[string]interface{}{
					"params": []interface{}{
						map[string]interface{}{
							"name":        "REPO",
							"description": "The repo to publish the status update for e.g. tektoncd/triggers",
							"type":        "string",
						},
						map[string]interface{}{
							"name":        "COMMIT_SHA",
							"description": "The specific commit to report a status for.",
							"type":        "string",
						},
						map[string]interface{}{
							"name":        "STATE",
							"description": "The state to report error, failure, pending, or success.",
							"type":        "string",
						},
						map[string]interface{}{
							"name":        "TARGET_URL",
							"description": "The target URL to associate with this status.",
							"default":     "",
							"type":        "string",
						},
						map[string]interface{}{
							"name":        "DESCRIPTION",
							"description": "A short description of the status.",
							"type":        "string",
						},
						map[string]interface{}{
							"name":        "CONTEXT",
							"description": "A string label to differentiate this status from the status of other systems.",
							"type":        "string",
						},
					},
				},
				"steps": map[string]interface{}{
					"name":       "start-status",
					"image":      "quay.io/kmcdermo/github-tool:latest",
					"workingDir": "/workspace/source",
					"env": map[string]interface{}{
						"name": "GITHUB_TOKEN",
						"valueFrom": map[string]interface{}{
							"secretKeyRef": map[string]interface{}{
								"name": "github-auth",
								"key":  "token",
							},
						},
					},
					"command": []string{"github-tool"},
					"args": []string{
						"create-status",
						"--repo",
						"$(inputs.params.REPO)",
						"--sha",
						"$(inputs.params.COMMIT_SHA)",
						"--state",
						"$(inputs.params.STATE)",
						"--target-url",
						"$(inputs.params.TARGET_URL)",
						"--description",
						"$(inputs.params.DESCRIPTION)",
						"--context",
						"$(inputs.params.CONTEXT)",
					},
				},
			},
		},
	}
	// task := &unstructured.Unstructured{}
	task.SetName("create-github-status-task")
	task.SetAPIVersion("tekton.dev/v1alpha1")
	task.SetKind("task")
	fmt.Println(task.GetName())
}

// apiVersion: tekton.dev/v1alpha1
// kind: Task
// metadata:
//   name: create-github-status-task
// spec:
//   inputs:
//     params:
//     - name: REPO
//       description: The repo to publish the status update for e.g. tektoncd/triggers
//       type: string
//     - name: COMMIT_SHA
//       description: The specific commit to report a status for.
//       type: string
//     - name: STATE
//       description: The state to report error, failure, pending, or success.
//       type: string
//     - name: TARGET_URL
//       type: string
//       default: ""
//       description: The target URL to associate with this status.
//     - name: DESCRIPTION
//       type: string
//       description: A short description of the status.
//     - name: CONTEXT
//       type: string
//       description: A string label to differentiate this status from the status of other systems.
//   steps:
//   - name: start-status
//     image: quay.io/kmcdermo/github-tool:latest
//     workingDir: /workspace/source
//     env:
//     - name: GITHUB_TOKEN
//       valueFrom:
//         secretKeyRef:
//           name: github-auth
//           key: token
//     command: ["github-tool"]
//     args:
//       - "create-status"
//       - "--repo"
//       - "$(inputs.params.REPO)"
//       - "--sha"
//       - "$(inputs.params.COMMIT_SHA)"
//       - "--state"
//       - "$(inputs.params.STATE)"
//       - "--target-url"
//       - "$(inputs.params.TARGET_URL)"
//       - "--description"
//       - "$(inputs.params.DESCRIPTION)"
//       - "--context"
//       - "$(inputs.params.CONTEXT)"

// func main() {
// 	// GenerateCreateGithubStatusTaskFromUnstrcutured()
// 	GenerateCreateGithubStatusTask()
// }
