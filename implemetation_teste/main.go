package main

import (
	"context"
	"fmt"
)

func main(coreClient kubernetes.Interface) {

	nsList, err := coreClient.CoreV1().
		Namespaces().
		List(context.Background(), metav1.ListOptions{})
	//checkErr(err)
	fmt.Println(err)

	for _, n := range nsList.Items {
		fmt.Println(n.Name)
	}
}
