# VCH Deployment Fails with Resource Pool Creation Error #
When you use the command line installer to deploy a virtual container host directly on an ESXi host, the installation fails with a resource pool creation error.

## Problem ##
Command line installation on an ESXi host fails during the validation of the configuration that you provided: 

<pre>
Creating resource pool failed with ServerFaultCode: 
Access to resource settings on the host is restricted to the server 
that is managing it: <i>vcenter_server_address</i>. 
Exiting ...
</pre>

## Cause ##
You set the `target` option to the address of an ESXi host that is managed by a vCenter Server instance.

## Solution ##
- Set the `target` option to the address of the vCenter Server instance that manages the ESXi host. 
- Disassociate the ESXi host from the vCenter Server instance and run the installer again. 
