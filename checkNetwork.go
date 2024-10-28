package main

import (
	"github.com/c9s/goprocinfo/linux"
	"github.com/infrasonar/go-libagent"
)

func readNetworkStats(state map[string][]map[string]any) error {
	net, err := linux.ReadNetworkStat("/proc/net/dev")
	if err != nil {
		return err
	}

	items := []map[string]any{}

	for _, n := range net {
		items = append(items, map[string]any{
			"name":         n.Iface,
			"rxBytes":      n.RxBytes,
			"rxPackets":    n.RxPackets,
			"rxErrs":       n.RxErrs,
			"rxDrop":       n.RxDrop,
			"rxFifo":       n.RxFifo,
			"rxFrame":      n.RxFrame,
			"rxCompressed": n.RxCompressed,
			"rxMulticast":  n.RxMulticast,
			"txBytes":      n.TxBytes,
			"txPackets":    n.TxPackets,
			"txErrs":       n.TxErrs,
			"txDrop":       n.TxDrop,
			"txFifo":       n.TxFifo,
			"txColls":      n.TxColls,
			"txCarrier":    n.TxCarrier,
			"txCompressed": n.TxCompressed,
		})
	}

	state["networkStats"] = items
	return nil
}

func CheckNetwork(check *libagent.Check) (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}

	err := readNetworkStats(state)
	if err != nil {
		return nil, err
	}

	return state, nil
}
