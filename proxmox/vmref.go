package proxmox

import "errors"

func NewVmRef(vmId int) (vmr *VmRef) {
	vmr = &VmRef{vmId: vmId, node: "", vmType: ""}
	return
}

// VmRef - virtual machine ref parts
// map[type:qemu node:proxmox1-xx id:qemu/132 diskread:5.57424738e+08 disk:0 netin:5.9297450593e+10 mem:3.3235968e+09 uptime:1.4567097e+07 vmid:132 template:0 maxcpu:2 netout:6.053310416e+09 maxdisk:3.4359738368e+10 maxmem:8.592031744e+09 diskwrite:1.49663619584e+12 status:running cpu:0.00386980694947209 name:appt-app1-dev.xxx.xx]
type VmRef struct {
	vmId    int
	node    string
	pool    string
	vmType  string
	haState string
	haGroup string
}

const (
	VmRef_Error_Nil string = "vm reference may not be nil"
)

func (vmr *VmRef) Check(c *Client) (err error) {
	err = vmr.nilCheck()
	if err != nil {
		return
	}
	if vmr.node == "" || vmr.vmType == "" {
		_, err = c.GetVmInfo(vmr)
	}
	return
}

func (vmr *VmRef) GetVmType() string {
	return vmr.vmType
}

func (vmr *VmRef) HaGroup() string {
	return vmr.haGroup
}

func (vmr *VmRef) HaState() string {
	return vmr.haState
}

func (vmr *VmRef) nilCheck() error {
	if vmr == nil {
		return errors.New(VmRef_Error_Nil)
	}
	return nil
}

func (vmr *VmRef) Node() string {
	return vmr.node
}

func (vmr *VmRef) Pool() string {
	return vmr.pool
}

func (vmr *VmRef) SetNode(node string) {
	vmr.node = node
}

func (vmr *VmRef) SetPool(pool string) {
	vmr.pool = pool
}

func (vmr *VmRef) SetVmType(vmType string) {
	vmr.vmType = vmType
}

func (vmr *VmRef) VmId() int {
	return vmr.vmId
}
