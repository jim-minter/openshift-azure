// Package plugin holds the implementation of a plugin.
package plugin

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	acsapi "github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/arm"
	"github.com/openshift/openshift-azure/pkg/config"
	"github.com/openshift/openshift-azure/pkg/healthcheck"
	"github.com/openshift/openshift-azure/pkg/validate"
)

type plugin struct {
	log *logrus.Entry
}

var _ api.Plugin = &plugin{}

func NewPlugin(log *logrus.Entry) api.Plugin {
	// always have a logger
	if log == nil {
		logger := logrus.New()
		log = logrus.NewEntry(logger)
	}

	return &plugin{
		log: log,
	}
}

func (p *plugin) ValidateInternal(new, old *acsapi.ContainerService) []error {
	p.log.Info("validating internal data models")
	return validate.ContainerService(new, old)
}

func (p *plugin) GenerateConfig(cs *acsapi.ContainerService) error {
	p.log.Info("generating configs")
	// TODO should we save off the original config here and if there are any errors we can restore it?
	if cs.Config == nil {
		cs.Config = &acsapi.Config{}
	}

	upgrader := config.NewSimpleUpgrader(p.log)
	err := upgrader.Upgrade(cs)
	if err != nil {
		return err
	}

	err = config.Generate(cs)
	if err != nil {
		return err
	}
	return nil
}

func (p *plugin) GenerateARM(cs *acsapi.ContainerService) ([]byte, error) {
	p.log.Info("generating arm templates")
	generator := arm.NewSimpleGenerator(p.log)
	return generator.Generate(cs)
}

func (p *plugin) HealthCheck(ctx context.Context, cs *acsapi.ContainerService) error {
	p.log.Info("starting health check")
	healthChecker := healthcheck.NewSimpleHealthChecker(p.log)
	return healthChecker.HealthCheck(ctx, cs)
}

type MasterUpgrade struct {
	*kubernetes.Clientset
}

var _ api.Upgrade = &MasterUpgrade{}

func (u *MasterUpgrade) IsReady(nodeName string) (bool, error) {
	etcd := fmt.Sprintf("etcd-%s", nodeName)
	etcdPod, err := u.Clientset.CoreV1().Pods("kube-system").Get(etcd, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	apiServer := fmt.Sprintf("api-%s", nodeName)
	apiPod, err := u.Clientset.CoreV1().Pods("kube-system").Get(apiServer, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	controllerManager := fmt.Sprintf("controllers-%s", nodeName)
	cmPod, err := u.Clientset.CoreV1().Pods("kube-system").Get(controllerManager, metav1.GetOptions{})
	if err != nil {
		return false, err
	}

	return isPodReady(etcdPod) && isPodReady(apiPod) && isPodReady(cmPod), nil
}

type InfraUpgrade struct {
	*kubernetes.Clientset
}

var _ api.Upgrade = &InfraUpgrade{}

func (u *InfraUpgrade) IsReady(nodeName string) (bool, error) {
	node, err := u.Clientset.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		return false, nil
	}
	return isNodeReady(node), nil
}

type ComputeUpgrade struct {
	*kubernetes.Clientset
}

var _ api.Upgrade = &ComputeUpgrade{}

func (u *ComputeUpgrade) IsReady(nodeName string) (bool, error) {
	node, err := u.Clientset.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		return false, nil
	}
	return isNodeReady(node), nil
}

func isPodReady(pod *corev1.Pod) bool {
	for _, c := range pod.Status.Conditions {
		if c.Type == corev1.PodReady {
			return c.Status == corev1.ConditionTrue
		}
	}
	return false
}

func isNodeReady(node *corev1.Node) bool {
	for _, c := range node.Status.Conditions {
		if c.Type == corev1.NodeReady {
			return c.Status == corev1.ConditionTrue
		}
	}
	return false
}
