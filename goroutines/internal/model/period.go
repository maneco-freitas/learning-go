package model

import "time"

type Period struct {
	Start time.Time
	End   time.Time
}

func NewPeriod(start, end time.Time) Period {
	return Period{
		Start: start,
		End:   end,
	}
}

func (p *Period) SplitIntoPeriodChunks(numChunks int) []Period {
	periods := make([]Period, numChunks)
	chunkDuration := p.End.Sub(p.Start) / time.Duration(numChunks)
	for i := 0; i < numChunks; i++ {
		periods[i] = Period{
			Start: p.Start.Add(time.Duration(i) * chunkDuration),
			End:   p.Start.Add(time.Duration(i+1) * chunkDuration),
		}
	}
	return periods
}
