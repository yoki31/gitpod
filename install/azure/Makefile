#
# Azure
#
AZ_SUBSCRIPTION_ID=AZ_SUBSCRIPTION_ID
AZ_LOCATION=AZ_LOCATION
AZ_RESOURCE_GROUP=AZ_RESOURCE_GROUP
AZ_DNS_TENANT_ID=AZ_DNS_TENANT_ID
AZ_APP_NAME=AZ_APP_NAME

#
# Gitpod
#
GITPOD_HOSTNAME=your-domain.com

#
# GitHub
#
GITHUB_CLIENT_ID=GITHUB_CLIENT_ID
GITHUB_CLIENT_SECRET=GITHUB_CLIENT_SECRET

.PHONY: subscription
subscription: 
	az account set \
		--subscription $(AZ_SUBSCRIPTION_ID)

.PHONY: resource-group
resource-group: subscription
	az group create \
		--location $(AZ_LOCATION) \
        --name $(AZ_RESOURCE_GROUP) \
        --subscription $(AZ_SUBSCRIPTION_ID)

vnet: resource-group
	az network vnet create \
		--name $(AZ_APP_NAME)-network \
        --resource-group $(AZ_RESOURCE_GROUP) \
        --address-prefixes 10.0.0.0/8 \
		--subnet-name $(AZ_APP_NAME)-subnet \
		--subnet-prefix 10.240.0.0/16 \
        --location $(AZ_LOCATION) \
    	--subscription $(AZ_SUBSCRIPTION_ID) \

AZ_SUBNET_ID=$(shell az network vnet show --resource-group $(AZ_RESOURCE_GROUP) --name $(AZ_APP_NAME)-network | jq -r '.subnets[0].id')
AZ_SUBNET_NAME=$(shell az network vnet show --resource-group $(AZ_RESOURCE_GROUP) --name $(AZ_APP_NAME)-network | jq -r '.subnets[0].name')

aks: resource-group vnet
	az aks create \
		--name $(AZ_APP_NAME)-cluster \
		--resource-group $(AZ_RESOURCE_GROUP) \
		--aks-custom-headers CustomizedUbuntu=aks-ubuntu-1804,ContainerRuntime=containerd \
		--vnet-subnet-id $(AZ_SUBNET_ID) \
		--network-plugin azure \
		--network-policy calico

.PHONY: aks-credentials
aks-credentials: aks
	az aks get-credentials \
		--name $(AZ_APP_NAME)-cluster \
        --resource-group $(AZ_RESOURCE_GROUP) \
		--subscription $(AZ_SUBSCRIPTION_ID)

.PHONY: values
values:
	cp $(PWD)/templates/values.yaml $(PWD)/values.yaml
	sed -i -e 's/GITPOD_HOSTNAME/$(GITPOD_HOSTNAME)/g' values.yaml
	sed -i -e 's/GITHUB_CLIENT_ID/$(GITHUB_CLIENT_ID)/g' values.yaml
	sed -i -e 's/GITHUB_CLIENT_SECRET/$(GITHUB_CLIENT_SECRET)/g' values.yaml

GITPOD_IP=$(shell kubectl get service proxy -o json | jq -r '.status.loadBalancer.ingress[0].ip')
.PHONY: gitpod
gitpod: values aks
	helm dependency update $(PWD)/../../chart
	helm upgrade --install gitpod $(PWD)/../../chart \
		--values $(PWD)/values.yaml \
		--timeout 10m0s
	@echo $(GITPOD_IP)


.PHONY: install
install: gitpod

.PHONY: uninstall
uninstall:
	az group delete \
		--name $(AZ_RESOURCE_GROUP) \
		--subscription $(AZ_SUBSCRIPTION_ID) \
		--yes
