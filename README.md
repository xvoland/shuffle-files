# SHUFFLE-FILES
This is a CLI tool which shuffles the files in the directory, their content, but without changing the file names.

*NOTE: If the path contains subdirectories, this has no effect on them. Also, the application ignores all files that begin with a dot "."*

<p align="right"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/qr_shuffle-files.png" alt="DOTOCA Ltd." height="50" width="50" /></a>
</p>

## Install

### 🍺 Homebrew.  Let’s try it!

```bash
brew tap xvoland/shuffle-files

brew install shuffle-files
```
or
```bash
brew install xvoland/shuffle-files
```

<br />



### 🐙 GitHub Releases

Alternatively, binaries are available in the [GitHub Releases][githubreleases]. Or you can visit the [home page][homepage]

<br />
<br />

# ℹ️ How it use

#### Path
```bash
shuffle-files ./some_path_to_files
```

*Shuffles the files in the specified path `./some_path_to_files`*
*For example, the directory contains:*

	- file1 (content1)
	- file2 (content2)
	- file3 (content3)

*After running the `shuffle_file`, now has content like this:*

	- file1 (content2)
	- file2 (content1)
	- file3 (content3)


#### -o (output)

```bash
shuffle-files ./some_path_to_files -o /output_path
```

*Shuffles the files in the specified path `./some_path_to_files` and copy output result to `/output_path`*



#### --debug
```bash
shuffle-files ./some_path_to_files --debug
```

*Shuffles the files in the specified path `./some_path_to_files` and display the result of the program on the screen*



#### --test
```bash
shuffle-files ./some_path_to_files --test
```

*Shuffles the files in the specified path `./some_path_to_files` and do nothing with the files*



#### Combination of flags
```bash
shuffle-files ./some_path_to_files --test --debug
```

*You have the ability to combine keys.
For example, shuffles the files in the specified path `./some_path_to_files` and do nothing with the files, and display the result of the program on the screen*



#### --help
```bash
shuffle-files --help
```

*Show help on the screen*



#### --version
```bash
shuffle-files --version
```

*Show program version*

<br />
<br />

# ⚠️ Donation

Whether I receive the money or not, I’ll keep improving the app because I genuinely love seeing people use it to achieve their goals. Every single dollar truly makes a difference for creators like me. It helps cover things like domain hosting and the countless hours I spend coding. Your support would mean the world to me and even give me a bit more time to spend with my family. Thank you so much!

### Crypto

**BTC (ERC20):** 0x17496b75d241d377334717f8cbc16cc1a5b80396<br />
**USDT (TRC20):** TAAsGXjNoQRJ7ewxSBL2W3DUCoG7h8LCT6


### Other

<p align="center">
  <a href="https://paypal.me/xvoland" target="blank"><img align="center" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/paypal.png" alt="PayPal" width="250" /></a>
</p>

<br />


# ☎️ Connect with me:

### Social
[<img align="left" alt="xVoLAnD" width="50px" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/logo-dotoca.svg" />][home]
[<img align="left" alt="xvoland | Instagram" width="50px" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/instagram.svg" />][instagram]
[<img align="left" alt="Vitalii Tereshchuk | YouTube" width="50px" src="https://raw.githubusercontent.com/xvoland/xvoland/main/images/youtube.svg" />][youtube]

<br />
<br />

# 📺 My Latest YouTube Videos:
<!-- YOUTUBE:START -->
- [99% of People Don’t Know This Shower Gel Hack](https://www.youtube.com/watch?v=aWfExS_jswA)
- [JSX Plugin v0.6.9 Nano Banana 2 Pro Gemini 3.1 - jsxNanaBanana Plugin for Adobe Photoshop](https://www.youtube.com/watch?v=rher71lpwcw)
- [​                                            𝓍𝑽𝕠𝑳𝑨𝔫𝔇](https://www.youtube.com/watch?v=HhTHIrUPJcc)
- [⚡ Simple DIY Repair Old Bike Battery And Install Separate USB Charging Port For Any Device](https://www.youtube.com/watch?v=dNi7ayBp2nk)
- [🔴 WWDC App&#39;s Wall Screensaver | Apple Store Apps Wall  #live #screensaver4k #relax](https://www.youtube.com/watch?v=tZ3UaYibMso)
<!-- YOUTUBE:END -->

➡️ [more videos...][youtube]
<br />
<br />

## ⛔ License

&copy; 2023, [Vitalii Tereshchuk][home] via Apache2.0 license.


[home]: http://dotoca.net
[homepage]: https://dotoca.net/shuffle-files
[githubreleases]: https://github.com/xvoland/shuffle-files/releases
[paypal]: https://paypal.me/xvoland
[youtube]: https://youtube.com/xvoland
[instagram]: https://www.instagram.com/xvoland/
[opencollective]: https://opencollective.com/extract/backers/0/website