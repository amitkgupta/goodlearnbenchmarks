package classifier_test

import (
	"github.com/amitkgupta/goodlearn/classifier/knn"
	"github.com/amitkgupta/goodlearn/csvparse"

	. "github.com/amitkgupta/goodlearnbenchmarks/classifier/sharedbehaviours"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KNN Classifier", func() {
	inputs := ClassifiesAccuratelyAndQuicklyBehaviorInputs{}

	const (
		basicDatasetNumNeighbours           = 1
		basicDatasetExpectedAccuracy        = 0.932
		basicDatasetMaxSecondsTimeThreshold = 0.02
		basicDatasetNumRepetitions          = 10

		manyFeaturesDatasetNumNeighbours           = 1
		manyFeaturesDatasetExpectedAccuracy        = 0.948
		manyFeaturesDatasetMaxSecondsTimeThreshold = 60
		manyFeaturesDatasetNumRepetitions          = 2
	)

	Context("When given a basic dataset", func() {
		BeforeEach(func() {
			trainingData, err := csvparse.DatasetFromPath("datasets/basic_training.csv", 4, 5)
			Ω(err).ShouldNot(HaveOccurred())
			testData, err := csvparse.DatasetFromPath("datasets/basic_test.csv", 4, 5)
			Ω(err).ShouldNot(HaveOccurred())

			inputs.TrainingData = trainingData
			inputs.TestData = testData
			inputs.ExpectedAccuracy = basicDatasetExpectedAccuracy
			inputs.MaxSecondsTimeThreshold = basicDatasetMaxSecondsTimeThreshold

			inputs.Classifier, _ = knn.NewKNNClassifier(basicDatasetNumNeighbours)
		})

		It("classifies without error", ClassifiesWithoutError(&inputs))
		It("classifies with deterministic accuracy", ClassifiesWithDeterministicAccuracy(&inputs))
		Measure("consistently classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), basicDatasetNumRepetitions)
	})

	XContext("When given a dataset with many features", func() {
		BeforeEach(func() {
			trainingData, err := csvparse.DatasetFromPath("datasets/many_features_training.csv", 0, 1)
			Ω(err).ShouldNot(HaveOccurred())
			testData, err := csvparse.DatasetFromPath("datasets/many_features_test.csv", 0, 1)
			Ω(err).ShouldNot(HaveOccurred())

			inputs.TrainingData = trainingData
			inputs.TestData = testData
			inputs.ExpectedAccuracy = manyFeaturesDatasetExpectedAccuracy
			inputs.MaxSecondsTimeThreshold = manyFeaturesDatasetMaxSecondsTimeThreshold

			inputs.Classifier, _ = knn.NewKNNClassifier(manyFeaturesDatasetNumNeighbours)
		})

		It("classifies without error", ClassifiesWithoutError(&inputs))
		It("classifies with deterministic accuracy", ClassifiesWithDeterministicAccuracy(&inputs))
		Measure("consistently classifies sufficiently quickly", ClassifiesSufficientlyQuickly(&inputs), manyFeaturesDatasetNumRepetitions)
	})
})
