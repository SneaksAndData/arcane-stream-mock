package main

// run controller-gen to generate CRDs into Helm chart templates
//go:generate controller-gen crd paths=./pkg/apis/... output:crd:dir=./.helm/templates output:stdout
