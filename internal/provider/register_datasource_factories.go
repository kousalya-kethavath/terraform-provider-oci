// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/oracle/oci-go-sdk/v65/common"
	tf_adm "github.com/oracle/terraform-provider-oci/internal/service/adm"
	tf_ai_data_platform "github.com/oracle/terraform-provider-oci/internal/service/ai_data_platform"
	tf_ai_document "github.com/oracle/terraform-provider-oci/internal/service/ai_document"
	tf_ai_language "github.com/oracle/terraform-provider-oci/internal/service/ai_language"
	tf_ai_vision "github.com/oracle/terraform-provider-oci/internal/service/ai_vision"
	tf_analytics "github.com/oracle/terraform-provider-oci/internal/service/analytics"
	tf_announcements_service "github.com/oracle/terraform-provider-oci/internal/service/announcements_service"
	tf_api_platform "github.com/oracle/terraform-provider-oci/internal/service/api_platform"
	tf_apiaccesscontrol "github.com/oracle/terraform-provider-oci/internal/service/apiaccesscontrol"
	tf_apigateway "github.com/oracle/terraform-provider-oci/internal/service/apigateway"
	tf_apm "github.com/oracle/terraform-provider-oci/internal/service/apm"
	tf_apm_config "github.com/oracle/terraform-provider-oci/internal/service/apm_config"
	tf_apm_synthetics "github.com/oracle/terraform-provider-oci/internal/service/apm_synthetics"
	tf_apm_traces "github.com/oracle/terraform-provider-oci/internal/service/apm_traces"
	tf_appmgmt_control "github.com/oracle/terraform-provider-oci/internal/service/appmgmt_control"
	tf_artifacts "github.com/oracle/terraform-provider-oci/internal/service/artifacts"
	tf_audit "github.com/oracle/terraform-provider-oci/internal/service/audit"
	tf_autoscaling "github.com/oracle/terraform-provider-oci/internal/service/autoscaling"
	tf_bastion "github.com/oracle/terraform-provider-oci/internal/service/bastion"
	tf_batch "github.com/oracle/terraform-provider-oci/internal/service/batch"
	tf_bds "github.com/oracle/terraform-provider-oci/internal/service/bds"
	tf_blockchain "github.com/oracle/terraform-provider-oci/internal/service/blockchain"
	tf_budget "github.com/oracle/terraform-provider-oci/internal/service/budget"
	tf_capacity_management "github.com/oracle/terraform-provider-oci/internal/service/capacity_management"
	tf_certificates_management "github.com/oracle/terraform-provider-oci/internal/service/certificates_management"
	tf_cloud_bridge "github.com/oracle/terraform-provider-oci/internal/service/cloud_bridge"
	tf_cloud_guard "github.com/oracle/terraform-provider-oci/internal/service/cloud_guard"
	tf_cloud_migrations "github.com/oracle/terraform-provider-oci/internal/service/cloud_migrations"
	tf_cluster_health "github.com/oracle/terraform-provider-oci/internal/service/cluster_health"
	tf_cluster_placement_groups "github.com/oracle/terraform-provider-oci/internal/service/cluster_placement_groups"
	tf_compute_cloud_at_customer "github.com/oracle/terraform-provider-oci/internal/service/compute_cloud_at_customer"
	tf_computeinstanceagent "github.com/oracle/terraform-provider-oci/internal/service/computeinstanceagent"
	tf_container_instances "github.com/oracle/terraform-provider-oci/internal/service/container_instances"
	tf_containerengine "github.com/oracle/terraform-provider-oci/internal/service/containerengine"
	tf_core "github.com/oracle/terraform-provider-oci/internal/service/core"
	tf_costad "github.com/oracle/terraform-provider-oci/internal/service/costad"
	tf_data_labeling_service "github.com/oracle/terraform-provider-oci/internal/service/data_labeling_service"
	tf_data_safe "github.com/oracle/terraform-provider-oci/internal/service/data_safe"
	tf_database "github.com/oracle/terraform-provider-oci/internal/service/database"
	tf_database_management "github.com/oracle/terraform-provider-oci/internal/service/database_management"
	tf_database_migration "github.com/oracle/terraform-provider-oci/internal/service/database_migration"
	tf_database_tools "github.com/oracle/terraform-provider-oci/internal/service/database_tools"
	tf_database_tools_runtime "github.com/oracle/terraform-provider-oci/internal/service/database_tools_runtime"
	tf_datacatalog "github.com/oracle/terraform-provider-oci/internal/service/datacatalog"
	tf_datacc "github.com/oracle/terraform-provider-oci/internal/service/datacc"
	tf_dataflow "github.com/oracle/terraform-provider-oci/internal/service/dataflow"
	tf_dataintegration "github.com/oracle/terraform-provider-oci/internal/service/dataintegration"
	tf_datascience "github.com/oracle/terraform-provider-oci/internal/service/datascience"
	tf_dblm "github.com/oracle/terraform-provider-oci/internal/service/dblm"
	tf_dbmulticloud "github.com/oracle/terraform-provider-oci/internal/service/dbmulticloud"
	tf_ddfs "github.com/oracle/terraform-provider-oci/internal/service/ddfs"
	tf_delegate_access_control "github.com/oracle/terraform-provider-oci/internal/service/delegate_access_control"
	tf_demand_signal "github.com/oracle/terraform-provider-oci/internal/service/demand_signal"
	tf_desktops "github.com/oracle/terraform-provider-oci/internal/service/desktops"
	tf_devops "github.com/oracle/terraform-provider-oci/internal/service/devops"
	tf_dif "github.com/oracle/terraform-provider-oci/internal/service/dif"
	tf_disaster_recovery "github.com/oracle/terraform-provider-oci/internal/service/disaster_recovery"
	tf_dns "github.com/oracle/terraform-provider-oci/internal/service/dns"
	tf_email "github.com/oracle/terraform-provider-oci/internal/service/email"
	tf_events "github.com/oracle/terraform-provider-oci/internal/service/events"
	tf_file_storage "github.com/oracle/terraform-provider-oci/internal/service/file_storage"
	tf_fleet_apps_management "github.com/oracle/terraform-provider-oci/internal/service/fleet_apps_management"
	tf_fleet_software_update "github.com/oracle/terraform-provider-oci/internal/service/fleet_software_update"
	tf_functions "github.com/oracle/terraform-provider-oci/internal/service/functions"
	tf_fusion_apps "github.com/oracle/terraform-provider-oci/internal/service/fusion_apps"
	tf_gdp "github.com/oracle/terraform-provider-oci/internal/service/gdp"
	tf_generative_ai "github.com/oracle/terraform-provider-oci/internal/service/generative_ai"
	tf_generative_ai_agent "github.com/oracle/terraform-provider-oci/internal/service/generative_ai_agent"
	tf_generic_artifacts_content "github.com/oracle/terraform-provider-oci/internal/service/generic_artifacts_content"
	tf_golden_gate "github.com/oracle/terraform-provider-oci/internal/service/golden_gate"
	tf_health_checks "github.com/oracle/terraform-provider-oci/internal/service/health_checks"
	tf_identity "github.com/oracle/terraform-provider-oci/internal/service/identity"
	tf_identity_domains "github.com/oracle/terraform-provider-oci/internal/service/identity_domains"
	tf_integration "github.com/oracle/terraform-provider-oci/internal/service/integration"
	tf_iot "github.com/oracle/terraform-provider-oci/internal/service/iot"
	tf_jms "github.com/oracle/terraform-provider-oci/internal/service/jms"
	tf_jms_java_downloads "github.com/oracle/terraform-provider-oci/internal/service/jms_java_downloads"
	tf_jms_utils "github.com/oracle/terraform-provider-oci/internal/service/jms_utils"
	tf_kms "github.com/oracle/terraform-provider-oci/internal/service/kms"
	tf_license_manager "github.com/oracle/terraform-provider-oci/internal/service/license_manager"
	tf_limits "github.com/oracle/terraform-provider-oci/internal/service/limits"
	tf_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/load_balancer"
	tf_log_analytics "github.com/oracle/terraform-provider-oci/internal/service/log_analytics"
	tf_logging "github.com/oracle/terraform-provider-oci/internal/service/logging"
	tf_lustre_file_storage "github.com/oracle/terraform-provider-oci/internal/service/lustre_file_storage"
	tf_managed_kafka "github.com/oracle/terraform-provider-oci/internal/service/managed_kafka"
	tf_management_agent "github.com/oracle/terraform-provider-oci/internal/service/management_agent"
	tf_management_dashboard "github.com/oracle/terraform-provider-oci/internal/service/management_dashboard"
	tf_marketplace "github.com/oracle/terraform-provider-oci/internal/service/marketplace"
	tf_media_services "github.com/oracle/terraform-provider-oci/internal/service/media_services"
	tf_metering_computation "github.com/oracle/terraform-provider-oci/internal/service/metering_computation"
	tf_monitoring "github.com/oracle/terraform-provider-oci/internal/service/monitoring"
	tf_multicloud "github.com/oracle/terraform-provider-oci/internal/service/multicloud"
	tf_mysql "github.com/oracle/terraform-provider-oci/internal/service/mysql"
	tf_network_firewall "github.com/oracle/terraform-provider-oci/internal/service/network_firewall"
	tf_network_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/network_load_balancer"
	tf_nosql "github.com/oracle/terraform-provider-oci/internal/service/nosql"
	tf_objectstorage "github.com/oracle/terraform-provider-oci/internal/service/objectstorage"
	tf_oce "github.com/oracle/terraform-provider-oci/internal/service/oce"
	tf_ocvp "github.com/oracle/terraform-provider-oci/internal/service/ocvp"
	tf_oda "github.com/oracle/terraform-provider-oci/internal/service/oda"
	tf_onesubscription "github.com/oracle/terraform-provider-oci/internal/service/onesubscription"
	tf_ons "github.com/oracle/terraform-provider-oci/internal/service/ons"
	tf_opa "github.com/oracle/terraform-provider-oci/internal/service/opa"
	tf_opensearch "github.com/oracle/terraform-provider-oci/internal/service/opensearch"
	tf_operator_access_control "github.com/oracle/terraform-provider-oci/internal/service/operator_access_control"
	tf_opsi "github.com/oracle/terraform-provider-oci/internal/service/opsi"
	tf_optimizer "github.com/oracle/terraform-provider-oci/internal/service/optimizer"
	tf_os_management_hub "github.com/oracle/terraform-provider-oci/internal/service/os_management_hub"
	tf_osp_gateway "github.com/oracle/terraform-provider-oci/internal/service/osp_gateway"
	tf_osub_billing_schedule "github.com/oracle/terraform-provider-oci/internal/service/osub_billing_schedule"
	tf_osub_organization_subscription "github.com/oracle/terraform-provider-oci/internal/service/osub_organization_subscription"
	tf_osub_subscription "github.com/oracle/terraform-provider-oci/internal/service/osub_subscription"
	tf_osub_usage "github.com/oracle/terraform-provider-oci/internal/service/osub_usage"
	tf_psa "github.com/oracle/terraform-provider-oci/internal/service/psa"
	tf_psql "github.com/oracle/terraform-provider-oci/internal/service/psql"
	tf_queue "github.com/oracle/terraform-provider-oci/internal/service/queue"
	tf_recovery "github.com/oracle/terraform-provider-oci/internal/service/recovery"
	tf_redis "github.com/oracle/terraform-provider-oci/internal/service/redis"
	tf_resource_analytics "github.com/oracle/terraform-provider-oci/internal/service/resource_analytics"
	tf_resource_scheduler "github.com/oracle/terraform-provider-oci/internal/service/resource_scheduler"
	tf_resource_search "github.com/oracle/terraform-provider-oci/internal/service/resource_search"
	tf_resourcemanager "github.com/oracle/terraform-provider-oci/internal/service/resourcemanager"
	tf_sch "github.com/oracle/terraform-provider-oci/internal/service/sch"
	tf_secrets "github.com/oracle/terraform-provider-oci/internal/service/secrets"
	tf_security_attribute "github.com/oracle/terraform-provider-oci/internal/service/security_attribute"
	tf_self "github.com/oracle/terraform-provider-oci/internal/service/self"
	tf_service_catalog "github.com/oracle/terraform-provider-oci/internal/service/service_catalog"
	tf_service_manager_proxy "github.com/oracle/terraform-provider-oci/internal/service/service_manager_proxy"
	tf_stack_monitoring "github.com/oracle/terraform-provider-oci/internal/service/stack_monitoring"
	tf_streaming "github.com/oracle/terraform-provider-oci/internal/service/streaming"
	tf_tenantmanagercontrolplane "github.com/oracle/terraform-provider-oci/internal/service/tenantmanagercontrolplane"
	tf_usage_proxy "github.com/oracle/terraform-provider-oci/internal/service/usage_proxy"
	tf_vault "github.com/oracle/terraform-provider-oci/internal/service/vault"
	tf_vbs_inst "github.com/oracle/terraform-provider-oci/internal/service/vbs_inst"
	tf_visual_builder "github.com/oracle/terraform-provider-oci/internal/service/visual_builder"
	tf_vn_monitoring "github.com/oracle/terraform-provider-oci/internal/service/vn_monitoring"
	tf_vulnerability_scanning "github.com/oracle/terraform-provider-oci/internal/service/vulnerability_scanning"
	tf_waa "github.com/oracle/terraform-provider-oci/internal/service/waa"
	tf_waas "github.com/oracle/terraform-provider-oci/internal/service/waas"
	tf_waf "github.com/oracle/terraform-provider-oci/internal/service/waf"
	tf_zpr "github.com/oracle/terraform-provider-oci/internal/service/zpr"
)

// generatedSDKv2DatasourceFactories returns the provider-owned data-source
// factory inventory. It intentionally does not mutate tfresource's
// process-global registration state.
func generatedSDKv2DatasourceFactories() map[string]func() *schema.Resource {
	factories := make(map[string]func() *schema.Resource)
	if common.CheckForEnabledServices("adm") {
		factories["oci_adm_knowledge_base"] = tf_adm.AdmKnowledgeBaseDataSource
		factories["oci_adm_knowledge_bases"] = tf_adm.AdmKnowledgeBasesDataSource
		factories["oci_adm_remediation_recipe"] = tf_adm.AdmRemediationRecipeDataSource
		factories["oci_adm_remediation_recipes"] = tf_adm.AdmRemediationRecipesDataSource
		factories["oci_adm_remediation_run"] = tf_adm.AdmRemediationRunDataSource
		factories["oci_adm_remediation_run_application_dependency_recommendations"] = tf_adm.AdmRemediationRunApplicationDependencyRecommendationsDataSource
		factories["oci_adm_remediation_run_stage"] = tf_adm.AdmRemediationRunStageDataSource
		factories["oci_adm_remediation_run_stages"] = tf_adm.AdmRemediationRunStagesDataSource
		factories["oci_adm_remediation_runs"] = tf_adm.AdmRemediationRunsDataSource
		factories["oci_adm_vulnerability_audit"] = tf_adm.AdmVulnerabilityAuditDataSource
		factories["oci_adm_vulnerability_audit_application_dependency_vulnerabilities"] = tf_adm.AdmVulnerabilityAuditApplicationDependencyVulnerabilitiesDataSource
		factories["oci_adm_vulnerability_audit_application_dependency_vulnerability"] = tf_adm.AdmVulnerabilityAuditApplicationDependencyVulnerabilityDataSource
		factories["oci_adm_vulnerability_audit_vulnerability"] = tf_adm.AdmVulnerabilityAuditVulnerabilityDataSource
		factories["oci_adm_vulnerability_audits"] = tf_adm.AdmVulnerabilityAuditsDataSource
	}
	if common.CheckForEnabledServices("aidataplatform") {
		factories["oci_ai_data_platform_ai_data_platform"] = tf_ai_data_platform.AiDataPlatformAiDataPlatformDataSource
		factories["oci_ai_data_platform_ai_data_platforms"] = tf_ai_data_platform.AiDataPlatformAiDataPlatformsDataSource
	}
	if common.CheckForEnabledServices("aidocument") {
		factories["oci_ai_document_model"] = tf_ai_document.AiDocumentModelDataSource
		factories["oci_ai_document_model_type"] = tf_ai_document.AiDocumentModelTypeDataSource
		factories["oci_ai_document_models"] = tf_ai_document.AiDocumentModelsDataSource
		factories["oci_ai_document_processor_job"] = tf_ai_document.AiDocumentProcessorJobDataSource
		factories["oci_ai_document_project"] = tf_ai_document.AiDocumentProjectDataSource
		factories["oci_ai_document_projects"] = tf_ai_document.AiDocumentProjectsDataSource
	}
	if common.CheckForEnabledServices("ailanguage") {
		factories["oci_ai_language_endpoint"] = tf_ai_language.AiLanguageEndpointDataSource
		factories["oci_ai_language_endpoints"] = tf_ai_language.AiLanguageEndpointsDataSource
		factories["oci_ai_language_job"] = tf_ai_language.AiLanguageJobDataSource
		factories["oci_ai_language_jobs"] = tf_ai_language.AiLanguageJobsDataSource
		factories["oci_ai_language_model"] = tf_ai_language.AiLanguageModelDataSource
		factories["oci_ai_language_model_evaluation_results"] = tf_ai_language.AiLanguageModelEvaluationResultsDataSource
		factories["oci_ai_language_model_type"] = tf_ai_language.AiLanguageModelTypeDataSource
		factories["oci_ai_language_models"] = tf_ai_language.AiLanguageModelsDataSource
		factories["oci_ai_language_project"] = tf_ai_language.AiLanguageProjectDataSource
		factories["oci_ai_language_projects"] = tf_ai_language.AiLanguageProjectsDataSource
	}
	if common.CheckForEnabledServices("aivision") {
		factories["oci_ai_vision_model"] = tf_ai_vision.AiVisionModelDataSource
		factories["oci_ai_vision_models"] = tf_ai_vision.AiVisionModelsDataSource
		factories["oci_ai_vision_project"] = tf_ai_vision.AiVisionProjectDataSource
		factories["oci_ai_vision_projects"] = tf_ai_vision.AiVisionProjectsDataSource
		factories["oci_ai_vision_stream_group"] = tf_ai_vision.AiVisionStreamGroupDataSource
		factories["oci_ai_vision_stream_groups"] = tf_ai_vision.AiVisionStreamGroupsDataSource
		factories["oci_ai_vision_stream_job"] = tf_ai_vision.AiVisionStreamJobDataSource
		factories["oci_ai_vision_stream_jobs"] = tf_ai_vision.AiVisionStreamJobsDataSource
		factories["oci_ai_vision_stream_source"] = tf_ai_vision.AiVisionStreamSourceDataSource
		factories["oci_ai_vision_stream_sources"] = tf_ai_vision.AiVisionStreamSourcesDataSource
		factories["oci_ai_vision_vision_private_endpoint"] = tf_ai_vision.AiVisionVisionPrivateEndpointDataSource
		factories["oci_ai_vision_vision_private_endpoints"] = tf_ai_vision.AiVisionVisionPrivateEndpointsDataSource
	}
	if common.CheckForEnabledServices("analytics") {
		factories["oci_analytics_analytics_instance"] = tf_analytics.AnalyticsAnalyticsInstanceDataSource
		factories["oci_analytics_analytics_instance_private_access_channel"] = tf_analytics.AnalyticsAnalyticsInstancePrivateAccessChannelDataSource
		factories["oci_analytics_analytics_instance_resource_group"] = tf_analytics.AnalyticsAnalyticsInstanceResourceGroupDataSource
		factories["oci_analytics_analytics_instance_resource_groups"] = tf_analytics.AnalyticsAnalyticsInstanceResourceGroupsDataSource
		factories["oci_analytics_analytics_instances"] = tf_analytics.AnalyticsAnalyticsInstancesDataSource
	}
	if common.CheckForEnabledServices("announcementsservice") {
		factories["oci_announcements_service_announcement_subscription"] = tf_announcements_service.AnnouncementsServiceAnnouncementSubscriptionDataSource
		factories["oci_announcements_service_announcement_subscriptions"] = tf_announcements_service.AnnouncementsServiceAnnouncementSubscriptionsDataSource
		factories["oci_announcements_service_services"] = tf_announcements_service.AnnouncementsServiceServicesDataSource
	}
	if common.CheckForEnabledServices("apiplatform") {
		factories["oci_api_platform_api_platform_instance"] = tf_api_platform.ApiPlatformApiPlatformInstanceDataSource
		factories["oci_api_platform_api_platform_instances"] = tf_api_platform.ApiPlatformApiPlatformInstancesDataSource
	}
	if common.CheckForEnabledServices("apiaccesscontrol") {
		factories["oci_apiaccesscontrol_api_metadata"] = tf_apiaccesscontrol.ApiaccesscontrolApiMetadataDataSource
		factories["oci_apiaccesscontrol_api_metadata_by_entity_types"] = tf_apiaccesscontrol.ApiaccesscontrolApiMetadataByEntityTypesDataSource
		factories["oci_apiaccesscontrol_api_metadatas"] = tf_apiaccesscontrol.ApiaccesscontrolApiMetadatasDataSource
		factories["oci_apiaccesscontrol_privileged_api_control"] = tf_apiaccesscontrol.ApiaccesscontrolPrivilegedApiControlDataSource
		factories["oci_apiaccesscontrol_privileged_api_controls"] = tf_apiaccesscontrol.ApiaccesscontrolPrivilegedApiControlsDataSource
		factories["oci_apiaccesscontrol_privileged_api_request"] = tf_apiaccesscontrol.ApiaccesscontrolPrivilegedApiRequestDataSource
		factories["oci_apiaccesscontrol_privileged_api_requests"] = tf_apiaccesscontrol.ApiaccesscontrolPrivilegedApiRequestsDataSource
	}
	if common.CheckForEnabledServices("apigateway") {
		factories["oci_apigateway_api"] = tf_apigateway.ApigatewayApiDataSource
		factories["oci_apigateway_api_content"] = tf_apigateway.ApigatewayApiContentDataSource
		factories["oci_apigateway_api_deployment_specification"] = tf_apigateway.ApigatewayApiDeploymentSpecificationDataSource
		factories["oci_apigateway_api_validation"] = tf_apigateway.ApigatewayApiValidationDataSource
		factories["oci_apigateway_apis"] = tf_apigateway.ApigatewayApisDataSource
		factories["oci_apigateway_certificate"] = tf_apigateway.ApigatewayCertificateDataSource
		factories["oci_apigateway_certificates"] = tf_apigateway.ApigatewayCertificatesDataSource
		factories["oci_apigateway_deployment"] = tf_apigateway.ApigatewayDeploymentDataSource
		factories["oci_apigateway_deployments"] = tf_apigateway.ApigatewayDeploymentsDataSource
		factories["oci_apigateway_gateway"] = tf_apigateway.ApigatewayGatewayDataSource
		factories["oci_apigateway_gateways"] = tf_apigateway.ApigatewayGatewaysDataSource
		factories["oci_apigateway_subscriber"] = tf_apigateway.ApigatewaySubscriberDataSource
		factories["oci_apigateway_subscribers"] = tf_apigateway.ApigatewaySubscribersDataSource
		factories["oci_apigateway_usage_plan"] = tf_apigateway.ApigatewayUsagePlanDataSource
		factories["oci_apigateway_usage_plans"] = tf_apigateway.ApigatewayUsagePlansDataSource
	}
	if common.CheckForEnabledServices("apm") {
		factories["oci_apm_apm_domain"] = tf_apm.ApmApmDomainDataSource
		factories["oci_apm_apm_domains"] = tf_apm.ApmApmDomainsDataSource
		factories["oci_apm_data_keys"] = tf_apm.ApmDataKeysDataSource
	}
	if common.CheckForEnabledServices("apmconfig") {
		factories["oci_apm_config_config"] = tf_apm_config.ApmConfigConfigDataSource
		factories["oci_apm_config_configs"] = tf_apm_config.ApmConfigConfigsDataSource
		factories["oci_apm_config_data_file"] = tf_apm_config.ApmConfigDataFileDataSource
		factories["oci_apm_config_data_files"] = tf_apm_config.ApmConfigDataFilesDataSource
	}
	if common.CheckForEnabledServices("apmsynthetics") {
		factories["oci_apm_synthetics_dedicated_vantage_point"] = tf_apm_synthetics.ApmSyntheticsDedicatedVantagePointDataSource
		factories["oci_apm_synthetics_dedicated_vantage_points"] = tf_apm_synthetics.ApmSyntheticsDedicatedVantagePointsDataSource
		factories["oci_apm_synthetics_monitor"] = tf_apm_synthetics.ApmSyntheticsMonitorDataSource
		factories["oci_apm_synthetics_monitors"] = tf_apm_synthetics.ApmSyntheticsMonitorsDataSource
		factories["oci_apm_synthetics_on_premise_vantage_point"] = tf_apm_synthetics.ApmSyntheticsOnPremiseVantagePointDataSource
		factories["oci_apm_synthetics_on_premise_vantage_point_worker"] = tf_apm_synthetics.ApmSyntheticsOnPremiseVantagePointWorkerDataSource
		factories["oci_apm_synthetics_on_premise_vantage_point_workers"] = tf_apm_synthetics.ApmSyntheticsOnPremiseVantagePointWorkersDataSource
		factories["oci_apm_synthetics_on_premise_vantage_points"] = tf_apm_synthetics.ApmSyntheticsOnPremiseVantagePointsDataSource
		factories["oci_apm_synthetics_public_vantage_point"] = tf_apm_synthetics.ApmSyntheticsPublicVantagePointDataSource
		factories["oci_apm_synthetics_public_vantage_points"] = tf_apm_synthetics.ApmSyntheticsPublicVantagePointsDataSource
		factories["oci_apm_synthetics_result"] = tf_apm_synthetics.ApmSyntheticsResultDataSource
		factories["oci_apm_synthetics_script"] = tf_apm_synthetics.ApmSyntheticsScriptDataSource
		factories["oci_apm_synthetics_scripts"] = tf_apm_synthetics.ApmSyntheticsScriptsDataSource
	}
	if common.CheckForEnabledServices("apmtraces") {
		factories["oci_apm_traces_attribute_auto_activate_status"] = tf_apm_traces.ApmTracesAttributeAutoActivateStatusDataSource
		factories["oci_apm_traces_log"] = tf_apm_traces.ApmTracesLogDataSource
		factories["oci_apm_traces_query_quick_picks"] = tf_apm_traces.ApmTracesQueryQuickPicksDataSource
		factories["oci_apm_traces_scheduled_queries"] = tf_apm_traces.ApmTracesScheduledQueriesDataSource
		factories["oci_apm_traces_scheduled_query"] = tf_apm_traces.ApmTracesScheduledQueryDataSource
		factories["oci_apm_traces_trace"] = tf_apm_traces.ApmTracesTraceDataSource
		factories["oci_apm_traces_trace_aggregated_snapshot_data"] = tf_apm_traces.ApmTracesTraceAggregatedSnapshotDataDataSource
		factories["oci_apm_traces_trace_snapshot_data"] = tf_apm_traces.ApmTracesTraceSnapshotDataDataSource
	}
	if common.CheckForEnabledServices("appmgmtcontrol") {
		factories["oci_appmgmt_control_monitored_instance"] = tf_appmgmt_control.AppmgmtControlMonitoredInstanceDataSource
		factories["oci_appmgmt_control_monitored_instances"] = tf_appmgmt_control.AppmgmtControlMonitoredInstancesDataSource
	}
	if common.CheckForEnabledServices("artifacts") {
		factories["oci_artifacts_container_configuration"] = tf_artifacts.ArtifactsContainerConfigurationDataSource
		factories["oci_artifacts_container_image"] = tf_artifacts.ArtifactsContainerImageDataSource
		factories["oci_artifacts_container_image_signature"] = tf_artifacts.ArtifactsContainerImageSignatureDataSource
		factories["oci_artifacts_container_image_signatures"] = tf_artifacts.ArtifactsContainerImageSignaturesDataSource
		factories["oci_artifacts_container_images"] = tf_artifacts.ArtifactsContainerImagesDataSource
		factories["oci_artifacts_container_repositories"] = tf_artifacts.ArtifactsContainerRepositoriesDataSource
		factories["oci_artifacts_container_repository"] = tf_artifacts.ArtifactsContainerRepositoryDataSource
		factories["oci_artifacts_generic_artifact"] = tf_artifacts.ArtifactsGenericArtifactDataSource
		factories["oci_artifacts_generic_artifacts"] = tf_artifacts.ArtifactsGenericArtifactsDataSource
		factories["oci_artifacts_repositories"] = tf_artifacts.ArtifactsRepositoriesDataSource
		factories["oci_artifacts_repository"] = tf_artifacts.ArtifactsRepositoryDataSource
	}
	if common.CheckForEnabledServices("audit") {
		factories["oci_audit_configuration"] = tf_audit.AuditConfigurationDataSource
		factories["oci_audit_events"] = tf_audit.AuditAuditEventsDataSource
	}
	if common.CheckForEnabledServices("autoscaling") {
		factories["oci_autoscaling_auto_scaling_configuration"] = tf_autoscaling.AutoScalingAutoScalingConfigurationDataSource
		factories["oci_autoscaling_auto_scaling_configurations"] = tf_autoscaling.AutoScalingAutoScalingConfigurationsDataSource
	}
	if common.CheckForEnabledServices("bastion") {
		factories["oci_bastion_bastion"] = tf_bastion.BastionBastionDataSource
		factories["oci_bastion_bastions"] = tf_bastion.BastionBastionsDataSource
		factories["oci_bastion_session"] = tf_bastion.BastionSessionDataSource
		factories["oci_bastion_sessions"] = tf_bastion.BastionSessionsDataSource
	}
	if common.CheckForEnabledServices("batch") {
		factories["oci_batch_batch_context"] = tf_batch.BatchBatchContextDataSource
		factories["oci_batch_batch_context_shapes"] = tf_batch.BatchBatchContextShapesDataSource
		factories["oci_batch_batch_contexts"] = tf_batch.BatchBatchContextsDataSource
		factories["oci_batch_batch_job_pool"] = tf_batch.BatchBatchJobPoolDataSource
		factories["oci_batch_batch_job_pools"] = tf_batch.BatchBatchJobPoolsDataSource
		factories["oci_batch_batch_task_environment"] = tf_batch.BatchBatchTaskEnvironmentDataSource
		factories["oci_batch_batch_task_environments"] = tf_batch.BatchBatchTaskEnvironmentsDataSource
		factories["oci_batch_batch_task_profile"] = tf_batch.BatchBatchTaskProfileDataSource
		factories["oci_batch_batch_task_profiles"] = tf_batch.BatchBatchTaskProfilesDataSource
	}
	if common.CheckForEnabledServices("bds") {
		factories["oci_bds_auto_scaling_configuration"] = tf_bds.BdsAutoScalingConfigurationDataSource
		factories["oci_bds_auto_scaling_configurations"] = tf_bds.BdsAutoScalingConfigurationsDataSource
		factories["oci_bds_bds_capacity_reservation"] = tf_bds.BdsBdsCapacityReservationDataSource
		factories["oci_bds_bds_capacity_reservation_associated_configurations"] = tf_bds.BdsBdsCapacityReservationAssociatedConfigurationsDataSource
		factories["oci_bds_bds_capacity_reservations"] = tf_bds.BdsBdsCapacityReservationsDataSource
		factories["oci_bds_bds_cluster_versions"] = tf_bds.BdsBdsClusterVersionsDataSource
		factories["oci_bds_bds_instance"] = tf_bds.BdsBdsInstanceDataSource
		factories["oci_bds_bds_instance_api_key"] = tf_bds.BdsBdsInstanceApiKeyDataSource
		factories["oci_bds_bds_instance_api_keys"] = tf_bds.BdsBdsInstanceApiKeysDataSource
		factories["oci_bds_bds_instance_bds_capacity_reservation_configuration"] = tf_bds.BdsBdsInstanceBdsCapacityReservationConfigurationDataSource
		factories["oci_bds_bds_instance_bds_capacity_reservation_configurations"] = tf_bds.BdsBdsInstanceBdsCapacityReservationConfigurationsDataSource
		factories["oci_bds_bds_instance_bds_certificate_configuration"] = tf_bds.BdsBdsInstanceBdsCertificateConfigurationDataSource
		factories["oci_bds_bds_instance_bds_certificate_configurations"] = tf_bds.BdsBdsInstanceBdsCertificateConfigurationsDataSource
		factories["oci_bds_bds_instance_get_os_patch"] = tf_bds.BdsBdsInstanceGetOsPatchDataSource
		factories["oci_bds_bds_instance_identity_configuration"] = tf_bds.BdsBdsInstanceIdentityConfigurationDataSource
		factories["oci_bds_bds_instance_identity_configurations"] = tf_bds.BdsBdsInstanceIdentityConfigurationsDataSource
		factories["oci_bds_bds_instance_list_os_patches"] = tf_bds.BdsBdsInstanceListOsPatchesDataSource
		factories["oci_bds_bds_instance_metastore_config"] = tf_bds.BdsBdsInstanceMetastoreConfigDataSource
		factories["oci_bds_bds_instance_metastore_configs"] = tf_bds.BdsBdsInstanceMetastoreConfigsDataSource
		factories["oci_bds_bds_instance_node_backup"] = tf_bds.BdsBdsInstanceNodeBackupDataSource
		factories["oci_bds_bds_instance_node_backup_configuration"] = tf_bds.BdsBdsInstanceNodeBackupConfigurationDataSource
		factories["oci_bds_bds_instance_node_backup_configurations"] = tf_bds.BdsBdsInstanceNodeBackupConfigurationsDataSource
		factories["oci_bds_bds_instance_node_backups"] = tf_bds.BdsBdsInstanceNodeBackupsDataSource
		factories["oci_bds_bds_instance_node_replace_configuration"] = tf_bds.BdsBdsInstanceNodeReplaceConfigurationDataSource
		factories["oci_bds_bds_instance_node_replace_configurations"] = tf_bds.BdsBdsInstanceNodeReplaceConfigurationsDataSource
		factories["oci_bds_bds_instance_patch_histories"] = tf_bds.BdsBdsInstancePatchHistoriesDataSource
		factories["oci_bds_bds_instance_patches"] = tf_bds.BdsBdsInstancePatchesDataSource
		factories["oci_bds_bds_instance_resource_principal_configuration"] = tf_bds.BdsBdsInstanceResourcePrincipalConfigurationDataSource
		factories["oci_bds_bds_instance_resource_principal_configurations"] = tf_bds.BdsBdsInstanceResourcePrincipalConfigurationsDataSource
		factories["oci_bds_bds_instance_software_update"] = tf_bds.BdsBdsInstanceSoftwareUpdateDataSource
		factories["oci_bds_bds_instance_software_updates"] = tf_bds.BdsBdsInstanceSoftwareUpdatesDataSource
		factories["oci_bds_bds_instances"] = tf_bds.BdsBdsInstancesDataSource
	}
	if common.CheckForEnabledServices("blockchain") {
		factories["oci_blockchain_blockchain_platform"] = tf_blockchain.BlockchainBlockchainPlatformDataSource
		factories["oci_blockchain_blockchain_platform_patches"] = tf_blockchain.BlockchainBlockchainPlatformPatchesDataSource
		factories["oci_blockchain_blockchain_platforms"] = tf_blockchain.BlockchainBlockchainPlatformsDataSource
		factories["oci_blockchain_osn"] = tf_blockchain.BlockchainOsnDataSource
		factories["oci_blockchain_osns"] = tf_blockchain.BlockchainOsnsDataSource
		factories["oci_blockchain_peer"] = tf_blockchain.BlockchainPeerDataSource
		factories["oci_blockchain_peers"] = tf_blockchain.BlockchainPeersDataSource
	}
	if common.CheckForEnabledServices("budget") {
		factories["oci_budget_alert_rule"] = tf_budget.BudgetAlertRuleDataSource
		factories["oci_budget_alert_rules"] = tf_budget.BudgetAlertRulesDataSource
		factories["oci_budget_budget"] = tf_budget.BudgetBudgetDataSource
		factories["oci_budget_budgets"] = tf_budget.BudgetBudgetsDataSource
		factories["oci_budget_cost_alert_subscription"] = tf_budget.BudgetCostAlertSubscriptionDataSource
		factories["oci_budget_cost_alert_subscriptions"] = tf_budget.BudgetCostAlertSubscriptionsDataSource
		factories["oci_budget_cost_anomaly_event"] = tf_budget.BudgetCostAnomalyEventDataSource
		factories["oci_budget_cost_anomaly_event_analytics"] = tf_budget.BudgetCostAnomalyEventAnalyticsDataSource
		factories["oci_budget_cost_anomaly_events"] = tf_budget.BudgetCostAnomalyEventsDataSource
		factories["oci_budget_cost_anomaly_monitor"] = tf_budget.BudgetCostAnomalyMonitorDataSource
		factories["oci_budget_cost_anomaly_monitors"] = tf_budget.BudgetCostAnomalyMonitorsDataSource
	}
	if common.CheckForEnabledServices("capacitymanagement") {
		factories["oci_capacity_management_internal_namespace_occ_overviews"] = tf_capacity_management.CapacityManagementInternalNamespaceOccOverviewsDataSource
		factories["oci_capacity_management_internal_occ_availability_catalogs"] = tf_capacity_management.CapacityManagementInternalOccAvailabilityCatalogsDataSource
		factories["oci_capacity_management_internal_occ_handover_resource_block_details"] = tf_capacity_management.CapacityManagementInternalOccHandoverResourceBlockDetailsDataSource
		factories["oci_capacity_management_internal_occ_handover_resource_blocks"] = tf_capacity_management.CapacityManagementInternalOccHandoverResourceBlocksDataSource
		factories["oci_capacity_management_internal_occm_demand_signal"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalDataSource
		factories["oci_capacity_management_internal_occm_demand_signal_catalog"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalCatalogDataSource
		factories["oci_capacity_management_internal_occm_demand_signal_catalog_resources"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalCatalogResourcesDataSource
		factories["oci_capacity_management_internal_occm_demand_signal_catalogs"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalCatalogsDataSource
		factories["oci_capacity_management_internal_occm_demand_signal_deliveries"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalDeliveriesDataSource
		factories["oci_capacity_management_internal_occm_demand_signal_delivery"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalDeliveryDataSource
		factories["oci_capacity_management_internal_occm_demand_signal_items"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalItemsDataSource
		factories["oci_capacity_management_internal_occm_demand_signals"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalsDataSource
		factories["oci_capacity_management_namespace_occ_overviews"] = tf_capacity_management.CapacityManagementNamespaceOccOverviewsDataSource
		factories["oci_capacity_management_occ_availability_catalog"] = tf_capacity_management.CapacityManagementOccAvailabilityCatalogDataSource
		factories["oci_capacity_management_occ_availability_catalog_content"] = tf_capacity_management.CapacityManagementOccAvailabilityCatalogContentDataSource
		factories["oci_capacity_management_occ_availability_catalog_occ_availabilities"] = tf_capacity_management.CapacityManagementOccAvailabilityCatalogOccAvailabilitiesDataSource
		factories["oci_capacity_management_occ_availability_catalogs"] = tf_capacity_management.CapacityManagementOccAvailabilityCatalogsDataSource
		factories["oci_capacity_management_occ_capacity_request"] = tf_capacity_management.CapacityManagementOccCapacityRequestDataSource
		factories["oci_capacity_management_occ_capacity_requests"] = tf_capacity_management.CapacityManagementOccCapacityRequestsDataSource
		factories["oci_capacity_management_occ_customer_group"] = tf_capacity_management.CapacityManagementOccCustomerGroupDataSource
		factories["oci_capacity_management_occ_customer_groups"] = tf_capacity_management.CapacityManagementOccCustomerGroupsDataSource
		factories["oci_capacity_management_occ_handover_resource_block_details"] = tf_capacity_management.CapacityManagementOccHandoverResourceBlockDetailsDataSource
		factories["oci_capacity_management_occ_handover_resource_blocks"] = tf_capacity_management.CapacityManagementOccHandoverResourceBlocksDataSource
		factories["oci_capacity_management_occm_demand_signal"] = tf_capacity_management.CapacityManagementOccmDemandSignalDataSource
		factories["oci_capacity_management_occm_demand_signal_catalog_resources"] = tf_capacity_management.CapacityManagementOccmDemandSignalCatalogResourcesDataSource
		factories["oci_capacity_management_occm_demand_signal_deliveries"] = tf_capacity_management.CapacityManagementOccmDemandSignalDeliveriesDataSource
		factories["oci_capacity_management_occm_demand_signal_item"] = tf_capacity_management.CapacityManagementOccmDemandSignalItemDataSource
		factories["oci_capacity_management_occm_demand_signal_items"] = tf_capacity_management.CapacityManagementOccmDemandSignalItemsDataSource
		factories["oci_capacity_management_occm_demand_signals"] = tf_capacity_management.CapacityManagementOccmDemandSignalsDataSource
	}
	if common.CheckForEnabledServices("certificatesmanagement") {
		factories["oci_certificates_management_association"] = tf_certificates_management.CertificatesManagementAssociationDataSource
		factories["oci_certificates_management_associations"] = tf_certificates_management.CertificatesManagementAssociationsDataSource
		factories["oci_certificates_management_ca_bundle"] = tf_certificates_management.CertificatesManagementCaBundleDataSource
		factories["oci_certificates_management_ca_bundles"] = tf_certificates_management.CertificatesManagementCaBundlesDataSource
		factories["oci_certificates_management_certificate"] = tf_certificates_management.CertificatesManagementCertificateDataSource
		factories["oci_certificates_management_certificate_authorities"] = tf_certificates_management.CertificatesManagementCertificateAuthoritiesDataSource
		factories["oci_certificates_management_certificate_authority"] = tf_certificates_management.CertificatesManagementCertificateAuthorityDataSource
		factories["oci_certificates_management_certificate_authority_version"] = tf_certificates_management.CertificatesManagementCertificateAuthorityVersionDataSource
		factories["oci_certificates_management_certificate_authority_versions"] = tf_certificates_management.CertificatesManagementCertificateAuthorityVersionsDataSource
		factories["oci_certificates_management_certificate_version"] = tf_certificates_management.CertificatesManagementCertificateVersionDataSource
		factories["oci_certificates_management_certificate_versions"] = tf_certificates_management.CertificatesManagementCertificateVersionsDataSource
		factories["oci_certificates_management_certificates"] = tf_certificates_management.CertificatesManagementCertificatesDataSource
	}
	if common.CheckForEnabledServices("cloudbridge") {
		factories["oci_cloud_bridge_agent"] = tf_cloud_bridge.CloudBridgeAgentDataSource
		factories["oci_cloud_bridge_agent_dependencies"] = tf_cloud_bridge.CloudBridgeAgentDependenciesDataSource
		factories["oci_cloud_bridge_agent_dependency"] = tf_cloud_bridge.CloudBridgeAgentDependencyDataSource
		factories["oci_cloud_bridge_agent_plugin"] = tf_cloud_bridge.CloudBridgeAgentPluginDataSource
		factories["oci_cloud_bridge_agents"] = tf_cloud_bridge.CloudBridgeAgentsDataSource
		factories["oci_cloud_bridge_appliance_image"] = tf_cloud_bridge.CloudBridgeApplianceImageDataSource
		factories["oci_cloud_bridge_appliance_images"] = tf_cloud_bridge.CloudBridgeApplianceImagesDataSource
		factories["oci_cloud_bridge_asset"] = tf_cloud_bridge.CloudBridgeAssetDataSource
		factories["oci_cloud_bridge_asset_source"] = tf_cloud_bridge.CloudBridgeAssetSourceDataSource
		factories["oci_cloud_bridge_asset_sources"] = tf_cloud_bridge.CloudBridgeAssetSourcesDataSource
		factories["oci_cloud_bridge_assets"] = tf_cloud_bridge.CloudBridgeAssetsDataSource
		factories["oci_cloud_bridge_discovery_schedule"] = tf_cloud_bridge.CloudBridgeDiscoveryScheduleDataSource
		factories["oci_cloud_bridge_discovery_schedules"] = tf_cloud_bridge.CloudBridgeDiscoverySchedulesDataSource
		factories["oci_cloud_bridge_environment"] = tf_cloud_bridge.CloudBridgeEnvironmentDataSource
		factories["oci_cloud_bridge_environments"] = tf_cloud_bridge.CloudBridgeEnvironmentsDataSource
		factories["oci_cloud_bridge_inventories"] = tf_cloud_bridge.CloudBridgeInventoriesDataSource
		factories["oci_cloud_bridge_inventory"] = tf_cloud_bridge.CloudBridgeInventoryDataSource
		factories["oci_cloud_bridge_supported_cloud_regions"] = tf_cloud_bridge.CloudBridgeSupportedCloudRegionsDataSource
	}
	if common.CheckForEnabledServices("cloudguard") {
		factories["oci_cloud_guard_adhoc_queries"] = tf_cloud_guard.CloudGuardAdhocQueriesDataSource
		factories["oci_cloud_guard_adhoc_query"] = tf_cloud_guard.CloudGuardAdhocQueryDataSource
		factories["oci_cloud_guard_cloud_guard_configuration"] = tf_cloud_guard.CloudGuardCloudGuardConfigurationDataSource
		factories["oci_cloud_guard_data_mask_rule"] = tf_cloud_guard.CloudGuardDataMaskRuleDataSource
		factories["oci_cloud_guard_data_mask_rules"] = tf_cloud_guard.CloudGuardDataMaskRulesDataSource
		factories["oci_cloud_guard_data_source"] = tf_cloud_guard.CloudGuardDataSourceDataSource
		factories["oci_cloud_guard_data_source_event"] = tf_cloud_guard.CloudGuardDataSourceEventDataSource
		factories["oci_cloud_guard_data_source_events"] = tf_cloud_guard.CloudGuardDataSourceEventsDataSource
		factories["oci_cloud_guard_data_sources"] = tf_cloud_guard.CloudGuardDataSourcesDataSource
		factories["oci_cloud_guard_detector_recipe"] = tf_cloud_guard.CloudGuardDetectorRecipeDataSource
		factories["oci_cloud_guard_detector_recipes"] = tf_cloud_guard.CloudGuardDetectorRecipesDataSource
		factories["oci_cloud_guard_managed_list"] = tf_cloud_guard.CloudGuardManagedListDataSource
		factories["oci_cloud_guard_managed_lists"] = tf_cloud_guard.CloudGuardManagedListsDataSource
		factories["oci_cloud_guard_problem_entities"] = tf_cloud_guard.CloudGuardProblemEntitiesDataSource
		factories["oci_cloud_guard_problem_entity"] = tf_cloud_guard.CloudGuardProblemEntityDataSource
		factories["oci_cloud_guard_responder_recipe"] = tf_cloud_guard.CloudGuardResponderRecipeDataSource
		factories["oci_cloud_guard_responder_recipes"] = tf_cloud_guard.CloudGuardResponderRecipesDataSource
		factories["oci_cloud_guard_saved_queries"] = tf_cloud_guard.CloudGuardSavedQueriesDataSource
		factories["oci_cloud_guard_saved_query"] = tf_cloud_guard.CloudGuardSavedQueryDataSource
		factories["oci_cloud_guard_security_policies"] = tf_cloud_guard.CloudGuardSecurityPoliciesDataSource
		factories["oci_cloud_guard_security_policy"] = tf_cloud_guard.CloudGuardSecurityPolicyDataSource
		factories["oci_cloud_guard_security_recipe"] = tf_cloud_guard.CloudGuardSecurityRecipeDataSource
		factories["oci_cloud_guard_security_recipes"] = tf_cloud_guard.CloudGuardSecurityRecipesDataSource
		factories["oci_cloud_guard_security_zone"] = tf_cloud_guard.CloudGuardSecurityZoneDataSource
		factories["oci_cloud_guard_security_zones"] = tf_cloud_guard.CloudGuardSecurityZonesDataSource
		factories["oci_cloud_guard_target"] = tf_cloud_guard.CloudGuardTargetDataSource
		factories["oci_cloud_guard_targets"] = tf_cloud_guard.CloudGuardTargetsDataSource
		factories["oci_cloud_guard_wlp_agent"] = tf_cloud_guard.CloudGuardWlpAgentDataSource
		factories["oci_cloud_guard_wlp_agents"] = tf_cloud_guard.CloudGuardWlpAgentsDataSource
	}
	if common.CheckForEnabledServices("cloudmigrations") {
		factories["oci_cloud_migrations_migration"] = tf_cloud_migrations.CloudMigrationsMigrationDataSource
		factories["oci_cloud_migrations_migration_asset"] = tf_cloud_migrations.CloudMigrationsMigrationAssetDataSource
		factories["oci_cloud_migrations_migration_assets"] = tf_cloud_migrations.CloudMigrationsMigrationAssetsDataSource
		factories["oci_cloud_migrations_migration_plan"] = tf_cloud_migrations.CloudMigrationsMigrationPlanDataSource
		factories["oci_cloud_migrations_migration_plan_available_shape"] = tf_cloud_migrations.CloudMigrationsMigrationPlanAvailableShapeDataSource
		factories["oci_cloud_migrations_migration_plan_available_shapes"] = tf_cloud_migrations.CloudMigrationsMigrationPlanAvailableShapesDataSource
		factories["oci_cloud_migrations_migration_plans"] = tf_cloud_migrations.CloudMigrationsMigrationPlansDataSource
		factories["oci_cloud_migrations_migrations"] = tf_cloud_migrations.CloudMigrationsMigrationsDataSource
		factories["oci_cloud_migrations_replication_schedule"] = tf_cloud_migrations.CloudMigrationsReplicationScheduleDataSource
		factories["oci_cloud_migrations_replication_schedules"] = tf_cloud_migrations.CloudMigrationsReplicationSchedulesDataSource
		factories["oci_cloud_migrations_target_asset"] = tf_cloud_migrations.CloudMigrationsTargetAssetDataSource
		factories["oci_cloud_migrations_target_assets"] = tf_cloud_migrations.CloudMigrationsTargetAssetsDataSource
	}
	if common.CheckForEnabledServices("clusterhealth") {
		factories["oci_cluster_health_diagnosis_store"] = tf_cluster_health.ClusterHealthDiagnosisStoreDataSource
		factories["oci_cluster_health_diagnosis_stores"] = tf_cluster_health.ClusterHealthDiagnosisStoresDataSource
	}
	if common.CheckForEnabledServices("clusterplacementgroups") {
		factories["oci_cluster_placement_groups_cluster_placement_group"] = tf_cluster_placement_groups.ClusterPlacementGroupsClusterPlacementGroupDataSource
		factories["oci_cluster_placement_groups_cluster_placement_groups"] = tf_cluster_placement_groups.ClusterPlacementGroupsClusterPlacementGroupsDataSource
	}
	if common.CheckForEnabledServices("computecloudatcustomer") {
		factories["oci_compute_cloud_at_customer_ccc_infrastructure"] = tf_compute_cloud_at_customer.ComputeCloudAtCustomerCccInfrastructureDataSource
		factories["oci_compute_cloud_at_customer_ccc_infrastructures"] = tf_compute_cloud_at_customer.ComputeCloudAtCustomerCccInfrastructuresDataSource
		factories["oci_compute_cloud_at_customer_ccc_upgrade_schedule"] = tf_compute_cloud_at_customer.ComputeCloudAtCustomerCccUpgradeScheduleDataSource
		factories["oci_compute_cloud_at_customer_ccc_upgrade_schedules"] = tf_compute_cloud_at_customer.ComputeCloudAtCustomerCccUpgradeSchedulesDataSource
	}
	if common.CheckForEnabledServices("computeinstanceagent") {
		factories["oci_computeinstanceagent_instance_agent_plugin"] = tf_computeinstanceagent.ComputeinstanceagentInstanceAgentPluginDataSource
		factories["oci_computeinstanceagent_instance_agent_plugins"] = tf_computeinstanceagent.ComputeinstanceagentInstanceAgentPluginsDataSource
		factories["oci_computeinstanceagent_instance_available_plugins"] = tf_computeinstanceagent.ComputeinstanceagentInstanceAvailablePluginsDataSource
	}
	if common.CheckForEnabledServices("containerinstances") {
		factories["oci_container_instances_container_instance"] = tf_container_instances.ContainerInstancesContainerInstanceDataSource
		factories["oci_container_instances_container_instance_shape"] = tf_container_instances.ContainerInstancesContainerInstanceShapeDataSource
		factories["oci_container_instances_container_instance_shapes"] = tf_container_instances.ContainerInstancesContainerInstanceShapesDataSource
		factories["oci_container_instances_container_instances"] = tf_container_instances.ContainerInstancesContainerInstancesDataSource
	}
	if common.CheckForEnabledServices("containerengine") {
		factories["oci_containerengine_addon"] = tf_containerengine.ContainerengineAddonDataSource
		factories["oci_containerengine_addon_options"] = tf_containerengine.ContainerengineAddonOptionsDataSource
		factories["oci_containerengine_addons"] = tf_containerengine.ContainerengineAddonsDataSource
		factories["oci_containerengine_cluster"] = tf_containerengine.ContainerengineClusterDataSource
		factories["oci_containerengine_cluster_credential_rotation_status"] = tf_containerengine.ContainerengineClusterCredentialRotationStatusDataSource
		factories["oci_containerengine_cluster_kube_config"] = tf_containerengine.ContainerengineClusterKubeConfigDataSource
		factories["oci_containerengine_cluster_option"] = tf_containerengine.ContainerengineClusterOptionDataSource
		factories["oci_containerengine_cluster_public_api_endpoint_decommission_status"] = tf_containerengine.ContainerengineClusterPublicApiEndpointDecommissionStatusDataSource
		factories["oci_containerengine_cluster_workload_mapping"] = tf_containerengine.ContainerengineClusterWorkloadMappingDataSource
		factories["oci_containerengine_cluster_workload_mappings"] = tf_containerengine.ContainerengineClusterWorkloadMappingsDataSource
		factories["oci_containerengine_clusters"] = tf_containerengine.ContainerengineClustersDataSource
		factories["oci_containerengine_migrate_to_native_vcn_status"] = tf_containerengine.ContainerengineMigrateToNativeVcnStatusDataSource
		factories["oci_containerengine_node_pool"] = tf_containerengine.ContainerengineNodePoolDataSource
		factories["oci_containerengine_node_pool_option"] = tf_containerengine.ContainerengineNodePoolOptionDataSource
		factories["oci_containerengine_node_pools"] = tf_containerengine.ContainerengineNodePoolsDataSource
		factories["oci_containerengine_pod_shapes"] = tf_containerengine.ContainerenginePodShapesDataSource
		factories["oci_containerengine_virtual_node_pool"] = tf_containerengine.ContainerengineVirtualNodePoolDataSource
		factories["oci_containerengine_virtual_node_pools"] = tf_containerengine.ContainerengineVirtualNodePoolsDataSource
		factories["oci_containerengine_work_request_errors"] = tf_containerengine.ContainerengineWorkRequestErrorsDataSource
		factories["oci_containerengine_work_request_log_entries"] = tf_containerengine.ContainerengineWorkRequestLogEntriesDataSource
		factories["oci_containerengine_work_requests"] = tf_containerengine.ContainerengineWorkRequestsDataSource
	}
	if common.CheckForEnabledServices("core") {
		factories["oci_core_app_catalog_listing"] = tf_core.CoreAppCatalogListingDataSource
		factories["oci_core_app_catalog_listing_resource_version"] = tf_core.CoreAppCatalogListingResourceVersionDataSource
		factories["oci_core_app_catalog_listing_resource_versions"] = tf_core.CoreAppCatalogListingResourceVersionsDataSource
		factories["oci_core_app_catalog_listings"] = tf_core.CoreAppCatalogListingsDataSource
		factories["oci_core_app_catalog_subscriptions"] = tf_core.CoreAppCatalogSubscriptionsDataSource
		factories["oci_core_block_volume_replica"] = tf_core.CoreBlockVolumeReplicaDataSource
		factories["oci_core_block_volume_replicas"] = tf_core.CoreBlockVolumeReplicasDataSource
		factories["oci_core_boot_volume"] = tf_core.CoreBootVolumeDataSource
		factories["oci_core_boot_volume_attachments"] = tf_core.CoreBootVolumeAttachmentsDataSource
		factories["oci_core_boot_volume_backup"] = tf_core.CoreBootVolumeBackupDataSource
		factories["oci_core_boot_volume_backups"] = tf_core.CoreBootVolumeBackupsDataSource
		factories["oci_core_boot_volume_replica"] = tf_core.CoreBootVolumeReplicaDataSource
		factories["oci_core_boot_volume_replicas"] = tf_core.CoreBootVolumeReplicasDataSource
		factories["oci_core_boot_volumes"] = tf_core.CoreBootVolumesDataSource
		factories["oci_core_byoasn"] = tf_core.CoreByoasnDataSource
		factories["oci_core_byoasns"] = tf_core.CoreByoasnsDataSource
		factories["oci_core_byoip_allocated_ranges"] = tf_core.CoreByoipAllocatedRangesDataSource
		factories["oci_core_byoip_range"] = tf_core.CoreByoipRangeDataSource
		factories["oci_core_byoip_ranges"] = tf_core.CoreByoipRangesDataSource
		factories["oci_core_capture_filter"] = tf_core.CoreCaptureFilterDataSource
		factories["oci_core_capture_filters"] = tf_core.CoreCaptureFiltersDataSource
		factories["oci_core_cluster_network"] = tf_core.CoreClusterNetworkDataSource
		factories["oci_core_cluster_network_instances"] = tf_core.CoreClusterNetworkInstancesDataSource
		factories["oci_core_cluster_networks"] = tf_core.CoreClusterNetworksDataSource
		factories["oci_core_compute_capacity_reservation"] = tf_core.CoreComputeCapacityReservationDataSource
		factories["oci_core_compute_capacity_reservation_instance_shapes"] = tf_core.CoreComputeCapacityReservationInstanceShapesDataSource
		factories["oci_core_compute_capacity_reservation_instances"] = tf_core.CoreComputeCapacityReservationInstancesDataSource
		factories["oci_core_compute_capacity_reservations"] = tf_core.CoreComputeCapacityReservationsDataSource
		factories["oci_core_compute_capacity_topologies"] = tf_core.CoreComputeCapacityTopologiesDataSource
		factories["oci_core_compute_capacity_topology"] = tf_core.CoreComputeCapacityTopologyDataSource
		factories["oci_core_compute_capacity_topology_compute_bare_metal_hosts"] = tf_core.CoreComputeCapacityTopologyComputeBareMetalHostsDataSource
		factories["oci_core_compute_capacity_topology_compute_hpc_islands"] = tf_core.CoreComputeCapacityTopologyComputeHpcIslandsDataSource
		factories["oci_core_compute_capacity_topology_compute_network_blocks"] = tf_core.CoreComputeCapacityTopologyComputeNetworkBlocksDataSource
		factories["oci_core_compute_cluster"] = tf_core.CoreComputeClusterDataSource
		factories["oci_core_compute_clusters"] = tf_core.CoreComputeClustersDataSource
		factories["oci_core_compute_global_image_capability_schema"] = tf_core.CoreComputeGlobalImageCapabilitySchemaDataSource
		factories["oci_core_compute_global_image_capability_schemas"] = tf_core.CoreComputeGlobalImageCapabilitySchemasDataSource
		factories["oci_core_compute_global_image_capability_schemas_version"] = tf_core.CoreComputeGlobalImageCapabilitySchemasVersionDataSource
		factories["oci_core_compute_global_image_capability_schemas_versions"] = tf_core.CoreComputeGlobalImageCapabilitySchemasVersionsDataSource
		factories["oci_core_compute_gpu_memory_cluster"] = tf_core.CoreComputeGpuMemoryClusterDataSource
		factories["oci_core_compute_gpu_memory_cluster_instances"] = tf_core.CoreComputeGpuMemoryClusterInstancesDataSource
		factories["oci_core_compute_gpu_memory_clusters"] = tf_core.CoreComputeGpuMemoryClustersDataSource
		factories["oci_core_compute_gpu_memory_fabric"] = tf_core.CoreComputeGpuMemoryFabricDataSource
		factories["oci_core_compute_gpu_memory_fabrics"] = tf_core.CoreComputeGpuMemoryFabricsDataSource
		factories["oci_core_compute_host"] = tf_core.CoreComputeHostDataSource
		factories["oci_core_compute_host_group"] = tf_core.CoreComputeHostGroupDataSource
		factories["oci_core_compute_host_groups"] = tf_core.CoreComputeHostGroupsDataSource
		factories["oci_core_compute_hosts"] = tf_core.CoreComputeHostsDataSource
		factories["oci_core_compute_image_capability_schema"] = tf_core.CoreComputeImageCapabilitySchemaDataSource
		factories["oci_core_compute_image_capability_schemas"] = tf_core.CoreComputeImageCapabilitySchemasDataSource
		factories["oci_core_console_histories"] = tf_core.CoreConsoleHistoriesDataSource
		factories["oci_core_console_history_data"] = tf_core.CoreConsoleHistoryContentDataSource
		factories["oci_core_cpe_device_shape"] = tf_core.CoreCpeDeviceShapeDataSource
		factories["oci_core_cpe_device_shapes"] = tf_core.CoreCpeDeviceShapesDataSource
		factories["oci_core_cpes"] = tf_core.CoreCpesDataSource
		factories["oci_core_cross_connect"] = tf_core.CoreCrossConnectDataSource
		factories["oci_core_cross_connect_group"] = tf_core.CoreCrossConnectGroupDataSource
		factories["oci_core_cross_connect_groups"] = tf_core.CoreCrossConnectGroupsDataSource
		factories["oci_core_cross_connect_locations"] = tf_core.CoreCrossConnectLocationsDataSource
		factories["oci_core_cross_connect_port_speed_shapes"] = tf_core.CoreCrossConnectPortSpeedShapesDataSource
		factories["oci_core_cross_connect_status"] = tf_core.CoreCrossConnectStatusDataSource
		factories["oci_core_cross_connects"] = tf_core.CoreCrossConnectsDataSource
		factories["oci_core_dedicated_vm_host"] = tf_core.CoreDedicatedVmHostDataSource
		factories["oci_core_dedicated_vm_host_instance_shapes"] = tf_core.CoreDedicatedVmHostInstanceShapesDataSource
		factories["oci_core_dedicated_vm_host_shapes"] = tf_core.CoreDedicatedVmHostShapesDataSource
		factories["oci_core_dedicated_vm_hosts"] = tf_core.CoreDedicatedVmHostsDataSource
		factories["oci_core_dedicated_vm_hosts_instances"] = tf_core.CoreDedicatedVmHostsInstancesDataSource
		factories["oci_core_dhcp_options"] = tf_core.CoreDhcpOptionsDataSource
		factories["oci_core_drg_attachments"] = tf_core.CoreDrgAttachmentsDataSource
		factories["oci_core_drg_route_distribution"] = tf_core.CoreDrgRouteDistributionDataSource
		factories["oci_core_drg_route_distribution_statements"] = tf_core.CoreDrgRouteDistributionStatementsDataSource
		factories["oci_core_drg_route_distributions"] = tf_core.CoreDrgRouteDistributionsDataSource
		factories["oci_core_drg_route_table"] = tf_core.CoreDrgRouteTableDataSource
		factories["oci_core_drg_route_table_route_rules"] = tf_core.CoreDrgRouteTableRouteRulesDataSource
		factories["oci_core_drg_route_tables"] = tf_core.CoreDrgRouteTablesDataSource
		factories["oci_core_drgs"] = tf_core.CoreDrgsDataSource
		factories["oci_core_fast_connect_provider_service"] = tf_core.CoreFastConnectProviderServiceDataSource
		factories["oci_core_fast_connect_provider_service_key"] = tf_core.CoreFastConnectProviderServiceKeyDataSource
		factories["oci_core_fast_connect_provider_services"] = tf_core.CoreFastConnectProviderServicesDataSource
		factories["oci_core_firmware_bundle"] = tf_core.CoreFirmwareBundleDataSource
		factories["oci_core_firmware_bundles"] = tf_core.CoreFirmwareBundlesDataSource
		factories["oci_core_image"] = tf_core.CoreImageDataSource
		factories["oci_core_image_shape"] = tf_core.CoreImageShapeDataSource
		factories["oci_core_image_shapes"] = tf_core.CoreImageShapesDataSource
		factories["oci_core_images"] = tf_core.CoreImagesDataSource
		factories["oci_core_instance"] = tf_core.CoreInstanceDataSource
		factories["oci_core_instance_configuration"] = tf_core.CoreInstanceConfigurationDataSource
		factories["oci_core_instance_configurations"] = tf_core.CoreInstanceConfigurationsDataSource
		factories["oci_core_instance_console_connections"] = tf_core.CoreInstanceConsoleConnectionsDataSource
		factories["oci_core_instance_credentials"] = tf_core.CoreInstanceCredentialDataSource
		factories["oci_core_instance_devices"] = tf_core.CoreInstanceDevicesDataSource
		factories["oci_core_instance_maintenance_event"] = tf_core.CoreInstanceMaintenanceEventDataSource
		factories["oci_core_instance_maintenance_events"] = tf_core.CoreInstanceMaintenanceEventsDataSource
		factories["oci_core_instance_maintenance_reboot"] = tf_core.CoreInstanceMaintenanceRebootDataSource
		factories["oci_core_instance_measured_boot_report"] = tf_core.CoreInstanceMeasuredBootReportDataSource
		factories["oci_core_instance_pool"] = tf_core.CoreInstancePoolDataSource
		factories["oci_core_instance_pool_instances"] = tf_core.CoreInstancePoolInstancesDataSource
		factories["oci_core_instance_pool_load_balancer_attachment"] = tf_core.CoreInstancePoolLoadBalancerAttachmentDataSource
		factories["oci_core_instance_pools"] = tf_core.CoreInstancePoolsDataSource
		factories["oci_core_instances"] = tf_core.CoreInstancesDataSource
		factories["oci_core_internet_gateways"] = tf_core.CoreInternetGatewaysDataSource
		factories["oci_core_ip_inventory_subnet"] = tf_core.CoreIpInventorySubnetDataSource
		factories["oci_core_ip_inventory_subnet_cidr"] = tf_core.CoreIpInventorySubnetCidrDataSource
		factories["oci_core_ip_inventory_vcn_overlaps"] = tf_core.CoreIpInventoryVcnOverlapsDataSource
		factories["oci_core_ipsec_algorithm"] = tf_core.CoreIpsecAlgorithmDataSource
		factories["oci_core_ipsec_config"] = tf_core.CoreIpSecConnectionDeviceConfigDataSource
		factories["oci_core_ipsec_connection_tunnel"] = tf_core.CoreIpSecConnectionTunnelDataSource
		factories["oci_core_ipsec_connection_tunnel_error"] = tf_core.CoreIpsecConnectionTunnelErrorDataSource
		factories["oci_core_ipsec_connection_tunnel_routes"] = tf_core.CoreIpsecConnectionTunnelRoutesDataSource
		factories["oci_core_ipsec_connection_tunnels"] = tf_core.CoreIpSecConnectionTunnelsDataSource
		factories["oci_core_ipsec_connections"] = tf_core.CoreIpSecConnectionsDataSource
		factories["oci_core_ipsec_status"] = tf_core.CoreIpSecConnectionDeviceStatusDataSource
		factories["oci_core_ipv6"] = tf_core.CoreIpv6DataSource
		factories["oci_core_ipv6s"] = tf_core.CoreIpv6sDataSource
		factories["oci_core_letter_of_authority"] = tf_core.CoreLetterOfAuthorityDataSource
		factories["oci_core_local_peering_gateways"] = tf_core.CoreLocalPeeringGatewaysDataSource
		factories["oci_core_nat_gateway"] = tf_core.CoreNatGatewayDataSource
		factories["oci_core_nat_gateways"] = tf_core.CoreNatGatewaysDataSource
		factories["oci_core_network_security_group"] = tf_core.CoreNetworkSecurityGroupDataSource
		factories["oci_core_network_security_group_security_rules"] = tf_core.CoreNetworkSecurityGroupSecurityRulesDataSource
		factories["oci_core_network_security_group_vnics"] = tf_core.CoreNetworkSecurityGroupVnicsDataSource
		factories["oci_core_network_security_groups"] = tf_core.CoreNetworkSecurityGroupsDataSource
		factories["oci_core_peer_region_for_remote_peerings"] = tf_core.CorePeerRegionForRemotePeeringsDataSource
		factories["oci_core_private_ip"] = tf_core.CorePrivateIpDataSource
		factories["oci_core_private_ips"] = tf_core.CorePrivateIpsDataSource
		factories["oci_core_public_ip"] = tf_core.CorePublicIpDataSource
		factories["oci_core_public_ip_pool"] = tf_core.CorePublicIpPoolDataSource
		factories["oci_core_public_ip_pools"] = tf_core.CorePublicIpPoolsDataSource
		factories["oci_core_public_ips"] = tf_core.CorePublicIpsDataSource
		factories["oci_core_remote_peering_connections"] = tf_core.CoreRemotePeeringConnectionsDataSource
		factories["oci_core_route_tables"] = tf_core.CoreRouteTablesDataSource
		factories["oci_core_security_lists"] = tf_core.CoreSecurityListsDataSource
		factories["oci_core_service_gateways"] = tf_core.CoreServiceGatewaysDataSource
		factories["oci_core_services"] = tf_core.CoreServicesDataSource
		factories["oci_core_shapes"] = tf_core.CoreShapesDataSource
		factories["oci_core_subnet"] = tf_core.CoreSubnetDataSource
		factories["oci_core_subnets"] = tf_core.CoreSubnetsDataSource
		factories["oci_core_tunnel_security_associations"] = tf_core.CoreTunnelSecurityAssociationsDataSource
		factories["oci_core_vcn"] = tf_core.CoreVcnDataSource
		factories["oci_core_vcn_dns_resolver_association"] = tf_core.CoreVcnDnsResolverAssociationDataSource
		factories["oci_core_vcns"] = tf_core.CoreVcnsDataSource
		factories["oci_core_virtual_circuit"] = tf_core.CoreVirtualCircuitDataSource
		factories["oci_core_virtual_circuit_associated_tunnels"] = tf_core.CoreVirtualCircuitAssociatedTunnelsDataSource
		factories["oci_core_virtual_circuit_bandwidth_shapes"] = tf_core.CoreVirtualCircuitBandwidthShapesDataSource
		factories["oci_core_virtual_circuit_public_prefixes"] = tf_core.CoreVirtualCircuitPublicPrefixesDataSource
		factories["oci_core_virtual_circuits"] = tf_core.CoreVirtualCircuitsDataSource
		factories["oci_core_vlan"] = tf_core.CoreVlanDataSource
		factories["oci_core_vlans"] = tf_core.CoreVlansDataSource
		factories["oci_core_vnic"] = tf_core.CoreVnicDataSource
		factories["oci_core_vnic_attachments"] = tf_core.CoreVnicAttachmentsDataSource
		factories["oci_core_volume"] = tf_core.CoreVolumeDataSource
		factories["oci_core_volume_attachments"] = tf_core.CoreVolumeAttachmentsDataSource
		factories["oci_core_volume_backup_policies"] = tf_core.CoreVolumeBackupPoliciesDataSource
		factories["oci_core_volume_backup_policy_assignments"] = tf_core.CoreVolumeBackupPolicyAssignmentsDataSource
		factories["oci_core_volume_backups"] = tf_core.CoreVolumeBackupsDataSource
		factories["oci_core_volume_group_backups"] = tf_core.CoreVolumeGroupBackupsDataSource
		factories["oci_core_volume_group_replica"] = tf_core.CoreVolumeGroupReplicaDataSource
		factories["oci_core_volume_group_replicas"] = tf_core.CoreVolumeGroupReplicasDataSource
		factories["oci_core_volume_groups"] = tf_core.CoreVolumeGroupsDataSource
		factories["oci_core_volumes"] = tf_core.CoreVolumesDataSource
		factories["oci_core_vtap"] = tf_core.CoreVtapDataSource
		factories["oci_core_vtaps"] = tf_core.CoreVtapsDataSource
	}
	if common.CheckForEnabledServices("costad") {
		factories["oci_costad_cost_alert_subscription"] = tf_costad.CostadCostAlertSubscriptionDataSource
		factories["oci_costad_cost_alert_subscriptions"] = tf_costad.CostadCostAlertSubscriptionsDataSource
		factories["oci_costad_cost_anomaly_event"] = tf_costad.CostadCostAnomalyEventDataSource
		factories["oci_costad_cost_anomaly_event_analytics"] = tf_costad.CostadCostAnomalyEventAnalyticsDataSource
		factories["oci_costad_cost_anomaly_events"] = tf_costad.CostadCostAnomalyEventsDataSource
		factories["oci_costad_cost_anomaly_monitor"] = tf_costad.CostadCostAnomalyMonitorDataSource
		factories["oci_costad_cost_anomaly_monitors"] = tf_costad.CostadCostAnomalyMonitorsDataSource
	}
	if common.CheckForEnabledServices("datalabelingservice") {
		factories["oci_data_labeling_service_annotation_format"] = tf_data_labeling_service.DataLabelingServiceAnnotationFormatDataSource
		factories["oci_data_labeling_service_annotation_formats"] = tf_data_labeling_service.DataLabelingServiceAnnotationFormatsDataSource
		factories["oci_data_labeling_service_dataset"] = tf_data_labeling_service.DataLabelingServiceDatasetDataSource
		factories["oci_data_labeling_service_datasets"] = tf_data_labeling_service.DataLabelingServiceDatasetsDataSource
	}
	if common.CheckForEnabledServices("datasafe") {
		factories["oci_data_safe_alert"] = tf_data_safe.DataSafeAlertDataSource
		factories["oci_data_safe_alert_analytic"] = tf_data_safe.DataSafeAlertAnalyticDataSource
		factories["oci_data_safe_alert_policies"] = tf_data_safe.DataSafeAlertPoliciesDataSource
		factories["oci_data_safe_alert_policy"] = tf_data_safe.DataSafeAlertPolicyDataSource
		factories["oci_data_safe_alert_policy_rule"] = tf_data_safe.DataSafeAlertPolicyRuleDataSource
		factories["oci_data_safe_alert_policy_rules"] = tf_data_safe.DataSafeAlertPolicyRulesDataSource
		factories["oci_data_safe_alerts"] = tf_data_safe.DataSafeAlertsDataSource
		factories["oci_data_safe_attribute_set"] = tf_data_safe.DataSafeAttributeSetDataSource
		factories["oci_data_safe_attribute_set_associated_resources"] = tf_data_safe.DataSafeAttributeSetAssociatedResourcesDataSource
		factories["oci_data_safe_attribute_sets"] = tf_data_safe.DataSafeAttributeSetsDataSource
		factories["oci_data_safe_audit_archive_retrieval"] = tf_data_safe.DataSafeAuditArchiveRetrievalDataSource
		factories["oci_data_safe_audit_archive_retrievals"] = tf_data_safe.DataSafeAuditArchiveRetrievalsDataSource
		factories["oci_data_safe_audit_event"] = tf_data_safe.DataSafeAuditEventDataSource
		factories["oci_data_safe_audit_event_analytic"] = tf_data_safe.DataSafeAuditEventAnalyticDataSource
		factories["oci_data_safe_audit_events"] = tf_data_safe.DataSafeAuditEventsDataSource
		factories["oci_data_safe_audit_policies"] = tf_data_safe.DataSafeAuditPoliciesDataSource
		factories["oci_data_safe_audit_policy"] = tf_data_safe.DataSafeAuditPolicyDataSource
		factories["oci_data_safe_audit_profile"] = tf_data_safe.DataSafeAuditProfileDataSource
		factories["oci_data_safe_audit_profile_analytic"] = tf_data_safe.DataSafeAuditProfileAnalyticDataSource
		factories["oci_data_safe_audit_profile_available_audit_volume"] = tf_data_safe.DataSafeAuditProfileAvailableAuditVolumeDataSource
		factories["oci_data_safe_audit_profile_available_audit_volumes"] = tf_data_safe.DataSafeAuditProfileAvailableAuditVolumesDataSource
		factories["oci_data_safe_audit_profile_collected_audit_volume"] = tf_data_safe.DataSafeAuditProfileCollectedAuditVolumeDataSource
		factories["oci_data_safe_audit_profile_collected_audit_volumes"] = tf_data_safe.DataSafeAuditProfileCollectedAuditVolumesDataSource
		factories["oci_data_safe_audit_profile_target_overrides"] = tf_data_safe.DataSafeAuditProfileTargetOverridesDataSource
		factories["oci_data_safe_audit_profiles"] = tf_data_safe.DataSafeAuditProfilesDataSource
		factories["oci_data_safe_audit_trail"] = tf_data_safe.DataSafeAuditTrailDataSource
		factories["oci_data_safe_audit_trail_analytic"] = tf_data_safe.DataSafeAuditTrailAnalyticDataSource
		factories["oci_data_safe_audit_trails"] = tf_data_safe.DataSafeAuditTrailsDataSource
		factories["oci_data_safe_compatible_formats_for_data_type"] = tf_data_safe.DataSafeCompatibleFormatsForDataTypeDataSource
		factories["oci_data_safe_compatible_formats_for_sensitive_type"] = tf_data_safe.DataSafeCompatibleFormatsForSensitiveTypeDataSource
		factories["oci_data_safe_crypto_assessment"] = tf_data_safe.DataSafeCryptoAssessmentDataSource
		factories["oci_data_safe_crypto_assessment_backup_sets"] = tf_data_safe.DataSafeCryptoAssessmentBackupSetsDataSource
		factories["oci_data_safe_crypto_assessment_cbom_items"] = tf_data_safe.DataSafeCryptoAssessmentCbomItemsDataSource
		factories["oci_data_safe_crypto_assessment_certificates"] = tf_data_safe.DataSafeCryptoAssessmentCertificatesDataSource
		factories["oci_data_safe_crypto_assessment_finding_analytics"] = tf_data_safe.DataSafeCryptoAssessmentFindingAnalyticsDataSource
		factories["oci_data_safe_crypto_assessment_finding_targets"] = tf_data_safe.DataSafeCryptoAssessmentFindingTargetsDataSource
		factories["oci_data_safe_crypto_assessment_findings"] = tf_data_safe.DataSafeCryptoAssessmentFindingsDataSource
		factories["oci_data_safe_crypto_assessment_keys"] = tf_data_safe.DataSafeCryptoAssessmentKeysDataSource
		factories["oci_data_safe_crypto_assessment_sqlnet_parameter"] = tf_data_safe.DataSafeCryptoAssessmentSqlnetParameterDataSource
		factories["oci_data_safe_crypto_assessment_tde_objects"] = tf_data_safe.DataSafeCryptoAssessmentTdeObjectsDataSource
		factories["oci_data_safe_crypto_assessment_wallets"] = tf_data_safe.DataSafeCryptoAssessmentWalletsDataSource
		factories["oci_data_safe_crypto_assessments"] = tf_data_safe.DataSafeCryptoAssessmentsDataSource
		factories["oci_data_safe_data_safe_configuration"] = tf_data_safe.DataSafeDataSafeConfigurationDataSource
		factories["oci_data_safe_data_safe_private_endpoint"] = tf_data_safe.DataSafeDataSafePrivateEndpointDataSource
		factories["oci_data_safe_data_safe_private_endpoints"] = tf_data_safe.DataSafeDataSafePrivateEndpointsDataSource
		factories["oci_data_safe_database_security_config"] = tf_data_safe.DataSafeDatabaseSecurityConfigDataSource
		factories["oci_data_safe_database_security_configs"] = tf_data_safe.DataSafeDatabaseSecurityConfigsDataSource
		factories["oci_data_safe_discovery_analytic"] = tf_data_safe.DataSafeDiscoveryAnalyticDataSource
		factories["oci_data_safe_discovery_analytics"] = tf_data_safe.DataSafeDiscoveryAnalyticsDataSource
		factories["oci_data_safe_discovery_job"] = tf_data_safe.DataSafeDiscoveryJobDataSource
		factories["oci_data_safe_discovery_jobs"] = tf_data_safe.DataSafeDiscoveryJobsDataSource
		factories["oci_data_safe_discovery_jobs_result"] = tf_data_safe.DataSafeDiscoveryJobsResultDataSource
		factories["oci_data_safe_discovery_jobs_results"] = tf_data_safe.DataSafeDiscoveryJobsResultsDataSource
		factories["oci_data_safe_library_masking_format"] = tf_data_safe.DataSafeLibraryMaskingFormatDataSource
		factories["oci_data_safe_library_masking_formats"] = tf_data_safe.DataSafeLibraryMaskingFormatsDataSource
		factories["oci_data_safe_list_user_grants"] = tf_data_safe.DataSafeListUserGrantsDataSource
		factories["oci_data_safe_masking_analytic"] = tf_data_safe.DataSafeMaskingAnalyticDataSource
		factories["oci_data_safe_masking_analytics"] = tf_data_safe.DataSafeMaskingAnalyticsDataSource
		factories["oci_data_safe_masking_policies"] = tf_data_safe.DataSafeMaskingPoliciesDataSource
		factories["oci_data_safe_masking_policies_masking_column"] = tf_data_safe.DataSafeMaskingPoliciesMaskingColumnDataSource
		factories["oci_data_safe_masking_policies_masking_columns"] = tf_data_safe.DataSafeMaskingPoliciesMaskingColumnsDataSource
		factories["oci_data_safe_masking_policy"] = tf_data_safe.DataSafeMaskingPolicyDataSource
		factories["oci_data_safe_masking_policy_health_report"] = tf_data_safe.DataSafeMaskingPolicyHealthReportDataSource
		factories["oci_data_safe_masking_policy_health_report_logs"] = tf_data_safe.DataSafeMaskingPolicyHealthReportLogsDataSource
		factories["oci_data_safe_masking_policy_health_reports"] = tf_data_safe.DataSafeMaskingPolicyHealthReportsDataSource
		factories["oci_data_safe_masking_policy_masking_objects"] = tf_data_safe.DataSafeMaskingPolicyMaskingObjectsDataSource
		factories["oci_data_safe_masking_policy_masking_schemas"] = tf_data_safe.DataSafeMaskingPolicyMaskingSchemasDataSource
		factories["oci_data_safe_masking_policy_referential_relations"] = tf_data_safe.DataSafeMaskingPolicyReferentialRelationsDataSource
		factories["oci_data_safe_masking_report"] = tf_data_safe.DataSafeMaskingReportDataSource
		factories["oci_data_safe_masking_report_masking_errors"] = tf_data_safe.DataSafeMaskingReportMaskingErrorsDataSource
		factories["oci_data_safe_masking_reports"] = tf_data_safe.DataSafeMaskingReportsDataSource
		factories["oci_data_safe_masking_reports_masked_column"] = tf_data_safe.DataSafeMaskingReportsMaskedColumnDataSource
		factories["oci_data_safe_masking_reports_masked_columns"] = tf_data_safe.DataSafeMaskingReportsMaskedColumnsDataSource
		factories["oci_data_safe_on_prem_connector"] = tf_data_safe.DataSafeOnPremConnectorDataSource
		factories["oci_data_safe_on_prem_connectors"] = tf_data_safe.DataSafeOnPremConnectorsDataSource
		factories["oci_data_safe_report"] = tf_data_safe.DataSafeReportDataSource
		factories["oci_data_safe_report_content"] = tf_data_safe.DataSafeReportContentDataSource
		factories["oci_data_safe_report_definition"] = tf_data_safe.DataSafeReportDefinitionDataSource
		factories["oci_data_safe_report_definitions"] = tf_data_safe.DataSafeReportDefinitionsDataSource
		factories["oci_data_safe_reports"] = tf_data_safe.DataSafeReportsDataSource
		factories["oci_data_safe_sdm_masking_policy_difference"] = tf_data_safe.DataSafeSdmMaskingPolicyDifferenceDataSource
		factories["oci_data_safe_sdm_masking_policy_difference_difference_column"] = tf_data_safe.DataSafeSdmMaskingPolicyDifferenceDifferenceColumnDataSource
		factories["oci_data_safe_sdm_masking_policy_difference_difference_columns"] = tf_data_safe.DataSafeSdmMaskingPolicyDifferenceDifferenceColumnsDataSource
		factories["oci_data_safe_sdm_masking_policy_differences"] = tf_data_safe.DataSafeSdmMaskingPolicyDifferencesDataSource
		factories["oci_data_safe_security_assessment"] = tf_data_safe.DataSafeSecurityAssessmentDataSource
		factories["oci_data_safe_security_assessment_checks"] = tf_data_safe.DataSafeSecurityAssessmentChecksDataSource
		factories["oci_data_safe_security_assessment_comparison"] = tf_data_safe.DataSafeSecurityAssessmentComparisonDataSource
		factories["oci_data_safe_security_assessment_finding"] = tf_data_safe.DataSafeSecurityAssessmentFindingsDataSource
		factories["oci_data_safe_security_assessment_finding_analytics"] = tf_data_safe.DataSafeSecurityAssessmentFindingAnalyticsDataSource
		factories["oci_data_safe_security_assessment_findings"] = tf_data_safe.DataSafeSecurityAssessmentFindingsDataSource
		factories["oci_data_safe_security_assessment_findings_change_audit_logs"] = tf_data_safe.DataSafeSecurityAssessmentFindingsChangeAuditLogsDataSource
		factories["oci_data_safe_security_assessment_security_feature_analytics"] = tf_data_safe.DataSafeSecurityAssessmentSecurityFeatureAnalyticsDataSource
		factories["oci_data_safe_security_assessment_security_features"] = tf_data_safe.DataSafeSecurityAssessmentSecurityFeaturesDataSource
		factories["oci_data_safe_security_assessment_template_analytics"] = tf_data_safe.DataSafeSecurityAssessmentTemplateAnalyticsDataSource
		factories["oci_data_safe_security_assessment_template_association_analytics"] = tf_data_safe.DataSafeSecurityAssessmentTemplateAssociationAnalyticsDataSource
		factories["oci_data_safe_security_assessment_template_baseline_comparison"] = tf_data_safe.DataSafeSecurityAssessmentTemplateBaselineComparisonDataSource
		factories["oci_data_safe_security_assessments"] = tf_data_safe.DataSafeSecurityAssessmentsDataSource
		factories["oci_data_safe_security_policies"] = tf_data_safe.DataSafeSecurityPoliciesDataSource
		factories["oci_data_safe_security_policy"] = tf_data_safe.DataSafeSecurityPolicyDataSource
		factories["oci_data_safe_security_policy_config"] = tf_data_safe.DataSafeSecurityPolicyConfigDataSource
		factories["oci_data_safe_security_policy_configs"] = tf_data_safe.DataSafeSecurityPolicyConfigsDataSource
		factories["oci_data_safe_security_policy_deployment"] = tf_data_safe.DataSafeSecurityPolicyDeploymentDataSource
		factories["oci_data_safe_security_policy_deployment_security_policy_entry_state"] = tf_data_safe.DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStateDataSource
		factories["oci_data_safe_security_policy_deployment_security_policy_entry_states"] = tf_data_safe.DataSafeSecurityPolicyDeploymentSecurityPolicyEntryStatesDataSource
		factories["oci_data_safe_security_policy_deployments"] = tf_data_safe.DataSafeSecurityPolicyDeploymentsDataSource
		factories["oci_data_safe_security_policy_report"] = tf_data_safe.DataSafeSecurityPolicyReportDataSource
		factories["oci_data_safe_security_policy_report_database_table_access_entries"] = tf_data_safe.DataSafeSecurityPolicyReportDatabaseTableAccessEntriesDataSource
		factories["oci_data_safe_security_policy_report_database_table_access_entry"] = tf_data_safe.DataSafeSecurityPolicyReportDatabaseTableAccessEntryDataSource
		factories["oci_data_safe_security_policy_report_database_view_access_entries"] = tf_data_safe.DataSafeSecurityPolicyReportDatabaseViewAccessEntriesDataSource
		factories["oci_data_safe_security_policy_report_database_view_access_entry"] = tf_data_safe.DataSafeSecurityPolicyReportDatabaseViewAccessEntryDataSource
		factories["oci_data_safe_security_policy_report_role_grant_paths"] = tf_data_safe.DataSafeSecurityPolicyReportRoleGrantPathsDataSource
		factories["oci_data_safe_security_policy_reports"] = tf_data_safe.DataSafeSecurityPolicyReportsDataSource
		factories["oci_data_safe_sensitive_column_analytics"] = tf_data_safe.DataSafeSensitiveColumnAnalyticsDataSource
		factories["oci_data_safe_sensitive_data_model"] = tf_data_safe.DataSafeSensitiveDataModelDataSource
		factories["oci_data_safe_sensitive_data_model_referential_relation"] = tf_data_safe.DataSafeSensitiveDataModelReferentialRelationDataSource
		factories["oci_data_safe_sensitive_data_model_referential_relations"] = tf_data_safe.DataSafeSensitiveDataModelReferentialRelationsDataSource
		factories["oci_data_safe_sensitive_data_model_sensitive_objects"] = tf_data_safe.DataSafeSensitiveDataModelSensitiveObjectsDataSource
		factories["oci_data_safe_sensitive_data_model_sensitive_schemas"] = tf_data_safe.DataSafeSensitiveDataModelSensitiveSchemasDataSource
		factories["oci_data_safe_sensitive_data_model_sensitive_types"] = tf_data_safe.DataSafeSensitiveDataModelSensitiveTypesDataSource
		factories["oci_data_safe_sensitive_data_models"] = tf_data_safe.DataSafeSensitiveDataModelsDataSource
		factories["oci_data_safe_sensitive_data_models_sensitive_column"] = tf_data_safe.DataSafeSensitiveDataModelsSensitiveColumnDataSource
		factories["oci_data_safe_sensitive_data_models_sensitive_columns"] = tf_data_safe.DataSafeSensitiveDataModelsSensitiveColumnsDataSource
		factories["oci_data_safe_sensitive_type"] = tf_data_safe.DataSafeSensitiveTypeDataSource
		factories["oci_data_safe_sensitive_type_group"] = tf_data_safe.DataSafeSensitiveTypeGroupDataSource
		factories["oci_data_safe_sensitive_type_group_grouped_sensitive_types"] = tf_data_safe.DataSafeSensitiveTypeGroupGroupedSensitiveTypesDataSource
		factories["oci_data_safe_sensitive_type_groups"] = tf_data_safe.DataSafeSensitiveTypeGroupsDataSource
		factories["oci_data_safe_sensitive_types"] = tf_data_safe.DataSafeSensitiveTypesDataSource
		factories["oci_data_safe_sensitive_types_export"] = tf_data_safe.DataSafeSensitiveTypesExportDataSource
		factories["oci_data_safe_sensitive_types_exports"] = tf_data_safe.DataSafeSensitiveTypesExportsDataSource
		factories["oci_data_safe_sql_collection"] = tf_data_safe.DataSafeSqlCollectionDataSource
		factories["oci_data_safe_sql_collection_analytics"] = tf_data_safe.DataSafeSqlCollectionAnalyticsDataSource
		factories["oci_data_safe_sql_collection_log_insights"] = tf_data_safe.DataSafeSqlCollectionLogInsightsDataSource
		factories["oci_data_safe_sql_collections"] = tf_data_safe.DataSafeSqlCollectionsDataSource
		factories["oci_data_safe_sql_firewall_allowed_sql"] = tf_data_safe.DataSafeSqlFirewallAllowedSqlDataSource
		factories["oci_data_safe_sql_firewall_allowed_sql_analytics"] = tf_data_safe.DataSafeSqlFirewallAllowedSqlAnalyticsDataSource
		factories["oci_data_safe_sql_firewall_allowed_sqls"] = tf_data_safe.DataSafeSqlFirewallAllowedSqlsDataSource
		factories["oci_data_safe_sql_firewall_policies"] = tf_data_safe.DataSafeSqlFirewallPoliciesDataSource
		factories["oci_data_safe_sql_firewall_policy"] = tf_data_safe.DataSafeSqlFirewallPolicyDataSource
		factories["oci_data_safe_sql_firewall_policy_analytics"] = tf_data_safe.DataSafeSqlFirewallPolicyAnalyticsDataSource
		factories["oci_data_safe_sql_firewall_violation_analytics"] = tf_data_safe.DataSafeSqlFirewallViolationAnalyticsDataSource
		factories["oci_data_safe_sql_firewall_violations"] = tf_data_safe.DataSafeSqlFirewallViolationsDataSource
		factories["oci_data_safe_target_alert_policy_association"] = tf_data_safe.DataSafeTargetAlertPolicyAssociationDataSource
		factories["oci_data_safe_target_alert_policy_association_unassociated_target_members"] = tf_data_safe.DataSafeTargetAlertPolicyAssociationUnassociatedTargetMembersDataSource
		factories["oci_data_safe_target_alert_policy_associations"] = tf_data_safe.DataSafeTargetAlertPolicyAssociationsDataSource
		factories["oci_data_safe_target_database"] = tf_data_safe.DataSafeTargetDatabaseDataSource
		factories["oci_data_safe_target_database_group"] = tf_data_safe.DataSafeTargetDatabaseGroupDataSource
		factories["oci_data_safe_target_database_group_group_member"] = tf_data_safe.DataSafeTargetDatabaseGroupGroupMemberDataSource
		factories["oci_data_safe_target_database_groups"] = tf_data_safe.DataSafeTargetDatabaseGroupsDataSource
		factories["oci_data_safe_target_database_peer_target_database"] = tf_data_safe.DataSafeTargetDatabasePeerTargetDatabaseDataSource
		factories["oci_data_safe_target_database_peer_target_databases"] = tf_data_safe.DataSafeTargetDatabasePeerTargetDatabasesDataSource
		factories["oci_data_safe_target_database_role"] = tf_data_safe.DataSafeTargetDatabaseRolesDataSource
		factories["oci_data_safe_target_database_roles"] = tf_data_safe.DataSafeTargetDatabaseRolesDataSource
		factories["oci_data_safe_target_databases"] = tf_data_safe.DataSafeTargetDatabasesDataSource
		factories["oci_data_safe_target_databases_columns"] = tf_data_safe.DataSafeTargetDatabasesColumnsDataSource
		factories["oci_data_safe_target_databases_schemas"] = tf_data_safe.DataSafeTargetDatabasesSchemasDataSource
		factories["oci_data_safe_target_databases_tables"] = tf_data_safe.DataSafeTargetDatabasesTablesDataSource
		factories["oci_data_safe_unified_audit_policies"] = tf_data_safe.DataSafeUnifiedAuditPoliciesDataSource
		factories["oci_data_safe_unified_audit_policy"] = tf_data_safe.DataSafeUnifiedAuditPolicyDataSource
		factories["oci_data_safe_unified_audit_policy_definition"] = tf_data_safe.DataSafeUnifiedAuditPolicyDefinitionDataSource
		factories["oci_data_safe_unified_audit_policy_definitions"] = tf_data_safe.DataSafeUnifiedAuditPolicyDefinitionsDataSource
		factories["oci_data_safe_user_assessment"] = tf_data_safe.DataSafeUserAssessmentDataSource
		factories["oci_data_safe_user_assessment_comparison"] = tf_data_safe.DataSafeUserAssessmentComparisonDataSource
		factories["oci_data_safe_user_assessment_password_expiry_date_analytics"] = tf_data_safe.DataSafeUserAssessmentPasswordExpiryDateAnalyticsDataSource
		factories["oci_data_safe_user_assessment_profile_analytics"] = tf_data_safe.DataSafeUserAssessmentProfileAnalyticsDataSource
		factories["oci_data_safe_user_assessment_profiles"] = tf_data_safe.DataSafeUserAssessmentProfilesDataSource
		factories["oci_data_safe_user_assessment_user_access_analytics"] = tf_data_safe.DataSafeUserAssessmentUserAccessAnalyticsDataSource
		factories["oci_data_safe_user_assessment_user_analytics"] = tf_data_safe.DataSafeUserAssessmentUserAnalyticsDataSource
		factories["oci_data_safe_user_assessment_users"] = tf_data_safe.DataSafeUserAssessmentUsersDataSource
		factories["oci_data_safe_user_assessments"] = tf_data_safe.DataSafeUserAssessmentsDataSource
	}
	if common.CheckForEnabledServices("database") {
		factories["oci_database_advanced_cluster_file_system"] = tf_database.DatabaseAdvancedClusterFileSystemDataSource
		factories["oci_database_advanced_cluster_file_systems"] = tf_database.DatabaseAdvancedClusterFileSystemsDataSource
		factories["oci_database_application_vip"] = tf_database.DatabaseApplicationVipDataSource
		factories["oci_database_application_vips"] = tf_database.DatabaseApplicationVipsDataSource
		factories["oci_database_autonomous_container_database"] = tf_database.DatabaseAutonomousContainerDatabaseDataSource
		factories["oci_database_autonomous_container_database_backup"] = tf_database.DatabaseAutonomousContainerDatabaseBackupDataSource
		factories["oci_database_autonomous_container_database_backup_list_autonomous_databases_in_backups"] = tf_database.DatabaseAutonomousContainerDatabaseBackupListAutonomousDatabasesInBackupsDataSource
		factories["oci_database_autonomous_container_database_backups"] = tf_database.DatabaseAutonomousContainerDatabaseBackupsDataSource
		factories["oci_database_autonomous_container_database_dataguard_association"] = tf_database.DatabaseAutonomousContainerDatabaseDataguardAssociationDataSource
		factories["oci_database_autonomous_container_database_dataguard_associations"] = tf_database.DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource
		factories["oci_database_autonomous_container_database_resource_usage"] = tf_database.DatabaseAutonomousContainerDatabaseResourceUsageDataSource
		factories["oci_database_autonomous_container_database_versions"] = tf_database.DatabaseAutonomousContainerDatabaseVersionsDataSource
		factories["oci_database_autonomous_container_databases"] = tf_database.DatabaseAutonomousContainerDatabasesDataSource
		factories["oci_database_autonomous_container_patches"] = tf_database.DatabaseAutonomousContainerPatchesDataSource
		factories["oci_database_autonomous_database"] = tf_database.DatabaseAutonomousDatabaseDataSource
		factories["oci_database_autonomous_database_available_maintenance_windows"] = tf_database.DatabaseAutonomousDatabaseAvailableMaintenanceWindowsDataSource
		factories["oci_database_autonomous_database_backup"] = tf_database.DatabaseAutonomousDatabaseBackupDataSource
		factories["oci_database_autonomous_database_backups"] = tf_database.DatabaseAutonomousDatabaseBackupsDataSource
		factories["oci_database_autonomous_database_character_sets"] = tf_database.DatabaseAutonomousDatabaseCharacterSetsDataSource
		factories["oci_database_autonomous_database_dataguard_association"] = tf_database.DatabaseAutonomousDatabaseDataguardAssociationDataSource
		factories["oci_database_autonomous_database_dataguard_associations"] = tf_database.DatabaseAutonomousDatabaseDataguardAssociationsDataSource
		factories["oci_database_autonomous_database_instance_wallet_management"] = tf_database.DatabaseAutonomousDatabaseInstanceWalletManagementDataSource
		factories["oci_database_autonomous_database_peers"] = tf_database.DatabaseAutonomousDatabasePeersDataSource
		factories["oci_database_autonomous_database_refreshable_clones"] = tf_database.DatabaseAutonomousDatabaseRefreshableClonesDataSource
		factories["oci_database_autonomous_database_regional_wallet_management"] = tf_database.DatabaseAutonomousDatabaseRegionalWalletManagementDataSource
		factories["oci_database_autonomous_database_resource_pool_members"] = tf_database.DatabaseAutonomousDatabaseResourcePoolMembersDataSource
		factories["oci_database_autonomous_database_software_image"] = tf_database.DatabaseAutonomousDatabaseSoftwareImageDataSource
		factories["oci_database_autonomous_database_software_images"] = tf_database.DatabaseAutonomousDatabaseSoftwareImagesDataSource
		factories["oci_database_autonomous_database_wallet"] = tf_database.DatabaseAutonomousDatabaseWalletDataSource
		factories["oci_database_autonomous_databases"] = tf_database.DatabaseAutonomousDatabasesDataSource
		factories["oci_database_autonomous_databases_clones"] = tf_database.DatabaseAutonomousDatabasesClonesDataSource
		factories["oci_database_autonomous_databases_estimate_cost_savings"] = tf_database.DatabaseAutonomousDatabasesEstimateCostSavingsDataSource
		factories["oci_database_autonomous_db_preview_versions"] = tf_database.DatabaseAutonomousDbPreviewVersionsDataSource
		factories["oci_database_autonomous_db_versions"] = tf_database.DatabaseAutonomousDbVersionsDataSource
		factories["oci_database_autonomous_exadata_infrastructure"] = tf_database.DatabaseAutonomousExadataInfrastructureDataSource
		factories["oci_database_autonomous_exadata_infrastructure_ocpu"] = tf_database.DatabaseAutonomousExadataInfrastructureOcpuDataSource
		factories["oci_database_autonomous_exadata_infrastructure_shapes"] = tf_database.DatabaseAutonomousExadataInfrastructureShapesDataSource
		factories["oci_database_autonomous_exadata_infrastructures"] = tf_database.DatabaseAutonomousExadataInfrastructuresDataSource
		factories["oci_database_autonomous_patch"] = tf_database.DatabaseAutonomousPatchDataSource
		factories["oci_database_autonomous_virtual_machine"] = tf_database.DatabaseAutonomousVirtualMachineDataSource
		factories["oci_database_autonomous_virtual_machines"] = tf_database.DatabaseAutonomousVirtualMachinesDataSource
		factories["oci_database_autonomous_vm_cluster"] = tf_database.DatabaseAutonomousVmClusterDataSource
		factories["oci_database_autonomous_vm_cluster_acd_resource_usages"] = tf_database.DatabaseAutonomousVmClusterAcdResourceUsagesDataSource
		factories["oci_database_autonomous_vm_cluster_resource_usage"] = tf_database.DatabaseAutonomousVmClusterResourceUsageDataSource
		factories["oci_database_autonomous_vm_clusters"] = tf_database.DatabaseAutonomousVmClustersDataSource
		factories["oci_database_backup_destination"] = tf_database.DatabaseBackupDestinationDataSource
		factories["oci_database_backup_destinations"] = tf_database.DatabaseBackupDestinationsDataSource
		factories["oci_database_backups"] = tf_database.DatabaseBackupsDataSource
		factories["oci_database_cloud_autonomous_vm_cluster"] = tf_database.DatabaseCloudAutonomousVmClusterDataSource
		factories["oci_database_cloud_autonomous_vm_cluster_acd_resource_usages"] = tf_database.DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSource
		factories["oci_database_cloud_autonomous_vm_cluster_resource_usage"] = tf_database.DatabaseCloudAutonomousVmClusterResourceUsageDataSource
		factories["oci_database_cloud_autonomous_vm_clusters"] = tf_database.DatabaseCloudAutonomousVmClustersDataSource
		factories["oci_database_cloud_exadata_infrastructure"] = tf_database.DatabaseCloudExadataInfrastructureDataSource
		factories["oci_database_cloud_exadata_infrastructure_un_allocated_resource"] = tf_database.DatabaseCloudExadataInfrastructureUnAllocatedResourceDataSource
		factories["oci_database_cloud_exadata_infrastructures"] = tf_database.DatabaseCloudExadataInfrastructuresDataSource
		factories["oci_database_cloud_vm_cluster"] = tf_database.DatabaseCloudVmClusterDataSource
		factories["oci_database_cloud_vm_cluster_iorm_config"] = tf_database.DatabaseCloudVmClusterIormConfigDataSource
		factories["oci_database_cloud_vm_clusters"] = tf_database.DatabaseCloudVmClustersDataSource
		factories["oci_database_data_guard_association"] = tf_database.DatabaseDataGuardAssociationDataSource
		factories["oci_database_data_guard_associations"] = tf_database.DatabaseDataGuardAssociationsDataSource
		factories["oci_database_database"] = tf_database.DatabaseDatabaseDataSource
		factories["oci_database_database_pdb_conversion_history_entries"] = tf_database.DatabaseDatabasePdbConversionHistoryEntriesDataSource
		factories["oci_database_database_pdb_conversion_history_entry"] = tf_database.DatabaseDatabasePdbConversionHistoryEntryDataSource
		factories["oci_database_database_software_image"] = tf_database.DatabaseDatabaseSoftwareImageDataSource
		factories["oci_database_database_software_images"] = tf_database.DatabaseDatabaseSoftwareImagesDataSource
		factories["oci_database_database_upgrade_history_entries"] = tf_database.DatabaseDatabaseUpgradeHistoryEntriesDataSource
		factories["oci_database_database_upgrade_history_entry"] = tf_database.DatabaseDatabaseUpgradeHistoryEntryDataSource
		factories["oci_database_databases"] = tf_database.DatabaseDatabasesDataSource
		factories["oci_database_db_connection_bundle"] = tf_database.DatabaseDbConnectionBundleDataSource
		factories["oci_database_db_connection_bundles"] = tf_database.DatabaseDbConnectionBundlesDataSource
		factories["oci_database_db_home"] = tf_database.DatabaseDbHomeDataSource
		factories["oci_database_db_home_patch_history_entries"] = tf_database.DatabaseDbHomePatchHistoryEntriesDataSource
		factories["oci_database_db_home_patches"] = tf_database.DatabaseDbHomePatchesDataSource
		factories["oci_database_db_homes"] = tf_database.DatabaseDbHomesDataSource
		factories["oci_database_db_node"] = tf_database.DatabaseDbNodeDataSource
		factories["oci_database_db_node_console_connection"] = tf_database.DatabaseDbNodeConsoleConnectionDataSource
		factories["oci_database_db_node_console_connections"] = tf_database.DatabaseDbNodeConsoleConnectionsDataSource
		factories["oci_database_db_node_console_histories"] = tf_database.DatabaseDbNodeConsoleHistoriesDataSource
		factories["oci_database_db_node_console_history"] = tf_database.DatabaseDbNodeConsoleHistoryDataSource
		factories["oci_database_db_node_console_history_content"] = tf_database.DatabaseDbNodeConsoleHistoryContentDataSource
		factories["oci_database_db_node_snapshot"] = tf_database.DatabaseDbNodeSnapshotDataSource
		factories["oci_database_db_node_snapshots"] = tf_database.DatabaseDbNodeSnapshotsDataSource
		factories["oci_database_db_nodes"] = tf_database.DatabaseDbNodesDataSource
		factories["oci_database_db_server"] = tf_database.DatabaseDbServerDataSource
		factories["oci_database_db_servers"] = tf_database.DatabaseDbServersDataSource
		factories["oci_database_db_system_compute_performances"] = tf_database.DatabaseDbSystemComputePerformancesDataSource
		factories["oci_database_db_system_os_patch_history_entries"] = tf_database.DatabaseDbSystemOsPatchHistoryEntriesDataSource
		factories["oci_database_db_system_os_patch_history_entry"] = tf_database.DatabaseDbSystemOsPatchHistoryEntryDataSource
		factories["oci_database_db_system_patch_history_entries"] = tf_database.DatabaseDbSystemPatchHistoryEntriesDataSource
		factories["oci_database_db_system_patches"] = tf_database.DatabaseDbSystemPatchesDataSource
		factories["oci_database_db_system_shapes"] = tf_database.DatabaseDbSystemShapesDataSource
		factories["oci_database_db_system_storage_performances"] = tf_database.DatabaseDbSystemStoragePerformancesDataSource
		factories["oci_database_db_systems"] = tf_database.DatabaseDbSystemsDataSource
		factories["oci_database_db_systems_upgrade_history_entries"] = tf_database.DatabaseDbSystemsUpgradeHistoryEntriesDataSource
		factories["oci_database_db_systems_upgrade_history_entry"] = tf_database.DatabaseDbSystemsUpgradeHistoryEntryDataSource
		factories["oci_database_db_versions"] = tf_database.DatabaseDbVersionsDataSource
		factories["oci_database_exadata_infrastructure"] = tf_database.DatabaseExadataInfrastructureDataSource
		factories["oci_database_exadata_infrastructure_download_config_file"] = tf_database.DatabaseExadataInfrastructureDownloadConfigFileDataSource
		factories["oci_database_exadata_infrastructure_un_allocated_resource"] = tf_database.DatabaseExadataInfrastructureUnAllocatedResourceDataSource
		factories["oci_database_exadata_infrastructures"] = tf_database.DatabaseExadataInfrastructuresDataSource
		factories["oci_database_exadata_iorm_config"] = tf_database.DatabaseExadataIormConfigDataSource
		factories["oci_database_exadb_vm_cluster"] = tf_database.DatabaseExadbVmClusterDataSource
		factories["oci_database_exadb_vm_cluster_update"] = tf_database.DatabaseExadbVmClusterUpdateDataSource
		factories["oci_database_exadb_vm_cluster_update_history_entries"] = tf_database.DatabaseExadbVmClusterUpdateHistoryEntriesDataSource
		factories["oci_database_exadb_vm_cluster_update_history_entry"] = tf_database.DatabaseExadbVmClusterUpdateHistoryEntryDataSource
		factories["oci_database_exadb_vm_cluster_updates"] = tf_database.DatabaseExadbVmClusterUpdatesDataSource
		factories["oci_database_exadb_vm_clusters"] = tf_database.DatabaseExadbVmClustersDataSource
		factories["oci_database_exascale_db_storage_vault"] = tf_database.DatabaseExascaleDbStorageVaultDataSource
		factories["oci_database_exascale_db_storage_vaults"] = tf_database.DatabaseExascaleDbStorageVaultsDataSource
		factories["oci_database_execution_action"] = tf_database.DatabaseExecutionActionDataSource
		factories["oci_database_execution_actions"] = tf_database.DatabaseExecutionActionsDataSource
		factories["oci_database_execution_window"] = tf_database.DatabaseExecutionWindowDataSource
		factories["oci_database_execution_windows"] = tf_database.DatabaseExecutionWindowsDataSource
		factories["oci_database_external_container_database"] = tf_database.DatabaseExternalContainerDatabaseDataSource
		factories["oci_database_external_container_databases"] = tf_database.DatabaseExternalContainerDatabasesDataSource
		factories["oci_database_external_database_connector"] = tf_database.DatabaseExternalDatabaseConnectorDataSource
		factories["oci_database_external_database_connectors"] = tf_database.DatabaseExternalDatabaseConnectorsDataSource
		factories["oci_database_external_non_container_database"] = tf_database.DatabaseExternalNonContainerDatabaseDataSource
		factories["oci_database_external_non_container_databases"] = tf_database.DatabaseExternalNonContainerDatabasesDataSource
		factories["oci_database_external_pluggable_database"] = tf_database.DatabaseExternalPluggableDatabaseDataSource
		factories["oci_database_external_pluggable_databases"] = tf_database.DatabaseExternalPluggableDatabasesDataSource
		factories["oci_database_flex_components"] = tf_database.DatabaseFlexComponentsDataSource
		factories["oci_database_gi_version_minor_versions"] = tf_database.DatabaseGiVersionMinorVersionsDataSource
		factories["oci_database_gi_versions"] = tf_database.DatabaseGiVersionsDataSource
		factories["oci_database_infrastructure_target_version"] = tf_database.DatabaseInfrastructureTargetVersionDataSource
		factories["oci_database_key_store"] = tf_database.DatabaseKeyStoreDataSource
		factories["oci_database_key_stores"] = tf_database.DatabaseKeyStoresDataSource
		factories["oci_database_maintenance_run"] = tf_database.DatabaseMaintenanceRunDataSource
		factories["oci_database_maintenance_run_histories"] = tf_database.DatabaseMaintenanceRunHistoriesDataSource
		factories["oci_database_maintenance_run_history"] = tf_database.DatabaseMaintenanceRunHistoryDataSource
		factories["oci_database_maintenance_runs"] = tf_database.DatabaseMaintenanceRunsDataSource
		factories["oci_database_oneoff_patch"] = tf_database.DatabaseOneoffPatchDataSource
		factories["oci_database_oneoff_patches"] = tf_database.DatabaseOneoffPatchesDataSource
		factories["oci_database_pluggable_database"] = tf_database.DatabasePluggableDatabaseDataSource
		factories["oci_database_pluggable_database_snapshot"] = tf_database.DatabasePluggableDatabaseSnapshotDataSource
		factories["oci_database_pluggable_database_snapshots"] = tf_database.DatabasePluggableDatabaseSnapshotsDataSource
		factories["oci_database_pluggable_databases"] = tf_database.DatabasePluggableDatabasesDataSource
		factories["oci_database_scheduled_action"] = tf_database.DatabaseScheduledActionDataSource
		factories["oci_database_scheduled_action_params"] = tf_database.DatabaseScheduledActionParamsDataSource
		factories["oci_database_scheduled_actions"] = tf_database.DatabaseScheduledActionsDataSource
		factories["oci_database_scheduling_plan"] = tf_database.DatabaseSchedulingPlanDataSource
		factories["oci_database_scheduling_plans"] = tf_database.DatabaseSchedulingPlansDataSource
		factories["oci_database_scheduling_policies"] = tf_database.DatabaseSchedulingPoliciesDataSource
		factories["oci_database_scheduling_policy"] = tf_database.DatabaseSchedulingPolicyDataSource
		factories["oci_database_scheduling_policy_recommended_scheduled_actions"] = tf_database.DatabaseSchedulingPolicyRecommendedScheduledActionsDataSource
		factories["oci_database_scheduling_policy_scheduling_window"] = tf_database.DatabaseSchedulingPolicySchedulingWindowDataSource
		factories["oci_database_scheduling_policy_scheduling_windows"] = tf_database.DatabaseSchedulingPolicySchedulingWindowsDataSource
		factories["oci_database_system_version_minor_versions"] = tf_database.DatabaseSystemVersionMinorVersionsDataSource
		factories["oci_database_system_versions"] = tf_database.DatabaseSystemVersionsDataSource
		factories["oci_database_vm_cluster"] = tf_database.DatabaseVmClusterDataSource
		factories["oci_database_vm_cluster_network"] = tf_database.DatabaseVmClusterNetworkDataSource
		factories["oci_database_vm_cluster_network_download_config_file"] = tf_database.DatabaseVmClusterNetworkDownloadConfigFileDataSource
		factories["oci_database_vm_cluster_networks"] = tf_database.DatabaseVmClusterNetworksDataSource
		factories["oci_database_vm_cluster_patch"] = tf_database.DatabaseVmClusterPatchDataSource
		factories["oci_database_vm_cluster_patch_history_entries"] = tf_database.DatabaseVmClusterPatchHistoryEntriesDataSource
		factories["oci_database_vm_cluster_patch_history_entry"] = tf_database.DatabaseVmClusterPatchHistoryEntryDataSource
		factories["oci_database_vm_cluster_patches"] = tf_database.DatabaseVmClusterPatchesDataSource
		factories["oci_database_vm_cluster_recommended_network"] = tf_database.DatabaseVmClusterRecommendedNetworkDataSource
		factories["oci_database_vm_cluster_update"] = tf_database.DatabaseVmClusterUpdateDataSource
		factories["oci_database_vm_cluster_update_history_entries"] = tf_database.DatabaseVmClusterUpdateHistoryEntriesDataSource
		factories["oci_database_vm_cluster_update_history_entry"] = tf_database.DatabaseVmClusterUpdateHistoryEntryDataSource
		factories["oci_database_vm_cluster_updates"] = tf_database.DatabaseVmClusterUpdatesDataSource
		factories["oci_database_vm_clusters"] = tf_database.DatabaseVmClustersDataSource
	}
	if common.CheckForEnabledServices("databasemanagement") {
		factories["oci_database_management_cloud_asm"] = tf_database_management.DatabaseManagementCloudAsmDataSource
		factories["oci_database_management_cloud_asm_configuration"] = tf_database_management.DatabaseManagementCloudAsmConfigurationDataSource
		factories["oci_database_management_cloud_asm_disk_groups"] = tf_database_management.DatabaseManagementCloudAsmDiskGroupsDataSource
		factories["oci_database_management_cloud_asm_instance"] = tf_database_management.DatabaseManagementCloudAsmInstanceDataSource
		factories["oci_database_management_cloud_asm_instances"] = tf_database_management.DatabaseManagementCloudAsmInstancesDataSource
		factories["oci_database_management_cloud_asm_users"] = tf_database_management.DatabaseManagementCloudAsmUsersDataSource
		factories["oci_database_management_cloud_asms"] = tf_database_management.DatabaseManagementCloudAsmsDataSource
		factories["oci_database_management_cloud_cluster"] = tf_database_management.DatabaseManagementCloudClusterDataSource
		factories["oci_database_management_cloud_cluster_instance"] = tf_database_management.DatabaseManagementCloudClusterInstanceDataSource
		factories["oci_database_management_cloud_cluster_instances"] = tf_database_management.DatabaseManagementCloudClusterInstancesDataSource
		factories["oci_database_management_cloud_clusters"] = tf_database_management.DatabaseManagementCloudClustersDataSource
		factories["oci_database_management_cloud_databases"] = tf_database_management.DatabaseManagementCloudDatabasesDataSource
		factories["oci_database_management_cloud_db_home"] = tf_database_management.DatabaseManagementCloudDbHomeDataSource
		factories["oci_database_management_cloud_db_homes"] = tf_database_management.DatabaseManagementCloudDbHomesDataSource
		factories["oci_database_management_cloud_db_node"] = tf_database_management.DatabaseManagementCloudDbNodeDataSource
		factories["oci_database_management_cloud_db_nodes"] = tf_database_management.DatabaseManagementCloudDbNodesDataSource
		factories["oci_database_management_cloud_db_system"] = tf_database_management.DatabaseManagementCloudDbSystemDataSource
		factories["oci_database_management_cloud_db_system_connector"] = tf_database_management.DatabaseManagementCloudDbSystemConnectorDataSource
		factories["oci_database_management_cloud_db_system_connectors"] = tf_database_management.DatabaseManagementCloudDbSystemConnectorsDataSource
		factories["oci_database_management_cloud_db_system_discoveries"] = tf_database_management.DatabaseManagementCloudDbSystemDiscoveriesDataSource
		factories["oci_database_management_cloud_db_system_discovery"] = tf_database_management.DatabaseManagementCloudDbSystemDiscoveryDataSource
		factories["oci_database_management_cloud_db_systems"] = tf_database_management.DatabaseManagementCloudDbSystemsDataSource
		factories["oci_database_management_cloud_exadata_infrastructure"] = tf_database_management.DatabaseManagementCloudExadataInfrastructureDataSource
		factories["oci_database_management_cloud_exadata_infrastructures"] = tf_database_management.DatabaseManagementCloudExadataInfrastructuresDataSource
		factories["oci_database_management_cloud_exadata_storage_connector"] = tf_database_management.DatabaseManagementCloudExadataStorageConnectorDataSource
		factories["oci_database_management_cloud_exadata_storage_connectors"] = tf_database_management.DatabaseManagementCloudExadataStorageConnectorsDataSource
		factories["oci_database_management_cloud_exadata_storage_grid"] = tf_database_management.DatabaseManagementCloudExadataStorageGridDataSource
		factories["oci_database_management_cloud_exadata_storage_server"] = tf_database_management.DatabaseManagementCloudExadataStorageServerDataSource
		factories["oci_database_management_cloud_exadata_storage_server_iorm_plan"] = tf_database_management.DatabaseManagementCloudExadataStorageServerIormPlanDataSource
		factories["oci_database_management_cloud_exadata_storage_server_open_alert_history"] = tf_database_management.DatabaseManagementCloudExadataStorageServerOpenAlertHistoryDataSource
		factories["oci_database_management_cloud_exadata_storage_servers"] = tf_database_management.DatabaseManagementCloudExadataStorageServersDataSource
		factories["oci_database_management_cloud_listener"] = tf_database_management.DatabaseManagementCloudListenerDataSource
		factories["oci_database_management_cloud_listener_services"] = tf_database_management.DatabaseManagementCloudListenerServicesDataSource
		factories["oci_database_management_cloud_listeners"] = tf_database_management.DatabaseManagementCloudListenersDataSource
		factories["oci_database_management_db_management_private_endpoint"] = tf_database_management.DatabaseManagementDbManagementPrivateEndpointDataSource
		factories["oci_database_management_db_management_private_endpoint_associated_database"] = tf_database_management.DatabaseManagementDbManagementPrivateEndpointAssociatedDatabaseDataSource
		factories["oci_database_management_db_management_private_endpoint_associated_databases"] = tf_database_management.DatabaseManagementDbManagementPrivateEndpointAssociatedDatabasesDataSource
		factories["oci_database_management_db_management_private_endpoints"] = tf_database_management.DatabaseManagementDbManagementPrivateEndpointsDataSource
		factories["oci_database_management_exadata_infrastructure_fleet_metric"] = tf_database_management.DatabaseManagementExadataInfrastructureFleetMetricDataSource
		factories["oci_database_management_external_asm"] = tf_database_management.DatabaseManagementExternalAsmDataSource
		factories["oci_database_management_external_asm_configuration"] = tf_database_management.DatabaseManagementExternalAsmConfigurationDataSource
		factories["oci_database_management_external_asm_disk_groups"] = tf_database_management.DatabaseManagementExternalAsmDiskGroupsDataSource
		factories["oci_database_management_external_asm_instance"] = tf_database_management.DatabaseManagementExternalAsmInstanceDataSource
		factories["oci_database_management_external_asm_instances"] = tf_database_management.DatabaseManagementExternalAsmInstancesDataSource
		factories["oci_database_management_external_asm_users"] = tf_database_management.DatabaseManagementExternalAsmUsersDataSource
		factories["oci_database_management_external_asms"] = tf_database_management.DatabaseManagementExternalAsmsDataSource
		factories["oci_database_management_external_cluster"] = tf_database_management.DatabaseManagementExternalClusterDataSource
		factories["oci_database_management_external_cluster_instance"] = tf_database_management.DatabaseManagementExternalClusterInstanceDataSource
		factories["oci_database_management_external_cluster_instances"] = tf_database_management.DatabaseManagementExternalClusterInstancesDataSource
		factories["oci_database_management_external_clusters"] = tf_database_management.DatabaseManagementExternalClustersDataSource
		factories["oci_database_management_external_databases"] = tf_database_management.DatabaseManagementExternalDatabasesDataSource
		factories["oci_database_management_external_db_home"] = tf_database_management.DatabaseManagementExternalDbHomeDataSource
		factories["oci_database_management_external_db_homes"] = tf_database_management.DatabaseManagementExternalDbHomesDataSource
		factories["oci_database_management_external_db_node"] = tf_database_management.DatabaseManagementExternalDbNodeDataSource
		factories["oci_database_management_external_db_nodes"] = tf_database_management.DatabaseManagementExternalDbNodesDataSource
		factories["oci_database_management_external_db_system"] = tf_database_management.DatabaseManagementExternalDbSystemDataSource
		factories["oci_database_management_external_db_system_connector"] = tf_database_management.DatabaseManagementExternalDbSystemConnectorDataSource
		factories["oci_database_management_external_db_system_connectors"] = tf_database_management.DatabaseManagementExternalDbSystemConnectorsDataSource
		factories["oci_database_management_external_db_system_discoveries"] = tf_database_management.DatabaseManagementExternalDbSystemDiscoveriesDataSource
		factories["oci_database_management_external_db_system_discovery"] = tf_database_management.DatabaseManagementExternalDbSystemDiscoveryDataSource
		factories["oci_database_management_external_db_systems"] = tf_database_management.DatabaseManagementExternalDbSystemsDataSource
		factories["oci_database_management_external_exadata_infrastructure"] = tf_database_management.DatabaseManagementExternalExadataInfrastructureDataSource
		factories["oci_database_management_external_exadata_infrastructures"] = tf_database_management.DatabaseManagementExternalExadataInfrastructuresDataSource
		factories["oci_database_management_external_exadata_storage_connector"] = tf_database_management.DatabaseManagementExternalExadataStorageConnectorDataSource
		factories["oci_database_management_external_exadata_storage_connectors"] = tf_database_management.DatabaseManagementExternalExadataStorageConnectorsDataSource
		factories["oci_database_management_external_exadata_storage_grid"] = tf_database_management.DatabaseManagementExternalExadataStorageGridDataSource
		factories["oci_database_management_external_exadata_storage_server"] = tf_database_management.DatabaseManagementExternalExadataStorageServerDataSource
		factories["oci_database_management_external_exadata_storage_server_iorm_plan"] = tf_database_management.DatabaseManagementExternalExadataStorageServerIormPlanDataSource
		factories["oci_database_management_external_exadata_storage_server_open_alert_history"] = tf_database_management.DatabaseManagementExternalExadataStorageServerOpenAlertHistoryDataSource
		factories["oci_database_management_external_exadata_storage_server_top_sql_cpu_activity"] = tf_database_management.DatabaseManagementExternalExadataStorageServerTopSqlCpuActivityDataSource
		factories["oci_database_management_external_exadata_storage_servers"] = tf_database_management.DatabaseManagementExternalExadataStorageServersDataSource
		factories["oci_database_management_external_listener"] = tf_database_management.DatabaseManagementExternalListenerDataSource
		factories["oci_database_management_external_listener_services"] = tf_database_management.DatabaseManagementExternalListenerServicesDataSource
		factories["oci_database_management_external_listeners"] = tf_database_management.DatabaseManagementExternalListenersDataSource
		factories["oci_database_management_external_my_sql_database"] = tf_database_management.DatabaseManagementExternalMySqlDatabaseDataSource
		factories["oci_database_management_external_my_sql_database_connector"] = tf_database_management.DatabaseManagementExternalMySqlDatabaseConnectorDataSource
		factories["oci_database_management_external_my_sql_database_connectors"] = tf_database_management.DatabaseManagementExternalMySqlDatabaseConnectorsDataSource
		factories["oci_database_management_external_my_sql_databases"] = tf_database_management.DatabaseManagementExternalMySqlDatabasesDataSource
		factories["oci_database_management_job_executions_status"] = tf_database_management.DatabaseManagementJobExecutionsStatusDataSource
		factories["oci_database_management_job_executions_statuses"] = tf_database_management.DatabaseManagementJobExecutionsStatusesDataSource
		factories["oci_database_management_managed_database"] = tf_database_management.DatabaseManagementManagedDatabaseDataSource
		factories["oci_database_management_managed_database_addm_task"] = tf_database_management.DatabaseManagementManagedDatabaseAddmTaskDataSource
		factories["oci_database_management_managed_database_addm_tasks"] = tf_database_management.DatabaseManagementManagedDatabaseAddmTasksDataSource
		factories["oci_database_management_managed_database_alert_log_count"] = tf_database_management.DatabaseManagementManagedDatabaseAlertLogCountDataSource
		factories["oci_database_management_managed_database_alert_log_counts"] = tf_database_management.DatabaseManagementManagedDatabaseAlertLogCountsDataSource
		factories["oci_database_management_managed_database_attention_log_count"] = tf_database_management.DatabaseManagementManagedDatabaseAttentionLogCountDataSource
		factories["oci_database_management_managed_database_attention_log_counts"] = tf_database_management.DatabaseManagementManagedDatabaseAttentionLogCountsDataSource
		factories["oci_database_management_managed_database_cursor_cache_statements"] = tf_database_management.DatabaseManagementManagedDatabaseCursorCacheStatementsDataSource
		factories["oci_database_management_managed_database_group"] = tf_database_management.DatabaseManagementManagedDatabaseGroupDataSource
		factories["oci_database_management_managed_database_groups"] = tf_database_management.DatabaseManagementManagedDatabaseGroupsDataSource
		factories["oci_database_management_managed_database_optimizer_statistics_advisor_execution"] = tf_database_management.DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionDataSource
		factories["oci_database_management_managed_database_optimizer_statistics_advisor_execution_script"] = tf_database_management.DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionScriptDataSource
		factories["oci_database_management_managed_database_optimizer_statistics_advisor_executions"] = tf_database_management.DatabaseManagementManagedDatabaseOptimizerStatisticsAdvisorExecutionsDataSource
		factories["oci_database_management_managed_database_optimizer_statistics_collection_aggregations"] = tf_database_management.DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionAggregationsDataSource
		factories["oci_database_management_managed_database_optimizer_statistics_collection_operation"] = tf_database_management.DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationDataSource
		factories["oci_database_management_managed_database_optimizer_statistics_collection_operations"] = tf_database_management.DatabaseManagementManagedDatabaseOptimizerStatisticsCollectionOperationsDataSource
		factories["oci_database_management_managed_database_preferred_credential"] = tf_database_management.DatabaseManagementManagedDatabasePreferredCredentialDataSource
		factories["oci_database_management_managed_database_preferred_credentials"] = tf_database_management.DatabaseManagementManagedDatabasePreferredCredentialsDataSource
		factories["oci_database_management_managed_database_sql_plan_baseline"] = tf_database_management.DatabaseManagementManagedDatabaseSqlPlanBaselineDataSource
		factories["oci_database_management_managed_database_sql_plan_baseline_configuration"] = tf_database_management.DatabaseManagementManagedDatabaseSqlPlanBaselineConfigurationDataSource
		factories["oci_database_management_managed_database_sql_plan_baseline_jobs"] = tf_database_management.DatabaseManagementManagedDatabaseSqlPlanBaselineJobsDataSource
		factories["oci_database_management_managed_database_sql_plan_baselines"] = tf_database_management.DatabaseManagementManagedDatabaseSqlPlanBaselinesDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_task"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTaskDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks_execution_plan_stats_comparision"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksExecutionPlanStatsComparisionDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks_finding"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks_findings"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksFindingsDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendation"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks_recommendations"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksRecommendationsDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks_sql_execution_plan"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSqlExecutionPlanDataSource
		factories["oci_database_management_managed_database_sql_tuning_advisor_tasks_summary_report"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningAdvisorTasksSummaryReportDataSource
		factories["oci_database_management_managed_database_sql_tuning_set"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningSetDataSource
		factories["oci_database_management_managed_database_sql_tuning_sets"] = tf_database_management.DatabaseManagementManagedDatabaseSqlTuningSetsDataSource
		factories["oci_database_management_managed_database_table_statistics"] = tf_database_management.DatabaseManagementManagedDatabaseTableStatisticsDataSource
		factories["oci_database_management_managed_database_user"] = tf_database_management.DatabaseManagementManagedDatabaseUserDataSource
		factories["oci_database_management_managed_database_user_consumer_group_privilege"] = tf_database_management.DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegeDataSource
		factories["oci_database_management_managed_database_user_consumer_group_privileges"] = tf_database_management.DatabaseManagementManagedDatabaseUserConsumerGroupPrivilegesDataSource
		factories["oci_database_management_managed_database_user_data_access_container"] = tf_database_management.DatabaseManagementManagedDatabaseUserDataAccessContainerDataSource
		factories["oci_database_management_managed_database_user_data_access_containers"] = tf_database_management.DatabaseManagementManagedDatabaseUserDataAccessContainersDataSource
		factories["oci_database_management_managed_database_user_object_privilege"] = tf_database_management.DatabaseManagementManagedDatabaseUserObjectPrivilegeDataSource
		factories["oci_database_management_managed_database_user_object_privileges"] = tf_database_management.DatabaseManagementManagedDatabaseUserObjectPrivilegesDataSource
		factories["oci_database_management_managed_database_user_proxied_for_user"] = tf_database_management.DatabaseManagementManagedDatabaseUserProxiedForUserDataSource
		factories["oci_database_management_managed_database_user_proxied_for_users"] = tf_database_management.DatabaseManagementManagedDatabaseUserProxiedForUsersDataSource
		factories["oci_database_management_managed_database_user_role"] = tf_database_management.DatabaseManagementManagedDatabaseUserRoleDataSource
		factories["oci_database_management_managed_database_user_roles"] = tf_database_management.DatabaseManagementManagedDatabaseUserRolesDataSource
		factories["oci_database_management_managed_database_users"] = tf_database_management.DatabaseManagementManagedDatabaseUsersDataSource
		factories["oci_database_management_managed_databases"] = tf_database_management.DatabaseManagementManagedDatabasesDataSource
		factories["oci_database_management_managed_databases_asm_properties"] = tf_database_management.DatabaseManagementManagedDatabasesAsmPropertiesDataSource
		factories["oci_database_management_managed_databases_asm_property"] = tf_database_management.DatabaseManagementManagedDatabasesAsmPropertyDataSource
		factories["oci_database_management_managed_databases_database_parameter"] = tf_database_management.DatabaseManagementManagedDatabasesDatabaseParameterDataSource
		factories["oci_database_management_managed_databases_database_parameters"] = tf_database_management.DatabaseManagementManagedDatabasesDatabaseParametersDataSource
		factories["oci_database_management_managed_databases_user_proxy_user"] = tf_database_management.DatabaseManagementManagedDatabasesUserProxyUserDataSource
		factories["oci_database_management_managed_databases_user_proxy_users"] = tf_database_management.DatabaseManagementManagedDatabasesUserProxyUsersDataSource
		factories["oci_database_management_managed_databases_user_system_privilege"] = tf_database_management.DatabaseManagementManagedDatabasesUserSystemPrivilegeDataSource
		factories["oci_database_management_managed_databases_user_system_privileges"] = tf_database_management.DatabaseManagementManagedDatabasesUserSystemPrivilegesDataSource
		factories["oci_database_management_managed_my_sql_database"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseDataSource
		factories["oci_database_management_managed_my_sql_database_binary_log_information"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseBinaryLogInformationDataSource
		factories["oci_database_management_managed_my_sql_database_configuration_data"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseConfigurationDataDataSource
		factories["oci_database_management_managed_my_sql_database_digest_errors"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseDigestErrorsDataSource
		factories["oci_database_management_managed_my_sql_database_general_replication_information"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseGeneralReplicationInformationDataSource
		factories["oci_database_management_managed_my_sql_database_high_availability_members"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseHighAvailabilityMembersDataSource
		factories["oci_database_management_managed_my_sql_database_inbound_replications"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseInboundReplicationsDataSource
		factories["oci_database_management_managed_my_sql_database_outbound_replications"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseOutboundReplicationsDataSource
		factories["oci_database_management_managed_my_sql_database_query_detail"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseQueryDetailDataSource
		factories["oci_database_management_managed_my_sql_database_sql_data"] = tf_database_management.DatabaseManagementManagedMySqlDatabaseSqlDataDataSource
		factories["oci_database_management_managed_my_sql_databases"] = tf_database_management.DatabaseManagementManagedMySqlDatabasesDataSource
		factories["oci_database_management_named_credential"] = tf_database_management.DatabaseManagementNamedCredentialDataSource
		factories["oci_database_management_named_credentials"] = tf_database_management.DatabaseManagementNamedCredentialsDataSource
	}
	if common.CheckForEnabledServices("databasemigration") {
		factories["oci_database_migration_assessment"] = tf_database_migration.DatabaseMigrationAssessmentDataSource
		factories["oci_database_migration_assessment_assessor"] = tf_database_migration.DatabaseMigrationAssessmentAssessorDataSource
		factories["oci_database_migration_assessment_assessor_check"] = tf_database_migration.DatabaseMigrationAssessmentAssessorCheckDataSource
		factories["oci_database_migration_assessment_assessor_check_affected_objects"] = tf_database_migration.DatabaseMigrationAssessmentAssessorCheckAffectedObjectsDataSource
		factories["oci_database_migration_assessment_assessor_checks"] = tf_database_migration.DatabaseMigrationAssessmentAssessorChecksDataSource
		factories["oci_database_migration_assessment_assessors"] = tf_database_migration.DatabaseMigrationAssessmentAssessorsDataSource
		factories["oci_database_migration_assessment_object_types"] = tf_database_migration.DatabaseMigrationAssessmentObjectTypesDataSource
		factories["oci_database_migration_assessments"] = tf_database_migration.DatabaseMigrationAssessmentsDataSource
		factories["oci_database_migration_connection"] = tf_database_migration.DatabaseMigrationConnectionDataSource
		factories["oci_database_migration_connection_databaseconnectiontypes"] = tf_database_migration.DatabaseMigrationConnectionDatabaseconnectiontypesDataSource
		factories["oci_database_migration_connections"] = tf_database_migration.DatabaseMigrationConnectionsDataSource
		factories["oci_database_migration_job"] = tf_database_migration.DatabaseMigrationJobDataSource
		factories["oci_database_migration_job_advisor_report"] = tf_database_migration.DatabaseMigrationJobAdvisorReportDataSource
		factories["oci_database_migration_job_advisor_report_check_objects"] = tf_database_migration.DatabaseMigrationJobAdvisorReportCheckObjectsDataSource
		factories["oci_database_migration_job_advisor_report_checks"] = tf_database_migration.DatabaseMigrationJobAdvisorReportChecksDataSource
		factories["oci_database_migration_job_output"] = tf_database_migration.DatabaseMigrationJobOutputDataSource
		factories["oci_database_migration_jobs"] = tf_database_migration.DatabaseMigrationJobsDataSource
		factories["oci_database_migration_migration"] = tf_database_migration.DatabaseMigrationMigrationDataSource
		factories["oci_database_migration_migration_object_types"] = tf_database_migration.DatabaseMigrationMigrationObjectTypesDataSource
		factories["oci_database_migration_migrations"] = tf_database_migration.DatabaseMigrationMigrationDataSource
		factories["oci_database_migration_script"] = tf_database_migration.DatabaseMigrationScriptDataSource
	}
	if common.CheckForEnabledServices("databasetools") {
		factories["oci_database_tools_database_tools_connection"] = tf_database_tools.DatabaseToolsDatabaseToolsConnectionDataSource
		factories["oci_database_tools_database_tools_connections"] = tf_database_tools.DatabaseToolsDatabaseToolsConnectionsDataSource
		factories["oci_database_tools_database_tools_database_api_gateway_config"] = tf_database_tools.DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigDataSource
		factories["oci_database_tools_database_tools_database_api_gateway_configs"] = tf_database_tools.DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigsDataSource
		factories["oci_database_tools_database_tools_endpoint_service"] = tf_database_tools.DatabaseToolsDatabaseToolsEndpointServiceDataSource
		factories["oci_database_tools_database_tools_endpoint_services"] = tf_database_tools.DatabaseToolsDatabaseToolsEndpointServicesDataSource
		factories["oci_database_tools_database_tools_identities"] = tf_database_tools.DatabaseToolsDatabaseToolsIdentitiesDataSource
		factories["oci_database_tools_database_tools_identity"] = tf_database_tools.DatabaseToolsDatabaseToolsIdentityDataSource
		factories["oci_database_tools_database_tools_mcp_server"] = tf_database_tools.DatabaseToolsDatabaseToolsMcpServerDataSource
		factories["oci_database_tools_database_tools_mcp_servers"] = tf_database_tools.DatabaseToolsDatabaseToolsMcpServersDataSource
		factories["oci_database_tools_database_tools_mcp_toolset"] = tf_database_tools.DatabaseToolsDatabaseToolsMcpToolsetDataSource
		factories["oci_database_tools_database_tools_mcp_toolset_versions"] = tf_database_tools.DatabaseToolsDatabaseToolsMcpToolsetVersionsDataSource
		factories["oci_database_tools_database_tools_mcp_toolsets"] = tf_database_tools.DatabaseToolsDatabaseToolsMcpToolsetsDataSource
		factories["oci_database_tools_database_tools_private_endpoint"] = tf_database_tools.DatabaseToolsDatabaseToolsPrivateEndpointDataSource
		factories["oci_database_tools_database_tools_private_endpoints"] = tf_database_tools.DatabaseToolsDatabaseToolsPrivateEndpointsDataSource
		factories["oci_database_tools_database_tools_sql_report"] = tf_database_tools.DatabaseToolsDatabaseToolsSqlReportDataSource
		factories["oci_database_tools_database_tools_sql_reports"] = tf_database_tools.DatabaseToolsDatabaseToolsSqlReportsDataSource
	}
	if common.CheckForEnabledServices("databasetoolsruntime") {
		factories["oci_database_tools_runtime_database_tools_connection_credential"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialDataSource
		factories["oci_database_tools_runtime_database_tools_connection_credential_execute_grantee"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeDataSource
		factories["oci_database_tools_runtime_database_tools_connection_credential_execute_grantees"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteesDataSource
		factories["oci_database_tools_runtime_database_tools_connection_credential_public_synonym"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymDataSource
		factories["oci_database_tools_runtime_database_tools_connection_credential_public_synonyms"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymsDataSource
		factories["oci_database_tools_runtime_database_tools_connection_credentials"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialsDataSource
		factories["oci_database_tools_runtime_database_tools_connection_property_set"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetDataSource
		factories["oci_database_tools_runtime_database_tools_connection_user_credential"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialDataSource
		factories["oci_database_tools_runtime_database_tools_connection_user_credentials"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionUserCredentialsDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_advanced_properties"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigAdvancedPropertiesDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_content"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigContentDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_global"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_specs"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecsDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_specs"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecsDataSource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pools"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolsDataSource
	}
	if common.CheckForEnabledServices("datacatalog") {
		factories["oci_datacatalog_catalog"] = tf_datacatalog.DatacatalogCatalogDataSource
		factories["oci_datacatalog_catalog_private_endpoint"] = tf_datacatalog.DatacatalogCatalogPrivateEndpointDataSource
		factories["oci_datacatalog_catalog_private_endpoints"] = tf_datacatalog.DatacatalogCatalogPrivateEndpointsDataSource
		factories["oci_datacatalog_catalog_type"] = tf_datacatalog.DatacatalogCatalogTypeDataSource
		factories["oci_datacatalog_catalog_types"] = tf_datacatalog.DatacatalogCatalogTypesDataSource
		factories["oci_datacatalog_catalogs"] = tf_datacatalog.DatacatalogCatalogsDataSource
		factories["oci_datacatalog_connection"] = tf_datacatalog.DatacatalogConnectionDataSource
		factories["oci_datacatalog_connections"] = tf_datacatalog.DatacatalogConnectionsDataSource
		factories["oci_datacatalog_data_asset"] = tf_datacatalog.DatacatalogDataAssetDataSource
		factories["oci_datacatalog_data_assets"] = tf_datacatalog.DatacatalogDataAssetsDataSource
		factories["oci_datacatalog_metastore"] = tf_datacatalog.DatacatalogMetastoreDataSource
		factories["oci_datacatalog_metastores"] = tf_datacatalog.DatacatalogMetastoresDataSource
	}
	if common.CheckForEnabledServices("datacc") {
		factories["oci_datacc_infrastructure"] = tf_datacc.DataccInfrastructureDataSource
		factories["oci_datacc_infrastructure_scale_option"] = tf_datacc.DataccInfrastructureScaleOptionDataSource
		factories["oci_datacc_infrastructures"] = tf_datacc.DataccInfrastructuresDataSource
		factories["oci_datacc_maintenance_execution"] = tf_datacc.DataccMaintenanceExecutionDataSource
		factories["oci_datacc_maintenance_executions"] = tf_datacc.DataccMaintenanceExecutionsDataSource
		factories["oci_datacc_vm_cluster_network"] = tf_datacc.DataccVmClusterNetworkDataSource
		factories["oci_datacc_vm_cluster_networks"] = tf_datacc.DataccVmClusterNetworksDataSource
		factories["oci_datacc_vm_instance"] = tf_datacc.DataccVmInstanceDataSource
		factories["oci_datacc_vm_instances"] = tf_datacc.DataccVmInstancesDataSource
	}
	if common.CheckForEnabledServices("dataflow") {
		factories["oci_dataflow_application"] = tf_dataflow.DataflowApplicationDataSource
		factories["oci_dataflow_applications"] = tf_dataflow.DataflowApplicationsDataSource
		factories["oci_dataflow_invoke_run"] = tf_dataflow.DataflowInvokeRunDataSource
		factories["oci_dataflow_invoke_runs"] = tf_dataflow.DataflowInvokeRunsDataSource
		factories["oci_dataflow_pool"] = tf_dataflow.DataflowPoolDataSource
		factories["oci_dataflow_pools"] = tf_dataflow.DataflowPoolsDataSource
		factories["oci_dataflow_private_endpoint"] = tf_dataflow.DataflowPrivateEndpointDataSource
		factories["oci_dataflow_private_endpoints"] = tf_dataflow.DataflowPrivateEndpointsDataSource
		factories["oci_dataflow_run_log"] = tf_dataflow.DataflowRunLogDataSource
		factories["oci_dataflow_run_logs"] = tf_dataflow.DataflowRunLogsDataSource
		factories["oci_dataflow_run_statement"] = tf_dataflow.DataflowRunStatementDataSource
		factories["oci_dataflow_run_statements"] = tf_dataflow.DataflowRunStatementsDataSource
		factories["oci_dataflow_sql_endpoint"] = tf_dataflow.DataflowSqlEndpointDataSource
		factories["oci_dataflow_sql_endpoints"] = tf_dataflow.DataflowSqlEndpointsDataSource
	}
	if common.CheckForEnabledServices("dataintegration") {
		factories["oci_dataintegration_workspace"] = tf_dataintegration.DataintegrationWorkspaceDataSource
		factories["oci_dataintegration_workspace_application"] = tf_dataintegration.DataintegrationWorkspaceApplicationDataSource
		factories["oci_dataintegration_workspace_application_patch"] = tf_dataintegration.DataintegrationWorkspaceApplicationPatchDataSource
		factories["oci_dataintegration_workspace_application_patches"] = tf_dataintegration.DataintegrationWorkspaceApplicationPatchesDataSource
		factories["oci_dataintegration_workspace_application_schedule"] = tf_dataintegration.DataintegrationWorkspaceApplicationScheduleDataSource
		factories["oci_dataintegration_workspace_application_schedules"] = tf_dataintegration.DataintegrationWorkspaceApplicationSchedulesDataSource
		factories["oci_dataintegration_workspace_application_task_schedule"] = tf_dataintegration.DataintegrationWorkspaceApplicationTaskScheduleDataSource
		factories["oci_dataintegration_workspace_application_task_schedules"] = tf_dataintegration.DataintegrationWorkspaceApplicationTaskSchedulesDataSource
		factories["oci_dataintegration_workspace_applications"] = tf_dataintegration.DataintegrationWorkspaceApplicationsDataSource
		factories["oci_dataintegration_workspace_export_request"] = tf_dataintegration.DataintegrationWorkspaceExportRequestDataSource
		factories["oci_dataintegration_workspace_export_requests"] = tf_dataintegration.DataintegrationWorkspaceExportRequestsDataSource
		factories["oci_dataintegration_workspace_folder"] = tf_dataintegration.DataintegrationWorkspaceFolderDataSource
		factories["oci_dataintegration_workspace_folders"] = tf_dataintegration.DataintegrationWorkspaceFoldersDataSource
		factories["oci_dataintegration_workspace_import_request"] = tf_dataintegration.DataintegrationWorkspaceImportRequestDataSource
		factories["oci_dataintegration_workspace_import_requests"] = tf_dataintegration.DataintegrationWorkspaceImportRequestsDataSource
		factories["oci_dataintegration_workspace_project"] = tf_dataintegration.DataintegrationWorkspaceProjectDataSource
		factories["oci_dataintegration_workspace_projects"] = tf_dataintegration.DataintegrationWorkspaceProjectsDataSource
		factories["oci_dataintegration_workspace_task"] = tf_dataintegration.DataintegrationWorkspaceTaskDataSource
		factories["oci_dataintegration_workspace_tasks"] = tf_dataintegration.DataintegrationWorkspaceTasksDataSource
		factories["oci_dataintegration_workspaces"] = tf_dataintegration.DataintegrationWorkspacesDataSource
	}
	if common.CheckForEnabledServices("datascience") {
		factories["oci_datascience_compute_target"] = tf_datascience.DatascienceComputeTargetDataSource
		factories["oci_datascience_compute_target_shapes"] = tf_datascience.DatascienceComputeTargetShapesDataSource
		factories["oci_datascience_compute_targets"] = tf_datascience.DatascienceComputeTargetsDataSource
		factories["oci_datascience_containers"] = tf_datascience.DatascienceContainersDataSource
		factories["oci_datascience_fast_launch_job_configs"] = tf_datascience.DatascienceFastLaunchJobConfigsDataSource
		factories["oci_datascience_job"] = tf_datascience.DatascienceJobDataSource
		factories["oci_datascience_job_run"] = tf_datascience.DatascienceJobRunDataSource
		factories["oci_datascience_job_runs"] = tf_datascience.DatascienceJobRunsDataSource
		factories["oci_datascience_job_shapes"] = tf_datascience.DatascienceJobShapesDataSource
		factories["oci_datascience_jobs"] = tf_datascience.DatascienceJobsDataSource
		factories["oci_datascience_ml_application"] = tf_datascience.DatascienceMlApplicationDataSource
		factories["oci_datascience_ml_application_implementation"] = tf_datascience.DatascienceMlApplicationImplementationDataSource
		factories["oci_datascience_ml_application_implementation_version"] = tf_datascience.DatascienceMlApplicationImplementationVersionDataSource
		factories["oci_datascience_ml_application_implementation_versions"] = tf_datascience.DatascienceMlApplicationImplementationVersionsDataSource
		factories["oci_datascience_ml_application_implementations"] = tf_datascience.DatascienceMlApplicationImplementationsDataSource
		factories["oci_datascience_ml_application_instance"] = tf_datascience.DatascienceMlApplicationInstanceDataSource
		factories["oci_datascience_ml_application_instances"] = tf_datascience.DatascienceMlApplicationInstancesDataSource
		factories["oci_datascience_ml_applications"] = tf_datascience.DatascienceMlApplicationsDataSource
		factories["oci_datascience_model"] = tf_datascience.DatascienceModelDataSource
		factories["oci_datascience_model_custom_metadata_artifact_content"] = tf_datascience.DatascienceModelCustomMetadataArtifactContentDataSource
		factories["oci_datascience_model_defined_metadata_artifact_content"] = tf_datascience.DatascienceModelDefinedMetadataArtifactContentDataSource
		factories["oci_datascience_model_deployment"] = tf_datascience.DatascienceModelDeploymentDataSource
		factories["oci_datascience_model_deployment_model_states"] = tf_datascience.DatascienceModelDeploymentModelStatesDataSource
		factories["oci_datascience_model_deployment_shapes"] = tf_datascience.DatascienceModelDeploymentShapesDataSource
		factories["oci_datascience_model_deployments"] = tf_datascience.DatascienceModelDeploymentsDataSource
		factories["oci_datascience_model_group"] = tf_datascience.DatascienceModelGroupDataSource
		factories["oci_datascience_model_group_artifact_content"] = tf_datascience.DatascienceModelGroupArtifactContentDataSource
		factories["oci_datascience_model_group_models"] = tf_datascience.DatascienceModelGroupModelsDataSource
		factories["oci_datascience_model_group_version_histories"] = tf_datascience.DatascienceModelGroupVersionHistoriesDataSource
		factories["oci_datascience_model_group_version_history"] = tf_datascience.DatascienceModelGroupVersionHistoryDataSource
		factories["oci_datascience_model_groups"] = tf_datascience.DatascienceModelGroupsDataSource
		factories["oci_datascience_model_provenance"] = tf_datascience.DatascienceModelProvenanceDataSource
		factories["oci_datascience_model_version_set"] = tf_datascience.DatascienceModelVersionSetDataSource
		factories["oci_datascience_model_version_sets"] = tf_datascience.DatascienceModelVersionSetsDataSource
		factories["oci_datascience_models"] = tf_datascience.DatascienceModelsDataSource
		factories["oci_datascience_notebook_session"] = tf_datascience.DatascienceNotebookSessionDataSource
		factories["oci_datascience_notebook_session_shapes"] = tf_datascience.DatascienceNotebookSessionShapesDataSource
		factories["oci_datascience_notebook_sessions"] = tf_datascience.DatascienceNotebookSessionsDataSource
		factories["oci_datascience_pipeline"] = tf_datascience.DatasciencePipelineDataSource
		factories["oci_datascience_pipeline_run"] = tf_datascience.DatasciencePipelineRunDataSource
		factories["oci_datascience_pipeline_runs"] = tf_datascience.DatasciencePipelineRunsDataSource
		factories["oci_datascience_pipelines"] = tf_datascience.DatasciencePipelinesDataSource
		factories["oci_datascience_private_endpoint"] = tf_datascience.DatasciencePrivateEndpointDataSource
		factories["oci_datascience_private_endpoints"] = tf_datascience.DatasciencePrivateEndpointsDataSource
		factories["oci_datascience_project"] = tf_datascience.DatascienceProjectDataSource
		factories["oci_datascience_projects"] = tf_datascience.DatascienceProjectsDataSource
		factories["oci_datascience_schedule"] = tf_datascience.DatascienceScheduleDataSource
		factories["oci_datascience_schedules"] = tf_datascience.DatascienceSchedulesDataSource
	}
	if common.CheckForEnabledServices("dbmulticloud") {
		factories["oci_dbmulticloud_multi_cloud_resource_discoveries"] = tf_dbmulticloud.DbmulticloudMultiCloudResourceDiscoveriesDataSource
		factories["oci_dbmulticloud_multi_cloud_resource_discovery"] = tf_dbmulticloud.DbmulticloudMultiCloudResourceDiscoveryDataSource
		factories["oci_dbmulticloud_oracle_db_aws_identity_connector"] = tf_dbmulticloud.DbmulticloudOracleDbAwsIdentityConnectorDataSource
		factories["oci_dbmulticloud_oracle_db_aws_identity_connectors"] = tf_dbmulticloud.DbmulticloudOracleDbAwsIdentityConnectorsDataSource
		factories["oci_dbmulticloud_oracle_db_aws_key"] = tf_dbmulticloud.DbmulticloudOracleDbAwsKeyDataSource
		factories["oci_dbmulticloud_oracle_db_aws_keys"] = tf_dbmulticloud.DbmulticloudOracleDbAwsKeysDataSource
		factories["oci_dbmulticloud_oracle_db_azure_blob_container"] = tf_dbmulticloud.DbmulticloudOracleDbAzureBlobContainerDataSource
		factories["oci_dbmulticloud_oracle_db_azure_blob_containers"] = tf_dbmulticloud.DbmulticloudOracleDbAzureBlobContainersDataSource
		factories["oci_dbmulticloud_oracle_db_azure_blob_mount"] = tf_dbmulticloud.DbmulticloudOracleDbAzureBlobMountDataSource
		factories["oci_dbmulticloud_oracle_db_azure_blob_mounts"] = tf_dbmulticloud.DbmulticloudOracleDbAzureBlobMountsDataSource
		factories["oci_dbmulticloud_oracle_db_azure_connector"] = tf_dbmulticloud.DbmulticloudOracleDbAzureConnectorDataSource
		factories["oci_dbmulticloud_oracle_db_azure_connectors"] = tf_dbmulticloud.DbmulticloudOracleDbAzureConnectorsDataSource
		factories["oci_dbmulticloud_oracle_db_azure_key"] = tf_dbmulticloud.DbmulticloudOracleDbAzureKeyDataSource
		factories["oci_dbmulticloud_oracle_db_azure_keys"] = tf_dbmulticloud.DbmulticloudOracleDbAzureKeysDataSource
		factories["oci_dbmulticloud_oracle_db_azure_vault"] = tf_dbmulticloud.DbmulticloudOracleDbAzureVaultDataSource
		factories["oci_dbmulticloud_oracle_db_azure_vault_association"] = tf_dbmulticloud.DbmulticloudOracleDbAzureVaultAssociationDataSource
		factories["oci_dbmulticloud_oracle_db_azure_vault_associations"] = tf_dbmulticloud.DbmulticloudOracleDbAzureVaultAssociationsDataSource
		factories["oci_dbmulticloud_oracle_db_azure_vaults"] = tf_dbmulticloud.DbmulticloudOracleDbAzureVaultsDataSource
		factories["oci_dbmulticloud_oracle_db_gcp_identity_connector"] = tf_dbmulticloud.DbmulticloudOracleDbGcpIdentityConnectorDataSource
		factories["oci_dbmulticloud_oracle_db_gcp_identity_connectors"] = tf_dbmulticloud.DbmulticloudOracleDbGcpIdentityConnectorsDataSource
		factories["oci_dbmulticloud_oracle_db_gcp_key"] = tf_dbmulticloud.DbmulticloudOracleDbGcpKeyDataSource
		factories["oci_dbmulticloud_oracle_db_gcp_key_ring"] = tf_dbmulticloud.DbmulticloudOracleDbGcpKeyRingDataSource
		factories["oci_dbmulticloud_oracle_db_gcp_key_rings"] = tf_dbmulticloud.DbmulticloudOracleDbGcpKeyRingsDataSource
		factories["oci_dbmulticloud_oracle_db_gcp_keys"] = tf_dbmulticloud.DbmulticloudOracleDbGcpKeysDataSource
	}
	if common.CheckForEnabledServices("dblm") {
		factories["oci_dblm_patch_management"] = tf_dblm.DblmPatchManagementDataSource
		factories["oci_dblm_patch_management_databases"] = tf_dblm.DblmPatchManagementDatabasesDataSource
		factories["oci_dblm_vulnerability"] = tf_dblm.DblmVulnerabilityDataSource
		factories["oci_dblm_vulnerability_aggregated_vulnerability_data"] = tf_dblm.DblmVulnerabilityAggregatedVulnerabilityDataDataSource
		factories["oci_dblm_vulnerability_notifications"] = tf_dblm.DblmVulnerabilityNotificationsDataSource
		factories["oci_dblm_vulnerability_resources"] = tf_dblm.DblmVulnerabilityResourcesDataSource
		factories["oci_dblm_vulnerability_scan"] = tf_dblm.DblmVulnerabilityScanDataSource
		factories["oci_dblm_vulnerability_scans"] = tf_dblm.DblmVulnerabilityScansDataSource
		factories["oci_dblm_vulnerability_vulnerabilities"] = tf_dblm.DblmVulnerabilityVulnerabilitiesDataSource
	}
	if common.CheckForEnabledServices("ddfs") {
		factories["oci_ddfs_instance"] = tf_ddfs.DdfsInstanceDataSource
		factories["oci_ddfs_instances"] = tf_ddfs.DdfsInstancesDataSource
	}
	if common.CheckForEnabledServices("delegateaccesscontrol") {
		factories["oci_delegate_access_control_delegated_resource_access_request"] = tf_delegate_access_control.DelegateAccessControlDelegatedResourceAccessRequestDataSource
		factories["oci_delegate_access_control_delegated_resource_access_request_audit_log_report"] = tf_delegate_access_control.DelegateAccessControlDelegatedResourceAccessRequestAuditLogReportDataSource
		factories["oci_delegate_access_control_delegated_resource_access_request_histories"] = tf_delegate_access_control.DelegateAccessControlDelegatedResourceAccessRequestHistoriesDataSource
		factories["oci_delegate_access_control_delegated_resource_access_requests"] = tf_delegate_access_control.DelegateAccessControlDelegatedResourceAccessRequestsDataSource
		factories["oci_delegate_access_control_delegation_control"] = tf_delegate_access_control.DelegateAccessControlDelegationControlDataSource
		factories["oci_delegate_access_control_delegation_control_resources"] = tf_delegate_access_control.DelegateAccessControlDelegationControlResourcesDataSource
		factories["oci_delegate_access_control_delegation_controls"] = tf_delegate_access_control.DelegateAccessControlDelegationControlsDataSource
		factories["oci_delegate_access_control_delegation_subscription"] = tf_delegate_access_control.DelegateAccessControlDelegationSubscriptionDataSource
		factories["oci_delegate_access_control_delegation_subscriptions"] = tf_delegate_access_control.DelegateAccessControlDelegationSubscriptionsDataSource
		factories["oci_delegate_access_control_service_provider"] = tf_delegate_access_control.DelegateAccessControlServiceProviderDataSource
		factories["oci_delegate_access_control_service_provider_action"] = tf_delegate_access_control.DelegateAccessControlServiceProviderActionDataSource
		factories["oci_delegate_access_control_service_provider_actions"] = tf_delegate_access_control.DelegateAccessControlServiceProviderActionsDataSource
		factories["oci_delegate_access_control_service_providers"] = tf_delegate_access_control.DelegateAccessControlServiceProvidersDataSource
	}
	if common.CheckForEnabledServices("demandsignal") {
		factories["oci_demand_signal_occ_demand_signal"] = tf_demand_signal.DemandSignalOccDemandSignalDataSource
		factories["oci_demand_signal_occ_demand_signals"] = tf_demand_signal.DemandSignalOccDemandSignalsDataSource
		factories["oci_demand_signal_occ_metric_alarm"] = tf_demand_signal.DemandSignalOccMetricAlarmDataSource
		factories["oci_demand_signal_occ_metric_alarms"] = tf_demand_signal.DemandSignalOccMetricAlarmsDataSource
	}
	if common.CheckForEnabledServices("desktops") {
		factories["oci_desktops_desktop"] = tf_desktops.DesktopsDesktopDataSource
		factories["oci_desktops_desktop_pool"] = tf_desktops.DesktopsDesktopPoolDataSource
		factories["oci_desktops_desktop_pool_desktops"] = tf_desktops.DesktopsDesktopPoolDesktopsDataSource
		factories["oci_desktops_desktop_pool_volumes"] = tf_desktops.DesktopsDesktopPoolVolumesDataSource
		factories["oci_desktops_desktop_pools"] = tf_desktops.DesktopsDesktopPoolsDataSource
		factories["oci_desktops_desktops"] = tf_desktops.DesktopsDesktopsDataSource
	}
	if common.CheckForEnabledServices("devops") {
		factories["oci_devops_build_pipeline"] = tf_devops.DevopsBuildPipelineDataSource
		factories["oci_devops_build_pipeline_stage"] = tf_devops.DevopsBuildPipelineStageDataSource
		factories["oci_devops_build_pipeline_stages"] = tf_devops.DevopsBuildPipelineStagesDataSource
		factories["oci_devops_build_pipelines"] = tf_devops.DevopsBuildPipelinesDataSource
		factories["oci_devops_build_run"] = tf_devops.DevopsBuildRunDataSource
		factories["oci_devops_build_runs"] = tf_devops.DevopsBuildRunsDataSource
		factories["oci_devops_connection"] = tf_devops.DevopsConnectionDataSource
		factories["oci_devops_connections"] = tf_devops.DevopsConnectionsDataSource
		factories["oci_devops_deploy_artifact"] = tf_devops.DevopsDeployArtifactDataSource
		factories["oci_devops_deploy_artifacts"] = tf_devops.DevopsDeployArtifactsDataSource
		factories["oci_devops_deploy_environment"] = tf_devops.DevopsDeployEnvironmentDataSource
		factories["oci_devops_deploy_environments"] = tf_devops.DevopsDeployEnvironmentsDataSource
		factories["oci_devops_deploy_pipeline"] = tf_devops.DevopsDeployPipelineDataSource
		factories["oci_devops_deploy_pipelines"] = tf_devops.DevopsDeployPipelinesDataSource
		factories["oci_devops_deploy_stage"] = tf_devops.DevopsDeployStageDataSource
		factories["oci_devops_deploy_stages"] = tf_devops.DevopsDeployStagesDataSource
		factories["oci_devops_deployment"] = tf_devops.DevopsDeploymentDataSource
		factories["oci_devops_deployments"] = tf_devops.DevopsDeploymentsDataSource
		factories["oci_devops_project"] = tf_devops.DevopsProjectDataSource
		factories["oci_devops_project_repository_setting"] = tf_devops.DevopsProjectRepositorySettingDataSource
		factories["oci_devops_projects"] = tf_devops.DevopsProjectsDataSource
		factories["oci_devops_repo_file_line"] = tf_devops.DevopsRepoFileLineDataSource
		factories["oci_devops_repositories"] = tf_devops.DevopsRepositoriesDataSource
		factories["oci_devops_repository"] = tf_devops.DevopsRepositoryDataSource
		factories["oci_devops_repository_archive_content"] = tf_devops.DevopsRepositoryArchiveContentDataSource
		factories["oci_devops_repository_author"] = tf_devops.DevopsRepositoryAuthorDataSource
		factories["oci_devops_repository_authors"] = tf_devops.DevopsRepositoryAuthorsDataSource
		factories["oci_devops_repository_commit"] = tf_devops.DevopsRepositoryCommitDataSource
		factories["oci_devops_repository_commits"] = tf_devops.DevopsRepositoryCommitsDataSource
		factories["oci_devops_repository_diff"] = tf_devops.DevopsRepositoryDiffDataSource
		factories["oci_devops_repository_diffs"] = tf_devops.DevopsRepositoryDiffsDataSource
		factories["oci_devops_repository_file_diff"] = tf_devops.DevopsRepositoryFileDiffDataSource
		factories["oci_devops_repository_file_line"] = tf_devops.DevopsRepositoryFileLineDataSource
		factories["oci_devops_repository_mirror_record"] = tf_devops.DevopsRepositoryMirrorRecordDataSource
		factories["oci_devops_repository_mirror_records"] = tf_devops.DevopsRepositoryMirrorRecordsDataSource
		factories["oci_devops_repository_mirrorrecord"] = tf_devops.DevopsRepositoryMirrorrecordDataSource
		factories["oci_devops_repository_object"] = tf_devops.DevopsRepositoryObjectDataSource
		factories["oci_devops_repository_object_content"] = tf_devops.DevopsRepositoryObjectContentDataSource
		factories["oci_devops_repository_path"] = tf_devops.DevopsRepositoryPathDataSource
		factories["oci_devops_repository_paths"] = tf_devops.DevopsRepositoryPathsDataSource
		factories["oci_devops_repository_protected_branches"] = tf_devops.DevopsRepositoryProtectedBranchesDataSource
		factories["oci_devops_repository_ref"] = tf_devops.DevopsRepositoryRefDataSource
		factories["oci_devops_repository_refs"] = tf_devops.DevopsRepositoryRefsDataSource
		factories["oci_devops_repository_setting"] = tf_devops.DevopsRepositorySettingDataSource
		factories["oci_devops_trigger"] = tf_devops.DevopsTriggerDataSource
		factories["oci_devops_triggers"] = tf_devops.DevopsTriggersDataSource
	}
	if common.CheckForEnabledServices("dif") {
		factories["oci_dif_stack"] = tf_dif.DifStackDataSource
		factories["oci_dif_stacks"] = tf_dif.DifStacksDataSource
	}
	if common.CheckForEnabledServices("disasterrecovery") {
		factories["oci_disaster_recovery_automatic_dr_configuration"] = tf_disaster_recovery.DisasterRecoveryAutomaticDrConfigurationDataSource
		factories["oci_disaster_recovery_automatic_dr_configurations"] = tf_disaster_recovery.DisasterRecoveryAutomaticDrConfigurationsDataSource
		factories["oci_disaster_recovery_dr_plan"] = tf_disaster_recovery.DisasterRecoveryDrPlanDataSource
		factories["oci_disaster_recovery_dr_plan_execution"] = tf_disaster_recovery.DisasterRecoveryDrPlanExecutionDataSource
		factories["oci_disaster_recovery_dr_plan_executions"] = tf_disaster_recovery.DisasterRecoveryDrPlanExecutionsDataSource
		factories["oci_disaster_recovery_dr_plans"] = tf_disaster_recovery.DisasterRecoveryDrPlansDataSource
		factories["oci_disaster_recovery_dr_protection_group"] = tf_disaster_recovery.DisasterRecoveryDrProtectionGroupDataSource
		factories["oci_disaster_recovery_dr_protection_groups"] = tf_disaster_recovery.DisasterRecoveryDrProtectionGroupsDataSource
	}
	if common.CheckForEnabledServices("dns") {
		factories["oci_dns_records"] = tf_dns.DnsRecordsDataSource
		factories["oci_dns_resolver"] = tf_dns.DnsResolverDataSource
		factories["oci_dns_resolver_endpoint"] = tf_dns.DnsResolverEndpointDataSource
		factories["oci_dns_resolver_endpoints"] = tf_dns.DnsResolverEndpointsDataSource
		factories["oci_dns_resolvers"] = tf_dns.DnsResolversDataSource
		factories["oci_dns_rrset"] = tf_dns.DnsRrsetDataSource
		factories["oci_dns_rrsets"] = tf_dns.DnsRrsetsDataSource
		factories["oci_dns_steering_policies"] = tf_dns.DnsSteeringPoliciesDataSource
		factories["oci_dns_steering_policy"] = tf_dns.DnsSteeringPolicyDataSource
		factories["oci_dns_steering_policy_attachment"] = tf_dns.DnsSteeringPolicyAttachmentDataSource
		factories["oci_dns_steering_policy_attachments"] = tf_dns.DnsSteeringPolicyAttachmentsDataSource
		factories["oci_dns_tsig_key"] = tf_dns.DnsTsigKeyDataSource
		factories["oci_dns_tsig_keys"] = tf_dns.DnsTsigKeysDataSource
		factories["oci_dns_view"] = tf_dns.DnsViewDataSource
		factories["oci_dns_views"] = tf_dns.DnsViewsDataSource
		factories["oci_dns_zone"] = tf_dns.DnsZoneDataSource
		factories["oci_dns_zones"] = tf_dns.DnsZonesDataSource
	}
	if common.CheckForEnabledServices("email") {
		factories["oci_email_configuration"] = tf_email.EmailConfigurationDataSource
		factories["oci_email_dkim"] = tf_email.EmailDkimDataSource
		factories["oci_email_dkims"] = tf_email.EmailDkimsDataSource
		factories["oci_email_email_domain"] = tf_email.EmailEmailDomainDataSource
		factories["oci_email_email_domains"] = tf_email.EmailEmailDomainsDataSource
		factories["oci_email_email_ip_pool"] = tf_email.EmailEmailIpPoolDataSource
		factories["oci_email_email_ip_pools"] = tf_email.EmailEmailIpPoolsDataSource
		factories["oci_email_email_outbound_ips"] = tf_email.EmailEmailOutboundIpsDataSource
		factories["oci_email_email_return_path"] = tf_email.EmailEmailReturnPathDataSource
		factories["oci_email_email_return_paths"] = tf_email.EmailEmailReturnPathsDataSource
		factories["oci_email_sender"] = tf_email.EmailSenderDataSource
		factories["oci_email_senders"] = tf_email.EmailSendersDataSource
		factories["oci_email_suppression"] = tf_email.EmailSuppressionDataSource
		factories["oci_email_suppressions"] = tf_email.EmailSuppressionsDataSource
	}
	if common.CheckForEnabledServices("events") {
		factories["oci_events_rule"] = tf_events.EventsRuleDataSource
		factories["oci_events_rules"] = tf_events.EventsRulesDataSource
	}
	if common.CheckForEnabledServices("filestorage") {
		factories["oci_file_storage_export_sets"] = tf_file_storage.FileStorageExportSetsDataSource
		factories["oci_file_storage_exports"] = tf_file_storage.FileStorageExportsDataSource
		factories["oci_file_storage_file_system_quota_rule"] = tf_file_storage.FileStorageFileSystemQuotaRuleDataSource
		factories["oci_file_storage_file_system_quota_rules"] = tf_file_storage.FileStorageFileSystemQuotaRulesDataSource
		factories["oci_file_storage_file_systems"] = tf_file_storage.FileStorageFileSystemsDataSource
		factories["oci_file_storage_filesystem_snapshot_policies"] = tf_file_storage.FileStorageFilesystemSnapshotPoliciesDataSource
		factories["oci_file_storage_filesystem_snapshot_policy"] = tf_file_storage.FileStorageFilesystemSnapshotPolicyDataSource
		factories["oci_file_storage_mount_targets"] = tf_file_storage.FileStorageMountTargetsDataSource
		factories["oci_file_storage_outbound_connector"] = tf_file_storage.FileStorageOutboundConnectorDataSource
		factories["oci_file_storage_outbound_connectors"] = tf_file_storage.FileStorageOutboundConnectorsDataSource
		factories["oci_file_storage_replication"] = tf_file_storage.FileStorageReplicationDataSource
		factories["oci_file_storage_replication_target"] = tf_file_storage.FileStorageReplicationTargetDataSource
		factories["oci_file_storage_replication_targets"] = tf_file_storage.FileStorageReplicationTargetsDataSource
		factories["oci_file_storage_replications"] = tf_file_storage.FileStorageReplicationsDataSource
		factories["oci_file_storage_snapshot"] = tf_file_storage.FileStorageSnapshotDataSource
		factories["oci_file_storage_snapshots"] = tf_file_storage.FileStorageSnapshotsDataSource
	}
	if common.CheckForEnabledServices("fleetappsmanagement") {
		factories["oci_fleet_apps_management_announcements"] = tf_fleet_apps_management.FleetAppsManagementAnnouncementsDataSource
		factories["oci_fleet_apps_management_catalog_item"] = tf_fleet_apps_management.FleetAppsManagementCatalogItemDataSource
		factories["oci_fleet_apps_management_catalog_item_variables_definition"] = tf_fleet_apps_management.FleetAppsManagementCatalogItemVariablesDefinitionDataSource
		factories["oci_fleet_apps_management_catalog_items"] = tf_fleet_apps_management.FleetAppsManagementCatalogItemsDataSource
		factories["oci_fleet_apps_management_compliance_policies"] = tf_fleet_apps_management.FleetAppsManagementCompliancePoliciesDataSource
		factories["oci_fleet_apps_management_compliance_policy"] = tf_fleet_apps_management.FleetAppsManagementCompliancePolicyDataSource
		factories["oci_fleet_apps_management_compliance_policy_rule"] = tf_fleet_apps_management.FleetAppsManagementCompliancePolicyRuleDataSource
		factories["oci_fleet_apps_management_compliance_policy_rules"] = tf_fleet_apps_management.FleetAppsManagementCompliancePolicyRulesDataSource
		factories["oci_fleet_apps_management_compliance_record_counts"] = tf_fleet_apps_management.FleetAppsManagementComplianceRecordCountsDataSource
		factories["oci_fleet_apps_management_compliance_records"] = tf_fleet_apps_management.FleetAppsManagementComplianceRecordsDataSource
		factories["oci_fleet_apps_management_fleet"] = tf_fleet_apps_management.FleetAppsManagementFleetDataSource
		factories["oci_fleet_apps_management_fleet_compliance"] = tf_fleet_apps_management.FleetAppsManagementFleetComplianceDataSource
		factories["oci_fleet_apps_management_fleet_compliance_report"] = tf_fleet_apps_management.FleetAppsManagementFleetComplianceReportDataSource
		factories["oci_fleet_apps_management_fleet_credential"] = tf_fleet_apps_management.FleetAppsManagementFleetCredentialDataSource
		factories["oci_fleet_apps_management_fleet_credentials"] = tf_fleet_apps_management.FleetAppsManagementFleetCredentialsDataSource
		factories["oci_fleet_apps_management_fleet_products"] = tf_fleet_apps_management.FleetAppsManagementFleetProductsDataSource
		factories["oci_fleet_apps_management_fleet_properties"] = tf_fleet_apps_management.FleetAppsManagementFleetPropertiesDataSource
		factories["oci_fleet_apps_management_fleet_property"] = tf_fleet_apps_management.FleetAppsManagementFleetPropertyDataSource
		factories["oci_fleet_apps_management_fleet_resource"] = tf_fleet_apps_management.FleetAppsManagementFleetResourceDataSource
		factories["oci_fleet_apps_management_fleet_resources"] = tf_fleet_apps_management.FleetAppsManagementFleetResourcesDataSource
		factories["oci_fleet_apps_management_fleet_targets"] = tf_fleet_apps_management.FleetAppsManagementFleetTargetsDataSource
		factories["oci_fleet_apps_management_fleets"] = tf_fleet_apps_management.FleetAppsManagementFleetsDataSource
		factories["oci_fleet_apps_management_installed_patches"] = tf_fleet_apps_management.FleetAppsManagementInstalledPatchesDataSource
		factories["oci_fleet_apps_management_inventory_records"] = tf_fleet_apps_management.FleetAppsManagementInventoryRecordsDataSource
		factories["oci_fleet_apps_management_inventory_resources"] = tf_fleet_apps_management.FleetAppsManagementInventoryResourcesDataSource
		factories["oci_fleet_apps_management_maintenance_window"] = tf_fleet_apps_management.FleetAppsManagementMaintenanceWindowDataSource
		factories["oci_fleet_apps_management_maintenance_windows"] = tf_fleet_apps_management.FleetAppsManagementMaintenanceWindowsDataSource
		factories["oci_fleet_apps_management_managed_entity_counts"] = tf_fleet_apps_management.FleetAppsManagementManagedEntityCountsDataSource
		factories["oci_fleet_apps_management_onboarding_policies"] = tf_fleet_apps_management.FleetAppsManagementOnboardingPoliciesDataSource
		factories["oci_fleet_apps_management_onboardings"] = tf_fleet_apps_management.FleetAppsManagementOnboardingsDataSource
		factories["oci_fleet_apps_management_patch"] = tf_fleet_apps_management.FleetAppsManagementPatchDataSource
		factories["oci_fleet_apps_management_patches"] = tf_fleet_apps_management.FleetAppsManagementPatchesDataSource
		factories["oci_fleet_apps_management_platform_configuration"] = tf_fleet_apps_management.FleetAppsManagementPlatformConfigurationDataSource
		factories["oci_fleet_apps_management_platform_configurations"] = tf_fleet_apps_management.FleetAppsManagementPlatformConfigurationsDataSource
		factories["oci_fleet_apps_management_properties"] = tf_fleet_apps_management.FleetAppsManagementPropertiesDataSource
		factories["oci_fleet_apps_management_property"] = tf_fleet_apps_management.FleetAppsManagementPropertyDataSource
		factories["oci_fleet_apps_management_provision"] = tf_fleet_apps_management.FleetAppsManagementProvisionDataSource
		factories["oci_fleet_apps_management_provisions"] = tf_fleet_apps_management.FleetAppsManagementProvisionsDataSource
		factories["oci_fleet_apps_management_recommended_patches"] = tf_fleet_apps_management.FleetAppsManagementRecommendedPatchesDataSource
		factories["oci_fleet_apps_management_report_metadata"] = tf_fleet_apps_management.FleetAppsManagementReportMetadataDataSource
		factories["oci_fleet_apps_management_runbook"] = tf_fleet_apps_management.FleetAppsManagementRunbookDataSource
		factories["oci_fleet_apps_management_runbook_export"] = tf_fleet_apps_management.FleetAppsManagementRunbookExportDataSource
		factories["oci_fleet_apps_management_runbook_export_statuses"] = tf_fleet_apps_management.FleetAppsManagementRunbookExportStatusesDataSource
		factories["oci_fleet_apps_management_runbook_import"] = tf_fleet_apps_management.FleetAppsManagementRunbookImportDataSource
		factories["oci_fleet_apps_management_runbook_import_statuses"] = tf_fleet_apps_management.FleetAppsManagementRunbookImportStatusesDataSource
		factories["oci_fleet_apps_management_runbook_version"] = tf_fleet_apps_management.FleetAppsManagementRunbookVersionDataSource
		factories["oci_fleet_apps_management_runbook_versions"] = tf_fleet_apps_management.FleetAppsManagementRunbookVersionsDataSource
		factories["oci_fleet_apps_management_runbooks"] = tf_fleet_apps_management.FleetAppsManagementRunbooksDataSource
		factories["oci_fleet_apps_management_scheduler_definition"] = tf_fleet_apps_management.FleetAppsManagementSchedulerDefinitionDataSource
		factories["oci_fleet_apps_management_scheduler_definition_scheduled_fleets"] = tf_fleet_apps_management.FleetAppsManagementSchedulerDefinitionScheduledFleetsDataSource
		factories["oci_fleet_apps_management_scheduler_definitions"] = tf_fleet_apps_management.FleetAppsManagementSchedulerDefinitionsDataSource
		factories["oci_fleet_apps_management_scheduler_executions"] = tf_fleet_apps_management.FleetAppsManagementSchedulerExecutionsDataSource
		factories["oci_fleet_apps_management_scheduler_job_counts"] = tf_fleet_apps_management.FleetAppsManagementSchedulerJobCountsDataSource
		factories["oci_fleet_apps_management_scheduler_job_job_activity_resources"] = tf_fleet_apps_management.FleetAppsManagementSchedulerJobJobActivityResourcesDataSource
		factories["oci_fleet_apps_management_scheduler_job_job_activity_steps"] = tf_fleet_apps_management.FleetAppsManagementSchedulerJobJobActivityStepsDataSource
		factories["oci_fleet_apps_management_target_components"] = tf_fleet_apps_management.FleetAppsManagementTargetComponentsDataSource
		factories["oci_fleet_apps_management_target_properties"] = tf_fleet_apps_management.FleetAppsManagementTargetPropertiesDataSource
		factories["oci_fleet_apps_management_task_record"] = tf_fleet_apps_management.FleetAppsManagementTaskRecordDataSource
		factories["oci_fleet_apps_management_task_records"] = tf_fleet_apps_management.FleetAppsManagementTaskRecordsDataSource
	}
	if common.CheckForEnabledServices("fleetsoftwareupdate") {
		factories["oci_fleet_software_update_fsu_collection"] = tf_fleet_software_update.FleetSoftwareUpdateFsuCollectionDataSource
		factories["oci_fleet_software_update_fsu_collections"] = tf_fleet_software_update.FleetSoftwareUpdateFsuCollectionsDataSource
		factories["oci_fleet_software_update_fsu_cycle"] = tf_fleet_software_update.FleetSoftwareUpdateFsuCycleDataSource
		factories["oci_fleet_software_update_fsu_cycles"] = tf_fleet_software_update.FleetSoftwareUpdateFsuCyclesDataSource
		factories["oci_fleet_software_update_fsu_readiness_check"] = tf_fleet_software_update.FleetSoftwareUpdateFsuReadinessCheckDataSource
		factories["oci_fleet_software_update_fsu_readiness_checks"] = tf_fleet_software_update.FleetSoftwareUpdateFsuReadinessChecksDataSource
	}
	if common.CheckForEnabledServices("functions") {
		factories["oci_functions_application"] = tf_functions.FunctionsApplicationDataSource
		factories["oci_functions_applications"] = tf_functions.FunctionsApplicationsDataSource
		factories["oci_functions_function"] = tf_functions.FunctionsFunctionDataSource
		factories["oci_functions_functions"] = tf_functions.FunctionsFunctionsDataSource
		factories["oci_functions_functions_runtime"] = tf_functions.FunctionsFunctionsRuntimeDataSource
		factories["oci_functions_functions_runtime_version"] = tf_functions.FunctionsFunctionsRuntimeVersionDataSource
		factories["oci_functions_functions_runtime_versions"] = tf_functions.FunctionsFunctionsRuntimeVersionsDataSource
		factories["oci_functions_functions_runtimes"] = tf_functions.FunctionsFunctionsRuntimesDataSource
		factories["oci_functions_pbf_listing"] = tf_functions.FunctionsPbfListingDataSource
		factories["oci_functions_pbf_listing_triggers"] = tf_functions.FunctionsPbfListingTriggersDataSource
		factories["oci_functions_pbf_listing_version"] = tf_functions.FunctionsPbfListingVersionDataSource
		factories["oci_functions_pbf_listing_versions"] = tf_functions.FunctionsPbfListingVersionsDataSource
		factories["oci_functions_pbf_listings"] = tf_functions.FunctionsPbfListingsDataSource
	}
	if common.CheckForEnabledServices("fusionapps") {
		factories["oci_fusion_apps_fusion_environment"] = tf_fusion_apps.FusionAppsFusionEnvironmentDataSource
		factories["oci_fusion_apps_fusion_environment_admin_user"] = tf_fusion_apps.FusionAppsFusionEnvironmentAdminUserDataSource
		factories["oci_fusion_apps_fusion_environment_admin_users"] = tf_fusion_apps.FusionAppsFusionEnvironmentAdminUsersDataSource
		factories["oci_fusion_apps_fusion_environment_data_masking_activities"] = tf_fusion_apps.FusionAppsFusionEnvironmentDataMaskingActivitiesDataSource
		factories["oci_fusion_apps_fusion_environment_data_masking_activity"] = tf_fusion_apps.FusionAppsFusionEnvironmentDataMaskingActivityDataSource
		factories["oci_fusion_apps_fusion_environment_families"] = tf_fusion_apps.FusionAppsFusionEnvironmentFamiliesDataSource
		factories["oci_fusion_apps_fusion_environment_family"] = tf_fusion_apps.FusionAppsFusionEnvironmentFamilyDataSource
		factories["oci_fusion_apps_fusion_environment_family_limits_and_usage"] = tf_fusion_apps.FusionAppsFusionEnvironmentFamilyLimitsAndUsageDataSource
		factories["oci_fusion_apps_fusion_environment_family_subscription_detail"] = tf_fusion_apps.FusionAppsFusionEnvironmentFamilySubscriptionDetailDataSource
		factories["oci_fusion_apps_fusion_environment_refresh_activities"] = tf_fusion_apps.FusionAppsFusionEnvironmentRefreshActivitiesDataSource
		factories["oci_fusion_apps_fusion_environment_refresh_activity"] = tf_fusion_apps.FusionAppsFusionEnvironmentRefreshActivityDataSource
		factories["oci_fusion_apps_fusion_environment_scheduled_activities"] = tf_fusion_apps.FusionAppsFusionEnvironmentScheduledActivitiesDataSource
		factories["oci_fusion_apps_fusion_environment_scheduled_activity"] = tf_fusion_apps.FusionAppsFusionEnvironmentScheduledActivityDataSource
		factories["oci_fusion_apps_fusion_environment_service_attachment"] = tf_fusion_apps.FusionAppsFusionEnvironmentServiceAttachmentDataSource
		factories["oci_fusion_apps_fusion_environment_service_attachments"] = tf_fusion_apps.FusionAppsFusionEnvironmentServiceAttachmentsDataSource
		factories["oci_fusion_apps_fusion_environment_status"] = tf_fusion_apps.FusionAppsFusionEnvironmentStatusDataSource
		factories["oci_fusion_apps_fusion_environment_time_available_for_refresh"] = tf_fusion_apps.FusionAppsFusionEnvironmentTimeAvailableForRefreshDataSource
		factories["oci_fusion_apps_fusion_environment_time_available_for_refreshs"] = tf_fusion_apps.FusionAppsFusionEnvironmentTimeAvailableForRefreshsDataSource
		factories["oci_fusion_apps_fusion_environments"] = tf_fusion_apps.FusionAppsFusionEnvironmentsDataSource
	}
	if common.CheckForEnabledServices("gdp") {
		factories["oci_gdp_gdp_pipeline"] = tf_gdp.GdpGdpPipelineDataSource
		factories["oci_gdp_gdp_pipelines"] = tf_gdp.GdpGdpPipelinesDataSource
	}
	if common.CheckForEnabledServices("generativeai") {
		factories["oci_generative_ai_dedicated_ai_cluster"] = tf_generative_ai.GenerativeAiDedicatedAiClusterDataSource
		factories["oci_generative_ai_dedicated_ai_clusters"] = tf_generative_ai.GenerativeAiDedicatedAiClustersDataSource
		factories["oci_generative_ai_endpoint"] = tf_generative_ai.GenerativeAiEndpointDataSource
		factories["oci_generative_ai_endpoints"] = tf_generative_ai.GenerativeAiEndpointsDataSource
		factories["oci_generative_ai_generative_ai_private_endpoint"] = tf_generative_ai.GenerativeAiGenerativeAiPrivateEndpointDataSource
		factories["oci_generative_ai_generative_ai_private_endpoints"] = tf_generative_ai.GenerativeAiGenerativeAiPrivateEndpointsDataSource
		factories["oci_generative_ai_hosted_application"] = tf_generative_ai.GenerativeAiHostedApplicationDataSource
		factories["oci_generative_ai_hosted_application_iam"] = tf_generative_ai.GenerativeAiHostedApplicationIamDataSource
		factories["oci_generative_ai_hosted_application_iams"] = tf_generative_ai.GenerativeAiHostedApplicationIamsDataSource
		factories["oci_generative_ai_hosted_application_storage"] = tf_generative_ai.GenerativeAiHostedApplicationStorageDataSource
		factories["oci_generative_ai_hosted_application_storages"] = tf_generative_ai.GenerativeAiHostedApplicationStoragesDataSource
		factories["oci_generative_ai_hosted_applications"] = tf_generative_ai.GenerativeAiHostedApplicationsDataSource
		factories["oci_generative_ai_hosted_deployment"] = tf_generative_ai.GenerativeAiHostedDeploymentDataSource
		factories["oci_generative_ai_hosted_deployments"] = tf_generative_ai.GenerativeAiHostedDeploymentsDataSource
		factories["oci_generative_ai_imported_model"] = tf_generative_ai.GenerativeAiImportedModelDataSource
		factories["oci_generative_ai_imported_models"] = tf_generative_ai.GenerativeAiImportedModelsDataSource
		factories["oci_generative_ai_model"] = tf_generative_ai.GenerativeAiModelDataSource
		factories["oci_generative_ai_model_discoveries"] = tf_generative_ai.GenerativeAiModelDiscoveriesDataSource
		factories["oci_generative_ai_models"] = tf_generative_ai.GenerativeAiModelsDataSource
		factories["oci_generative_ai_project"] = tf_generative_ai.GenerativeAiProjectDataSource
		factories["oci_generative_ai_projects"] = tf_generative_ai.GenerativeAiProjectsDataSource
		factories["oci_generative_ai_routing_profile"] = tf_generative_ai.GenerativeAiRoutingProfileDataSource
		factories["oci_generative_ai_routing_profiles"] = tf_generative_ai.GenerativeAiRoutingProfilesDataSource
		factories["oci_generative_ai_semantic_store"] = tf_generative_ai.GenerativeAiSemanticStoreDataSource
		factories["oci_generative_ai_semantic_stores"] = tf_generative_ai.GenerativeAiSemanticStoresDataSource
	}
	if common.CheckForEnabledServices("generativeaiagent") {
		factories["oci_generative_ai_agent_agent"] = tf_generative_ai_agent.GenerativeAiAgentAgentDataSource
		factories["oci_generative_ai_agent_agent_endpoint"] = tf_generative_ai_agent.GenerativeAiAgentAgentEndpointDataSource
		factories["oci_generative_ai_agent_agent_endpoints"] = tf_generative_ai_agent.GenerativeAiAgentAgentEndpointsDataSource
		factories["oci_generative_ai_agent_agents"] = tf_generative_ai_agent.GenerativeAiAgentAgentsDataSource
		factories["oci_generative_ai_agent_data_ingestion_job"] = tf_generative_ai_agent.GenerativeAiAgentDataIngestionJobDataSource
		factories["oci_generative_ai_agent_data_ingestion_job_log_content"] = tf_generative_ai_agent.GenerativeAiAgentDataIngestionJobLogContentDataSource
		factories["oci_generative_ai_agent_data_ingestion_jobs"] = tf_generative_ai_agent.GenerativeAiAgentDataIngestionJobsDataSource
		factories["oci_generative_ai_agent_data_source"] = tf_generative_ai_agent.GenerativeAiAgentDataSourceDataSource
		factories["oci_generative_ai_agent_data_sources"] = tf_generative_ai_agent.GenerativeAiAgentDataSourcesDataSource
		factories["oci_generative_ai_agent_knowledge_base"] = tf_generative_ai_agent.GenerativeAiAgentKnowledgeBaseDataSource
		factories["oci_generative_ai_agent_knowledge_bases"] = tf_generative_ai_agent.GenerativeAiAgentKnowledgeBasesDataSource
		factories["oci_generative_ai_agent_provisioned_capacities"] = tf_generative_ai_agent.GenerativeAiAgentProvisionedCapacitiesDataSource
		factories["oci_generative_ai_agent_provisioned_capacity"] = tf_generative_ai_agent.GenerativeAiAgentProvisionedCapacityDataSource
		factories["oci_generative_ai_agent_tool"] = tf_generative_ai_agent.GenerativeAiAgentToolDataSource
		factories["oci_generative_ai_agent_tools"] = tf_generative_ai_agent.GenerativeAiAgentToolsDataSource
	}
	if common.CheckForEnabledServices("genericartifactscontent") {
		factories["oci_generic_artifacts_content_artifact_by_path"] = tf_generic_artifacts_content.GenericArtifactsContentArtifactByPathDataSource
		factories["oci_generic_artifacts_content_generic_artifacts_content"] = tf_generic_artifacts_content.GenericArtifactsContentGenericArtifactsContentDataSource
	}
	if common.CheckForEnabledServices("goldengate") {
		factories["oci_golden_gate_ai_models"] = tf_golden_gate.GoldenGateAiModelsDataSource
		factories["oci_golden_gate_ai_providers"] = tf_golden_gate.GoldenGateAiProvidersDataSource
		factories["oci_golden_gate_connection"] = tf_golden_gate.GoldenGateConnectionDataSource
		factories["oci_golden_gate_connection_assignment"] = tf_golden_gate.GoldenGateConnectionAssignmentDataSource
		factories["oci_golden_gate_connection_assignments"] = tf_golden_gate.GoldenGateConnectionAssignmentsDataSource
		factories["oci_golden_gate_connections"] = tf_golden_gate.GoldenGateConnectionsDataSource
		factories["oci_golden_gate_database_registration"] = tf_golden_gate.GoldenGateDatabaseRegistrationDataSource
		factories["oci_golden_gate_database_registrations"] = tf_golden_gate.GoldenGateDatabaseRegistrationsDataSource
		factories["oci_golden_gate_deployment"] = tf_golden_gate.GoldenGateDeploymentDataSource
		factories["oci_golden_gate_deployment_backup"] = tf_golden_gate.GoldenGateDeploymentBackupDataSource
		factories["oci_golden_gate_deployment_backups"] = tf_golden_gate.GoldenGateDeploymentBackupsDataSource
		factories["oci_golden_gate_deployment_certificate"] = tf_golden_gate.GoldenGateDeploymentCertificateDataSource
		factories["oci_golden_gate_deployment_certificates"] = tf_golden_gate.GoldenGateDeploymentCertificatesDataSource
		factories["oci_golden_gate_deployment_disaster_recovery_precheck_report"] = tf_golden_gate.GoldenGateDeploymentDisasterRecoveryPrecheckReportDataSource
		factories["oci_golden_gate_deployment_environments"] = tf_golden_gate.GoldenGateDeploymentEnvironmentsDataSource
		factories["oci_golden_gate_deployment_peers"] = tf_golden_gate.GoldenGateDeploymentPeersDataSource
		factories["oci_golden_gate_deployment_type"] = tf_golden_gate.GoldenGateDeploymentTypeDataSource
		factories["oci_golden_gate_deployment_types"] = tf_golden_gate.GoldenGateDeploymentTypesDataSource
		factories["oci_golden_gate_deployment_upgrade"] = tf_golden_gate.GoldenGateDeploymentUpgradeDataSource
		factories["oci_golden_gate_deployment_upgrades"] = tf_golden_gate.GoldenGateDeploymentUpgradesDataSource
		factories["oci_golden_gate_deployment_versions"] = tf_golden_gate.GoldenGateDeploymentVersionsDataSource
		factories["oci_golden_gate_deployments"] = tf_golden_gate.GoldenGateDeploymentsDataSource
		factories["oci_golden_gate_message"] = tf_golden_gate.GoldenGateMessageDataSource
		factories["oci_golden_gate_messages"] = tf_golden_gate.GoldenGateMessagesDataSource
		factories["oci_golden_gate_pipeline"] = tf_golden_gate.GoldenGatePipelineDataSource
		factories["oci_golden_gate_pipeline_running_processes"] = tf_golden_gate.GoldenGatePipelineRunningProcessesDataSource
		factories["oci_golden_gate_pipeline_schema_tables"] = tf_golden_gate.GoldenGatePipelineSchemaTablesDataSource
		factories["oci_golden_gate_pipeline_schemas"] = tf_golden_gate.GoldenGatePipelineSchemasDataSource
		factories["oci_golden_gate_pipelines"] = tf_golden_gate.GoldenGatePipelinesDataSource
		factories["oci_golden_gate_recipes"] = tf_golden_gate.GoldenGateRecipesDataSource
		factories["oci_golden_gate_trail_file"] = tf_golden_gate.GoldenGateTrailFileDataSource
		factories["oci_golden_gate_trail_files"] = tf_golden_gate.GoldenGateTrailFilesDataSource
		factories["oci_golden_gate_trail_sequence"] = tf_golden_gate.GoldenGateTrailSequenceDataSource
		factories["oci_golden_gate_trail_sequences"] = tf_golden_gate.GoldenGateTrailSequencesDataSource
	}
	if common.CheckForEnabledServices("healthchecks") {
		factories["oci_health_checks_http_monitor"] = tf_health_checks.HealthChecksHttpMonitorDataSource
		factories["oci_health_checks_http_monitors"] = tf_health_checks.HealthChecksHttpMonitorsDataSource
		factories["oci_health_checks_http_probe_results"] = tf_health_checks.HealthChecksHttpProbeResultsDataSource
		factories["oci_health_checks_ping_monitor"] = tf_health_checks.HealthChecksPingMonitorDataSource
		factories["oci_health_checks_ping_monitors"] = tf_health_checks.HealthChecksPingMonitorsDataSource
		factories["oci_health_checks_ping_probe_results"] = tf_health_checks.HealthChecksPingProbeResultsDataSource
		factories["oci_health_checks_vantage_points"] = tf_health_checks.HealthChecksVantagePointsDataSource
	}
	if common.CheckForEnabledServices("identity") {
		factories["oci_identity_allowed_domain_license_types"] = tf_identity.IdentityAllowedDomainLicenseTypesDataSource
		factories["oci_identity_api_keys"] = tf_identity.IdentityApiKeysDataSource
		factories["oci_identity_auth_tokens"] = tf_identity.IdentityAuthTokensDataSource
		factories["oci_identity_authentication_policy"] = tf_identity.IdentityAuthenticationPolicyDataSource
		factories["oci_identity_availability_domain"] = tf_identity.IdentityAvailabilityDomainDataSource
		factories["oci_identity_availability_domains"] = tf_identity.IdentityAvailabilityDomainsDataSource
		factories["oci_identity_compartment"] = tf_identity.IdentityCompartmentDataSource
		factories["oci_identity_compartments"] = tf_identity.IdentityCompartmentsDataSource
		factories["oci_identity_cost_tracking_tags"] = tf_identity.IdentityCostTrackingTagsDataSource
		factories["oci_identity_customer_secret_keys"] = tf_identity.IdentityCustomerSecretKeysDataSource
		factories["oci_identity_db_credentials"] = tf_identity.IdentityDbCredentialsDataSource
		factories["oci_identity_domain"] = tf_identity.IdentityDomainDataSource
		factories["oci_identity_domains"] = tf_identity.IdentityDomainsDataSource
		factories["oci_identity_dynamic_groups"] = tf_identity.IdentityDynamicGroupsDataSource
		factories["oci_identity_fault_domains"] = tf_identity.IdentityFaultDomainsDataSource
		factories["oci_identity_group"] = tf_identity.IdentityGroupDataSource
		factories["oci_identity_groups"] = tf_identity.IdentityGroupsDataSource
		factories["oci_identity_iam_work_request"] = tf_identity.IdentityIamWorkRequestDataSource
		factories["oci_identity_iam_work_request_errors"] = tf_identity.IdentityIamWorkRequestErrorsDataSource
		factories["oci_identity_iam_work_request_logs"] = tf_identity.IdentityIamWorkRequestLogsDataSource
		factories["oci_identity_iam_work_requests"] = tf_identity.IdentityIamWorkRequestsDataSource
		factories["oci_identity_identity_provider_groups"] = tf_identity.IdentityIdentityProviderGroupsDataSource
		factories["oci_identity_identity_providers"] = tf_identity.IdentityIdentityProvidersDataSource
		factories["oci_identity_idp_group_mappings"] = tf_identity.IdentityIdpGroupMappingsDataSource
		factories["oci_identity_network_source"] = tf_identity.IdentityNetworkSourceDataSource
		factories["oci_identity_network_sources"] = tf_identity.IdentityNetworkSourcesDataSource
		factories["oci_identity_policies"] = tf_identity.IdentityPoliciesDataSource
		factories["oci_identity_region_subscriptions"] = tf_identity.IdentityRegionSubscriptionsDataSource
		factories["oci_identity_regions"] = tf_identity.IdentityRegionsDataSource
		factories["oci_identity_smtp_credentials"] = tf_identity.IdentitySmtpCredentialsDataSource
		factories["oci_identity_tag"] = tf_identity.IdentityTagDataSource
		factories["oci_identity_tag_default"] = tf_identity.IdentityTagDefaultDataSource
		factories["oci_identity_tag_defaults"] = tf_identity.IdentityTagDefaultsDataSource
		factories["oci_identity_tag_namespaces"] = tf_identity.IdentityTagNamespacesDataSource
		factories["oci_identity_tag_standard_tag_namespace_template"] = tf_identity.IdentityTagStandardTagNamespaceTemplateDataSource
		factories["oci_identity_tag_standard_tag_namespace_templates"] = tf_identity.IdentityTagStandardTagNamespaceTemplatesDataSource
		factories["oci_identity_tags"] = tf_identity.IdentityTagsDataSource
		factories["oci_identity_tenancy"] = tf_identity.IdentityTenancyDataSource
		factories["oci_identity_ui_password"] = tf_identity.IdentityUiPasswordDataSource
		factories["oci_identity_user"] = tf_identity.IdentityUserDataSource
		factories["oci_identity_user_group_memberships"] = tf_identity.IdentityUserGroupMembershipsDataSource
		factories["oci_identity_users"] = tf_identity.IdentityUsersDataSource
	}
	if common.CheckForEnabledServices("identitydomains") {
		factories["oci_identity_domains_account_mgmt_info"] = tf_identity_domains.IdentityDomainsAccountMgmtInfoDataSource
		factories["oci_identity_domains_account_mgmt_infos"] = tf_identity_domains.IdentityDomainsAccountMgmtInfosDataSource
		factories["oci_identity_domains_account_recovery_setting"] = tf_identity_domains.IdentityDomainsAccountRecoverySettingDataSource
		factories["oci_identity_domains_account_recovery_settings"] = tf_identity_domains.IdentityDomainsAccountRecoverySettingsDataSource
		factories["oci_identity_domains_api_key"] = tf_identity_domains.IdentityDomainsApiKeyDataSource
		factories["oci_identity_domains_api_keys"] = tf_identity_domains.IdentityDomainsApiKeysDataSource
		factories["oci_identity_domains_app"] = tf_identity_domains.IdentityDomainsAppDataSource
		factories["oci_identity_domains_app_role"] = tf_identity_domains.IdentityDomainsAppRoleDataSource
		factories["oci_identity_domains_app_roles"] = tf_identity_domains.IdentityDomainsAppRolesDataSource
		factories["oci_identity_domains_approval_workflow"] = tf_identity_domains.IdentityDomainsApprovalWorkflowDataSource
		factories["oci_identity_domains_approval_workflow_assignment"] = tf_identity_domains.IdentityDomainsApprovalWorkflowAssignmentDataSource
		factories["oci_identity_domains_approval_workflow_assignments"] = tf_identity_domains.IdentityDomainsApprovalWorkflowAssignmentsDataSource
		factories["oci_identity_domains_approval_workflow_step"] = tf_identity_domains.IdentityDomainsApprovalWorkflowStepDataSource
		factories["oci_identity_domains_approval_workflow_steps"] = tf_identity_domains.IdentityDomainsApprovalWorkflowStepsDataSource
		factories["oci_identity_domains_approval_workflows"] = tf_identity_domains.IdentityDomainsApprovalWorkflowsDataSource
		factories["oci_identity_domains_apps"] = tf_identity_domains.IdentityDomainsAppsDataSource
		factories["oci_identity_domains_auth_token"] = tf_identity_domains.IdentityDomainsAuthTokenDataSource
		factories["oci_identity_domains_auth_tokens"] = tf_identity_domains.IdentityDomainsAuthTokensDataSource
		factories["oci_identity_domains_authentication_factor_setting"] = tf_identity_domains.IdentityDomainsAuthenticationFactorSettingDataSource
		factories["oci_identity_domains_authentication_factor_settings"] = tf_identity_domains.IdentityDomainsAuthenticationFactorSettingsDataSource
		factories["oci_identity_domains_branding_setting"] = tf_identity_domains.IdentityDomainsBrandingSettingDataSource
		factories["oci_identity_domains_branding_settings"] = tf_identity_domains.IdentityDomainsBrandingSettingsDataSource
		factories["oci_identity_domains_cloud_gate"] = tf_identity_domains.IdentityDomainsCloudGateDataSource
		factories["oci_identity_domains_cloud_gate_mapping"] = tf_identity_domains.IdentityDomainsCloudGateMappingDataSource
		factories["oci_identity_domains_cloud_gate_mappings"] = tf_identity_domains.IdentityDomainsCloudGateMappingsDataSource
		factories["oci_identity_domains_cloud_gate_server"] = tf_identity_domains.IdentityDomainsCloudGateServerDataSource
		factories["oci_identity_domains_cloud_gate_servers"] = tf_identity_domains.IdentityDomainsCloudGateServersDataSource
		factories["oci_identity_domains_cloud_gates"] = tf_identity_domains.IdentityDomainsCloudGatesDataSource
		factories["oci_identity_domains_condition"] = tf_identity_domains.IdentityDomainsConditionDataSource
		factories["oci_identity_domains_conditions"] = tf_identity_domains.IdentityDomainsConditionsDataSource
		factories["oci_identity_domains_customer_secret_key"] = tf_identity_domains.IdentityDomainsCustomerSecretKeyDataSource
		factories["oci_identity_domains_customer_secret_keys"] = tf_identity_domains.IdentityDomainsCustomerSecretKeysDataSource
		factories["oci_identity_domains_dynamic_resource_group"] = tf_identity_domains.IdentityDomainsDynamicResourceGroupDataSource
		factories["oci_identity_domains_dynamic_resource_groups"] = tf_identity_domains.IdentityDomainsDynamicResourceGroupsDataSource
		factories["oci_identity_domains_grant"] = tf_identity_domains.IdentityDomainsGrantDataSource
		factories["oci_identity_domains_grants"] = tf_identity_domains.IdentityDomainsGrantsDataSource
		factories["oci_identity_domains_group"] = tf_identity_domains.IdentityDomainsGroupDataSource
		factories["oci_identity_domains_groups"] = tf_identity_domains.IdentityDomainsGroupsDataSource
		factories["oci_identity_domains_identity_proofing_provider"] = tf_identity_domains.IdentityDomainsIdentityProofingProviderDataSource
		factories["oci_identity_domains_identity_proofing_provider_template"] = tf_identity_domains.IdentityDomainsIdentityProofingProviderTemplateDataSource
		factories["oci_identity_domains_identity_proofing_provider_templates"] = tf_identity_domains.IdentityDomainsIdentityProofingProviderTemplatesDataSource
		factories["oci_identity_domains_identity_proofing_providers"] = tf_identity_domains.IdentityDomainsIdentityProofingProvidersDataSource
		factories["oci_identity_domains_identity_propagation_trust"] = tf_identity_domains.IdentityDomainsIdentityPropagationTrustDataSource
		factories["oci_identity_domains_identity_propagation_trusts"] = tf_identity_domains.IdentityDomainsIdentityPropagationTrustsDataSource
		factories["oci_identity_domains_identity_provider"] = tf_identity_domains.IdentityDomainsIdentityProviderDataSource
		factories["oci_identity_domains_identity_providers"] = tf_identity_domains.IdentityDomainsIdentityProvidersDataSource
		factories["oci_identity_domains_identity_setting"] = tf_identity_domains.IdentityDomainsIdentitySettingDataSource
		factories["oci_identity_domains_identity_settings"] = tf_identity_domains.IdentityDomainsIdentitySettingsDataSource
		factories["oci_identity_domains_kmsi_setting"] = tf_identity_domains.IdentityDomainsKmsiSettingDataSource
		factories["oci_identity_domains_kmsi_settings"] = tf_identity_domains.IdentityDomainsKmsiSettingsDataSource
		factories["oci_identity_domains_mapped_attribute"] = tf_identity_domains.IdentityDomainsMappedAttributeDataSource
		factories["oci_identity_domains_mapped_attributes"] = tf_identity_domains.IdentityDomainsMappedAttributesDataSource
		factories["oci_identity_domains_my_api_key"] = tf_identity_domains.IdentityDomainsMyApiKeyDataSource
		factories["oci_identity_domains_my_api_keys"] = tf_identity_domains.IdentityDomainsMyApiKeysDataSource
		factories["oci_identity_domains_my_apps"] = tf_identity_domains.IdentityDomainsMyAppsDataSource
		factories["oci_identity_domains_my_auth_token"] = tf_identity_domains.IdentityDomainsMyAuthTokenDataSource
		factories["oci_identity_domains_my_auth_tokens"] = tf_identity_domains.IdentityDomainsMyAuthTokensDataSource
		factories["oci_identity_domains_my_completed_approval"] = tf_identity_domains.IdentityDomainsMyCompletedApprovalDataSource
		factories["oci_identity_domains_my_completed_approvals"] = tf_identity_domains.IdentityDomainsMyCompletedApprovalsDataSource
		factories["oci_identity_domains_my_customer_secret_key"] = tf_identity_domains.IdentityDomainsMyCustomerSecretKeyDataSource
		factories["oci_identity_domains_my_customer_secret_keys"] = tf_identity_domains.IdentityDomainsMyCustomerSecretKeysDataSource
		factories["oci_identity_domains_my_device"] = tf_identity_domains.IdentityDomainsMyDeviceDataSource
		factories["oci_identity_domains_my_devices"] = tf_identity_domains.IdentityDomainsMyDevicesDataSource
		factories["oci_identity_domains_my_groups"] = tf_identity_domains.IdentityDomainsMyGroupsDataSource
		factories["oci_identity_domains_my_oauth2client_credential"] = tf_identity_domains.IdentityDomainsMyOAuth2ClientCredentialDataSource
		factories["oci_identity_domains_my_oauth2client_credentials"] = tf_identity_domains.IdentityDomainsMyOAuth2ClientCredentialsDataSource
		factories["oci_identity_domains_my_pending_approval"] = tf_identity_domains.IdentityDomainsMyPendingApprovalDataSource
		factories["oci_identity_domains_my_pending_approvals"] = tf_identity_domains.IdentityDomainsMyPendingApprovalsDataSource
		factories["oci_identity_domains_my_requestable_groups"] = tf_identity_domains.IdentityDomainsMyRequestableGroupsDataSource
		factories["oci_identity_domains_my_requests"] = tf_identity_domains.IdentityDomainsMyRequestsDataSource
		factories["oci_identity_domains_my_smtp_credential"] = tf_identity_domains.IdentityDomainsMySmtpCredentialDataSource
		factories["oci_identity_domains_my_smtp_credentials"] = tf_identity_domains.IdentityDomainsMySmtpCredentialsDataSource
		factories["oci_identity_domains_my_support_account"] = tf_identity_domains.IdentityDomainsMySupportAccountDataSource
		factories["oci_identity_domains_my_support_accounts"] = tf_identity_domains.IdentityDomainsMySupportAccountsDataSource
		factories["oci_identity_domains_my_trusted_user_agent"] = tf_identity_domains.IdentityDomainsMyTrustedUserAgentDataSource
		factories["oci_identity_domains_my_trusted_user_agents"] = tf_identity_domains.IdentityDomainsMyTrustedUserAgentsDataSource
		factories["oci_identity_domains_my_user_db_credential"] = tf_identity_domains.IdentityDomainsMyUserDbCredentialDataSource
		factories["oci_identity_domains_my_user_db_credentials"] = tf_identity_domains.IdentityDomainsMyUserDbCredentialsDataSource
		factories["oci_identity_domains_network_perimeter"] = tf_identity_domains.IdentityDomainsNetworkPerimeterDataSource
		factories["oci_identity_domains_network_perimeters"] = tf_identity_domains.IdentityDomainsNetworkPerimetersDataSource
		factories["oci_identity_domains_notification_setting"] = tf_identity_domains.IdentityDomainsNotificationSettingDataSource
		factories["oci_identity_domains_notification_settings"] = tf_identity_domains.IdentityDomainsNotificationSettingsDataSource
		factories["oci_identity_domains_oauth2client_credential"] = tf_identity_domains.IdentityDomainsOAuth2ClientCredentialDataSource
		factories["oci_identity_domains_oauth2client_credentials"] = tf_identity_domains.IdentityDomainsOAuth2ClientCredentialsDataSource
		factories["oci_identity_domains_oauth_client_certificate"] = tf_identity_domains.IdentityDomainsOAuthClientCertificateDataSource
		factories["oci_identity_domains_oauth_client_certificates"] = tf_identity_domains.IdentityDomainsOAuthClientCertificatesDataSource
		factories["oci_identity_domains_oauth_partner_certificate"] = tf_identity_domains.IdentityDomainsOAuthPartnerCertificateDataSource
		factories["oci_identity_domains_oauth_partner_certificates"] = tf_identity_domains.IdentityDomainsOAuthPartnerCertificatesDataSource
		factories["oci_identity_domains_oci_console_sign_on_policy_consent"] = tf_identity_domains.IdentityDomainsOciConsoleSignOnPolicyConsentDataSource
		factories["oci_identity_domains_oci_console_sign_on_policy_consents"] = tf_identity_domains.IdentityDomainsOciConsoleSignOnPolicyConsentsDataSource
		factories["oci_identity_domains_password_policies"] = tf_identity_domains.IdentityDomainsPasswordPoliciesDataSource
		factories["oci_identity_domains_password_policy"] = tf_identity_domains.IdentityDomainsPasswordPolicyDataSource
		factories["oci_identity_domains_policies"] = tf_identity_domains.IdentityDomainsPoliciesDataSource
		factories["oci_identity_domains_policy"] = tf_identity_domains.IdentityDomainsPolicyDataSource
		factories["oci_identity_domains_resource_type_schema_attributes"] = tf_identity_domains.IdentityDomainsResourceTypeSchemaAttributesDataSource
		factories["oci_identity_domains_rule"] = tf_identity_domains.IdentityDomainsRuleDataSource
		factories["oci_identity_domains_rules"] = tf_identity_domains.IdentityDomainsRulesDataSource
		factories["oci_identity_domains_security_question"] = tf_identity_domains.IdentityDomainsSecurityQuestionDataSource
		factories["oci_identity_domains_security_question_setting"] = tf_identity_domains.IdentityDomainsSecurityQuestionSettingDataSource
		factories["oci_identity_domains_security_question_settings"] = tf_identity_domains.IdentityDomainsSecurityQuestionSettingsDataSource
		factories["oci_identity_domains_security_questions"] = tf_identity_domains.IdentityDomainsSecurityQuestionsDataSource
		factories["oci_identity_domains_self_registration_profile"] = tf_identity_domains.IdentityDomainsSelfRegistrationProfileDataSource
		factories["oci_identity_domains_self_registration_profiles"] = tf_identity_domains.IdentityDomainsSelfRegistrationProfilesDataSource
		factories["oci_identity_domains_setting"] = tf_identity_domains.IdentityDomainsSettingDataSource
		factories["oci_identity_domains_settings"] = tf_identity_domains.IdentityDomainsSettingsDataSource
		factories["oci_identity_domains_smtp_credential"] = tf_identity_domains.IdentityDomainsSmtpCredentialDataSource
		factories["oci_identity_domains_smtp_credentials"] = tf_identity_domains.IdentityDomainsSmtpCredentialsDataSource
		factories["oci_identity_domains_social_identity_provider"] = tf_identity_domains.IdentityDomainsSocialIdentityProviderDataSource
		factories["oci_identity_domains_social_identity_providers"] = tf_identity_domains.IdentityDomainsSocialIdentityProvidersDataSource
		factories["oci_identity_domains_user"] = tf_identity_domains.IdentityDomainsUserDataSource
		factories["oci_identity_domains_user_attributes_setting"] = tf_identity_domains.IdentityDomainsUserAttributesSettingDataSource
		factories["oci_identity_domains_user_attributes_settings"] = tf_identity_domains.IdentityDomainsUserAttributesSettingsDataSource
		factories["oci_identity_domains_user_db_credential"] = tf_identity_domains.IdentityDomainsUserDbCredentialDataSource
		factories["oci_identity_domains_user_db_credentials"] = tf_identity_domains.IdentityDomainsUserDbCredentialsDataSource
		factories["oci_identity_domains_users"] = tf_identity_domains.IdentityDomainsUsersDataSource
	}
	if common.CheckForEnabledServices("integration") {
		factories["oci_integration_integration_instance"] = tf_integration.IntegrationIntegrationInstanceDataSource
		factories["oci_integration_integration_instances"] = tf_integration.IntegrationIntegrationInstancesDataSource
	}
	if common.CheckForEnabledServices("iot") {
		factories["oci_iot_digital_twin_adapter"] = tf_iot.IotDigitalTwinAdapterDataSource
		factories["oci_iot_digital_twin_adapters"] = tf_iot.IotDigitalTwinAdaptersDataSource
		factories["oci_iot_digital_twin_instance"] = tf_iot.IotDigitalTwinInstanceDataSource
		factories["oci_iot_digital_twin_instance_content"] = tf_iot.IotDigitalTwinInstanceContentDataSource
		factories["oci_iot_digital_twin_instances"] = tf_iot.IotDigitalTwinInstancesDataSource
		factories["oci_iot_digital_twin_model"] = tf_iot.IotDigitalTwinModelDataSource
		factories["oci_iot_digital_twin_model_spec"] = tf_iot.IotDigitalTwinModelSpecDataSource
		factories["oci_iot_digital_twin_models"] = tf_iot.IotDigitalTwinModelsDataSource
		factories["oci_iot_digital_twin_relationship"] = tf_iot.IotDigitalTwinRelationshipDataSource
		factories["oci_iot_digital_twin_relationships"] = tf_iot.IotDigitalTwinRelationshipsDataSource
		factories["oci_iot_iot_domain"] = tf_iot.IotIotDomainDataSource
		factories["oci_iot_iot_domain_group"] = tf_iot.IotIotDomainGroupDataSource
		factories["oci_iot_iot_domain_groups"] = tf_iot.IotIotDomainGroupsDataSource
		factories["oci_iot_iot_domains"] = tf_iot.IotIotDomainsDataSource
		factories["oci_iot_iot_flow_runtime"] = tf_iot.IotIotFlowRuntimeDataSource
		factories["oci_iot_iot_flow_runtime_flow"] = tf_iot.IotIotFlowRuntimeFlowDataSource
		factories["oci_iot_iot_flow_runtimes"] = tf_iot.IotIotFlowRuntimesDataSource
	}
	if common.CheckForEnabledServices("jms") {
		factories["oci_jms_agent_installers"] = tf_jms.JmsAgentInstallersDataSource
		factories["oci_jms_announcements"] = tf_jms.JmsAnnouncementsDataSource
		factories["oci_jms_fleet"] = tf_jms.JmsFleetDataSource
		factories["oci_jms_fleet_advanced_feature_configuration"] = tf_jms.JmsFleetAdvancedFeatureConfigurationDataSource
		factories["oci_jms_fleet_agent_configuration"] = tf_jms.JmsFleetAgentConfigurationDataSource
		factories["oci_jms_fleet_blocklists"] = tf_jms.JmsFleetBlocklistsDataSource
		factories["oci_jms_fleet_containers"] = tf_jms.JmsFleetContainersDataSource
		factories["oci_jms_fleet_crypto_analysis_result"] = tf_jms.JmsFleetCryptoAnalysisResultDataSource
		factories["oci_jms_fleet_crypto_analysis_results"] = tf_jms.JmsFleetCryptoAnalysisResultsDataSource
		factories["oci_jms_fleet_diagnoses"] = tf_jms.JmsFleetDiagnosesDataSource
		factories["oci_jms_fleet_drs_file"] = tf_jms.JmsFleetDrsFileDataSource
		factories["oci_jms_fleet_drs_files"] = tf_jms.JmsFleetDrsFilesDataSource
		factories["oci_jms_fleet_error_analytics"] = tf_jms.JmsFleetErrorAnalyticsDataSource
		factories["oci_jms_fleet_errors"] = tf_jms.JmsFleetErrorsDataSource
		factories["oci_jms_fleet_export_setting"] = tf_jms.JmsFleetExportSettingDataSource
		factories["oci_jms_fleet_export_status"] = tf_jms.JmsFleetExportStatusDataSource
		factories["oci_jms_fleet_installation_site"] = tf_jms.JmsFleetInstallationSiteDataSource
		factories["oci_jms_fleet_installation_sites"] = tf_jms.JmsFleetInstallationSitesDataSource
		factories["oci_jms_fleet_java_migration_analysis_result"] = tf_jms.JmsFleetJavaMigrationAnalysisResultDataSource
		factories["oci_jms_fleet_java_migration_analysis_results"] = tf_jms.JmsFleetJavaMigrationAnalysisResultsDataSource
		factories["oci_jms_fleet_library_applications"] = tf_jms.JmsFleetLibraryApplicationsDataSource
		factories["oci_jms_fleet_library_managed_instances"] = tf_jms.JmsFleetLibraryManagedInstancesDataSource
		factories["oci_jms_fleet_performance_tuning_analysis_result"] = tf_jms.JmsFleetPerformanceTuningAnalysisResultDataSource
		factories["oci_jms_fleet_performance_tuning_analysis_results"] = tf_jms.JmsFleetPerformanceTuningAnalysisResultsDataSource
		factories["oci_jms_fleet_summarize_library_inventory"] = tf_jms.JmsSummarizeLibraryInventoryDataSource
		factories["oci_jms_fleet_uncorrelated_package_applications"] = tf_jms.JmsFleetUncorrelatedPackageApplicationsDataSource
		factories["oci_jms_fleet_uncorrelated_package_managed_instances"] = tf_jms.JmsFleetUncorrelatedPackageManagedInstancesDataSource
		factories["oci_jms_fleet_uncorrelated_packages"] = tf_jms.JmsFleetUncorrelatedPackagesDataSource
		factories["oci_jms_fleets"] = tf_jms.JmsFleetsDataSource
		factories["oci_jms_java_families"] = tf_jms.JmsJavaFamiliesDataSource
		factories["oci_jms_java_family"] = tf_jms.JmsJavaFamilyDataSource
		factories["oci_jms_java_release"] = tf_jms.JmsJavaReleaseDataSource
		factories["oci_jms_java_releases"] = tf_jms.JmsJavaReleasesDataSource
		factories["oci_jms_jms_plugin"] = tf_jms.JmsJmsPluginDataSource
		factories["oci_jms_jms_plugins"] = tf_jms.JmsJmsPluginsDataSource
		factories["oci_jms_list_jre_usage"] = tf_jms.JmsListJreUsageDataSource
		factories["oci_jms_plugin_error_analytics"] = tf_jms.JmsPluginErrorAnalyticsDataSource
		factories["oci_jms_plugin_errors"] = tf_jms.JmsPluginErrorsDataSource
		factories["oci_jms_summarize_resource_inventory"] = tf_jms.JmsSummarizeResourceInventoryDataSource
		factories["oci_jms_task_schedule"] = tf_jms.JmsTaskScheduleDataSource
		factories["oci_jms_task_schedules"] = tf_jms.JmsTaskSchedulesDataSource
	}
	if common.CheckForEnabledServices("jmsjavadownloads") {
		factories["oci_jms_java_downloads_java_download_records"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadRecordsDataSource
		factories["oci_jms_java_downloads_java_download_report"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadReportDataSource
		factories["oci_jms_java_downloads_java_download_report_content"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadReportContentDataSource
		factories["oci_jms_java_downloads_java_download_reports"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadReportsDataSource
		factories["oci_jms_java_downloads_java_download_token"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadTokenDataSource
		factories["oci_jms_java_downloads_java_download_tokens"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadTokensDataSource
		factories["oci_jms_java_downloads_java_license"] = tf_jms_java_downloads.JmsJavaDownloadsJavaLicenseDataSource
		factories["oci_jms_java_downloads_java_license_acceptance_record"] = tf_jms_java_downloads.JmsJavaDownloadsJavaLicenseAcceptanceRecordDataSource
		factories["oci_jms_java_downloads_java_license_acceptance_records"] = tf_jms_java_downloads.JmsJavaDownloadsJavaLicenseAcceptanceRecordsDataSource
		factories["oci_jms_java_downloads_java_licenses"] = tf_jms_java_downloads.JmsJavaDownloadsJavaLicensesDataSource
	}
	if common.CheckForEnabledServices("jmsutils") {
		factories["oci_jms_utils_analyze_applications_configuration"] = tf_jms_utils.JmsUtilsAnalyzeApplicationsConfigurationDataSource
		factories["oci_jms_utils_java_migration_analysi"] = tf_jms_utils.JmsUtilsJavaMigrationAnalysiDataSource
		factories["oci_jms_utils_java_migration_analysis"] = tf_jms_utils.JmsUtilsJavaMigrationAnalysisDataSource
		factories["oci_jms_utils_performance_tuning_analysi"] = tf_jms_utils.JmsUtilsPerformanceTuningAnalysiDataSource
		factories["oci_jms_utils_performance_tuning_analysis"] = tf_jms_utils.JmsUtilsPerformanceTuningAnalysisDataSource
		factories["oci_jms_utils_subscription_acknowledgment_configuration"] = tf_jms_utils.JmsUtilsSubscriptionAcknowledgmentConfigurationDataSource
	}
	if common.CheckForEnabledServices("kms") {
		factories["oci_kms_decrypted_data"] = tf_kms.KmsDecryptedDataDataSource
		factories["oci_kms_ekms_private_endpoint"] = tf_kms.KmsEkmsPrivateEndpointDataSource
		factories["oci_kms_ekms_private_endpoints"] = tf_kms.KmsEkmsPrivateEndpointsDataSource
		factories["oci_kms_encrypted_data"] = tf_kms.KmsEncryptedDataDataSource
		factories["oci_kms_key"] = tf_kms.KmsKeyDataSource
		factories["oci_kms_key_version"] = tf_kms.KmsKeyVersionDataSource
		factories["oci_kms_key_versions"] = tf_kms.KmsKeyVersionsDataSource
		factories["oci_kms_keys"] = tf_kms.KmsKeysDataSource
		factories["oci_kms_replication_status"] = tf_kms.KmsReplicationStatusDataSource
		factories["oci_kms_vault"] = tf_kms.KmsVaultDataSource
		factories["oci_kms_vault_replicas"] = tf_kms.KmsVaultReplicasDataSource
		factories["oci_kms_vault_usage"] = tf_kms.KmsVaultUsageDataSource
		factories["oci_kms_vaults"] = tf_kms.KmsVaultsDataSource
	}
	if common.CheckForEnabledServices("licensemanager") {
		factories["oci_license_manager_configuration"] = tf_license_manager.LicenseManagerConfigurationDataSource
		factories["oci_license_manager_license_metric"] = tf_license_manager.LicenseManagerLicenseMetricDataSource
		factories["oci_license_manager_license_record"] = tf_license_manager.LicenseManagerLicenseRecordDataSource
		factories["oci_license_manager_license_records"] = tf_license_manager.LicenseManagerLicenseRecordsDataSource
		factories["oci_license_manager_product_license"] = tf_license_manager.LicenseManagerProductLicenseDataSource
		factories["oci_license_manager_product_license_consumers"] = tf_license_manager.LicenseManagerProductLicenseConsumersDataSource
		factories["oci_license_manager_product_licenses"] = tf_license_manager.LicenseManagerProductLicensesDataSource
		factories["oci_license_manager_top_utilized_product_licenses"] = tf_license_manager.LicenseManagerTopUtilizedProductLicensesDataSource
		factories["oci_license_manager_top_utilized_resources"] = tf_license_manager.LicenseManagerTopUtilizedResourcesDataSource
	}
	if common.CheckForEnabledServices("limits") {
		factories["oci_limits_limit_definitions"] = tf_limits.LimitsLimitDefinitionsDataSource
		factories["oci_limits_limit_values"] = tf_limits.LimitsLimitValuesDataSource
		factories["oci_limits_quota"] = tf_limits.LimitsQuotaDataSource
		factories["oci_limits_quotas"] = tf_limits.LimitsQuotasDataSource
		factories["oci_limits_resource_availability"] = tf_limits.LimitsResourceAvailabilityDataSource
		factories["oci_limits_services"] = tf_limits.LimitsServicesDataSource
	}
	if common.CheckForEnabledServices("loadbalancer") {
		factories["oci_load_balancer_backend_health"] = tf_load_balancer.LoadBalancerBackendHealthDataSource
		factories["oci_load_balancer_backend_set_health"] = tf_load_balancer.LoadBalancerBackendSetHealthDataSource
		factories["oci_load_balancer_backend_sets"] = tf_load_balancer.LoadBalancerBackendSetsDataSource
		factories["oci_load_balancer_backends"] = tf_load_balancer.LoadBalancerBackendsDataSource
		factories["oci_load_balancer_certificates"] = tf_load_balancer.LoadBalancerCertificatesDataSource
		factories["oci_load_balancer_health"] = tf_load_balancer.LoadBalancerLoadBalancerHealthDataSource
		factories["oci_load_balancer_hostnames"] = tf_load_balancer.LoadBalancerHostnamesDataSource
		factories["oci_load_balancer_listener_rules"] = tf_load_balancer.LoadBalancerListenerRulesDataSource
		factories["oci_load_balancer_load_balancer_routing_policies"] = tf_load_balancer.LoadBalancerLoadBalancerRoutingPoliciesDataSource
		factories["oci_load_balancer_load_balancer_routing_policy"] = tf_load_balancer.LoadBalancerLoadBalancerRoutingPolicyDataSource
		factories["oci_load_balancer_load_balancers"] = tf_load_balancer.LoadBalancerLoadBalancersDataSource
		factories["oci_load_balancer_path_route_sets"] = tf_load_balancer.LoadBalancerPathRouteSetsDataSource
		factories["oci_load_balancer_policies"] = tf_load_balancer.LoadBalancerLoadBalancerPoliciesDataSource
		factories["oci_load_balancer_protocols"] = tf_load_balancer.LoadBalancerLoadBalancerProtocolsDataSource
		factories["oci_load_balancer_rule_set"] = tf_load_balancer.LoadBalancerRuleSetDataSource
		factories["oci_load_balancer_rule_sets"] = tf_load_balancer.LoadBalancerRuleSetsDataSource
		factories["oci_load_balancer_shapes"] = tf_load_balancer.LoadBalancerLoadBalancerShapesDataSource
		factories["oci_load_balancer_ssl_cipher_suite"] = tf_load_balancer.LoadBalancerSslCipherSuiteDataSource
		factories["oci_load_balancer_ssl_cipher_suites"] = tf_load_balancer.LoadBalancerSslCipherSuitesDataSource
	}
	if common.CheckForEnabledServices("loganalytics") {
		factories["oci_log_analytics_log_analytics_categories_list"] = tf_log_analytics.LogAnalyticsLogAnalyticsCategoriesListDataSource
		factories["oci_log_analytics_log_analytics_category"] = tf_log_analytics.LogAnalyticsLogAnalyticsCategoryDataSource
		factories["oci_log_analytics_log_analytics_entities"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntitiesDataSource
		factories["oci_log_analytics_log_analytics_entities_summary"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntitiesSummaryDataSource
		factories["oci_log_analytics_log_analytics_entity"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityDataSource
		factories["oci_log_analytics_log_analytics_entity_associations_list"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityAssociationsListDataSource
		factories["oci_log_analytics_log_analytics_entity_topology"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityTopologyDataSource
		factories["oci_log_analytics_log_analytics_entity_type"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityTypeDataSource
		factories["oci_log_analytics_log_analytics_entity_types"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityTypesDataSource
		factories["oci_log_analytics_log_analytics_log_group"] = tf_log_analytics.LogAnalyticsLogAnalyticsLogGroupDataSource
		factories["oci_log_analytics_log_analytics_log_groups"] = tf_log_analytics.LogAnalyticsLogAnalyticsLogGroupsDataSource
		factories["oci_log_analytics_log_analytics_log_groups_summary"] = tf_log_analytics.LogAnalyticsLogAnalyticsLogGroupsSummaryDataSource
		factories["oci_log_analytics_log_analytics_object_collection_rule"] = tf_log_analytics.LogAnalyticsLogAnalyticsObjectCollectionRuleDataSource
		factories["oci_log_analytics_log_analytics_object_collection_rules"] = tf_log_analytics.LogAnalyticsLogAnalyticsObjectCollectionRulesDataSource
		factories["oci_log_analytics_log_analytics_preference"] = tf_log_analytics.LogAnalyticsLogAnalyticsPreferenceDataSource
		factories["oci_log_analytics_log_analytics_resource_categories_list"] = tf_log_analytics.LogAnalyticsLogAnalyticsResourceCategoriesListDataSource
		factories["oci_log_analytics_log_analytics_unprocessed_data_bucket"] = tf_log_analytics.LogAnalyticsLogAnalyticsUnprocessedDataBucketDataSource
		factories["oci_log_analytics_log_sets_count"] = tf_log_analytics.LogAnalyticsLogSetsCountDataSource
		factories["oci_log_analytics_namespace"] = tf_log_analytics.LogAnalyticsNamespaceDataSource
		factories["oci_log_analytics_namespace_effective_properties"] = tf_log_analytics.LogAnalyticsNamespaceEffectivePropertiesDataSource
		factories["oci_log_analytics_namespace_field_usage"] = tf_log_analytics.LogAnalyticsNamespaceFieldUsageDataSource
		factories["oci_log_analytics_namespace_ingest_time_rule"] = tf_log_analytics.LogAnalyticsNamespaceIngestTimeRuleDataSource
		factories["oci_log_analytics_namespace_ingest_time_rules"] = tf_log_analytics.LogAnalyticsNamespaceIngestTimeRulesDataSource
		factories["oci_log_analytics_namespace_lookup"] = tf_log_analytics.LogAnalyticsNamespaceLookupDataSource
		factories["oci_log_analytics_namespace_parser_actions"] = tf_log_analytics.LogAnalyticsNamespaceParserActionsDataSource
		factories["oci_log_analytics_namespace_properties_metadata"] = tf_log_analytics.LogAnalyticsNamespacePropertiesMetadataDataSource
		factories["oci_log_analytics_namespace_rules"] = tf_log_analytics.LogAnalyticsNamespaceRulesDataSource
		factories["oci_log_analytics_namespace_rules_summary"] = tf_log_analytics.LogAnalyticsNamespaceRulesSummaryDataSource
		factories["oci_log_analytics_namespace_scheduled_task"] = tf_log_analytics.LogAnalyticsNamespaceScheduledTaskDataSource
		factories["oci_log_analytics_namespace_scheduled_tasks"] = tf_log_analytics.LogAnalyticsNamespaceScheduledTasksDataSource
		factories["oci_log_analytics_namespace_storage_archival_config"] = tf_log_analytics.LogAnalyticsNamespaceStorageArchivalConfigDataSource
		factories["oci_log_analytics_namespace_storage_encryption_key_info"] = tf_log_analytics.LogAnalyticsNamespaceStorageEncryptionKeyInfoDataSource
		factories["oci_log_analytics_namespace_storage_overlapping_recalls"] = tf_log_analytics.LogAnalyticsNamespaceStorageOverlappingRecallsDataSource
		factories["oci_log_analytics_namespace_storage_recall_count"] = tf_log_analytics.LogAnalyticsNamespaceStorageRecallCountDataSource
		factories["oci_log_analytics_namespace_storage_recalled_data_size"] = tf_log_analytics.LogAnalyticsNamespaceStorageRecalledDataSizeDataSource
		factories["oci_log_analytics_namespace_template"] = tf_log_analytics.LogAnalyticsNamespaceTemplateDataSource
		factories["oci_log_analytics_namespace_templates"] = tf_log_analytics.LogAnalyticsNamespaceTemplatesDataSource
		factories["oci_log_analytics_namespaces"] = tf_log_analytics.LogAnalyticsNamespacesDataSource
	}
	if common.CheckForEnabledServices("logging") {
		factories["oci_logging_log"] = tf_logging.LoggingLogDataSource
		factories["oci_logging_log_group"] = tf_logging.LoggingLogGroupDataSource
		factories["oci_logging_log_groups"] = tf_logging.LoggingLogGroupsDataSource
		factories["oci_logging_log_saved_search"] = tf_logging.LoggingLogSavedSearchDataSource
		factories["oci_logging_log_saved_searches"] = tf_logging.LoggingLogSavedSearchesDataSource
		factories["oci_logging_logs"] = tf_logging.LoggingLogsDataSource
		factories["oci_logging_unified_agent_configuration"] = tf_logging.LoggingUnifiedAgentConfigurationDataSource
		factories["oci_logging_unified_agent_configurations"] = tf_logging.LoggingUnifiedAgentConfigurationsDataSource
	}
	if common.CheckForEnabledServices("lustrefilestorage") {
		factories["oci_lustre_file_storage_available_maintenance_schedule_start_times"] = tf_lustre_file_storage.LustreFileStorageAvailableMaintenanceScheduleStartTimesDataSource
		factories["oci_lustre_file_storage_available_override_maintenance_start_times"] = tf_lustre_file_storage.LustreFileStorageAvailableOverrideMaintenanceStartTimesDataSource
		factories["oci_lustre_file_storage_lustre_file_system"] = tf_lustre_file_storage.LustreFileStorageLustreFileSystemDataSource
		factories["oci_lustre_file_storage_lustre_file_systems"] = tf_lustre_file_storage.LustreFileStorageLustreFileSystemsDataSource
		factories["oci_lustre_file_storage_object_storage_link"] = tf_lustre_file_storage.LustreFileStorageObjectStorageLinkDataSource
		factories["oci_lustre_file_storage_object_storage_link_sync_job"] = tf_lustre_file_storage.LustreFileStorageObjectStorageLinkSyncJobDataSource
		factories["oci_lustre_file_storage_object_storage_link_sync_jobs"] = tf_lustre_file_storage.LustreFileStorageObjectStorageLinkSyncJobsDataSource
		factories["oci_lustre_file_storage_object_storage_links"] = tf_lustre_file_storage.LustreFileStorageObjectStorageLinksDataSource
	}
	if common.CheckForEnabledServices("managedkafka") {
		factories["oci_managed_kafka_addon_options"] = tf_managed_kafka.ManagedKafkaAddonOptionsDataSource
		factories["oci_managed_kafka_kafka_cluster"] = tf_managed_kafka.ManagedKafkaKafkaClusterDataSource
		factories["oci_managed_kafka_kafka_cluster_addon"] = tf_managed_kafka.ManagedKafkaKafkaClusterAddonDataSource
		factories["oci_managed_kafka_kafka_cluster_addons"] = tf_managed_kafka.ManagedKafkaKafkaClusterAddonsDataSource
		factories["oci_managed_kafka_kafka_cluster_config"] = tf_managed_kafka.ManagedKafkaKafkaClusterConfigDataSource
		factories["oci_managed_kafka_kafka_cluster_config_version"] = tf_managed_kafka.ManagedKafkaKafkaClusterConfigVersionDataSource
		factories["oci_managed_kafka_kafka_cluster_config_versions"] = tf_managed_kafka.ManagedKafkaKafkaClusterConfigVersionsDataSource
		factories["oci_managed_kafka_kafka_cluster_configs"] = tf_managed_kafka.ManagedKafkaKafkaClusterConfigsDataSource
		factories["oci_managed_kafka_kafka_clusters"] = tf_managed_kafka.ManagedKafkaKafkaClustersDataSource
		factories["oci_managed_kafka_node_shapes"] = tf_managed_kafka.ManagedKafkaNodeShapesDataSource
	}
	if common.CheckForEnabledServices("managementagent") {
		factories["oci_management_agent_management_agent"] = tf_management_agent.ManagementAgentManagementAgentDataSource
		factories["oci_management_agent_management_agent_available_histories"] = tf_management_agent.ManagementAgentManagementAgentAvailableHistoriesDataSource
		factories["oci_management_agent_management_agent_count"] = tf_management_agent.ManagementAgentManagementAgentCountDataSource
		factories["oci_management_agent_management_agent_data_source"] = tf_management_agent.ManagementAgentManagementAgentDataSourceDataSource
		factories["oci_management_agent_management_agent_data_sources"] = tf_management_agent.ManagementAgentManagementAgentDataSourcesDataSource
		factories["oci_management_agent_management_agent_get_auto_upgradable_config"] = tf_management_agent.ManagementAgentManagementAgentGetAutoUpgradableConfigDataSource
		factories["oci_management_agent_management_agent_images"] = tf_management_agent.ManagementAgentManagementAgentImagesDataSource
		factories["oci_management_agent_management_agent_install_key"] = tf_management_agent.ManagementAgentManagementAgentInstallKeyDataSource
		factories["oci_management_agent_management_agent_install_keys"] = tf_management_agent.ManagementAgentManagementAgentInstallKeysDataSource
		factories["oci_management_agent_management_agent_named_credentials_metadata"] = tf_management_agent.ManagementAgentManagementAgentNamedCredentialsMetadataDataSource
		factories["oci_management_agent_management_agent_plugin_count"] = tf_management_agent.ManagementAgentManagementAgentPluginCountDataSource
		factories["oci_management_agent_management_agent_plugins"] = tf_management_agent.ManagementAgentManagementAgentPluginsDataSource
		factories["oci_management_agent_management_agents"] = tf_management_agent.ManagementAgentManagementAgentsDataSource
		factories["oci_management_agent_named_credential"] = tf_management_agent.ManagementAgentNamedCredentialDataSource
		factories["oci_management_agent_named_credentials"] = tf_management_agent.ManagementAgentNamedCredentialsDataSource
	}
	if common.CheckForEnabledServices("managementdashboard") {
		factories["oci_management_dashboard_management_dashboards_export"] = tf_management_dashboard.ManagementDashboardManagementDashboardsExportDataSource
		factories["oci_management_dashboard_management_saved_search"] = tf_management_dashboard.ManagementDashboardManagementSavedSearchDataSource
		factories["oci_management_dashboard_management_saved_searches"] = tf_management_dashboard.ManagementDashboardManagementSavedSearchesDataSource
	}
	if common.CheckForEnabledServices("marketplace") {
		factories["oci_marketplace_accepted_agreement"] = tf_marketplace.MarketplaceAcceptedAgreementDataSource
		factories["oci_marketplace_accepted_agreements"] = tf_marketplace.MarketplaceAcceptedAgreementsDataSource
		factories["oci_marketplace_categories"] = tf_marketplace.MarketplaceCategoriesDataSource
		factories["oci_marketplace_listing"] = tf_marketplace.MarketplaceListingDataSource
		factories["oci_marketplace_listing_package"] = tf_marketplace.MarketplaceListingPackageDataSource
		factories["oci_marketplace_listing_package_agreements"] = tf_marketplace.MarketplaceListingPackageAgreementsDataSource
		factories["oci_marketplace_listing_packages"] = tf_marketplace.MarketplaceListingPackagesDataSource
		factories["oci_marketplace_listing_taxes"] = tf_marketplace.MarketplaceListingTaxesDataSource
		factories["oci_marketplace_listings"] = tf_marketplace.MarketplaceListingsDataSource
		factories["oci_marketplace_marketplace_metadata_public_keys"] = tf_marketplace.MarketplaceMarketplaceMetadataPublicKeysDataSource
		factories["oci_marketplace_publication"] = tf_marketplace.MarketplacePublicationDataSource
		factories["oci_marketplace_publication_package"] = tf_marketplace.MarketplacePublicationPackageDataSource
		factories["oci_marketplace_publication_packages"] = tf_marketplace.MarketplacePublicationPackagesDataSource
		factories["oci_marketplace_publications"] = tf_marketplace.MarketplacePublicationsDataSource
		factories["oci_marketplace_publishers"] = tf_marketplace.MarketplacePublishersDataSource
	}
	if common.CheckForEnabledServices("mediaservices") {
		factories["oci_media_services_media_asset"] = tf_media_services.MediaServicesMediaAssetDataSource
		factories["oci_media_services_media_asset_distribution_channel_attachment"] = tf_media_services.MediaServicesMediaAssetDistributionChannelAttachmentDataSource
		factories["oci_media_services_media_assets"] = tf_media_services.MediaServicesMediaAssetsDataSource
		factories["oci_media_services_media_workflow"] = tf_media_services.MediaServicesMediaWorkflowDataSource
		factories["oci_media_services_media_workflow_configuration"] = tf_media_services.MediaServicesMediaWorkflowConfigurationDataSource
		factories["oci_media_services_media_workflow_configurations"] = tf_media_services.MediaServicesMediaWorkflowConfigurationsDataSource
		factories["oci_media_services_media_workflow_job"] = tf_media_services.MediaServicesMediaWorkflowJobDataSource
		factories["oci_media_services_media_workflow_job_fact"] = tf_media_services.MediaServicesMediaWorkflowJobFactDataSource
		factories["oci_media_services_media_workflow_job_facts"] = tf_media_services.MediaServicesMediaWorkflowJobFactsDataSource
		factories["oci_media_services_media_workflow_jobs"] = tf_media_services.MediaServicesMediaWorkflowJobsDataSource
		factories["oci_media_services_media_workflow_task_declaration"] = tf_media_services.MediaServicesMediaWorkflowTaskDeclarationDataSource
		factories["oci_media_services_media_workflows"] = tf_media_services.MediaServicesMediaWorkflowsDataSource
		factories["oci_media_services_stream_cdn_config"] = tf_media_services.MediaServicesStreamCdnConfigDataSource
		factories["oci_media_services_stream_cdn_configs"] = tf_media_services.MediaServicesStreamCdnConfigsDataSource
		factories["oci_media_services_stream_distribution_channel"] = tf_media_services.MediaServicesStreamDistributionChannelDataSource
		factories["oci_media_services_stream_distribution_channels"] = tf_media_services.MediaServicesStreamDistributionChannelsDataSource
		factories["oci_media_services_stream_packaging_config"] = tf_media_services.MediaServicesStreamPackagingConfigDataSource
		factories["oci_media_services_stream_packaging_configs"] = tf_media_services.MediaServicesStreamPackagingConfigsDataSource
		factories["oci_media_services_system_media_workflow"] = tf_media_services.MediaServicesSystemMediaWorkflowDataSource
	}
	if common.CheckForEnabledServices("meteringcomputation") {
		factories["oci_metering_computation_average_carbon_emission"] = tf_metering_computation.MeteringComputationAverageCarbonEmissionDataSource
		factories["oci_metering_computation_clean_energy_usage"] = tf_metering_computation.MeteringComputationCleanEnergyUsageDataSource
		factories["oci_metering_computation_configuration"] = tf_metering_computation.MeteringComputationConfigurationDataSource
		factories["oci_metering_computation_custom_table"] = tf_metering_computation.MeteringComputationCustomTableDataSource
		factories["oci_metering_computation_custom_tables"] = tf_metering_computation.MeteringComputationCustomTablesDataSource
		factories["oci_metering_computation_queries"] = tf_metering_computation.MeteringComputationQueriesDataSource
		factories["oci_metering_computation_query"] = tf_metering_computation.MeteringComputationQueryDataSource
		factories["oci_metering_computation_schedule"] = tf_metering_computation.MeteringComputationScheduleDataSource
		factories["oci_metering_computation_scheduled_run"] = tf_metering_computation.MeteringComputationScheduledRunDataSource
		factories["oci_metering_computation_scheduled_runs"] = tf_metering_computation.MeteringComputationScheduledRunsDataSource
		factories["oci_metering_computation_schedules"] = tf_metering_computation.MeteringComputationSchedulesDataSource
		factories["oci_metering_computation_usage_carbon_emissions_config"] = tf_metering_computation.MeteringComputationUsageCarbonEmissionsConfigDataSource
		factories["oci_metering_computation_usage_carbon_emissions_queries"] = tf_metering_computation.MeteringComputationUsageCarbonEmissionsQueriesDataSource
		factories["oci_metering_computation_usage_carbon_emissions_query"] = tf_metering_computation.MeteringComputationUsageCarbonEmissionsQueryDataSource
		factories["oci_metering_computation_usage_statement_email_recipients_group"] = tf_metering_computation.MeteringComputationUsageStatementEmailRecipientsGroupDataSource
		factories["oci_metering_computation_usage_statement_email_recipients_groups"] = tf_metering_computation.MeteringComputationUsageStatementEmailRecipientsGroupsDataSource
	}
	if common.CheckForEnabledServices("monitoring") {
		factories["oci_monitoring_alarm"] = tf_monitoring.MonitoringAlarmDataSource
		factories["oci_monitoring_alarm_history_collection"] = tf_monitoring.MonitoringAlarmHistoryCollectionDataSource
		factories["oci_monitoring_alarm_statuses"] = tf_monitoring.MonitoringAlarmStatusesDataSource
		factories["oci_monitoring_alarm_suppression"] = tf_monitoring.MonitoringAlarmSuppressionDataSource
		factories["oci_monitoring_alarm_suppressions"] = tf_monitoring.MonitoringAlarmSuppressionsDataSource
		factories["oci_monitoring_alarms"] = tf_monitoring.MonitoringAlarmsDataSource
		factories["oci_monitoring_metric_data"] = tf_monitoring.MonitoringMetricDataDataSource
		factories["oci_monitoring_metrics"] = tf_monitoring.MonitoringMetricsDataSource
	}
	if common.CheckForEnabledServices("multicloud") {
		factories["oci_multicloud_external_location_mapping_metadata"] = tf_multicloud.MulticloudExternalLocationMappingMetadataDataSource
		factories["oci_multicloud_external_location_summaries_metadata"] = tf_multicloud.MulticloudExternalLocationSummariesMetadataDataSource
		factories["oci_multicloud_external_locations_metadata"] = tf_multicloud.MulticloudExternalLocationsMetadataDataSource
		factories["oci_multicloud_multicloudalerts"] = tf_multicloud.MulticloudMulticloudalertsDataSource
		factories["oci_multicloud_multicloudpolicies"] = tf_multicloud.MulticloudMulticloudpoliciesDataSource
		factories["oci_multicloud_multicloudsubscriptions"] = tf_multicloud.MulticloudMulticloudsubscriptionsDataSource
		factories["oci_multicloud_network_anchor"] = tf_multicloud.MulticloudNetworkAnchorDataSource
		factories["oci_multicloud_network_anchors"] = tf_multicloud.MulticloudNetworkAnchorsDataSource
		factories["oci_multicloud_om_hub_multi_cloud_metadata"] = tf_multicloud.MulticloudOmHubMultiCloudMetadataDataSource
		factories["oci_multicloud_om_hub_multi_clouds_metadata"] = tf_multicloud.MulticloudOmHubMultiCloudsMetadataDataSource
		factories["oci_multicloud_om_hub_multicloud_resources"] = tf_multicloud.MulticloudOmHubMulticloudResourcesDataSource
		factories["oci_multicloud_resource_anchor"] = tf_multicloud.MulticloudResourceAnchorDataSource
		factories["oci_multicloud_resource_anchors"] = tf_multicloud.MulticloudResourceAnchorsDataSource
	}
	if common.CheckForEnabledServices("mysql") {
		factories["oci_mysql_blue_green_deployment"] = tf_mysql.MysqlBlueGreenDeploymentDataSource
		factories["oci_mysql_blue_green_deployments"] = tf_mysql.MysqlBlueGreenDeploymentsDataSource
		factories["oci_mysql_channel"] = tf_mysql.MysqlChannelDataSource
		factories["oci_mysql_channels"] = tf_mysql.MysqlChannelsDataSource
		factories["oci_mysql_db_system_maintenance_events"] = tf_mysql.MysqlDbSystemMaintenanceEventsDataSource
		factories["oci_mysql_heat_wave_cluster"] = tf_mysql.MysqlHeatWaveClusterDataSource
		factories["oci_mysql_mysql_backup"] = tf_mysql.MysqlMysqlBackupDataSource
		factories["oci_mysql_mysql_backups"] = tf_mysql.MysqlMysqlBackupsDataSource
		factories["oci_mysql_mysql_configuration"] = tf_mysql.MysqlMysqlConfigurationDataSource
		factories["oci_mysql_mysql_configurations"] = tf_mysql.MysqlMysqlConfigurationsDataSource
		factories["oci_mysql_mysql_db_system"] = tf_mysql.MysqlMysqlDbSystemDataSource
		factories["oci_mysql_mysql_db_systems"] = tf_mysql.MysqlMysqlDbSystemsDataSource
		factories["oci_mysql_mysql_versions"] = tf_mysql.MysqlMysqlVersionsDataSource
		factories["oci_mysql_replica"] = tf_mysql.MysqlReplicaDataSource
		factories["oci_mysql_replicas"] = tf_mysql.MysqlReplicasDataSource
		factories["oci_mysql_shapes"] = tf_mysql.MysqlShapesDataSource
	}
	if common.CheckForEnabledServices("networkfirewall") {
		factories["oci_network_firewall_network_firewall"] = tf_network_firewall.NetworkFirewallNetworkFirewallDataSource
		factories["oci_network_firewall_network_firewall_health_status"] = tf_network_firewall.NetworkFirewallNetworkFirewallHealthStatusDataSource
		factories["oci_network_firewall_network_firewall_policies"] = tf_network_firewall.NetworkFirewallNetworkFirewallPoliciesDataSource
		factories["oci_network_firewall_network_firewall_policy"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyDataSource
		factories["oci_network_firewall_network_firewall_policy_address_list"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyAddressListDataSource
		factories["oci_network_firewall_network_firewall_policy_address_lists"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyAddressListsDataSource
		factories["oci_network_firewall_network_firewall_policy_application"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyApplicationDataSource
		factories["oci_network_firewall_network_firewall_policy_application_group"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyApplicationGroupDataSource
		factories["oci_network_firewall_network_firewall_policy_application_groups"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyApplicationGroupsDataSource
		factories["oci_network_firewall_network_firewall_policy_applications"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyApplicationsDataSource
		factories["oci_network_firewall_network_firewall_policy_decryption_profile"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyDecryptionProfileDataSource
		factories["oci_network_firewall_network_firewall_policy_decryption_profiles"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyDecryptionProfilesDataSource
		factories["oci_network_firewall_network_firewall_policy_decryption_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyDecryptionRuleDataSource
		factories["oci_network_firewall_network_firewall_policy_decryption_rules"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyDecryptionRulesDataSource
		factories["oci_network_firewall_network_firewall_policy_mapped_secret"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyMappedSecretDataSource
		factories["oci_network_firewall_network_firewall_policy_mapped_secrets"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyMappedSecretsDataSource
		factories["oci_network_firewall_network_firewall_policy_nat_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyNatRuleDataSource
		factories["oci_network_firewall_network_firewall_policy_nat_rules"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyNatRulesDataSource
		factories["oci_network_firewall_network_firewall_policy_security_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicySecurityRuleDataSource
		factories["oci_network_firewall_network_firewall_policy_security_rules"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicySecurityRulesDataSource
		factories["oci_network_firewall_network_firewall_policy_service"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyServiceDataSource
		factories["oci_network_firewall_network_firewall_policy_service_list"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyServiceListDataSource
		factories["oci_network_firewall_network_firewall_policy_service_lists"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyServiceListsDataSource
		factories["oci_network_firewall_network_firewall_policy_services"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyServicesDataSource
		factories["oci_network_firewall_network_firewall_policy_tunnel_inspection_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleDataSource
		factories["oci_network_firewall_network_firewall_policy_tunnel_inspection_rules"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyTunnelInspectionRulesDataSource
		factories["oci_network_firewall_network_firewall_policy_url_list"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyUrlListDataSource
		factories["oci_network_firewall_network_firewall_policy_url_lists"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyUrlListsDataSource
		factories["oci_network_firewall_network_firewalls"] = tf_network_firewall.NetworkFirewallNetworkFirewallsDataSource
	}
	if common.CheckForEnabledServices("networkloadbalancer") {
		factories["oci_network_load_balancer_backend_health"] = tf_network_load_balancer.NetworkLoadBalancerBackendHealthDataSource
		factories["oci_network_load_balancer_backend_set"] = tf_network_load_balancer.NetworkLoadBalancerBackendSetDataSource
		factories["oci_network_load_balancer_backend_set_health"] = tf_network_load_balancer.NetworkLoadBalancerBackendSetHealthDataSource
		factories["oci_network_load_balancer_backend_sets"] = tf_network_load_balancer.NetworkLoadBalancerBackendSetsDataSource
		factories["oci_network_load_balancer_backends"] = tf_network_load_balancer.NetworkLoadBalancerBackendsDataSource
		factories["oci_network_load_balancer_listener"] = tf_network_load_balancer.NetworkLoadBalancerListenerDataSource
		factories["oci_network_load_balancer_listeners"] = tf_network_load_balancer.NetworkLoadBalancerListenersDataSource
		factories["oci_network_load_balancer_network_load_balancer"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancerDataSource
		factories["oci_network_load_balancer_network_load_balancer_backend_set_backend_operational_status"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancerBackendSetBackendOperationalStatusDataSource
		factories["oci_network_load_balancer_network_load_balancer_health"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancerHealthDataSource
		factories["oci_network_load_balancer_network_load_balancers"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancersDataSource
		factories["oci_network_load_balancer_network_load_balancers_policies"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancersPoliciesDataSource
		factories["oci_network_load_balancer_network_load_balancers_protocols"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancersProtocolsDataSource
	}
	if common.CheckForEnabledServices("nosql") {
		factories["oci_nosql_configuration"] = tf_nosql.NosqlConfigurationDataSource
		factories["oci_nosql_index"] = tf_nosql.NosqlIndexDataSource
		factories["oci_nosql_indexes"] = tf_nosql.NosqlIndexesDataSource
		factories["oci_nosql_table"] = tf_nosql.NosqlTableDataSource
		factories["oci_nosql_tables"] = tf_nosql.NosqlTablesDataSource
	}
	if common.CheckForEnabledServices("objectstorage") {
		factories["oci_objectstorage_bucket"] = tf_objectstorage.ObjectStorageBucketDataSource
		factories["oci_objectstorage_bucket_summaries"] = tf_objectstorage.ObjectStorageBucketsDataSource
		factories["oci_objectstorage_namespace"] = tf_objectstorage.ObjectStorageNamespaceDataSource
		factories["oci_objectstorage_namespace_metadata"] = tf_objectstorage.ObjectStorageNamespaceMetadataDataSource
		factories["oci_objectstorage_object"] = tf_objectstorage.ObjectStorageObjectDataSource
		factories["oci_objectstorage_object_head"] = tf_objectstorage.ObjectStorageObjectHeadDataSource
		factories["oci_objectstorage_object_lifecycle_policy"] = tf_objectstorage.ObjectStorageObjectLifecyclePolicyDataSource
		factories["oci_objectstorage_object_versions"] = tf_objectstorage.ObjectStorageObjectVersionsDataSource
		factories["oci_objectstorage_objects"] = tf_objectstorage.ObjectStorageObjectsDataSource
		factories["oci_objectstorage_preauthrequest"] = tf_objectstorage.ObjectStoragePreauthenticatedRequestDataSource
		factories["oci_objectstorage_preauthrequests"] = tf_objectstorage.ObjectStoragePreauthenticatedRequestsDataSource
		factories["oci_objectstorage_private_endpoint"] = tf_objectstorage.ObjectStoragePrivateEndpointDataSource
		factories["oci_objectstorage_private_endpoint_summaries"] = tf_objectstorage.ObjectStoragePrivateEndpointsDataSource
		factories["oci_objectstorage_replication_policies"] = tf_objectstorage.ObjectStorageReplicationPoliciesDataSource
		factories["oci_objectstorage_replication_policy"] = tf_objectstorage.ObjectStorageReplicationPolicyDataSource
		factories["oci_objectstorage_replication_sources"] = tf_objectstorage.ObjectStorageReplicationSourcesDataSource
	}
	if common.CheckForEnabledServices("oce") {
		factories["oci_oce_oce_instance"] = tf_oce.OceOceInstanceDataSource
		factories["oci_oce_oce_instances"] = tf_oce.OceOceInstancesDataSource
	}
	if common.CheckForEnabledServices("ocvp") {
		factories["oci_ocvp_byol"] = tf_ocvp.OcvpByolDataSource
		factories["oci_ocvp_byol_allocation"] = tf_ocvp.OcvpByolAllocationDataSource
		factories["oci_ocvp_byol_allocations"] = tf_ocvp.OcvpByolAllocationsDataSource
		factories["oci_ocvp_byols"] = tf_ocvp.OcvpByolsDataSource
		factories["oci_ocvp_cluster"] = tf_ocvp.OcvpClusterDataSource
		factories["oci_ocvp_clusters"] = tf_ocvp.OcvpClustersDataSource
		factories["oci_ocvp_datastore"] = tf_ocvp.OcvpDatastoreDataSource
		factories["oci_ocvp_datastore_cluster"] = tf_ocvp.OcvpDatastoreClusterDataSource
		factories["oci_ocvp_datastore_clusters"] = tf_ocvp.OcvpDatastoreClustersDataSource
		factories["oci_ocvp_datastores"] = tf_ocvp.OcvpDatastoresDataSource
		factories["oci_ocvp_esxi_host"] = tf_ocvp.OcvpEsxiHostDataSource
		factories["oci_ocvp_esxi_hosts"] = tf_ocvp.OcvpEsxiHostsDataSource
		factories["oci_ocvp_generate_vmware_binary_download_info"] = tf_ocvp.OcvpGenerateVmwareBinaryDownloadInfoDataSource
		factories["oci_ocvp_management_appliance"] = tf_ocvp.OcvpManagementApplianceDataSource
		factories["oci_ocvp_management_appliances"] = tf_ocvp.OcvpManagementAppliancesDataSource
		factories["oci_ocvp_retrieve_password"] = tf_ocvp.OcvpRetrievePasswordDataSource
		factories["oci_ocvp_retrieve_vmware_binaries"] = tf_ocvp.OcvpRetrieveVmwareBinariesDataSource
		factories["oci_ocvp_sddc"] = tf_ocvp.OcvpSddcDataSource
		factories["oci_ocvp_sddcs"] = tf_ocvp.OcvpSddcsDataSource
		factories["oci_ocvp_supported_commitments"] = tf_ocvp.OcvpSupportedCommitmentsDataSource
		factories["oci_ocvp_supported_host_shapes"] = tf_ocvp.OcvpSupportedHostShapesDataSource
		factories["oci_ocvp_supported_skus"] = tf_ocvp.OcvpSupportedSkusDataSource
		factories["oci_ocvp_supported_vmware_software_versions"] = tf_ocvp.OcvpSupportedVmwareSoftwareVersionsDataSource
	}
	if common.CheckForEnabledServices("oda") {
		factories["oci_oda_oda_instance"] = tf_oda.OdaOdaInstanceDataSource
		factories["oci_oda_oda_instances"] = tf_oda.OdaOdaInstancesDataSource
		factories["oci_oda_oda_private_endpoint"] = tf_oda.OdaOdaPrivateEndpointDataSource
		factories["oci_oda_oda_private_endpoint_attachment"] = tf_oda.OdaOdaPrivateEndpointAttachmentDataSource
		factories["oci_oda_oda_private_endpoint_attachments"] = tf_oda.OdaOdaPrivateEndpointAttachmentsDataSource
		factories["oci_oda_oda_private_endpoint_scan_proxies"] = tf_oda.OdaOdaPrivateEndpointScanProxiesDataSource
		factories["oci_oda_oda_private_endpoint_scan_proxy"] = tf_oda.OdaOdaPrivateEndpointScanProxyDataSource
		factories["oci_oda_oda_private_endpoints"] = tf_oda.OdaOdaPrivateEndpointsDataSource
	}
	if common.CheckForEnabledServices("onesubscription") {
		factories["oci_onesubscription_aggregated_computed_usages"] = tf_onesubscription.OnesubscriptionAggregatedComputedUsagesDataSource
		factories["oci_onesubscription_billing_schedules"] = tf_onesubscription.OnesubscriptionBillingSchedulesDataSource
		factories["oci_onesubscription_commitment"] = tf_onesubscription.OnesubscriptionCommitmentDataSource
		factories["oci_onesubscription_commitments"] = tf_onesubscription.OnesubscriptionCommitmentsDataSource
		factories["oci_onesubscription_computed_usage"] = tf_onesubscription.OnesubscriptionComputedUsageDataSource
		factories["oci_onesubscription_computed_usages"] = tf_onesubscription.OnesubscriptionComputedUsagesDataSource
		factories["oci_onesubscription_invoice_line_computed_usages"] = tf_onesubscription.OnesubscriptionInvoiceLineComputedUsagesDataSource
		factories["oci_onesubscription_invoices"] = tf_onesubscription.OnesubscriptionInvoicesDataSource
		factories["oci_onesubscription_organization_subscriptions"] = tf_onesubscription.OnesubscriptionOrganizationSubscriptionsDataSource
		factories["oci_onesubscription_ratecards"] = tf_onesubscription.OnesubscriptionRatecardsDataSource
		factories["oci_onesubscription_subscribed_service"] = tf_onesubscription.OnesubscriptionSubscribedServiceDataSource
		factories["oci_onesubscription_subscribed_services"] = tf_onesubscription.OnesubscriptionSubscribedServicesDataSource
		factories["oci_onesubscription_subscriptions"] = tf_onesubscription.OnesubscriptionSubscriptionsDataSource
	}
	if common.CheckForEnabledServices("ons") {
		factories["oci_ons_notification_topic"] = tf_ons.OnsNotificationTopicDataSource
		factories["oci_ons_notification_topics"] = tf_ons.OnsNotificationTopicsDataSource
		factories["oci_ons_subscription"] = tf_ons.OnsSubscriptionDataSource
		factories["oci_ons_subscriptions"] = tf_ons.OnsSubscriptionsDataSource
	}
	if common.CheckForEnabledServices("opa") {
		factories["oci_opa_opa_instance"] = tf_opa.OpaOpaInstanceDataSource
		factories["oci_opa_opa_instances"] = tf_opa.OpaOpaInstancesDataSource
	}
	if common.CheckForEnabledServices("opensearch") {
		factories["oci_opensearch_opensearch_cluster"] = tf_opensearch.OpensearchOpensearchClusterDataSource
		factories["oci_opensearch_opensearch_cluster_pipeline"] = tf_opensearch.OpensearchOpensearchClusterPipelineDataSource
		factories["oci_opensearch_opensearch_cluster_pipelines"] = tf_opensearch.OpensearchOpensearchClusterPipelinesDataSource
		factories["oci_opensearch_opensearch_clusters"] = tf_opensearch.OpensearchOpensearchClustersDataSource
		factories["oci_opensearch_opensearch_version"] = tf_opensearch.OpensearchOpensearchVersionDataSource
		factories["oci_opensearch_opensearch_versions"] = tf_opensearch.OpensearchOpensearchVersionsDataSource
	}
	if common.CheckForEnabledServices("operatoraccesscontrol") {
		factories["oci_operator_access_control_access_request"] = tf_operator_access_control.OperatorAccessControlAccessRequestDataSource
		factories["oci_operator_access_control_access_request_audit_log_report"] = tf_operator_access_control.OperatorAccessControlAccessRequestAuditLogReportDataSource
		factories["oci_operator_access_control_access_request_history"] = tf_operator_access_control.OperatorAccessControlAccessRequestHistoryDataSource
		factories["oci_operator_access_control_access_requests"] = tf_operator_access_control.OperatorAccessControlAccessRequestsDataSource
		factories["oci_operator_access_control_operator_action"] = tf_operator_access_control.OperatorAccessControlOperatorActionDataSource
		factories["oci_operator_access_control_operator_actions"] = tf_operator_access_control.OperatorAccessControlOperatorActionsDataSource
		factories["oci_operator_access_control_operator_control"] = tf_operator_access_control.OperatorAccessControlOperatorControlDataSource
		factories["oci_operator_access_control_operator_control_assignment"] = tf_operator_access_control.OperatorAccessControlOperatorControlAssignmentDataSource
		factories["oci_operator_access_control_operator_control_assignments"] = tf_operator_access_control.OperatorAccessControlOperatorControlAssignmentsDataSource
		factories["oci_operator_access_control_operator_controls"] = tf_operator_access_control.OperatorAccessControlOperatorControlsDataSource
	}
	if common.CheckForEnabledServices("opsi") {
		factories["oci_opsi_awr_hub"] = tf_opsi.OpsiAwrHubDataSource
		factories["oci_opsi_awr_hub_awr_snapshot"] = tf_opsi.OpsiAwrHubAwrSnapshotDataSource
		factories["oci_opsi_awr_hub_awr_snapshots"] = tf_opsi.OpsiAwrHubAwrSnapshotsDataSource
		factories["oci_opsi_awr_hub_awr_sources_summary"] = tf_opsi.OpsiAwrHubAwrSourcesSummaryDataSource
		factories["oci_opsi_awr_hub_source"] = tf_opsi.OpsiAwrHubSourceDataSource
		factories["oci_opsi_awr_hub_sources"] = tf_opsi.OpsiAwrHubSourcesDataSource
		factories["oci_opsi_awr_hubs"] = tf_opsi.OpsiAwrHubsDataSource
		factories["oci_opsi_chargeback_plan"] = tf_opsi.OpsiChargebackPlanDataSource
		factories["oci_opsi_chargeback_plans"] = tf_opsi.OpsiChargebackPlansDataSource
		factories["oci_opsi_database_insight"] = tf_opsi.OpsiDatabaseInsightDataSource
		factories["oci_opsi_database_insights"] = tf_opsi.OpsiDatabaseInsightsDataSource
		factories["oci_opsi_enterprise_manager_bridge"] = tf_opsi.OpsiEnterpriseManagerBridgeDataSource
		factories["oci_opsi_enterprise_manager_bridges"] = tf_opsi.OpsiEnterpriseManagerBridgesDataSource
		factories["oci_opsi_exadata_insight"] = tf_opsi.OpsiExadataInsightDataSource
		factories["oci_opsi_exadata_insights"] = tf_opsi.OpsiExadataInsightsDataSource
		factories["oci_opsi_host_insight"] = tf_opsi.OpsiHostInsightDataSource
		factories["oci_opsi_host_insights"] = tf_opsi.OpsiHostInsightsDataSource
		factories["oci_opsi_importable_agent_entities"] = tf_opsi.OpsiImportableAgentEntitiesDataSource
		factories["oci_opsi_importable_agent_entity"] = tf_opsi.OpsiImportableAgentEntityDataSource
		factories["oci_opsi_importable_compute_entities"] = tf_opsi.OpsiImportableComputeEntitiesDataSource
		factories["oci_opsi_importable_compute_entity"] = tf_opsi.OpsiImportableComputeEntityDataSource
		factories["oci_opsi_news_report"] = tf_opsi.OpsiNewsReportDataSource
		factories["oci_opsi_news_reports"] = tf_opsi.OpsiNewsReportsDataSource
		factories["oci_opsi_operations_insights_private_endpoint"] = tf_opsi.OpsiOperationsInsightsPrivateEndpointDataSource
		factories["oci_opsi_operations_insights_private_endpoints"] = tf_opsi.OpsiOperationsInsightsPrivateEndpointsDataSource
		factories["oci_opsi_operations_insights_warehouse"] = tf_opsi.OpsiOperationsInsightsWarehouseDataSource
		factories["oci_opsi_operations_insights_warehouse_resource_usage_summary"] = tf_opsi.OpsiOperationsInsightsWarehouseResourceUsageSummaryDataSource
		factories["oci_opsi_operations_insights_warehouse_user"] = tf_opsi.OpsiOperationsInsightsWarehouseUserDataSource
		factories["oci_opsi_operations_insights_warehouse_users"] = tf_opsi.OpsiOperationsInsightsWarehouseUsersDataSource
		factories["oci_opsi_operations_insights_warehouses"] = tf_opsi.OpsiOperationsInsightsWarehousesDataSource
		factories["oci_opsi_opsi_configuration"] = tf_opsi.OpsiOpsiConfigurationDataSource
		factories["oci_opsi_opsi_configuration_configuration_item"] = tf_opsi.OpsiOpsiConfigurationConfigurationItemDataSource
		factories["oci_opsi_opsi_configurations"] = tf_opsi.OpsiOpsiConfigurationsDataSource
	}
	if common.CheckForEnabledServices("optimizer") {
		factories["oci_optimizer_categories"] = tf_optimizer.OptimizerCategoriesDataSource
		factories["oci_optimizer_category"] = tf_optimizer.OptimizerCategoryDataSource
		factories["oci_optimizer_enrollment_status"] = tf_optimizer.OptimizerEnrollmentStatusDataSource
		factories["oci_optimizer_enrollment_statuses"] = tf_optimizer.OptimizerEnrollmentStatusesDataSource
		factories["oci_optimizer_histories"] = tf_optimizer.OptimizerHistoriesDataSource
		factories["oci_optimizer_profile"] = tf_optimizer.OptimizerProfileDataSource
		factories["oci_optimizer_profile_level"] = tf_optimizer.OptimizerProfileLevelDataSource
		factories["oci_optimizer_profile_levels"] = tf_optimizer.OptimizerProfileLevelsDataSource
		factories["oci_optimizer_profiles"] = tf_optimizer.OptimizerProfilesDataSource
		factories["oci_optimizer_recommendation"] = tf_optimizer.OptimizerRecommendationDataSource
		factories["oci_optimizer_recommendation_strategies"] = tf_optimizer.OptimizerRecommendationStrategiesDataSource
		factories["oci_optimizer_recommendation_strategy"] = tf_optimizer.OptimizerRecommendationStrategyDataSource
		factories["oci_optimizer_recommendations"] = tf_optimizer.OptimizerRecommendationsDataSource
		factories["oci_optimizer_resource_action"] = tf_optimizer.OptimizerResourceActionDataSource
		factories["oci_optimizer_resource_actions"] = tf_optimizer.OptimizerResourceActionsDataSource
	}
	if common.CheckForEnabledServices("osmanagementhub") {
		factories["oci_os_management_hub_dynamic_set"] = tf_os_management_hub.OsManagementHubDynamicSetDataSource
		factories["oci_os_management_hub_dynamic_set_managed_instances"] = tf_os_management_hub.OsManagementHubDynamicSetManagedInstancesDataSource
		factories["oci_os_management_hub_dynamic_sets"] = tf_os_management_hub.OsManagementHubDynamicSetsDataSource
		factories["oci_os_management_hub_entitlements"] = tf_os_management_hub.OsManagementHubEntitlementsDataSource
		factories["oci_os_management_hub_errata"] = tf_os_management_hub.OsManagementHubErrataDataSource
		factories["oci_os_management_hub_erratum"] = tf_os_management_hub.OsManagementHubErratumDataSource
		factories["oci_os_management_hub_event"] = tf_os_management_hub.OsManagementHubEventDataSource
		factories["oci_os_management_hub_events"] = tf_os_management_hub.OsManagementHubEventsDataSource
		factories["oci_os_management_hub_lifecycle_environment"] = tf_os_management_hub.OsManagementHubLifecycleEnvironmentDataSource
		factories["oci_os_management_hub_lifecycle_environments"] = tf_os_management_hub.OsManagementHubLifecycleEnvironmentsDataSource
		factories["oci_os_management_hub_lifecycle_stage"] = tf_os_management_hub.OsManagementHubLifecycleStageDataSource
		factories["oci_os_management_hub_lifecycle_stages"] = tf_os_management_hub.OsManagementHubLifecycleStagesDataSource
		factories["oci_os_management_hub_managed_instance"] = tf_os_management_hub.OsManagementHubManagedInstanceDataSource
		factories["oci_os_management_hub_managed_instance_available_packages"] = tf_os_management_hub.OsManagementHubManagedInstanceAvailablePackagesDataSource
		factories["oci_os_management_hub_managed_instance_available_software_sources"] = tf_os_management_hub.OsManagementHubManagedInstanceAvailableSoftwareSourcesDataSource
		factories["oci_os_management_hub_managed_instance_available_windows_updates"] = tf_os_management_hub.OsManagementHubManagedInstanceAvailableWindowsUpdatesDataSource
		factories["oci_os_management_hub_managed_instance_errata"] = tf_os_management_hub.OsManagementHubManagedInstanceErrataDataSource
		factories["oci_os_management_hub_managed_instance_group"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupDataSource
		factories["oci_os_management_hub_managed_instance_group_available_modules"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupAvailableModulesDataSource
		factories["oci_os_management_hub_managed_instance_group_available_packages"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupAvailablePackagesDataSource
		factories["oci_os_management_hub_managed_instance_group_available_software_sources"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupAvailableSoftwareSourcesDataSource
		factories["oci_os_management_hub_managed_instance_group_installed_packages"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupInstalledPackagesDataSource
		factories["oci_os_management_hub_managed_instance_group_managed_instances"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupManagedInstancesDataSource
		factories["oci_os_management_hub_managed_instance_group_modules"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupModulesDataSource
		factories["oci_os_management_hub_managed_instance_groups"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupsDataSource
		factories["oci_os_management_hub_managed_instance_installed_packages"] = tf_os_management_hub.OsManagementHubManagedInstanceInstalledPackagesDataSource
		factories["oci_os_management_hub_managed_instance_installed_windows_updates"] = tf_os_management_hub.OsManagementHubManagedInstanceInstalledWindowsUpdatesDataSource
		factories["oci_os_management_hub_managed_instance_modules"] = tf_os_management_hub.OsManagementHubManagedInstanceModulesDataSource
		factories["oci_os_management_hub_managed_instance_snaps"] = tf_os_management_hub.OsManagementHubManagedInstanceSnapsDataSource
		factories["oci_os_management_hub_managed_instance_updatable_packages"] = tf_os_management_hub.OsManagementHubManagedInstanceUpdatablePackagesDataSource
		factories["oci_os_management_hub_managed_instances"] = tf_os_management_hub.OsManagementHubManagedInstancesDataSource
		factories["oci_os_management_hub_management_station"] = tf_os_management_hub.OsManagementHubManagementStationDataSource
		factories["oci_os_management_hub_management_station_mirrors"] = tf_os_management_hub.OsManagementHubManagementStationMirrorsDataSource
		factories["oci_os_management_hub_management_stations"] = tf_os_management_hub.OsManagementHubManagementStationsDataSource
		factories["oci_os_management_hub_profile"] = tf_os_management_hub.OsManagementHubProfileDataSource
		factories["oci_os_management_hub_profile_available_software_sources"] = tf_os_management_hub.OsManagementHubProfileAvailableSoftwareSourcesDataSource
		factories["oci_os_management_hub_profile_version"] = tf_os_management_hub.OsManagementHubProfileVersionDataSource
		factories["oci_os_management_hub_profiles"] = tf_os_management_hub.OsManagementHubProfilesDataSource
		factories["oci_os_management_hub_scheduled_job"] = tf_os_management_hub.OsManagementHubScheduledJobDataSource
		factories["oci_os_management_hub_scheduled_jobs"] = tf_os_management_hub.OsManagementHubScheduledJobsDataSource
		factories["oci_os_management_hub_software_package"] = tf_os_management_hub.OsManagementHubSoftwarePackageDataSource
		factories["oci_os_management_hub_software_package_software_source"] = tf_os_management_hub.OsManagementHubSoftwarePackageSoftwareSourceDataSource
		factories["oci_os_management_hub_software_packages"] = tf_os_management_hub.OsManagementHubSoftwarePackagesDataSource
		factories["oci_os_management_hub_software_source"] = tf_os_management_hub.OsManagementHubSoftwareSourceDataSource
		factories["oci_os_management_hub_software_source_available_software_packages"] = tf_os_management_hub.OsManagementHubSoftwareSourceAvailableSoftwarePackagesDataSource
		factories["oci_os_management_hub_software_source_manifest"] = tf_os_management_hub.OsManagementHubSoftwareSourceManifestDataSource
		factories["oci_os_management_hub_software_source_module_stream"] = tf_os_management_hub.OsManagementHubSoftwareSourceModuleStreamDataSource
		factories["oci_os_management_hub_software_source_module_stream_profile"] = tf_os_management_hub.OsManagementHubSoftwareSourceModuleStreamProfileDataSource
		factories["oci_os_management_hub_software_source_module_stream_profiles"] = tf_os_management_hub.OsManagementHubSoftwareSourceModuleStreamProfilesDataSource
		factories["oci_os_management_hub_software_source_module_streams"] = tf_os_management_hub.OsManagementHubSoftwareSourceModuleStreamsDataSource
		factories["oci_os_management_hub_software_source_package_group"] = tf_os_management_hub.OsManagementHubSoftwareSourcePackageGroupDataSource
		factories["oci_os_management_hub_software_source_package_groups"] = tf_os_management_hub.OsManagementHubSoftwareSourcePackageGroupsDataSource
		factories["oci_os_management_hub_software_source_software_package"] = tf_os_management_hub.OsManagementHubSoftwareSourceSoftwarePackageDataSource
		factories["oci_os_management_hub_software_source_software_packages"] = tf_os_management_hub.OsManagementHubSoftwareSourceSoftwarePackagesDataSource
		factories["oci_os_management_hub_software_source_vendors"] = tf_os_management_hub.OsManagementHubSoftwareSourceVendorsDataSource
		factories["oci_os_management_hub_software_sources"] = tf_os_management_hub.OsManagementHubSoftwareSourcesDataSource
		factories["oci_os_management_hub_windows_update"] = tf_os_management_hub.OsManagementHubWindowsUpdateDataSource
		factories["oci_os_management_hub_windows_updates"] = tf_os_management_hub.OsManagementHubWindowsUpdatesDataSource
	}
	if common.CheckForEnabledServices("ospgateway") {
		factories["oci_osp_gateway_address"] = tf_osp_gateway.OspGatewayAddressDataSource
		factories["oci_osp_gateway_address_rule"] = tf_osp_gateway.OspGatewayAddressRuleDataSource
		factories["oci_osp_gateway_invoice"] = tf_osp_gateway.OspGatewayInvoiceDataSource
		factories["oci_osp_gateway_invoices"] = tf_osp_gateway.OspGatewayInvoicesDataSource
		factories["oci_osp_gateway_invoices_invoice_line"] = tf_osp_gateway.OspGatewayInvoicesInvoiceLineDataSource
		factories["oci_osp_gateway_invoices_invoice_lines"] = tf_osp_gateway.OspGatewayInvoicesInvoiceLinesDataSource
		factories["oci_osp_gateway_subscription"] = tf_osp_gateway.OspGatewaySubscriptionDataSource
		factories["oci_osp_gateway_subscriptions"] = tf_osp_gateway.OspGatewaySubscriptionsDataSource
	}
	if common.CheckForEnabledServices("osubbillingschedule") {
		factories["oci_osub_billing_schedule_billing_schedules"] = tf_osub_billing_schedule.OsubBillingScheduleBillingSchedulesDataSource
	}
	if common.CheckForEnabledServices("osuborganizationsubscription") {
		factories["oci_osub_organization_subscription_organization_subscriptions"] = tf_osub_organization_subscription.OsubOrganizationSubscriptionOrganizationSubscriptionsDataSource
	}
	if common.CheckForEnabledServices("osubsubscription") {
		factories["oci_osub_subscription_commitment"] = tf_osub_subscription.OsubSubscriptionCommitmentDataSource
		factories["oci_osub_subscription_commitments"] = tf_osub_subscription.OsubSubscriptionCommitmentsDataSource
		factories["oci_osub_subscription_ratecards"] = tf_osub_subscription.OsubSubscriptionRatecardsDataSource
		factories["oci_osub_subscription_subscriptions"] = tf_osub_subscription.OsubSubscriptionSubscriptionsDataSource
	}
	if common.CheckForEnabledServices("osubusage") {
		factories["oci_osub_usage_computed_usage"] = tf_osub_usage.OsubUsageComputedUsageDataSource
		factories["oci_osub_usage_computed_usage_aggregateds"] = tf_osub_usage.OsubUsageComputedUsageAggregatedsDataSource
		factories["oci_osub_usage_computed_usages"] = tf_osub_usage.OsubUsageComputedUsagesDataSource
	}
	if common.CheckForEnabledServices("psa") {
		factories["oci_psa_private_service_access"] = tf_psa.PsaPrivateServiceAccessDataSource
		factories["oci_psa_private_service_accesses"] = tf_psa.PsaPrivateServiceAccessesDataSource
		factories["oci_psa_psa_services"] = tf_psa.PsaPsaServicesDataSource
		factories["oci_psa_psa_work_request"] = tf_psa.PsaPsaWorkRequestDataSource
		factories["oci_psa_psa_work_request_errors"] = tf_psa.PsaWorkRequestErrorsDataSource
		factories["oci_psa_psa_work_request_logs"] = tf_psa.PsaWorkRequestLogEntriesDataSource
		factories["oci_psa_psa_work_requests"] = tf_psa.PsaPsaWorkRequestsDataSource
	}
	if common.CheckForEnabledServices("psql") {
		factories["oci_psql_backup"] = tf_psql.PsqlBackupDataSource
		factories["oci_psql_backups"] = tf_psql.PsqlBackupsDataSource
		factories["oci_psql_configuration"] = tf_psql.PsqlConfigurationDataSource
		factories["oci_psql_configurations"] = tf_psql.PsqlConfigurationsDataSource
		factories["oci_psql_db_system"] = tf_psql.PsqlDbSystemDataSource
		factories["oci_psql_db_system_connection_detail"] = tf_psql.PsqlDbSystemConnectionDetailDataSource
		factories["oci_psql_db_system_pitr_detail"] = tf_psql.PsqlDbSystemPitrDetailDataSource
		factories["oci_psql_db_system_primary_db_instance"] = tf_psql.PsqlDbSystemPrimaryDbInstanceDataSource
		factories["oci_psql_db_system_replicas"] = tf_psql.PsqlDbSystemReplicasDataSource
		factories["oci_psql_db_systems"] = tf_psql.PsqlDbSystemsDataSource
		factories["oci_psql_default_configuration"] = tf_psql.PsqlDefaultConfigurationDataSource
		factories["oci_psql_default_configurations"] = tf_psql.PsqlDefaultConfigurationsDataSource
		factories["oci_psql_insight_capabilities"] = tf_psql.PsqlInsightCapabilitiesDataSource
		factories["oci_psql_shapes"] = tf_psql.PsqlShapesDataSource
	}
	if common.CheckForEnabledServices("queue") {
		factories["oci_queue_consumer_group"] = tf_queue.QueueConsumerGroupDataSource
		factories["oci_queue_consumer_groups"] = tf_queue.QueueConsumerGroupsDataSource
		factories["oci_queue_queue"] = tf_queue.QueueQueueDataSource
		factories["oci_queue_queues"] = tf_queue.QueueQueuesDataSource
	}
	if common.CheckForEnabledServices("recovery") {
		factories["oci_recovery_protected_database"] = tf_recovery.RecoveryProtectedDatabaseDataSource
		factories["oci_recovery_protected_database_fetch_configuration"] = tf_recovery.RecoveryProtectedDatabaseFetchConfigurationDataSource
		factories["oci_recovery_protected_databases"] = tf_recovery.RecoveryProtectedDatabasesDataSource
		factories["oci_recovery_protection_policies"] = tf_recovery.RecoveryProtectionPoliciesDataSource
		factories["oci_recovery_protection_policy"] = tf_recovery.RecoveryProtectionPolicyDataSource
		factories["oci_recovery_recovery_service_subnet"] = tf_recovery.RecoveryRecoveryServiceSubnetDataSource
		factories["oci_recovery_recovery_service_subnets"] = tf_recovery.RecoveryRecoveryServiceSubnetsDataSource
	}
	if common.CheckForEnabledServices("redis") {
		factories["oci_redis_oci_cache_backup"] = tf_redis.RedisOciCacheBackupDataSource
		factories["oci_redis_oci_cache_backups"] = tf_redis.RedisOciCacheBackupsDataSource
		factories["oci_redis_oci_cache_config_set"] = tf_redis.RedisOciCacheConfigSetDataSource
		factories["oci_redis_oci_cache_config_sets"] = tf_redis.RedisOciCacheConfigSetsDataSource
		factories["oci_redis_oci_cache_default_config_set"] = tf_redis.RedisOciCacheDefaultConfigSetDataSource
		factories["oci_redis_oci_cache_default_config_sets"] = tf_redis.RedisOciCacheDefaultConfigSetsDataSource
		factories["oci_redis_oci_cache_engine_options"] = tf_redis.RedisOciCacheEngineOptionsDataSource
		factories["oci_redis_oci_cache_user"] = tf_redis.RedisOciCacheUserDataSource
		factories["oci_redis_oci_cache_users"] = tf_redis.RedisOciCacheUsersDataSource
		factories["oci_redis_redis_cluster"] = tf_redis.RedisRedisClusterDataSource
		factories["oci_redis_redis_cluster_nodes"] = tf_redis.RedisRedisClusterNodesDataSource
		factories["oci_redis_redis_clusters"] = tf_redis.RedisRedisClustersDataSource
	}
	if common.CheckForEnabledServices("resourceanalytics") {
		factories["oci_resource_analytics_monitored_region"] = tf_resource_analytics.ResourceAnalyticsMonitoredRegionDataSource
		factories["oci_resource_analytics_monitored_regions"] = tf_resource_analytics.ResourceAnalyticsMonitoredRegionsDataSource
		factories["oci_resource_analytics_resource_analytics_instance"] = tf_resource_analytics.ResourceAnalyticsResourceAnalyticsInstanceDataSource
		factories["oci_resource_analytics_resource_analytics_instances"] = tf_resource_analytics.ResourceAnalyticsResourceAnalyticsInstancesDataSource
		factories["oci_resource_analytics_tenancy_attachment"] = tf_resource_analytics.ResourceAnalyticsTenancyAttachmentDataSource
		factories["oci_resource_analytics_tenancy_attachments"] = tf_resource_analytics.ResourceAnalyticsTenancyAttachmentsDataSource
	}
	if common.CheckForEnabledServices("resourcescheduler") {
		factories["oci_resource_scheduler_schedule"] = tf_resource_scheduler.ResourceSchedulerScheduleDataSource
		factories["oci_resource_scheduler_schedules"] = tf_resource_scheduler.ResourceSchedulerSchedulesDataSource
	}
	if common.CheckForEnabledServices("resourcesearch") {
		factories["oci_resource_search"] = tf_resource_search.ResourceSearchDataSource
	}
	if common.CheckForEnabledServices("resourcemanager") {
		factories["oci_resourcemanager_private_endpoint"] = tf_resourcemanager.ResourcemanagerPrivateEndpointDataSource
		factories["oci_resourcemanager_private_endpoint_reachable_ip"] = tf_resourcemanager.ResourcemanagerPrivateEndpointReachableIpDataSource
		factories["oci_resourcemanager_private_endpoints"] = tf_resourcemanager.ResourcemanagerPrivateEndpointsDataSource
		factories["oci_resourcemanager_stack"] = tf_resourcemanager.ResourcemanagerStackDataSource
		factories["oci_resourcemanager_stack_tf_state"] = tf_resourcemanager.ResourcemanagerStackTfStateDataSource
		factories["oci_resourcemanager_stacks"] = tf_resourcemanager.ResourcemanagerStacksDataSource
	}
	if common.CheckForEnabledServices("sch") {
		factories["oci_sch_connector_plugin"] = tf_sch.SchConnectorPluginDataSource
		factories["oci_sch_connector_plugins"] = tf_sch.SchConnectorPluginsDataSource
		factories["oci_sch_service_connector"] = tf_sch.SchServiceConnectorDataSource
		factories["oci_sch_service_connectors"] = tf_sch.SchServiceConnectorsDataSource
	}
	if common.CheckForEnabledServices("secrets") {
		factories["oci_secrets_secretbundle"] = tf_secrets.SecretsSecretbundleDataSource
		factories["oci_secrets_secretbundle_versions"] = tf_secrets.SecretsSecretbundleVersionsDataSource
	}
	if common.CheckForEnabledServices("securityattribute") {
		factories["oci_security_attribute_security_attribute"] = tf_security_attribute.SecurityAttributeSecurityAttributeDataSource
		factories["oci_security_attribute_security_attribute_namespace"] = tf_security_attribute.SecurityAttributeSecurityAttributeNamespaceDataSource
		factories["oci_security_attribute_security_attribute_namespaces"] = tf_security_attribute.SecurityAttributeSecurityAttributeNamespacesDataSource
		factories["oci_security_attribute_security_attributes"] = tf_security_attribute.SecurityAttributeSecurityAttributesDataSource
	}
	if common.CheckForEnabledServices("self") {
		factories["oci_self_partner_subscriptions"] = tf_self.SelfPartnerSubscriptionsDataSource
		factories["oci_self_partners"] = tf_self.SelfPartnersDataSource
		factories["oci_self_self_partner_subscriptions"] = tf_self.SelfSelfPartnerSubscriptionsDataSource
		factories["oci_self_subscription"] = tf_self.SelfSubscriptionDataSource
		factories["oci_self_subscription_token"] = tf_self.SelfSubscriptionTokenDataSource
		factories["oci_self_subscriptions"] = tf_self.SelfSubscriptionsDataSource
	}
	if common.CheckForEnabledServices("servicecatalog") {
		factories["oci_service_catalog_all_applications"] = tf_service_catalog.ServiceCatalogAllApplicationsDataSource
		factories["oci_service_catalog_configuration"] = tf_service_catalog.ServiceCatalogConfigurationDataSource
		factories["oci_service_catalog_private_application"] = tf_service_catalog.ServiceCatalogPrivateApplicationDataSource
		factories["oci_service_catalog_private_application_package"] = tf_service_catalog.ServiceCatalogPrivateApplicationPackageDataSource
		factories["oci_service_catalog_private_application_packages"] = tf_service_catalog.ServiceCatalogPrivateApplicationPackagesDataSource
		factories["oci_service_catalog_private_applications"] = tf_service_catalog.ServiceCatalogPrivateApplicationsDataSource
		factories["oci_service_catalog_service_catalog"] = tf_service_catalog.ServiceCatalogServiceCatalogDataSource
		factories["oci_service_catalog_service_catalog_association"] = tf_service_catalog.ServiceCatalogServiceCatalogAssociationDataSource
		factories["oci_service_catalog_service_catalog_associations"] = tf_service_catalog.ServiceCatalogServiceCatalogAssociationsDataSource
		factories["oci_service_catalog_service_catalogs"] = tf_service_catalog.ServiceCatalogServiceCatalogsDataSource
	}
	if common.CheckForEnabledServices("servicemanagerproxy") {
		factories["oci_service_manager_proxy_service_environment"] = tf_service_manager_proxy.ServiceManagerProxyServiceEnvironmentDataSource
		factories["oci_service_manager_proxy_service_environments"] = tf_service_manager_proxy.ServiceManagerProxyServiceEnvironmentsDataSource
	}
	if common.CheckForEnabledServices("stackmonitoring") {
		factories["oci_stack_monitoring_baselineable_metric"] = tf_stack_monitoring.StackMonitoringBaselineableMetricDataSource
		factories["oci_stack_monitoring_baselineable_metrics"] = tf_stack_monitoring.StackMonitoringBaselineableMetricsDataSource
		factories["oci_stack_monitoring_baselineable_metrics_evaluate"] = tf_stack_monitoring.StackMonitoringBaselineableMetricsEvaluateDataSource
		factories["oci_stack_monitoring_config"] = tf_stack_monitoring.StackMonitoringConfigDataSource
		factories["oci_stack_monitoring_configs"] = tf_stack_monitoring.StackMonitoringConfigsDataSource
		factories["oci_stack_monitoring_defined_monitoring_templates"] = tf_stack_monitoring.StackMonitoringDefinedMonitoringTemplatesDataSource
		factories["oci_stack_monitoring_discovery_job"] = tf_stack_monitoring.StackMonitoringDiscoveryJobDataSource
		factories["oci_stack_monitoring_discovery_job_logs"] = tf_stack_monitoring.StackMonitoringDiscoveryJobLogsDataSource
		factories["oci_stack_monitoring_discovery_jobs"] = tf_stack_monitoring.StackMonitoringDiscoveryJobsDataSource
		factories["oci_stack_monitoring_maintenance_window"] = tf_stack_monitoring.StackMonitoringMaintenanceWindowDataSource
		factories["oci_stack_monitoring_maintenance_windows"] = tf_stack_monitoring.StackMonitoringMaintenanceWindowsDataSource
		factories["oci_stack_monitoring_metric_extension"] = tf_stack_monitoring.StackMonitoringMetricExtensionDataSource
		factories["oci_stack_monitoring_metric_extensions"] = tf_stack_monitoring.StackMonitoringMetricExtensionsDataSource
		factories["oci_stack_monitoring_monitored_resource"] = tf_stack_monitoring.StackMonitoringMonitoredResourceDataSource
		factories["oci_stack_monitoring_monitored_resource_task"] = tf_stack_monitoring.StackMonitoringMonitoredResourceTaskDataSource
		factories["oci_stack_monitoring_monitored_resource_tasks"] = tf_stack_monitoring.StackMonitoringMonitoredResourceTasksDataSource
		factories["oci_stack_monitoring_monitored_resource_type"] = tf_stack_monitoring.StackMonitoringMonitoredResourceTypeDataSource
		factories["oci_stack_monitoring_monitored_resource_types"] = tf_stack_monitoring.StackMonitoringMonitoredResourceTypesDataSource
		factories["oci_stack_monitoring_monitored_resources"] = tf_stack_monitoring.StackMonitoringMonitoredResourcesDataSource
		factories["oci_stack_monitoring_monitoring_template"] = tf_stack_monitoring.StackMonitoringMonitoringTemplateDataSource
		factories["oci_stack_monitoring_monitoring_template_alarm_condition"] = tf_stack_monitoring.StackMonitoringMonitoringTemplateAlarmConditionDataSource
		factories["oci_stack_monitoring_monitoring_template_alarm_conditions"] = tf_stack_monitoring.StackMonitoringMonitoringTemplateAlarmConditionsDataSource
		factories["oci_stack_monitoring_monitoring_templates"] = tf_stack_monitoring.StackMonitoringMonitoringTemplatesDataSource
		factories["oci_stack_monitoring_process_set"] = tf_stack_monitoring.StackMonitoringProcessSetDataSource
		factories["oci_stack_monitoring_process_sets"] = tf_stack_monitoring.StackMonitoringProcessSetsDataSource
	}
	if common.CheckForEnabledServices("streaming") {
		factories["oci_streaming_connect_harness"] = tf_streaming.StreamingConnectHarnessDataSource
		factories["oci_streaming_connect_harnesses"] = tf_streaming.StreamingConnectHarnessesDataSource
		factories["oci_streaming_stream"] = tf_streaming.StreamingStreamDataSource
		factories["oci_streaming_stream_pool"] = tf_streaming.StreamingStreamPoolDataSource
		factories["oci_streaming_stream_pools"] = tf_streaming.StreamingStreamPoolsDataSource
		factories["oci_streaming_streams"] = tf_streaming.StreamingStreamsDataSource
	}
	if common.CheckForEnabledServices("tenantmanagercontrolplane") {
		factories["oci_tenantmanagercontrolplane_assigned_subscription"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneAssignedSubscriptionDataSource
		factories["oci_tenantmanagercontrolplane_assigned_subscription_line_items"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneAssignedSubscriptionLineItemsDataSource
		factories["oci_tenantmanagercontrolplane_assigned_subscriptions"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneAssignedSubscriptionsDataSource
		factories["oci_tenantmanagercontrolplane_domain"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneDomainDataSource
		factories["oci_tenantmanagercontrolplane_domain_governance"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneDomainGovernanceDataSource
		factories["oci_tenantmanagercontrolplane_domain_governances"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneDomainGovernancesDataSource
		factories["oci_tenantmanagercontrolplane_domains"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneDomainsDataSource
		factories["oci_tenantmanagercontrolplane_link"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneLinkDataSource
		factories["oci_tenantmanagercontrolplane_link_features"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneLinkFeaturesDataSource
		factories["oci_tenantmanagercontrolplane_link_tenancy_name"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneLinkTenancyNameDataSource
		factories["oci_tenantmanagercontrolplane_links"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneLinksDataSource
		factories["oci_tenantmanagercontrolplane_organization"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneOrganizationDataSource
		factories["oci_tenantmanagercontrolplane_organization_tenancies"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneOrganizationTenanciesDataSource
		factories["oci_tenantmanagercontrolplane_organization_tenancy"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneOrganizationTenancyDataSource
		factories["oci_tenantmanagercontrolplane_organizations"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneOrganizationsDataSource
		factories["oci_tenantmanagercontrolplane_recipient_invitation"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneRecipientInvitationDataSource
		factories["oci_tenantmanagercontrolplane_recipient_invitations"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneRecipientInvitationsDataSource
		factories["oci_tenantmanagercontrolplane_sender_invitation"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSenderInvitationDataSource
		factories["oci_tenantmanagercontrolplane_sender_invitations"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSenderInvitationsDataSource
		factories["oci_tenantmanagercontrolplane_subscription"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSubscriptionDataSource
		factories["oci_tenantmanagercontrolplane_subscription_available_regions"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSubscriptionAvailableRegionsDataSource
		factories["oci_tenantmanagercontrolplane_subscription_line_items"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSubscriptionLineItemsDataSource
		factories["oci_tenantmanagercontrolplane_subscription_mapping"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSubscriptionMappingDataSource
		factories["oci_tenantmanagercontrolplane_subscription_mappings"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSubscriptionMappingsDataSource
		factories["oci_tenantmanagercontrolplane_subscriptions"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSubscriptionsDataSource
	}
	if common.CheckForEnabledServices("usageproxy") {
		factories["oci_usage_proxy_resource_quotas"] = tf_usage_proxy.UsageProxyResourceQuotasDataSource
		factories["oci_usage_proxy_resources"] = tf_usage_proxy.UsageProxyResourcesDataSource
		factories["oci_usage_proxy_subscription_product"] = tf_usage_proxy.UsageProxySubscriptionProductDataSource
		factories["oci_usage_proxy_subscription_products"] = tf_usage_proxy.UsageProxySubscriptionProductsDataSource
		factories["oci_usage_proxy_subscription_redeemable_user"] = tf_usage_proxy.UsageProxySubscriptionRedeemableUserDataSource
		factories["oci_usage_proxy_subscription_redeemable_users"] = tf_usage_proxy.UsageProxySubscriptionRedeemableUsersDataSource
		factories["oci_usage_proxy_subscription_redemption"] = tf_usage_proxy.UsageProxySubscriptionRedemptionDataSource
		factories["oci_usage_proxy_subscription_redemptions"] = tf_usage_proxy.UsageProxySubscriptionRedemptionsDataSource
		factories["oci_usage_proxy_subscription_reward"] = tf_usage_proxy.UsageProxySubscriptionRewardDataSource
		factories["oci_usage_proxy_subscription_rewards"] = tf_usage_proxy.UsageProxySubscriptionRewardsDataSource
		factories["oci_usage_proxy_usagelimits"] = tf_usage_proxy.UsageProxyUsagelimitsDataSource
	}
	if common.CheckForEnabledServices("vault") {
		factories["oci_vault_secret"] = tf_vault.VaultSecretDataSource
		factories["oci_vault_secret_version_sdk_v2"] = tf_vault.VaultSecretVersionDataSource
		factories["oci_vault_secrets"] = tf_vault.VaultSecretsDataSource
	}
	if common.CheckForEnabledServices("vbsinst") {
		factories["oci_vbs_inst_vbs_instance"] = tf_vbs_inst.VbsInstVbsInstanceDataSource
		factories["oci_vbs_inst_vbs_instances"] = tf_vbs_inst.VbsInstVbsInstancesDataSource
	}
	if common.CheckForEnabledServices("visualbuilder") {
		factories["oci_visual_builder_vb_instance"] = tf_visual_builder.VisualBuilderVbInstanceDataSource
		factories["oci_visual_builder_vb_instance_applications"] = tf_visual_builder.VisualBuilderVbInstanceApplicationsDataSource
		factories["oci_visual_builder_vb_instances"] = tf_visual_builder.VisualBuilderVbInstancesDataSource
	}
	if common.CheckForEnabledServices("vnmonitoring") {
		factories["oci_vn_monitoring_path_analyzer_test"] = tf_vn_monitoring.VnMonitoringPathAnalyzerTestDataSource
		factories["oci_vn_monitoring_path_analyzer_tests"] = tf_vn_monitoring.VnMonitoringPathAnalyzerTestsDataSource
	}
	if common.CheckForEnabledServices("vulnerabilityscanning") {
		factories["oci_vulnerability_scanning_container_scan_recipe"] = tf_vulnerability_scanning.VulnerabilityScanningContainerScanRecipeDataSource
		factories["oci_vulnerability_scanning_container_scan_recipes"] = tf_vulnerability_scanning.VulnerabilityScanningContainerScanRecipesDataSource
		factories["oci_vulnerability_scanning_container_scan_target"] = tf_vulnerability_scanning.VulnerabilityScanningContainerScanTargetDataSource
		factories["oci_vulnerability_scanning_container_scan_targets"] = tf_vulnerability_scanning.VulnerabilityScanningContainerScanTargetsDataSource
		factories["oci_vulnerability_scanning_host_scan_recipe"] = tf_vulnerability_scanning.VulnerabilityScanningHostScanRecipeDataSource
		factories["oci_vulnerability_scanning_host_scan_recipes"] = tf_vulnerability_scanning.VulnerabilityScanningHostScanRecipesDataSource
		factories["oci_vulnerability_scanning_host_scan_target"] = tf_vulnerability_scanning.VulnerabilityScanningHostScanTargetDataSource
		factories["oci_vulnerability_scanning_host_scan_target_errors"] = tf_vulnerability_scanning.VulnerabilityScanningHostScanTargetErrorsDataSource
		factories["oci_vulnerability_scanning_host_scan_targets"] = tf_vulnerability_scanning.VulnerabilityScanningHostScanTargetsDataSource
	}
	if common.CheckForEnabledServices("waa") {
		factories["oci_waa_web_app_acceleration"] = tf_waa.WaaWebAppAccelerationDataSource
		factories["oci_waa_web_app_acceleration_policies"] = tf_waa.WaaWebAppAccelerationPoliciesDataSource
		factories["oci_waa_web_app_acceleration_policy"] = tf_waa.WaaWebAppAccelerationPolicyDataSource
		factories["oci_waa_web_app_accelerations"] = tf_waa.WaaWebAppAccelerationsDataSource
	}
	if common.CheckForEnabledServices("waas") {
		factories["oci_waas_address_list"] = tf_waas.WaasAddressListDataSource
		factories["oci_waas_address_lists"] = tf_waas.WaasAddressListsDataSource
		factories["oci_waas_certificate"] = tf_waas.WaasCertificateDataSource
		factories["oci_waas_certificates"] = tf_waas.WaasCertificatesDataSource
		factories["oci_waas_custom_protection_rule"] = tf_waas.WaasCustomProtectionRuleDataSource
		factories["oci_waas_custom_protection_rules"] = tf_waas.WaasCustomProtectionRulesDataSource
		factories["oci_waas_edge_subnets"] = tf_waas.WaasEdgeSubnetsDataSource
		factories["oci_waas_http_redirect"] = tf_waas.WaasHttpRedirectDataSource
		factories["oci_waas_http_redirects"] = tf_waas.WaasHttpRedirectsDataSource
		factories["oci_waas_protection_rule"] = tf_waas.WaasProtectionRuleDataSource
		factories["oci_waas_protection_rules"] = tf_waas.WaasProtectionRulesDataSource
		factories["oci_waas_waas_policies"] = tf_waas.WaasWaasPoliciesDataSource
		factories["oci_waas_waas_policy"] = tf_waas.WaasWaasPolicyDataSource
	}
	if common.CheckForEnabledServices("waf") {
		factories["oci_waf_network_address_list"] = tf_waf.WafNetworkAddressListDataSource
		factories["oci_waf_network_address_lists"] = tf_waf.WafNetworkAddressListsDataSource
		factories["oci_waf_protection_capabilities"] = tf_waf.WafProtectionCapabilitiesDataSource
		factories["oci_waf_protection_capability_group_tags"] = tf_waf.WafProtectionCapabilityGroupTagsDataSource
		factories["oci_waf_web_app_firewall"] = tf_waf.WafWebAppFirewallDataSource
		factories["oci_waf_web_app_firewall_policies"] = tf_waf.WafWebAppFirewallPoliciesDataSource
		factories["oci_waf_web_app_firewall_policy"] = tf_waf.WafWebAppFirewallPolicyDataSource
		factories["oci_waf_web_app_firewalls"] = tf_waf.WafWebAppFirewallsDataSource
	}
	if common.CheckForEnabledServices("zpr") {
		factories["oci_zpr_configuration"] = tf_zpr.ZprConfigurationDataSource
		factories["oci_zpr_zpr_policies"] = tf_zpr.ZprZprPoliciesDataSource
		factories["oci_zpr_zpr_policy"] = tf_zpr.ZprZprPolicyDataSource
	}
	return factories
}
