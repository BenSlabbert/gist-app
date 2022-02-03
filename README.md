# Gist Client

This is a [Cobra](https://github.com/spf13/cobra) cli app which optionally serves a [Svelte](https://svelte.dev/) web app.

This app uses a [Github](https://github.com/) access token to do simple CRUD on private [gists](https://gist.github.com/).

With no access token, only public gists are available.

create token: https://github.com/settings/tokens/new

## Testing

Make sure to set environment variable `GITHUB_API_TOKEN` for testing.
