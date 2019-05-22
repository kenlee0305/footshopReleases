# Footshop Releases Bot

Automated raffle script made to easily enter raffles on releases.footshop.com
If you don't have Go language installed follow [this](https://golang.org/doc/install) link.

## Installation
Tested on macOS & Ubuntu 18.04 Desktop

1. Clone this repo to your machine.
```
git clone https://github.com/washedsamo/footshopReleases-raffle-bot
```
2. Navigate to your cloned repository
```
cd $HOME/footshopReleases-raffle-bot
```
3. Edit your people.json file with your Billing details and save the file (⌘+S)

3.1. Edit your footshop.go file with your selected SizeID and ProductID (which could be found in the product url) and save the file (⌘+S)

4. Make the file executable by running this command
```
go build footshop.go
```
5. Run the script by typing the following command to your terminal.
```
./footshop
```
or
```
go run footshop.go
```
Enjoy!


## Contact me
*Feel free to DM me on Twitter or Discord if you need any help*.

* [Twitter](https://twitter.com/washedsamo)
* [Mail](mailto:iam@samuelmikula.com)

*DISCLAIMER*
*Your sensitive informations are stored locally on your machine.*

