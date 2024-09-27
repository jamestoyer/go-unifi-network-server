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

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
)

type NetworkConf struct {
	ID     *string `json:"_id,omitempty"`
	SiteID *string `json:"site_id,omitempty"`

	Hidden   *bool   `json:"attr_hidden,omitempty"`
	HiddenID *string `json:"attr_hidden_id,omitempty"`
	NoDelete *bool   `json:"attr_no_delete,omitempty"`
	NoEdit   *bool   `json:"attr_no_edit,omitempty"`

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

func (s *NetworkConf) GetAutoScaleEnabled() bool {
	if s == nil {
		return false
	}

	return *s.AutoScaleEnabled
}

func (s *NetworkConf) GetDhcpRelayEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpRelayEnabled
}

func (s *NetworkConf) GetDhcpdBootEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdBootEnabled
}

func (s *NetworkConf) GetDhcpdBootFilename() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdBootFilename
}

func (s *NetworkConf) GetDhcpdBootServer() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdBootServer
}

func (s *NetworkConf) GetDhcpdConflictChecking() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdConflictChecking
}

func (s *NetworkConf) GetDhcpdDns1() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdDns1
}

func (s *NetworkConf) GetDhcpdDns2() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdDns2
}

func (s *NetworkConf) GetDhcpdDns3() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdDns3
}

func (s *NetworkConf) GetDhcpdDns4() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdDns4
}

func (s *NetworkConf) GetDhcpdDnsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdDnsEnabled
}

func (s *NetworkConf) GetDhcpdEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdEnabled
}

func (s *NetworkConf) GetDhcpdGateway() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdGateway
}

func (s *NetworkConf) GetDhcpdGatewayEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdGatewayEnabled
}

func (s *NetworkConf) GetDhcpdIp1() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdIp1
}

func (s *NetworkConf) GetDhcpdIp2() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdIp2
}

func (s *NetworkConf) GetDhcpdIp3() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdIp3
}

func (s *NetworkConf) GetDhcpdLeasetime() int64 {
	if s == nil {
		return 0
	}

	return *s.DhcpdLeasetime
}

func (s *NetworkConf) GetDhcpdMac1() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdMac1
}

func (s *NetworkConf) GetDhcpdMac2() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdMac2
}

func (s *NetworkConf) GetDhcpdMac3() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdMac3
}

func (s *NetworkConf) GetDhcpdNtp1() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdNtp1
}

func (s *NetworkConf) GetDhcpdNtp2() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdNtp2
}

func (s *NetworkConf) GetDhcpdNtpEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdNtpEnabled
}

func (s *NetworkConf) GetDhcpdStart() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdStart
}

func (s *NetworkConf) GetDhcpdStop() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdStop
}

func (s *NetworkConf) GetDhcpdTftpServer() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdTftpServer
}

func (s *NetworkConf) GetDhcpdTimeOffset() float64 {
	if s == nil {
		return 0
	}

	return *s.DhcpdTimeOffset
}

func (s *NetworkConf) GetDhcpdTimeOffsetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdTimeOffsetEnabled
}

func (s *NetworkConf) GetDhcpdUnifiController() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdUnifiController
}

func (s *NetworkConf) GetDhcpdWins1() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdWins1
}

func (s *NetworkConf) GetDhcpdWins2() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdWins2
}

func (s *NetworkConf) GetDhcpdWinsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpdWinsEnabled
}

func (s *NetworkConf) GetDhcpdWpadUrl() string {
	if s == nil {
		return ""
	}

	return *s.DhcpdWpadUrl
}

func (s *NetworkConf) GetDhcpdv6AllowSlaac() bool {
	if s == nil {
		return false
	}

	return *s.Dhcpdv6AllowSlaac
}

func (s *NetworkConf) GetDhcpdv6Dns1() string {
	if s == nil {
		return ""
	}

	return *s.Dhcpdv6Dns1
}

func (s *NetworkConf) GetDhcpdv6Dns2() string {
	if s == nil {
		return ""
	}

	return *s.Dhcpdv6Dns2
}

func (s *NetworkConf) GetDhcpdv6Dns3() string {
	if s == nil {
		return ""
	}

	return *s.Dhcpdv6Dns3
}

func (s *NetworkConf) GetDhcpdv6Dns4() string {
	if s == nil {
		return ""
	}

	return *s.Dhcpdv6Dns4
}

func (s *NetworkConf) GetDhcpdv6DnsAuto() bool {
	if s == nil {
		return false
	}

	return *s.Dhcpdv6DnsAuto
}

func (s *NetworkConf) GetDhcpdv6Enabled() bool {
	if s == nil {
		return false
	}

	return *s.Dhcpdv6Enabled
}

func (s *NetworkConf) GetDhcpdv6Leasetime() int64 {
	if s == nil {
		return 0
	}

	return *s.Dhcpdv6Leasetime
}

func (s *NetworkConf) GetDhcpdv6Start() string {
	if s == nil {
		return ""
	}

	return *s.Dhcpdv6Start
}

func (s *NetworkConf) GetDhcpdv6Stop() string {
	if s == nil {
		return ""
	}

	return *s.Dhcpdv6Stop
}

func (s *NetworkConf) GetDhcpguardEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DhcpguardEnabled
}

func (s *NetworkConf) GetDomainName() string {
	if s == nil {
		return ""
	}

	return *s.DomainName
}

func (s *NetworkConf) GetDpiEnabled() bool {
	if s == nil {
		return false
	}

	return *s.DpiEnabled
}

func (s *NetworkConf) GetDpigroupId() string {
	if s == nil {
		return ""
	}

	return *s.DpigroupId
}

func (s *NetworkConf) GetEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Enabled
}

func (s *NetworkConf) GetExposedToSiteVpn() bool {
	if s == nil {
		return false
	}

	return *s.ExposedToSiteVpn
}

func (s *NetworkConf) GetGatewayDevice() string {
	if s == nil {
		return ""
	}

	return *s.GatewayDevice
}

func (s *NetworkConf) GetGatewayType() string {
	if s == nil {
		return ""
	}

	return *s.GatewayType
}

func (s *NetworkConf) GetIgmpFastleave() bool {
	if s == nil {
		return false
	}

	return *s.IgmpFastleave
}

func (s *NetworkConf) GetIgmpForwardUnknownMulticast() bool {
	if s == nil {
		return false
	}

	return *s.IgmpForwardUnknownMulticast
}

func (s *NetworkConf) GetIgmpGroupmembership() int64 {
	if s == nil {
		return 0
	}

	return *s.IgmpGroupmembership
}

func (s *NetworkConf) GetIgmpMaxresponse() int64 {
	if s == nil {
		return 0
	}

	return *s.IgmpMaxresponse
}

func (s *NetworkConf) GetIgmpMcrtrexpiretime() int64 {
	if s == nil {
		return 0
	}

	return *s.IgmpMcrtrexpiretime
}

func (s *NetworkConf) GetIgmpProxyDownstreamNetworkconfIds() []string {
	if s == nil || s.IgmpProxyDownstreamNetworkconfIds == nil {
		return nil
	}

	return *s.IgmpProxyDownstreamNetworkconfIds
}

func (s *NetworkConf) GetIgmpProxyFor() string {
	if s == nil {
		return ""
	}

	return *s.IgmpProxyFor
}

func (s *NetworkConf) GetIgmpProxyUpstream() bool {
	if s == nil {
		return false
	}

	return *s.IgmpProxyUpstream
}

func (s *NetworkConf) GetIgmpQuerierSwitches() []NetworkConfIgmpQuerierSwitches {
	if s == nil || s.IgmpQuerierSwitches == nil {
		return nil
	}

	return *s.IgmpQuerierSwitches
}

func (s *NetworkConf) GetIgmpSnooping() bool {
	if s == nil {
		return false
	}

	return *s.IgmpSnooping
}

func (s *NetworkConf) GetIgmpSupression() bool {
	if s == nil {
		return false
	}

	return *s.IgmpSupression
}

func (s *NetworkConf) GetInterfaceMtu() float64 {
	if s == nil {
		return 0
	}

	return *s.InterfaceMtu
}

func (s *NetworkConf) GetInterfaceMtuEnabled() bool {
	if s == nil {
		return false
	}

	return *s.InterfaceMtuEnabled
}

func (s *NetworkConf) GetInternetAccessEnabled() bool {
	if s == nil {
		return false
	}

	return *s.InternetAccessEnabled
}

func (s *NetworkConf) GetIpSubnet() string {
	if s == nil {
		return ""
	}

	return *s.IpSubnet
}

func (s *NetworkConf) GetIpsecDhGroup() int64 {
	if s == nil {
		return 0
	}

	return *s.IpsecDhGroup
}

func (s *NetworkConf) GetIpsecDynamicRouting() bool {
	if s == nil {
		return false
	}

	return *s.IpsecDynamicRouting
}

func (s *NetworkConf) GetIpsecEncryption() string {
	if s == nil {
		return ""
	}

	return *s.IpsecEncryption
}

func (s *NetworkConf) GetIpsecEspDhGroup() float64 {
	if s == nil {
		return 0
	}

	return *s.IpsecEspDhGroup
}

func (s *NetworkConf) GetIpsecEspEncryption() string {
	if s == nil {
		return ""
	}

	return *s.IpsecEspEncryption
}

func (s *NetworkConf) GetIpsecEspHash() string {
	if s == nil {
		return ""
	}

	return *s.IpsecEspHash
}

func (s *NetworkConf) GetIpsecEspLifetime() string {
	if s == nil {
		return ""
	}

	return *s.IpsecEspLifetime
}

func (s *NetworkConf) GetIpsecHash() string {
	if s == nil {
		return ""
	}

	return *s.IpsecHash
}

func (s *NetworkConf) GetIpsecIkeDhGroup() float64 {
	if s == nil {
		return 0
	}

	return *s.IpsecIkeDhGroup
}

func (s *NetworkConf) GetIpsecIkeEncryption() string {
	if s == nil {
		return ""
	}

	return *s.IpsecIkeEncryption
}

func (s *NetworkConf) GetIpsecIkeHash() string {
	if s == nil {
		return ""
	}

	return *s.IpsecIkeHash
}

func (s *NetworkConf) GetIpsecIkeLifetime() string {
	if s == nil {
		return ""
	}

	return *s.IpsecIkeLifetime
}

func (s *NetworkConf) GetIpsecInterface() string {
	if s == nil {
		return ""
	}

	return *s.IpsecInterface
}

func (s *NetworkConf) GetIpsecKeyExchange() string {
	if s == nil {
		return ""
	}

	return *s.IpsecKeyExchange
}

func (s *NetworkConf) GetIpsecLocalIdentifier() string {
	if s == nil {
		return ""
	}

	return *s.IpsecLocalIdentifier
}

func (s *NetworkConf) GetIpsecLocalIdentifierEnabled() bool {
	if s == nil {
		return false
	}

	return *s.IpsecLocalIdentifierEnabled
}

func (s *NetworkConf) GetIpsecLocalIp() string {
	if s == nil {
		return ""
	}

	return *s.IpsecLocalIp
}

func (s *NetworkConf) GetIpsecPeerIp() string {
	if s == nil {
		return ""
	}

	return *s.IpsecPeerIp
}

func (s *NetworkConf) GetIpsecPfs() bool {
	if s == nil {
		return false
	}

	return *s.IpsecPfs
}

func (s *NetworkConf) GetIpsecProfile() string {
	if s == nil {
		return ""
	}

	return *s.IpsecProfile
}

func (s *NetworkConf) GetIpsecRemoteIdentifier() string {
	if s == nil {
		return ""
	}

	return *s.IpsecRemoteIdentifier
}

func (s *NetworkConf) GetIpsecRemoteIdentifierEnabled() bool {
	if s == nil {
		return false
	}

	return *s.IpsecRemoteIdentifierEnabled
}

func (s *NetworkConf) GetIpsecSeparateIkev2Networks() bool {
	if s == nil {
		return false
	}

	return *s.IpsecSeparateIkev2Networks
}

func (s *NetworkConf) GetIpsecTunnelIp() string {
	if s == nil {
		return ""
	}

	return *s.IpsecTunnelIp
}

func (s *NetworkConf) GetIpsecTunnelIpEnabled() bool {
	if s == nil {
		return false
	}

	return *s.IpsecTunnelIpEnabled
}

func (s *NetworkConf) GetIpv6ClientAddressAssignment() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6ClientAddressAssignment
}

func (s *NetworkConf) GetIpv6InterfaceType() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6InterfaceType
}

func (s *NetworkConf) GetIpv6PdAutoPrefixidEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Ipv6PdAutoPrefixidEnabled
}

func (s *NetworkConf) GetIpv6PdInterface() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6PdInterface
}

func (s *NetworkConf) GetIpv6PdPrefixid() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6PdPrefixid
}

func (s *NetworkConf) GetIpv6PdStart() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6PdStart
}

func (s *NetworkConf) GetIpv6PdStop() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6PdStop
}

func (s *NetworkConf) GetIpv6RaEnabled() bool {
	if s == nil {
		return false
	}

	return *s.Ipv6RaEnabled
}

func (s *NetworkConf) GetIpv6RaPreferredLifetime() float64 {
	if s == nil {
		return 0
	}

	return *s.Ipv6RaPreferredLifetime
}

func (s *NetworkConf) GetIpv6RaPriority() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6RaPriority
}

func (s *NetworkConf) GetIpv6RaValidLifetime() float64 {
	if s == nil {
		return 0
	}

	return *s.Ipv6RaValidLifetime
}

func (s *NetworkConf) GetIpv6SettingPreference() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6SettingPreference
}

func (s *NetworkConf) GetIpv6SingleNetworkInterface() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6SingleNetworkInterface
}

func (s *NetworkConf) GetIpv6Subnet() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6Subnet
}

func (s *NetworkConf) GetIpv6WanDelegationType() string {
	if s == nil {
		return ""
	}

	return *s.Ipv6WanDelegationType
}

func (s *NetworkConf) GetIsNat() bool {
	if s == nil {
		return false
	}

	return *s.IsNat
}

func (s *NetworkConf) GetL2TpAllowWeakCiphers() bool {
	if s == nil {
		return false
	}

	return *s.L2TpAllowWeakCiphers
}

func (s *NetworkConf) GetL2TpInterface() string {
	if s == nil {
		return ""
	}

	return *s.L2TpInterface
}

func (s *NetworkConf) GetL2TpLocalWanIp() string {
	if s == nil {
		return ""
	}

	return *s.L2TpLocalWanIp
}

func (s *NetworkConf) GetLocalPort() float64 {
	if s == nil {
		return 0
	}

	return *s.LocalPort
}

func (s *NetworkConf) GetLteLanEnabled() bool {
	if s == nil {
		return false
	}

	return *s.LteLanEnabled
}

func (s *NetworkConf) GetMacOverride() string {
	if s == nil {
		return ""
	}

	return *s.MacOverride
}

func (s *NetworkConf) GetMacOverrideEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MacOverrideEnabled
}

func (s *NetworkConf) GetMdnsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.MdnsEnabled
}

func (s *NetworkConf) GetName() string {
	if s == nil {
		return ""
	}

	return *s.Name
}

func (s *NetworkConf) GetNatOutboundIpAddresses() []NetworkConfNatOutboundIpAddresses {
	if s == nil || s.NatOutboundIpAddresses == nil {
		return nil
	}

	return *s.NatOutboundIpAddresses
}

func (s *NetworkConf) GetNetworkIsolationEnabled() bool {
	if s == nil {
		return false
	}

	return *s.NetworkIsolationEnabled
}

func (s *NetworkConf) GetNetworkgroup() string {
	if s == nil {
		return ""
	}

	return *s.Networkgroup
}

func (s *NetworkConf) GetOpenvpnConfiguration() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnConfiguration
}

func (s *NetworkConf) GetOpenvpnConfigurationFilename() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnConfigurationFilename
}

func (s *NetworkConf) GetOpenvpnEncryptionCipher() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnEncryptionCipher
}

func (s *NetworkConf) GetOpenvpnInterface() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnInterface
}

func (s *NetworkConf) GetOpenvpnLocalAddress() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnLocalAddress
}

func (s *NetworkConf) GetOpenvpnLocalPort() float64 {
	if s == nil {
		return 0
	}

	return *s.OpenvpnLocalPort
}

func (s *NetworkConf) GetOpenvpnLocalWanIp() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnLocalWanIp
}

func (s *NetworkConf) GetOpenvpnMode() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnMode
}

func (s *NetworkConf) GetOpenvpnRemoteAddress() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnRemoteAddress
}

func (s *NetworkConf) GetOpenvpnRemoteHost() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnRemoteHost
}

func (s *NetworkConf) GetOpenvpnRemotePort() float64 {
	if s == nil {
		return 0
	}

	return *s.OpenvpnRemotePort
}

func (s *NetworkConf) GetOpenvpnUsername() string {
	if s == nil {
		return ""
	}

	return *s.OpenvpnUsername
}

func (s *NetworkConf) GetPptpcRequireMppe() bool {
	if s == nil {
		return false
	}

	return *s.PptpcRequireMppe
}

func (s *NetworkConf) GetPptpcRouteDistance() int64 {
	if s == nil {
		return 0
	}

	return *s.PptpcRouteDistance
}

func (s *NetworkConf) GetPptpcServerIp() string {
	if s == nil {
		return ""
	}

	return *s.PptpcServerIp
}

func (s *NetworkConf) GetPptpcUsername() string {
	if s == nil {
		return ""
	}

	return *s.PptpcUsername
}

func (s *NetworkConf) GetPriority() int64 {
	if s == nil {
		return 0
	}

	return *s.Priority
}

func (s *NetworkConf) GetPurpose() string {
	if s == nil {
		return ""
	}

	return *s.Purpose
}

func (s *NetworkConf) GetRadiusprofileId() string {
	if s == nil {
		return ""
	}

	return *s.RadiusprofileId
}

func (s *NetworkConf) GetRemoteSiteId() string {
	if s == nil {
		return ""
	}

	return *s.RemoteSiteId
}

func (s *NetworkConf) GetRemoteSiteSubnets() []string {
	if s == nil || s.RemoteSiteSubnets == nil {
		return nil
	}

	return *s.RemoteSiteSubnets
}

func (s *NetworkConf) GetRemoteVpnDynamicSubnetsEnabled() bool {
	if s == nil {
		return false
	}

	return *s.RemoteVpnDynamicSubnetsEnabled
}

func (s *NetworkConf) GetRemoteVpnSubnets() []string {
	if s == nil || s.RemoteVpnSubnets == nil {
		return nil
	}

	return *s.RemoteVpnSubnets
}

func (s *NetworkConf) GetReportWanEvent() bool {
	if s == nil {
		return false
	}

	return *s.ReportWanEvent
}

func (s *NetworkConf) GetRequireMschapv2() bool {
	if s == nil {
		return false
	}

	return *s.RequireMschapv2
}

func (s *NetworkConf) GetRouteDistance() int64 {
	if s == nil {
		return 0
	}

	return *s.RouteDistance
}

func (s *NetworkConf) GetSettingPreference() string {
	if s == nil {
		return ""
	}

	return *s.SettingPreference
}

func (s *NetworkConf) GetSingleNetworkLan() string {
	if s == nil {
		return ""
	}

	return *s.SingleNetworkLan
}

func (s *NetworkConf) GetUidPolicyEnabled() bool {
	if s == nil {
		return false
	}

	return *s.UidPolicyEnabled
}

func (s *NetworkConf) GetUidPolicyName() string {
	if s == nil {
		return ""
	}

	return *s.UidPolicyName
}

func (s *NetworkConf) GetUidPublicGatewayPort() float64 {
	if s == nil {
		return 0
	}

	return *s.UidPublicGatewayPort
}

func (s *NetworkConf) GetUidTrafficRulesAllowedIpsAndHostnames() []string {
	if s == nil || s.UidTrafficRulesAllowedIpsAndHostnames == nil {
		return nil
	}

	return *s.UidTrafficRulesAllowedIpsAndHostnames
}

func (s *NetworkConf) GetUidTrafficRulesEnabled() bool {
	if s == nil {
		return false
	}

	return *s.UidTrafficRulesEnabled
}

func (s *NetworkConf) GetUidVpnCustomRouting() []string {
	if s == nil || s.UidVpnCustomRouting == nil {
		return nil
	}

	return *s.UidVpnCustomRouting
}

func (s *NetworkConf) GetUidVpnDefaultDnsSuffix() string {
	if s == nil {
		return ""
	}

	return *s.UidVpnDefaultDnsSuffix
}

func (s *NetworkConf) GetUidVpnMasqueradeEnabled() bool {
	if s == nil {
		return false
	}

	return *s.UidVpnMasqueradeEnabled
}

func (s *NetworkConf) GetUidVpnMaxConnectionTimeSeconds() int64 {
	if s == nil {
		return 0
	}

	return *s.UidVpnMaxConnectionTimeSeconds
}

func (s *NetworkConf) GetUidVpnSyncPublicIp() bool {
	if s == nil {
		return false
	}

	return *s.UidVpnSyncPublicIp
}

func (s *NetworkConf) GetUidVpnType() string {
	if s == nil {
		return ""
	}

	return *s.UidVpnType
}

func (s *NetworkConf) GetUidWorkspaceUrl() string {
	if s == nil {
		return ""
	}

	return *s.UidWorkspaceUrl
}

func (s *NetworkConf) GetUpnpLanEnabled() bool {
	if s == nil {
		return false
	}

	return *s.UpnpLanEnabled
}

func (s *NetworkConf) GetUsergroupId() string {
	if s == nil {
		return ""
	}

	return *s.UsergroupId
}

func (s *NetworkConf) GetVlan() float64 {
	if s == nil {
		return 0
	}

	return *s.Vlan
}

func (s *NetworkConf) GetVlanEnabled() bool {
	if s == nil {
		return false
	}

	return *s.VlanEnabled
}

func (s *NetworkConf) GetVpnClientConfigurationRemoteIpOverride() string {
	if s == nil {
		return ""
	}

	return *s.VpnClientConfigurationRemoteIpOverride
}

func (s *NetworkConf) GetVpnClientConfigurationRemoteIpOverrideEnabled() bool {
	if s == nil {
		return false
	}

	return *s.VpnClientConfigurationRemoteIpOverrideEnabled
}

func (s *NetworkConf) GetVpnClientDefaultRoute() bool {
	if s == nil {
		return false
	}

	return *s.VpnClientDefaultRoute
}

func (s *NetworkConf) GetVpnClientPullDns() bool {
	if s == nil {
		return false
	}

	return *s.VpnClientPullDns
}

func (s *NetworkConf) GetVpnProtocol() string {
	if s == nil {
		return ""
	}

	return *s.VpnProtocol
}

func (s *NetworkConf) GetVpnType() string {
	if s == nil {
		return ""
	}

	return *s.VpnType
}

func (s *NetworkConf) GetVrrpIpSubnetGw1() string {
	if s == nil {
		return ""
	}

	return *s.VrrpIpSubnetGw1
}

func (s *NetworkConf) GetVrrpIpSubnetGw2() string {
	if s == nil {
		return ""
	}

	return *s.VrrpIpSubnetGw2
}

func (s *NetworkConf) GetVrrpVrid() int64 {
	if s == nil {
		return 0
	}

	return *s.VrrpVrid
}

func (s *NetworkConf) GetWanDhcpCos() int64 {
	if s == nil {
		return 0
	}

	return *s.WanDhcpCos
}

func (s *NetworkConf) GetWanDhcpOptions() []NetworkConfWanDhcpOptions {
	if s == nil || s.WanDhcpOptions == nil {
		return nil
	}

	return *s.WanDhcpOptions
}

func (s *NetworkConf) GetWanDhcpv6PdSize() int64 {
	if s == nil {
		return 0
	}

	return *s.WanDhcpv6PdSize
}

func (s *NetworkConf) GetWanDns1() string {
	if s == nil {
		return ""
	}

	return *s.WanDns1
}

func (s *NetworkConf) GetWanDns2() string {
	if s == nil {
		return ""
	}

	return *s.WanDns2
}

func (s *NetworkConf) GetWanDns3() string {
	if s == nil {
		return ""
	}

	return *s.WanDns3
}

func (s *NetworkConf) GetWanDns4() string {
	if s == nil {
		return ""
	}

	return *s.WanDns4
}

func (s *NetworkConf) GetWanDnsPreference() string {
	if s == nil {
		return ""
	}

	return *s.WanDnsPreference
}

func (s *NetworkConf) GetWanDsliteRemoteHost() string {
	if s == nil {
		return ""
	}

	return *s.WanDsliteRemoteHost
}

func (s *NetworkConf) GetWanEgressQos() int64 {
	if s == nil {
		return 0
	}

	return *s.WanEgressQos
}

func (s *NetworkConf) GetWanGateway() string {
	if s == nil {
		return ""
	}

	return *s.WanGateway
}

func (s *NetworkConf) GetWanGatewayV6() string {
	if s == nil {
		return ""
	}

	return *s.WanGatewayV6
}

func (s *NetworkConf) GetWanIp() string {
	if s == nil {
		return ""
	}

	return *s.WanIp
}

func (s *NetworkConf) GetWanIpAliases() []string {
	if s == nil || s.WanIpAliases == nil {
		return nil
	}

	return *s.WanIpAliases
}

func (s *NetworkConf) GetWanIpv6() string {
	if s == nil {
		return ""
	}

	return *s.WanIpv6
}

func (s *NetworkConf) GetWanIpv6Dns1() string {
	if s == nil {
		return ""
	}

	return *s.WanIpv6Dns1
}

func (s *NetworkConf) GetWanIpv6Dns2() string {
	if s == nil {
		return ""
	}

	return *s.WanIpv6Dns2
}

func (s *NetworkConf) GetWanIpv6DnsPreference() string {
	if s == nil {
		return ""
	}

	return *s.WanIpv6DnsPreference
}

func (s *NetworkConf) GetWanLoadBalanceType() string {
	if s == nil {
		return ""
	}

	return *s.WanLoadBalanceType
}

func (s *NetworkConf) GetWanLoadBalanceWeight() int64 {
	if s == nil {
		return 0
	}

	return *s.WanLoadBalanceWeight
}

func (s *NetworkConf) GetWanNetmask() string {
	if s == nil {
		return ""
	}

	return *s.WanNetmask
}

func (s *NetworkConf) GetWanNetworkgroup() string {
	if s == nil {
		return ""
	}

	return *s.WanNetworkgroup
}

func (s *NetworkConf) GetWanPppoePasswordEnabled() bool {
	if s == nil {
		return false
	}

	return *s.WanPppoePasswordEnabled
}

func (s *NetworkConf) GetWanPppoeUsernameEnabled() bool {
	if s == nil {
		return false
	}

	return *s.WanPppoeUsernameEnabled
}

func (s *NetworkConf) GetWanPrefixlen() int64 {
	if s == nil {
		return 0
	}

	return *s.WanPrefixlen
}

func (s *NetworkConf) GetWanProviderCapabilities() *NetworkConfWanProviderCapabilities {
	if s == nil || s.WanProviderCapabilities == nil {
		return nil
	}

	return s.WanProviderCapabilities
}

func (s *NetworkConf) GetWanSmartqDownRate() int64 {
	if s == nil {
		return 0
	}

	return *s.WanSmartqDownRate
}

func (s *NetworkConf) GetWanSmartqEnabled() bool {
	if s == nil {
		return false
	}

	return *s.WanSmartqEnabled
}

func (s *NetworkConf) GetWanSmartqUpRate() int64 {
	if s == nil {
		return 0
	}

	return *s.WanSmartqUpRate
}

func (s *NetworkConf) GetWanType() string {
	if s == nil {
		return ""
	}

	return *s.WanType
}

func (s *NetworkConf) GetWanTypeV6() string {
	if s == nil {
		return ""
	}

	return *s.WanTypeV6
}

func (s *NetworkConf) GetWanUsername() string {
	if s == nil {
		return ""
	}

	return *s.WanUsername
}

func (s *NetworkConf) GetWanVlan() float64 {
	if s == nil {
		return 0
	}

	return *s.WanVlan
}

func (s *NetworkConf) GetWanVlanEnabled() bool {
	if s == nil {
		return false
	}

	return *s.WanVlanEnabled
}

func (s *NetworkConf) GetWireguardClientConfigurationFile() string {
	if s == nil {
		return ""
	}

	return *s.WireguardClientConfigurationFile
}

func (s *NetworkConf) GetWireguardClientConfigurationFilename() string {
	if s == nil {
		return ""
	}

	return *s.WireguardClientConfigurationFilename
}

func (s *NetworkConf) GetWireguardClientMode() string {
	if s == nil {
		return ""
	}

	return *s.WireguardClientMode
}

func (s *NetworkConf) GetWireguardClientPeerIp() string {
	if s == nil {
		return ""
	}

	return *s.WireguardClientPeerIp
}

func (s *NetworkConf) GetWireguardClientPeerPort() float64 {
	if s == nil {
		return 0
	}

	return *s.WireguardClientPeerPort
}

func (s *NetworkConf) GetWireguardClientPeerPublicKey() string {
	if s == nil {
		return ""
	}

	return *s.WireguardClientPeerPublicKey
}

func (s *NetworkConf) GetWireguardClientPresharedKey() string {
	if s == nil {
		return ""
	}

	return *s.WireguardClientPresharedKey
}

func (s *NetworkConf) GetWireguardClientPresharedKeyEnabled() bool {
	if s == nil {
		return false
	}

	return *s.WireguardClientPresharedKeyEnabled
}

func (s *NetworkConf) GetWireguardInterface() string {
	if s == nil {
		return ""
	}

	return *s.WireguardInterface
}

func (s *NetworkConf) GetWireguardLocalWanIp() string {
	if s == nil {
		return ""
	}

	return *s.WireguardLocalWanIp
}

func (s *NetworkConf) GetWireguardPublicKey() string {
	if s == nil {
		return ""
	}

	return *s.WireguardPublicKey
}

func (s *NetworkConf) GetXAuthKey() string {
	if s == nil {
		return ""
	}

	return *s.XAuthKey
}

func (s *NetworkConf) GetXCaCrt() string {
	if s == nil {
		return ""
	}

	return *s.XCaCrt
}

func (s *NetworkConf) GetXCaKey() string {
	if s == nil {
		return ""
	}

	return *s.XCaKey
}

func (s *NetworkConf) GetXDhKey() string {
	if s == nil {
		return ""
	}

	return *s.XDhKey
}

func (s *NetworkConf) GetXIpsecPreSharedKey() string {
	if s == nil {
		return ""
	}

	return *s.XIpsecPreSharedKey
}

func (s *NetworkConf) GetXOpenvpnPassword() string {
	if s == nil {
		return ""
	}

	return *s.XOpenvpnPassword
}

func (s *NetworkConf) GetXOpenvpnSharedSecretKey() string {
	if s == nil {
		return ""
	}

	return *s.XOpenvpnSharedSecretKey
}

func (s *NetworkConf) GetXPptpcPassword() string {
	if s == nil {
		return ""
	}

	return *s.XPptpcPassword
}

func (s *NetworkConf) GetXServerCrt() string {
	if s == nil {
		return ""
	}

	return *s.XServerCrt
}

func (s *NetworkConf) GetXServerKey() string {
	if s == nil {
		return ""
	}

	return *s.XServerKey
}

func (s *NetworkConf) GetXSharedClientCrt() string {
	if s == nil {
		return ""
	}

	return *s.XSharedClientCrt
}

func (s *NetworkConf) GetXSharedClientKey() string {
	if s == nil {
		return ""
	}

	return *s.XSharedClientKey
}

func (s *NetworkConf) GetXWanPassword() string {
	if s == nil {
		return ""
	}

	return *s.XWanPassword
}

func (s *NetworkConf) GetXWireguardPrivateKey() string {
	if s == nil {
		return ""
	}

	return *s.XWireguardPrivateKey
}

type NetworkConfIgmpQuerierSwitches struct {
	QuerierAddress *string `json:"querier_address,omitempty"`
	SwitchMac      *string `json:"switch_mac,omitempty"`
}

func (s *NetworkConfIgmpQuerierSwitches) GetQuerierAddress() string {
	if s == nil {
		return ""
	}

	return *s.QuerierAddress
}

func (s *NetworkConfIgmpQuerierSwitches) GetSwitchMac() string {
	if s == nil {
		return ""
	}

	return *s.SwitchMac
}

type NetworkConfNatOutboundIpAddresses struct {
	IpAddress       *string   `json:"ip_address,omitempty"`
	IpAddressPool   *[]string `json:"ip_address_pool,omitempty"`
	Mode            *string   `json:"mode,omitempty"`
	WanNetworkGroup *string   `json:"wan_network_group,omitempty"`
}

func (s *NetworkConfNatOutboundIpAddresses) GetIpAddress() string {
	if s == nil {
		return ""
	}

	return *s.IpAddress
}

func (s *NetworkConfNatOutboundIpAddresses) GetIpAddressPool() []string {
	if s == nil || s.IpAddressPool == nil {
		return nil
	}

	return *s.IpAddressPool
}

func (s *NetworkConfNatOutboundIpAddresses) GetMode() string {
	if s == nil {
		return ""
	}

	return *s.Mode
}

func (s *NetworkConfNatOutboundIpAddresses) GetWanNetworkGroup() string {
	if s == nil {
		return ""
	}

	return *s.WanNetworkGroup
}

type NetworkConfWanDhcpOptions struct {
	OptionNumber *float64 `json:"optionNumber,omitempty"`
	Value        *string  `json:"value,omitempty"`
}

func (s *NetworkConfWanDhcpOptions) GetOptionNumber() float64 {
	if s == nil {
		return 0
	}

	return *s.OptionNumber
}

func (s *NetworkConfWanDhcpOptions) GetValue() string {
	if s == nil {
		return ""
	}

	return *s.Value
}

type NetworkConfWanProviderCapabilities struct {
	DownloadKilobitsPerSecond *int64 `json:"download_kilobits_per_second,omitempty"`
	UploadKilobitsPerSecond   *int64 `json:"upload_kilobits_per_second,omitempty"`
}

func (s *NetworkConfWanProviderCapabilities) GetDownloadKilobitsPerSecond() int64 {
	if s == nil {
		return 0
	}

	return *s.DownloadKilobitsPerSecond
}

func (s *NetworkConfWanProviderCapabilities) GetUploadKilobitsPerSecond() int64 {
	if s == nil {
		return 0
	}

	return *s.UploadKilobitsPerSecond
}

type responseBodyNetworkConf struct {
	Metadata json.RawMessage `json:"meta"`
	Payload  []NetworkConf   `json:"data"`
}

func (c *Client) CreateNetworkConf(ctx context.Context, site string, data *NetworkConf) (*NetworkConf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "networkconf")
	req, err := c.NewRequest(ctx, http.MethodPost, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyNetworkConf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to create networkconf: %w`, err)
	}

	var item NetworkConf
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to create NetworkConf`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) DeleteNetworkConf(ctx context.Context, site string, id string) (*http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "networkconf", id)
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, err
	}

	var body responseBodyNetworkConf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return resp, fmt.Errorf(`unable to delete NetworkConf: %w`, err)
	}

	return nil, nil
}

func (c *Client) GetNetworkConf(ctx context.Context, site, id string) (*NetworkConf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "networkconf", id)
	req, err := c.NewRequest(ctx, http.MethodDelete, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyNetworkConf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get NetworkConf: %w`, err)
	}

	var item NetworkConf
	switch len(body.Payload) {
	case 0:
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}

func (c *Client) ListNetworkConf(ctx context.Context, site string) ([]NetworkConf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "networkconf")
	req, err := c.NewRequest(ctx, http.MethodGet, endpointPath, nil)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyNetworkConf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to get NetworkConf: %w`, err)
	}

	return body.Payload, resp, nil
}

func (c *Client) UpdateNetworkConf(ctx context.Context, site string, data *NetworkConf) (*NetworkConf, *http.Response, error) {
	endpointPath := path.Join("api/s/", site, "rest", "networkconf", *data.ID)
	req, err := c.NewRequest(ctx, http.MethodPut, endpointPath, data)
	if err != nil {
		return nil, nil, err
	}

	var body responseBodyNetworkConf
	resp, err := c.Do(ctx, req, &body)
	if err != nil {
		return nil, resp, fmt.Errorf(`unable to update networkconf: %w`, err)
	}

	var item NetworkConf
	switch len(body.Payload) {
	case 0:
		err = errors.New(`failed to update NetworkConf`)
	case 1:
		item = body.Payload[0]
	default:
		err = fmt.Errorf("unexpected number of results: %v", len(body.Payload))
	}

	return &item, resp, err
}
