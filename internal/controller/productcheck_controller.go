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
	"time"

	logging "log"

	eolv1 "github.com/asafdavid23/endoflife-operator/api/v1"
	eolctl "github.com/asafdavid23/eolctl/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// ProductCheckReconciler reconciles a ProductCheck object
type ProductCheckReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=eol.endoflife.io,resources=productchecks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eol.endoflife.io,resources=productchecks/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=eol.endoflife.io,resources=productchecks/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ProductCheck object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *ProductCheckReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// Fetch the ProductCheck instance
	var productCheck eolv1.ProductCheck
	if err := r.Get(ctx, req.NamespacedName, &productCheck); err != nil {
		logger := log.FromContext(ctx)
		logger.Error(err, "unable to fetch ProductCheck")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Call your Go CLI logic here - you might need to run a subprocess or integrate with your existing CLI function.
	// For example, if your CLI tool is invoked using flags, you can replicate that logic in the operator
	isEOL, message, err := eolctl.CheckProductEOL(productCheck.Spec.Name, productCheck.Spec.Version)

	if err != nil {
		return ctrl.Result{}, err
	}

	// Update the status of the CR
	productCheck.Status.IsEOL = isEOL
	productCheck.Status.Message = message
	productCheck.Status.CheckdAt = metav1.Now()

	if err := r.Status().Update(ctx, &productCheck); err != nil {
		return ctrl.Result{}, err
	}

	logging.Printf("ProductCheck %s is updated", productCheck.Spec.Name)
	return ctrl.Result{RequeueAfter: 24 * time.Hour}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProductCheckReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&eolv1.ProductCheck{}).
		Complete(r)
}
