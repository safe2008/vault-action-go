# vault-actions

[![GitHub Marketplace](https://img.shields.io/badge/Marketplace-Find%20and%20Replace-blue.svg?colorA=24292e&colorB=0366d6&style=flat&longCache=true&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAA4AAAAOCAYAAAAfSC3RAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAM6wAADOsB5dZE0gAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAERSURBVCiRhZG/SsMxFEZPfsVJ61jbxaF0cRQRcRJ9hlYn30IHN/+9iquDCOIsblIrOjqKgy5aKoJQj4O3EEtbPwhJbr6Te28CmdSKeqzeqr0YbfVIrTBKakvtOl5dtTkK+v4HfA9PEyBFCY9AGVgCBLaBp1jPAyfAJ/AAdIEG0dNAiyP7+K1qIfMdonZic6+WJoBJvQlvuwDqcXadUuqPA1NKAlexbRTAIMvMOCjTbMwl1LtI/6KWJ5Q6rT6Ht1MA58AX8Apcqqt5r2qhrgAXQC3CZ6i1+KMd9TRu3MvA3aH/fFPnBodb6oe6HM8+lYHrGdRXW8M9bMZtPXUji69lmf5Cmamq7quNLFZXD9Rq7v0Bpc1o/tp0fisAAAAASUVORK5CYII=)](https://github.com/marketplace/actions/vault-actions)
[![Actions Status](https://github.com/safe2008/vault-actions/actions/workflows/test.yml/badge.svg)](https://github.com/safe2008/vault-actions/actions/workflows/test.yml)
[![Actions Status](https://github.com/safe2008/vault-actions/actions/workflows/integration.yml/badge.svg)](https://github.com/safe2008/vault-actions/actions/workflows/integration.yml)

This action will sync Vault application.

## Usage

### Example workflow

This example replaces syncs Vault application.

```yaml
name: My Workflow
on: [ push, pull_request ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Get vault secret
        uses: safe2008/vault-actions@main
        with:
          address: "https://vault.example.com"
          token: ${{ secrets.ARGOCD_TOKEN }}
          secrets: "argocd/data/token admin | ARGOCD_TOKEN"
```

### Inputs

| Input | Description|
| --- | --- |
| `address` | Vault server address. |
| `token` | Vault Token. |
| `secrets` | path of the secret in vault and set output variable bame. |

## Examples

### Get Vautl secret

You can get Vault application after building an image etc.

```yaml
name: My Workflow
on: [ push, pull_request ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@master

      - name: Import Vault Secrets
        id: secrets
        uses: safe2008/vault-actions@main
        with:
          address: ${{ secrets.VAULT_ADDR }}
          token: ${{ secrets.VAULT_TOKEN }}
          secrets: "argocd/data/token admin | ARGOCD_TOKEN"

      - name: Self test
        id: selftest
        uses: safe2008/argocd-app-actions@main
        with:
          address: ${{ secrets.ARGOCD_ADDR }}
          token: ${{ steps.secrets.outputs.ARGOCD_TOKEN }}
          insecure: false
          appName: horusec-platform
```

## Publishing

To publish a new version of this Action we need to update the Docker image tag in `action.yml` and also create a new
release on GitHub.

- Work out the next tag version number.
- Update the Docker image in `action.yml`.
- Create a new release on GitHub with the same tag.
