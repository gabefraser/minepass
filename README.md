# MinePass - Minecraft Whitelister

**MinePass** is a powerful Minecraft whitelisting tool that allows you to manage your server's whitelist remotely over the internet. With MinePass, you can easily control access to your Minecraft server, add or remove players from the whitelist, and ensure a safe and secure gaming experience for your community.

## Features

- Easy-to-use web interface for managing your server's whitelist.
- Secure authentication to protect your server from unauthorized access.
- Real-time updates to the Minecraft whitelist without the need for server restarts.
- Compatibility with popular Minecraft server software.
- User-friendly design for both server administrators and players.

## Getting Started

### Running MinePass in a Docker Container

1. Pull the MinePass Docker image from Docker Hub
    ```shell
   docker pull gabefraser/minepass:latest
    ```
   
2. Run MinePass in a Docker container, specifying the necessary environment variables and ports. Replace the 
   placeholders with your configuration.
    ```shell
   docker run -d \
    -e MP_HOSTNAME=YOUR_SERVER_IP \ # required
    -e MP_PASSWORD=YOUR_SERVER_RCON_PASSWORD \ # required
    -e MP_UI_PASSWORD=YOUR_UI_PASSWORD \ # required
    -e MP_TITLE=YOUR_TITLE \ # optional, will default to "MinePass"
    -p 8080:8080 \
    gabefraser/minepass:latest
    ```

## Credits

MinePass is made possible thanks to the contributions and support from the following individuals and projects:

- [gin-gonic/gin](https://github.com/gin-gonic/gin): Gin is a web framework written in Go.
- [gorcon/rcon](https://github.com/gorcon/rcon): Source RCON Protocol implementation in Go.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Feedback and Support

If you encounter any issues or have suggestions for improving MinePass, please open an issue on the [GitHub repository](https://github.com/gabefraser/minepass/issues).

We hope you find MinePass to be a valuable tool for managing your Minecraft server's whitelist. Happy gaming!
