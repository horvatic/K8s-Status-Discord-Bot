package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type payload struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Content   string `json:"content"`
}

type Pod struct {
	Metadata struct {
		Name              string    `json:"name"`
		Namespace         string    `json:"namespace"`
		Resourceversion   string    `json:"resourceVersion"`
		Creationtimestamp time.Time `json:"creationTimestamp"`
		Labels            struct {
			App             string `json:"app"`
			PodTemplateHash string `json:"pod-template-hash"`
		} `json:"labels"`
	} `json:"metadata"`
	Spec struct {
		Volumes []struct {
			Name      string `json:"name"`
			Projected struct {
				Sources []struct {
					Serviceaccounttoken struct {
						Expirationseconds int    `json:"expirationSeconds"`
						Path              string `json:"path"`
					} `json:"serviceAccountToken,omitempty"`
					Configmap struct {
						Name  string `json:"name"`
						Items []struct {
							Key  string `json:"key"`
							Path string `json:"path"`
						} `json:"items"`
					} `json:"configMap,omitempty"`
					Downwardapi struct {
						Items []struct {
							Path     string `json:"path"`
							Fieldref struct {
								Apiversion string `json:"apiVersion"`
								Fieldpath  string `json:"fieldPath"`
							} `json:"fieldRef"`
						} `json:"items"`
					} `json:"downwardAPI,omitempty"`
				} `json:"sources"`
				Defaultmode int `json:"defaultMode"`
			} `json:"projected"`
		} `json:"volumes"`
		Containers []struct {
			Name  string `json:"name"`
			Image string `json:"image"`
			Ports []struct {
				Name          string `json:"name"`
				Containerport int    `json:"containerPort"`
				Protocol      string `json:"protocol"`
			} `json:"ports"`
			Env []struct {
				Name      string `json:"name"`
				Valuefrom struct {
					Configmapkeyref struct {
						Name string `json:"name"`
						Key  string `json:"key"`
					} `json:"configMapKeyRef"`
				} `json:"valueFrom"`
			} `json:"env"`
			Resources struct {
			} `json:"resources"`
			Volumemounts []struct {
				Name      string `json:"name"`
				Readonly  bool   `json:"readOnly"`
				Mountpath string `json:"mountPath"`
			} `json:"volumeMounts"`
			Livenessprobe struct {
				Httpget struct {
					Path   string `json:"path"`
					Port   int    `json:"port"`
					Scheme string `json:"scheme"`
				} `json:"httpGet"`
				Initialdelayseconds int `json:"initialDelaySeconds"`
				Timeoutseconds      int `json:"timeoutSeconds"`
				Periodseconds       int `json:"periodSeconds"`
				Successthreshold    int `json:"successThreshold"`
				Failurethreshold    int `json:"failureThreshold"`
			} `json:"livenessProbe"`
			Terminationmessagepath   string `json:"terminationMessagePath"`
			Terminationmessagepolicy string `json:"terminationMessagePolicy"`
			Imagepullpolicy          string `json:"imagePullPolicy"`
		} `json:"containers"`
		Restartpolicy                 string `json:"restartPolicy"`
		Terminationgraceperiodseconds int    `json:"terminationGracePeriodSeconds"`
		Dnspolicy                     string `json:"dnsPolicy"`
		Serviceaccountname            string `json:"serviceAccountName"`
		Serviceaccount                string `json:"serviceAccount"`
		Nodename                      string `json:"nodeName"`
		Securitycontext               struct {
		} `json:"securityContext"`
		Schedulername string `json:"schedulerName"`
		Tolerations   []struct {
			Key               string `json:"key"`
			Operator          string `json:"operator"`
			Effect            string `json:"effect"`
			Tolerationseconds int    `json:"tolerationSeconds"`
		} `json:"tolerations"`
		Priority           int    `json:"priority"`
		Enableservicelinks bool   `json:"enableServiceLinks"`
		Preemptionpolicy   string `json:"preemptionPolicy"`
	} `json:"spec"`
	Status struct {
		Phase      string `json:"phase"`
		Conditions []struct {
			Type               string      `json:"type"`
			Status             string      `json:"status"`
			Lastprobetime      interface{} `json:"lastProbeTime"`
			Lasttransitiontime time.Time   `json:"lastTransitionTime"`
		} `json:"conditions"`
		Hostip string `json:"hostIP"`
		Podip  string `json:"podIP"`
		Podips []struct {
			IP string `json:"ip"`
		} `json:"podIPs"`
		Starttime         time.Time `json:"startTime"`
		Containerstatuses []struct {
			Name  string `json:"name"`
			State struct {
				Running struct {
					Startedat time.Time `json:"startedAt"`
				} `json:"running"`
			} `json:"state"`
			Laststate struct {
			} `json:"lastState"`
			Ready        bool   `json:"ready"`
			Restartcount int    `json:"restartCount"`
			Image        string `json:"image"`
			Imageid      string `json:"imageID"`
			Containerid  string `json:"containerID"`
			Started      bool   `json:"started"`
		} `json:"containerStatuses"`
		Qosclass string `json:"qosClass"`
	} `json:"status"`
}

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

func main() {

	report := ""

	// uses the current context in kubeconfig
	// path-to-kubeconfig -- for example, /root/.kube/config
	config, _ := clientcmd.BuildConfigFromFlags("", "<path-to-kubeconfig>")

	clientset, _ := kubernetes.NewForConfig(config)
	report = report + reportPods("dev", clientset)
	report = report + reportPods("prod", clientset)
	report = report + reportNodes(clientset)

	sendPayload(report)

}

func reportNodes(clientset *kubernetes.Clientset) string {
	report := ""
	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	report = report + "Nodes\n\n"
	for _, n := range nodes.Items {
		b, _ := json.Marshal(n)
		node := Node{}
		json.Unmarshal(b, &node)
		report = report + fmt.Sprintf("Node Name: %s\n", node.Metadata.Name)
		for _, c := range node.Status.Conditions {
			report = report + fmt.Sprintf("%s: %s\n", c.Type, c.Status)
		}
		report = report + fmt.Sprintf("\n")
	}
	report = report + fmt.Sprintf("--------------------------\n\n")
	return report
}

func reportPods(namespace string, clientset *kubernetes.Clientset) string {
	report := ""
	devPods, _ := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	report = report + fmt.Sprintf("Pods in namespace: %s\n\n", namespace)
	for _, p := range devPods.Items {
		b, _ := json.Marshal(p)
		pod := Pod{}
		json.Unmarshal(b, &pod)
		report = report + fmt.Sprintf("Pod Name: %s\n", pod.Metadata.Name)
		for _, c := range pod.Status.Conditions {
			report = report + fmt.Sprintf("%s: %s\n", c.Type, c.Status)
		}
		report = report + fmt.Sprintf("\n")
	}
	report = report + fmt.Sprintf("--------------------------\n\n")
	return report
}

func sendPayload(content string) {
	//discord web hook, used to post to discord
	discordUrl := "<discord-web-hook>"

	jsonStr, _ := json.Marshal(&payload{
		Username:  "k8s-healthcheck",
		AvatarUrl: "",
		Content:   content,
	})
	req, err := http.NewRequest("POST", discordUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

//https://kubernetes.io/docs/tasks/administer-cluster/access-cluster-api/
