# .bashrc

# Source global definitions
if [ -f /etc/bashrc ]; then
	. /etc/bashrc
fi

# Uncomment the following line if you don't like systemctl's auto-paging feature:
# export SYSTEMD_PAGER=

# User specific aliases and functions
export PS1="[\e[94m\u\033[0m@\e[92m\h\e[0m \e[33m\t\033[0m \e[96m\W\033[0m]$ "
export GOPATH=~/work/git/go
export GOROOT=/usr/lib/golang
export PATH=$PATH:$GOPATH/bin:$GOROOT/bin

