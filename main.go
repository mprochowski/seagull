package main

import (
    "flag"
    "fmt"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
    "log"
    "net/http"
    "path/filepath"
)

func main() {

    var kubeconfig *string
    if home := homedir.HomeDir(); home != "" {
        kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
        kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    // create the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    fmt.Printf("Starting server at port 8080\n")

    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

        deployments, _ := getDeployments(*clientset)

        fmt.Fprintf(writer, "<table><tr><th>Name</th><th>Version</th><th>Latest version</th></tr>")
        for _, deployment := range deployments {
            fmt.Fprintf(writer, "<tr><td>%s</td><td>%s</td><td>%s</td></tr>", deployment.Name, deployment.Version, deployment.LatestVersion)
        }
        fmt.Fprintf(writer, "</table>")
    })

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
