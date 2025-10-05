package services

import (
	"errors"
	"sync"
	"time"

	"harmancioglue/url-shortener/internal/domain/service"
)

const (
	// Number of bits for each component
	timestampBits = 41
	workerIDBits  = 10
	sequenceBits  = 12

	// Maximum values for each component
	maxWorkerID = -1 ^ (-1 << workerIDBits) // 1023
	maxSequence = -1 ^ (-1 << sequenceBits) // 4095

	workerIDShift  = sequenceBits
	timestampShift = sequenceBits + workerIDBits

	customEpoch = 1609459200000 // January 1, 2021 00:00:00 UTC in milliseconds
)

type SnowflakeIDGenerator struct {
	mu       sync.Mutex
	workerID int64
	sequence int64
	lastTime int64
}

func NewSnowflakeIDGenerator(workerID int64) (service.IDGenerator, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, errors.New("worker ID must be between 0 and 1023")
	}

	return &SnowflakeIDGenerator{
		workerID: workerID,
		sequence: 0,
		lastTime: 0,
	}, nil
}

func (s *SnowflakeIDGenerator) GenerateID() (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentTime := s.currentTimeMillis()

	if currentTime < s.lastTime {
		return 0, errors.New("clock moved backwards")
	}

	if currentTime == s.lastTime {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			currentTime = s.waitNextMillisecond(s.lastTime)
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = currentTime

	id := ((currentTime - customEpoch) << timestampShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return id, nil
}

func (s *SnowflakeIDGenerator) currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}

func (s *SnowflakeIDGenerator) waitNextMillisecond(lastTime int64) int64 {
	currentTime := s.currentTimeMillis()
	for currentTime <= lastTime {
		currentTime = s.currentTimeMillis()
	}
	return currentTime
}
