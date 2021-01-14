/*


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


//https://kubernetes.io/zh/docs/tasks/access-application-cluster/connecting-frontend-backend/ 参考文章
package controllers

import (
	"context"
	"gxin0426/springboot-vue-operator/util"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	devopsv1beta1 "gxin0426/springboot-vue-operator/api/v1beta1"
)

// BackAndFrontReconciler reconciles a BackAndFront object
type BackAndFrontReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

const(
	BackAndFrontFinalizer string = "BackAndFront.finalizers.devops.gxin0426"
)

// +kubebuilder:rbac:groups=devops.gxin0426,resources=backandfronts,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=devops.gxin0426,resources=backandfronts/status,verbs=get;update;patch

func (r *BackAndFrontReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("backandfront", req.NamespacedName)

	bf := &devopsv1beta1.BackAndFront{}

	if err := r.Get(ctx, req.NamespacedName, bf); err != nil {
		if err := client.IgnoreNotFound(err); err == nil {
			log.Info("not found")
			return  ctrl.Result{}, nil
		}else {
			log.Error(err, "未知错误")
			return ctrl.Result{}, err
		}
	}

	if !bf.ObjectMeta.DeletionTimestamp.IsZero(){
		if util.ContainsString(bf.ObjectMeta.Finalizers, BackAndFrontFinalizer) {
			if err := r.
		}
	}

	// your logic here

	return ctrl.Result{}, nil
}

func (r *BackAndFrontReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&devopsv1beta1.BackAndFront{}).
		Complete(r)
}
