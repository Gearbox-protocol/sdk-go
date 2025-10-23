# sdk-go


### Important information for contributors
As a contributor to the Gearbox Protocol GitHub repository, your pull requests indicate acceptance of our Gearbox Contribution Agreement. This agreement outlines that you assign the Intellectual Property Rights of your contributions to the Gearbox Foundation. This helps safeguard the Gearbox protocol and ensure the accumulation of its intellectual property. Contributions become part of the repository and may be used for various purposes, including commercial. As recognition for your expertise and work, you receive the opportunity to participate in the protocol's development and the potential to see your work integrated within it. The full Gearbox Contribution Agreement is accessible within the [repository](/ContributionAgreement) for comprehensive understanding. [Let's innovate together!]



# Environment variables

**ADDRESS_PROVIDER**
**MARKET_CONFIGURATORS**
**ETHERSCAN_PROXY_URL** in ts and logs from etherscan via this url
**ETHERSCAN_API_KEY** is for getting ts and logs from etherscan, 
**MORALIS_API_KEY** api key for getting block number
**ETHERSCAN_API_KEY** should be set if in first log , if the discoveredat is zero and network has etherscan
**ETHERSCAN_GETLOG_DISABLED** if not set in the filter logs , the firstlog of each address is used and min is used for getting logs
**OPTIMISTIC_LIQUIDATION** if opt liq used wbtcPrice can be zero for liquidator in the gearbox_oracle
**REDSTONE_URL**, if set the redstone server is not used for getting price on demand.