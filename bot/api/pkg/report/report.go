package report

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/horvatic/k8s-status-discord-bot/pkg/data"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KReport struct {
	clientset *kubernetes.Clientset
}

func NewKReport(clientset *kubernetes.Clientset) *KReport {
	return &KReport{clientset: clientset}
}

func (k KReport) GetKInfoNamespace(w http.ResponseWriter, r *http.Request) {
	k8Report := processNameSpaceReport(k.clientset, r.URL.Query()["namespace"])
	fmt.Fprintf(w, "%s", k8Report)
}

func (k KReport) GetKInfoNode(w http.ResponseWriter, r *http.Request) {
	k8Report := processNodeReport(k.clientset)
	fmt.Fprintf(w, "%s", k8Report)
}

func setHeader(report *strings.Builder) {
	fmt.Fprintf(report, "**###########################**\nReport For: %s\n**--------------------------**\n\n", time.Now().Format("01-02-2006"))
}

func setFooter(report *strings.Builder) {
	fmt.Fprintf(report, "**###########################**\n")
}

func processNameSpaceReport(clientset *kubernetes.Clientset, targets []string) string {

	var report strings.Builder
	setHeader(&report)

	for _, n := range targets {
		reportPods(&report, n, clientset)
	}

	setFooter(&report)
	return report.String()

}

func processNodeReport(clientset *kubernetes.Clientset) string {
	var report strings.Builder
	setHeader(&report)
	reportNodes(&report, clientset)
	setFooter(&report)
	return report.String()

}

func reportNodes(report *strings.Builder, clientset *kubernetes.Clientset) {
	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	report.WriteString("__**Nodes**__\n\n")
	for _, n := range nodes.Items {
		b, _ := json.Marshal(n)
		node := data.Node{}
		json.Unmarshal(b, &node)
		fmt.Fprintf(report, "Node Name: %s\n", node.Metadata.Name)
		report.WriteString("- - - - - - - - - - - - - -\n")
		for _, c := range node.Status.Conditions {
			fmt.Fprintf(report, "%s: %s\n", c.Type, c.Status)
			report.WriteString("- - - - - - - - - - - - - -\n")
		}
		report.WriteString("\n\n")
	}
	report.WriteString("**--------------------------**\n\n")
}

func reportPods(report *strings.Builder, namespace string, clientset *kubernetes.Clientset) {
	pods, _ := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	fmt.Fprintf(report, "__**Pods In Namespace: %s**__\n\n", namespace)
	for _, p := range pods.Items {
		b, _ := json.Marshal(p)
		pod := data.Pod{}
		json.Unmarshal(b, &pod)
		fmt.Fprintf(report, "Pod Name: %s\n", pod.Metadata.Name)
		report.WriteString("- - - - - - - - - - - - - -\n")
		for _, c := range pod.Status.Conditions {
			fmt.Fprintf(report, "%s: %s\n", c.Type, c.Status)
			report.WriteString("- - - - - - - - - - - - - -\n")
		}
		report.WriteString("\n\n")
	}
	report.WriteString("**--------------------------**\n\n")
}
