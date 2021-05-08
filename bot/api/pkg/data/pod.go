package data

import (
	"time"
)

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
