package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

var samples = `

{
    "API": {
        "createAsset": {
            "args": [
                {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                }
            ],
            "function": "createAsset"
        }
    },
    "Model": {
        "asset": {
            "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
            "carrier": "The carrier in possession of this asset",
            "common": {
                "appdata": [
                    {
                        "K": "carpe noctem",
                        "V": "carpe noctem"
                    }
                ],
                "deviceID": "A unique identifier for the device that sent the current event",
                "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                "location": {
                    "latitude": 123.456,
                    "longitude": 123.456
                }
            },
            "temperature": 123.456
        },
        "assetstate": {
            "alerts": [
                "An alert name"
            ],
            "assetID": "This asset's world state asset ID",
            "assetIDpath": "Qualified property path to the asset's ID, declared in the contract code",
            "class": "The asset's asset class",
            "compliant": true,
            "eventin": {
                "asset": {
                    "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                    "carrier": "The carrier in possession of this asset",
                    "common": {
                        "appdata": [
                            {
                                "K": "carpe noctem",
                                "V": "carpe noctem"
                            }
                        ],
                        "deviceID": "A unique identifier for the device that sent the current event",
                        "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                        "location": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        }
                    },
                    "temperature": 123.456
                }
            },
            "eventout": {
                "asset": {
                    "name": "The chaincode event's name",
                    "payload": {}
                }
            },
            "prefix": "The asset's asset class prefix in world state",
            "state": {
                "asset": {
                    "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                    "carrier": "The carrier in possession of this asset",
                    "common": {
                        "appdata": [
                            {
                                "K": "carpe noctem",
                                "V": "carpe noctem"
                            }
                        ],
                        "deviceID": "A unique identifier for the device that sent the current event",
                        "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                        "location": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        }
                    },
                    "temperature": 123.456
                }
            },
            "txnid": "Transaction UUID matching the blockchain",
            "txnts": "Transaction timestamp matching the blockchain"
        },
        "assetstatearray": [
            {
                "alerts": [
                    "An alert name"
                ],
                "assetID": "This asset's world state asset ID",
                "assetIDpath": "Qualified property path to the asset's ID, declared in the contract code",
                "class": "The asset's asset class",
                "compliant": true,
                "eventin": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "eventout": {
                    "asset": {
                        "name": "The chaincode event's name",
                        "payload": {}
                    }
                },
                "prefix": "The asset's asset class prefix in world state",
                "state": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "txnid": "Transaction UUID matching the blockchain",
                "txnts": "Transaction timestamp matching the blockchain"
            }
        ],
        "assetstateexternal": {
            "^DEF": {
                "alerts": [
                    "An alert name"
                ],
                "assetID": "This asset's world state asset ID",
                "assetIDpath": "Qualified property path to the asset's ID, declared in the contract code",
                "class": "The asset's asset class",
                "compliant": true,
                "eventin": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "eventout": {
                    "asset": {
                        "name": "The chaincode event's name",
                        "payload": {}
                    }
                },
                "prefix": "The asset's asset class prefix in world state",
                "state": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "txnid": "Transaction UUID matching the blockchain",
                "txnts": "Transaction timestamp matching the blockchain"
            }
        },
        "invokeevent": {
            "name": "The chaincode event's name",
            "payload": {}
        },
        "ioteventcommon": {
            "appdata": [
                {
                    "K": "carpe noctem",
                    "V": "carpe noctem"
                }
            ],
            "deviceID": "A unique identifier for the device that sent the current event",
            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            }
        },
        "stateFilter": {
            "match": "all",
            "select": [
                {
                    "qprop": "Qualified property to compare, for example 'asset.assetID'",
                    "value": "Value to be compared"
                }
            ]
        }
    }
}`


	var readAssetSamples iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(samples), nil
	}

	func init() {
		iot.AddRoute("readAssetSamples", "query", iot.SystemClass, readAssetSamples)
	}
	