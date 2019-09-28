package main

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	_ = corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-example",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "ubuntu",
					Image:   "ubuntu:trusty",
					Command: []string{"echo"},
					Args:    []string{"Hello World"},
				},
			},
		},
	}
}
