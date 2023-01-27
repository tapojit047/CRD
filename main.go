package main

import (
	"context"
	"flag"
	myv1 "github.com/tapojit047/CRD/pkg/apis/fullmetal.com/v1"
	alchemistclientset "github.com/tapojit047/CRD/pkg/client/clientset/versioned"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crdclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	_ "k8s.io/code-generator"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	crdClient, err := crdclientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	newCRD := &v1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "alchemists.fullmetal.com",
		},
		Spec: v1.CustomResourceDefinitionSpec{
			Group: "fullmetal.com",
			Names: v1.CustomResourceDefinitionNames{
				Kind:     "Alchemist",
				Plural:   "alchemists",
				Singular: "alchemist",
				ShortNames: []string{
					"ac",
				},
			},
			Scope: "Namespaced",
			Versions: []v1.CustomResourceDefinitionVersion{
				{
					Name:    "v1",
					Served:  true,
					Storage: true,
					Schema: &v1.CustomResourceValidation{
						OpenAPIV3Schema: &v1.JSONSchemaProps{
							Type: "object",

							Properties: map[string]v1.JSONSchemaProps{
								"spec": {
									Type: "object",
									Properties: map[string]v1.JSONSchemaProps{
										"name": {
											Type: "string",
										},
										"replicas": {
											Type: "integer",
										},
										"containerPort": {
											Type: "integer",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	_ = crdClient.ApiextensionsV1().CustomResourceDefinitions().Delete(context.TODO(), newCRD.Name, metav1.DeleteOptions{})
	_, err = crdClient.ApiextensionsV1().CustomResourceDefinitions().Create(context.TODO(), newCRD, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	log.Println("Custom Resource Definition, 'Alchemist' Created!")

	log.Println("Press Ctrl+C to create an instance of Alchemist...")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	log.Println("Creating a new alchemist...")
	client, err := alchemistclientset.NewForConfig(config)
	alchemistobject := &myv1.Alchemist{
		ObjectMeta: metav1.ObjectMeta{
			Name: "alchemist",
		},
		Spec: myv1.AlchemistSpec{
			Replicas:      intptr(2),
			Name:          "edward-elric",
			ContainerPort: 8080,
		},
	}
	_, err = client.FullmetalV1().Alchemists("default").Create(context.TODO(), alchemistobject, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	log.Println("New Alchemist Created!!!")

	log.Println("Press CRTL+C to delete...")
	ch = make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	err = client.FullmetalV1().Alchemists("default").Delete(context.TODO(), alchemistobject.Name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}

	err = crdClient.ApiextensionsV1().CustomResourceDefinitions().Delete(context.TODO(), newCRD.Name, metav1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	log.Println("Clean up success")
}

func intptr(i int32) *int32 {
	return &i
}
