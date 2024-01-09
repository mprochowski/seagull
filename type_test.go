package main

import "testing"

func Test_getVersionFromGithub(t *testing.T) {
    type args struct {
        path string
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {"test", args{path: "argoproj/argo-cd"}, "2.9.3", false}}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := getVersionFromGithub(tt.args.path)
            if (err != nil) != tt.wantErr {
                t.Errorf("getVersionFromGithub() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("getVersionFromGithub() got = %v, want %v", got, tt.want)
            }
        })
    }
}
