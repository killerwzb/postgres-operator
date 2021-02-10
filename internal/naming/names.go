/*
 Copyright 2021 Crunchy Data Solutions, Inc.
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

package naming

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/rand"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crunchydata/postgres-operator/pkg/apis/postgres-operator.crunchydata.com/v1alpha1"
)

const (
	ContainerDatabase   = "database"
	ContainerPostgreSQL = ContainerDatabase
)

const (
	PortPostgreSQL = "postgres"
)

// AsObjectKey converts the ObjectMeta API type to a client.ObjectKey.
// When you have a client.Object, use client.ObjectKeyFromObject() instead.
func AsObjectKey(m metav1.ObjectMeta) client.ObjectKey {
	return client.ObjectKey{Namespace: m.Namespace, Name: m.Name}
}

// ClusterConfigMap returns the ObjectMeta necessary to lookup
// cluster's shared ConfigMap.
func ClusterConfigMap(cluster *v1alpha1.PostgresCluster) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      cluster.Name + "-config",
	}
}

// ClusterPodService returns the ObjectMeta necessary to lookup the Service
// that is responsible for the network identity of Pods.
func ClusterPodService(cluster *v1alpha1.PostgresCluster) metav1.ObjectMeta {
	// The hyphen below ensures that the DNS name will not be interpreted as a
	// top-level domain. Partially qualified requests for "{pod}.{cluster}-pods"
	// should not leave the Kubernetes cluster, and if they do they are less
	// likely to resolve.
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      cluster.Name + "-pods",
	}
}

// ClusterPrimaryService returns the ObjectMeta necessary to lookup the Service
// that exposes the PostgreSQL primary instance.
func ClusterPrimaryService(cluster *v1alpha1.PostgresCluster) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      cluster.Name + "-primary",
	}
}

// GenerateInstance returns a random name for a member of cluster and set.
func GenerateInstance(
	cluster *v1alpha1.PostgresCluster, set *v1alpha1.PostgresInstanceSetSpec,
) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      cluster.Name + "-" + set.Name + "-" + rand.String(4),
	}
}

// InstanceConfigMap returns the ObjectMeta necessary to lookup
// instance's shared ConfigMap.
func InstanceConfigMap(instance metav1.Object) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: instance.GetNamespace(),
		Name:      instance.GetName() + "-config",
	}
}

// InstanceCertificates returns the ObjectMeta necessary to lookup the Secret
// containing instance's certificates.
func InstanceCertificates(instance metav1.Object) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: instance.GetNamespace(),
		Name:      instance.GetName() + "-certs",
	}
}

// PatroniDistributedConfiguration returns the ObjectMeta necessary to lookup
// the DCS created by Patroni for cluster. This same name is used for both
// ConfigMap and Endpoints. See Patroni DCS "config_path".
func PatroniDistributedConfiguration(cluster *v1alpha1.PostgresCluster) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      PatroniScope(cluster) + "-config",
	}
}

// PatroniLeaderConfigMap returns the ObjectMeta necessary to lookup the
// ConfigMap created by Patroni for the leader election of cluster.
// See Patroni DCS "leader_path".
func PatroniLeaderConfigMap(cluster *v1alpha1.PostgresCluster) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      PatroniScope(cluster) + "-leader",
	}
}

// PatroniLeaderEndpoints returns the ObjectMeta necessary to lookup the
// Endpoints created by Patroni for the leader election of cluster.
// See Patroni DCS "leader_path".
func PatroniLeaderEndpoints(cluster *v1alpha1.PostgresCluster) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      PatroniScope(cluster),
	}
}

// PatroniScope returns the "scope" Patroni uses for cluster.
func PatroniScope(cluster *v1alpha1.PostgresCluster) string {
	return cluster.Name + "-ha"
}

// PatroniTrigger returns the ObjectMeta necessary to lookup the ConfigMap or
// Endpoints Patroni creates for cluster to initiate a controlled change of the
// leader. See Patroni DCS "failover_path".
func PatroniTrigger(cluster *v1alpha1.PostgresCluster) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Namespace: cluster.Namespace,
		Name:      PatroniScope(cluster) + "-failover",
	}
}
