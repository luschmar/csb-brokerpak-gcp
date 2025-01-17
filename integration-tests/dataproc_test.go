package integration_test

import (
	testframework "github.com/cloudfoundry/cloud-service-broker/brokerpaktestframework"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dataproc", func() {
	const serviceName = "csb-google-dataproc"

	BeforeEach(func() {
		Expect(mockTerraform.SetTFState([]testframework.TFStateValue{})).To(Succeed())
	})

	AfterEach(func() {
		Expect(mockTerraform.Reset()).To(Succeed())
	})

	Describe("provisioning", func() {
		It("should check region constraints", func() {
			_, err := broker.Provision(serviceName, "standard", map[string]any{"region": "-Asia-northeast1"})
			Expect(err).To(MatchError(ContainSubstring("region: Does not match pattern '^[a-z][a-z0-9-]+$'")))
		})
	})

	Describe("updating instance", func() {
		var instanceID string

		BeforeEach(func() {
			var err error
			instanceID, err = broker.Provision(serviceName, "standard", nil)

			Expect(err).NotTo(HaveOccurred())
			Expect(mockTerraform.Reset()).To(Succeed())
		})

		It("should allow updating region because it is not flagged as `prohibit_update` and not specified in the plan", func() {
			err := broker.Update(instanceID, serviceName, "standard", map[string]any{"region": "asia-southeast1"})

			Expect(err).NotTo(HaveOccurred())
		})
	})
})
