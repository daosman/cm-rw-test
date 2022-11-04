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
  "os"
  "strconv"
  "time"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/rest"
  "k8s.io/api/core/v1"
)

const (
  CM_TEST_NAMESPACE = "cm-test-ns"
  CM_TEST_NAME      = "cm-test-cm"
  CM_TEST_IT_KEY    = "CurrentIteration"
)

func getDelay(defaultVal int) time.Duration {
  valueStr := os.Getenv("CM_TEST_DELAY")
  if value, err := strconv.Atoi(valueStr); err == nil {
    return time.Duration(value)
  }

  return time.Duration(defaultVal)
}

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

  iteration := 0
  cm, err := clientset.CoreV1().ConfigMaps(CM_TEST_NAMESPACE).Create(context.TODO(), &v1.ConfigMap{
    ObjectMeta: metav1.ObjectMeta{
      Name:      CM_TEST_NAME,
      Namespace: CM_TEST_NAMESPACE,
      Annotations: map[string]string{
        CM_TEST_IT_KEY: string(iteration),
      },
    },
  }, metav1.CreateOptions{})
  if err != nil {
    panic(err.Error())
  }

  fmt.Printf("Created ConfigMap:\"%s\" in namespace:\"%s\" it:%s\n",
             cm.ObjectMeta.Name, cm.ObjectMeta.Namespace, cm.ObjectMeta.Annotations[CM_TEST_IT_KEY])

  for {
    time.Sleep(getDelay(1) * time.Second)

    // get the cm-test-cm ConfigMap in the cm-test-ns namespace
    cur_cm, err := clientset.CoreV1().ConfigMaps(CM_TEST_NAMESPACE).Get(context.TODO(), CM_TEST_NAME, metav1.GetOptions{})
    if err != nil {
      panic(err.Error())
    }

    // Update the iteration in the Annotation section
    iteration++
    cur_cm.ObjectMeta.Annotations[CM_TEST_IT_KEY] = string(iteration)
    _, err = clientset.CoreV1().ConfigMaps(CM_TEST_NAMESPACE).Update(context.TODO(), cur_cm, metav1.UpdateOptions{})
    if err != nil {
      panic(err.Error())
    }
    fmt.Printf("Updated ConfigMap:\"%s\" in namespace:\"%s\" it:%s\n",
               cur_cm.ObjectMeta.Name, cur_cm.ObjectMeta.Namespace, cur_cm.ObjectMeta.Annotations[CM_TEST_IT_KEY])

  }
}
