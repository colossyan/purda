/*
Copyright 2025.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	purdav1alpha1 "colossyan.com/purda/api/v1alpha1"

	"colossyan.com/purda/internal/repository"
)

// PurdaReconciler reconciles a Purda object
type PurdaReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	syncer repository.Syncer
}

func New(client client.Client, scheme *runtime.Scheme) *PurdaReconciler {
	config := getConfig()

	return &PurdaReconciler{
		Client: client,
		Scheme: scheme,
		syncer: repository.New(context.Background(), config.Token),
	}
}

// +kubebuilder:rbac:groups=purda.colossyan.com,resources=purda,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=purda.colossyan.com,resources=purda/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=purda.colossyan.com,resources=purda/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Purda object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.20.0/pkg/reconcile
func (r *PurdaReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Enriching context with namespace and name
	ctx = context.WithValue(ctx, "namespace", req.NamespacedName.Namespace)
	ctx = context.WithValue(ctx, "name", req.NamespacedName.Name)

	logger := log.FromContext(ctx)
	logger.Info("reconciling Purda")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PurdaReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&purdav1alpha1.Purda{}).
		Named("purda").
		Complete(r)
}
