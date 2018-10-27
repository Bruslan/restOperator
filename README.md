# Operators

Source: https://github.com/operator-framework/operator-sdk/

### Operators can watch, query, and mutate Kubernetes resources.

# Installation Guide

### To install the operator follow [this](https://github.com/operator-framework/operator-sdk) installation guide

## Operator Logic

### The Go File in pkg/controller/demo/demo_controller.go which is generated for us, is intended for the user to modify with their own Controller business logic.

```go
    func (r \*ReconcileExampleOperator) Reconcile(request reconcile.Request) (reconcile.Result, error) {
    instance := &appv1alpha1.ExampleOperator{}



    // Here the main businesslogic ...

    // Reconcile for any reason than error after 5 seconds
    // Perfect oportunity to make GET Calls this way
    return reconcile.Result{RequeueAfter: time.Second*5}, nil

    }
```

## Reconcile loop types

```go
    // Reconcile successful - don't requeue
    return reconcile.Result{}, nil
    // Reconcile failed due to error - requeue
    return reconcile.Result{}, err
    // Requeue for any reason other than error
    return reconcile.Result{Requeue: true}, nil
```

## Define the spec and status

### Modify the spec and status of the Memcached Custom Resource(CR) at pkg/apis/cache/v1alpha1/demo_types.go:

```go
type DemoSpec struct {
	// Size is the size of the memcached deployment
	Size int32 `json:"size"`
}
type DemoStatus struct {
	// Nodes are the names of the memcached pods
	Nodes []string `json:"nodes"`
}
```

# Example Operator Logic

### This Operator will create a CR every time a new Virtual Machine is Initialized

### To simulate the Initialization of a Virtual Machine, a simple Get request from a Rest API Server is used

```go
    func (r \*ReconcileExampleOperator) Reconcile(request reconcile.Request) (reconcile.Result, error) {
    instance := &appv1alpha1.ExampleOperator{}

    // ... scafolded Logic here

    response, err := http.Get("here_your_api")

    for item in response{

            as := &appv1alpha1.RestOperator{
                ObjectMeta: metav1.ObjectMeta{
                    Name:      item.name,
                    Namespace: "default",
                },
                Status: appv1alpha1.RestOperatorStatus{
                    Node: item.ip,
                },
            }

            err = r.client.Create(context.TODO(), as)
            if err != nil {
                fmt.Println("CR existiert bereits")
            }

    }

    //request every 60 sek
    return reconcile.Result{RequeueAfter: time.Second*60}, nil

    }
```

## Create a Pod

```go
pod := &corev1.Pod{
    ObjectMeta: metav1.ObjectMeta{
        Name:      cr.Name + "-pod",
        Namespace: cr.Namespace,
        Labels:    labels,
    },
    Spec: corev1.PodSpec{
        Containers: []corev1.Container{
            {
                Name:    "busybox",
                Image:   "busybox",
                Command: []string{"sleep", "3600"},
            },
        },
    },
}
err = r.client.Create(context.TODO(), pod)
    if err != nil {
        fmt.Println("Pod existiert bereits")
    }
```

## Create a Custom Recource

```go
cr := &appv1alpha1.RestOperator{
    ObjectMeta: metav1.ObjectMeta{
        Name:      item.name,
        Namespace: "default",
    },
    Status: appv1alpha1.RestOperatorStatus{
        Node: item.ip,
    },
}

err = r.client.Create(context.TODO(), cr)
    if err != nil {
        fmt.Println("CR existiert bereits")
    }
```
