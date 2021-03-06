// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package proxy

import (
	"github.com/docker/libnetwork/ipvs"
	"github.com/vishvananda/netlink"
	"net"
	"sync"
)

var (
	lockLinuxNetworkingMockcleanupMangleTableRule         sync.RWMutex
	lockLinuxNetworkingMockgetKubeDummyInterface          sync.RWMutex
	lockLinuxNetworkingMockipAddrAdd                      sync.RWMutex
	lockLinuxNetworkingMockipAddrDel                      sync.RWMutex
	lockLinuxNetworkingMockipvsAddFWMarkService           sync.RWMutex
	lockLinuxNetworkingMockipvsAddServer                  sync.RWMutex
	lockLinuxNetworkingMockipvsAddService                 sync.RWMutex
	lockLinuxNetworkingMockipvsDelDestination             sync.RWMutex
	lockLinuxNetworkingMockipvsDelService                 sync.RWMutex
	lockLinuxNetworkingMockipvsGetDestinations            sync.RWMutex
	lockLinuxNetworkingMockipvsGetServices                sync.RWMutex
	lockLinuxNetworkingMockipvsNewDestination             sync.RWMutex
	lockLinuxNetworkingMockipvsNewService                 sync.RWMutex
	lockLinuxNetworkingMockipvsUpdateDestination          sync.RWMutex
	lockLinuxNetworkingMockipvsUpdateService              sync.RWMutex
	lockLinuxNetworkingMockprepareEndpointForDsr          sync.RWMutex
	lockLinuxNetworkingMocksetupPolicyRoutingForDSR       sync.RWMutex
	lockLinuxNetworkingMocksetupRoutesForExternalIPForDSR sync.RWMutex
)

// LinuxNetworkingMock is a mock implementation of LinuxNetworking.
//
//     func TestSomethingThatUsesLinuxNetworking(t *testing.T) {
//
//         // make and configure a mocked LinuxNetworking
//         mockedLinuxNetworking := &LinuxNetworkingMock{
//             cleanupMangleTableRuleFunc: func(ip string, protocol string, port string, fwmark string) error {
// 	               panic("TODO: mock out the cleanupMangleTableRule method")
//             },
//             getKubeDummyInterfaceFunc: func() (netlink.Link, error) {
// 	               panic("TODO: mock out the getKubeDummyInterface method")
//             },
//             ipAddrAddFunc: func(iface netlink.Link, ip string, addRoute bool) error {
// 	               panic("TODO: mock out the ipAddrAdd method")
//             },
//             ipAddrDelFunc: func(iface netlink.Link, ip string) error {
// 	               panic("TODO: mock out the ipAddrDel method")
//             },
//             ipvsAddFWMarkServiceFunc: func(vip net.IP, protocol uint16, port uint16, persistent bool, scheduler string, flags schedFlags) (*ipvs.Service, error) {
// 	               panic("TODO: mock out the ipvsAddFWMarkService method")
//             },
//             ipvsAddServerFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
// 	               panic("TODO: mock out the ipvsAddServer method")
//             },
//             ipvsAddServiceFunc: func(svcs []*ipvs.Service, vip net.IP, protocol uint16, port uint16, persistent bool, scheduler string, flags schedFlags) (*ipvs.Service, error) {
// 	               panic("TODO: mock out the ipvsAddService method")
//             },
//             ipvsDelDestinationFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
// 	               panic("TODO: mock out the ipvsDelDestination method")
//             },
//             ipvsDelServiceFunc: func(ipvsSvc *ipvs.Service) error {
// 	               panic("TODO: mock out the ipvsDelService method")
//             },
//             ipvsGetDestinationsFunc: func(ipvsSvc *ipvs.Service) ([]*ipvs.Destination, error) {
// 	               panic("TODO: mock out the ipvsGetDestinations method")
//             },
//             ipvsGetServicesFunc: func() ([]*ipvs.Service, error) {
// 	               panic("TODO: mock out the ipvsGetServices method")
//             },
//             ipvsNewDestinationFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
// 	               panic("TODO: mock out the ipvsNewDestination method")
//             },
//             ipvsNewServiceFunc: func(ipvsSvc *ipvs.Service) error {
// 	               panic("TODO: mock out the ipvsNewService method")
//             },
//             ipvsUpdateDestinationFunc: func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
// 	               panic("TODO: mock out the ipvsUpdateDestination method")
//             },
//             ipvsUpdateServiceFunc: func(ipvsSvc *ipvs.Service) error {
// 	               panic("TODO: mock out the ipvsUpdateService method")
//             },
//             prepareEndpointForDsrFunc: func(containerId string, endpointIP string, vip string) error {
// 	               panic("TODO: mock out the prepareEndpointForDsr method")
//             },
//             setupPolicyRoutingForDSRFunc: func() error {
// 	               panic("TODO: mock out the setupPolicyRoutingForDSR method")
//             },
//             setupRoutesForExternalIPForDSRFunc: func(in1 serviceInfoMap) error {
// 	               panic("TODO: mock out the setupRoutesForExternalIPForDSR method")
//             },
//         }
//
//         // TODO: use mockedLinuxNetworking in code that requires LinuxNetworking
//         //       and then make assertions.
//
//     }
type LinuxNetworkingMock struct {
	// cleanupMangleTableRuleFunc mocks the cleanupMangleTableRule method.
	cleanupMangleTableRuleFunc func(ip string, protocol string, port string, fwmark string) error

	// getKubeDummyInterfaceFunc mocks the getKubeDummyInterface method.
	getKubeDummyInterfaceFunc func() (netlink.Link, error)

	// ipAddrAddFunc mocks the ipAddrAdd method.
	ipAddrAddFunc func(iface netlink.Link, ip string, addRoute bool) error

	// ipAddrDelFunc mocks the ipAddrDel method.
	ipAddrDelFunc func(iface netlink.Link, ip string) error

	// ipvsAddFWMarkServiceFunc mocks the ipvsAddFWMarkService method.
	ipvsAddFWMarkServiceFunc func(vip net.IP, protocol uint16, port uint16, persistent bool, scheduler string, flags schedFlags) (*ipvs.Service, error)

	// ipvsAddServerFunc mocks the ipvsAddServer method.
	ipvsAddServerFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsAddServiceFunc mocks the ipvsAddService method.
	ipvsAddServiceFunc func(svcs []*ipvs.Service, vip net.IP, protocol uint16, port uint16, persistent bool, scheduler string, flags schedFlags) (*ipvs.Service, error)

	// ipvsDelDestinationFunc mocks the ipvsDelDestination method.
	ipvsDelDestinationFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsDelServiceFunc mocks the ipvsDelService method.
	ipvsDelServiceFunc func(ipvsSvc *ipvs.Service) error

	// ipvsGetDestinationsFunc mocks the ipvsGetDestinations method.
	ipvsGetDestinationsFunc func(ipvsSvc *ipvs.Service) ([]*ipvs.Destination, error)

	// ipvsGetServicesFunc mocks the ipvsGetServices method.
	ipvsGetServicesFunc func() ([]*ipvs.Service, error)

	// ipvsNewDestinationFunc mocks the ipvsNewDestination method.
	ipvsNewDestinationFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsNewServiceFunc mocks the ipvsNewService method.
	ipvsNewServiceFunc func(ipvsSvc *ipvs.Service) error

	// ipvsUpdateDestinationFunc mocks the ipvsUpdateDestination method.
	ipvsUpdateDestinationFunc func(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error

	// ipvsUpdateServiceFunc mocks the ipvsUpdateService method.
	ipvsUpdateServiceFunc func(ipvsSvc *ipvs.Service) error

	// prepareEndpointForDsrFunc mocks the prepareEndpointForDsr method.
	prepareEndpointForDsrFunc func(containerId string, endpointIP string, vip string) error

	// setupPolicyRoutingForDSRFunc mocks the setupPolicyRoutingForDSR method.
	setupPolicyRoutingForDSRFunc func() error

	// setupRoutesForExternalIPForDSRFunc mocks the setupRoutesForExternalIPForDSR method.
	setupRoutesForExternalIPForDSRFunc func(in1 serviceInfoMap) error

	// calls tracks calls to the methods.
	calls struct {
		// cleanupMangleTableRule holds details about calls to the cleanupMangleTableRule method.
		cleanupMangleTableRule []struct {
			// IP is the ip argument value.
			IP string
			// Protocol is the protocol argument value.
			Protocol string
			// Port is the port argument value.
			Port string
			// Fwmark is the fwmark argument value.
			Fwmark string
		}
		// getKubeDummyInterface holds details about calls to the getKubeDummyInterface method.
		getKubeDummyInterface []struct {
		}
		// ipAddrAdd holds details about calls to the ipAddrAdd method.
		ipAddrAdd []struct {
			// Iface is the iface argument value.
			Iface netlink.Link
			// IP is the ip argument value.
			IP string
			// AddRoute is the addRoute argument value.
			AddRoute bool
		}
		// ipAddrDel holds details about calls to the ipAddrDel method.
		ipAddrDel []struct {
			// Iface is the iface argument value.
			Iface netlink.Link
			// IP is the ip argument value.
			IP string
		}
		// ipvsAddFWMarkService holds details about calls to the ipvsAddFWMarkService method.
		ipvsAddFWMarkService []struct {
			// Vip is the vip argument value.
			Vip net.IP
			// Protocol is the protocol argument value.
			Protocol uint16
			// Port is the port argument value.
			Port uint16
			// Persistent is the persistent argument value.
			Persistent bool
			// Scheduler is the scheduler argument value.
			Scheduler string
			// Flags is the flags argument value.
			Flags schedFlags
		}
		// ipvsAddServer holds details about calls to the ipvsAddServer method.
		ipvsAddServer []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsAddService holds details about calls to the ipvsAddService method.
		ipvsAddService []struct {
			// Svcs is the svcs argument value.
			Svcs []*ipvs.Service
			// Vip is the vip argument value.
			Vip net.IP
			// Protocol is the protocol argument value.
			Protocol uint16
			// Port is the port argument value.
			Port uint16
			// Persistent is the persistent argument value.
			Persistent bool
			// Scheduler is the scheduler argument value.
			Scheduler string
			// Flags is the flags argument value.
			Flags schedFlags
		}
		// ipvsDelDestination holds details about calls to the ipvsDelDestination method.
		ipvsDelDestination []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsDelService holds details about calls to the ipvsDelService method.
		ipvsDelService []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// ipvsGetDestinations holds details about calls to the ipvsGetDestinations method.
		ipvsGetDestinations []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// ipvsGetServices holds details about calls to the ipvsGetServices method.
		ipvsGetServices []struct {
		}
		// ipvsNewDestination holds details about calls to the ipvsNewDestination method.
		ipvsNewDestination []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsNewService holds details about calls to the ipvsNewService method.
		ipvsNewService []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// ipvsUpdateDestination holds details about calls to the ipvsUpdateDestination method.
		ipvsUpdateDestination []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
			// IpvsDst is the ipvsDst argument value.
			IpvsDst *ipvs.Destination
		}
		// ipvsUpdateService holds details about calls to the ipvsUpdateService method.
		ipvsUpdateService []struct {
			// IpvsSvc is the ipvsSvc argument value.
			IpvsSvc *ipvs.Service
		}
		// prepareEndpointForDsr holds details about calls to the prepareEndpointForDsr method.
		prepareEndpointForDsr []struct {
			// ContainerId is the containerId argument value.
			ContainerId string
			// EndpointIP is the endpointIP argument value.
			EndpointIP string
			// Vip is the vip argument value.
			Vip string
		}
		// setupPolicyRoutingForDSR holds details about calls to the setupPolicyRoutingForDSR method.
		setupPolicyRoutingForDSR []struct {
		}
		// setupRoutesForExternalIPForDSR holds details about calls to the setupRoutesForExternalIPForDSR method.
		setupRoutesForExternalIPForDSR []struct {
			// In1 is the in1 argument value.
			In1 serviceInfoMap
		}
	}
}

// cleanupMangleTableRule calls cleanupMangleTableRuleFunc.
func (mock *LinuxNetworkingMock) cleanupMangleTableRule(ip string, protocol string, port string, fwmark string) error {
	if mock.cleanupMangleTableRuleFunc == nil {
		panic("LinuxNetworkingMock.cleanupMangleTableRuleFunc: method is nil but LinuxNetworking.cleanupMangleTableRule was just called")
	}
	callInfo := struct {
		IP       string
		Protocol string
		Port     string
		Fwmark   string
	}{
		IP:       ip,
		Protocol: protocol,
		Port:     port,
		Fwmark:   fwmark,
	}
	lockLinuxNetworkingMockcleanupMangleTableRule.Lock()
	mock.calls.cleanupMangleTableRule = append(mock.calls.cleanupMangleTableRule, callInfo)
	lockLinuxNetworkingMockcleanupMangleTableRule.Unlock()
	return mock.cleanupMangleTableRuleFunc(ip, protocol, port, fwmark)
}

// cleanupMangleTableRuleCalls gets all the calls that were made to cleanupMangleTableRule.
// Check the length with:
//     len(mockedLinuxNetworking.cleanupMangleTableRuleCalls())
func (mock *LinuxNetworkingMock) cleanupMangleTableRuleCalls() []struct {
	IP       string
	Protocol string
	Port     string
	Fwmark   string
} {
	var calls []struct {
		IP       string
		Protocol string
		Port     string
		Fwmark   string
	}
	lockLinuxNetworkingMockcleanupMangleTableRule.RLock()
	calls = mock.calls.cleanupMangleTableRule
	lockLinuxNetworkingMockcleanupMangleTableRule.RUnlock()
	return calls
}

// getKubeDummyInterface calls getKubeDummyInterfaceFunc.
func (mock *LinuxNetworkingMock) getKubeDummyInterface() (netlink.Link, error) {
	if mock.getKubeDummyInterfaceFunc == nil {
		panic("LinuxNetworkingMock.getKubeDummyInterfaceFunc: method is nil but LinuxNetworking.getKubeDummyInterface was just called")
	}
	callInfo := struct {
	}{}
	lockLinuxNetworkingMockgetKubeDummyInterface.Lock()
	mock.calls.getKubeDummyInterface = append(mock.calls.getKubeDummyInterface, callInfo)
	lockLinuxNetworkingMockgetKubeDummyInterface.Unlock()
	return mock.getKubeDummyInterfaceFunc()
}

// getKubeDummyInterfaceCalls gets all the calls that were made to getKubeDummyInterface.
// Check the length with:
//     len(mockedLinuxNetworking.getKubeDummyInterfaceCalls())
func (mock *LinuxNetworkingMock) getKubeDummyInterfaceCalls() []struct {
} {
	var calls []struct {
	}
	lockLinuxNetworkingMockgetKubeDummyInterface.RLock()
	calls = mock.calls.getKubeDummyInterface
	lockLinuxNetworkingMockgetKubeDummyInterface.RUnlock()
	return calls
}

// ipAddrAdd calls ipAddrAddFunc.
func (mock *LinuxNetworkingMock) ipAddrAdd(iface netlink.Link, ip string, addRoute bool) error {
	if mock.ipAddrAddFunc == nil {
		panic("LinuxNetworkingMock.ipAddrAddFunc: method is nil but LinuxNetworking.ipAddrAdd was just called")
	}
	callInfo := struct {
		Iface    netlink.Link
		IP       string
		AddRoute bool
	}{
		Iface:    iface,
		IP:       ip,
		AddRoute: addRoute,
	}
	lockLinuxNetworkingMockipAddrAdd.Lock()
	mock.calls.ipAddrAdd = append(mock.calls.ipAddrAdd, callInfo)
	lockLinuxNetworkingMockipAddrAdd.Unlock()
	return mock.ipAddrAddFunc(iface, ip, addRoute)
}

// ipAddrAddCalls gets all the calls that were made to ipAddrAdd.
// Check the length with:
//     len(mockedLinuxNetworking.ipAddrAddCalls())
func (mock *LinuxNetworkingMock) ipAddrAddCalls() []struct {
	Iface    netlink.Link
	IP       string
	AddRoute bool
} {
	var calls []struct {
		Iface    netlink.Link
		IP       string
		AddRoute bool
	}
	lockLinuxNetworkingMockipAddrAdd.RLock()
	calls = mock.calls.ipAddrAdd
	lockLinuxNetworkingMockipAddrAdd.RUnlock()
	return calls
}

// ipAddrDel calls ipAddrDelFunc.
func (mock *LinuxNetworkingMock) ipAddrDel(iface netlink.Link, ip string) error {
	if mock.ipAddrDelFunc == nil {
		panic("LinuxNetworkingMock.ipAddrDelFunc: method is nil but LinuxNetworking.ipAddrDel was just called")
	}
	callInfo := struct {
		Iface netlink.Link
		IP    string
	}{
		Iface: iface,
		IP:    ip,
	}
	lockLinuxNetworkingMockipAddrDel.Lock()
	mock.calls.ipAddrDel = append(mock.calls.ipAddrDel, callInfo)
	lockLinuxNetworkingMockipAddrDel.Unlock()
	return mock.ipAddrDelFunc(iface, ip)
}

// ipAddrDelCalls gets all the calls that were made to ipAddrDel.
// Check the length with:
//     len(mockedLinuxNetworking.ipAddrDelCalls())
func (mock *LinuxNetworkingMock) ipAddrDelCalls() []struct {
	Iface netlink.Link
	IP    string
} {
	var calls []struct {
		Iface netlink.Link
		IP    string
	}
	lockLinuxNetworkingMockipAddrDel.RLock()
	calls = mock.calls.ipAddrDel
	lockLinuxNetworkingMockipAddrDel.RUnlock()
	return calls
}

// ipvsAddFWMarkService calls ipvsAddFWMarkServiceFunc.
func (mock *LinuxNetworkingMock) ipvsAddFWMarkService(vip net.IP, protocol uint16, port uint16, persistent bool, scheduler string, flags schedFlags) (*ipvs.Service, error) {
	if mock.ipvsAddFWMarkServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsAddFWMarkServiceFunc: method is nil but LinuxNetworking.ipvsAddFWMarkService was just called")
	}
	callInfo := struct {
		Vip        net.IP
		Protocol   uint16
		Port       uint16
		Persistent bool
		Scheduler  string
		Flags      schedFlags
	}{
		Vip:        vip,
		Protocol:   protocol,
		Port:       port,
		Persistent: persistent,
		Scheduler:  scheduler,
		Flags:      flags,
	}
	lockLinuxNetworkingMockipvsAddFWMarkService.Lock()
	mock.calls.ipvsAddFWMarkService = append(mock.calls.ipvsAddFWMarkService, callInfo)
	lockLinuxNetworkingMockipvsAddFWMarkService.Unlock()
	return mock.ipvsAddFWMarkServiceFunc(vip, protocol, port, persistent, scheduler, flags)
}

// ipvsAddFWMarkServiceCalls gets all the calls that were made to ipvsAddFWMarkService.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsAddFWMarkServiceCalls())
func (mock *LinuxNetworkingMock) ipvsAddFWMarkServiceCalls() []struct {
	Vip        net.IP
	Protocol   uint16
	Port       uint16
	Persistent bool
	Scheduler  string
	Flags      schedFlags
} {
	var calls []struct {
		Vip        net.IP
		Protocol   uint16
		Port       uint16
		Persistent bool
		Scheduler  string
		Flags      schedFlags
	}
	lockLinuxNetworkingMockipvsAddFWMarkService.RLock()
	calls = mock.calls.ipvsAddFWMarkService
	lockLinuxNetworkingMockipvsAddFWMarkService.RUnlock()
	return calls
}

// ipvsAddServer calls ipvsAddServerFunc.
func (mock *LinuxNetworkingMock) ipvsAddServer(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsAddServerFunc == nil {
		panic("LinuxNetworkingMock.ipvsAddServerFunc: method is nil but LinuxNetworking.ipvsAddServer was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	lockLinuxNetworkingMockipvsAddServer.Lock()
	mock.calls.ipvsAddServer = append(mock.calls.ipvsAddServer, callInfo)
	lockLinuxNetworkingMockipvsAddServer.Unlock()
	return mock.ipvsAddServerFunc(ipvsSvc, ipvsDst)
}

// ipvsAddServerCalls gets all the calls that were made to ipvsAddServer.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsAddServerCalls())
func (mock *LinuxNetworkingMock) ipvsAddServerCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	lockLinuxNetworkingMockipvsAddServer.RLock()
	calls = mock.calls.ipvsAddServer
	lockLinuxNetworkingMockipvsAddServer.RUnlock()
	return calls
}

// ipvsAddService calls ipvsAddServiceFunc.
func (mock *LinuxNetworkingMock) ipvsAddService(svcs []*ipvs.Service, vip net.IP, protocol uint16, port uint16, persistent bool, scheduler string, flags schedFlags) (*ipvs.Service, error) {
	if mock.ipvsAddServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsAddServiceFunc: method is nil but LinuxNetworking.ipvsAddService was just called")
	}
	callInfo := struct {
		Svcs       []*ipvs.Service
		Vip        net.IP
		Protocol   uint16
		Port       uint16
		Persistent bool
		Scheduler  string
		Flags      schedFlags
	}{
		Svcs:       svcs,
		Vip:        vip,
		Protocol:   protocol,
		Port:       port,
		Persistent: persistent,
		Scheduler:  scheduler,
		Flags:      flags,
	}
	lockLinuxNetworkingMockipvsAddService.Lock()
	mock.calls.ipvsAddService = append(mock.calls.ipvsAddService, callInfo)
	lockLinuxNetworkingMockipvsAddService.Unlock()
	return mock.ipvsAddServiceFunc(svcs, vip, protocol, port, persistent, scheduler, flags)
}

// ipvsAddServiceCalls gets all the calls that were made to ipvsAddService.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsAddServiceCalls())
func (mock *LinuxNetworkingMock) ipvsAddServiceCalls() []struct {
	Svcs       []*ipvs.Service
	Vip        net.IP
	Protocol   uint16
	Port       uint16
	Persistent bool
	Scheduler  string
	Flags      schedFlags
} {
	var calls []struct {
		Svcs       []*ipvs.Service
		Vip        net.IP
		Protocol   uint16
		Port       uint16
		Persistent bool
		Scheduler  string
		Flags      schedFlags
	}
	lockLinuxNetworkingMockipvsAddService.RLock()
	calls = mock.calls.ipvsAddService
	lockLinuxNetworkingMockipvsAddService.RUnlock()
	return calls
}

// ipvsDelDestination calls ipvsDelDestinationFunc.
func (mock *LinuxNetworkingMock) ipvsDelDestination(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsDelDestinationFunc == nil {
		panic("LinuxNetworkingMock.ipvsDelDestinationFunc: method is nil but LinuxNetworking.ipvsDelDestination was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	lockLinuxNetworkingMockipvsDelDestination.Lock()
	mock.calls.ipvsDelDestination = append(mock.calls.ipvsDelDestination, callInfo)
	lockLinuxNetworkingMockipvsDelDestination.Unlock()
	return mock.ipvsDelDestinationFunc(ipvsSvc, ipvsDst)
}

// ipvsDelDestinationCalls gets all the calls that were made to ipvsDelDestination.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsDelDestinationCalls())
func (mock *LinuxNetworkingMock) ipvsDelDestinationCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	lockLinuxNetworkingMockipvsDelDestination.RLock()
	calls = mock.calls.ipvsDelDestination
	lockLinuxNetworkingMockipvsDelDestination.RUnlock()
	return calls
}

// ipvsDelService calls ipvsDelServiceFunc.
func (mock *LinuxNetworkingMock) ipvsDelService(ipvsSvc *ipvs.Service) error {
	if mock.ipvsDelServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsDelServiceFunc: method is nil but LinuxNetworking.ipvsDelService was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	lockLinuxNetworkingMockipvsDelService.Lock()
	mock.calls.ipvsDelService = append(mock.calls.ipvsDelService, callInfo)
	lockLinuxNetworkingMockipvsDelService.Unlock()
	return mock.ipvsDelServiceFunc(ipvsSvc)
}

// ipvsDelServiceCalls gets all the calls that were made to ipvsDelService.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsDelServiceCalls())
func (mock *LinuxNetworkingMock) ipvsDelServiceCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	lockLinuxNetworkingMockipvsDelService.RLock()
	calls = mock.calls.ipvsDelService
	lockLinuxNetworkingMockipvsDelService.RUnlock()
	return calls
}

// ipvsGetDestinations calls ipvsGetDestinationsFunc.
func (mock *LinuxNetworkingMock) ipvsGetDestinations(ipvsSvc *ipvs.Service) ([]*ipvs.Destination, error) {
	if mock.ipvsGetDestinationsFunc == nil {
		panic("LinuxNetworkingMock.ipvsGetDestinationsFunc: method is nil but LinuxNetworking.ipvsGetDestinations was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	lockLinuxNetworkingMockipvsGetDestinations.Lock()
	mock.calls.ipvsGetDestinations = append(mock.calls.ipvsGetDestinations, callInfo)
	lockLinuxNetworkingMockipvsGetDestinations.Unlock()
	return mock.ipvsGetDestinationsFunc(ipvsSvc)
}

// ipvsGetDestinationsCalls gets all the calls that were made to ipvsGetDestinations.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsGetDestinationsCalls())
func (mock *LinuxNetworkingMock) ipvsGetDestinationsCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	lockLinuxNetworkingMockipvsGetDestinations.RLock()
	calls = mock.calls.ipvsGetDestinations
	lockLinuxNetworkingMockipvsGetDestinations.RUnlock()
	return calls
}

// ipvsGetServices calls ipvsGetServicesFunc.
func (mock *LinuxNetworkingMock) ipvsGetServices() ([]*ipvs.Service, error) {
	if mock.ipvsGetServicesFunc == nil {
		panic("LinuxNetworkingMock.ipvsGetServicesFunc: method is nil but LinuxNetworking.ipvsGetServices was just called")
	}
	callInfo := struct {
	}{}
	lockLinuxNetworkingMockipvsGetServices.Lock()
	mock.calls.ipvsGetServices = append(mock.calls.ipvsGetServices, callInfo)
	lockLinuxNetworkingMockipvsGetServices.Unlock()
	return mock.ipvsGetServicesFunc()
}

// ipvsGetServicesCalls gets all the calls that were made to ipvsGetServices.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsGetServicesCalls())
func (mock *LinuxNetworkingMock) ipvsGetServicesCalls() []struct {
} {
	var calls []struct {
	}
	lockLinuxNetworkingMockipvsGetServices.RLock()
	calls = mock.calls.ipvsGetServices
	lockLinuxNetworkingMockipvsGetServices.RUnlock()
	return calls
}

// ipvsNewDestination calls ipvsNewDestinationFunc.
func (mock *LinuxNetworkingMock) ipvsNewDestination(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsNewDestinationFunc == nil {
		panic("LinuxNetworkingMock.ipvsNewDestinationFunc: method is nil but LinuxNetworking.ipvsNewDestination was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	lockLinuxNetworkingMockipvsNewDestination.Lock()
	mock.calls.ipvsNewDestination = append(mock.calls.ipvsNewDestination, callInfo)
	lockLinuxNetworkingMockipvsNewDestination.Unlock()
	return mock.ipvsNewDestinationFunc(ipvsSvc, ipvsDst)
}

// ipvsNewDestinationCalls gets all the calls that were made to ipvsNewDestination.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsNewDestinationCalls())
func (mock *LinuxNetworkingMock) ipvsNewDestinationCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	lockLinuxNetworkingMockipvsNewDestination.RLock()
	calls = mock.calls.ipvsNewDestination
	lockLinuxNetworkingMockipvsNewDestination.RUnlock()
	return calls
}

// ipvsNewService calls ipvsNewServiceFunc.
func (mock *LinuxNetworkingMock) ipvsNewService(ipvsSvc *ipvs.Service) error {
	if mock.ipvsNewServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsNewServiceFunc: method is nil but LinuxNetworking.ipvsNewService was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	lockLinuxNetworkingMockipvsNewService.Lock()
	mock.calls.ipvsNewService = append(mock.calls.ipvsNewService, callInfo)
	lockLinuxNetworkingMockipvsNewService.Unlock()
	return mock.ipvsNewServiceFunc(ipvsSvc)
}

// ipvsNewServiceCalls gets all the calls that were made to ipvsNewService.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsNewServiceCalls())
func (mock *LinuxNetworkingMock) ipvsNewServiceCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	lockLinuxNetworkingMockipvsNewService.RLock()
	calls = mock.calls.ipvsNewService
	lockLinuxNetworkingMockipvsNewService.RUnlock()
	return calls
}

// ipvsUpdateDestination calls ipvsUpdateDestinationFunc.
func (mock *LinuxNetworkingMock) ipvsUpdateDestination(ipvsSvc *ipvs.Service, ipvsDst *ipvs.Destination) error {
	if mock.ipvsUpdateDestinationFunc == nil {
		panic("LinuxNetworkingMock.ipvsUpdateDestinationFunc: method is nil but LinuxNetworking.ipvsUpdateDestination was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}{
		IpvsSvc: ipvsSvc,
		IpvsDst: ipvsDst,
	}
	lockLinuxNetworkingMockipvsUpdateDestination.Lock()
	mock.calls.ipvsUpdateDestination = append(mock.calls.ipvsUpdateDestination, callInfo)
	lockLinuxNetworkingMockipvsUpdateDestination.Unlock()
	return mock.ipvsUpdateDestinationFunc(ipvsSvc, ipvsDst)
}

// ipvsUpdateDestinationCalls gets all the calls that were made to ipvsUpdateDestination.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsUpdateDestinationCalls())
func (mock *LinuxNetworkingMock) ipvsUpdateDestinationCalls() []struct {
	IpvsSvc *ipvs.Service
	IpvsDst *ipvs.Destination
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
		IpvsDst *ipvs.Destination
	}
	lockLinuxNetworkingMockipvsUpdateDestination.RLock()
	calls = mock.calls.ipvsUpdateDestination
	lockLinuxNetworkingMockipvsUpdateDestination.RUnlock()
	return calls
}

// ipvsUpdateService calls ipvsUpdateServiceFunc.
func (mock *LinuxNetworkingMock) ipvsUpdateService(ipvsSvc *ipvs.Service) error {
	if mock.ipvsUpdateServiceFunc == nil {
		panic("LinuxNetworkingMock.ipvsUpdateServiceFunc: method is nil but LinuxNetworking.ipvsUpdateService was just called")
	}
	callInfo := struct {
		IpvsSvc *ipvs.Service
	}{
		IpvsSvc: ipvsSvc,
	}
	lockLinuxNetworkingMockipvsUpdateService.Lock()
	mock.calls.ipvsUpdateService = append(mock.calls.ipvsUpdateService, callInfo)
	lockLinuxNetworkingMockipvsUpdateService.Unlock()
	return mock.ipvsUpdateServiceFunc(ipvsSvc)
}

// ipvsUpdateServiceCalls gets all the calls that were made to ipvsUpdateService.
// Check the length with:
//     len(mockedLinuxNetworking.ipvsUpdateServiceCalls())
func (mock *LinuxNetworkingMock) ipvsUpdateServiceCalls() []struct {
	IpvsSvc *ipvs.Service
} {
	var calls []struct {
		IpvsSvc *ipvs.Service
	}
	lockLinuxNetworkingMockipvsUpdateService.RLock()
	calls = mock.calls.ipvsUpdateService
	lockLinuxNetworkingMockipvsUpdateService.RUnlock()
	return calls
}

// prepareEndpointForDsr calls prepareEndpointForDsrFunc.
func (mock *LinuxNetworkingMock) prepareEndpointForDsr(containerId string, endpointIP string, vip string) error {
	if mock.prepareEndpointForDsrFunc == nil {
		panic("LinuxNetworkingMock.prepareEndpointForDsrFunc: method is nil but LinuxNetworking.prepareEndpointForDsr was just called")
	}
	callInfo := struct {
		ContainerId string
		EndpointIP  string
		Vip         string
	}{
		ContainerId: containerId,
		EndpointIP:  endpointIP,
		Vip:         vip,
	}
	lockLinuxNetworkingMockprepareEndpointForDsr.Lock()
	mock.calls.prepareEndpointForDsr = append(mock.calls.prepareEndpointForDsr, callInfo)
	lockLinuxNetworkingMockprepareEndpointForDsr.Unlock()
	return mock.prepareEndpointForDsrFunc(containerId, endpointIP, vip)
}

// prepareEndpointForDsrCalls gets all the calls that were made to prepareEndpointForDsr.
// Check the length with:
//     len(mockedLinuxNetworking.prepareEndpointForDsrCalls())
func (mock *LinuxNetworkingMock) prepareEndpointForDsrCalls() []struct {
	ContainerId string
	EndpointIP  string
	Vip         string
} {
	var calls []struct {
		ContainerId string
		EndpointIP  string
		Vip         string
	}
	lockLinuxNetworkingMockprepareEndpointForDsr.RLock()
	calls = mock.calls.prepareEndpointForDsr
	lockLinuxNetworkingMockprepareEndpointForDsr.RUnlock()
	return calls
}

// setupPolicyRoutingForDSR calls setupPolicyRoutingForDSRFunc.
func (mock *LinuxNetworkingMock) setupPolicyRoutingForDSR() error {
	if mock.setupPolicyRoutingForDSRFunc == nil {
		panic("LinuxNetworkingMock.setupPolicyRoutingForDSRFunc: method is nil but LinuxNetworking.setupPolicyRoutingForDSR was just called")
	}
	callInfo := struct {
	}{}
	lockLinuxNetworkingMocksetupPolicyRoutingForDSR.Lock()
	mock.calls.setupPolicyRoutingForDSR = append(mock.calls.setupPolicyRoutingForDSR, callInfo)
	lockLinuxNetworkingMocksetupPolicyRoutingForDSR.Unlock()
	return mock.setupPolicyRoutingForDSRFunc()
}

// setupPolicyRoutingForDSRCalls gets all the calls that were made to setupPolicyRoutingForDSR.
// Check the length with:
//     len(mockedLinuxNetworking.setupPolicyRoutingForDSRCalls())
func (mock *LinuxNetworkingMock) setupPolicyRoutingForDSRCalls() []struct {
} {
	var calls []struct {
	}
	lockLinuxNetworkingMocksetupPolicyRoutingForDSR.RLock()
	calls = mock.calls.setupPolicyRoutingForDSR
	lockLinuxNetworkingMocksetupPolicyRoutingForDSR.RUnlock()
	return calls
}

// setupRoutesForExternalIPForDSR calls setupRoutesForExternalIPForDSRFunc.
func (mock *LinuxNetworkingMock) setupRoutesForExternalIPForDSR(in1 serviceInfoMap) error {
	if mock.setupRoutesForExternalIPForDSRFunc == nil {
		panic("LinuxNetworkingMock.setupRoutesForExternalIPForDSRFunc: method is nil but LinuxNetworking.setupRoutesForExternalIPForDSR was just called")
	}
	callInfo := struct {
		In1 serviceInfoMap
	}{
		In1: in1,
	}
	lockLinuxNetworkingMocksetupRoutesForExternalIPForDSR.Lock()
	mock.calls.setupRoutesForExternalIPForDSR = append(mock.calls.setupRoutesForExternalIPForDSR, callInfo)
	lockLinuxNetworkingMocksetupRoutesForExternalIPForDSR.Unlock()
	return mock.setupRoutesForExternalIPForDSRFunc(in1)
}

// setupRoutesForExternalIPForDSRCalls gets all the calls that were made to setupRoutesForExternalIPForDSR.
// Check the length with:
//     len(mockedLinuxNetworking.setupRoutesForExternalIPForDSRCalls())
func (mock *LinuxNetworkingMock) setupRoutesForExternalIPForDSRCalls() []struct {
	In1 serviceInfoMap
} {
	var calls []struct {
		In1 serviceInfoMap
	}
	lockLinuxNetworkingMocksetupRoutesForExternalIPForDSR.RLock()
	calls = mock.calls.setupRoutesForExternalIPForDSR
	lockLinuxNetworkingMocksetupRoutesForExternalIPForDSR.RUnlock()
	return calls
}
