
##### xmansay
a poor duplicate of xcowsay.
it just display a poor **being** with a bubble of whatever you told it to say :)

## Usage: 

first [install qt bindings](https://github.com/therecipe/qt/wiki/Installation) completely. (you can skip the build process and deploy your application with therecipe/qt [docker containers](https://github.com/therecipe/qt/wiki/Deploying-Application))

`go get github.com/smf8/xmansay` 

**don't worry, because of [qt binding for go](https://github.com/therecipe/qt) it might take sometime to download project**

for building desired output for your enviroment use `qtdeploy build [linux/windwos/etc]`
don't forget to copy `res` folder to the generated `deploy` folder

#### If you want to use it with cron:

- create a link to original executable file inside one of `PATH` addresses.

- add these lines to your `~/.profile` to save `$DISPLAY` and `$XAUTHORITY` on each login. [ref](https://unix.stackexchange.com/a/10126)

  ```shell
  case $DISPLAY in                                                                                                                                                                   
  :*) export | grep -E '(^| )(DISPLAY|XAUTHORITY)=' | cut -d = -f 2 > ~/.local-display-setup.sh;;                                                                                                              
  esac  
  ```

- (Optional) create a file of sentences you want `xmansay` to say.

- add this to your crontab

  ```shell
  */10 * * * * export XAUTHORITY=$(tail -n 1 ~/.local-display-setup.sh) && export DISPLAY=$(head -n 1 ~/.local-display-setup.sh) && shuf -n 1 /path/to/sentences/list/sentences.txt | xmansay -time 10
  
  ```

------

**Demo:**
<img src="./demo.gif" width="100%">

as there is no official API for rendering persian texts in go [#27281](https://github.com/golang/go/issues/27281), multiline farsi texts are not rendered correctly in this project.


farsi text available because of [goarabic](https://github.com/01walid/goarabic)
