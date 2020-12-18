---
url: /docs/self-hosted/latest/install/oauth/
---

# How To integrate Gitpod with OAuth providers

Gitpod does not implement user authentication itself, but integrates with other auth provider using [OAuth2](https://oauth.net/2/).
Usually your Git hosting solution (e.g. GitHub or GitLab) acts as the OAuth auth provider. This way we control access to Gitpod while at
the same time making sure every user has proper access to their Git repository.
You can add multiple providers.

Gitpod supports the following authentication providers:
* [GitHub](https://github.com)
* [GitHub Enterprise](https://github.com/enterprise) in version 2.16.x and higher
* [GitLab](https://gitlab.com)
* [GitLab Community Edition](https://about.gitlab.com/community/) in version 11.7.x and higher
* [GitLab Enterprise Edition](https://about.gitlab.com/enterprise/) in version 11.7.x and higher
* [Bitbucket](https://bitbucket.org) (coming soon)
* Custom OAuth Providers (please send an email to [contact@gitpod.io](mailto:contact@gitpod.io) for more information)

If no OAuth provider is set up using the `values.yaml`, Gitpod will ask for the configuration in the dashboard.

## GitHub
To authenticate your users with GitHub you need to create a [GitHub OAuth App](https://developer.github.com/apps/building-oauth-apps/creating-an-oauth-app/).

- set "Authentication callback URL" to: 
    ```     
    https://<your-domain.com>/auth/github/callback
    ```
 - copy the following values and configure them in `values.yaml`:
    - set the "Client ID" in `oauth.clientId` and `oauth.settingsUrl`
    - set the "Client Secret" in `oauth.clientSecret`
    ```
    authProviders:
      - id: "Example Github"
        host: "github.com"
        protocol: "https"
        type: "GitHub"
        oauth:
            clientId: "<clientId>"
            clientSecret: "<clientSecret>"
            callBackUrl: "https://your-domain.com/auth/github/callback"
            settingsUrl: "https://github.com/settings/connections/applications/<clientId>"
    ```

## GitLab
To authenticate your users with GitLab you need to create an [GitLab OAuth application](https://docs.gitlab.com/ee/integration/oauth_provider.html).
Follow the guide linked above and:
- set "Authentication callback URL" to: 
    ```
    https://<your-domain.com>/auth/<gitlab.com-OR-your-gitlab.com>/callback
    ```
- set "Scopes" to `api`, `read_user` and `read_repository`.
- copy the following values and configure them in `values.yaml`:
    - set the "Application ID" in `oauth.clientId`
    - set the "Secret" in `oauth.clientSecret`
    - set the "Authentication callback URL" of your GitLab installation in `oauth.callBackUrl`
    - set your the Gitlan domain in `oauth.settingsUrl`

    ```
    authProviders:
      - id: "Example Gitlab"
        host: "gitlab.com"
        protocol: "https"
        type: "GitLab"
        oauth:
            clientId: "<ApplicationID>"
            clientSecret: "<Secret>"
            callBackUrl: "https://<your-domain.com>/auth/<gitlab.com-OR-your-gitlab.com>/callback"
            settingsUrl: "<gitlab.com-OR-your-gitlab.com>/profile/applications"
    ```
