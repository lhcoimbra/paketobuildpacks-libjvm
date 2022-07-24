/*
 * Copyright 2018-2020 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package calc

import (
	"fmt"
	"regexp"
	"strings"
)

var MinHeapRE = regexp.MustCompile(fmt.Sprintf("^-Xms(%s)$", SizePattern))

type MinHeap Size

func (h MinHeap) String() string {
	return fmt.Sprintf("-Xms%s", Size(h))
}

func MatchMinHeap(s string) bool {
	return MinHeapRE.MatchString(strings.TrimSpace(s))
}

func ParseMinHeap(s string) (*MinHeap, error) {
	g := MinHeapRE.FindStringSubmatch(s)
	if g == nil {
		return nil, fmt.Errorf("%s does not match min heap pattern %s", s, MinHeapRE.String())
	}

	z, err := ParseSize(g[1])
	if err != nil {
		return nil, fmt.Errorf("unable to parse min heap size\n%w", err)
	}

	h := MinHeap(z)
	return &h, nil
}
