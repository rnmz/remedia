<div align="center">
  <img src="logo.png" alt="ReMedia" width="360">

  <h1>ReMedia</h1>

  <p>A cross-platform reader for comics, manga and webtoons<br>with progress synced across all your devices.</p>

  <img src="https://img.shields.io/badge/status-work_in_progress-orange" alt="Status: work in progress">
</div>

---

## About

ReMedia is an ecosystem of a backend and client apps for reading comics, manga and webtoons. One account, one library, and your reading progress follows you from phone to desktop.

The project is under active development. Right now the focus is on the **backend**; screenshots, public builds and release notes will appear closer to the first release.

## Features

**Reading**
- Page caching in RAM and on disk, so pages open instantly and don't re-download traffic
- Three reading modes: classic comic pages, right-to-left manga, vertical webtoon scroll
- Offline downloads of chapters *(planned)*

**Catalog**
- Home feed with news, new releases and fresh chapters
- Search by name, format, release status and tags
- Title pages with synopsis, cover, chapter list and your progress

**Account**
- Cloud sync of reading progress
- Custom library folders such as *Fav. isekai*, *Too boring*, *Plan to read*
- Subscription management
- 2 factor authentication *(planned)*

## Platforms

| Platform | Status | Notes |
| :--- | :--- | :--- |
| Backend | In progress | REST API, authentication, content delivery, OpenAPI docs |
| Android | In progress | Jetpack Compose, will be migrated to Flutter |
| iOS, Desktop, Web | Planned | Flutter multiplatform |

## Repository structure

```text
ReMedia-Backend/   server application
ReMedia/           client application (Android, later multiplatform)
```

## Getting started

> Setup instructions will be added once the first stable version is ready.

<!-- TODO: backend requirements and run command -->
<!-- TODO: Android build instructions -->

## Contributing

The project is not yet open for code contributions, but ideas and bug reports are welcome. Please open an issue to discuss a change before sending a pull request.

## License

**[MIT License](./license.md)**

Copyright (c) 2026-present, ReMedia project
