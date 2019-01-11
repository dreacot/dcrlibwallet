package addresshelper

import (
	"fmt"
	"github.com/decred/dcrd/dcrutil"
	"github.com/decred/dcrd/txscript"
	"github.com/decred/dcrwallet/netparams"
)

func PkScript(address string) ([]byte, error) {
	addr, err := dcrutil.DecodeAddress(address)
	if err != nil {
		return nil, fmt.Errorf("error decoding address '%s': %s", address, err.Error())
	}

	return txscript.PayToAddrScript(addr)
}

func DecodeForNetwork(address string, params *netparams.Params) (dcrutil.Address, error) {
	addr, err := dcrutil.DecodeAddress(address)
	if err != nil {
		return nil, err
	}
	if !addr.IsForNet(params.Params) {
		return nil, fmt.Errorf("address %s is not intended for use on %s", address, params.Name)
	}
	return addr, nil
}

func PkScriptAddresses(params *netparams.Params, pkScript []byte) ([]string, error) {
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(txscript.DefaultScriptVersion, pkScript, params.Params)
	if err != nil {
		return nil, err
	}

	encodedAddresses := make([]string, len(addresses))
	for i, address := range addresses {
		encodedAddresses[i] = address.EncodeAddress()
	}

	return encodedAddresses, nil
}