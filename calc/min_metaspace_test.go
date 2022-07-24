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

func testMinMetaspace(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect
	)

	it("formats", func() {
		Expect(calc.MinMetaspace{Value: calc.Kibi}.String()).To(Equal("-XX:MetaspaceSize=1K"))
	})

	it("matches -XX:MetaspaceSize", func() {
		Expect(calc.MatchMinMetaspace("-XX:MetaspaceSize=1K")).To(BeTrue())
	})

	it("does not match non -XX:MetaspaceSize", func() {
		Expect(calc.MatchMinMetaspace("-Xss1K")).To(BeFalse())
	})

	it("parses", func() {
		Expect(calc.ParseMinMetaspace("-XX:MetaspaceSize=1K")).To(Equal(&calc.MinMetaspace{Value: calc.Kibi}))
	})

	it("does not parse -XX:MaxMetaspaceSize", func() {
		s := "-XX:MaxMetaspaceSize=1K"
		err := fmt.Errorf("%s does not match metaspace pattern %s", s, calc.MinMetaspaceRE.String())
		Expect(calc.ParseMinMetaspace(s)).Error().To(MatchError(err))
	})

	it("does not parse overflown -XX:MetaspaceSize", func() {
		_, err := calc.ParseMinMetaspace("-XX:MetaspaceSize=92233720368547758070K")
		Expect(err).Error().To(HaveOccurred())
	})
}
