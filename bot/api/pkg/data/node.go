package data

import (
	"time"
)

type Node struct {
	Metadata struct {
		Name              string    `json:"name"`
		Resourceversion   string    `json:"resourceVersion"`
		Creationtimestamp time.Time `json:"creationTimestamp"`
		Labels            struct {
			BetaKubernetesIoArch string `json:"beta.kubernetes.io/arch"`
			BetaKubernetesIoOs   string `json:"beta.kubernetes.io/os"`
			KubernetesIoArch     string `json:"kubernetes.io/arch"`
			KubernetesIoHostname string `json:"kubernetes.io/hostname"`
			KubernetesIoOs       string `json:"kubernetes.io/os"`
			Microk8SIoCluster    string `json:"microk8s.io/cluster"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
	} `json:"spec"`
	Status struct {
		Capacity struct {
			CPU              string `json:"cpu"`
			EphemeralStorage string `json:"ephemeral-storage"`
			Memory           string `json:"memory"`
			Pods             string `json:"pods"`
		} `json:"capacity"`
		Allocatable struct {
			CPU              string `json:"cpu"`
			EphemeralStorage string `json:"ephemeral-storage"`
			Memory           string `json:"memory"`
			Pods             string `json:"pods"`
		} `json:"allocatable"`
		Conditions []struct {
			Type               string    `json:"type"`
			Status             string    `json:"status"`
			Lastheartbeattime  time.Time `json:"lastHeartbeatTime"`
			Lasttransitiontime time.Time `json:"lastTransitionTime"`
			Reason             string    `json:"reason"`
			Message            string    `json:"message"`
		} `json:"conditions"`
		Addresses []struct {
			Type    string `json:"type"`
			Address string `json:"address"`
		} `json:"addresses"`
		Daemonendpoints struct {
			Kubeletendpoint struct {
				Port int `json:"Port"`
			} `json:"kubeletEndpoint"`
		} `json:"daemonEndpoints"`
		Nodeinfo struct {
			Machineid               string `json:"machineID"`
			Systemuuid              string `json:"systemUUID"`
			Bootid                  string `json:"bootID"`
			Kernelversion           string `json:"kernelVersion"`
			Osimage                 string `json:"osImage"`
			Containerruntimeversion string `json:"containerRuntimeVersion"`
			Kubeletversion          string `json:"kubeletVersion"`
			Kubeproxyversion        string `json:"kubeProxyVersion"`
			Operatingsystem         string `json:"operatingSystem"`
			Architecture            string `json:"architecture"`
		} `json:"nodeInfo"`
		Images []struct {
			Names     []string `json:"names"`
			Sizebytes int      `json:"sizeBytes"`
		} `json:"images"`
	} `json:"status"`
}
