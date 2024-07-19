package main

import (
	"fmt"

	"github.com/goryszewski/operator-test/pkg/apis/v1alpha1"
)

func main() {

	test := v1alpha1.Cluster{}
	fmt.Printf(" %+v \n", test)
}
