// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package operator

const (
	AutoPortForwardFlag                  = "auto-port-forward"
	CADirFlag                            = "ca-dir"
	CACertRotateBeforeFlag               = "ca-cert-rotate-before"
	CACertValidityFlag                   = "ca-cert-validity"
	CertRotateBeforeFlag                 = "cert-rotate-before"
	CertValidityFlag                     = "cert-validity"
	ConfigFlag                           = "config"
	ContainerRegistryFlag                = "container-registry"
	ContainerSuffixFlag                  = "container-suffix"
	DebugHTTPListenFlag                  = "debug-http-listen"
	DisableConfigWatch                   = "disable-config-watch"
	DisableTelemetryFlag                 = "disable-telemetry"
	DistributionChannelFlag              = "distribution-channel"
	ElasticsearchClientTimeout           = "elasticsearch-client-timeout"
	ElasticsearchObservationIntervalFlag = "elasticsearch-observation-interval"
	EnableLeaderElection                 = "enable-leader-election"
	EnableTracingFlag                    = "enable-tracing"
	EnableWebhookFlag                    = "enable-webhook"
	EnforceRBACOnRefsFlag                = "enforce-rbac-on-refs"
	ExposedNodeLabels                    = "exposed-node-labels"
	PasswordHashCacheSize                = "password-hash-cache-size"
	IPFamilyFlag                         = "ip-family"
	KubeClientTimeout                    = "kube-client-timeout"
	KubeClientQPS                        = "kube-client-qps"
	ManageWebhookCertsFlag               = "manage-webhook-certs"
	MaxConcurrentReconcilesFlag          = "max-concurrent-reconciles"
	MetricsPortFlag                      = "metrics-port"
	NamespacesFlag                       = "namespaces"
	OperatorNamespaceFlag                = "operator-namespace"
	SetDefaultSecurityContextFlag        = "set-default-security-context"
	TelemetryIntervalFlag                = "telemetry-interval"
	UBIOnlyFlag                          = "ubi-only"
	ValidateStorageClassFlag             = "validate-storage-class"
	WebhookCertDirFlag                   = "webhook-cert-dir"
	WebhookNameFlag                      = "webhook-name"
	WebhookSecretFlag                    = "webhook-secret"
)
