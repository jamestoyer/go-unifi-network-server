// Code generated from Unifi Network Server API definitions.
// DO NOT EDIT.

// Copyright 2024 James Toyer
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package networkserver

type NetworkConf struct {
	AutoScaleEnabled                              *bool                                `json:"auto_scale_enabled,omitempty"`
	DhcpRelayEnabled                              *bool                                `json:"dhcp_relay_enabled,omitempty"`
	DhcpdBootEnabled                              *bool                                `json:"dhcpd_boot_enabled,omitempty"`
	DhcpdBootFilename                             *string                              `json:"dhcpd_boot_filename,omitempty"`
	DhcpdBootServer                               *string                              `json:"dhcpd_boot_server,omitempty"`
	DhcpdConflictChecking                         *bool                                `json:"dhcpd_conflict_checking,omitempty"`
	DhcpdDns1                                     *string                              `json:"dhcpd_dns_1,omitempty"`
	DhcpdDns2                                     *string                              `json:"dhcpd_dns_2,omitempty"`
	DhcpdDns3                                     *string                              `json:"dhcpd_dns_3,omitempty"`
	DhcpdDns4                                     *string                              `json:"dhcpd_dns_4,omitempty"`
	DhcpdDnsEnabled                               *bool                                `json:"dhcpd_dns_enabled,omitempty"`
	DhcpdEnabled                                  *bool                                `json:"dhcpd_enabled,omitempty"`
	DhcpdGateway                                  *string                              `json:"dhcpd_gateway,omitempty"`
	DhcpdGatewayEnabled                           *bool                                `json:"dhcpd_gateway_enabled,omitempty"`
	DhcpdIp1                                      *string                              `json:"dhcpd_ip_1,omitempty"`
	DhcpdIp2                                      *string                              `json:"dhcpd_ip_2,omitempty"`
	DhcpdIp3                                      *string                              `json:"dhcpd_ip_3,omitempty"`
	DhcpdLeasetime                                *int64                               `json:"dhcpd_leasetime,omitempty"`
	DhcpdMac1                                     *string                              `json:"dhcpd_mac_1,omitempty"`
	DhcpdMac2                                     *string                              `json:"dhcpd_mac_2,omitempty"`
	DhcpdMac3                                     *string                              `json:"dhcpd_mac_3,omitempty"`
	DhcpdNtp1                                     *string                              `json:"dhcpd_ntp_1,omitempty"`
	DhcpdNtp2                                     *string                              `json:"dhcpd_ntp_2,omitempty"`
	DhcpdNtpEnabled                               *bool                                `json:"dhcpd_ntp_enabled,omitempty"`
	DhcpdStart                                    *string                              `json:"dhcpd_start,omitempty"`
	DhcpdStop                                     *string                              `json:"dhcpd_stop,omitempty"`
	DhcpdTftpServer                               *string                              `json:"dhcpd_tftp_server,omitempty"`
	DhcpdTimeOffset                               *float64                             `json:"dhcpd_time_offset,omitempty"`
	DhcpdTimeOffsetEnabled                        *bool                                `json:"dhcpd_time_offset_enabled,omitempty"`
	DhcpdUnifiController                          *string                              `json:"dhcpd_unifi_controller,omitempty"`
	DhcpdWins1                                    *string                              `json:"dhcpd_wins_1,omitempty"`
	DhcpdWins2                                    *string                              `json:"dhcpd_wins_2,omitempty"`
	DhcpdWinsEnabled                              *bool                                `json:"dhcpd_wins_enabled,omitempty"`
	DhcpdWpadUrl                                  *string                              `json:"dhcpd_wpad_url,omitempty"`
	Dhcpdv6AllowSlaac                             *bool                                `json:"dhcpdv6_allow_slaac,omitempty"`
	Dhcpdv6Dns1                                   *string                              `json:"dhcpdv6_dns_1,omitempty"`
	Dhcpdv6Dns2                                   *string                              `json:"dhcpdv6_dns_2,omitempty"`
	Dhcpdv6Dns3                                   *string                              `json:"dhcpdv6_dns_3,omitempty"`
	Dhcpdv6Dns4                                   *string                              `json:"dhcpdv6_dns_4,omitempty"`
	Dhcpdv6DnsAuto                                *bool                                `json:"dhcpdv6_dns_auto,omitempty"`
	Dhcpdv6Enabled                                *bool                                `json:"dhcpdv6_enabled,omitempty"`
	Dhcpdv6Leasetime                              *int64                               `json:"dhcpdv6_leasetime,omitempty"`
	Dhcpdv6Start                                  *string                              `json:"dhcpdv6_start,omitempty"`
	Dhcpdv6Stop                                   *string                              `json:"dhcpdv6_stop,omitempty"`
	DhcpguardEnabled                              *bool                                `json:"dhcpguard_enabled,omitempty"`
	DomainName                                    *string                              `json:"domain_name,omitempty"`
	DpiEnabled                                    *bool                                `json:"dpi_enabled,omitempty"`
	DpigroupId                                    *string                              `json:"dpigroup_id,omitempty"`
	Enabled                                       *bool                                `json:"enabled,omitempty"`
	ExposedToSiteVpn                              *bool                                `json:"exposed_to_site_vpn,omitempty"`
	GatewayDevice                                 *string                              `json:"gateway_device,omitempty"`
	GatewayType                                   *string                              `json:"gateway_type,omitempty"`
	IgmpFastleave                                 *bool                                `json:"igmp_fastleave,omitempty"`
	IgmpForwardUnknownMulticast                   *bool                                `json:"igmp_forward_unknown_multicast,omitempty"`
	IgmpGroupmembership                           *int64                               `json:"igmp_groupmembership,omitempty"`
	IgmpMaxresponse                               *int64                               `json:"igmp_maxresponse,omitempty"`
	IgmpMcrtrexpiretime                           *int64                               `json:"igmp_mcrtrexpiretime,omitempty"`
	IgmpProxyDownstreamNetworkconfIds             *[]string                            `json:"igmp_proxy_downstream_networkconf_ids,omitempty"`
	IgmpProxyFor                                  *string                              `json:"igmp_proxy_for,omitempty"`
	IgmpProxyUpstream                             *bool                                `json:"igmp_proxy_upstream,omitempty"`
	IgmpQuerierSwitches                           *[]NetworkConfIgmpQuerierSwitches    `json:"igmp_querier_switches,omitempty"`
	IgmpSnooping                                  *bool                                `json:"igmp_snooping,omitempty"`
	IgmpSupression                                *bool                                `json:"igmp_supression,omitempty"`
	InterfaceMtu                                  *float64                             `json:"interface_mtu,omitempty"`
	InterfaceMtuEnabled                           *bool                                `json:"interface_mtu_enabled,omitempty"`
	InternetAccessEnabled                         *bool                                `json:"internet_access_enabled,omitempty"`
	IpSubnet                                      *string                              `json:"ip_subnet,omitempty"`
	IpsecDhGroup                                  *int64                               `json:"ipsec_dh_group,omitempty"`
	IpsecDynamicRouting                           *bool                                `json:"ipsec_dynamic_routing,omitempty"`
	IpsecEncryption                               *string                              `json:"ipsec_encryption,omitempty"`
	IpsecEspDhGroup                               *float64                             `json:"ipsec_esp_dh_group,omitempty"`
	IpsecEspEncryption                            *string                              `json:"ipsec_esp_encryption,omitempty"`
	IpsecEspHash                                  *string                              `json:"ipsec_esp_hash,omitempty"`
	IpsecEspLifetime                              *string                              `json:"ipsec_esp_lifetime,omitempty"`
	IpsecHash                                     *string                              `json:"ipsec_hash,omitempty"`
	IpsecIkeDhGroup                               *float64                             `json:"ipsec_ike_dh_group,omitempty"`
	IpsecIkeEncryption                            *string                              `json:"ipsec_ike_encryption,omitempty"`
	IpsecIkeHash                                  *string                              `json:"ipsec_ike_hash,omitempty"`
	IpsecIkeLifetime                              *string                              `json:"ipsec_ike_lifetime,omitempty"`
	IpsecInterface                                *string                              `json:"ipsec_interface,omitempty"`
	IpsecKeyExchange                              *string                              `json:"ipsec_key_exchange,omitempty"`
	IpsecLocalIdentifier                          *string                              `json:"ipsec_local_identifier,omitempty"`
	IpsecLocalIdentifierEnabled                   *bool                                `json:"ipsec_local_identifier_enabled,omitempty"`
	IpsecLocalIp                                  *string                              `json:"ipsec_local_ip,omitempty"`
	IpsecPeerIp                                   *string                              `json:"ipsec_peer_ip,omitempty"`
	IpsecPfs                                      *bool                                `json:"ipsec_pfs,omitempty"`
	IpsecProfile                                  *string                              `json:"ipsec_profile,omitempty"`
	IpsecRemoteIdentifier                         *string                              `json:"ipsec_remote_identifier,omitempty"`
	IpsecRemoteIdentifierEnabled                  *bool                                `json:"ipsec_remote_identifier_enabled,omitempty"`
	IpsecSeparateIkev2Networks                    *bool                                `json:"ipsec_separate_ikev2_networks,omitempty"`
	IpsecTunnelIp                                 *string                              `json:"ipsec_tunnel_ip,omitempty"`
	IpsecTunnelIpEnabled                          *bool                                `json:"ipsec_tunnel_ip_enabled,omitempty"`
	Ipv6ClientAddressAssignment                   *string                              `json:"ipv6_client_address_assignment,omitempty"`
	Ipv6InterfaceType                             *string                              `json:"ipv6_interface_type,omitempty"`
	Ipv6PdAutoPrefixidEnabled                     *bool                                `json:"ipv6_pd_auto_prefixid_enabled,omitempty"`
	Ipv6PdInterface                               *string                              `json:"ipv6_pd_interface,omitempty"`
	Ipv6PdPrefixid                                *string                              `json:"ipv6_pd_prefixid,omitempty"`
	Ipv6PdStart                                   *string                              `json:"ipv6_pd_start,omitempty"`
	Ipv6PdStop                                    *string                              `json:"ipv6_pd_stop,omitempty"`
	Ipv6RaEnabled                                 *bool                                `json:"ipv6_ra_enabled,omitempty"`
	Ipv6RaPreferredLifetime                       *float64                             `json:"ipv6_ra_preferred_lifetime,omitempty"`
	Ipv6RaPriority                                *string                              `json:"ipv6_ra_priority,omitempty"`
	Ipv6RaValidLifetime                           *float64                             `json:"ipv6_ra_valid_lifetime,omitempty"`
	Ipv6SettingPreference                         *string                              `json:"ipv6_setting_preference,omitempty"`
	Ipv6SingleNetworkInterface                    *string                              `json:"ipv6_single_network_interface,omitempty"`
	Ipv6Subnet                                    *string                              `json:"ipv6_subnet,omitempty"`
	Ipv6WanDelegationType                         *string                              `json:"ipv6_wan_delegation_type,omitempty"`
	IsNat                                         *bool                                `json:"is_nat,omitempty"`
	L2TpAllowWeakCiphers                          *bool                                `json:"l2tp_allow_weak_ciphers,omitempty"`
	L2TpInterface                                 *string                              `json:"l2tp_interface,omitempty"`
	L2TpLocalWanIp                                *string                              `json:"l2tp_local_wan_ip,omitempty"`
	LocalPort                                     *float64                             `json:"local_port,omitempty"`
	LteLanEnabled                                 *bool                                `json:"lte_lan_enabled,omitempty"`
	MacOverride                                   *string                              `json:"mac_override,omitempty"`
	MacOverrideEnabled                            *bool                                `json:"mac_override_enabled,omitempty"`
	MdnsEnabled                                   *bool                                `json:"mdns_enabled,omitempty"`
	Name                                          *string                              `json:"name,omitempty"`
	NatOutboundIpAddresses                        *[]NetworkConfNatOutboundIpAddresses `json:"nat_outbound_ip_addresses,omitempty"`
	NetworkIsolationEnabled                       *bool                                `json:"network_isolation_enabled,omitempty"`
	Networkgroup                                  *string                              `json:"networkgroup,omitempty"`
	OpenvpnConfiguration                          *string                              `json:"openvpn_configuration,omitempty"`
	OpenvpnConfigurationFilename                  *string                              `json:"openvpn_configuration_filename,omitempty"`
	OpenvpnEncryptionCipher                       *string                              `json:"openvpn_encryption_cipher,omitempty"`
	OpenvpnInterface                              *string                              `json:"openvpn_interface,omitempty"`
	OpenvpnLocalAddress                           *string                              `json:"openvpn_local_address,omitempty"`
	OpenvpnLocalPort                              *float64                             `json:"openvpn_local_port,omitempty"`
	OpenvpnLocalWanIp                             *string                              `json:"openvpn_local_wan_ip,omitempty"`
	OpenvpnMode                                   *string                              `json:"openvpn_mode,omitempty"`
	OpenvpnRemoteAddress                          *string                              `json:"openvpn_remote_address,omitempty"`
	OpenvpnRemoteHost                             *string                              `json:"openvpn_remote_host,omitempty"`
	OpenvpnRemotePort                             *float64                             `json:"openvpn_remote_port,omitempty"`
	OpenvpnUsername                               *string                              `json:"openvpn_username,omitempty"`
	PptpcRequireMppe                              *bool                                `json:"pptpc_require_mppe,omitempty"`
	PptpcRouteDistance                            *int64                               `json:"pptpc_route_distance,omitempty"`
	PptpcServerIp                                 *string                              `json:"pptpc_server_ip,omitempty"`
	PptpcUsername                                 *string                              `json:"pptpc_username,omitempty"`
	Priority                                      *int64                               `json:"priority,omitempty"`
	Purpose                                       *string                              `json:"purpose,omitempty"`
	RadiusprofileId                               *string                              `json:"radiusprofile_id,omitempty"`
	RemoteSiteId                                  *string                              `json:"remote_site_id,omitempty"`
	RemoteSiteSubnets                             *[]string                            `json:"remote_site_subnets,omitempty"`
	RemoteVpnDynamicSubnetsEnabled                *bool                                `json:"remote_vpn_dynamic_subnets_enabled,omitempty"`
	RemoteVpnSubnets                              *[]string                            `json:"remote_vpn_subnets,omitempty"`
	ReportWanEvent                                *bool                                `json:"report_wan_event,omitempty"`
	RequireMschapv2                               *bool                                `json:"require_mschapv2,omitempty"`
	RouteDistance                                 *int64                               `json:"route_distance,omitempty"`
	SettingPreference                             *string                              `json:"setting_preference,omitempty"`
	SingleNetworkLan                              *string                              `json:"single_network_lan,omitempty"`
	UidPolicyEnabled                              *bool                                `json:"uid_policy_enabled,omitempty"`
	UidPolicyName                                 *string                              `json:"uid_policy_name,omitempty"`
	UidPublicGatewayPort                          *float64                             `json:"uid_public_gateway_port,omitempty"`
	UidTrafficRulesAllowedIpsAndHostnames         *[]string                            `json:"uid_traffic_rules_allowed_ips_and_hostnames,omitempty"`
	UidTrafficRulesEnabled                        *bool                                `json:"uid_traffic_rules_enabled,omitempty"`
	UidVpnCustomRouting                           *[]string                            `json:"uid_vpn_custom_routing,omitempty"`
	UidVpnDefaultDnsSuffix                        *string                              `json:"uid_vpn_default_dns_suffix,omitempty"`
	UidVpnMasqueradeEnabled                       *bool                                `json:"uid_vpn_masquerade_enabled,omitempty"`
	UidVpnMaxConnectionTimeSeconds                *int64                               `json:"uid_vpn_max_connection_time_seconds,omitempty"`
	UidVpnSyncPublicIp                            *bool                                `json:"uid_vpn_sync_public_ip,omitempty"`
	UidVpnType                                    *string                              `json:"uid_vpn_type,omitempty"`
	UidWorkspaceUrl                               *string                              `json:"uid_workspace_url,omitempty"`
	UpnpLanEnabled                                *bool                                `json:"upnp_lan_enabled,omitempty"`
	UsergroupId                                   *string                              `json:"usergroup_id,omitempty"`
	Vlan                                          *float64                             `json:"vlan,omitempty"`
	VlanEnabled                                   *bool                                `json:"vlan_enabled,omitempty"`
	VpnClientConfigurationRemoteIpOverride        *string                              `json:"vpn_client_configuration_remote_ip_override,omitempty"`
	VpnClientConfigurationRemoteIpOverrideEnabled *bool                                `json:"vpn_client_configuration_remote_ip_override_enabled,omitempty"`
	VpnClientDefaultRoute                         *bool                                `json:"vpn_client_default_route,omitempty"`
	VpnClientPullDns                              *bool                                `json:"vpn_client_pull_dns,omitempty"`
	VpnProtocol                                   *string                              `json:"vpn_protocol,omitempty"`
	VpnType                                       *string                              `json:"vpn_type,omitempty"`
	VrrpIpSubnetGw1                               *string                              `json:"vrrp_ip_subnet_gw1,omitempty"`
	VrrpIpSubnetGw2                               *string                              `json:"vrrp_ip_subnet_gw2,omitempty"`
	VrrpVrid                                      *int64                               `json:"vrrp_vrid,omitempty"`
	WanDhcpCos                                    *int64                               `json:"wan_dhcp_cos,omitempty"`
	WanDhcpOptions                                *[]NetworkConfWanDhcpOptions         `json:"wan_dhcp_options,omitempty"`
	WanDhcpv6PdSize                               *int64                               `json:"wan_dhcpv6_pd_size,omitempty"`
	WanDns1                                       *string                              `json:"wan_dns1,omitempty"`
	WanDns2                                       *string                              `json:"wan_dns2,omitempty"`
	WanDns3                                       *string                              `json:"wan_dns3,omitempty"`
	WanDns4                                       *string                              `json:"wan_dns4,omitempty"`
	WanDnsPreference                              *string                              `json:"wan_dns_preference,omitempty"`
	WanDsliteRemoteHost                           *string                              `json:"wan_dslite_remote_host,omitempty"`
	WanEgressQos                                  *int64                               `json:"wan_egress_qos,omitempty"`
	WanGateway                                    *string                              `json:"wan_gateway,omitempty"`
	WanGatewayV6                                  *string                              `json:"wan_gateway_v6,omitempty"`
	WanIp                                         *string                              `json:"wan_ip,omitempty"`
	WanIpAliases                                  *[]string                            `json:"wan_ip_aliases,omitempty"`
	WanIpv6                                       *string                              `json:"wan_ipv6,omitempty"`
	WanIpv6Dns1                                   *string                              `json:"wan_ipv6_dns1,omitempty"`
	WanIpv6Dns2                                   *string                              `json:"wan_ipv6_dns2,omitempty"`
	WanIpv6DnsPreference                          *string                              `json:"wan_ipv6_dns_preference,omitempty"`
	WanLoadBalanceType                            *string                              `json:"wan_load_balance_type,omitempty"`
	WanLoadBalanceWeight                          *int64                               `json:"wan_load_balance_weight,omitempty"`
	WanNetmask                                    *string                              `json:"wan_netmask,omitempty"`
	WanNetworkgroup                               *string                              `json:"wan_networkgroup,omitempty"`
	WanPppoePasswordEnabled                       *bool                                `json:"wan_pppoe_password_enabled,omitempty"`
	WanPppoeUsernameEnabled                       *bool                                `json:"wan_pppoe_username_enabled,omitempty"`
	WanPrefixlen                                  *int64                               `json:"wan_prefixlen,omitempty"`
	WanProviderCapabilities                       *NetworkConfWanProviderCapabilities  `json:"wan_provider_capabilities,omitempty"`
	WanSmartqDownRate                             *int64                               `json:"wan_smartq_down_rate,omitempty"`
	WanSmartqEnabled                              *bool                                `json:"wan_smartq_enabled,omitempty"`
	WanSmartqUpRate                               *int64                               `json:"wan_smartq_up_rate,omitempty"`
	WanType                                       *string                              `json:"wan_type,omitempty"`
	WanTypeV6                                     *string                              `json:"wan_type_v6,omitempty"`
	WanUsername                                   *string                              `json:"wan_username,omitempty"`
	WanVlan                                       *float64                             `json:"wan_vlan,omitempty"`
	WanVlanEnabled                                *bool                                `json:"wan_vlan_enabled,omitempty"`
	WireguardClientConfigurationFile              *string                              `json:"wireguard_client_configuration_file,omitempty"`
	WireguardClientConfigurationFilename          *string                              `json:"wireguard_client_configuration_filename,omitempty"`
	WireguardClientMode                           *string                              `json:"wireguard_client_mode,omitempty"`
	WireguardClientPeerIp                         *string                              `json:"wireguard_client_peer_ip,omitempty"`
	WireguardClientPeerPort                       *float64                             `json:"wireguard_client_peer_port,omitempty"`
	WireguardClientPeerPublicKey                  *string                              `json:"wireguard_client_peer_public_key,omitempty"`
	WireguardClientPresharedKey                   *string                              `json:"wireguard_client_preshared_key,omitempty"`
	WireguardClientPresharedKeyEnabled            *bool                                `json:"wireguard_client_preshared_key_enabled,omitempty"`
	WireguardInterface                            *string                              `json:"wireguard_interface,omitempty"`
	WireguardLocalWanIp                           *string                              `json:"wireguard_local_wan_ip,omitempty"`
	WireguardPublicKey                            *string                              `json:"wireguard_public_key,omitempty"`
	XAuthKey                                      *string                              `json:"x_auth_key,omitempty"`
	XCaCrt                                        *string                              `json:"x_ca_crt,omitempty"`
	XCaKey                                        *string                              `json:"x_ca_key,omitempty"`
	XDhKey                                        *string                              `json:"x_dh_key,omitempty"`
	XIpsecPreSharedKey                            *string                              `json:"x_ipsec_pre_shared_key,omitempty"`
	XOpenvpnPassword                              *string                              `json:"x_openvpn_password,omitempty"`
	XOpenvpnSharedSecretKey                       *string                              `json:"x_openvpn_shared_secret_key,omitempty"`
	XPptpcPassword                                *string                              `json:"x_pptpc_password,omitempty"`
	XServerCrt                                    *string                              `json:"x_server_crt,omitempty"`
	XServerKey                                    *string                              `json:"x_server_key,omitempty"`
	XSharedClientCrt                              *string                              `json:"x_shared_client_crt,omitempty"`
	XSharedClientKey                              *string                              `json:"x_shared_client_key,omitempty"`
	XWanPassword                                  *string                              `json:"x_wan_password,omitempty"`
	XWireguardPrivateKey                          *string                              `json:"x_wireguard_private_key,omitempty"`
}

type NetworkConfIgmpQuerierSwitches struct {
	QuerierAddress *string `json:"querier_address,omitempty"`
	SwitchMac      *string `json:"switch_mac,omitempty"`
}

type NetworkConfNatOutboundIpAddresses struct {
	IpAddress       *string   `json:"ip_address,omitempty"`
	IpAddressPool   *[]string `json:"ip_address_pool,omitempty"`
	Mode            *string   `json:"mode,omitempty"`
	WanNetworkGroup *string   `json:"wan_network_group,omitempty"`
}

type NetworkConfWanDhcpOptions struct {
	OptionNumber *float64 `json:"optionNumber,omitempty"`
	Value        *string  `json:"value,omitempty"`
}

type NetworkConfWanProviderCapabilities struct {
	DownloadKilobitsPerSecond *int64 `json:"download_kilobits_per_second,omitempty"`
	UploadKilobitsPerSecond   *int64 `json:"upload_kilobits_per_second,omitempty"`
}
