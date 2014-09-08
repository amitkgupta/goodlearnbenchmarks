package sharedbehaviours

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"runtime"
	"sync"

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

		runtime.GOMAXPROCS(runtime.NumCPU())

		var wg sync.WaitGroup
		for i := 0; i < inputs.TestData.NumRows(); i++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				testRow, err := inputs.TestData.Row(j)
				Ω(err).ShouldNot(HaveOccurred())

				_, err = inputs.Classifier.Classify(testRow)
				Ω(err).ShouldNot(HaveOccurred())
			}(i)
		}
		wg.Wait()
	}
}

var ClassifiesWithDeterministicAccuracy = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func() {
	return func() {
		inputs.Classifier.Train(inputs.TrainingData)

		runtime.GOMAXPROCS(runtime.NumCPU())

		correctChan := make(chan bool, inputs.TestData.NumRows())
		for i := 0; i < inputs.TestData.NumRows(); i++ {
			go func(j int) {
				testRow, _ := inputs.TestData.Row(j)

				class, _ := inputs.Classifier.Classify(testRow)
				correctChan <- class.Equals(testRow.Target())
			}(i)
		}

		totalCorrect := 0
		for i := 0; i < inputs.TestData.NumRows(); i++ {
			if <-correctChan {
				totalCorrect++
			}
		}

		accuracy := float64(totalCorrect) / float64(inputs.TestData.NumRows())

		Ω(accuracy).Should(BeNumerically("~", inputs.ExpectedAccuracy, 0.001))
	}
}

var ClassifiesSufficientlyAccurately = func(inputs *ClassifiesAccuratelyAndQuicklyBehaviorInputs) func() {
	return func() {
		inputs.Classifier.Train(inputs.TrainingData)

		runtime.GOMAXPROCS(runtime.NumCPU())

		correctChan := make(chan bool, inputs.TestData.NumRows())
		for i := 0; i < inputs.TestData.NumRows(); i++ {
			go func(j int) {
				testRow, _ := inputs.TestData.Row(j)

				class, _ := inputs.Classifier.Classify(testRow)
				correctChan <- class.Equals(testRow.Target())
			}(i)
		}

		totalCorrect := 0
		for i := 0; i < inputs.TestData.NumRows(); i++ {
			if <-correctChan {
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

			runtime.GOMAXPROCS(runtime.NumCPU())

			var wg sync.WaitGroup
			for i := 0; i < inputs.TestData.NumRows(); i++ {
				wg.Add(1)
				go func(j int) {
					defer wg.Done()
					testRow, _ := inputs.TestData.Row(j)
					inputs.Classifier.Classify(testRow)
				}(i)
			}
			wg.Wait()
		})

		Ω(trainAndClassifyTime.Seconds()).Should(BeNumerically("<", inputs.MaxSecondsTimeThreshold))
	}
}
