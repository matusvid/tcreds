# tcreds

`tcreds` is a simple command-line tool written in Go to manage Terraform credentials. It allows you to create, update, switch, delete, and list Terraform credentials in an easy and organized way.

## Features

- **Use**: Switch to a specific set of Terraform credentials.
- **Create**: Generate and store new Terraform credentials.
- **Update**: Refresh existing credentials.
- **Delete**: Remove unwanted credentials.
- **List**: Display all stored credentials with their creation date.

## Installation

## Using Homebrew

You can install `tcreds` using the Homebrew package manager:

1. Tap the repository:

```bash
brew tap matusvid/tcreds
```

2. Install tcreds:

```bash
brew install tcreds
```

## Manual Installation
Alternatively, you can download the binary directly from the releases page and place it in your $PATH.

1. Download the latest version for your platform.
2. Make the binary executable:

```bash
chmod +x tcreds
```

3. Move it to a directory in your $PATH (e.g., /usr/local/bin):

```bash
mv tcreds /usr/local/bin/
```

## Usage

### Use Credentials:

```bash
tcreds use <credentials_name>
```
Switch to the specified Terraform credentials.

### Create Credentials:

```bash
tcreds create <credentials_name>
```
Generate new Terraform credentials and store them under the provided name.

### Update Credentials:

```bash
tcreds update <credentials_name>
```
Refresh the existing Terraform credentials for the specified name.

### Delete Credentials:

```bash
tcreds delete <credentials_name>
```
Delete the Terraform credentials associated with the specified name.

### List Credentials:

```bash
tcreds list
```
List all stored Terraform credentials with their creation date.

### Help:

```bash
tcreds -h
```
Display the help menu.

### Updating
To update tcreds to the latest version, run:

```bash
brew upgrade tcreds
```

## Contributing
Contributions are welcome! If you have suggestions for improvements or want to report bugs, please open an issue or submit a pull request.

Clone the repository:

```bash
git clone https://github.com/matusvid/tcreds.git
```
Make your changes in a new branch:

```bash
git checkout -b feature/new-feature
```

Open a pull request on GitHub.
