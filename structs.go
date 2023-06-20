package main

import "encoding/xml"

// MultiRoutingEngineResults was generated 2023-06-13 15:03:43 by https://xml-to-go.github.io/ in Ukraine.
type MultiRoutingEngineResults struct {
	XMLName                xml.Name `xml:"multi-routing-engine-results"`
	Text                   string   `xml:",chardata"`
	MultiRoutingEngineItem struct {
		Text             string `xml:",chardata"`
		ReName           string `xml:"re-name"`
		SecurityPolicies struct {
			Text                      string `xml:",chardata"`
			Style                     string `xml:"style,attr"`
			DefaultPolicy             string `xml:"default-policy"`
			DefaultPolicyLogProfileID string `xml:"default-policy-log-profile-id"`
			PreIDDefaultPolicy        string `xml:"pre-id-default-policy"`
			SecurityContext           []struct {
				Text               string `xml:",chardata"`
				ContextInformation struct {
					Text                string `xml:",chardata"`
					SourceZoneName      string `xml:"source-zone-name"`
					DestinationZoneName string `xml:"destination-zone-name"`
					GlobalContext       string `xml:"global-context"`
				} `xml:"context-information"`
				Policies []struct {
					Text              string `xml:",chardata"`
					PolicyInformation struct {
						Text                  string `xml:",chardata"`
						PolicyName            string `xml:"policy-name"`
						PolicyState           string `xml:"policy-state"`
						PolicyIdentifier      string `xml:"policy-identifier"`
						ScopePolicyIdentifier string `xml:"scope-policy-identifier"`
						PolicySequenceNumber  string `xml:"policy-sequence-number"`
						LogProfileID          string `xml:"log-profile-id"`
						MultipleSourceZones   struct {
							Text       string `xml:",chardata"`
							Style      string `xml:"style,attr"`
							SourceZone []struct {
								Text           string `xml:",chardata"`
								SourceZoneName string `xml:"source-zone-name"`
							} `xml:"source-zone"`
						} `xml:"multiple-source-zones"`
						MultipleDestinationZones struct {
							Text            string `xml:",chardata"`
							Style           string `xml:"style,attr"`
							DestinationZone []struct {
								Text                string `xml:",chardata"`
								DestinationZoneName string `xml:"destination-zone-name"`
							} `xml:"destination-zone"`
						} `xml:"multiple-destination-zones"`
						SourceVrfs struct {
							Text      string `xml:",chardata"`
							Style     string `xml:"style,attr"`
							SourceVrf struct {
								Text          string `xml:",chardata"`
								SourceVrfName string `xml:"source-vrf-name"`
							} `xml:"source-vrf"`
						} `xml:"source-vrfs"`
						DestinationVrfs struct {
							Text           string `xml:",chardata"`
							Style          string `xml:"style,attr"`
							DestinationVrf struct {
								Text               string `xml:",chardata"`
								DestinationVrfName string `xml:"destination-vrf-name"`
							} `xml:"destination-vrf"`
						} `xml:"destination-vrfs"`
						SourceAddresses struct {
							Text          string `xml:",chardata"`
							Style         string `xml:"style,attr"`
							SourceAddress []struct {
								Text        string `xml:",chardata"`
								AddressName string `xml:"address-name"`
							} `xml:"source-address"`
						} `xml:"source-addresses"`
						DestinationAddresses struct {
							Text               string `xml:",chardata"`
							Style              string `xml:"style,attr"`
							DestinationAddress []struct {
								Text        string `xml:",chardata"`
								AddressName string `xml:"address-name"`
							} `xml:"destination-address"`
						} `xml:"destination-addresses"`
						Applications struct {
							Text        string `xml:",chardata"`
							Style       string `xml:"style,attr"`
							Application []struct {
								Text            string `xml:",chardata"`
								ApplicationName string `xml:"application-name"`
							} `xml:"application"`
						} `xml:"applications"`
						SourceIdentities struct {
							Text  string `xml:",chardata"`
							Style string `xml:"style,attr"`
						} `xml:"source-identities"`
						SourceIdentitiesFeeds struct {
							Text               string `xml:",chardata"`
							Style              string `xml:"style,attr"`
							SourceIdentityFeed struct {
								Text     string `xml:",chardata"`
								FeedName string `xml:"feed-name"`
							} `xml:"source-identity-feed"`
						} `xml:"source-identities-feeds"`
						DestinationIdentitiesFeeds struct {
							Text                    string `xml:",chardata"`
							Style                   string `xml:"style,attr"`
							DestinationIdentityFeed struct {
								Text     string `xml:",chardata"`
								FeedName string `xml:"feed-name"`
							} `xml:"destination-identity-feed"`
						} `xml:"destination-identities-feeds"`
						PolicyAction struct {
							Text             string `xml:",chardata"`
							ActionType       string `xml:"action-type"`
							PolicyTcpOptions struct {
								Text                          string `xml:",chardata"`
								PolicyTcpOptionsSynCheck      string `xml:"policy-tcp-options-syn-check"`
								PolicyTcpOptionsSequenceCheck string `xml:"policy-tcp-options-sequence-check"`
								PolicyTcpOptionsWindowScale   string `xml:"policy-tcp-options-window-scale"`
							} `xml:"policy-tcp-options"`
							Log string `xml:"log"`
						} `xml:"policy-action"`
						PolicyApplicationServices string `xml:"policy-application-services"`
					} `xml:"policy-information"`
				} `xml:"policies"`
			} `xml:"security-context"`
		} `xml:"security-policies"`
	} `xml:"multi-routing-engine-item"`
}

type SecurityPolicies struct {
	XMLName                   xml.Name `xml:"security-policies"`
	Text                      string   `xml:",chardata"`
	Style                     string   `xml:"style,attr"`
	DefaultPolicy             string   `xml:"default-policy"`
	DefaultPolicyLogProfileID string   `xml:"default-policy-log-profile-id"`
	PreIDDefaultPolicy        string   `xml:"pre-id-default-policy"`
	SecurityContext           []struct {
		Text               string `xml:",chardata"`
		ContextInformation struct {
			Text                string `xml:",chardata"`
			SourceZoneName      string `xml:"source-zone-name"`
			DestinationZoneName string `xml:"destination-zone-name"`
			GlobalContext       string `xml:"global-context"`
		} `xml:"context-information"`
		Policies []struct {
			Text              string `xml:",chardata"`
			PolicyInformation struct {
				Text                  string `xml:",chardata"`
				PolicyName            string `xml:"policy-name"`
				PolicyState           string `xml:"policy-state"`
				PolicyIdentifier      string `xml:"policy-identifier"`
				ScopePolicyIdentifier string `xml:"scope-policy-identifier"`
				PolicySequenceNumber  string `xml:"policy-sequence-number"`
				LogProfileID          string `xml:"log-profile-id"`
				MultipleSourceZones   struct {
					Text       string `xml:",chardata"`
					Style      string `xml:"style,attr"`
					SourceZone []struct {
						Text           string `xml:",chardata"`
						SourceZoneName string `xml:"source-zone-name"`
					} `xml:"source-zone"`
				} `xml:"multiple-source-zones"`
				MultipleDestinationZones struct {
					Text            string `xml:",chardata"`
					Style           string `xml:"style,attr"`
					DestinationZone []struct {
						Text                string `xml:",chardata"`
						DestinationZoneName string `xml:"destination-zone-name"`
					} `xml:"destination-zone"`
				} `xml:"multiple-destination-zones"`
				SourceVrfs struct {
					Text      string `xml:",chardata"`
					Style     string `xml:"style,attr"`
					SourceVrf struct {
						Text          string `xml:",chardata"`
						SourceVrfName string `xml:"source-vrf-name"`
					} `xml:"source-vrf"`
				} `xml:"source-vrfs"`
				DestinationVrfs struct {
					Text           string `xml:",chardata"`
					Style          string `xml:"style,attr"`
					DestinationVrf struct {
						Text               string `xml:",chardata"`
						DestinationVrfName string `xml:"destination-vrf-name"`
					} `xml:"destination-vrf"`
				} `xml:"destination-vrfs"`
				SourceAddresses struct {
					Text          string `xml:",chardata"`
					Style         string `xml:"style,attr"`
					SourceAddress []struct {
						Text        string `xml:",chardata"`
						AddressName string `xml:"address-name"`
					} `xml:"source-address"`
				} `xml:"source-addresses"`
				DestinationAddresses struct {
					Text               string `xml:",chardata"`
					Style              string `xml:"style,attr"`
					DestinationAddress []struct {
						Text        string `xml:",chardata"`
						AddressName string `xml:"address-name"`
					} `xml:"destination-address"`
				} `xml:"destination-addresses"`
				Applications struct {
					Text        string `xml:",chardata"`
					Style       string `xml:"style,attr"`
					Application []struct {
						Text            string `xml:",chardata"`
						ApplicationName string `xml:"application-name"`
					} `xml:"application"`
				} `xml:"applications"`
				SourceIdentities struct {
					Text  string `xml:",chardata"`
					Style string `xml:"style,attr"`
				} `xml:"source-identities"`
				SourceIdentitiesFeeds struct {
					Text               string `xml:",chardata"`
					Style              string `xml:"style,attr"`
					SourceIdentityFeed struct {
						Text     string `xml:",chardata"`
						FeedName string `xml:"feed-name"`
					} `xml:"source-identity-feed"`
				} `xml:"source-identities-feeds"`
				DestinationIdentitiesFeeds struct {
					Text                    string `xml:",chardata"`
					Style                   string `xml:"style,attr"`
					DestinationIdentityFeed struct {
						Text     string `xml:",chardata"`
						FeedName string `xml:"feed-name"`
					} `xml:"destination-identity-feed"`
				} `xml:"destination-identities-feeds"`
				PolicyAction struct {
					Text             string `xml:",chardata"`
					ActionType       string `xml:"action-type"`
					PolicyTcpOptions struct {
						Text                          string `xml:",chardata"`
						PolicyTcpOptionsSynCheck      string `xml:"policy-tcp-options-syn-check"`
						PolicyTcpOptionsSequenceCheck string `xml:"policy-tcp-options-sequence-check"`
						PolicyTcpOptionsWindowScale   string `xml:"policy-tcp-options-window-scale"`
					} `xml:"policy-tcp-options"`
				} `xml:"policy-action"`
				PolicyApplicationServices string `xml:"policy-application-services"`
			} `xml:"policy-information"`
		} `xml:"policies"`
	} `xml:"security-context"`
}

type MultiRoutingEngineResultsDetails struct {
	XMLName                   xml.Name `xml:"rpc-reply"`
	Text                      string   `xml:",chardata"`
	Junos                     string   `xml:"junos,attr"`
	MessageID                 string   `xml:"message-id,attr"`
	Xmlns                     string   `xml:"xmlns,attr"`
	MultiRoutingEngineResults struct {
		Text                   string `xml:",chardata"`
		MultiRoutingEngineItem struct {
			Text             string `xml:",chardata"`
			ReName           string `xml:"re-name"`
			SecurityPolicies struct {
				Text                      string `xml:",chardata"`
				Style                     string `xml:"style,attr"`
				DefaultPolicy             string `xml:"default-policy"`
				DefaultPolicyLogProfileID string `xml:"default-policy-log-profile-id"`
				PreIDDefaultPolicy        string `xml:"pre-id-default-policy"`
				SecurityContext           []struct {
					Text               string `xml:",chardata"`
					ContextInformation struct {
						Text                string `xml:",chardata"`
						SourceZoneName      string `xml:"source-zone-name"`
						DestinationZoneName string `xml:"destination-zone-name"`
						GlobalContext       string `xml:"global-context"`
					} `xml:"context-information"`
					Policies []struct {
						Text              string `xml:",chardata"`
						PolicyInformation struct {
							Text         string `xml:",chardata"`
							PolicyName   string `xml:"policy-name"`
							PolicyAction struct {
								Text            string `xml:",chardata"`
								ActionType      string `xml:"action-type"`
								ServicesOffload string `xml:"services-offload"`
							} `xml:"policy-action"`
							PolicyState           string `xml:"policy-state"`
							PolicyIdentifier      string `xml:"policy-identifier"`
							ScopePolicyIdentifier string `xml:"scope-policy-identifier"`
							PolicyTypeInformation struct {
								Text             string `xml:",chardata"`
								PolicyTypeString string `xml:"policy-type-string"`
							} `xml:"policy-type-information"`
							PolicyDescription    string `xml:"policy-description"`
							PolicySequenceNumber string `xml:"policy-sequence-number"`
							ContextInformation   struct {
								Text                string `xml:",chardata"`
								SourceZoneName      string `xml:"source-zone-name"`
								DestinationZoneName string `xml:"destination-zone-name"`
							} `xml:"context-information"`
							MultipleSourceZones struct {
								Text       string `xml:",chardata"`
								Style      string `xml:"style,attr"`
								SourceZone []struct {
									Text           string `xml:",chardata"`
									SourceZoneName string `xml:"source-zone-name"`
								} `xml:"source-zone"`
							} `xml:"multiple-source-zones"`
							MultipleDestinationZones struct {
								Text            string `xml:",chardata"`
								Style           string `xml:"style,attr"`
								DestinationZone []struct {
									Text                string `xml:",chardata"`
									DestinationZoneName string `xml:"destination-zone-name"`
								} `xml:"destination-zone"`
							} `xml:"multiple-destination-zones"`
							SourceVrfs struct {
								Text      string `xml:",chardata"`
								Style     string `xml:"style,attr"`
								SourceVrf struct {
									Text          string `xml:",chardata"`
									SourceVrfName string `xml:"source-vrf-name"`
								} `xml:"source-vrf"`
							} `xml:"source-vrfs"`
							DestinationVrfs struct {
								Text           string `xml:",chardata"`
								Style          string `xml:"style,attr"`
								DestinationVrf struct {
									Text               string `xml:",chardata"`
									DestinationVrfName string `xml:"destination-vrf-name"`
								} `xml:"destination-vrf"`
							} `xml:"destination-vrfs"`
							SourceAddresses struct {
								Text          string `xml:",chardata"`
								Style         string `xml:"style,attr"`
								SourceAddress []struct {
									Text        string `xml:",chardata"`
									AddressName string `xml:"address-name"`
									Prefixes    struct {
										Text          string `xml:",chardata"`
										AddressPrefix string `xml:"address-prefix"`
									} `xml:"prefixes"`
									AddressRange struct {
										Text        string `xml:",chardata"`
										AddressLow  string `xml:"address-low"`
										AddressHigh string `xml:"address-high"`
									} `xml:"address-range"`
								} `xml:"source-address"`
							} `xml:"source-addresses"`
							DestinationAddresses struct {
								Text               string `xml:",chardata"`
								Style              string `xml:"style,attr"`
								DestinationAddress []struct {
									Text        string `xml:",chardata"`
									AddressName string `xml:"address-name"`
									Prefixes    struct {
										Text          string `xml:",chardata"`
										AddressPrefix string `xml:"address-prefix"`
									} `xml:"prefixes"`
								} `xml:"destination-address"`
							} `xml:"destination-addresses"`
							Applications struct {
								Text        string `xml:",chardata"`
								Style       string `xml:"style,attr"`
								Application []struct {
									Text            string `xml:",chardata"`
									ApplicationName string `xml:"application-name"`
									ApplicationTerm []struct {
										Text              string `xml:",chardata"`
										Protocol          string `xml:"protocol"`
										AlgName           string `xml:"alg-name"`
										InactivityTimeout string `xml:"inactivity-timeout"`
										SourcePortRange   struct {
											Text string `xml:",chardata"`
											Low  string `xml:"low"`
											High string `xml:"high"`
										} `xml:"source-port-range"`
										DestinationPortRange struct {
											Text   string `xml:",chardata"`
											Single string `xml:"single"`
											Low    string `xml:"low"`
											High   string `xml:"high"`
										} `xml:"destination-port-range"`
										IcmpInfo struct {
											Text     string `xml:",chardata"`
											IcmpType string `xml:"icmp-type"`
											IcmpCode string `xml:"icmp-code"`
										} `xml:"icmp-info"`
									} `xml:"application-term"`
								} `xml:"application"`
							} `xml:"applications"`
							PolicyDynamicApplications struct {
								Text  string `xml:",chardata"`
								Style string `xml:"style,attr"`
							} `xml:"policy-dynamic-applications"`
							PolicyURLCategories struct {
								Text  string `xml:",chardata"`
								Style string `xml:"style,attr"`
							} `xml:"policy-url-categories"`
							SourceIdentities struct {
								Text  string `xml:",chardata"`
								Style string `xml:"style,attr"`
							} `xml:"source-identities"`
							SourceIdentitiesFeeds struct {
								Text               string `xml:",chardata"`
								Style              string `xml:"style,attr"`
								SourceIdentityFeed struct {
									Text     string `xml:",chardata"`
									FeedName string `xml:"feed-name"`
								} `xml:"source-identity-feed"`
							} `xml:"source-identities-feeds"`
							DestinationIdentitiesFeeds struct {
								Text                    string `xml:",chardata"`
								Style                   string `xml:"style,attr"`
								DestinationIdentityFeed struct {
									Text     string `xml:",chardata"`
									FeedName string `xml:"feed-name"`
								} `xml:"destination-identity-feed"`
							} `xml:"destination-identities-feeds"`
							PolicyTcpOptions struct {
								Text                          string `xml:",chardata"`
								PolicyTcpOptionsSynCheck      string `xml:"policy-tcp-options-syn-check"`
								PolicyTcpOptionsSequenceCheck string `xml:"policy-tcp-options-sequence-check"`
								PolicyTcpOptionsWindowScale   string `xml:"policy-tcp-options-window-scale"`
							} `xml:"policy-tcp-options"`
							PolicyApplicationServices struct {
								Text      string `xml:",chardata"`
								Idp       string `xml:"idp"`
								UacPolicy string `xml:"uac-policy"`
							} `xml:"policy-application-services"`
							PolicyLog struct {
								Text               string `xml:",chardata"`
								LogSessionCreation string `xml:"log-session-creation"`
								LogSessionClose    string `xml:"log-session-close"`
							} `xml:"policy-log"`
							PolicyFeed      string `xml:"policy-feed"`
							PolicyScheduler string `xml:"policy-scheduler"`
						} `xml:"policy-information"`
					} `xml:"policies"`
				} `xml:"security-context"`
			} `xml:"security-policies"`
		} `xml:"multi-routing-engine-item"`
	} `xml:"multi-routing-engine-results"`
}

type SecurityPoliciesDetails struct {
	XMLName          xml.Name `xml:"rpc-reply"`
	Text             string   `xml:",chardata"`
	Junos            string   `xml:"junos,attr"`
	MessageID        string   `xml:"message-id,attr"`
	Xmlns            string   `xml:"xmlns,attr"`
	SecurityPolicies struct {
		Text            string `xml:",chardata"`
		Style           string `xml:"style,attr"`
		DefaultPolicy   string `xml:"default-policy"`
		SecurityContext []struct {
			Text               string `xml:",chardata"`
			ContextInformation struct {
				Text                string `xml:",chardata"`
				SourceZoneName      string `xml:"source-zone-name"`
				DestinationZoneName string `xml:"destination-zone-name"`
				GlobalContext       string `xml:"global-context"`
			} `xml:"context-information"`
			Policies []struct {
				Text              string `xml:",chardata"`
				PolicyInformation struct {
					Text         string `xml:",chardata"`
					PolicyName   string `xml:"policy-name"`
					PolicyAction struct {
						Text       string `xml:",chardata"`
						ActionType string `xml:"action-type"`
					} `xml:"policy-action"`
					PolicyState           string `xml:"policy-state"`
					PolicyIdentifier      string `xml:"policy-identifier"`
					ScopePolicyIdentifier string `xml:"scope-policy-identifier"`
					PolicyTypeInformation struct {
						Text             string `xml:",chardata"`
						PolicyTypeString string `xml:"policy-type-string"`
					} `xml:"policy-type-information"`
					PolicySequenceNumber string `xml:"policy-sequence-number"`
					ContextInformation   struct {
						Text                string `xml:",chardata"`
						SourceZoneName      string `xml:"source-zone-name"`
						DestinationZoneName string `xml:"destination-zone-name"`
					} `xml:"context-information"`
					MultipleSourceZones struct {
						Text       string `xml:",chardata"`
						Style      string `xml:"style,attr"`
						SourceZone []struct {
							Text           string `xml:",chardata"`
							SourceZoneName string `xml:"source-zone-name"`
						} `xml:"source-zone"`
					} `xml:"multiple-source-zones"`
					MultipleDestinationZones struct {
						Text            string `xml:",chardata"`
						Style           string `xml:"style,attr"`
						DestinationZone []struct {
							Text                string `xml:",chardata"`
							DestinationZoneName string `xml:"destination-zone-name"`
						} `xml:"destination-zone"`
					} `xml:"multiple-destination-zones"`
					SourceAddresses struct {
						Text          string `xml:",chardata"`
						Style         string `xml:"style,attr"`
						SourceAddress []struct {
							Text        string `xml:",chardata"`
							AddressName string `xml:"address-name"`
							Prefixes    struct {
								Text          string `xml:",chardata"`
								AddressPrefix string `xml:"address-prefix"`
							} `xml:"prefixes"`
						} `xml:"source-address"`
					} `xml:"source-addresses"`
					DestinationAddresses struct {
						Text               string `xml:",chardata"`
						Style              string `xml:"style,attr"`
						DestinationAddress []struct {
							Text        string `xml:",chardata"`
							AddressName string `xml:"address-name"`
							Prefixes    struct {
								Text          string `xml:",chardata"`
								AddressPrefix string `xml:"address-prefix"`
							} `xml:"prefixes"`
						} `xml:"destination-address"`
					} `xml:"destination-addresses"`
					Applications struct {
						Text        string `xml:",chardata"`
						Style       string `xml:"style,attr"`
						Application []struct {
							Text            string `xml:",chardata"`
							ApplicationName string `xml:"application-name"`
							ApplicationTerm struct {
								Text              string `xml:",chardata"`
								Protocol          string `xml:"protocol"`
								AlgName           string `xml:"alg-name"`
								InactivityTimeout string `xml:"inactivity-timeout"`
								SourcePortRange   struct {
									Text string `xml:",chardata"`
									Low  string `xml:"low"`
									High string `xml:"high"`
								} `xml:"source-port-range"`
								DestinationPortRange struct {
									Text string `xml:",chardata"`
									Low  string `xml:"low"`
									High string `xml:"high"`
								} `xml:"destination-port-range"`
								IcmpInfo struct {
									Text     string `xml:",chardata"`
									IcmpType string `xml:"icmp-type"`
									IcmpCode string `xml:"icmp-code"`
								} `xml:"icmp-info"`
							} `xml:"application-term"`
						} `xml:"application"`
					} `xml:"applications"`
					SourceIdentities struct {
						Text  string `xml:",chardata"`
						Style string `xml:"style,attr"`
					} `xml:"source-identities"`
					PolicyTcpOptions struct {
						Text                          string `xml:",chardata"`
						PolicyTcpOptionsSynCheck      string `xml:"policy-tcp-options-syn-check"`
						PolicyTcpOptionsSequenceCheck string `xml:"policy-tcp-options-sequence-check"`
					} `xml:"policy-tcp-options"`
					PolicyApplicationServices string `xml:"policy-application-services"`
					PolicyLog                 struct {
						Text               string `xml:",chardata"`
						LogSessionCreation string `xml:"log-session-creation"`
						LogSessionClose    string `xml:"log-session-close"`
					} `xml:"policy-log"`
					PolicyScheduler   string `xml:"policy-scheduler"`
					PolicyDescription string `xml:"policy-description"`
				} `xml:"policy-information"`
			} `xml:"policies"`
		} `xml:"security-context"`
	} `xml:"security-policies"`
}
