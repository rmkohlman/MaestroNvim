// Package plugin provides types and utilities for Neovim plugin management.
package plugin

import (
	"encoding/json"
	"fmt"
	"sort"
)

// HealthChecker runs health checks for plugins.
type HealthChecker struct{}

// NewHealthChecker creates a new HealthChecker.
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{}
}

// StaticCheck performs static health checks that don't require Neovim.
// It validates that health check definitions are well-formed and
// returns reports for plugins that have no checks defined.
func (hc *HealthChecker) StaticCheck(plugins []*Plugin) []*PluginHealthReport {
	var reports []*PluginHealthReport

	for _, p := range plugins {
		report := &PluginHealthReport{
			PluginName: p.Name,
			Enabled:    p.Enabled,
		}

		if !p.Enabled {
			report.Status = HealthStatusSkipped
			reports = append(reports, report)
			continue
		}

		checks := DefaultHealthChecks(p)
		if len(checks) == 0 {
			report.Status = HealthStatusUnknown
			reports = append(reports, report)
			continue
		}

		// Validate check definitions
		allValid := true
		for _, check := range checks {
			result := HealthCheckResult{
				Plugin: p.Name,
				Check:  check,
			}
			if err := ValidateHealthCheckType(string(check.Type)); err != nil {
				result.Status = HealthStatusUnhealthy
				result.Message = err.Error()
				allValid = false
			} else if check.Value == "" {
				result.Status = HealthStatusUnhealthy
				result.Message = "health check value is empty"
				allValid = false
			} else {
				result.Status = HealthStatusUnknown
				result.Message = "requires Neovim to verify"
			}
			report.Results = append(report.Results, result)
		}

		if allValid {
			report.Status = HealthStatusUnknown
		} else {
			report.Status = HealthStatusUnhealthy
		}

		reports = append(reports, report)
	}

	// Sort by plugin name for consistent output
	sort.Slice(reports, func(i, j int) bool {
		return reports[i].PluginName < reports[j].PluginName
	})

	return reports
}

// nvimCheckResult matches the JSON structure output by the Lua health script.
type nvimCheckResult struct {
	Plugin     string `json:"plugin"`
	CheckType  string `json:"check_type"`
	CheckValue string `json:"check_value"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

// ParseNvimResults parses JSON output from the health check Lua script
// and returns structured health reports grouped by plugin.
func (hc *HealthChecker) ParseNvimResults(jsonOutput []byte) ([]*PluginHealthReport, error) {
	var rawResults []nvimCheckResult
	if err := json.Unmarshal(jsonOutput, &rawResults); err != nil {
		return nil, fmt.Errorf("failed to parse health check results: %w", err)
	}

	// Group results by plugin
	reportMap := make(map[string]*PluginHealthReport)
	for _, r := range rawResults {
		report, ok := reportMap[r.Plugin]
		if !ok {
			report = &PluginHealthReport{
				PluginName: r.Plugin,
				Enabled:    true,
				Status:     HealthStatusHealthy,
			}
			reportMap[r.Plugin] = report
		}

		result := HealthCheckResult{
			Plugin: r.Plugin,
			Check: HealthCheck{
				Type:  HealthCheckType(r.CheckType),
				Value: r.CheckValue,
			},
			Status:  HealthStatus(r.Status),
			Message: r.Message,
		}
		report.Results = append(report.Results, result)

		// Update overall status — unhealthy if any check fails
		if result.Status == HealthStatusUnhealthy {
			report.Status = HealthStatusUnhealthy
		}
	}

	// Convert to sorted slice
	var reports []*PluginHealthReport
	for _, r := range reportMap {
		reports = append(reports, r)
	}
	sort.Slice(reports, func(i, j int) bool {
		return reports[i].PluginName < reports[j].PluginName
	})

	return reports, nil
}
