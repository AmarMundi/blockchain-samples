package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

var samples = `

{
    "API": {
        "createAssetContainer": {
            "args": [
                {
                    "container": {
                        "barcode": "A container's ID",
                        "carrier": "The carrier in possession of this container",
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
            "function": "createAssetContainer"
        },
        "updateAssetContainer": {
            "args": [
                {
                    "container": {
                        "barcode": "A container's ID",
                        "carrier": "The carrier in possession of this container",
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
            "function": "updateAssetContainer"
        }
    },
    "Model": {
        "container": {
            "barcode": "A container's ID",
            "carrier": "The carrier in possession of this container",
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
        "containerstate": {
            "AssetKey": "This container's world state container ID",
            "alerts": [
                "An alert name"
            ],
            "assetIDpath": "Qualified property path to the container's ID, declared in the contract code",
            "class": "The container's asset class",
            "compliant": true,
            "eventin": {
                "container": {
                    "barcode": "A container's ID",
                    "carrier": "The carrier in possession of this container",
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
                "container": {
                    "name": "The chaincode event's name",
                    "payload": {}
                }
            },
            "prefix": "The container's asset class prefix in world state",
            "state": {
                "container": {
                    "barcode": "A container's ID",
                    "carrier": "The carrier in possession of this container",
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
        "containerstatearray": [
            {
                "^CON": "INVALID OBJECT - MISSING PROPERTIES"
            }
        ],
        "containerstateexternal": {
            "^CON": "INVALID OBJECT - MISSING PROPERTIES"
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
	