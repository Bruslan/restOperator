package restoperator

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	appv1alpha1 "github.com/example-inc/rest-operator/pkg/apis/app/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// var restStarted bool = false

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new RestOperator Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileRestOperator{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("restoperator-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}
	// Watch for changes to primary resource RestOperator
	err = c.Watch(&source.Kind{Type: &appv1alpha1.RestOperator{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appv1alpha1.RestOperator{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileRestOperator{}

// ReconcileRestOperator reconciles a RestOperator object
type ReconcileRestOperator struct {
	client client.Client
	scheme *runtime.Scheme
}

func (r *ReconcileRestOperator) Reconcile(request reconcile.Request) (reconcile.Result, error) {

	log.Printf("Reconciling RestOperator %s/%s\n", request.Namespace, request.Name)

	// Fetch the RestOperator instance
	instance := &appv1alpha1.RestOperator{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if the deployment already exists, if not create a new one

	instance.Status.Node = "das ist ein request"

	fmt.Println("neue Instanz wird erzeugt")

	// hier die Azure abfrage
	response, err := http.Get("hier dein Azure Call")

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}

	// hier die Azure response Parameter
	as := &appv1alpha1.RestOperator{
		ObjectMeta: metav1.ObjectMeta{
			Name:      response.Body.name,
			Namespace: "default",
		},
		Status: appv1alpha1.RestOperatorStatus{
			Node: response.Body.ip,
		},
	}
	err = r.client.Create(context.TODO(), as)
	if err != nil {
		fmt.Println("CR existiert bereits")
		return reconcile.Result{}, nil
	}

	//Hier stellst du das Intervall ein
	return reconcile.Result{RequeueAfter: time.Second * 60}, nil
}
