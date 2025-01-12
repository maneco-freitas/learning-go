package main

import (
	"context"
	"fmt"
	"goroutines/internal/model"
	"math/rand"
	"sync"

	"time"
)

func fetchDataForPeriodWithRetry(ctx context.Context, period model.Period, maxRetries int) (string, error) {
	for attempt := 0; attempt < maxRetries; attempt++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			data, err := fetchDataForPeriod(period)
			if err == nil {
				return data, nil
			}
			if attempt < maxRetries-1 {
				time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond) // backoff de 1 segundo
				fmt.Printf("Retrying period %v, attempt %d\n", period, attempt+1)
			}
		}
	}
	return "", fmt.Errorf("failed to fetch data after %d attempts", maxRetries)
}

func fetchDataForPeriod(period model.Period) (string, error) {
	if rand.Float32() < 0.2 {
		return "", fmt.Errorf("failed to fetch data")
	}
	if rand.Float32() < 0.1 {
		time.Sleep(1 * time.Second)
	}
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("Date for period: %s - %s",
		period.Start.Format(time.RFC3339), period.End.Format(time.RFC3339)), nil
}

func main() {
	endDate := time.Now()
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	period := model.NewPeriod(startDate, endDate)
	periods := period.SplitIntoPeriodChunks(10)

	maxConcurrent := 5
	maxRetries := 4
	sem := make(chan struct{}, maxConcurrent)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	var wg sync.WaitGroup
	results := make(chan model.Result, len(periods))
	defer cancel()
	for _, period := range periods {
		wg.Add(1)
		sem <- struct{}{}

		go func(p model.Period) {
			defer wg.Done()
			defer func() { <-sem }()
			select {
			case <-ctx.Done():
				results <- model.Result{
					Data:   "",
					Err:    ctx.Err(),
					Period: p,
				}
				fmt.Printf("Timeout occurred for period %s - %s\n", p.Start.Format(time.RFC3339), p.End.Format(time.RFC3339))
				return
			default:
				data, err := fetchDataForPeriodWithRetry(ctx, p, maxRetries)
				if err != nil {
					fmt.Printf("Error fetching data for period %s - %s: %v\n", p.Start.Format(time.RFC3339), p.End.Format(time.RFC3339), err)
					results <- model.Result{
						Data:   "",
						Err:    err,
						Period: p,
					}
					return
				}
				results <- model.Result{
					Data:   data,
					Err:    err,
					Period: p,
				}
			}

		}(period)
	}

	for i := 0; i < len(periods); i++ {
		result := <-results
		fmt.Println(result)
	}

	close(results)

}
