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

package calc_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"

	"github.com/paketo-buildpacks/libjvm/calc"
)

func testMinHeap(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	it("formats", func() {
		Expect(calc.MinHeap{Value: calc.Kibi}.String()).To(Equal("-Xms1K"))
	})

	it("matches -Xms", func() {
		Expect(calc.MatchMinHeap("-Xms1K")).To(BeTrue())
	})

	it("does not match non -Xms", func() {
		Expect(calc.MatchMinHeap("-Xss1K")).To(BeFalse())
	})

	it("parses", func() {
		Expect(calc.ParseMinHeap("-Xms1K")).To(Equal(&calc.MinHeap{Value: calc.Kibi}))
	})

	it("does not parse -Xmx", func() {
		Expect(calc.ParseMinHeap("-Xmx1K")).Error().To(MatchError(
			fmt.Errorf("-Xmx1K does not match min heap pattern %s", calc.MinHeapRE.String())))
	})

	it("does not parse overflown int64", func() {
		Expect(calc.ParseMinHeap("-Xms92233720368547758070k")).Error().To(HaveOccurred())
	})
}
