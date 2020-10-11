# toc-md

Generate or inject "Table of Contents" for markdown file.

## Contents

- [toc-md](#toc-md)
  - [Installation](#installation)
    - [Install toc-md](#install-toc-md)
    - [Verify installation](#verify-installation)
  - [Usage](#usage)

## Installation

### Install toc-md

1. Download the zip from [releases](https://github.com/yankeexe/toc-md/releases)

2. unzip the tar file

   ```bash
   $ tar -xzf toc
   ```

3. Copy the binary to `bin` directory.

   ```bash
   $ sudo cp toc /usr/local/bin
   ```

### Verify installation

Run `toc` to see the help message.

```bash
$ toc
```

## Usage

Generate table of contents

```bash
$ toc gen "README.md"
```

Inject table of contents to the file.

```bash
$ toc gen "README.md" -i
```
