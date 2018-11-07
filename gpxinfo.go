// Copyright 2013, 2014 Peter Vasil, Tomo Krajina. All
// rights reserved. Use of this source code is governed
// by a BSD-style license that can be found in the
// LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mbecker/gpxgo/gpx"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gpxDirectory := filepath.Join(currentDirectory, "gpx_files")
	files, readDirErr := ioutil.ReadDir(gpxDirectory)
	if readDirErr != nil {
		panic(readDirErr)
	}

	for _, file := range files {
		fmt.Printf("Reading gpx file: %s\n", file.Name())
		gpxFile, gpxParseFileErr := gpx.ParseFile(filepath.Join(gpxDirectory, file.Name()))
		if gpxParseFileErr != nil {
			fmt.Printf("%s error parsing file: %s\n", file.Name(), gpxParseFileErr)
			//continue
		}
		fmt.Println(gpxFile.GetGpxInfo())

		// csvFilename := filepath.Join(currentDirectory, "gpx_files", file.Name()+".csv")

		// err := os.Remove(csvFilename)
		// file, err := os.Create(csvFilename)
		// if err != nil {
		// 	fmt.Println("Cannot create CSV file")
		// 	continue
		// }
		// writer := csv.NewWriter(file)
		// defer writer.Flush()

		// for index := 1; index <= len(distancesAndSpeed); index++ {
		// 	// fmt.Printf("%d - %f - %f\n", index, distancesAndSpeed[index].Distance, distancesAndSpeed[index].Speed)
		// 	number := fmt.Sprintf("%d", index)
		// 	distance := fmt.Sprintf("%f", distancesAndSpeed[index].Distance)
		// 	speed := fmt.Sprintf("%f", distancesAndSpeed[index].Speed)

		// 	sumAllSpeeds += distancesAndSpeed[index].Speed

		// 	x := []string{number, distance, speed}
		// 	err := writer.Write(x)
		// 	if err != nil {
		// 		fmt.Println("Cannot write to CSV file")
		// 		continue
		// 	}
		// }

		// // Proability
		// var sumAllSpeeds float64
		// var populationMeanMu float64
		// var sumXMinusMuPow2 float64
		// var varianceSigmaPow2 float64
		// var varianceSigmaPow float64

		// for index := 1; index <= len(distancesAndSpeed); index++ {
		// 	sumAllSpeeds += distancesAndSpeed[index].Speed
		// }

		// populationMeanMu = sumAllSpeeds / float64(len(distancesAndSpeed))

		// for index := 1; index <= len(distancesAndSpeed); index++ {
		// 	sumXMinusMuPow2 += math.Pow(distancesAndSpeed[index].Speed-populationMeanMu, 2)
		// }

		// varianceSigmaPow2 = sumXMinusMuPow2 / float64(len(distancesAndSpeed))
		// varianceSigmaPow = math.Sqrt(varianceSigmaPow2)

		// fmt.Printf("populationMeanMu: %f\n", populationMeanMu)
		// fmt.Printf("varianceSigmaPow2: %f\n", varianceSigmaPow2)
		// fmt.Printf("varianceSigmaPow: %f\n", varianceSigmaPow)

		// standardDeviationA := populationMeanMu - 2*varianceSigmaPow
		// standardDeviationB := populationMeanMu + 2*varianceSigmaPow
		// fmt.Printf("standardDeviationA: %f\n", standardDeviationA)
		// fmt.Printf("standardDeviationB: %f\n", standardDeviationB)

		// for index := 1; index <= len(distancesAndSpeed); index++ {
		// 	if distancesAndSpeed[index].Speed < standardDeviationA || distancesAndSpeed[index].Speed > standardDeviationB {
		// 		// fmt.Printf("Out ouf range: %d - %f\n", index, distancesAndSpeed[index].Speed)
		// 	}
		// }

		// var sumAllTimeDiffs float64
		// var populationMeanMuTimeDiffs float64
		// var sumXMinusMuPow2TimeDiffs float64
		// var varianceSigmaPow2TimeDiffs float64
		// var varianceSigmaPowTimeDiffs float64

		// var numberOfPoints int

		// var allPoints []NewPointWithDuration
		// var previousPoint gpx.GPXPoint
		// for _, track := range gpxFile.Tracks {
		// 	for _, segment := range track.Segments {
		// 		for pointNumber, point := range segment.Points {
		// 			if pointNumber > 0 {
		// 				previousPoint = segment.Points[pointNumber-1]
		// 				duration := point.TimeDiff(&previousPoint)
		// 				sumAllTimeDiffs += duration

		// 				numberOfPoints++
		// 				newPointWithDuration := NewPointWithDuration{
		// 					point,
		// 					duration,
		// 				}
		// 				allPoints = append(allPoints, newPointWithDuration)
		// 			}
		// 		}
		// 	}
		// }

		// populationMeanMuTimeDiffs = sumAllTimeDiffs / float64(numberOfPoints) // The mean mu for a population series

		// for _, point := range allPoints {
		// 	sumXMinusMuPow2TimeDiffs += math.Pow(point.Duration-populationMeanMuTimeDiffs, 2)
		// }

		// varianceSigmaPow2TimeDiffs = sumXMinusMuPow2TimeDiffs / float64(numberOfPoints)
		// varianceSigmaPowTimeDiffs = math.Sqrt(varianceSigmaPow2TimeDiffs)

		// // µ - 1σ = 68%  (1σ ~ 68.2689492%)
		// // µ - 1.644854σ = 90% (1.644854σ ~ 90%)
		// // µ - 2σ = 95% (2σ ~ 95.4499736%)
		// // µ - 2σ = 99,7% (3σ ~ 99.7300204%)
		// // https://matheguru.com/stochastik/normalverteilung.html
		// // https://en.wikipedia.org/wiki/Standard_deviation
		// var σMultiplier float64 = 1.644854

		// standardDeviationTimeDiffsA := populationMeanMuTimeDiffs - σMultiplier*varianceSigmaPowTimeDiffs
		// standardDeviationTimeDiffsB := populationMeanMuTimeDiffs + σMultiplier*varianceSigmaPowTimeDiffs

		// fmt.Println("--- Duration: ", file.Name())
		// fmt.Printf("sumAllTimeDiffs: %f sec\n", sumAllTimeDiffs) // sumAllTimeDiffsNew of all points in sec
		// t1, _ := time.ParseDuration(fmt.Sprintf("%ds", int64(sumAllTimeDiffs)))
		// fmt.Printf("duration: %s\n", t1) // Duration of all points in sec
		// fmt.Printf("populationMeanMuTimeDiffs: %f\n", populationMeanMuTimeDiffs)
		// fmt.Printf("sumXMinusMuPow2TimeDiffs: %f\n", sumXMinusMuPow2TimeDiffs)
		// fmt.Printf("varianceSigmaPow2TimeDiffs: %f\n", varianceSigmaPow2TimeDiffs)
		// fmt.Printf("varianceSigmaPowTimeDiffs: %f\n", varianceSigmaPowTimeDiffs)
		// fmt.Printf("standardDeviationTimeDiffsA: %f\n", standardDeviationTimeDiffsA)
		// fmt.Printf("standardDeviationTimeDiffsB: %f\n", standardDeviationTimeDiffsB)

		// var sumAllTimeDiffsNew float64
		// var sumAllTimeDiffsDurationGreaterThan float64
		// for _, point := range allPoints {
		// 	if !(point.Duration < standardDeviationTimeDiffsA || point.Duration > standardDeviationTimeDiffsB) {
		// 		sumAllTimeDiffsNew += point.Duration
		// 	}
		// 	if point.Duration < 6.0 {
		// 		sumAllTimeDiffsDurationGreaterThan += point.Duration
		// 	}
		// }
		// fmt.Println("--- Duration New: ", file.Name())
		// fmt.Printf("sumAllTimeDiffsNew: %f sec\n", sumAllTimeDiffsNew) // sumAllTimeDiffsNew of all points in sec
		// t, _ := time.ParseDuration(fmt.Sprintf("%ds", int64(sumAllTimeDiffsNew)))
		// fmt.Printf("duration: %s\n", t) // sumAllTimeDiffsNew of all points in sec

		// fmt.Println("--- Duration New 2:", file.Name())
		// fmt.Printf("sumAllTimeDiffsDurationGreaterThan: %f sec\n", sumAllTimeDiffsDurationGreaterThan) // sumAllTimeDiffsNew of all points in sec
		// t2, _ := time.ParseDuration(fmt.Sprintf("%ds", int64(sumAllTimeDiffsDurationGreaterThan)))
		// fmt.Printf("duration greater than: %s\n", t2) // sumAllTimeDiffsNew of all points in sec
	}

	/* --- Gaussian --- */
	// distribution := gaussian.NewGaussian(1.7, math.Pow(0.1, 2))

	// fmt.Println("CDF: ", distribution.Cdf(1.65))
	// fmt.Println("CDF: ", 1-distribution.Cdf(1.69))

	// P(µ − 2σ ≤ x ≤ µ + 2σ) ≈ 0,9545
	// 1.7-2*0.1 < x
}

type NewPointWithDuration struct {
	gpx.GPXPoint
	Duration float64
}
