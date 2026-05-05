package probe

import (
	"log"

	"github.com/Haameed/bigip_exporter/pkg/http"
	"github.com/prometheus/client_golang/prometheus"
)

func GetLicenseProbe(c http.BigIPHTTP, target string) ([]prometheus.Metric, bool) {
	var (
		licensedOnDate = prometheus.NewDesc(
			"bigip_system_licensed_on_date",
			"License expiration time (date)",
			[]string{"target", "date"}, nil,
		)
		licensedVersion = prometheus.NewDesc(
			"bigip_system_license_version",
			"License version",
			[]string{"target", "version"}, nil,
		)
		platformId = prometheus.NewDesc(
			"bigip_system_platform_id",
			"platform ID",
			[]string{"target", "platform"}, nil,
		)
	)

	type LicenseInfo struct {
		DailyRenewNotifPeriod StringBased    `json:"dailyRenewNotifPeriod"`
		LicensedOnDate        StringBased    `json:"licensedOnDate"`
		LicensedVersion       StringBased    `json:"licensedVersion"`
		PlatformID            StringBased    `json:"platformId"`
		RegistrationKey       StringBased    `json:"registrationKey"`
		ServiceCheckDate      StringBased    `json:"serviceCheckDate"`
		ActiveModules         map[string]any `json:"-"`
	}

	type NestedEntries struct {
		Entries LicenseInfo `json:"entries"`
	}

	type NestedStats struct {
		NestedStats NestedEntries `json:"nestedStats"`
	}

	type LicenseRespnse struct {
		Kind     string                 `json:"kind"`
		SelfLink string                 `json:"selfLink"`
		Entries  map[string]NestedStats `json:"entries"`
	}

	var licenseResponse LicenseRespnse
	if err := c.Get("/mgmt/tm/sys/license", &licenseResponse); err != nil {
		log.Printf("Error: %v", err)
		return nil, false
	}

	var m []prometheus.Metric
	for _, item := range licenseResponse.Entries {

		m = append(m, prometheus.MustNewConstMetric(licensedOnDate, prometheus.GaugeValue, 1.0, target, item.NestedStats.Entries.LicensedOnDate.Description))
		m = append(m, prometheus.MustNewConstMetric(licensedVersion, prometheus.GaugeValue, 1.0, target, item.NestedStats.Entries.LicensedVersion.Description))
		m = append(m, prometheus.MustNewConstMetric(platformId, prometheus.GaugeValue, 1.0, target, item.NestedStats.Entries.PlatformID.Description))
	}
	return m, true
}
