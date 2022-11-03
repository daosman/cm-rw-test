/*
Copyright 2016 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		// get the cm-test-cm ConfigMap in the cm-test-ns namespace
		_, err := clientset.CoreV1().ConfigMaps("cm-test-ns").Get(context.TODO(), "cm-test-cm", metav1.GetOptions{})
		if err != nil {
			panic(err.Error())
		}
		//fmt.Printf("There are %d ConfigMaps in the cm-test namespace\n", len(cms.Items))
		fmt.Printf("Found cm-test-cm ConfigMap\n")

		time.Sleep(10 * time.Second)
	}
}
