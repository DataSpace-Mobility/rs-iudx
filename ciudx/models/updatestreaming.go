/*
 * IUDX Resource Server APIs
 *
 * The Resource Server is IUDX's data store which allows publication, subscription and discovery of data. For search and discovery, it allows users to search through temporal, geo-based and attribute queries.  For publication and subscription, it allows users to use AMQP streaming protocol over TLS. It enables *Providers* of datasources to publish data as per the IUDX data descriptor. It enables *Consumers* of datasources to search and query for data using HTTPs APIs. It enables *Subscribers* a.k.a [Streaming Consumer] of datasources to stream data using AMQP streaming protocol over TLS.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ModelUpdatestreaming struct {
	Entities []string `json:"entities"`
}