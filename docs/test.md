## Test input file
Sample of input test file.
```
{
    blocks:{ 
        3:{
            events: [
                address: "",
                data: [],
                topics: [],
                txHash: "",
            ],
            calls: {
                pools: [],
                cms: [],
                accounts: [],
                masks:[],
                others: {
                    "sig" :{"target": "data"},
                }
            }
        }
    },
    states: {
        oracles: {
            oracle: [{
                oracle: "",
                feed: "",
                block: "",
            }]
        },
        otherCalls: {
            "sig" :{"target": "data"},
        }
    }
    mocks:{
        tokens: "",
        syncAdapters: "",
    }
}
```