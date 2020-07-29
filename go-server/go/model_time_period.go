/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type TimePeriod struct {
	// End date of the catalog
	EndDateTime time.Time `json:"endDateTime,omitempty"`
	// Start date of the catalog
	StartDateTime time.Time `json:"startDateTime,omitempty"`
}