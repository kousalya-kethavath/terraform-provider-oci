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
	tf_identity_data_plane "github.com/oracle/terraform-provider-oci/internal/service/identity_data_plane"
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
	tf_mysql "github.com/oracle/terraform-provider-oci/internal/service/mysql"
	tf_network_firewall "github.com/oracle/terraform-provider-oci/internal/service/network_firewall"
	tf_network_load_balancer "github.com/oracle/terraform-provider-oci/internal/service/network_load_balancer"
	tf_nosql "github.com/oracle/terraform-provider-oci/internal/service/nosql"
	tf_objectstorage "github.com/oracle/terraform-provider-oci/internal/service/objectstorage"
	tf_oce "github.com/oracle/terraform-provider-oci/internal/service/oce"
	tf_ocvp "github.com/oracle/terraform-provider-oci/internal/service/ocvp"
	tf_oda "github.com/oracle/terraform-provider-oci/internal/service/oda"
	tf_ons "github.com/oracle/terraform-provider-oci/internal/service/ons"
	tf_opa "github.com/oracle/terraform-provider-oci/internal/service/opa"
	tf_opensearch "github.com/oracle/terraform-provider-oci/internal/service/opensearch"
	tf_operator_access_control "github.com/oracle/terraform-provider-oci/internal/service/operator_access_control"
	tf_opsi "github.com/oracle/terraform-provider-oci/internal/service/opsi"
	tf_optimizer "github.com/oracle/terraform-provider-oci/internal/service/optimizer"
	tf_os_management_hub "github.com/oracle/terraform-provider-oci/internal/service/os_management_hub"
	tf_osp_gateway "github.com/oracle/terraform-provider-oci/internal/service/osp_gateway"
	tf_psa "github.com/oracle/terraform-provider-oci/internal/service/psa"
	tf_psql "github.com/oracle/terraform-provider-oci/internal/service/psql"
	tf_queue "github.com/oracle/terraform-provider-oci/internal/service/queue"
	tf_recovery "github.com/oracle/terraform-provider-oci/internal/service/recovery"
	tf_redis "github.com/oracle/terraform-provider-oci/internal/service/redis"
	tf_resource_analytics "github.com/oracle/terraform-provider-oci/internal/service/resource_analytics"
	tf_resource_scheduler "github.com/oracle/terraform-provider-oci/internal/service/resource_scheduler"
	tf_resourcemanager "github.com/oracle/terraform-provider-oci/internal/service/resourcemanager"
	tf_sch "github.com/oracle/terraform-provider-oci/internal/service/sch"
	tf_security_attribute "github.com/oracle/terraform-provider-oci/internal/service/security_attribute"
	tf_self "github.com/oracle/terraform-provider-oci/internal/service/self"
	tf_service_catalog "github.com/oracle/terraform-provider-oci/internal/service/service_catalog"
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

// generatedSDKv2ResourceFactories returns the provider-owned resource factory
// inventory. It intentionally does not mutate tfresource's process-global
// registration state.
func generatedSDKv2ResourceFactories() map[string]func() *schema.Resource {
	factories := make(map[string]func() *schema.Resource)
	if common.CheckForEnabledServices("adm") {
		factories["oci_adm_knowledge_base"] = tf_adm.AdmKnowledgeBaseResource
		factories["oci_adm_remediation_recipe"] = tf_adm.AdmRemediationRecipeResource
		factories["oci_adm_remediation_run"] = tf_adm.AdmRemediationRunResource
		factories["oci_adm_vulnerability_audit"] = tf_adm.AdmVulnerabilityAuditResource
	}
	if common.CheckForEnabledServices("aidataplatform") {
		factories["oci_ai_data_platform_ai_data_platform"] = tf_ai_data_platform.AiDataPlatformAiDataPlatformResource
	}
	if common.CheckForEnabledServices("aidocument") {
		factories["oci_ai_document_model"] = tf_ai_document.AiDocumentModelResource
		factories["oci_ai_document_processor_job"] = tf_ai_document.AiDocumentProcessorJobResource
		factories["oci_ai_document_project"] = tf_ai_document.AiDocumentProjectResource
	}
	if common.CheckForEnabledServices("ailanguage") {
		factories["oci_ai_language_endpoint"] = tf_ai_language.AiLanguageEndpointResource
		factories["oci_ai_language_job"] = tf_ai_language.AiLanguageJobResource
		factories["oci_ai_language_model"] = tf_ai_language.AiLanguageModelResource
		factories["oci_ai_language_project"] = tf_ai_language.AiLanguageProjectResource
	}
	if common.CheckForEnabledServices("aivision") {
		factories["oci_ai_vision_model"] = tf_ai_vision.AiVisionModelResource
		factories["oci_ai_vision_project"] = tf_ai_vision.AiVisionProjectResource
		factories["oci_ai_vision_stream_group"] = tf_ai_vision.AiVisionStreamGroupResource
		factories["oci_ai_vision_stream_job"] = tf_ai_vision.AiVisionStreamJobResource
		factories["oci_ai_vision_stream_source"] = tf_ai_vision.AiVisionStreamSourceResource
		factories["oci_ai_vision_vision_private_endpoint"] = tf_ai_vision.AiVisionVisionPrivateEndpointResource
	}
	if common.CheckForEnabledServices("analytics") {
		factories["oci_analytics_analytics_instance"] = tf_analytics.AnalyticsAnalyticsInstanceResource
		factories["oci_analytics_analytics_instance_private_access_channel"] = tf_analytics.AnalyticsAnalyticsInstancePrivateAccessChannelResource
		factories["oci_analytics_analytics_instance_resource_group"] = tf_analytics.AnalyticsAnalyticsInstanceResourceGroupResource
		factories["oci_analytics_analytics_instance_vanity_url"] = tf_analytics.AnalyticsAnalyticsInstanceVanityUrlResource
	}
	if common.CheckForEnabledServices("announcementsservice") {
		factories["oci_announcements_service_announcement_subscription"] = tf_announcements_service.AnnouncementsServiceAnnouncementSubscriptionResource
		factories["oci_announcements_service_announcement_subscriptions_actions_change_compartment"] = tf_announcements_service.AnnouncementsServiceAnnouncementSubscriptionsActionsChangeCompartmentResource
		factories["oci_announcements_service_announcement_subscriptions_filter_group"] = tf_announcements_service.AnnouncementsServiceAnnouncementSubscriptionsFilterGroupResource
	}
	if common.CheckForEnabledServices("apiplatform") {
		factories["oci_api_platform_api_platform_instance"] = tf_api_platform.ApiPlatformApiPlatformInstanceResource
	}
	if common.CheckForEnabledServices("apiaccesscontrol") {
		factories["oci_apiaccesscontrol_privileged_api_control"] = tf_apiaccesscontrol.ApiaccesscontrolPrivilegedApiControlResource
		factories["oci_apiaccesscontrol_privileged_api_request"] = tf_apiaccesscontrol.ApiaccesscontrolPrivilegedApiRequestResource
	}
	if common.CheckForEnabledServices("apigateway") {
		factories["oci_apigateway_api"] = tf_apigateway.ApigatewayApiResource
		factories["oci_apigateway_certificate"] = tf_apigateway.ApigatewayCertificateResource
		factories["oci_apigateway_deployment"] = tf_apigateway.ApigatewayDeploymentResource
		factories["oci_apigateway_gateway"] = tf_apigateway.ApigatewayGatewayResource
		factories["oci_apigateway_subscriber"] = tf_apigateway.ApigatewaySubscriberResource
		factories["oci_apigateway_usage_plan"] = tf_apigateway.ApigatewayUsagePlanResource
	}
	if common.CheckForEnabledServices("apm") {
		factories["oci_apm_apm_domain"] = tf_apm.ApmApmDomainResource
	}
	if common.CheckForEnabledServices("apmconfig") {
		factories["oci_apm_config_config"] = tf_apm_config.ApmConfigConfigResource
		factories["oci_apm_config_data_file"] = tf_apm_config.ApmConfigDataFileResource
	}
	if common.CheckForEnabledServices("apmsynthetics") {
		factories["oci_apm_synthetics_dedicated_vantage_point"] = tf_apm_synthetics.ApmSyntheticsDedicatedVantagePointResource
		factories["oci_apm_synthetics_monitor"] = tf_apm_synthetics.ApmSyntheticsMonitorResource
		factories["oci_apm_synthetics_on_premise_vantage_point"] = tf_apm_synthetics.ApmSyntheticsOnPremiseVantagePointResource
		factories["oci_apm_synthetics_on_premise_vantage_point_worker"] = tf_apm_synthetics.ApmSyntheticsOnPremiseVantagePointWorkerResource
		factories["oci_apm_synthetics_script"] = tf_apm_synthetics.ApmSyntheticsScriptResource
	}
	if common.CheckForEnabledServices("apmtraces") {
		factories["oci_apm_traces_scheduled_query"] = tf_apm_traces.ApmTracesScheduledQueryResource
	}
	if common.CheckForEnabledServices("appmgmtcontrol") {
		factories["oci_appmgmt_control_monitor_plugin_management"] = tf_appmgmt_control.AppmgmtControlMonitorPluginManagementResource
	}
	if common.CheckForEnabledServices("artifacts") {
		factories["oci_artifacts_container_configuration"] = tf_artifacts.ArtifactsContainerConfigurationResource
		factories["oci_artifacts_container_image_signature"] = tf_artifacts.ArtifactsContainerImageSignatureResource
		factories["oci_artifacts_container_repository"] = tf_artifacts.ArtifactsContainerRepositoryResource
		factories["oci_artifacts_generic_artifact"] = tf_artifacts.ArtifactsGenericArtifactResource
		factories["oci_artifacts_repository"] = tf_artifacts.ArtifactsRepositoryResource
	}
	if common.CheckForEnabledServices("audit") {
		factories["oci_audit_configuration"] = tf_audit.AuditConfigurationResource
	}
	if common.CheckForEnabledServices("autoscaling") {
		factories["oci_autoscaling_auto_scaling_configuration"] = tf_autoscaling.AutoScalingAutoScalingConfigurationResource
	}
	if common.CheckForEnabledServices("bastion") {
		factories["oci_bastion_bastion"] = tf_bastion.BastionBastionResource
		factories["oci_bastion_session"] = tf_bastion.BastionSessionResource
	}
	if common.CheckForEnabledServices("batch") {
		factories["oci_batch_batch_context"] = tf_batch.BatchBatchContextResource
		factories["oci_batch_batch_job_pool"] = tf_batch.BatchBatchJobPoolResource
		factories["oci_batch_batch_task_environment"] = tf_batch.BatchBatchTaskEnvironmentResource
		factories["oci_batch_batch_task_profile"] = tf_batch.BatchBatchTaskProfileResource
	}
	if common.CheckForEnabledServices("bds") {
		factories["oci_bds_auto_scaling_configuration"] = tf_bds.BdsAutoScalingConfigurationResource
		factories["oci_bds_bds_capacity_report"] = tf_bds.BdsBdsCapacityReportResource
		factories["oci_bds_bds_capacity_reservation"] = tf_bds.BdsBdsCapacityReservationResource
		factories["oci_bds_bds_cluster_admin_password_reset_action"] = tf_bds.BdsBdsClusterAdminPasswordResetActionResource
		factories["oci_bds_bds_instance"] = tf_bds.BdsBdsInstanceResource
		factories["oci_bds_bds_instance_api_key"] = tf_bds.BdsBdsInstanceApiKeyResource
		factories["oci_bds_bds_instance_bds_capacity_reservation_configuration"] = tf_bds.BdsBdsInstanceBdsCapacityReservationConfigurationResource
		factories["oci_bds_bds_instance_bds_certificate_configuration"] = tf_bds.BdsBdsInstanceBdsCertificateConfigurationResource
		factories["oci_bds_bds_instance_execute_bootstrap_script_action"] = tf_bds.BdsBdsInstanceExecuteBootstrapScriptActionResource
		factories["oci_bds_bds_instance_identity_configuration"] = tf_bds.BdsBdsInstanceIdentityConfigurationResource
		factories["oci_bds_bds_instance_metastore_config"] = tf_bds.BdsBdsInstanceMetastoreConfigResource
		factories["oci_bds_bds_instance_node_backup"] = tf_bds.BdsBdsInstanceNodeBackupResource
		factories["oci_bds_bds_instance_node_backup_configuration"] = tf_bds.BdsBdsInstanceNodeBackupConfigurationResource
		factories["oci_bds_bds_instance_node_replace_configuration"] = tf_bds.BdsBdsInstanceNodeReplaceConfigurationResource
		factories["oci_bds_bds_instance_operation_certificate_managements_management"] = tf_bds.BdsBdsInstanceOperationCertificateManagementsManagementResource
		factories["oci_bds_bds_instance_os_patch_action"] = tf_bds.BdsBdsInstanceOSPatchActionResource
		factories["oci_bds_bds_instance_patch_action"] = tf_bds.BdsBdsInstancePatchActionResource
		factories["oci_bds_bds_instance_replace_node_action"] = tf_bds.BdsBdsInstanceReplaceNodeActionResource
		factories["oci_bds_bds_instance_resource_principal_configuration"] = tf_bds.BdsBdsInstanceResourcePrincipalConfigurationResource
		factories["oci_bds_bds_instance_software_update_action"] = tf_bds.BdsBdsInstanceSoftwareUpdateResource
	}
	if common.CheckForEnabledServices("blockchain") {
		factories["oci_blockchain_blockchain_platform"] = tf_blockchain.BlockchainBlockchainPlatformResource
		factories["oci_blockchain_osn"] = tf_blockchain.BlockchainOsnResource
		factories["oci_blockchain_peer"] = tf_blockchain.BlockchainPeerResource
	}
	if common.CheckForEnabledServices("budget") {
		factories["oci_budget_alert_rule"] = tf_budget.BudgetAlertRuleResource
		factories["oci_budget_budget"] = tf_budget.BudgetBudgetResource
		factories["oci_budget_cost_alert_subscription"] = tf_budget.BudgetCostAlertSubscriptionResource
		factories["oci_budget_cost_anomaly_event"] = tf_budget.BudgetCostAnomalyEventResource
		factories["oci_budget_cost_anomaly_monitor"] = tf_budget.BudgetCostAnomalyMonitorResource
		factories["oci_budget_cost_anomaly_monitor_costanomalymonitorenabletoggles_management"] = tf_budget.BudgetCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource
	}
	if common.CheckForEnabledServices("capacitymanagement") {
		factories["oci_capacity_management_internal_occm_demand_signal"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalResource
		factories["oci_capacity_management_internal_occm_demand_signal_delivery"] = tf_capacity_management.CapacityManagementInternalOccmDemandSignalDeliveryResource
		factories["oci_capacity_management_occ_availability_catalog"] = tf_capacity_management.CapacityManagementOccAvailabilityCatalogResource
		factories["oci_capacity_management_occ_capacity_request"] = tf_capacity_management.CapacityManagementOccCapacityRequestResource
		factories["oci_capacity_management_occ_customer_group"] = tf_capacity_management.CapacityManagementOccCustomerGroupResource
		factories["oci_capacity_management_occ_customer_group_occ_customer"] = tf_capacity_management.CapacityManagementOccCustomerGroupOccCustomerResource
		factories["oci_capacity_management_occm_demand_signal"] = tf_capacity_management.CapacityManagementOccmDemandSignalResource
		factories["oci_capacity_management_occm_demand_signal_item"] = tf_capacity_management.CapacityManagementOccmDemandSignalItemResource
	}
	if common.CheckForEnabledServices("certificatesmanagement") {
		factories["oci_certificates_management_ca_bundle"] = tf_certificates_management.CertificatesManagementCaBundleResource
		factories["oci_certificates_management_certificate"] = tf_certificates_management.CertificatesManagementCertificateResource
		factories["oci_certificates_management_certificate_authority"] = tf_certificates_management.CertificatesManagementCertificateAuthorityResource
	}
	if common.CheckForEnabledServices("cloudbridge") {
		factories["oci_cloud_bridge_agent"] = tf_cloud_bridge.CloudBridgeAgentResource
		factories["oci_cloud_bridge_agent_dependency"] = tf_cloud_bridge.CloudBridgeAgentDependencyResource
		factories["oci_cloud_bridge_agent_plugin"] = tf_cloud_bridge.CloudBridgeAgentPluginResource
		factories["oci_cloud_bridge_asset"] = tf_cloud_bridge.CloudBridgeAssetResource
		factories["oci_cloud_bridge_asset_source"] = tf_cloud_bridge.CloudBridgeAssetSourceResource
		factories["oci_cloud_bridge_discovery_schedule"] = tf_cloud_bridge.CloudBridgeDiscoveryScheduleResource
		factories["oci_cloud_bridge_environment"] = tf_cloud_bridge.CloudBridgeEnvironmentResource
		factories["oci_cloud_bridge_inventory"] = tf_cloud_bridge.CloudBridgeInventoryResource
	}
	if common.CheckForEnabledServices("cloudguard") {
		factories["oci_cloud_guard_adhoc_query"] = tf_cloud_guard.CloudGuardAdhocQueryResource
		factories["oci_cloud_guard_cloud_guard_configuration"] = tf_cloud_guard.CloudGuardCloudGuardConfigurationResource
		factories["oci_cloud_guard_data_mask_rule"] = tf_cloud_guard.CloudGuardDataMaskRuleResource
		factories["oci_cloud_guard_data_source"] = tf_cloud_guard.CloudGuardDataSourceResource
		factories["oci_cloud_guard_detector_recipe"] = tf_cloud_guard.CloudGuardDetectorRecipeResource
		factories["oci_cloud_guard_managed_list"] = tf_cloud_guard.CloudGuardManagedListResource
		factories["oci_cloud_guard_responder_recipe"] = tf_cloud_guard.CloudGuardResponderRecipeResource
		factories["oci_cloud_guard_saved_query"] = tf_cloud_guard.CloudGuardSavedQueryResource
		factories["oci_cloud_guard_security_recipe"] = tf_cloud_guard.CloudGuardSecurityRecipeResource
		factories["oci_cloud_guard_security_zone"] = tf_cloud_guard.CloudGuardSecurityZoneResource
		factories["oci_cloud_guard_target"] = tf_cloud_guard.CloudGuardTargetResource
		factories["oci_cloud_guard_wlp_agent"] = tf_cloud_guard.CloudGuardWlpAgentResource
	}
	if common.CheckForEnabledServices("cloudmigrations") {
		factories["oci_cloud_migrations_migration"] = tf_cloud_migrations.CloudMigrationsMigrationResource
		factories["oci_cloud_migrations_migration_asset"] = tf_cloud_migrations.CloudMigrationsMigrationAssetResource
		factories["oci_cloud_migrations_migration_plan"] = tf_cloud_migrations.CloudMigrationsMigrationPlanResource
		factories["oci_cloud_migrations_replication_schedule"] = tf_cloud_migrations.CloudMigrationsReplicationScheduleResource
		factories["oci_cloud_migrations_target_asset"] = tf_cloud_migrations.CloudMigrationsTargetAssetResource
	}
	if common.CheckForEnabledServices("clusterhealth") {
		factories["oci_cluster_health_diagnosis_store"] = tf_cluster_health.ClusterHealthDiagnosisStoreResource
	}
	if common.CheckForEnabledServices("clusterplacementgroups") {
		factories["oci_cluster_placement_groups_cluster_placement_group"] = tf_cluster_placement_groups.ClusterPlacementGroupsClusterPlacementGroupResource
	}
	if common.CheckForEnabledServices("computecloudatcustomer") {
		factories["oci_compute_cloud_at_customer_ccc_infrastructure"] = tf_compute_cloud_at_customer.ComputeCloudAtCustomerCccInfrastructureResource
		factories["oci_compute_cloud_at_customer_ccc_upgrade_schedule"] = tf_compute_cloud_at_customer.ComputeCloudAtCustomerCccUpgradeScheduleResource
	}
	if common.CheckForEnabledServices("containerinstances") {
		factories["oci_container_instances_container_instance"] = tf_container_instances.ContainerInstancesContainerInstanceResource
	}
	if common.CheckForEnabledServices("containerengine") {
		factories["oci_containerengine_addon"] = tf_containerengine.ContainerengineAddonResource
		factories["oci_containerengine_cluster"] = tf_containerengine.ContainerengineClusterResource
		factories["oci_containerengine_cluster_complete_credential_rotation_management"] = tf_containerengine.ContainerengineClusterCompleteCredentialRotationManagementResource
		factories["oci_containerengine_cluster_public_api_endpoint_decommission_manager"] = tf_containerengine.ContainerengineClusterPublicApiEndpointDecommissionManagerResource
		factories["oci_containerengine_cluster_start_credential_rotation_management"] = tf_containerengine.ContainerengineClusterStartCredentialRotationManagementResource
		factories["oci_containerengine_cluster_workload_mapping"] = tf_containerengine.ContainerengineClusterWorkloadMappingResource
		factories["oci_containerengine_node_pool"] = tf_containerengine.ContainerengineNodePoolResource
		factories["oci_containerengine_virtual_node_pool"] = tf_containerengine.ContainerengineVirtualNodePoolResource
	}
	if common.CheckForEnabledServices("core") {
		factories["oci_core_app_catalog_listing_resource_version_agreement"] = tf_core.AppCatalogListingResourceVersionAgreementResource
		factories["oci_core_app_catalog_subscription"] = tf_core.CoreAppCatalogSubscriptionResource
		factories["oci_core_boot_volume"] = tf_core.CoreBootVolumeResource
		factories["oci_core_boot_volume_backup"] = tf_core.CoreBootVolumeBackupResource
		factories["oci_core_byoasn"] = tf_core.CoreByoasnResource
		factories["oci_core_capture_filter"] = tf_core.CoreCaptureFilterResource
		factories["oci_core_cluster_network"] = tf_core.CoreClusterNetworkResource
		factories["oci_core_compute_capacity_report"] = tf_core.CoreComputeCapacityReportResource
		factories["oci_core_compute_capacity_reservation"] = tf_core.CoreComputeCapacityReservationResource
		factories["oci_core_compute_capacity_topology"] = tf_core.CoreComputeCapacityTopologyResource
		factories["oci_core_compute_cluster"] = tf_core.CoreComputeClusterResource
		factories["oci_core_compute_gpu_memory_cluster"] = tf_core.CoreComputeGpuMemoryClusterResource
		factories["oci_core_compute_gpu_memory_fabric"] = tf_core.CoreComputeGpuMemoryFabricResource
		factories["oci_core_compute_host"] = tf_core.CoreComputeHostResource
		factories["oci_core_compute_host_group"] = tf_core.CoreComputeHostGroupResource
		factories["oci_core_compute_image_capability_schema"] = tf_core.CoreComputeImageCapabilitySchemaResource
		factories["oci_core_console_history"] = tf_core.CoreConsoleHistoryResource
		factories["oci_core_cpe"] = tf_core.CoreCpeResource
		factories["oci_core_cross_connect"] = tf_core.CoreCrossConnectResource
		factories["oci_core_cross_connect_group"] = tf_core.CoreCrossConnectGroupResource
		factories["oci_core_dedicated_vm_host"] = tf_core.CoreDedicatedVmHostResource
		factories["oci_core_default_dhcp_options"] = tf_core.DefaultCoreDhcpOptionsResource
		factories["oci_core_default_drg_route_table"] = tf_core.DefaultCoreDrgRouteTableResource
		factories["oci_core_default_route_table"] = tf_core.DefaultCoreRouteTableResource
		factories["oci_core_default_security_list"] = tf_core.CoreDefaultSecurityListResource
		factories["oci_core_dhcp_options"] = tf_core.CoreDhcpOptionsResource
		factories["oci_core_drg"] = tf_core.CoreDrgResource
		factories["oci_core_drg_attachment"] = tf_core.CoreDrgAttachmentResource
		factories["oci_core_drg_attachment_management"] = tf_core.CoreDrgAttachmentManagementResource
		factories["oci_core_drg_attachments_list"] = tf_core.CoreDrgAttachmentsListResource
		factories["oci_core_drg_route_distribution"] = tf_core.CoreDrgRouteDistributionResource
		factories["oci_core_drg_route_distribution_statement"] = tf_core.CoreDrgRouteDistributionStatementResource
		factories["oci_core_drg_route_table"] = tf_core.CoreDrgRouteTableResource
		factories["oci_core_drg_route_table_route_rule"] = tf_core.CoreDrgRouteTableRouteRuleResource
		factories["oci_core_image"] = tf_core.CoreImageResource
		factories["oci_core_instance"] = tf_core.CoreInstanceResource
		factories["oci_core_instance_configuration"] = tf_core.CoreInstanceConfigurationResource
		factories["oci_core_instance_console_connection"] = tf_core.CoreInstanceConsoleConnectionResource
		factories["oci_core_instance_maintenance_event"] = tf_core.CoreInstanceMaintenanceEventResource
		factories["oci_core_instance_pool"] = tf_core.CoreInstancePoolResource
		factories["oci_core_instance_pool_instance"] = tf_core.CoreInstancePoolInstanceResource
		factories["oci_core_internet_gateway"] = tf_core.CoreInternetGatewayResource
		factories["oci_core_ipsec"] = tf_core.CoreIpSecConnectionResource
		factories["oci_core_ipsec_connection_tunnel_management"] = tf_core.CoreIpSecConnectionTunnelManagementResource
		factories["oci_core_ipv6"] = tf_core.CoreIpv6Resource
		factories["oci_core_listing_resource_version_agreement"] = tf_core.AppCatalogListingResourceVersionAgreementResource
		factories["oci_core_local_peering_gateway"] = tf_core.CoreLocalPeeringGatewayResource
		factories["oci_core_nat_gateway"] = tf_core.CoreNatGatewayResource
		factories["oci_core_network_security_group"] = tf_core.CoreNetworkSecurityGroupResource
		factories["oci_core_network_security_group_security_rule"] = tf_core.CoreNetworkSecurityGroupSecurityRuleResource
		factories["oci_core_private_ip"] = tf_core.CorePrivateIpResource
		factories["oci_core_public_ip"] = tf_core.CorePublicIpResource
		factories["oci_core_public_ip_pool"] = tf_core.CorePublicIpPoolResource
		factories["oci_core_public_ip_pool_capacity"] = tf_core.PublicIpPoolCapacityResource
		factories["oci_core_remote_peering_connection"] = tf_core.CoreRemotePeeringConnectionResource
		factories["oci_core_route_table"] = tf_core.CoreRouteTableResource
		factories["oci_core_route_table_attachment"] = tf_core.CoreRouteTableAttachmentResource
		factories["oci_core_security_list"] = tf_core.CoreSecurityListResource
		factories["oci_core_service_gateway"] = tf_core.CoreServiceGatewayResource
		factories["oci_core_shape_management"] = tf_core.CoreShapeResource
		factories["oci_core_subnet"] = tf_core.CoreSubnetResource
		factories["oci_core_vcn"] = tf_core.CoreVcnResource
		factories["oci_core_virtual_circuit"] = tf_core.CoreVirtualCircuitResource
		factories["oci_core_vlan"] = tf_core.CoreVlanResource
		factories["oci_core_vnic_attachment"] = tf_core.CoreVnicAttachmentResource
		factories["oci_core_volume"] = tf_core.CoreVolumeResource
		factories["oci_core_volume_attachment"] = tf_core.CoreVolumeAttachmentResource
		factories["oci_core_volume_backup"] = tf_core.CoreVolumeBackupResource
		factories["oci_core_volume_backup_policy"] = tf_core.CoreVolumeBackupPolicyResource
		factories["oci_core_volume_backup_policy_assignment"] = tf_core.CoreVolumeBackupPolicyAssignmentResource
		factories["oci_core_volume_group"] = tf_core.CoreVolumeGroupResource
		factories["oci_core_volume_group_backup"] = tf_core.CoreVolumeGroupBackupResource
		factories["oci_core_vtap"] = tf_core.CoreVtapResource
	}
	if common.CheckForEnabledServices("costad") {
		factories["oci_costad_cost_alert_subscription"] = tf_costad.CostadCostAlertSubscriptionResource
		factories["oci_costad_cost_anomaly_event"] = tf_costad.CostadCostAnomalyEventResource
		factories["oci_costad_cost_anomaly_monitor"] = tf_costad.CostadCostAnomalyMonitorResource
		factories["oci_costad_cost_anomaly_monitor_costanomalymonitorenabletoggles_management"] = tf_costad.CostadCostAnomalyMonitorCostanomalymonitorenabletogglesManagementResource
	}
	if common.CheckForEnabledServices("datalabelingservice") {
		factories["oci_data_labeling_service_dataset"] = tf_data_labeling_service.DataLabelingServiceDatasetResource
	}
	if common.CheckForEnabledServices("datasafe") {
		factories["oci_data_safe_add_sdm_columns"] = tf_data_safe.DataSafeAddColumnsFromSdmResource
		factories["oci_data_safe_alert"] = tf_data_safe.DataSafeAlertResource
		factories["oci_data_safe_alert_policy"] = tf_data_safe.DataSafeAlertPolicyResource
		factories["oci_data_safe_alert_policy_rule"] = tf_data_safe.DataSafeAlertPolicyRuleResource
		factories["oci_data_safe_attribute_set"] = tf_data_safe.DataSafeAttributeSetResource
		factories["oci_data_safe_audit_archive_retrieval"] = tf_data_safe.DataSafeAuditArchiveRetrievalResource
		factories["oci_data_safe_audit_policy"] = tf_data_safe.DataSafeAuditPolicyResource
		factories["oci_data_safe_audit_policy_management"] = tf_data_safe.DataSafeAuditPolicyManagementResource
		factories["oci_data_safe_audit_profile"] = tf_data_safe.DataSafeAuditProfileResource
		factories["oci_data_safe_audit_profile_management"] = tf_data_safe.DataSafeAuditProfileManagementResource
		factories["oci_data_safe_audit_trail"] = tf_data_safe.DataSafeAuditTrailResource
		factories["oci_data_safe_audit_trail_management"] = tf_data_safe.DataSafeAuditTrailManagementResource
		factories["oci_data_safe_calculate_audit_volume_available"] = tf_data_safe.DataSafeCalculateAuditVolumeAvailableResource
		factories["oci_data_safe_calculate_audit_volume_collected"] = tf_data_safe.DataSafeCalculateAuditVolumeCollectedResource
		factories["oci_data_safe_compare_security_assessment"] = tf_data_safe.DataSafeCompareSecurityAssessmentResource
		factories["oci_data_safe_compare_user_assessment"] = tf_data_safe.DataSafeCompareUserAssessmentResource
		factories["oci_data_safe_crypto_assessment"] = tf_data_safe.DataSafeCryptoAssessmentResource
		factories["oci_data_safe_crypto_assessment_management"] = tf_data_safe.DataSafeCryptoAssessmentManagementResource
		factories["oci_data_safe_data_safe_configuration"] = tf_data_safe.DataSafeDataSafeConfigurationResource
		factories["oci_data_safe_data_safe_private_endpoint"] = tf_data_safe.DataSafeDataSafePrivateEndpointResource
		factories["oci_data_safe_database_security_config"] = tf_data_safe.DataSafeDatabaseSecurityConfigResource
		factories["oci_data_safe_database_security_config_management"] = tf_data_safe.DataSafeDatabaseSecurityConfigManagementResource
		factories["oci_data_safe_discovery_job"] = tf_data_safe.DataSafeDiscoveryJobResource
		factories["oci_data_safe_discovery_jobs_result"] = tf_data_safe.DataSafeDiscoveryJobsResultResource
		factories["oci_data_safe_generate_on_prem_connector_configuration"] = tf_data_safe.DataSafeGenerateOnPremConnectorConfigurationResource
		factories["oci_data_safe_library_masking_format"] = tf_data_safe.DataSafeLibraryMaskingFormatResource
		factories["oci_data_safe_mask_data"] = tf_data_safe.DataSafeMaskDataResource
		factories["oci_data_safe_masking_policies_apply_difference_to_masking_columns"] = tf_data_safe.DataSafeMaskingPolicyApplyDifferenceToMaskingColumnsResource
		factories["oci_data_safe_masking_policies_masking_column"] = tf_data_safe.DataSafeMaskingPoliciesMaskingColumnResource
		factories["oci_data_safe_masking_policy"] = tf_data_safe.DataSafeMaskingPolicyResource
		factories["oci_data_safe_masking_policy_health_report_management"] = tf_data_safe.DataSafeMaskingPolicyHealthReportManagementResource
		factories["oci_data_safe_masking_report_management"] = tf_data_safe.DataSafeMaskingReportManagementResource
		factories["oci_data_safe_on_prem_connector"] = tf_data_safe.DataSafeOnPremConnectorResource
		factories["oci_data_safe_report"] = tf_data_safe.DataSafeReportResource
		factories["oci_data_safe_report_definition"] = tf_data_safe.DataSafeReportDefinitionResource
		factories["oci_data_safe_sdm_masking_policy_difference"] = tf_data_safe.DataSafeSdmMaskingPolicyDifferenceResource
		factories["oci_data_safe_security_assessment"] = tf_data_safe.DataSafeSecurityAssessmentResource
		factories["oci_data_safe_security_assessment_check"] = tf_data_safe.DataSafeSecurityAssessmentCheckResource
		factories["oci_data_safe_security_assessment_finding"] = tf_data_safe.DataSafeSecurityAssessmentFindingResource
		factories["oci_data_safe_security_policy"] = tf_data_safe.DataSafeSecurityPolicyResource
		factories["oci_data_safe_security_policy_config"] = tf_data_safe.DataSafeSecurityPolicyConfigResource
		factories["oci_data_safe_security_policy_deployment"] = tf_data_safe.DataSafeSecurityPolicyDeploymentResource
		factories["oci_data_safe_security_policy_deployment_management"] = tf_data_safe.DataSafeSecurityPolicyDeploymentManagementResource
		factories["oci_data_safe_security_policy_management"] = tf_data_safe.DataSafeSecurityPolicyManagementResource
		factories["oci_data_safe_sensitive_data_model"] = tf_data_safe.DataSafeSensitiveDataModelResource
		factories["oci_data_safe_sensitive_data_model_referential_relation"] = tf_data_safe.DataSafeSensitiveDataModelReferentialRelationResource
		factories["oci_data_safe_sensitive_data_models_apply_discovery_job_results"] = tf_data_safe.DataSafeSensitiveDataModelsApplyDiscoveryJobResultsResource
		factories["oci_data_safe_sensitive_data_models_sensitive_column"] = tf_data_safe.DataSafeSensitiveDataModelsSensitiveColumnResource
		factories["oci_data_safe_sensitive_type"] = tf_data_safe.DataSafeSensitiveTypeResource
		factories["oci_data_safe_sensitive_type_group"] = tf_data_safe.DataSafeSensitiveTypeGroupResource
		factories["oci_data_safe_sensitive_type_group_grouped_sensitive_type"] = tf_data_safe.DataSafeSensitiveTypeGroupGroupedSensitiveTypeResource
		factories["oci_data_safe_sensitive_types_export"] = tf_data_safe.DataSafeSensitiveTypesExportResource
		factories["oci_data_safe_set_security_assessment_baseline"] = tf_data_safe.DataSafeSetSecurityAssessmentBaselineResource
		factories["oci_data_safe_set_security_assessment_baseline_management"] = tf_data_safe.DataSafeSetSecurityAssessmentBaselineManagementResource
		factories["oci_data_safe_set_user_assessment_baseline"] = tf_data_safe.DataSafeSetUserAssessmentBaselineResource
		factories["oci_data_safe_set_user_assessment_baseline_management"] = tf_data_safe.DataSafeSetUserAssessmentBaselineManagementResource
		factories["oci_data_safe_sql_collection"] = tf_data_safe.DataSafeSqlCollectionResource
		factories["oci_data_safe_sql_firewall_policy"] = tf_data_safe.DataSafeSqlFirewallPolicyResource
		factories["oci_data_safe_sql_firewall_policy_management"] = tf_data_safe.DataSafeSqlFirewallPolicyManagementResource
		factories["oci_data_safe_target_alert_policy_association"] = tf_data_safe.DataSafeTargetAlertPolicyAssociationResource
		factories["oci_data_safe_target_database"] = tf_data_safe.DataSafeTargetDatabaseResource
		factories["oci_data_safe_target_database_group"] = tf_data_safe.DataSafeTargetDatabaseGroupResource
		factories["oci_data_safe_target_database_peer_target_database"] = tf_data_safe.DataSafeTargetDatabasePeerTargetDatabaseResource
		factories["oci_data_safe_unified_audit_policy"] = tf_data_safe.DataSafeUnifiedAuditPolicyResource
		factories["oci_data_safe_unified_audit_policy_definition"] = tf_data_safe.DataSafeUnifiedAuditPolicyDefinitionResource
		factories["oci_data_safe_unset_security_assessment_baseline"] = tf_data_safe.DataSafeUnsetSecurityAssessmentBaselineResource
		factories["oci_data_safe_unset_security_assessment_baseline_management"] = tf_data_safe.DataSafeUnsetSecurityAssessmentBaselineManagementResource
		factories["oci_data_safe_unset_user_assessment_baseline"] = tf_data_safe.DataSafeUnsetUserAssessmentBaselineResource
		factories["oci_data_safe_unset_user_assessment_baseline_management"] = tf_data_safe.DataSafeUnsetUserAssessmentBaselineManagementResource
		factories["oci_data_safe_user_assessment"] = tf_data_safe.DataSafeUserAssessmentResource
	}
	if common.CheckForEnabledServices("database") {
		factories["oci_database_advanced_cluster_file_system"] = tf_database.DatabaseAdvancedClusterFileSystemResource
		factories["oci_database_advanced_cluster_file_system_mount"] = tf_database.DatabaseAdvancedClusterFileSystemMountResource
		factories["oci_database_advanced_cluster_file_system_unmount"] = tf_database.DatabaseAdvancedClusterFileSystemUnmountResource
		factories["oci_database_application_vip"] = tf_database.DatabaseApplicationVipResource
		factories["oci_database_autonomous_container_database"] = tf_database.DatabaseAutonomousContainerDatabaseResource
		factories["oci_database_autonomous_container_database_add_standby"] = tf_database.DatabaseAutonomousContainerDatabaseAddStandbyResource
		factories["oci_database_autonomous_container_database_dataguard_association"] = tf_database.DatabaseAutonomousContainerDatabaseDataguardAssociationResource
		factories["oci_database_autonomous_container_database_dataguard_association_operation"] = tf_database.DatabaseAutonomousContainerDatabaseDataguardAssociationOperationResource
		factories["oci_database_autonomous_container_database_dataguard_role_change"] = tf_database.DatabaseAutonomousContainerDatabaseDataguardRoleChangeResource
		factories["oci_database_autonomous_container_database_snapshot_standby"] = tf_database.DatabaseAutonomousContainerDatabaseSnapshotStandbyResource
		factories["oci_database_autonomous_database"] = tf_database.DatabaseAutonomousDatabaseResource
		factories["oci_database_autonomous_database_backup"] = tf_database.DatabaseAutonomousDatabaseBackupResource
		factories["oci_database_autonomous_database_instance_wallet_management"] = tf_database.DatabaseAutonomousDatabaseInstanceWalletManagementResource
		factories["oci_database_autonomous_database_regional_wallet_management"] = tf_database.DatabaseAutonomousDatabaseRegionalWalletManagementResource
		factories["oci_database_autonomous_database_saas_admin_user"] = tf_database.DatabaseAutonomousDatabaseSaasAdminUserResource
		factories["oci_database_autonomous_database_software_image"] = tf_database.DatabaseAutonomousDatabaseSoftwareImageResource
		factories["oci_database_autonomous_database_wallet"] = tf_database.DatabaseAutonomousDatabaseWalletResource
		factories["oci_database_autonomous_exadata_infrastructure"] = tf_database.DatabaseAutonomousExadataInfrastructureResource
		factories["oci_database_autonomous_vm_cluster"] = tf_database.DatabaseAutonomousVmClusterResource
		factories["oci_database_autonomous_vm_cluster_ords_certificate_management"] = tf_database.DatabaseAutonomousVmClusterOrdsCertificateManagementResource
		factories["oci_database_autonomous_vm_cluster_ssl_certificate_management"] = tf_database.DatabaseAutonomousVmClusterSslCertificateManagementResource
		factories["oci_database_backup"] = tf_database.DatabaseBackupResource
		factories["oci_database_backup_cancel_management"] = tf_database.DatabaseBackupCancelManagementResource
		factories["oci_database_backup_destination"] = tf_database.DatabaseBackupDestinationResource
		factories["oci_database_cloud_autonomous_vm_cluster"] = tf_database.DatabaseCloudAutonomousVmClusterResource
		factories["oci_database_cloud_database_management"] = tf_database.DatabaseCloudDatabaseManagementResource
		factories["oci_database_cloud_exadata_infrastructure"] = tf_database.DatabaseCloudExadataInfrastructureResource
		factories["oci_database_cloud_exadata_infrastructure_configure_exascale_management"] = tf_database.DatabaseCloudExadataInfrastructureConfigureExascaleManagementResource
		factories["oci_database_cloud_vm_cluster"] = tf_database.DatabaseCloudVmClusterResource
		factories["oci_database_cloud_vm_cluster_iorm_config"] = tf_database.DatabaseCloudVmClusterIormConfigResource
		factories["oci_database_data_guard_association"] = tf_database.DatabaseDataGuardAssociationResource
		factories["oci_database_data_patch"] = tf_database.DatabaseDatabaseDataPatchResource
		factories["oci_database_database"] = tf_database.DatabaseDatabaseResource
		factories["oci_database_database_snapshot_standby"] = tf_database.DatabaseDatabaseSnapshotStandbyResource
		factories["oci_database_database_software_image"] = tf_database.DatabaseDatabaseSoftwareImageResource
		factories["oci_database_database_software_schedule_management"] = tf_database.DatabaseDatabaseSoftwareScheduleManagementResource
		factories["oci_database_database_upgrade"] = tf_database.DatabaseDatabaseUpgradeResource
		factories["oci_database_db_home"] = tf_database.DatabaseDbHomeResource
		factories["oci_database_db_node"] = tf_database.DatabaseDbNodeResource
		factories["oci_database_db_node_console_connection"] = tf_database.DatabaseDbNodeConsoleConnectionResource
		factories["oci_database_db_node_console_history"] = tf_database.DatabaseDbNodeConsoleHistoryResource
		factories["oci_database_db_node_snapshot"] = tf_database.DatabaseDbNodeSnapshotResource
		factories["oci_database_db_node_snapshot_management"] = tf_database.DatabaseDbNodeSnapshotManagementResource
		factories["oci_database_db_system"] = tf_database.DatabaseDbSystemResource
		factories["oci_database_db_systems_upgrade"] = tf_database.DatabaseDbSystemsUpgradeResource
		factories["oci_database_exadata_infrastructure"] = tf_database.DatabaseExadataInfrastructureResource
		factories["oci_database_exadata_infrastructure_compute"] = tf_database.DatabaseExadataInfrastructureComputeManagedResource
		factories["oci_database_exadata_infrastructure_configure_exascale_management"] = tf_database.DatabaseExadataInfrastructureConfigureExascaleManagementResource
		factories["oci_database_exadata_infrastructure_storage"] = tf_database.DatabaseExadataInfrastructureStorageResource
		factories["oci_database_exadata_iorm_config"] = tf_database.DatabaseExadataIormConfigResource
		factories["oci_database_exadb_vm_cluster"] = tf_database.DatabaseExadbVmClusterResource
		factories["oci_database_exascale_db_storage_vault"] = tf_database.DatabaseExascaleDbStorageVaultResource
		factories["oci_database_execution_action"] = tf_database.DatabaseExecutionActionResource
		factories["oci_database_execution_window"] = tf_database.DatabaseExecutionWindowResource
		factories["oci_database_external_container_database"] = tf_database.DatabaseExternalContainerDatabaseResource
		factories["oci_database_external_container_database_management"] = tf_database.DatabaseExternalContainerDatabaseManagementResource
		factories["oci_database_external_database_connector"] = tf_database.DatabaseExternalDatabaseConnectorResource
		factories["oci_database_external_non_container_database"] = tf_database.DatabaseExternalNonContainerDatabaseResource
		factories["oci_database_external_non_container_database_management"] = tf_database.DatabaseExternalNonContainerDatabaseManagementResource
		factories["oci_database_external_non_container_database_operations_insights_management"] = tf_database.DatabaseExternalNonContainerDatabaseOperationsInsightsManagementResource
		factories["oci_database_external_pluggable_database"] = tf_database.DatabaseExternalPluggableDatabaseResource
		factories["oci_database_external_pluggable_database_management"] = tf_database.DatabaseExternalPluggableDatabaseManagementResource
		factories["oci_database_external_pluggable_database_operations_insights_management"] = tf_database.DatabaseExternalPluggableDatabaseOperationsInsightsManagementResource
		factories["oci_database_externalcontainerdatabases_stack_monitoring"] = tf_database.DatabaseExternalcontainerdatabasesStackMonitoringResource
		factories["oci_database_externalnoncontainerdatabases_stack_monitoring"] = tf_database.DatabaseExternalnoncontainerdatabasesStackMonitoringResource
		factories["oci_database_externalpluggabledatabases_stack_monitoring"] = tf_database.DatabaseExternalpluggabledatabasesStackMonitoringResource
		factories["oci_database_key_store"] = tf_database.DatabaseKeyStoreResource
		factories["oci_database_maintenance_run"] = tf_database.DatabaseMaintenanceRunResource
		factories["oci_database_migration"] = tf_database.DatabaseMigrationResource
		factories["oci_database_oneoff_patch"] = tf_database.DatabaseOneoffPatchResource
		factories["oci_database_pluggable_database"] = tf_database.DatabasePluggableDatabaseResource
		factories["oci_database_pluggable_database_pluggabledatabasemanagements_management"] = tf_database.DatabasePluggableDatabasePluggabledatabasemanagementsManagementResource
		factories["oci_database_pluggable_database_snapshot"] = tf_database.DatabasePluggableDatabaseSnapshotResource
		factories["oci_database_pluggable_databases_local_clone"] = tf_database.DatabasePluggableDatabasesLocalCloneResource
		factories["oci_database_pluggable_databases_remote_clone"] = tf_database.DatabasePluggableDatabasesRemoteCloneResource
		factories["oci_database_scheduled_action"] = tf_database.DatabaseScheduledActionResource
		factories["oci_database_scheduling_plan"] = tf_database.DatabaseSchedulingPlanResource
		factories["oci_database_scheduling_policy"] = tf_database.DatabaseSchedulingPolicyResource
		factories["oci_database_scheduling_policy_scheduling_window"] = tf_database.DatabaseSchedulingPolicySchedulingWindowResource
		factories["oci_database_vm_cluster"] = tf_database.DatabaseVmClusterResource
		factories["oci_database_vm_cluster_add_virtual_machine"] = tf_database.DatabaseVmClusterAddVirtualMachineResource
		factories["oci_database_vm_cluster_network"] = tf_database.DatabaseVmClusterNetworkResource
		factories["oci_database_vm_cluster_remove_virtual_machine"] = tf_database.DatabaseVmClusterRemoveVirtualMachineResource
	}
	if common.CheckForEnabledServices("databasemanagement") {
		factories["oci_database_management_autonomous_database_autonomous_database_dbm_features_management"] = tf_database_management.DatabaseManagementAutonomousDatabaseAutonomousDatabaseDbmFeaturesManagementResource
		factories["oci_database_management_cloud_asm"] = tf_database_management.DatabaseManagementCloudAsmResource
		factories["oci_database_management_cloud_asm_instance"] = tf_database_management.DatabaseManagementCloudAsmInstanceResource
		factories["oci_database_management_cloud_cluster"] = tf_database_management.DatabaseManagementCloudClusterResource
		factories["oci_database_management_cloud_cluster_instance"] = tf_database_management.DatabaseManagementCloudClusterInstanceResource
		factories["oci_database_management_cloud_db_home"] = tf_database_management.DatabaseManagementCloudDbHomeResource
		factories["oci_database_management_cloud_db_node"] = tf_database_management.DatabaseManagementCloudDbNodeResource
		factories["oci_database_management_cloud_db_system"] = tf_database_management.DatabaseManagementCloudDbSystemResource
		factories["oci_database_management_cloud_db_system_cloud_database_managements_management"] = tf_database_management.DatabaseManagementCloudDbSystemCloudDatabaseManagementsManagementResource
		factories["oci_database_management_cloud_db_system_cloud_stack_monitorings_management"] = tf_database_management.DatabaseManagementCloudDbSystemCloudStackMonitoringsManagementResource
		factories["oci_database_management_cloud_db_system_connector"] = tf_database_management.DatabaseManagementCloudDbSystemConnectorResource
		factories["oci_database_management_cloud_db_system_discovery"] = tf_database_management.DatabaseManagementCloudDbSystemDiscoveryResource
		factories["oci_database_management_cloud_exadata_infrastructure"] = tf_database_management.DatabaseManagementCloudExadataInfrastructureResource
		factories["oci_database_management_cloud_exadata_infrastructure_managedexadata_management"] = tf_database_management.DatabaseManagementCloudExadataInfrastructureManagedexadataManagementResource
		factories["oci_database_management_cloud_exadata_storage_connector"] = tf_database_management.DatabaseManagementCloudExadataStorageConnectorResource
		factories["oci_database_management_cloud_exadata_storage_grid"] = tf_database_management.DatabaseManagementCloudExadataStorageGridResource
		factories["oci_database_management_cloud_exadata_storage_server"] = tf_database_management.DatabaseManagementCloudExadataStorageServerResource
		factories["oci_database_management_cloud_listener"] = tf_database_management.DatabaseManagementCloudListenerResource
		factories["oci_database_management_database_dbm_features_management"] = tf_database_management.DatabaseManagementDatabaseDbmFeaturesManagementResource
		factories["oci_database_management_db_management_private_endpoint"] = tf_database_management.DatabaseManagementDbManagementPrivateEndpointResource
		factories["oci_database_management_external_asm"] = tf_database_management.DatabaseManagementExternalAsmResource
		factories["oci_database_management_external_asm_instance"] = tf_database_management.DatabaseManagementExternalAsmInstanceResource
		factories["oci_database_management_external_cluster"] = tf_database_management.DatabaseManagementExternalClusterResource
		factories["oci_database_management_external_cluster_instance"] = tf_database_management.DatabaseManagementExternalClusterInstanceResource
		factories["oci_database_management_external_db_home"] = tf_database_management.DatabaseManagementExternalDbHomeResource
		factories["oci_database_management_external_db_node"] = tf_database_management.DatabaseManagementExternalDbNodeResource
		factories["oci_database_management_external_db_system"] = tf_database_management.DatabaseManagementExternalDbSystemResource
		factories["oci_database_management_external_db_system_connector"] = tf_database_management.DatabaseManagementExternalDbSystemConnectorResource
		factories["oci_database_management_external_db_system_database_managements_management"] = tf_database_management.DatabaseManagementExternalDbSystemDatabaseManagementsManagementResource
		factories["oci_database_management_external_db_system_discovery"] = tf_database_management.DatabaseManagementExternalDbSystemDiscoveryResource
		factories["oci_database_management_external_db_system_stack_monitorings_management"] = tf_database_management.DatabaseManagementExternalDbSystemStackMonitoringsManagementResource
		factories["oci_database_management_external_exadata_infrastructure"] = tf_database_management.DatabaseManagementExternalExadataInfrastructureResource
		factories["oci_database_management_external_exadata_infrastructure_exadata_management"] = tf_database_management.DatabaseManagementExternalExadataInfrastructureExadataManagementResource
		factories["oci_database_management_external_exadata_storage_connector"] = tf_database_management.DatabaseManagementExternalExadataStorageConnectorResource
		factories["oci_database_management_external_exadata_storage_grid"] = tf_database_management.DatabaseManagementExternalExadataStorageGridResource
		factories["oci_database_management_external_exadata_storage_server"] = tf_database_management.DatabaseManagementExternalExadataStorageServerResource
		factories["oci_database_management_external_listener"] = tf_database_management.DatabaseManagementExternalListenerResource
		factories["oci_database_management_external_my_sql_database"] = tf_database_management.DatabaseManagementExternalMySqlDatabaseResource
		factories["oci_database_management_external_my_sql_database_connector"] = tf_database_management.DatabaseManagementExternalMySqlDatabaseConnectorResource
		factories["oci_database_management_external_my_sql_database_external_mysql_databases_management"] = tf_database_management.DatabaseManagementExternalMySqlDatabaseExternalMysqlDatabasesManagementResource
		factories["oci_database_management_externalcontainerdatabase_external_container_dbm_features_management"] = tf_database_management.DatabaseManagementExternalcontainerdatabaseExternalContainerDbmFeaturesManagementResource
		factories["oci_database_management_externalnoncontainerdatabase_external_non_container_dbm_features_management"] = tf_database_management.DatabaseManagementExternalnoncontainerdatabaseExternalNonContainerDbmFeaturesManagementResource
		factories["oci_database_management_externalpluggabledatabase_external_pluggable_dbm_features_management"] = tf_database_management.DatabaseManagementExternalpluggabledatabaseExternalPluggableDbmFeaturesManagementResource
		factories["oci_database_management_managed_database"] = tf_database_management.DatabaseManagementManagedDatabaseResource
		factories["oci_database_management_managed_database_group"] = tf_database_management.DatabaseManagementManagedDatabaseGroupResource
		factories["oci_database_management_managed_databases_change_database_parameter"] = tf_database_management.DatabaseManagementManagedDatabasesChangeDatabaseParameterResource
		factories["oci_database_management_managed_databases_reset_database_parameter"] = tf_database_management.DatabaseManagementManagedDatabasesResetDatabaseParameterResource
		factories["oci_database_management_named_credential"] = tf_database_management.DatabaseManagementNamedCredentialResource
		factories["oci_database_management_pluggabledatabase_pluggable_database_dbm_features_management"] = tf_database_management.DatabaseManagementPluggabledatabasePluggableDatabaseDbmFeaturesManagementResource
	}
	if common.CheckForEnabledServices("databasemigration") {
		factories["oci_database_migration_assessment"] = tf_database_migration.DatabaseMigrationAssessmentResource
		factories["oci_database_migration_assessment_assessor_action"] = tf_database_migration.DatabaseMigrationAssessmentAssessorActionResource
		factories["oci_database_migration_connection"] = tf_database_migration.DatabaseMigrationConnectionResource
		factories["oci_database_migration_job"] = tf_database_migration.DatabaseMigrationJobResource
		factories["oci_database_migration_job_advisor_report_check"] = tf_database_migration.DatabaseMigrationJobAdvisorReportCheckResource
		factories["oci_database_migration_migration"] = tf_database_migration.DatabaseMigrationMigrationResource
	}
	if common.CheckForEnabledServices("databasetools") {
		factories["oci_database_tools_database_tools_connection"] = tf_database_tools.DatabaseToolsDatabaseToolsConnectionResource
		factories["oci_database_tools_database_tools_database_api_gateway_config"] = tf_database_tools.DatabaseToolsDatabaseToolsDatabaseApiGatewayConfigResource
		factories["oci_database_tools_database_tools_identity"] = tf_database_tools.DatabaseToolsDatabaseToolsIdentityResource
		factories["oci_database_tools_database_tools_mcp_server"] = tf_database_tools.DatabaseToolsDatabaseToolsMcpServerResource
		factories["oci_database_tools_database_tools_mcp_toolset"] = tf_database_tools.DatabaseToolsDatabaseToolsMcpToolsetResource
		factories["oci_database_tools_database_tools_private_endpoint"] = tf_database_tools.DatabaseToolsDatabaseToolsPrivateEndpointResource
		factories["oci_database_tools_database_tools_sql_report"] = tf_database_tools.DatabaseToolsDatabaseToolsSqlReportResource
	}
	if common.CheckForEnabledServices("databasetoolsruntime") {
		factories["oci_database_tools_runtime_database_tools_connection_credential"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialResource
		factories["oci_database_tools_runtime_database_tools_connection_credential_execute_grantee"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialExecuteGranteeResource
		factories["oci_database_tools_runtime_database_tools_connection_credential_public_synonym"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionCredentialPublicSynonymResource
		factories["oci_database_tools_runtime_database_tools_connection_property_set"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsConnectionPropertySetResource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_global"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigGlobalResource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolResource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_api_spec"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolApiSpecResource
		factories["oci_database_tools_runtime_database_tools_database_api_gateway_config_pool_auto_api_spec"] = tf_database_tools_runtime.DatabaseToolsRuntimeDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResource
	}
	if common.CheckForEnabledServices("datacatalog") {
		factories["oci_datacatalog_catalog"] = tf_datacatalog.DatacatalogCatalogResource
		factories["oci_datacatalog_catalog_private_endpoint"] = tf_datacatalog.DatacatalogCatalogPrivateEndpointResource
		factories["oci_datacatalog_connection"] = tf_datacatalog.DatacatalogConnectionResource
		factories["oci_datacatalog_data_asset"] = tf_datacatalog.DatacatalogDataAssetResource
		factories["oci_datacatalog_metastore"] = tf_datacatalog.DatacatalogMetastoreResource
	}
	if common.CheckForEnabledServices("datacc") {
		factories["oci_datacc_infrastructure"] = tf_datacc.DataccInfrastructureResource
		factories["oci_datacc_vm_cluster_network"] = tf_datacc.DataccVmClusterNetworkResource
		factories["oci_datacc_vm_instance"] = tf_datacc.DataccVmInstanceResource
	}
	if common.CheckForEnabledServices("dataflow") {
		factories["oci_dataflow_application"] = tf_dataflow.DataflowApplicationResource
		factories["oci_dataflow_invoke_run"] = tf_dataflow.DataflowInvokeRunResource
		factories["oci_dataflow_pool"] = tf_dataflow.DataflowPoolResource
		factories["oci_dataflow_private_endpoint"] = tf_dataflow.DataflowPrivateEndpointResource
		factories["oci_dataflow_run_statement"] = tf_dataflow.DataflowRunStatementResource
		factories["oci_dataflow_sql_endpoint"] = tf_dataflow.DataflowSqlEndpointResource
	}
	if common.CheckForEnabledServices("dataintegration") {
		factories["oci_dataintegration_workspace"] = tf_dataintegration.DataintegrationWorkspaceResource
		factories["oci_dataintegration_workspace_application"] = tf_dataintegration.DataintegrationWorkspaceApplicationResource
		factories["oci_dataintegration_workspace_application_patch"] = tf_dataintegration.DataintegrationWorkspaceApplicationPatchResource
		factories["oci_dataintegration_workspace_application_schedule"] = tf_dataintegration.DataintegrationWorkspaceApplicationScheduleResource
		factories["oci_dataintegration_workspace_application_task_schedule"] = tf_dataintegration.DataintegrationWorkspaceApplicationTaskScheduleResource
		factories["oci_dataintegration_workspace_export_request"] = tf_dataintegration.DataintegrationWorkspaceExportRequestResource
		factories["oci_dataintegration_workspace_folder"] = tf_dataintegration.DataintegrationWorkspaceFolderResource
		factories["oci_dataintegration_workspace_import_request"] = tf_dataintegration.DataintegrationWorkspaceImportRequestResource
		factories["oci_dataintegration_workspace_project"] = tf_dataintegration.DataintegrationWorkspaceProjectResource
		factories["oci_dataintegration_workspace_task"] = tf_dataintegration.DataintegrationWorkspaceTaskResource
	}
	if common.CheckForEnabledServices("datascience") {
		factories["oci_datascience_compute_target"] = tf_datascience.DatascienceComputeTargetResource
		factories["oci_datascience_job"] = tf_datascience.DatascienceJobResource
		factories["oci_datascience_job_run"] = tf_datascience.DatascienceJobRunResource
		factories["oci_datascience_ml_application"] = tf_datascience.DatascienceMlApplicationResource
		factories["oci_datascience_ml_application_implementation"] = tf_datascience.DatascienceMlApplicationImplementationResource
		factories["oci_datascience_ml_application_instance"] = tf_datascience.DatascienceMlApplicationInstanceResource
		factories["oci_datascience_model"] = tf_datascience.DatascienceModelResource
		factories["oci_datascience_model_artifact_export"] = tf_datascience.DatascienceModelArtifactExportResource
		factories["oci_datascience_model_artifact_import"] = tf_datascience.DatascienceModelArtifactImportResource
		factories["oci_datascience_model_custom_metadata_artifact"] = tf_datascience.DatascienceModelCustomMetadataArtifactResource
		factories["oci_datascience_model_defined_metadata_artifact"] = tf_datascience.DatascienceModelDefinedMetadataArtifactResource
		factories["oci_datascience_model_deployment"] = tf_datascience.DatascienceModelDeploymentResource
		factories["oci_datascience_model_group"] = tf_datascience.DatascienceModelGroupResource
		factories["oci_datascience_model_group_artifact"] = tf_datascience.DatascienceModelGroupArtifactResource
		factories["oci_datascience_model_group_version_history"] = tf_datascience.DatascienceModelGroupVersionHistoryResource
		factories["oci_datascience_model_provenance"] = tf_datascience.DatascienceModelProvenanceResource
		factories["oci_datascience_model_version_set"] = tf_datascience.DatascienceModelVersionSetResource
		factories["oci_datascience_notebook_session"] = tf_datascience.DatascienceNotebookSessionResource
		factories["oci_datascience_pipeline"] = tf_datascience.DatasciencePipelineResource
		factories["oci_datascience_pipeline_run"] = tf_datascience.DatasciencePipelineRunResource
		factories["oci_datascience_private_endpoint"] = tf_datascience.DatasciencePrivateEndpointResource
		factories["oci_datascience_project"] = tf_datascience.DatascienceProjectResource
		factories["oci_datascience_schedule"] = tf_datascience.DatascienceScheduleResource
	}
	if common.CheckForEnabledServices("dbmulticloud") {
		factories["oci_dbmulticloud_multi_cloud_resource_discovery"] = tf_dbmulticloud.DbmulticloudMultiCloudResourceDiscoveryResource
		factories["oci_dbmulticloud_oracle_db_aws_identity_connector"] = tf_dbmulticloud.DbmulticloudOracleDbAwsIdentityConnectorResource
		factories["oci_dbmulticloud_oracle_db_aws_key"] = tf_dbmulticloud.DbmulticloudOracleDbAwsKeyResource
		factories["oci_dbmulticloud_oracle_db_azure_blob_container"] = tf_dbmulticloud.DbmulticloudOracleDbAzureBlobContainerResource
		factories["oci_dbmulticloud_oracle_db_azure_blob_mount"] = tf_dbmulticloud.DbmulticloudOracleDbAzureBlobMountResource
		factories["oci_dbmulticloud_oracle_db_azure_connector"] = tf_dbmulticloud.DbmulticloudOracleDbAzureConnectorResource
		factories["oci_dbmulticloud_oracle_db_azure_vault"] = tf_dbmulticloud.DbmulticloudOracleDbAzureVaultResource
		factories["oci_dbmulticloud_oracle_db_azure_vault_association"] = tf_dbmulticloud.DbmulticloudOracleDbAzureVaultAssociationResource
		factories["oci_dbmulticloud_oracle_db_gcp_identity_connector"] = tf_dbmulticloud.DbmulticloudOracleDbGcpIdentityConnectorResource
		factories["oci_dbmulticloud_oracle_db_gcp_key_ring"] = tf_dbmulticloud.DbmulticloudOracleDbGcpKeyRingResource
	}
	if common.CheckForEnabledServices("dblm") {
		factories["oci_dblm_vulnerability_scan"] = tf_dblm.DblmVulnerabilityScanResource
	}
	if common.CheckForEnabledServices("ddfs") {
		factories["oci_ddfs_instance"] = tf_ddfs.DdfsInstanceResource
	}
	if common.CheckForEnabledServices("delegateaccesscontrol") {
		factories["oci_delegate_access_control_delegation_control"] = tf_delegate_access_control.DelegateAccessControlDelegationControlResource
		factories["oci_delegate_access_control_delegation_subscription"] = tf_delegate_access_control.DelegateAccessControlDelegationSubscriptionResource
	}
	if common.CheckForEnabledServices("demandsignal") {
		factories["oci_demand_signal_occ_demand_signal"] = tf_demand_signal.DemandSignalOccDemandSignalResource
		factories["oci_demand_signal_occ_metric_alarm"] = tf_demand_signal.DemandSignalOccMetricAlarmResource
	}
	if common.CheckForEnabledServices("desktops") {
		factories["oci_desktops_desktop_pool"] = tf_desktops.DesktopsDesktopPoolResource
	}
	if common.CheckForEnabledServices("devops") {
		factories["oci_devops_build_pipeline"] = tf_devops.DevopsBuildPipelineResource
		factories["oci_devops_build_pipeline_stage"] = tf_devops.DevopsBuildPipelineStageResource
		factories["oci_devops_build_run"] = tf_devops.DevopsBuildRunResource
		factories["oci_devops_connection"] = tf_devops.DevopsConnectionResource
		factories["oci_devops_deploy_artifact"] = tf_devops.DevopsDeployArtifactResource
		factories["oci_devops_deploy_environment"] = tf_devops.DevopsDeployEnvironmentResource
		factories["oci_devops_deploy_pipeline"] = tf_devops.DevopsDeployPipelineResource
		factories["oci_devops_deploy_stage"] = tf_devops.DevopsDeployStageResource
		factories["oci_devops_deployment"] = tf_devops.DevopsDeploymentResource
		factories["oci_devops_project"] = tf_devops.DevopsProjectResource
		factories["oci_devops_project_repository_setting"] = tf_devops.DevopsProjectRepositorySettingResource
		factories["oci_devops_repository"] = tf_devops.DevopsRepositoryResource
		factories["oci_devops_repository_mirror"] = tf_devops.DevopsRepositoryMirrorResource
		factories["oci_devops_repository_protected_branch_management"] = tf_devops.DevopsRepositoryProtectedBranchManagementResource
		factories["oci_devops_repository_ref"] = tf_devops.DevopsRepositoryRefResource
		factories["oci_devops_repository_setting"] = tf_devops.DevopsRepositorySettingResource
		factories["oci_devops_trigger"] = tf_devops.DevopsTriggerResource
	}
	if common.CheckForEnabledServices("dif") {
		factories["oci_dif_stack"] = tf_dif.DifStackResource
	}
	if common.CheckForEnabledServices("disasterrecovery") {
		factories["oci_disaster_recovery_automatic_dr_configuration"] = tf_disaster_recovery.DisasterRecoveryAutomaticDrConfigurationResource
		factories["oci_disaster_recovery_dr_plan"] = tf_disaster_recovery.DisasterRecoveryDrPlanResource
		factories["oci_disaster_recovery_dr_plan_execution"] = tf_disaster_recovery.DisasterRecoveryDrPlanExecutionResource
		factories["oci_disaster_recovery_dr_protection_group"] = tf_disaster_recovery.DisasterRecoveryDrProtectionGroupResource
	}
	if common.CheckForEnabledServices("dns") {
		factories["oci_dns_action_create_zone_from_zone_file"] = tf_dns.DnsActionCreateZoneFromZoneFileResource
		factories["oci_dns_record"] = tf_dns.DnsRecordResource
		factories["oci_dns_resolver"] = tf_dns.DnsResolverResource
		factories["oci_dns_resolver_endpoint"] = tf_dns.DnsResolverEndpointResource
		factories["oci_dns_rrset"] = tf_dns.DnsRrsetResource
		factories["oci_dns_steering_policy"] = tf_dns.DnsSteeringPolicyResource
		factories["oci_dns_steering_policy_attachment"] = tf_dns.DnsSteeringPolicyAttachmentResource
		factories["oci_dns_tsig_key"] = tf_dns.DnsTsigKeyResource
		factories["oci_dns_view"] = tf_dns.DnsViewResource
		factories["oci_dns_zone"] = tf_dns.DnsZoneResource
		factories["oci_dns_zone_promote_dnssec_key_version"] = tf_dns.DnsZonePromoteDnssecKeyVersionResource
		factories["oci_dns_zone_stage_dnssec_key_version"] = tf_dns.DnsZoneStageDnssecKeyVersionResource
	}
	if common.CheckForEnabledServices("email") {
		factories["oci_email_dkim"] = tf_email.EmailDkimResource
		factories["oci_email_email_domain"] = tf_email.EmailEmailDomainResource
		factories["oci_email_email_ip_pool"] = tf_email.EmailEmailIpPoolResource
		factories["oci_email_email_return_path"] = tf_email.EmailEmailReturnPathResource
		factories["oci_email_sender"] = tf_email.EmailSenderResource
		factories["oci_email_suppression"] = tf_email.EmailSuppressionResource
	}
	if common.CheckForEnabledServices("events") {
		factories["oci_events_rule"] = tf_events.EventsRuleResource
	}
	if common.CheckForEnabledServices("filestorage") {
		factories["oci_file_storage_export"] = tf_file_storage.FileStorageExportResource
		factories["oci_file_storage_export_set"] = tf_file_storage.FileStorageExportSetResource
		factories["oci_file_storage_file_system"] = tf_file_storage.FileStorageFileSystemResource
		factories["oci_file_storage_file_system_quota_rule"] = tf_file_storage.FileStorageFileSystemQuotaRuleResource
		factories["oci_file_storage_filesystem_snapshot_policy"] = tf_file_storage.FileStorageFilesystemSnapshotPolicyResource
		factories["oci_file_storage_mount_target"] = tf_file_storage.FileStorageMountTargetResource
		factories["oci_file_storage_outbound_connector"] = tf_file_storage.FileStorageOutboundConnectorResource
		factories["oci_file_storage_replication"] = tf_file_storage.FileStorageReplicationResource
		factories["oci_file_storage_snapshot"] = tf_file_storage.FileStorageSnapshotResource
	}
	if common.CheckForEnabledServices("fleetappsmanagement") {
		factories["oci_fleet_apps_management_catalog_item"] = tf_fleet_apps_management.FleetAppsManagementCatalogItemResource
		factories["oci_fleet_apps_management_compliance_policy_rule"] = tf_fleet_apps_management.FleetAppsManagementCompliancePolicyRuleResource
		factories["oci_fleet_apps_management_fleet"] = tf_fleet_apps_management.FleetAppsManagementFleetResource
		factories["oci_fleet_apps_management_fleet_credential"] = tf_fleet_apps_management.FleetAppsManagementFleetCredentialResource
		factories["oci_fleet_apps_management_fleet_property"] = tf_fleet_apps_management.FleetAppsManagementFleetPropertyResource
		factories["oci_fleet_apps_management_fleet_resource"] = tf_fleet_apps_management.FleetAppsManagementFleetResourceResource
		factories["oci_fleet_apps_management_maintenance_window"] = tf_fleet_apps_management.FleetAppsManagementMaintenanceWindowResource
		factories["oci_fleet_apps_management_onboarding"] = tf_fleet_apps_management.FleetAppsManagementOnboardingResource
		factories["oci_fleet_apps_management_patch"] = tf_fleet_apps_management.FleetAppsManagementPatchResource
		factories["oci_fleet_apps_management_platform_configuration"] = tf_fleet_apps_management.FleetAppsManagementPlatformConfigurationResource
		factories["oci_fleet_apps_management_property"] = tf_fleet_apps_management.FleetAppsManagementPropertyResource
		factories["oci_fleet_apps_management_provision"] = tf_fleet_apps_management.FleetAppsManagementProvisionResource
		factories["oci_fleet_apps_management_runbook"] = tf_fleet_apps_management.FleetAppsManagementRunbookResource
		factories["oci_fleet_apps_management_runbook_version"] = tf_fleet_apps_management.FleetAppsManagementRunbookVersionResource
		factories["oci_fleet_apps_management_scheduler_definition"] = tf_fleet_apps_management.FleetAppsManagementSchedulerDefinitionResource
		factories["oci_fleet_apps_management_task_record"] = tf_fleet_apps_management.FleetAppsManagementTaskRecordResource
	}
	if common.CheckForEnabledServices("fleetsoftwareupdate") {
		factories["oci_fleet_software_update_fsu_collection"] = tf_fleet_software_update.FleetSoftwareUpdateFsuCollectionResource
		factories["oci_fleet_software_update_fsu_cycle"] = tf_fleet_software_update.FleetSoftwareUpdateFsuCycleResource
		factories["oci_fleet_software_update_fsu_readiness_check"] = tf_fleet_software_update.FleetSoftwareUpdateFsuReadinessCheckResource
	}
	if common.CheckForEnabledServices("functions") {
		factories["oci_functions_application"] = tf_functions.FunctionsApplicationResource
		factories["oci_functions_function"] = tf_functions.FunctionsFunctionResource
		factories["oci_functions_invoke_function"] = tf_functions.FunctionsInvokeFunctionResource
	}
	if common.CheckForEnabledServices("fusionapps") {
		factories["oci_fusion_apps_fusion_environment"] = tf_fusion_apps.FusionAppsFusionEnvironmentResource
		factories["oci_fusion_apps_fusion_environment_admin_user"] = tf_fusion_apps.FusionAppsFusionEnvironmentAdminUserResource
		factories["oci_fusion_apps_fusion_environment_data_masking_activity"] = tf_fusion_apps.FusionAppsFusionEnvironmentDataMaskingActivityResource
		factories["oci_fusion_apps_fusion_environment_family"] = tf_fusion_apps.FusionAppsFusionEnvironmentFamilyResource
		factories["oci_fusion_apps_fusion_environment_refresh_activity"] = tf_fusion_apps.FusionAppsFusionEnvironmentRefreshActivityResource
		factories["oci_fusion_apps_fusion_environment_service_attachment"] = tf_fusion_apps.FusionAppsFusionEnvironmentServiceAttachmentResource
	}
	if common.CheckForEnabledServices("gdp") {
		factories["oci_gdp_gdp_pipeline"] = tf_gdp.GdpGdpPipelineResource
	}
	if common.CheckForEnabledServices("generativeai") {
		factories["oci_generative_ai_dedicated_ai_cluster"] = tf_generative_ai.GenerativeAiDedicatedAiClusterResource
		factories["oci_generative_ai_endpoint"] = tf_generative_ai.GenerativeAiEndpointResource
		factories["oci_generative_ai_generative_ai_private_endpoint"] = tf_generative_ai.GenerativeAiGenerativeAiPrivateEndpointResource
		factories["oci_generative_ai_hosted_application"] = tf_generative_ai.GenerativeAiHostedApplicationResource
		factories["oci_generative_ai_hosted_application_iam"] = tf_generative_ai.GenerativeAiHostedApplicationIamResource
		factories["oci_generative_ai_hosted_application_storage"] = tf_generative_ai.GenerativeAiHostedApplicationStorageResource
		factories["oci_generative_ai_hosted_deployment"] = tf_generative_ai.GenerativeAiHostedDeploymentResource
		factories["oci_generative_ai_imported_model"] = tf_generative_ai.GenerativeAiImportedModelResource
		factories["oci_generative_ai_model"] = tf_generative_ai.GenerativeAiModelResource
		factories["oci_generative_ai_project"] = tf_generative_ai.GenerativeAiProjectResource
		factories["oci_generative_ai_routing_profile"] = tf_generative_ai.GenerativeAiRoutingProfileResource
		factories["oci_generative_ai_semantic_store"] = tf_generative_ai.GenerativeAiSemanticStoreResource
	}
	if common.CheckForEnabledServices("generativeaiagent") {
		factories["oci_generative_ai_agent_agent"] = tf_generative_ai_agent.GenerativeAiAgentAgentResource
		factories["oci_generative_ai_agent_agent_endpoint"] = tf_generative_ai_agent.GenerativeAiAgentAgentEndpointResource
		factories["oci_generative_ai_agent_data_ingestion_job"] = tf_generative_ai_agent.GenerativeAiAgentDataIngestionJobResource
		factories["oci_generative_ai_agent_data_source"] = tf_generative_ai_agent.GenerativeAiAgentDataSourceResource
		factories["oci_generative_ai_agent_knowledge_base"] = tf_generative_ai_agent.GenerativeAiAgentKnowledgeBaseResource
		factories["oci_generative_ai_agent_provisioned_capacity"] = tf_generative_ai_agent.GenerativeAiAgentProvisionedCapacityResource
		factories["oci_generative_ai_agent_tool"] = tf_generative_ai_agent.GenerativeAiAgentToolResource
	}
	if common.CheckForEnabledServices("genericartifactscontent") {
		factories["oci_generic_artifacts_content_artifact_by_path"] = tf_generic_artifacts_content.GenericArtifactsContentArtifactByPathResource
	}
	if common.CheckForEnabledServices("goldengate") {
		factories["oci_golden_gate_connection"] = tf_golden_gate.GoldenGateConnectionResource
		factories["oci_golden_gate_connection_assignment"] = tf_golden_gate.GoldenGateConnectionAssignmentResource
		factories["oci_golden_gate_database_registration"] = tf_golden_gate.GoldenGateDatabaseRegistrationResource
		factories["oci_golden_gate_deployment"] = tf_golden_gate.GoldenGateDeploymentResource
		factories["oci_golden_gate_deployment_backup"] = tf_golden_gate.GoldenGateDeploymentBackupResource
		factories["oci_golden_gate_deployment_certificate"] = tf_golden_gate.GoldenGateDeploymentCertificateResource
		factories["oci_golden_gate_pipeline"] = tf_golden_gate.GoldenGatePipelineResource
	}
	if common.CheckForEnabledServices("healthchecks") {
		factories["oci_health_checks_http_monitor"] = tf_health_checks.HealthChecksHttpMonitorResource
		factories["oci_health_checks_http_probe"] = tf_health_checks.HealthChecksHttpProbeResource
		factories["oci_health_checks_ping_monitor"] = tf_health_checks.HealthChecksPingMonitorResource
		factories["oci_health_checks_ping_probe"] = tf_health_checks.HealthChecksPingProbeResource
	}
	if common.CheckForEnabledServices("identity") {
		factories["oci_identity_api_key"] = tf_identity.IdentityApiKeyResource
		factories["oci_identity_auth_token"] = tf_identity.IdentityAuthTokenResource
		factories["oci_identity_authentication_policy"] = tf_identity.IdentityAuthenticationPolicyResource
		factories["oci_identity_compartment"] = tf_identity.IdentityCompartmentResource
		factories["oci_identity_customer_secret_key"] = tf_identity.IdentityCustomerSecretKeyResource
		factories["oci_identity_db_credential"] = tf_identity.IdentityDbCredentialResource
		factories["oci_identity_domain"] = tf_identity.IdentityDomainResource
		factories["oci_identity_domain_replication_to_region"] = tf_identity.IdentityDomainReplicationToRegionResource
		factories["oci_identity_dynamic_group"] = tf_identity.IdentityDynamicGroupResource
		factories["oci_identity_group"] = tf_identity.IdentityGroupResource
		factories["oci_identity_identity_provider"] = tf_identity.IdentityIdentityProviderResource
		factories["oci_identity_idp_group_mapping"] = tf_identity.IdentityIdpGroupMappingResource
		factories["oci_identity_import_standard_tags_management"] = tf_identity.IdentityImportStandardTagsManagementResource
		factories["oci_identity_network_source"] = tf_identity.IdentityNetworkSourceResource
		factories["oci_identity_policy"] = tf_identity.IdentityPolicyResource
		factories["oci_identity_smtp_credential"] = tf_identity.IdentitySmtpCredentialResource
		factories["oci_identity_tag"] = tf_identity.IdentityTagResource
		factories["oci_identity_tag_default"] = tf_identity.IdentityTagDefaultResource
		factories["oci_identity_tag_namespace"] = tf_identity.IdentityTagNamespaceResource
		factories["oci_identity_ui_password"] = tf_identity.IdentityUiPasswordResource
		factories["oci_identity_user"] = tf_identity.IdentityUserResource
		factories["oci_identity_user_capabilities_management"] = tf_identity.IdentityUserCapabilitiesManagementResource
		factories["oci_identity_user_group_membership"] = tf_identity.IdentityUserGroupMembershipResource
	}
	if common.CheckForEnabledServices("identitydataplane") {
		factories["oci_identity_data_plane_generate_scoped_access_token"] = tf_identity_data_plane.IdentityDataPlaneGenerateScopedAccessTokenResource
	}
	if common.CheckForEnabledServices("identitydomains") {
		factories["oci_identity_domains_account_recovery_setting"] = tf_identity_domains.IdentityDomainsAccountRecoverySettingResource
		factories["oci_identity_domains_api_key"] = tf_identity_domains.IdentityDomainsApiKeyResource
		factories["oci_identity_domains_app"] = tf_identity_domains.IdentityDomainsAppResource
		factories["oci_identity_domains_app_role"] = tf_identity_domains.IdentityDomainsAppRoleResource
		factories["oci_identity_domains_approval_workflow"] = tf_identity_domains.IdentityDomainsApprovalWorkflowResource
		factories["oci_identity_domains_approval_workflow_assignment"] = tf_identity_domains.IdentityDomainsApprovalWorkflowAssignmentResource
		factories["oci_identity_domains_approval_workflow_step"] = tf_identity_domains.IdentityDomainsApprovalWorkflowStepResource
		factories["oci_identity_domains_auth_token"] = tf_identity_domains.IdentityDomainsAuthTokenResource
		factories["oci_identity_domains_authentication_factor_setting"] = tf_identity_domains.IdentityDomainsAuthenticationFactorSettingResource
		factories["oci_identity_domains_cloud_gate"] = tf_identity_domains.IdentityDomainsCloudGateResource
		factories["oci_identity_domains_cloud_gate_mapping"] = tf_identity_domains.IdentityDomainsCloudGateMappingResource
		factories["oci_identity_domains_cloud_gate_server"] = tf_identity_domains.IdentityDomainsCloudGateServerResource
		factories["oci_identity_domains_condition"] = tf_identity_domains.IdentityDomainsConditionResource
		factories["oci_identity_domains_customer_secret_key"] = tf_identity_domains.IdentityDomainsCustomerSecretKeyResource
		factories["oci_identity_domains_dynamic_resource_group"] = tf_identity_domains.IdentityDomainsDynamicResourceGroupResource
		factories["oci_identity_domains_grant"] = tf_identity_domains.IdentityDomainsGrantResource
		factories["oci_identity_domains_group"] = tf_identity_domains.IdentityDomainsGroupResource
		factories["oci_identity_domains_identity_proofing_provider"] = tf_identity_domains.IdentityDomainsIdentityProofingProviderResource
		factories["oci_identity_domains_identity_proofing_provider_template"] = tf_identity_domains.IdentityDomainsIdentityProofingProviderTemplateResource
		factories["oci_identity_domains_identity_propagation_trust"] = tf_identity_domains.IdentityDomainsIdentityPropagationTrustResource
		factories["oci_identity_domains_identity_provider"] = tf_identity_domains.IdentityDomainsIdentityProviderResource
		factories["oci_identity_domains_identity_setting"] = tf_identity_domains.IdentityDomainsIdentitySettingResource
		factories["oci_identity_domains_kmsi_setting"] = tf_identity_domains.IdentityDomainsKmsiSettingResource
		factories["oci_identity_domains_mapped_attribute"] = tf_identity_domains.IdentityDomainsMappedAttributeResource
		factories["oci_identity_domains_my_api_key"] = tf_identity_domains.IdentityDomainsMyApiKeyResource
		factories["oci_identity_domains_my_auth_token"] = tf_identity_domains.IdentityDomainsMyAuthTokenResource
		factories["oci_identity_domains_my_customer_secret_key"] = tf_identity_domains.IdentityDomainsMyCustomerSecretKeyResource
		factories["oci_identity_domains_my_oauth2client_credential"] = tf_identity_domains.IdentityDomainsMyOAuth2ClientCredentialResource
		factories["oci_identity_domains_my_request"] = tf_identity_domains.IdentityDomainsMyRequestResource
		factories["oci_identity_domains_my_smtp_credential"] = tf_identity_domains.IdentityDomainsMySmtpCredentialResource
		factories["oci_identity_domains_my_support_account"] = tf_identity_domains.IdentityDomainsMySupportAccountResource
		factories["oci_identity_domains_my_user_db_credential"] = tf_identity_domains.IdentityDomainsMyUserDbCredentialResource
		factories["oci_identity_domains_network_perimeter"] = tf_identity_domains.IdentityDomainsNetworkPerimeterResource
		factories["oci_identity_domains_notification_setting"] = tf_identity_domains.IdentityDomainsNotificationSettingResource
		factories["oci_identity_domains_oauth2client_credential"] = tf_identity_domains.IdentityDomainsOAuth2ClientCredentialResource
		factories["oci_identity_domains_oauth_client_certificate"] = tf_identity_domains.IdentityDomainsOAuthClientCertificateResource
		factories["oci_identity_domains_oauth_partner_certificate"] = tf_identity_domains.IdentityDomainsOAuthPartnerCertificateResource
		factories["oci_identity_domains_password_policy"] = tf_identity_domains.IdentityDomainsPasswordPolicyResource
		factories["oci_identity_domains_policy"] = tf_identity_domains.IdentityDomainsPolicyResource
		factories["oci_identity_domains_rule"] = tf_identity_domains.IdentityDomainsRuleResource
		factories["oci_identity_domains_security_question"] = tf_identity_domains.IdentityDomainsSecurityQuestionResource
		factories["oci_identity_domains_security_question_setting"] = tf_identity_domains.IdentityDomainsSecurityQuestionSettingResource
		factories["oci_identity_domains_self_registration_profile"] = tf_identity_domains.IdentityDomainsSelfRegistrationProfileResource
		factories["oci_identity_domains_setting"] = tf_identity_domains.IdentityDomainsSettingResource
		factories["oci_identity_domains_smtp_credential"] = tf_identity_domains.IdentityDomainsSmtpCredentialResource
		factories["oci_identity_domains_social_identity_provider"] = tf_identity_domains.IdentityDomainsSocialIdentityProviderResource
		factories["oci_identity_domains_user"] = tf_identity_domains.IdentityDomainsUserResource
		factories["oci_identity_domains_user_db_credential"] = tf_identity_domains.IdentityDomainsUserDbCredentialResource
	}
	if common.CheckForEnabledServices("integration") {
		factories["oci_integration_integration_instance"] = tf_integration.IntegrationIntegrationInstanceResource
		factories["oci_integration_oracle_managed_custom_endpoint"] = tf_integration.IntegrationCustomEndpointResource
		factories["oci_integration_private_endpoint_outbound_connection"] = tf_integration.IntegrationPrivateEndpointOutboundConnectionResource
	}
	if common.CheckForEnabledServices("iot") {
		factories["oci_iot_digital_twin_adapter"] = tf_iot.IotDigitalTwinAdapterResource
		factories["oci_iot_digital_twin_instance"] = tf_iot.IotDigitalTwinInstanceResource
		factories["oci_iot_digital_twin_instance_invoke_raw_command"] = tf_iot.IotDigitalTwinInstanceInvokeRawCommandResource
		factories["oci_iot_digital_twin_model"] = tf_iot.IotDigitalTwinModelResource
		factories["oci_iot_digital_twin_relationship"] = tf_iot.IotDigitalTwinRelationshipResource
		factories["oci_iot_iot_domain"] = tf_iot.IotIotDomainResource
		factories["oci_iot_iot_domain_change_data_retention_period"] = tf_iot.IotIotDomainChangeDataRetentionPeriodResource
		factories["oci_iot_iot_domain_configure_data_access"] = tf_iot.IotIotDomainConfigureDataAccessResource
		factories["oci_iot_iot_domain_group"] = tf_iot.IotIotDomainGroupResource
		factories["oci_iot_iot_domain_group_configure_data_access"] = tf_iot.IotIotDomainGroupConfigureDataAccessResource
		factories["oci_iot_iot_flow_runtime"] = tf_iot.IotIotFlowRuntimeResource
		factories["oci_iot_iot_flow_runtime_activate"] = tf_iot.IotIotFlowRuntimeActivateResource
		factories["oci_iot_iot_flow_runtime_deactivate"] = tf_iot.IotIotFlowRuntimeDeactivateResource
		factories["oci_iot_iot_flow_runtime_flow"] = tf_iot.IotIotFlowRuntimeFlowResource
	}
	if common.CheckForEnabledServices("jms") {
		factories["oci_jms_fleet"] = tf_jms.JmsFleetResource
		factories["oci_jms_fleet_advanced_feature_configuration"] = tf_jms.JmsFleetAdvancedFeatureConfigurationResource
		factories["oci_jms_fleet_agent_configuration"] = tf_jms.JmsFleetAgentConfigurationResource
		factories["oci_jms_jms_plugin"] = tf_jms.JmsJmsPluginResource
		factories["oci_jms_task_schedule"] = tf_jms.JmsTaskScheduleResource
	}
	if common.CheckForEnabledServices("jmsjavadownloads") {
		factories["oci_jms_java_downloads_java_download_report"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadReportResource
		factories["oci_jms_java_downloads_java_download_token"] = tf_jms_java_downloads.JmsJavaDownloadsJavaDownloadTokenResource
		factories["oci_jms_java_downloads_java_license_acceptance_record"] = tf_jms_java_downloads.JmsJavaDownloadsJavaLicenseAcceptanceRecordResource
	}
	if common.CheckForEnabledServices("jmsutils") {
		factories["oci_jms_utils_analyze_applications_configuration"] = tf_jms_utils.JmsUtilsAnalyzeApplicationsConfigurationResource
		factories["oci_jms_utils_subscription_acknowledgment_configuration"] = tf_jms_utils.JmsUtilsSubscriptionAcknowledgmentConfigurationResource
	}
	if common.CheckForEnabledServices("kms") {
		factories["oci_kms_ekms_private_endpoint"] = tf_kms.KmsEkmsPrivateEndpointResource
		factories["oci_kms_encrypted_data"] = tf_kms.KmsEncryptedDataResource
		factories["oci_kms_generated_key"] = tf_kms.KmsGeneratedKeyResource
		factories["oci_kms_key"] = tf_kms.KmsKeyResource
		factories["oci_kms_key_version"] = tf_kms.KmsKeyVersionResource
		factories["oci_kms_sign"] = tf_kms.KmsSignResource
		factories["oci_kms_vault"] = tf_kms.KmsVaultResource
		factories["oci_kms_vault_replication"] = tf_kms.KmsVaultReplicationResource
		factories["oci_kms_verify"] = tf_kms.KmsVerifyResource
	}
	if common.CheckForEnabledServices("licensemanager") {
		factories["oci_license_manager_configuration"] = tf_license_manager.LicenseManagerConfigurationResource
		factories["oci_license_manager_license_record"] = tf_license_manager.LicenseManagerLicenseRecordResource
		factories["oci_license_manager_product_license"] = tf_license_manager.LicenseManagerProductLicenseResource
	}
	if common.CheckForEnabledServices("limits") {
		factories["oci_limits_quota"] = tf_limits.LimitsQuotaResource
	}
	if common.CheckForEnabledServices("loadbalancer") {
		factories["oci_load_balancer_backend"] = tf_load_balancer.LoadBalancerBackendResource
		factories["oci_load_balancer_backend_set"] = tf_load_balancer.LoadBalancerBackendSetResource
		factories["oci_load_balancer_certificate"] = tf_load_balancer.LoadBalancerCertificateResource
		factories["oci_load_balancer_hostname"] = tf_load_balancer.LoadBalancerHostnameResource
		factories["oci_load_balancer_listener"] = tf_load_balancer.LoadBalancerListenerResource
		factories["oci_load_balancer_load_balancer"] = tf_load_balancer.LoadBalancerLoadBalancerResource
		factories["oci_load_balancer_load_balancer_routing_policy"] = tf_load_balancer.LoadBalancerLoadBalancerRoutingPolicyResource
		factories["oci_load_balancer_path_route_set"] = tf_load_balancer.LoadBalancerPathRouteSetResource
		factories["oci_load_balancer_rule_set"] = tf_load_balancer.LoadBalancerRuleSetResource
		factories["oci_load_balancer_ssl_cipher_suite"] = tf_load_balancer.LoadBalancerSslCipherSuiteResource
	}
	if common.CheckForEnabledServices("loganalytics") {
		factories["oci_log_analytics_log_analytics_entity"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityResource
		factories["oci_log_analytics_log_analytics_entity_associations_add"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityAssociationsAddResource
		factories["oci_log_analytics_log_analytics_entity_associations_remove"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityAssociationsRemoveResource
		factories["oci_log_analytics_log_analytics_entity_type"] = tf_log_analytics.LogAnalyticsLogAnalyticsEntityTypeResource
		factories["oci_log_analytics_log_analytics_import_custom_content"] = tf_log_analytics.LogAnalyticsLogAnalyticsImportCustomContentResource
		factories["oci_log_analytics_log_analytics_log_group"] = tf_log_analytics.LogAnalyticsLogAnalyticsLogGroupResource
		factories["oci_log_analytics_log_analytics_object_collection_rule"] = tf_log_analytics.LogAnalyticsLogAnalyticsObjectCollectionRuleResource
		factories["oci_log_analytics_log_analytics_preferences_management"] = tf_log_analytics.LogAnalyticsLogAnalyticsPreferencesManagementResource
		factories["oci_log_analytics_log_analytics_resource_categories_management"] = tf_log_analytics.LogAnalyticsLogAnalyticsResourceCategoriesManagementResource
		factories["oci_log_analytics_log_analytics_unprocessed_data_bucket_management"] = tf_log_analytics.LogAnalyticsLogAnalyticsUnprocessedDataBucketManagementResource
		factories["oci_log_analytics_namespace"] = tf_log_analytics.LogAnalyticsNamespaceResource
		factories["oci_log_analytics_namespace_association"] = tf_log_analytics.LogAnalyticsNamespaceAssociationResource
		factories["oci_log_analytics_namespace_ingest_time_rule"] = tf_log_analytics.LogAnalyticsNamespaceIngestTimeRuleResource
		factories["oci_log_analytics_namespace_ingest_time_rules_management"] = tf_log_analytics.LogAnalyticsNamespaceIngestTimeRulesManagementResource
		factories["oci_log_analytics_namespace_lookup"] = tf_log_analytics.LogAnalyticsNamespaceLookupResource
		factories["oci_log_analytics_namespace_lookups_append_data_management"] = tf_log_analytics.LogAnalyticsNamespaceLookupsAppendDataManagementResource
		factories["oci_log_analytics_namespace_lookups_update_data_management"] = tf_log_analytics.LogAnalyticsNamespaceLookupsUpdateDataManagementResource
		factories["oci_log_analytics_namespace_scheduled_task"] = tf_log_analytics.LogAnalyticsNamespaceScheduledTaskResource
		factories["oci_log_analytics_namespace_storage_archival_config"] = tf_log_analytics.LogAnalyticsNamespaceStorageArchivalConfigResource
		factories["oci_log_analytics_namespace_storage_enable_disable_archiving"] = tf_log_analytics.LogAnalyticsNamespaceStorageEnableDisableArchivingResource
	}
	if common.CheckForEnabledServices("logging") {
		factories["oci_logging_log"] = tf_logging.LoggingLogResource
		factories["oci_logging_log_group"] = tf_logging.LoggingLogGroupResource
		factories["oci_logging_log_saved_search"] = tf_logging.LoggingLogSavedSearchResource
		factories["oci_logging_unified_agent_configuration"] = tf_logging.LoggingUnifiedAgentConfigurationResource
	}
	if common.CheckForEnabledServices("lustrefilestorage") {
		factories["oci_lustre_file_storage_lustre_file_system"] = tf_lustre_file_storage.LustreFileStorageLustreFileSystemResource
		factories["oci_lustre_file_storage_object_storage_link"] = tf_lustre_file_storage.LustreFileStorageObjectStorageLinkResource
	}
	if common.CheckForEnabledServices("managedkafka") {
		factories["oci_managed_kafka_kafka_cluster"] = tf_managed_kafka.ManagedKafkaKafkaClusterResource
		factories["oci_managed_kafka_kafka_cluster_addon"] = tf_managed_kafka.ManagedKafkaKafkaClusterAddonResource
		factories["oci_managed_kafka_kafka_cluster_config"] = tf_managed_kafka.ManagedKafkaKafkaClusterConfigResource
		factories["oci_managed_kafka_kafka_cluster_superusers_management"] = tf_managed_kafka.ManagedKafkaKafkaClusterSuperusersManagementResource
	}
	if common.CheckForEnabledServices("managementagent") {
		factories["oci_management_agent_management_agent"] = tf_management_agent.ManagementAgentManagementAgentResource
		factories["oci_management_agent_management_agent_data_source"] = tf_management_agent.ManagementAgentManagementAgentDataSourceResource
		factories["oci_management_agent_management_agent_install_key"] = tf_management_agent.ManagementAgentManagementAgentInstallKeyResource
		factories["oci_management_agent_named_credential"] = tf_management_agent.ManagementAgentNamedCredentialResource
	}
	if common.CheckForEnabledServices("managementdashboard") {
		factories["oci_management_dashboard_management_dashboards_import"] = tf_management_dashboard.ManagementDashboardManagementDashboardsImportResource
		factories["oci_management_dashboard_management_saved_search"] = tf_management_dashboard.ManagementDashboardManagementSavedSearchResource
	}
	if common.CheckForEnabledServices("marketplace") {
		factories["oci_marketplace_accepted_agreement"] = tf_marketplace.MarketplaceAcceptedAgreementResource
		factories["oci_marketplace_listing_package_agreement"] = tf_marketplace.MarketplaceListingPackageAgreementResource
		factories["oci_marketplace_marketplace_external_attested_metadata"] = tf_marketplace.MarketplaceMarketplaceExternalAttestedMetadataResource
		factories["oci_marketplace_publication"] = tf_marketplace.MarketplacePublicationResource
	}
	if common.CheckForEnabledServices("mediaservices") {
		factories["oci_media_services_media_asset"] = tf_media_services.MediaServicesMediaAssetResource
		factories["oci_media_services_media_workflow"] = tf_media_services.MediaServicesMediaWorkflowResource
		factories["oci_media_services_media_workflow_configuration"] = tf_media_services.MediaServicesMediaWorkflowConfigurationResource
		factories["oci_media_services_media_workflow_job"] = tf_media_services.MediaServicesMediaWorkflowJobResource
		factories["oci_media_services_stream_cdn_config"] = tf_media_services.MediaServicesStreamCdnConfigResource
		factories["oci_media_services_stream_distribution_channel"] = tf_media_services.MediaServicesStreamDistributionChannelResource
		factories["oci_media_services_stream_packaging_config"] = tf_media_services.MediaServicesStreamPackagingConfigResource
	}
	if common.CheckForEnabledServices("meteringcomputation") {
		factories["oci_metering_computation_custom_table"] = tf_metering_computation.MeteringComputationCustomTableResource
		factories["oci_metering_computation_query"] = tf_metering_computation.MeteringComputationQueryResource
		factories["oci_metering_computation_schedule"] = tf_metering_computation.MeteringComputationScheduleResource
		factories["oci_metering_computation_usage"] = tf_metering_computation.MeteringComputationUsageResource
		factories["oci_metering_computation_usage_carbon_emission"] = tf_metering_computation.MeteringComputationUsageCarbonEmissionResource
		factories["oci_metering_computation_usage_carbon_emissions_query"] = tf_metering_computation.MeteringComputationUsageCarbonEmissionsQueryResource
		factories["oci_metering_computation_usage_statement_email_recipients_group"] = tf_metering_computation.MeteringComputationUsageStatementEmailRecipientsGroupResource
	}
	if common.CheckForEnabledServices("monitoring") {
		factories["oci_monitoring_alarm"] = tf_monitoring.MonitoringAlarmResource
		factories["oci_monitoring_alarm_suppression"] = tf_monitoring.MonitoringAlarmSuppressionResource
	}
	if common.CheckForEnabledServices("mysql") {
		factories["oci_mysql_blue_green_deployment"] = tf_mysql.MysqlBlueGreenDeploymentResource
		factories["oci_mysql_channel"] = tf_mysql.MysqlChannelResource
		factories["oci_mysql_heat_wave_cluster"] = tf_mysql.MysqlHeatWaveClusterResource
		factories["oci_mysql_mysql_backup"] = tf_mysql.MysqlMysqlBackupResource
		factories["oci_mysql_mysql_configuration"] = tf_mysql.MysqlMysqlConfigurationResource
		factories["oci_mysql_mysql_db_system"] = tf_mysql.MysqlMysqlDbSystemResource
		factories["oci_mysql_replica"] = tf_mysql.MysqlReplicaResource
	}
	if common.CheckForEnabledServices("networkfirewall") {
		factories["oci_network_firewall_network_firewall"] = tf_network_firewall.NetworkFirewallNetworkFirewallResource
		factories["oci_network_firewall_network_firewall_policy"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyResource
		factories["oci_network_firewall_network_firewall_policy_address_list"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyAddressListResource
		factories["oci_network_firewall_network_firewall_policy_application"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyApplicationResource
		factories["oci_network_firewall_network_firewall_policy_application_group"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyApplicationGroupResource
		factories["oci_network_firewall_network_firewall_policy_decryption_profile"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyDecryptionProfileResource
		factories["oci_network_firewall_network_firewall_policy_decryption_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyDecryptionRuleResource
		factories["oci_network_firewall_network_firewall_policy_mapped_secret"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyMappedSecretResource
		factories["oci_network_firewall_network_firewall_policy_nat_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyNatRuleResource
		factories["oci_network_firewall_network_firewall_policy_security_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicySecurityRuleResource
		factories["oci_network_firewall_network_firewall_policy_service"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyServiceResource
		factories["oci_network_firewall_network_firewall_policy_service_list"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyServiceListResource
		factories["oci_network_firewall_network_firewall_policy_tunnel_inspection_rule"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyTunnelInspectionRuleResource
		factories["oci_network_firewall_network_firewall_policy_url_list"] = tf_network_firewall.NetworkFirewallNetworkFirewallPolicyUrlListResource
	}
	if common.CheckForEnabledServices("networkloadbalancer") {
		factories["oci_network_load_balancer_backend"] = tf_network_load_balancer.NetworkLoadBalancerBackendResource
		factories["oci_network_load_balancer_backend_set"] = tf_network_load_balancer.NetworkLoadBalancerBackendSetResource
		factories["oci_network_load_balancer_listener"] = tf_network_load_balancer.NetworkLoadBalancerListenerResource
		factories["oci_network_load_balancer_network_load_balancer"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancerResource
		factories["oci_network_load_balancer_network_load_balancers_backend_sets_unified"] = tf_network_load_balancer.NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource
	}
	if common.CheckForEnabledServices("nosql") {
		factories["oci_nosql_configuration"] = tf_nosql.NosqlConfigurationResource
		factories["oci_nosql_index"] = tf_nosql.NosqlIndexResource
		factories["oci_nosql_table"] = tf_nosql.NosqlTableResource
		factories["oci_nosql_table_replica"] = tf_nosql.NosqlTableReplicaResource
	}
	if common.CheckForEnabledServices("objectstorage") {
		factories["oci_objectstorage_bucket"] = tf_objectstorage.ObjectStorageBucketResource
		factories["oci_objectstorage_namespace_metadata"] = tf_objectstorage.ObjectStorageNamespaceMetadataResource
		factories["oci_objectstorage_object"] = tf_objectstorage.ObjectStorageObjectResource
		factories["oci_objectstorage_object_lifecycle_policy"] = tf_objectstorage.ObjectStorageObjectLifecyclePolicyResource
		factories["oci_objectstorage_preauthrequest"] = tf_objectstorage.ObjectStoragePreauthenticatedRequestResource
		factories["oci_objectstorage_private_endpoint"] = tf_objectstorage.ObjectStoragePrivateEndpointResource
		factories["oci_objectstorage_replication_policy"] = tf_objectstorage.ObjectStorageReplicationPolicyResource
	}
	if common.CheckForEnabledServices("oce") {
		factories["oci_oce_oce_instance"] = tf_oce.OceOceInstanceResource
	}
	if common.CheckForEnabledServices("ocvp") {
		factories["oci_ocvp_byol"] = tf_ocvp.OcvpByolResource
		factories["oci_ocvp_byol_allocation"] = tf_ocvp.OcvpByolAllocationResource
		factories["oci_ocvp_cluster"] = tf_ocvp.OcvpClusterResource
		factories["oci_ocvp_datastore"] = tf_ocvp.OcvpDatastoreResource
		factories["oci_ocvp_datastore_cluster"] = tf_ocvp.OcvpDatastoreClusterResource
		factories["oci_ocvp_esxi_host"] = tf_ocvp.OcvpEsxiHostResource
		factories["oci_ocvp_management_appliance"] = tf_ocvp.OcvpManagementApplianceResource
		factories["oci_ocvp_sddc"] = tf_ocvp.OcvpSddcResource
	}
	if common.CheckForEnabledServices("oda") {
		factories["oci_oda_oda_instance"] = tf_oda.OdaOdaInstanceResource
		factories["oci_oda_oda_private_endpoint"] = tf_oda.OdaOdaPrivateEndpointResource
		factories["oci_oda_oda_private_endpoint_attachment"] = tf_oda.OdaOdaPrivateEndpointAttachmentResource
		factories["oci_oda_oda_private_endpoint_scan_proxy"] = tf_oda.OdaOdaPrivateEndpointScanProxyResource
	}
	if common.CheckForEnabledServices("ons") {
		factories["oci_ons_notification_topic"] = tf_ons.OnsNotificationTopicResource
		factories["oci_ons_subscription"] = tf_ons.OnsSubscriptionResource
	}
	if common.CheckForEnabledServices("opa") {
		factories["oci_opa_opa_instance"] = tf_opa.OpaOpaInstanceResource
	}
	if common.CheckForEnabledServices("opensearch") {
		factories["oci_opensearch_opensearch_cluster"] = tf_opensearch.OpensearchOpensearchClusterResource
		factories["oci_opensearch_opensearch_cluster_pipeline"] = tf_opensearch.OpensearchOpensearchClusterPipelineResource
	}
	if common.CheckForEnabledServices("operatoraccesscontrol") {
		factories["oci_operator_access_control_operator_control"] = tf_operator_access_control.OperatorAccessControlOperatorControlResource
		factories["oci_operator_access_control_operator_control_assignment"] = tf_operator_access_control.OperatorAccessControlOperatorControlAssignmentResource
	}
	if common.CheckForEnabledServices("opsi") {
		factories["oci_opsi_awr_hub"] = tf_opsi.OpsiAwrHubResource
		factories["oci_opsi_awr_hub_source"] = tf_opsi.OpsiAwrHubSourceResource
		factories["oci_opsi_awr_hub_source_awrhubsources_management"] = tf_opsi.OpsiAwrHubSourceAwrhubsourcesManagementResource
		factories["oci_opsi_chargeback_plan"] = tf_opsi.OpsiChargebackPlanResource
		factories["oci_opsi_database_insight"] = tf_opsi.OpsiDatabaseInsightResource
		factories["oci_opsi_enterprise_manager_bridge"] = tf_opsi.OpsiEnterpriseManagerBridgeResource
		factories["oci_opsi_exadata_insight"] = tf_opsi.OpsiExadataInsightResource
		factories["oci_opsi_host_insight"] = tf_opsi.OpsiHostInsightResource
		factories["oci_opsi_news_report"] = tf_opsi.OpsiNewsReportResource
		factories["oci_opsi_operations_insights_private_endpoint"] = tf_opsi.OpsiOperationsInsightsPrivateEndpointResource
		factories["oci_opsi_operations_insights_warehouse"] = tf_opsi.OpsiOperationsInsightsWarehouseResource
		factories["oci_opsi_operations_insights_warehouse_download_warehouse_wallet"] = tf_opsi.OpsiOperationsInsightsWarehouseDownloadWarehouseWalletResource
		factories["oci_opsi_operations_insights_warehouse_rotate_warehouse_wallet"] = tf_opsi.OpsiOperationsInsightsWarehouseRotateWarehouseWalletResource
		factories["oci_opsi_operations_insights_warehouse_user"] = tf_opsi.OpsiOperationsInsightsWarehouseUserResource
		factories["oci_opsi_opsi_configuration"] = tf_opsi.OpsiOpsiConfigurationResource
	}
	if common.CheckForEnabledServices("optimizer") {
		factories["oci_optimizer_enrollment_status"] = tf_optimizer.OptimizerEnrollmentStatusResource
		factories["oci_optimizer_profile"] = tf_optimizer.OptimizerProfileResource
		factories["oci_optimizer_recommendation"] = tf_optimizer.OptimizerRecommendationResource
		factories["oci_optimizer_resource_action"] = tf_optimizer.OptimizerResourceActionResource
	}
	if common.CheckForEnabledServices("osmanagementhub") {
		factories["oci_os_management_hub_dynamic_set"] = tf_os_management_hub.OsManagementHubDynamicSetResource
		factories["oci_os_management_hub_dynamic_set_install_packages_management"] = tf_os_management_hub.OsManagementHubDynamicSetInstallPackagesManagementResource
		factories["oci_os_management_hub_dynamic_set_reboot_management"] = tf_os_management_hub.OsManagementHubDynamicSetRebootManagementResource
		factories["oci_os_management_hub_dynamic_set_remove_packages_management"] = tf_os_management_hub.OsManagementHubDynamicSetRemovePackagesManagementResource
		factories["oci_os_management_hub_dynamic_set_update_packages_management"] = tf_os_management_hub.OsManagementHubDynamicSetUpdatePackagesManagementResource
		factories["oci_os_management_hub_event"] = tf_os_management_hub.OsManagementHubEventResource
		factories["oci_os_management_hub_lifecycle_environment"] = tf_os_management_hub.OsManagementHubLifecycleEnvironmentResource
		factories["oci_os_management_hub_lifecycle_stage_attach_managed_instances_management"] = tf_os_management_hub.OsManagementHubLifecycleStageAttachManagedInstancesManagementResource
		factories["oci_os_management_hub_lifecycle_stage_detach_managed_instances_management"] = tf_os_management_hub.OsManagementHubLifecycleStageDetachManagedInstancesManagementResource
		factories["oci_os_management_hub_lifecycle_stage_promote_software_source_management"] = tf_os_management_hub.OsManagementHubLifecycleStagePromoteSoftwareSourceManagementResource
		factories["oci_os_management_hub_lifecycle_stage_reboot_management"] = tf_os_management_hub.OsManagementHubLifecycleStageRebootManagementResource
		factories["oci_os_management_hub_managed_instance"] = tf_os_management_hub.OsManagementHubManagedInstanceResource
		factories["oci_os_management_hub_managed_instance_attach_profile_management"] = tf_os_management_hub.OsManagementHubManagedInstanceAttachProfileManagementResource
		factories["oci_os_management_hub_managed_instance_attach_software_sources_management"] = tf_os_management_hub.OsManagementHubManagedInstanceAttachSoftwareSourcesManagementResource
		factories["oci_os_management_hub_managed_instance_detach_profile_management"] = tf_os_management_hub.OsManagementHubManagedInstanceDetachProfileManagementResource
		factories["oci_os_management_hub_managed_instance_detach_software_sources_management"] = tf_os_management_hub.OsManagementHubManagedInstanceDetachSoftwareSourcesManagementResource
		factories["oci_os_management_hub_managed_instance_group"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupResource
		factories["oci_os_management_hub_managed_instance_group_attach_managed_instances_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupAttachManagedInstancesManagementResource
		factories["oci_os_management_hub_managed_instance_group_attach_software_sources_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupAttachSoftwareSourcesManagementResource
		factories["oci_os_management_hub_managed_instance_group_detach_managed_instances_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupDetachManagedInstancesManagementResource
		factories["oci_os_management_hub_managed_instance_group_detach_software_sources_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupDetachSoftwareSourcesManagementResource
		factories["oci_os_management_hub_managed_instance_group_install_packages_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupInstallPackagesManagementResource
		factories["oci_os_management_hub_managed_instance_group_install_windows_updates_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupInstallWindowsUpdatesManagementResource
		factories["oci_os_management_hub_managed_instance_group_manage_module_streams_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupManageModuleStreamsManagementResource
		factories["oci_os_management_hub_managed_instance_group_reboot_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupRebootManagementResource
		factories["oci_os_management_hub_managed_instance_group_remove_packages_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupRemovePackagesManagementResource
		factories["oci_os_management_hub_managed_instance_group_update_all_packages_management"] = tf_os_management_hub.OsManagementHubManagedInstanceGroupUpdateAllPackagesManagementResource
		factories["oci_os_management_hub_managed_instance_install_packages_management"] = tf_os_management_hub.OsManagementHubManagedInstanceInstallPackagesManagementResource
		factories["oci_os_management_hub_managed_instance_install_snaps_management"] = tf_os_management_hub.OsManagementHubManagedInstanceInstallSnapsManagementResource
		factories["oci_os_management_hub_managed_instance_install_windows_updates_management"] = tf_os_management_hub.OsManagementHubManagedInstanceInstallWindowsUpdatesManagementResource
		factories["oci_os_management_hub_managed_instance_reboot_management"] = tf_os_management_hub.OsManagementHubManagedInstanceRebootManagementResource
		factories["oci_os_management_hub_managed_instance_refresh_software_management"] = tf_os_management_hub.OsManagementHubManagedInstanceRefreshSoftwareManagementResource
		factories["oci_os_management_hub_managed_instance_remove_packages_management"] = tf_os_management_hub.OsManagementHubManagedInstanceRemovePackagesManagementResource
		factories["oci_os_management_hub_managed_instance_remove_snaps_management"] = tf_os_management_hub.OsManagementHubManagedInstanceRemoveSnapsManagementResource
		factories["oci_os_management_hub_managed_instance_switch_snap_channel_management"] = tf_os_management_hub.OsManagementHubManagedInstanceSwitchSnapChannelManagementResource
		factories["oci_os_management_hub_managed_instance_update_packages_management"] = tf_os_management_hub.OsManagementHubManagedInstanceUpdatePackagesManagementResource
		factories["oci_os_management_hub_managed_instances_install_windows_updates_management"] = tf_os_management_hub.OsManagementHubManagedInstancesInstallWindowsUpdatesManagementResource
		factories["oci_os_management_hub_managed_instances_update_packages_management"] = tf_os_management_hub.OsManagementHubManagedInstancesUpdatePackagesManagementResource
		factories["oci_os_management_hub_management_station"] = tf_os_management_hub.OsManagementHubManagementStationResource
		factories["oci_os_management_hub_management_station_associate_managed_instances_management"] = tf_os_management_hub.OsManagementHubManagementStationAssociateManagedInstancesManagementResource
		factories["oci_os_management_hub_management_station_mirror_synchronize_management"] = tf_os_management_hub.OsManagementHubManagementStationMirrorSynchronizeManagementResource
		factories["oci_os_management_hub_management_station_refresh_management"] = tf_os_management_hub.OsManagementHubManagementStationRefreshManagementResource
		factories["oci_os_management_hub_management_station_synchronize_mirrors_management"] = tf_os_management_hub.OsManagementHubManagementStationSynchronizeMirrorsManagementResource
		factories["oci_os_management_hub_profile"] = tf_os_management_hub.OsManagementHubProfileResource
		factories["oci_os_management_hub_profile_attach_lifecycle_stage_management"] = tf_os_management_hub.OsManagementHubProfileAttachLifecycleStageManagementResource
		factories["oci_os_management_hub_profile_attach_managed_instance_group_management"] = tf_os_management_hub.OsManagementHubProfileAttachManagedInstanceGroupManagementResource
		factories["oci_os_management_hub_profile_attach_management_station_management"] = tf_os_management_hub.OsManagementHubProfileAttachManagementStationManagementResource
		factories["oci_os_management_hub_profile_attach_software_sources_management"] = tf_os_management_hub.OsManagementHubProfileAttachSoftwareSourcesManagementResource
		factories["oci_os_management_hub_profile_detach_management_station_management"] = tf_os_management_hub.OsManagementHubProfileDetachManagementStationManagementResource
		factories["oci_os_management_hub_profile_detach_software_sources_management"] = tf_os_management_hub.OsManagementHubProfileDetachSoftwareSourcesManagementResource
		factories["oci_os_management_hub_scheduled_job"] = tf_os_management_hub.OsManagementHubScheduledJobResource
		factories["oci_os_management_hub_software_source"] = tf_os_management_hub.OsManagementHubSoftwareSourceResource
		factories["oci_os_management_hub_software_source_add_packages_management"] = tf_os_management_hub.OsManagementHubSoftwareSourceAddPackagesManagementResource
		factories["oci_os_management_hub_software_source_change_availability_management"] = tf_os_management_hub.OsManagementHubSoftwareSourceChangeAvailabilityManagementResource
		factories["oci_os_management_hub_software_source_generate_metadata_management"] = tf_os_management_hub.OsManagementHubSoftwareSourceGenerateMetadataManagementResource
		factories["oci_os_management_hub_software_source_manifest"] = tf_os_management_hub.OsManagementHubSoftwareSourceManifestResource
		factories["oci_os_management_hub_software_source_remove_packages_management"] = tf_os_management_hub.OsManagementHubSoftwareSourceRemovePackagesManagementResource
		factories["oci_os_management_hub_software_source_replace_packages_management"] = tf_os_management_hub.OsManagementHubSoftwareSourceReplacePackagesManagementResource
		factories["oci_os_management_hub_work_request_rerun_management"] = tf_os_management_hub.OsManagementHubWorkRequestRerunManagementResource
	}
	if common.CheckForEnabledServices("ospgateway") {
		factories["oci_osp_gateway_address_action_verification"] = tf_osp_gateway.OspGatewayAddressActionVerificationResource
		factories["oci_osp_gateway_subscription"] = tf_osp_gateway.OspGatewaySubscriptionResource
	}
	if common.CheckForEnabledServices("psa") {
		factories["oci_psa_private_service_access"] = tf_psa.PsaPrivateServiceAccessResource
	}
	if common.CheckForEnabledServices("psql") {
		factories["oci_psql_backup"] = tf_psql.PsqlBackupResource
		factories["oci_psql_configuration"] = tf_psql.PsqlConfigurationResource
		factories["oci_psql_db_system"] = tf_psql.PsqlDbSystemResource
	}
	if common.CheckForEnabledServices("queue") {
		factories["oci_queue_consumer_group"] = tf_queue.QueueConsumerGroupResource
		factories["oci_queue_queue"] = tf_queue.QueueQueueResource
	}
	if common.CheckForEnabledServices("recovery") {
		factories["oci_recovery_protected_database"] = tf_recovery.RecoveryProtectedDatabaseResource
		factories["oci_recovery_protection_policy"] = tf_recovery.RecoveryProtectionPolicyResource
		factories["oci_recovery_recovery_service_subnet"] = tf_recovery.RecoveryRecoveryServiceSubnetResource
	}
	if common.CheckForEnabledServices("redis") {
		factories["oci_redis_oci_cache_backup"] = tf_redis.RedisOciCacheBackupResource
		factories["oci_redis_oci_cache_backup_export_to_object_storage"] = tf_redis.RedisOciCacheBackupExportToObjectStorageResource
		factories["oci_redis_oci_cache_config_set"] = tf_redis.RedisOciCacheConfigSetResource
		factories["oci_redis_oci_cache_config_setlist_associated_oci_cache_cluster"] = tf_redis.RedisOciCacheConfigSetlistAssociatedOciCacheClusterResource
		factories["oci_redis_oci_cache_user"] = tf_redis.RedisOciCacheUserResource
		factories["oci_redis_oci_cache_user_get_redis_cluster"] = tf_redis.RedisOciCacheUserGetRedisClusterResource
		factories["oci_redis_redis_cluster"] = tf_redis.RedisRedisClusterResource
		factories["oci_redis_redis_cluster_attach_oci_cache_user"] = tf_redis.RedisRedisClusterAttachOciCacheUserResource
		factories["oci_redis_redis_cluster_create_identity_token"] = tf_redis.RedisRedisClusterCreateIdentityTokenResource
		factories["oci_redis_redis_cluster_detach_oci_cache_user"] = tf_redis.RedisRedisClusterDetachOciCacheUserResource
		factories["oci_redis_redis_cluster_get_oci_cache_user"] = tf_redis.RedisRedisClusterGetOciCacheUserResource
	}
	if common.CheckForEnabledServices("resourceanalytics") {
		factories["oci_resource_analytics_monitored_region"] = tf_resource_analytics.ResourceAnalyticsMonitoredRegionResource
		factories["oci_resource_analytics_resource_analytics_instance"] = tf_resource_analytics.ResourceAnalyticsResourceAnalyticsInstanceResource
		factories["oci_resource_analytics_resource_analytics_instance_oac_management"] = tf_resource_analytics.ResourceAnalyticsResourceAnalyticsInstanceOacManagementResource
		factories["oci_resource_analytics_tenancy_attachment"] = tf_resource_analytics.ResourceAnalyticsTenancyAttachmentResource
	}
	if common.CheckForEnabledServices("resourcescheduler") {
		factories["oci_resource_scheduler_schedule"] = tf_resource_scheduler.ResourceSchedulerScheduleResource
	}
	if common.CheckForEnabledServices("resourcemanager") {
		factories["oci_resourcemanager_private_endpoint"] = tf_resourcemanager.ResourcemanagerPrivateEndpointResource
	}
	if common.CheckForEnabledServices("sch") {
		factories["oci_sch_service_connector"] = tf_sch.SchServiceConnectorResource
	}
	if common.CheckForEnabledServices("securityattribute") {
		factories["oci_security_attribute_security_attribute"] = tf_security_attribute.SecurityAttributeSecurityAttributeResource
		factories["oci_security_attribute_security_attribute_namespace"] = tf_security_attribute.SecurityAttributeSecurityAttributeNamespaceResource
	}
	if common.CheckForEnabledServices("self") {
		factories["oci_self_subscription"] = tf_self.SelfSubscriptionResource
	}
	if common.CheckForEnabledServices("servicecatalog") {
		factories["oci_service_catalog_private_application"] = tf_service_catalog.ServiceCatalogPrivateApplicationResource
		factories["oci_service_catalog_service_catalog"] = tf_service_catalog.ServiceCatalogServiceCatalogResource
		factories["oci_service_catalog_service_catalog_association"] = tf_service_catalog.ServiceCatalogServiceCatalogAssociationResource
	}
	if common.CheckForEnabledServices("stackmonitoring") {
		factories["oci_stack_monitoring_baselineable_metric"] = tf_stack_monitoring.StackMonitoringBaselineableMetricResource
		factories["oci_stack_monitoring_config"] = tf_stack_monitoring.StackMonitoringConfigResource
		factories["oci_stack_monitoring_discovery_job"] = tf_stack_monitoring.StackMonitoringDiscoveryJobResource
		factories["oci_stack_monitoring_maintenance_window"] = tf_stack_monitoring.StackMonitoringMaintenanceWindowResource
		factories["oci_stack_monitoring_maintenance_windows_retry_failed_operation"] = tf_stack_monitoring.StackMonitoringMaintenanceWindowsRetryFailedOperationResource
		factories["oci_stack_monitoring_maintenance_windows_stop"] = tf_stack_monitoring.StackMonitoringMaintenanceWindowsStopResource
		factories["oci_stack_monitoring_metric_extension"] = tf_stack_monitoring.StackMonitoringMetricExtensionResource
		factories["oci_stack_monitoring_metric_extension_metric_extension_on_given_resources_management"] = tf_stack_monitoring.StackMonitoringMetricExtensionMetricExtensionOnGivenResourcesManagementResource
		factories["oci_stack_monitoring_metric_extensions_test_management"] = tf_stack_monitoring.StackMonitoringMetricExtensionsTestManagementResource
		factories["oci_stack_monitoring_monitored_resource"] = tf_stack_monitoring.StackMonitoringMonitoredResourceResource
		factories["oci_stack_monitoring_monitored_resource_task"] = tf_stack_monitoring.StackMonitoringMonitoredResourceTaskResource
		factories["oci_stack_monitoring_monitored_resource_type"] = tf_stack_monitoring.StackMonitoringMonitoredResourceTypeResource
		factories["oci_stack_monitoring_monitored_resources_associate_monitored_resource"] = tf_stack_monitoring.StackMonitoringMonitoredResourcesAssociateMonitoredResourceResource
		factories["oci_stack_monitoring_monitored_resources_list_member"] = tf_stack_monitoring.StackMonitoringMonitoredResourcesListMemberResource
		factories["oci_stack_monitoring_monitored_resources_search"] = tf_stack_monitoring.StackMonitoringMonitoredResourcesSearchResource
		factories["oci_stack_monitoring_monitored_resources_search_association"] = tf_stack_monitoring.StackMonitoringMonitoredResourcesSearchAssociationResource
		factories["oci_stack_monitoring_monitoring_template"] = tf_stack_monitoring.StackMonitoringMonitoringTemplateResource
		factories["oci_stack_monitoring_monitoring_template_alarm_condition"] = tf_stack_monitoring.StackMonitoringMonitoringTemplateAlarmConditionResource
		factories["oci_stack_monitoring_monitoring_template_monitoring_template_on_given_resources_management"] = tf_stack_monitoring.StackMonitoringMonitoringTemplateMonitoringTemplateOnGivenResourcesManagementResource
		factories["oci_stack_monitoring_process_set"] = tf_stack_monitoring.StackMonitoringProcessSetResource
	}
	if common.CheckForEnabledServices("streaming") {
		factories["oci_streaming_connect_harness"] = tf_streaming.StreamingConnectHarnessResource
		factories["oci_streaming_stream"] = tf_streaming.StreamingStreamResource
		factories["oci_streaming_stream_pool"] = tf_streaming.StreamingStreamPoolResource
	}
	if common.CheckForEnabledServices("tenantmanagercontrolplane") {
		factories["oci_tenantmanagercontrolplane_subscription_mapping"] = tf_tenantmanagercontrolplane.TenantmanagercontrolplaneSubscriptionMappingResource
	}
	if common.CheckForEnabledServices("usageproxy") {
		factories["oci_usage_proxy_subscription_redeemable_user"] = tf_usage_proxy.UsageProxySubscriptionRedeemableUserResource
	}
	if common.CheckForEnabledServices("vault") {
		factories["oci_vault_secret"] = tf_vault.VaultSecretResource
	}
	if common.CheckForEnabledServices("vbsinst") {
		factories["oci_vbs_inst_vbs_instance"] = tf_vbs_inst.VbsInstVbsInstanceResource
	}
	if common.CheckForEnabledServices("visualbuilder") {
		factories["oci_visual_builder_vb_instance"] = tf_visual_builder.VisualBuilderVbInstanceResource
	}
	if common.CheckForEnabledServices("vnmonitoring") {
		factories["oci_vn_monitoring_path_analysi"] = tf_vn_monitoring.VnMonitoringPathAnalysiResource
		factories["oci_vn_monitoring_path_analyzer_test"] = tf_vn_monitoring.VnMonitoringPathAnalyzerTestResource
	}
	if common.CheckForEnabledServices("vulnerabilityscanning") {
		factories["oci_vulnerability_scanning_container_scan_recipe"] = tf_vulnerability_scanning.VulnerabilityScanningContainerScanRecipeResource
		factories["oci_vulnerability_scanning_container_scan_target"] = tf_vulnerability_scanning.VulnerabilityScanningContainerScanTargetResource
		factories["oci_vulnerability_scanning_host_scan_recipe"] = tf_vulnerability_scanning.VulnerabilityScanningHostScanRecipeResource
		factories["oci_vulnerability_scanning_host_scan_target"] = tf_vulnerability_scanning.VulnerabilityScanningHostScanTargetResource
	}
	if common.CheckForEnabledServices("waa") {
		factories["oci_waa_web_app_acceleration"] = tf_waa.WaaWebAppAccelerationResource
		factories["oci_waa_web_app_acceleration_policy"] = tf_waa.WaaWebAppAccelerationPolicyResource
	}
	if common.CheckForEnabledServices("waas") {
		factories["oci_waas_address_list"] = tf_waas.WaasAddressListResource
		factories["oci_waas_certificate"] = tf_waas.WaasCertificateResource
		factories["oci_waas_custom_protection_rule"] = tf_waas.WaasCustomProtectionRuleResource
		factories["oci_waas_http_redirect"] = tf_waas.WaasHttpRedirectResource
		factories["oci_waas_protection_rule"] = tf_waas.WaasProtectionRuleResource
		factories["oci_waas_purge_cache"] = tf_waas.WaasPurgeCacheResource
		factories["oci_waas_waas_policy"] = tf_waas.WaasWaasPolicyResource
	}
	if common.CheckForEnabledServices("waf") {
		factories["oci_waf_network_address_list"] = tf_waf.WafNetworkAddressListResource
		factories["oci_waf_web_app_firewall"] = tf_waf.WafWebAppFirewallResource
		factories["oci_waf_web_app_firewall_policy"] = tf_waf.WafWebAppFirewallPolicyResource
	}
	if common.CheckForEnabledServices("zpr") {
		factories["oci_zpr_configuration"] = tf_zpr.ZprConfigurationResource
		factories["oci_zpr_zpr_policy"] = tf_zpr.ZprZprPolicyResource
	}
	return factories
}
