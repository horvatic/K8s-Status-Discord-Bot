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
		Generatename      string    `json:"generateName"`
		Namespace         string    `json:"namespace"`
		Selflink          string    `json:"selfLink"`
		UID               string    `json:"uid"`
		Resourceversion   string    `json:"resourceVersion"`
		Creationtimestamp time.Time `json:"creationTimestamp"`
		Labels            struct {
			App             string `json:"app"`
			PodTemplateHash string `json:"pod-template-hash"`
		} `json:"labels"`
		Annotations struct {
			CniProjectcalicoOrgPodip  string `json:"cni.projectcalico.org/podIP"`
			CniProjectcalicoOrgPodips string `json:"cni.projectcalico.org/podIPs"`
		} `json:"annotations"`
		Ownerreferences []struct {
			Apiversion         string `json:"apiVersion"`
			Kind               string `json:"kind"`
			Name               string `json:"name"`
			UID                string `json:"uid"`
			Controller         bool   `json:"controller"`
			Blockownerdeletion bool   `json:"blockOwnerDeletion"`
		} `json:"ownerReferences"`
		Managedfields []struct {
			Manager    string    `json:"manager"`
			Operation  string    `json:"operation"`
			Apiversion string    `json:"apiVersion"`
			Time       time.Time `json:"time"`
			Fieldstype string    `json:"fieldsType"`
			Fieldsv1   struct {
				FMetadata struct {
					FGeneratename struct {
					} `json:"f:generateName"`
					FLabels struct {
						NAMING_FAILED struct {
						} `json:"."`
						FApp struct {
						} `json:"f:app"`
						FPodTemplateHash struct {
						} `json:"f:pod-template-hash"`
					} `json:"f:labels"`
					FOwnerreferences struct {
						NAMING_FAILED struct {
						} `json:"."`
						KUID8Afdf0B5201B49C589063530Ab5Fe5C0 struct {
							NAMING_FAILED struct {
							} `json:"."`
							FApiversion struct {
							} `json:"f:apiVersion"`
							FBlockownerdeletion struct {
							} `json:"f:blockOwnerDeletion"`
							FController struct {
							} `json:"f:controller"`
							FKind struct {
							} `json:"f:kind"`
							FName struct {
							} `json:"f:name"`
							FUID struct {
							} `json:"f:uid"`
						} `json:"k:{"uid":"8afdf0b5-201b-49c5-8906-3530ab5fe5c0"}"`
					} `json:"f:ownerReferences"`
				} `json:"f:metadata"`
				FSpec struct {
					FContainers struct {
						KNameGokubTester struct {
							NAMING_FAILED struct {
							} `json:"."`
							FEnv struct {
								NAMING_FAILED struct {
								} `json:"."`
								KNameNamespace struct {
									NAMING_FAILED struct {
									} `json:"."`
									FName struct {
									} `json:"f:name"`
									FValuefrom struct {
										NAMING_FAILED struct {
										} `json:"."`
										FConfigmapkeyref struct {
											NAMING_FAILED struct {
											} `json:"."`
											FKey struct {
											} `json:"f:key"`
											FName struct {
											} `json:"f:name"`
										} `json:"f:configMapKeyRef"`
									} `json:"f:valueFrom"`
								} `json:"k:{"name":"NAMESPACE"}"`
								KNameService struct {
									NAMING_FAILED struct {
									} `json:"."`
									FName struct {
									} `json:"f:name"`
									FValuefrom struct {
										NAMING_FAILED struct {
										} `json:"."`
										FConfigmapkeyref struct {
											NAMING_FAILED struct {
											} `json:"."`
											FKey struct {
											} `json:"f:key"`
											FName struct {
											} `json:"f:name"`
										} `json:"f:configMapKeyRef"`
									} `json:"f:valueFrom"`
								} `json:"k:{"name":"SERVICE"}"`
							} `json:"f:env"`
							FImage struct {
							} `json:"f:image"`
							FImagepullpolicy struct {
							} `json:"f:imagePullPolicy"`
							FLivenessprobe struct {
								NAMING_FAILED struct {
								} `json:"."`
								FFailurethreshold struct {
								} `json:"f:failureThreshold"`
								FHttpget struct {
									NAMING_FAILED struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FPort struct {
									} `json:"f:port"`
									FScheme struct {
									} `json:"f:scheme"`
								} `json:"f:httpGet"`
								FInitialdelayseconds struct {
								} `json:"f:initialDelaySeconds"`
								FPeriodseconds struct {
								} `json:"f:periodSeconds"`
								FSuccessthreshold struct {
								} `json:"f:successThreshold"`
								FTimeoutseconds struct {
								} `json:"f:timeoutSeconds"`
							} `json:"f:livenessProbe"`
							FName struct {
							} `json:"f:name"`
							FPorts struct {
								NAMING_FAILED struct {
								} `json:"."`
								KContainerport8080ProtocolTCP struct {
									NAMING_FAILED struct {
									} `json:"."`
									FContainerport struct {
									} `json:"f:containerPort"`
									FName struct {
									} `json:"f:name"`
									FProtocol struct {
									} `json:"f:protocol"`
								} `json:"k:{"containerPort":8080,"protocol":"TCP"}"`
							} `json:"f:ports"`
							FResources struct {
							} `json:"f:resources"`
							FTerminationmessagepath struct {
							} `json:"f:terminationMessagePath"`
							FTerminationmessagepolicy struct {
							} `json:"f:terminationMessagePolicy"`
						} `json:"k:{"name":"gokub-tester"}"`
					} `json:"f:containers"`
					FDnspolicy struct {
					} `json:"f:dnsPolicy"`
					FEnableservicelinks struct {
					} `json:"f:enableServiceLinks"`
					FRestartpolicy struct {
					} `json:"f:restartPolicy"`
					FSchedulername struct {
					} `json:"f:schedulerName"`
					FSecuritycontext struct {
					} `json:"f:securityContext"`
					FTerminationgraceperiodseconds struct {
					} `json:"f:terminationGracePeriodSeconds"`
				} `json:"f:spec"`
				FStatus struct {
					FConditions struct {
						KTypeContainersready struct {
							NAMING_FAILED struct {
							} `json:"."`
							FLastprobetime struct {
							} `json:"f:lastProbeTime"`
							FLasttransitiontime struct {
							} `json:"f:lastTransitionTime"`
							FStatus struct {
							} `json:"f:status"`
							FType struct {
							} `json:"f:type"`
						} `json:"k:{"type":"ContainersReady"}"`
						KTypeInitialized struct {
							NAMING_FAILED struct {
							} `json:"."`
							FLastprobetime struct {
							} `json:"f:lastProbeTime"`
							FLasttransitiontime struct {
							} `json:"f:lastTransitionTime"`
							FStatus struct {
							} `json:"f:status"`
							FType struct {
							} `json:"f:type"`
						} `json:"k:{"type":"Initialized"}"`
						KTypeReady struct {
							NAMING_FAILED struct {
							} `json:"."`
							FLastprobetime struct {
							} `json:"f:lastProbeTime"`
							FLasttransitiontime struct {
							} `json:"f:lastTransitionTime"`
							FStatus struct {
							} `json:"f:status"`
							FType struct {
							} `json:"f:type"`
						} `json:"k:{"type":"Ready"}"`
					} `json:"f:conditions"`
					FContainerstatuses struct {
					} `json:"f:containerStatuses"`
					FHostip struct {
					} `json:"f:hostIP"`
					FPhase struct {
					} `json:"f:phase"`
					FPodip struct {
					} `json:"f:podIP"`
					FPodips struct {
						NAMING_FAILED struct {
						} `json:"."`
						KIP101212 struct {
							NAMING_FAILED struct {
							} `json:"."`
							FIP struct {
							} `json:"f:ip"`
						} `json:"k:{"ip":"10.1.21.2"}"`
					} `json:"f:podIPs"`
					FStarttime struct {
					} `json:"f:startTime"`
				} `json:"f:status"`
			} `json:"fieldsV1,omitempty"`
		} `json:"managedFields"`
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
		Selflink          string    `json:"selfLink"`
		UID               string    `json:"uid"`
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
		Annotations struct {
			NodeAlphaKubernetesIoTTL                         string `json:"node.alpha.kubernetes.io/ttl"`
			ProjectcalicoOrgIpv4Address                      string `json:"projectcalico.org/IPv4Address"`
			ProjectcalicoOrgIpv4Vxlantunneladdr              string `json:"projectcalico.org/IPv4VXLANTunnelAddr"`
			VolumesKubernetesIoControllerManagedAttachDetach string `json:"volumes.kubernetes.io/controller-managed-attach-detach"`
		} `json:"annotations"`
		Managedfields []struct {
			Manager    string    `json:"manager"`
			Operation  string    `json:"operation"`
			Apiversion string    `json:"apiVersion"`
			Time       time.Time `json:"time"`
			Fieldstype string    `json:"fieldsType"`
			Fieldsv1   struct {
				FMetadata struct {
					FAnnotations struct {
						FProjectcalicoOrgIpv4Address struct {
						} `json:"f:projectcalico.org/IPv4Address"`
						FProjectcalicoOrgIpv4Vxlantunneladdr struct {
						} `json:"f:projectcalico.org/IPv4VXLANTunnelAddr"`
					} `json:"f:annotations"`
				} `json:"f:metadata"`
				FStatus struct {
					FConditions struct {
						KTypeNetworkunavailable struct {
							NAMING_FAILED struct {
							} `json:"."`
							FLastheartbeattime struct {
							} `json:"f:lastHeartbeatTime"`
							FLasttransitiontime struct {
							} `json:"f:lastTransitionTime"`
							FMessage struct {
							} `json:"f:message"`
							FReason struct {
							} `json:"f:reason"`
							FStatus struct {
							} `json:"f:status"`
							FType struct {
							} `json:"f:type"`
						} `json:"k:{"type":"NetworkUnavailable"}"`
					} `json:"f:conditions"`
				} `json:"f:status"`
			} `json:"fieldsV1"`
		} `json:"managedFields"`
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
	devPods, _ := clientset.CoreV1().Pods("dev").List(context.TODO(), v1.ListOptions{})
	report = report + "Dev Pods\n\n"
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
	prodPods, _ := clientset.CoreV1().Pods("prod").List(context.TODO(), v1.ListOptions{})
	report = report + "Prod Pods\n\n"
	for _, p := range prodPods.Items {
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

	sendPayload(report)

}

func sendPayload(content string) {
	//discord web hook, used to post to discord
	discordUrl := "<discord-webhook>"

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
