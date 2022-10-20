/*
Copyright 2016 The Kubernetes Authors.

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

package privatepathexact

import (
	"log"
	"sort"
	"strings"

	networking "k8s.io/api/networking/v1"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

type path struct {
	r resolver.Resolver
}

// NewParser creates a new whitelist annotation parser
func NewParser(r resolver.Resolver) parser.IngressAnnotation {
	return path{r}
}

// ParseAnnotations parses the annotations contained in the ingress
// rule used to limit access to certain client addresses or networks.
// Multiple ranges can specified using commas as separator
// e.g. `18.0.0.0/8,56.0.0.0/8`
func (a path) Parse(ing *networking.Ingress) (interface{}, error) {

	val, _ := parser.GetStringAnnotation("private-exact-path", ing)

	values := strings.Split(val, ",")

	paths := []string{}
	for _, v := range values {
		paths = append(paths, v)
	}

	sort.Strings(paths)
	log.Println("init started")
	log.Println(paths)
	return paths, nil
}
