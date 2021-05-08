package report

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/horvatic/k8s-status-discord-bot/pkg/config"
	"github.com/horvatic/k8s-status-discord-bot/pkg/data"
	"github.com/horvatic/k8s-status-discord-bot/pkg/discord"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ProcessReport(clientset *kubernetes.Clientset) {

	systemConfig := config.GetConfig()

	discord.SendPayload(fmt.Sprintf("**###########################**\nReport For: %s\n**--------------------------**\n\n", time.Now().Format("01-02-2006")), systemConfig.DiscordHook)

	for _, namespace := range systemConfig.Namespaces {
		reportPods(namespace, clientset, systemConfig.DiscordHook)
	}
	reportNodes(clientset, systemConfig.DiscordHook)

	discord.SendPayload("**###########################**\n", systemConfig.DiscordHook)

}

func reportNodes(clientset *kubernetes.Clientset, discordUri string) {
	var report strings.Builder
	nodes, _ := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	report.WriteString("__**Nodes**__\n\n")
	for _, n := range nodes.Items {
		b, _ := json.Marshal(n)
		node := data.Node{}
		json.Unmarshal(b, &node)
		fmt.Fprintf(&report, "Node Name: %s\n", node.Metadata.Name)
		report.WriteString("- - - - - - - - - - - - - -\n")
		for _, c := range node.Status.Conditions {
			fmt.Fprintf(&report, "%s: %s\n", c.Type, c.Status)
			report.WriteString("- - - - - - - - - - - - - -\n")
		}
		report.WriteString("\n\n")
		if report.Len() >= discord.MessageLimit {
			discord.SendPayload(report.String(), discordUri)
			report.Reset()
		}
	}
	report.WriteString("**--------------------------**\n\n")
	discord.SendPayload(report.String(), discordUri)
}

func reportPods(namespace string, clientset *kubernetes.Clientset, discordUri string) {
	var report strings.Builder
	pods, _ := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	fmt.Fprintf(&report, "__**Pods In Namespace: %s**__\n\n", namespace)
	for _, p := range pods.Items {
		b, _ := json.Marshal(p)
		pod := data.Pod{}
		json.Unmarshal(b, &pod)
		fmt.Fprintf(&report, "Pod Name: %s\n", pod.Metadata.Name)
		report.WriteString("- - - - - - - - - - - - - -\n")
		for _, c := range pod.Status.Conditions {
			fmt.Fprintf(&report, "%s: %s\n", c.Type, c.Status)
			report.WriteString("- - - - - - - - - - - - - -\n")
		}
		report.WriteString("\n\n")
		if report.Len() >= discord.MessageLimit {
			discord.SendPayload(report.String(), discordUri)
			report.Reset()
		}
	}
	report.WriteString("**--------------------------**\n\n")
	discord.SendPayload(report.String(), discordUri)
}
