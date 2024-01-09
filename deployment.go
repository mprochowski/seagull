package main

import (
    "context"
    "fmt"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
)

type Annotation struct {
    Name          string
    Version       string
    Type          string
    Repository    string
    LatestVersion string
}

func getDeployments(clientset kubernetes.Clientset) ([]Annotation, error) {

    fmt.Println("Log: checking 12")

    var annotations []Annotation

    deployments, err := clientset.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})

    for _, deployment := range deployments.Items {
        if metav1.HasAnnotation(deployment.ObjectMeta, "seagull.prochowski.dev/version") {
            annotation := Annotation{
                Name:    deployment.Name,
                Version: deployment.GetAnnotations()["seagull.prochowski.dev/version"]}

            if metav1.HasAnnotation(deployment.ObjectMeta, "seagull.prochowski.dev/type") {
                annotation.Type = deployment.GetAnnotations()["seagull.prochowski.dev/type"]
            }

            if metav1.HasAnnotation(deployment.ObjectMeta, "seagull.prochowski.dev/repo") {
                annotation.Repository = deployment.GetAnnotations()["seagull.prochowski.dev/repo"]
            }

            annotation.LatestVersion, _ = getVersionFromGithub(annotation.Repository)

            annotations = append(annotations, annotation)
        }
    }

    return annotations, err
}
