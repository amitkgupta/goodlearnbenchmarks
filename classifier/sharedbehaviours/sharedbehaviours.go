package sharedbehaviours

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/amitkgupta/goodlearn/classifier"
	"github.com/amitkgupta/goodlearn/data/dataset"
)

type ClassifiesAccuratelyAndQuicklyBehaviorInputs struct {
	Classifier              classifier.Classifier
	TrainingData            dataset.Dataset
	TestData                dataset.Dataset
	ExpectedAccuracy        float64
	MinAccuracyThreshold    float64
	MaxSecondsTimeThreshold float64
}

var ClassifiesWithoutError = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func() {
	return func() {
		err := inputs.Classifier.Train(inputs.TrainingData)
		Ω(err).ShouldNot(HaveOccurred())

		for i := 0; i < inputs.TestData.NumRows(); i++ {
			testRow, err := inputs.TestData.Row(i)
			Ω(err).ShouldNot(HaveOccurred())

			_, err = inputs.Classifier.Classify(testRow)
			Ω(err).ShouldNot(HaveOccurred())
		}
	}
}

var ClassifiesWithDeterministicAccuracy = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func() {
	return func() {
		inputs.Classifier.Train(inputs.TrainingData)

		totalCorrect := 0
		for i := 0; i < inputs.TestData.NumRows(); i++ {
			testRow, _ := inputs.TestData.Row(i)

			class, _ := inputs.Classifier.Classify(testRow)
			if class.Equals(testRow.Target()) {
				totalCorrect++
			}
		}
		accuracy := float64(totalCorrect) / float64(inputs.TestData.NumRows())

		Ω(accuracy).Should(BeNumerically("~", inputs.ExpectedAccuracy, 0.001))
	}
}

var ClassifiesSufficientlyAccurately = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func(Benchmarker) {
	return func(b Benchmarker) {
		inputs.Classifier.Train(inputs.TrainingData)

		totalCorrect := 0
		for i := 0; i < inputs.TestData.NumRows(); i++ {
			testRow, _ := inputs.TestData.Row(i)

			class, _ := inputs.Classifier.Classify(testRow)
			if class.Equals(testRow.Target()) {
				totalCorrect++
			}
		}
		accuracy := float64(totalCorrect) / float64(inputs.TestData.NumRows())

		Ω(accuracy).Should(BeNumerically(">", inputs.MinAccuracyThreshold))
	}
}

var ClassifiesSufficientlyQuickly = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func(Benchmarker) {
	return func(b Benchmarker) {
		trainAndClassifyTime := b.Time("train and classify", func() {
			inputs.Classifier.Train(inputs.TrainingData)

			for i := 0; i < inputs.TestData.NumRows(); i++ {
				testRow, _ := inputs.TestData.Row(i)
				inputs.Classifier.Classify(testRow)
			}
		})

		Ω(trainAndClassifyTime.Seconds()).Should(BeNumerically("<", inputs.MaxSecondsTimeThreshold))
	}
}
