package main

type Question struct {
	Text          string   `json:"text"`
	Options       []string `json:"options"`
	CorrectAnswer int      `json:"-"`
}

var Questions = []Question{
	{
		Text: "What is the primary goal of a Microsoft Security Operations Analyst?",
		Options: []string{
			"To reduce organizational risk",
			"To increase sales",
			"To manage human resources",
			"To develop software",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "Which tool is used for threat mitigation in Microsoft 365 Defender?",
		Options: []string{
			"Azure Sentinel",
			"Microsoft Defender for Endpoint",
			"Microsoft Teams",
			"Power BI",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the purpose of Microsoft Defender for Cloud?",
		Options: []string{
			"To provide email protection",
			"To secure cloud resources across multiple cloud providers",
			"To manage Active Directory",
			"To create virtual machines",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which component of Microsoft 365 Defender focuses on email and collaboration security?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the main function of Microsoft Defender for Identity?",
		Options: []string{
			"To protect cloud applications",
			"To secure on-premises identity infrastructure",
			"To manage mobile devices",
			"To encrypt data at rest",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which tool is used for Security Information and Event Management (SIEM) in Microsoft's security ecosystem?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Intune",
			"Microsoft Sentinel",
			"Microsoft Exchange Online Protection",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the purpose of Microsoft Defender Vulnerability Management?",
		Options: []string{
			"To manage software licenses",
			"To discover, assess, and remediate vulnerabilities and misconfigurations",
			"To create backups of critical data",
			"To monitor network traffic",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Cloud provides recommendations to improve your security posture?",
		Options: []string{
			"Secure Score",
			"Just-In-Time VM Access",
			"Adaptive Network Hardening",
			"Regulatory Compliance",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud Apps?",
		Options: []string{
			"To protect on-premises servers",
			"To secure and manage cloud applications",
			"To monitor physical security in data centers",
			"To create virtual networks",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which component of Microsoft 365 Defender provides protection against advanced threats in email and collaboration tools?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the purpose of Microsoft Sentinel's Fusion technology?",
		Options: []string{
			"To encrypt data in transit",
			"To correlate alerts from multiple sources into high-confidence incidents",
			"To manage user access rights",
			"To create virtual machines",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Endpoint provides risk-based assessment of vulnerabilities?",
		Options: []string{
			"Threat & Vulnerability Management",
			"Automated Investigation and Remediation",
			"Advanced Hunting",
			"Device Control",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Office 365's Safe Attachments feature?",
		Options: []string{
			"To encrypt email attachments",
			"To scan attachments for malware in a virtual environment",
			"To compress large attachments",
			"To archive old attachments",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender component helps detect and investigate advanced attacks on-premises and in the cloud?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Exchange Online Protection",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Just-In-Time VM Access feature?",
		Options: []string{
			"To automatically update virtual machines",
			"To provide temporary, on-demand access to VMs",
			"To create snapshots of virtual machines",
			"To monitor VM performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to create custom detection rules?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Analytics Rules",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the purpose of Microsoft Defender for Endpoint's Live Response feature?",
		Options: []string{
			"To provide real-time video streaming",
			"To allow real-time investigation and remediation on devices",
			"To enable live chat with support",
			"To stream live security events",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender feature helps prevent phishing attacks by analyzing email content and links?",
		Options: []string{
			"Safe Links",
			"Data Loss Prevention",
			"eDiscovery",
			"Information Rights Management",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main function of Microsoft Defender for Cloud's Adaptive Network Hardening?",
		Options: []string{
			"To automatically adjust firewall rules",
			"To encrypt network traffic",
			"To monitor network performance",
			"To create virtual networks",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to automate incident response actions?",
		Options: []string{
			"Workbooks",
			"Playbooks",
			"Hunting Queries",
			"Log Analytics",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the primary purpose of Microsoft Defender for Cloud Apps' Cloud Discovery feature?",
		Options: []string{
			"To create new cloud applications",
			"To identify shadow IT usage in an organization",
			"To discover on-premises servers",
			"To find lost cloud data",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Defender for Endpoint feature provides a visual representation of the attack chain?",
		Options: []string{
			"Advanced Hunting",
			"Automated Investigation and Remediation",
			"Attack Surface Reduction",
			"Incident Graph",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Identity's Lateral Movement Paths?",
		Options: []string{
			"To visualize potential attack paths in the network",
			"To create network diagrams",
			"To manage network traffic",
			"To optimize network performance",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to visualize and analyze data across multiple dimensions?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Regulatory Compliance dashboard?",
		Options: []string{
			"To create compliance reports",
			"To assess and track compliance with various standards and regulations",
			"To manage software licenses",
			"To schedule compliance audits",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender component provides protection for cloud-based applications?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Automated Investigation and Remediation feature?",
		Options: []string{
			"To automatically update software",
			"To investigate and remediate threats without human intervention",
			"To schedule system maintenance",
			"To create backup copies of files",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to proactively search for threats across your data?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's File Integrity Monitoring?",
		Options: []string{
			"To encrypt files",
			"To detect changes to important files and registry settings",
			"To compress large files",
			"To archive old files",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Office 365 helps protect against malicious URLs in email messages and Office documents?",
		Options: []string{
			"Safe Attachments",
			"Safe Links",
			"Anti-phishing",
			"Data Loss Prevention",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Identity's Entity Behavior Analytics?",
		Options: []string{
			"To analyze user login patterns",
			"To detect anomalous behavior and potential security threats",
			"To manage user accounts",
			"To create user profiles",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to connect and collect data from various sources?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Adaptive Application Controls?",
		Options: []string{
			"To automatically update applications",
			"To whitelist applications and prevent unauthorized software execution",
			"To create application backups",
			"To monitor application performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender feature helps prevent data leaks by monitoring and protecting sensitive information?",
		Options: []string{
			"Safe Links",
			"Data Loss Prevention",
			"eDiscovery",
			"Information Rights Management",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Threat & Vulnerability Management?",
		Options: []string{
			"To create backups of critical data",
			"To discover, prioritize, and remediate vulnerabilities and misconfigurations",
			"To manage software licenses",
			"To monitor network traffic",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to create custom dashboards and reports?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud Apps' Policy Management?",
		Options: []string{
			"To create IT policies",
			"To define and enforce security policies for cloud applications",
			"To manage human resources policies",
			"To set up backup policies",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Identity helps detect reconnaissance activities?",
		Options: []string{
			"Lateral Movement Paths",
			"Entity Behavior Analytics",
			"Domain Dominance",
			"Reconnaissance Detection",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Advanced Hunting feature?",
		Options: []string{
			"To automatically hunt for threats",
			"To allow custom queries for proactive threat hunting",
			"To schedule hunting activities",
			"To manage hunting licenses",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature helps you understand the scope and impact of an incident?",
		Options: []string{
			"Incident Investigation",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Security Alerts?",
		Options: []string{
			"To notify about system updates",
			"To provide actionable information about detected threats and anomalies",
			"To alert about network outages",
			"To remind about scheduled maintenance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender component provides protection for endpoints such as desktops, laptops, and mobile devices?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Office 365's Anti-phishing policies?",
		Options: []string{
			"To encrypt emails",
			"To detect and prevent phishing attempts in email communications",
			"To compress email attachments",
			"To archive old emails",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to create custom alert rules?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Analytics Rules",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Just-In-Time VM Access?",
		Options: []string{
			"To automatically update virtual machines",
			"To provide temporary, controlled access to VM management ports",
			"To create VM snapshots",
			"To monitor VM performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Identity helps detect attempts to compromise credentials?",
		Options: []string{
			"Lateral Movement Paths",
			"Entity Behavior Analytics",
			"Compromise Detection",
			"Domain Dominance",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Device Control?",
		Options: []string{
			"To manage device inventory",
			"To control and restrict usage of removable storage and peripherals",
			"To update device drivers",
			"To monitor device performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to automate common tasks and orchestrate your security workflows?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Adaptive Network Hardening?",
		Options: []string{
			"To automatically adjust network settings",
			"To recommend and enforce network security group rules",
			"To encrypt network traffic",
			"To optimize network performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender feature helps protect against malware in email attachments?",
		Options: []string{
			"Safe Attachments",
			"Safe Links",
			"Anti-phishing",
			"Data Loss Prevention",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Identity's Honeytoken accounts?",
		Options: []string{
			"To create fake user accounts",
			"To detect and alert on potential malicious activities",
			"To manage privileged accounts",
			"To automate account creation",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to visualize attack patterns and trends over time?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Investigation Graph",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Regulatory Compliance dashboard?",
		Options: []string{
			"To create compliance reports",
			"To assess and track compliance with various standards and regulations",
			"To manage software licenses",
			"To schedule compliance audits",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Endpoint provides risk-based vulnerability management?",
		Options: []string{
			"Threat & Vulnerability Management",
			"Automated Investigation and Remediation",
			"Advanced Hunting",
			"Device Control",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Office 365's Threat Explorer?",
		Options: []string{
			"To create threat reports",
			"To provide real-time and historical views of threats in your environment",
			"To automatically block threats",
			"To schedule threat scans",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to create custom detection rules using a domain-specific language?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Kusto Query Language (KQL)",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Secure Score?",
		Options: []string{
			"To rate the performance of cloud services",
			"To measure and report on your security posture",
			"To score the difficulty of security tasks",
			"To rank cloud providers",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender component focuses on protecting cloud applications and services?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Network Protection feature?",
		Options: []string{
			"To encrypt network traffic",
			"To block outbound connections to untrusted hosts and malicious IP addresses",
			"To optimize network performance",
			"To create virtual networks",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to search for specific events or patterns across your data?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Just-In-Time VM Access?",
		Options: []string{
			"To automatically update virtual machines",
			"To provide temporary, controlled access to VM management ports",
			"To create VM snapshots",
			"To monitor VM performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Identity helps detect lateral movement attempts?",
		Options: []string{
			"Lateral Movement Paths",
			"Entity Behavior Analytics",
			"Compromise Detection",
			"Domain Dominance",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Automated Investigation and Remediation?",
		Options: []string{
			"To automatically update software",
			"To investigate and remediate threats without human intervention",
			"To schedule system maintenance",
			"To create backup copies of files",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to create interactive reports and dashboards?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Adaptive Application Controls?",
		Options: []string{
			"To automatically update applications",
			"To whitelist applications and prevent unauthorized software execution",
			"To create application backups",
			"To monitor application performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender feature helps prevent data leaks by monitoring and protecting sensitive information?",
		Options: []string{
			"Safe Links",
			"Data Loss Prevention",
			"eDiscovery",
			"Information Rights Management",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Identity's Suspicious Activity Timeline?",
		Options: []string{
			"To schedule security tasks",
			"To provide a chronological view of suspicious activities and alerts",
			"To manage user schedules",
			"To track system updates",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to automate incident response actions?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's File Integrity Monitoring?",
		Options: []string{
			"To encrypt files",
			"To detect changes to important files and registry settings",
			"To compress large files",
			"To archive old files",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Endpoint provides a visual representation of the attack chain?",
		Options: []string{
			"Advanced Hunting",
			"Automated Investigation and Remediation",
			"Attack Surface Reduction",
			"Incident Graph",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Office 365's Safe Links feature?",
		Options: []string{
			"To encrypt email links",
			"To scan and rewrite URLs to protect users when they click on links",
			"To create shortened URLs",
			"To track link clicks",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to connect and collect data from various sources?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Security Center?",
		Options: []string{
			"To manage cloud resources",
			"To provide a unified view of security across all your cloud and on-premises resources",
			"To create virtual machines",
			"To monitor network traffic",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender component provides protection for on-premises and cloud-based identity infrastructure?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Threat Analytics?",
		Options: []string{
			"To create threat reports",
			"To provide actionable intelligence on active threats and mitigation recommendations",
			"To automatically block threats",
			"To schedule threat scans",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to create custom alert rules?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Analytics Rules",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Regulatory Compliance dashboard?",
		Options: []string{
			"To create compliance reports",
			"To assess and track compliance with various standards and regulations",
			"To manage software licenses",
			"To schedule compliance audits",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Identity helps detect attempts to exploit vulnerabilities?",
		Options: []string{
			"Lateral Movement Paths",
			"Entity Behavior Analytics",
			"Exploit Detection",
			"Domain Dominance",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Attack Surface Reduction rules?",
		Options: []string{
			"To automatically update software",
			"To reduce vulnerabilities and block common attack techniques",
			"To create system restore points",
			"To monitor system performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature helps you understand the scope and impact of an incident?",
		Options: []string{
			"Incident Investigation",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Network Map?",
		Options: []string{
			"To create network diagrams",
			"To visualize network topology and identify security issues",
			"To optimize network traffic",
			"To manage IP addresses",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender feature helps protect against malicious URLs in email messages and Office documents?",
		Options: []string{
			"Safe Attachments",
			"Safe Links",
			"Anti-phishing",
			"Data Loss Prevention",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Identity's Domain Dominance alerts?",
		Options: []string{
			"To manage domain controllers",
			"To detect potential domain takeover attempts",
			"To optimize domain performance",
			"To create domain backups",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to create custom dashboards and reports?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Adaptive Network Hardening?",
		Options: []string{
			"To automatically adjust network settings",
			"To recommend and enforce network security group rules",
			"To encrypt network traffic",
			"To optimize network performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Endpoint provides risk-based vulnerability management?",
		Options: []string{
			"Threat & Vulnerability Management",
			"Automated Investigation and Remediation",
			"Advanced Hunting",
			"Device Control",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Office 365's Threat Explorer?",
		Options: []string{
			"To create threat reports",
			"To provide real-time and historical views of threats in your environment",
			"To automatically block threats",
			"To schedule threat scans",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to create custom detection rules using a domain-specific language?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Kusto Query Language (KQL)",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Secure Score?",
		Options: []string{
			"To rate the performance of cloud services",
			"To measure and report on your security posture",
			"To score the difficulty of security tasks",
			"To rank cloud providers",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender component focuses on protecting cloud applications and services?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's Network Protection feature?",
		Options: []string{
			"To encrypt network traffic",
			"To block outbound connections to untrusted hosts and malicious IP addresses",
			"To optimize network performance",
			"To create virtual networks",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to search for specific events or patterns across your data?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Just-In-Time VM Access?",
		Options: []string{
			"To automatically update virtual machines",
			"To provide temporary, controlled access to VM management ports",
			"To create VM snapshots",
			"To monitor VM performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Identity helps detect lateral movement attempts?",
		Options: []string{
			"Lateral Movement Paths",
			"Entity Behavior Analytics",
			"Compromise Detection",
			"Domain Dominance",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Endpoint's EDR in block mode?",
		Options: []string{
			"To automatically update software",
			"To block malicious artifacts even when an antivirus is not present",
			"To create system restore points",
			"To monitor system performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to visualize attack patterns and trends over time?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Investigation Graph",
		},
		CorrectAnswer: 3,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Workflow automation?",
		Options: []string{
			"To automate software updates",
			"To create automated responses to specific security alerts and events",
			"To manage workflow schedules",
			"To optimize cloud resource allocation",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender feature helps prevent data exfiltration to unallowed apps and locations?",
		Options: []string{
			"Safe Links",
			"Data Loss Prevention",
			"eDiscovery",
			"Information Barriers",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Identity's Reconnaissance Detection?",
		Options: []string{
			"To scan for new devices",
			"To detect potential adversary reconnaissance activities",
			"To discover network topology",
			"To map user relationships",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel capability allows you to automate common tasks and orchestrate your security workflows?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Hunting Queries",
			"Data Connectors",
		},
		CorrectAnswer: 0,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Adaptive Application Controls?",
		Options: []string{
			"To automatically update applications",
			"To whitelist applications and prevent unauthorized software execution",
			"To create application backups",
			"To monitor application performance",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which feature in Microsoft Defender for Endpoint provides automated investigation and remediation capabilities?",
		Options: []string{
			"Advanced Hunting",
			"Automated Investigation and Remediation (AIR)",
			"Attack Surface Reduction",
			"Device Control",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "What is the main purpose of Microsoft Defender for Office 365's Campaign Views?",
		Options: []string{
			"To create marketing campaigns",
			"To provide a consolidated view of email-based attacks targeting your organization",
			"To manage email distribution lists",
			"To schedule email broadcasts",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft Sentinel feature allows you to create custom alert rules?",
		Options: []string{
			"Playbooks",
			"Workbooks",
			"Analytics Rules",
			"Data Connectors",
		},
		CorrectAnswer: 2,
	},
	{
		Text: "What is the primary function of Microsoft Defender for Cloud's Regulatory Compliance dashboard?",
		Options: []string{
			"To create compliance reports",
			"To assess and track compliance with various standards and regulations",
			"To manage software licenses",
			"To schedule compliance audits",
		},
		CorrectAnswer: 1,
	},
	{
		Text: "Which Microsoft 365 Defender component provides protection for endpoints such as desktops, laptops, and mobile devices?",
		Options: []string{
			"Microsoft Defender for Endpoint",
			"Microsoft Defender for Office 365",
			"Microsoft Defender for Identity",
			"Microsoft Defender for Cloud Apps",
		},
		CorrectAnswer: 0,
	},
}
