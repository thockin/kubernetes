/*
Copyright 2015 Google Inc. All rights reserved.

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

package e2e

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client/remotecommand"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("emptyDir", func() {
	var (
		c         *client.Client
		podClient client.PodInterface
	)

	BeforeEach(func() {
		var err error
		c, err = loadClient()
		expectNoError(err)

		podClient = c.Pods(api.NamespaceDefault)
	})

	It("should support tmpfs in emptyDir", func() {
	// This test for tmpfs volueme can't run on GCE enviroment because
	// the GCE environment does not support remote command execution.
		if testContext.provider == "GCE" {
			By("Skipping test which is broken for GCE")
			return
		}

		By("creating the pod")
		name := "pod-" + string(util.NewUUID())
		value := strconv.Itoa(time.Now().Nanosecond())
		pod := &api.Pod{
			TypeMeta: api.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1beta1",
			},
			ObjectMeta: api.ObjectMeta{
				Name: name,
				Labels: map[string]string{
					"name": "foo",
					"time": value,
				},
			},
			Spec: api.PodSpec{
				Containers: []api.Container{
					{
						Name:  "nginx",
						Image: "dockerfile/nginx",
						VolumeMounts: []api.VolumeMount{
							{
								Name:      "testvol",
								MountPath: "/testvol",
							},
						},
					},
				},
				Volumes: []api.Volume{
					{
						Name: "testvol",
						VolumeSource: api.VolumeSource{
							EmptyDir: &api.EmptyDirVolumeSource{
								Medium: api.StorageTypeMemory,
							},
						},
					},
				},
			},
		}

		By("submitting the pod to kubernetes")
		defer podClient.Delete(pod.Name)
		_, err := podClient.Create(pod)
		if err != nil {
			Fail(fmt.Sprintf("Failed to create pod: %v", err))
		}

		By("waiting for the pod to start running")
		expectNoError(waitForPodRunning(c, pod.Name))

		By("verifying the pod is in kubernetes")
		pods, err := podClient.List(labels.SelectorFromSet(labels.Set(map[string]string{"time": value})))
		if err != nil {
			Fail(fmt.Sprintf("Failed to query for pods: %v", err))
		}
		Expect(len(pods.Items)).To(Equal(1))

		pod = &pods.Items[0]
		By(fmt.Sprintf("executing command on host %s pod %s in container %s",
			pod.Status.Host, pod.Name, pod.Spec.Containers[0].Name))
		req := c.Get().
			Prefix("proxy").
			Resource("minions").
			Name(pod.Status.Host).
			Suffix("exec", api.NamespaceDefault, pod.Name, pod.Spec.Containers[0].Name)

		out := &bytes.Buffer{}
		clientConfig, err := loadConfig()
		if err != nil {
			Fail(fmt.Sprintf("Failed to create client config: %v", err))
		}
		e := remotecommand.New(req, clientConfig, []string{"df"}, nil, out, nil, false)
		err = e.Execute()
		if err != nil {
			Fail(fmt.Sprintf("Failed to execute command on host %s pod %s in container %s: %v",
				pod.Status.Host, pod.Name, pod.Spec.Containers[0].Name, err))
		}
		// Check result
		volmatch := regexp.MustCompile("tmpfs.*/testvol\n")
		match := volmatch.FindAllString(out.String(), -1)
		if len(match) != 1 {
			Fail(fmt.Sprintf("Fail to mount tmpfs volume in container"))
		}
		return
	})
})
