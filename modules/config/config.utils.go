package config

import (
	"fmt"
	"strconv"
	"time"
)

type EnvParser struct {
	Err error
}

func (p *EnvParser) Int(val, key string) int {
	if p.Err != nil || val == "" {
		return 0
	}
	res, err := strconv.Atoi(val)
	if err != nil {
		p.Err = fmt.Errorf("unable to parse %s: %w", key, err)
	}
	return res
}
func (p *EnvParser) Float(val, key string) float64 {
	if p.Err != nil || val == "" {
		return 0
	}
	res, err := strconv.ParseFloat(val, 64)
	if err != nil {
		p.Err = fmt.Errorf("unable to parse %s: %w", key, err)
	}
	return res
}
func (p *EnvParser) Duration(val, key string) time.Duration {
	if p.Err != nil || val == "" {
		return 0
	}
	// Parses strings like "1500ms", "5s", "1m" natively!
	res, err := time.ParseDuration(val)
	if err != nil {
		p.Err = fmt.Errorf("unable to parse duration %s: %w", key, err)
	}
	return res
}
