#! /bin/bash

home_path=$HOME
gov_path="${home_path}/.gov"
bin_path="./.gov/bin"
version_path="./.gov/gos"

if [[ ! -e $gov_path ]];then
    mkdir $gov_path
fi

if [[ ! -e $bin_path ]];then
    mkdir $bin_path
fi

if [[ ! -e $version_path ]];then
    mkdir $version_path
fi

echo "export PATH="