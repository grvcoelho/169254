# 169254
:pager: A simple API that simulate AWS EC2's metadata endpoint

## Why

Every EC2 Instance from AWS have access to a special metadata service that is used to retrieve metadata information about the instance such as `ami-id` and `local-ipv4`. The metadata service is tipically accessible on the private endpoint http://169.254.169.254 and cannot be accessed outside of the instance.

Whether you're writing some tests or building an application that data given back by the service metadata, you'll need to run it inside the instance or use some kind of mock API. 

This project is a mock API deployed at https://169254.now.sh that you can use to assert the service metadata response structure and possible fields. **Note:** all data returned by https://169254.now.sh is hard-coded, so the results will be different when you access the real service metadata within your EC2 Instances.

## Usage

The following are the current implemented enpdoints of the metadata service. Feel free to open a Pull Request to add any new endpoints:


* https://169254.now.sh/latest/meta-data/ami-id
* https://169254.now.sh/latest/meta-data/ami-launch-index
* https://169254.now.sh/latest/meta-data/ami-manifest-path
* https://169254.now.sh/latest/meta-data/block-device-mapping/ami
* https://169254.now.sh/latest/meta-data/block-device-mapping/root
* https://169254.now.sh/latest/meta-data/instance-id
* https://169254.now.sh/latest/meta-data/instance-type
* https://169254.now.sh/latest/meta-data/placement/availability-zone
* https://169254.now.sh/latest/meta-data/profile
* https://169254.now.sh/latest/meta-data/hostname
* https://169254.now.sh/latest/meta-data/local-hostname
* https://169254.now.sh/latest/meta-data/local-ipv4
* https://169254.now.sh/latest/meta-data/public-hostname
* https://169254.now.sh/latest/meta-data/public-ipv4
* https://169254.now.sh/latest/meta-data/reservation-id
* https://169254.now.sh/latest/meta-data/security-groups
* https://169254.now.sh/latest/meta-data/services/domain
* https://169254.now.sh/latest/meta-data/services/partition
