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

No matter if I get the money or not, I'm gonna keep making the app better 'cause I love seeing folks use it and reach their goals.<br />
And you know what? Every $1 really makes a difference for folks like me.<br />
It helps cover stuff like domain hosting and the hours I put into coding, which would be super awesome and free up some more family time. Thanks a bunch!

### Crypto

**BTC (ERC20):** 0x17496b75d241d377334717f8cbc16cc1a5b80396<br />
**USDT (TRC20):** TAAsGXjNoQRJ7ewxSBL2W3DUCoG7h8LCT6


### Other

[Donate here for my projects][paypal]

<a href='https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=9D4YBRWH8QURU'><img alt='Click here to lend your support to my projects and make a donation!' src='https://www.paypalobjects.com/en_US/GB/i/btn/btn_donateCC_LG.gif' border='0' /></a>

<br />
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
- [👀 WONLINK 1200Mb WL-NE3501 wireless repeater 802.11bgnac 1200 Megabits](https://www.youtube.com/watch?v=FEk1HdZMX2g)
- [ AirDrop Missing on macOS Sequoia? Easy Fix for Finder and AirDrop Issues!](https://www.youtube.com/watch?v=Baj8RbF_cTA)
- [👀 Review UGREEN HiTune H6 Pro Earbuds 2024 Best Deal: Hi-Res Wireless Headphones ANC LDAC](https://www.youtube.com/watch?v=2NksaUJAU8I)
- [⚒️ How to Fix Drawer Slides with Ball Bearings – Easy DIY Guide. Lifetime guarantee](https://www.youtube.com/watch?v=VuP4gkVt_-0)
- [🔥 Lifehack Of Open a Sealed Tube. Did You Know?](https://www.youtube.com/watch?v=YlD9b6KnfdM)
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