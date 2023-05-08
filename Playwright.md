# playwright issues



To run the playwright in a Linux/ Ubuntu 22.04 (see `lsb_release -a`), you'll need to add some
configs in the APT, since the playwright script used to install the browsers in the current godog
version is outdated, throwing [this error](https://github.com/microsoft/playwright/issues/13738).

- [Adding an old repo that works](https://askubuntu.com/a/659345)
- [To fix the signature error](https://chrisjean.com/fix-apt-get-update-the-following-signatures-couldnt-be-verified-because-the-public-key-is-not-available/)
- [Fix apt warning related to multi files gpg](https://itsfoss.com/key-is-stored-in-legacy-trusted-gpg/)


Create a new dep update file

```bash
sudo vim /etc/apt/sources.list.d/playwright.list
```

Paste this into it

<pre>
# Focal 20.04.6 LTS
deb http://archive.ubuntu.com/ubuntu focal main restricted universe
deb http://archive.ubuntu.com/ubuntu focal-updates main restricted universe
deb http://security.ubuntu.com/ubuntu focal-security main restricted universe multiverse
</pre>

* If you install an older version, you'll need to do the others steps as well,

Update the package list and try to install playwright again

```bash
sudo apt update
```

These changes will add an old repository to your linux, be caferull with this change

