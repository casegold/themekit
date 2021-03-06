---
layout: default
---
# Shopify Theme Kit

Theme Kit is a single binary that has no dependencies. Download the application
and with a tiny bit of setup you’re off to the theme creation races.

Using Theme Kit will enable you to

* Upload Themes to Multiple Environments
* Fast Uploads and Downloads
* Watch for local changes and upload automatically to Shopify
* Works on Windows, Linux and macOS

## Installation

### Automatic Installation

If you are on Mac or Linux you can use the following installation script to automatically
download and install Theme Kit for you. Please follow the directions outputted to your
console to change your bash profile so that you will have access to the `theme` command

```bash
curl -s https://raw.githubusercontent.com/Shopify/themekit/master/scripts/install | sudo python
```

### Homebrew

If you have [homebrew](http://brew.sh/) installed you can install Theme Kit by running the following commands.

```bash
brew tap shopify/shopify
brew install themekit
```

### Windows Installer

Download and run the installer:

- [windows-64](https://github.com/Shopify/themekit/releases/download/{{ site.themekitversion }}/themekit-setup-64.exe)
- [windows-32](https://github.com/Shopify/themekit/releases/download/{{ site.themekitversion }}/themekit-setup-32.exe)

### Manual Installation

Download and unzip the latest release.

| OS     | Architecture |          |
| :------| :------------| :------- |
| macOS  | 64-bit       |  [download](https://github.com/Shopify/themekit/releases/download/{{ site.themekitversion }}/darwin-amd64.zip)
| Windows| 64-bit       |  [download](https://github.com/Shopify/themekit/releases/download/{{ site.themekitversion }}/windows-amd64.zip)
| Windows| 32-bit       |  [download](https://github.com/Shopify/themekit/releases/download/{{ site.themekitversion }}/windows-386.zip)
| Linux  | 64-bit       |  [download](https://github.com/Shopify/themekit/releases/download/{{ site.themekitversion }}/linux-amd64.zip)
| Linux  | 32-bit       |  [download](https://github.com/Shopify/themekit/releases/download/{{ site.themekitversion }}/linux-386.zip)

## Get API Access

You will need to set up an API key to add to our configuration and create a connection
between your store and Theme Kit. The API key allows Theme Kit to talk to and access
your store, as well as its theme files.

To do so, log into the Shopify store, and create a private app. In the Shopify
Admin, go to Apps and click on View private apps. From there, click Generate API
credentials to create your private app. Make sure to set the permissions of **Theme
templates and theme assets** to have ***Read and write*** access in order to generate the
appropriate API credentials, then click Save.

<img src="{{ "/assets/images/shopify-local-theme-development-generate-api.gif" | prepend: site.baseurl }}" />

Fill out the information at the top and set the permissions of `Theme templates and theme assets` to
read and write access. Press `Save` and you will be presented with the next screen. In it you will
see your access credentials. Please make note of the password. You will need it later.

<img src="{{ "/assets/images/private-app-password.png" | prepend: site.baseurl }}" />

## Use a new theme.

If you are starting from scratch and want to get a quick start, run the following:

```bash
theme bootstrap --password=[your-password] --store=[your-store.myshopify.com]
```

This will create a new theme for your online store from the [Timber](https://shopify.github.io/Timber/) template. Then
it will download all those assets from Shopify and automatically create a `config.yml` file for you.

## Configure an existing theme.

To connect an existing theme, you need the theme’s ID number. The easiest way to
get your theme’s ID number is to go to the Theme Editor click on Edit HTML/CSS and
copy the theme ID number from the URL — it will be last several digits after mystore.myshopify.com/admin/themes/.

<img src="{{ "/assets/images/shopify-local-theme-development-theme-id.gif" | prepend: site.baseurl }}" />

Then once you have noted your theme ID, run the following commands:

```bash
# create configuration
theme configure --password=[your-password] --store=[you-store.myshopify.com] --themeid=[your-theme-id]
# download and setup project in the current directory
theme download
```
