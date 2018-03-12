/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2018 Red Hat, Inc.
 *
 */

package tests_test

import (
	"flag"
	"fmt"
	//"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	k8sv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"kubevirt.io/kubevirt/pkg/api/v1"
	"kubevirt.io/kubevirt/pkg/kubecli"
	"kubevirt.io/kubevirt/tests"
)

const (
	windowsDisk     = "windows-disk"
	windowsFirmware = "5d307ca9-b3ef-428c-8861-06e72d69f223"
	windowsVm       = "windows-vm"
	windowsVmUser   = "Administrator"
	windowsVmPassword = "Heslo123"
)

const (
	winrmCli    = "winrmcli"
	winrmCliCmd = "/go/bin/winrm-cli"
)

var _ = FDescribe("Windows VM", func() {
	flag.Parse()

	virtClient, err := kubecli.GetKubevirtClient()
	tests.PanicOnError(err)

	gracePeriod := int64(0)
	spinlocks := uint32(8191)
	firmware := types.UID(windowsFirmware)
	_false := false

	windowsVm := &v1.VirtualMachine{
		ObjectMeta: metav1.ObjectMeta{Name: windowsVm},
		Spec: v1.VirtualMachineSpec{
			TerminationGracePeriodSeconds: &gracePeriod,
			Domain: v1.DomainSpec{
				CPU: &v1.CPU{Cores: 2},
				Features: &v1.Features{
					ACPI: v1.FeatureState{},
					APIC: &v1.FeatureAPIC{},
					Hyperv: &v1.FeatureHyperv{
						Relaxed:   &v1.FeatureState{},
						VAPIC:     &v1.FeatureState{},
						Spinlocks: &v1.FeatureSpinlocks{Retries: &spinlocks},
					},
				},
				Clock: &v1.Clock{
					ClockOffset: v1.ClockOffset{UTC: &v1.ClockOffsetUTC{}},
					Timer: &v1.Timer{
						HPET:   &v1.HPETTimer{Enabled: &_false},
						PIT:    &v1.PITTimer{TickPolicy: v1.PITTickPolicyDelay},
						RTC:    &v1.RTCTimer{TickPolicy: v1.RTCTickPolicyCatchup},
						Hyperv: &v1.HypervTimer{},
					},
				},
				Firmware: &v1.Firmware{UUID: firmware},
				Resources: v1.ResourceRequirements{
					Requests: k8sv1.ResourceList{
						k8sv1.ResourceMemory: resource.MustParse("4096Mi"),
					},
				},
				Devices: v1.Devices{
					Disks: []v1.Disk{
						{
							Name:       windowsDisk,
							VolumeName: windowsDisk,
							DiskDevice: v1.DiskDevice{Disk: &v1.DiskTarget{Bus: "sata"}},
						},
					},
				},
			},
			Volumes: []v1.Volume{
				{
					Name: windowsDisk,
					VolumeSource: v1.VolumeSource{
						Ephemeral: &v1.EphemeralVolumeSource{
							PersistentVolumeClaim: &k8sv1.PersistentVolumeClaimVolumeSource{
								ClaimName: tests.DiskWindows,
							},
						},
					},
				},
			},
		},
	}

	tests.BeforeAll(func() {
		windowsPv, err := virtClient.CoreV1().PersistentVolumes().Get(tests.DiskWindows, metav1.GetOptions{})
		if err != nil {
			Skip(fmt.Sprintf("Skip Windows tests that requires PVC %s", tests.DiskWindows))
		}
		windowsPv.Spec.ClaimRef = nil
		_, err = virtClient.CoreV1().PersistentVolumes().Update(windowsPv)
		Expect(err).ToNot(HaveOccurred())
	})

	BeforeEach(func() {
		tests.BeforeTestCleanup()
	})

	It("should success to start", func() {
		vm, err := virtClient.VM(tests.NamespaceTestDefault).Create(windowsVm)
		Expect(err).To(BeNil())
		tests.WaitForSuccessfulVMStartWithTimeout(vm, 180)
	}, 300)

	It("should success to stop", func() {
		By("Starting a VM")
		vm, err := virtClient.VM(tests.NamespaceTestDefault).Create(windowsVm)
		Expect(err).To(BeNil())
		tests.WaitForSuccessfulVMStartWithTimeout(vm, 180)

		By("Stopping the VM")
		err = virtClient.VM(tests.NamespaceTestDefault).Delete(vm.Name, &metav1.DeleteOptions{})
		Expect(err).To(BeNil())
	}, 300)

	Context("with winrm connection", func() {
		var winrmcliPod *k8sv1.Pod
		var cli string

		BeforeEach(func() {
			tests.BeforeTestCleanup()
			winrmcliPod = &k8sv1.Pod{
				ObjectMeta: metav1.ObjectMeta{Name: winrmCli},
				Spec: k8sv1.PodSpec{
					Containers: []k8sv1.Container{
						{
							Name:    winrmCli,
							Image:   fmt.Sprintf("%s/%s:%s", tests.KubeVirtRepoPrefix, winrmCli, tests.KubeVirtVersionTag),
							Command: []string{"sleep"},
							Args:    []string{"3600"},
						},
					},
				},
			}
			winrmcliPod, err = virtClient.CoreV1().Pods(tests.NamespaceTestDefault).Create(winrmcliPod)
			Expect(err).ToNot(HaveOccurred())

			vm, err := virtClient.VM(tests.NamespaceTestDefault).Create(windowsVm)
			Expect(err).To(BeNil())
			tests.WaitForSuccessfulVMStartWithTimeout(vm, 180)

			vm, err = virtClient.VM(tests.NamespaceTestDefault).Get(vm.Name, metav1.GetOptions{})
			cli = fmt.Sprintf("%s -hostname %s -username \"%s\" -password \"%s\"",
				winrmCliCmd,
				vm.Status.Interfaces[0].IP,
				windowsVmUser,
				windowsVmPassword,
			)
		})

		XIt("should have correct UUID", func() {
			command := fmt.Sprintf("%s 'wmic csproduct get \"UUID\"'", cli)
			out, err := tests.ExecuteCommandOnPod(
				virtClient,
				winrmcliPod,
				winrmcliPod.Spec.Containers[0].Name,
				command,
			)
			Expect(err).To(BeNil())
			Expect(out).To(Equal(windowsFirmware))
		}, 300)

		XIt("should have pod IP", func() {})
	})
})
