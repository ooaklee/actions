# Ooaklee Actions

This monorepo houses all my actions.

A release is done by pushing to version branches, such as `v1` based on the `base` branch for new versions. The `main` branch is the development branch and should not be used by consumers.

## Actions
- [go-example](./go-example)

## Development
To start developing in this repo, you must enable asdf and `node corepack`.

Install the necessary ASDF plugins and install the required versions.


```sh
# Install asdf plugins
$ asdf plugin-add golang
$ asdf plugin-add nodejs
$ asdf plugin-add yarn

# Install necessary versions
$ asdf install
````

Complete the setup of the repo.

```sh
$ corepack enable
$ asdf reshim nodejs
$ yarn
yarn install v1.22.22
[1/4] 🔍  Resolving packages...
[2/4] 🚚  Fetching packages...
[3/4] 🔗  Linking dependencies...
[4/4] 🔨  Building fresh packages...
\$ husky
✨ Done in 0.16s.
```


<!-- CONTRIBUTING -->
## Contributing

Contributions make the open-source community a fantastic place to learn, inspire, and create. Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch 
3. Commit your Changes 
4. Push to the Branch 
5. Ensure you test code added
6. Open a Pull Request


> **DO NOT** make any changes to binaries in the `dist/` as Husky automatically generates it (using yarn package) on pre-commit.

> Any changes made on the GitHub Action source code MUST be reflected in the binaries located in the `dist` directory, too, as the workflow runs the `invoke-binary.js`, which will invoke the binary based on the respective runner platform and architecture, NOT the `src/main.go` of the respective action.

> **Run `yarn package`** before you push any changes made on the source code of the GitHub Action `src/main.go` 

## Support

Any support is appreciated :raised_hands:! You can show your support by staring this project. If you wish to show support in any other way, don't hesitate to contact me via the email below.

<!-- LICENSE -->
## License

Distributed under the MIT License. See [`LICENSE`](./LICENSE) for more information.



<!-- CONTACT -->
## Contact

Leon Silcott - leon@boasi.io
