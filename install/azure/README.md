# Gitpod on Azure

This is a simple proof-of-concept which uses the aks-preview using containerd as runtime.

## Prerequisites

To install Gitpod on Azure an account has to be configured. https://azure.microsoft.com/en-us/

The Azure client is used to run the necessary commands. https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest

A simple `Makefile` is used to document all steps needed to create a gitpod environment. To run the commands in these
steps successfully, the AKS-Preview has to be enabled.

https://docs.microsoft.com/en-us/azure/aks/cluster-configuration#container-runtime-configuration-preview

## Configuration

Configure the variables in the `Makefile` according to your Azure account.

## Installation

If the command `make install` is run, several commands will be executed based on the previous configuration. These command could also be run step-by-step or directly on the command line.

Gitpod is installed using the Helm chart in the root folder. The `values.yaml` is generated during the process.

## Outlook

This method is very simple ans has no claims to be expanded or used as a basis for futher development. A support for Azure using Terraform is planned and will be released when its support of containerd is officially released.